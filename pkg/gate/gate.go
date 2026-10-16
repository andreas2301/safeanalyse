package gate

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/user/safeanalyze/pkg/eval"
)

// Decisions.
const (
	Accept = "ACCEPT"
	Revert = "REVERT"
)

// Decision kinds.
const (
	KindGain       = "gain"       // significant, generalizing PRIMARY dev F1 gain
	KindLatency    = "latency"    // latency-only improvement, labeled counts identical
	KindRegression = "regression" // a rule was violated
	KindNoGain     = "no_gain"    // nothing regressed, nothing improved
)

const eps = 1e-12

// SetRun is one labeled set's eval result in one run directory.
type SetRun struct {
	Counts eval.Counts
	P95Ms  float64
	// Pred holds per-sample predictions (predictions-<set>.jsonl), or nil.
	Pred []eval.Prediction
	// FP and FN are sha256(text) multisets from fp-<set>.jsonl and
	// fn-<set>.jsonl, or nil. They are the fallback for pairing two runs.
	FP, FN map[string]int
}

// Run maps set names (e.g. "llmail", "llmail-holdout") to their results.
type Run map[string]*SetRun

// LoadRun reads every eval-<set>.json in dir, plus predictions-<set>.jsonl,
// fp-<set>.jsonl and fn-<set>.jsonl when present.
func LoadRun(dir string) (Run, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "eval-*.json"))
	if err != nil {
		return nil, err
	}
	if len(paths) == 0 {
		return nil, fmt.Errorf("no eval-<set>.json files in %s", dir)
	}
	run := Run{}
	for _, p := range paths {
		name := strings.TrimSuffix(strings.TrimPrefix(filepath.Base(p), "eval-"), ".json")
		data, err := os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", p, err)
		}
		var doc struct {
			eval.Counts
			LatencyP95Ms float64 `json:"latency_p95_ms"`
		}
		if err := json.Unmarshal(data, &doc); err != nil {
			return nil, fmt.Errorf("parsing %s: %w", p, err)
		}
		sr := &SetRun{Counts: doc.Counts, P95Ms: doc.LatencyP95Ms}
		if sr.Pred, err = loadPredictions(filepath.Join(dir, "predictions-"+name+".jsonl")); err != nil {
			return nil, err
		}
		if sr.FP, err = loadTextHashes(filepath.Join(dir, "fp-"+name+".jsonl")); err != nil {
			return nil, err
		}
		if sr.FN, err = loadTextHashes(filepath.Join(dir, "fn-"+name+".jsonl")); err != nil {
			return nil, err
		}
		run[name] = sr
	}
	return run, nil
}

func loadPredictions(path string) ([]eval.Prediction, error) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("opening %s: %w", path, err)
	}
	defer f.Close()
	preds := []eval.Prediction{}
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		if strings.TrimSpace(sc.Text()) == "" {
			continue
		}
		var p eval.Prediction
		if err := json.Unmarshal(sc.Bytes(), &p); err != nil {
			return nil, fmt.Errorf("parsing %s: %w", path, err)
		}
		preds = append(preds, p)
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	return preds, nil
}

func loadTextHashes(path string) (map[string]int, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, nil
	}
	samples, err := eval.LoadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	m := map[string]int{}
	for _, s := range samples {
		sum := sha256.Sum256([]byte(s.Text))
		m[hex.EncodeToString(sum[:])]++
	}
	return m, nil
}

// Bootstrap is the paired-bootstrap confidence interval of a dev F1 gain.
type Bootstrap struct {
	Source   string  `json:"source"` // "predictions" or "fp_fn"
	Observed float64 `json:"observed_gain"`
	Lower    float64 `json:"ci_lower"`
	Upper    float64 `json:"ci_upper"`
}

// SetDelta compares one set between baseline and candidate.
type SetDelta struct {
	Set       string      `json:"set"`
	Role      Role        `json:"role"`
	Holdout   bool        `json:"holdout"`
	Base      eval.Counts `json:"baseline"`
	Cand      eval.Counts `json:"candidate"`
	BaseF1    float64     `json:"baseline_f1"`
	CandF1    float64     `json:"candidate_f1"`
	BaseP     float64     `json:"baseline_precision"`
	CandP     float64     `json:"candidate_precision"`
	BaseR     float64     `json:"baseline_recall"`
	CandR     float64     `json:"candidate_recall"`
	BaseP95Ms float64     `json:"baseline_p95_ms"`
	CandP95Ms float64     `json:"candidate_p95_ms"`
	Status    string      `json:"status"`
	Bootstrap *Bootstrap  `json:"bootstrap,omitempty"`
}

// Decision is the gate's verdict.
type Decision struct {
	Decision string     `json:"decision"`
	Kind     string     `json:"kind"`
	Reasons  []string   `json:"reasons"`
	Notes    []string   `json:"notes,omitempty"`
	Sets     []SetDelta `json:"sets"`
	Ignored  []string   `json:"ignored,omitempty"`
}

func (d *Decision) reason(format string, a ...any) {
	d.Reasons = append(d.Reasons, fmt.Sprintf(format, a...))
}

func (d *Decision) note(format string, a ...any) {
	d.Notes = append(d.Notes, fmt.Sprintf(format, a...))
}

// Evaluate applies the decision rules of pol to a baseline and a candidate
// run. It returns an error when the runs cannot be compared (a gated set is
// missing, or the two runs evaluated different datasets).
func Evaluate(pol *Policy, base, cand Run) (*Decision, error) {
	d := &Decision{Reasons: []string{}}
	idx := map[string]int{}
	gated := []Role{RolePrimary, RoleRecallOnly, RoleGuardOnly}
	roles := pol.roleSets()

	// Collect deltas: gated sets must exist in both runs; report-only sets
	// are compared when present; sealed and unknown sets are ignored.
	known := map[string]bool{}
	for _, role := range append(gated, RoleReportOnly) {
		for _, name := range roles[role] {
			for _, holdout := range []bool{false, true} {
				if role == RoleReportOnly && holdout {
					continue
				}
				set := name
				if holdout {
					set += pol.HoldoutSuffix
				}
				known[set] = true
				b, c := base[set], cand[set]
				if b == nil || c == nil {
					if role == RoleReportOnly {
						continue
					}
					return nil, fmt.Errorf("set %s (%s) is missing from the %s run", set, role, missingSide(b, c))
				}
				if err := sameDataset(set, b, c); err != nil {
					return nil, err
				}
				idx[set] = len(d.Sets)
				d.Sets = append(d.Sets, newDelta(set, role, holdout, b, c))
			}
		}
	}
	for _, run := range []Run{base, cand} {
		for set := range run {
			if !known[set] && !contains(d.Ignored, set) {
				d.Ignored = append(d.Ignored, set)
			}
		}
	}
	sort.Strings(d.Ignored)
	for _, set := range d.Ignored {
		if !strings.HasSuffix(set, pol.SealedSuffix) {
			d.note("%s is not in the policy and was ignored", set)
		}
	}

	// Regressions.
	var tolerated []int
	for i := range d.Sets {
		s := &d.Sets[i]
		s.Status = "ok"
		switch s.Role {
		case RoleReportOnly:
			s.Status = "report-only"
		case RoleRecallOnly:
			if s.Cand.TP < s.Base.TP {
				s.Status = "regressed"
				d.reason("%s (RECALL-ONLY): recall dropped, TP %d -> %d", s.Set, s.Base.TP, s.Cand.TP)
			}
		case RolePrimary, RoleGuardOnly:
			precDrop := s.CandP < s.BaseP-eps
			f1Drop := s.CandF1 < s.BaseF1-eps
			lost := s.Base.TP - s.Cand.TP
			switch {
			case precDrop:
				s.Status = "regressed"
				d.reason("%s (%s): precision dropped %.4f -> %.4f (F1 %.4f -> %.4f)", s.Set, s.Role, s.BaseP, s.CandP, s.BaseF1, s.CandF1)
			case f1Drop && lost <= pol.Tolerance.MaxTPLost:
				s.Status = "tolerated"
				tolerated = append(tolerated, i)
			case f1Drop:
				s.Status = "regressed"
				d.reason("%s (%s): F1 dropped %.4f -> %.4f (TP %d -> %d)", s.Set, s.Role, s.BaseF1, s.CandF1, s.Base.TP, s.Cand.TP)
			}
		}
	}
	if len(tolerated) > 0 {
		var sumB, sumC float64
		for _, name := range pol.Roles.Primary {
			s := d.Sets[idx[name+pol.HoldoutSuffix]]
			sumB += s.BaseF1
			sumC += s.CandF1
		}
		ok := len(tolerated) <= pol.Tolerance.MaxSets && sumC > sumB+eps
		for _, i := range tolerated {
			s := &d.Sets[i]
			if ok {
				d.note("%s (%s): F1 %.4f -> %.4f tolerated (TP %d -> %d, precision %.4f -> %.4f, PRIMARY holdout F1 sum %.4f -> %.4f)",
					s.Set, s.Role, s.BaseF1, s.CandF1, s.Base.TP, s.Cand.TP, s.BaseP, s.CandP, sumB, sumC)
				continue
			}
			s.Status = "regressed"
			why := fmt.Sprintf("PRIMARY holdout F1 sum did not rise (%.4f -> %.4f)", sumB, sumC)
			if len(tolerated) > pol.Tolerance.MaxSets {
				why = fmt.Sprintf("%d sets need the single-sample tolerance, at most %d allowed", len(tolerated), pol.Tolerance.MaxSets)
			}
			d.reason("%s (%s): F1 dropped %.4f -> %.4f (TP %d -> %d); tolerance not applicable: %s",
				s.Set, s.Role, s.BaseF1, s.CandF1, s.Base.TP, s.Cand.TP, why)
		}
	}

	// PRIMARY dev F1 gains: generalization ratio and paired bootstrap.
	var gains []string
	for _, name := range pol.Roles.Primary {
		dev := &d.Sets[idx[name]]
		hold := d.Sets[idx[name+pol.HoldoutSuffix]]
		gain := dev.CandF1 - dev.BaseF1
		if gain <= eps {
			continue
		}
		holdGain := hold.CandF1 - hold.BaseF1
		generalizes := true
		required := pol.Generalization.MinHoldoutRatio * gain
		if step := oneTPStep(hold.Base); required > step {
			if holdGain < required-eps {
				generalizes = false
				d.reason("%s: dev F1 gain %+.4f does not generalize: holdout F1 gain %+.4f < %.2f x dev gain (%.4f)",
					name, gain, holdGain, pol.Generalization.MinHoldoutRatio, required)
			}
		} else {
			d.note("%s: required holdout gain %.4f is below one holdout TP (%.4f); ratio not applied, holdout must only not drop",
				name, required, step)
		}

		codes, src, err := pairCodes(base[name], cand[name])
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		if codes == nil {
			d.note("%s: dev F1 gain %+.4f not counted: no per-sample predictions or fp/fn files in both runs (run eval with --predictions)", name, gain)
			continue
		}
		bs := pairedBootstrap(codes, pol.Significance.Resamples, pol.Significance.Seed, pol.Significance.Confidence)
		bs.Source = src
		dev.Bootstrap = &bs
		if bs.Lower <= 0 {
			d.note("%s: dev F1 gain %+.4f not significant (%.0f%% CI [%+.4f, %+.4f])", name, gain, 100*pol.Significance.Confidence, bs.Lower, bs.Upper)
			continue
		}
		if generalizes {
			gains = append(gains, fmt.Sprintf("%s dev F1 %.4f -> %.4f (%.0f%% CI [%+.4f, %+.4f]), holdout %+.4f",
				name, dev.BaseF1, dev.CandF1, 100*pol.Significance.Confidence, bs.Lower, bs.Upper, holdGain))
		}
	}

	// Latency on gated dev sets.
	lat := pol.Latency
	identical := true
	var drops []string
	for i := range d.Sets {
		s := &d.Sets[i]
		if s.Role == RoleReportOnly {
			continue
		}
		if s.Base.TP != s.Cand.TP || s.Base.FP != s.Cand.FP {
			identical = false
		}
		if s.Holdout {
			continue
		}
		if s.CandP95Ms > lat.BudgetP95Ms {
			d.reason("%s: p95 latency %.2f ms is over the %.0f ms budget", s.Set, s.CandP95Ms, lat.BudgetP95Ms)
		}
		diff := s.CandP95Ms - s.BaseP95Ms
		if diff > lat.NoiseFloorMs && s.CandP95Ms > s.BaseP95Ms*(1+lat.MaxRisePct/100) {
			if len(gains) == 0 {
				d.reason("%s: p95 latency rose %.2f -> %.2f ms (> %.0f%%) without a gain", s.Set, s.BaseP95Ms, s.CandP95Ms, lat.MaxRisePct)
			} else {
				d.note("%s: p95 latency rose %.2f -> %.2f ms, accepted with the F1 gain", s.Set, s.BaseP95Ms, s.CandP95Ms)
			}
		}
		if -diff >= lat.NoiseFloorMs && s.CandP95Ms <= s.BaseP95Ms*(1-lat.MinDropPct/100) {
			drops = append(drops, fmt.Sprintf("%s p95 %.2f -> %.2f ms", s.Set, s.BaseP95Ms, s.CandP95Ms))
		}
	}

	switch {
	case len(d.Reasons) > 0:
		d.Decision, d.Kind = Revert, KindRegression
	case len(gains) > 0:
		d.Decision, d.Kind = Accept, KindGain
		d.Reasons = append(d.Reasons, gains...)
	case identical && len(drops) > 0:
		d.Decision, d.Kind = Accept, KindLatency
		d.Reasons = append(d.Reasons, "latency-only: TP/FP identical on every gated set; "+strings.Join(drops, ", "))
	default:
		d.Decision, d.Kind = Revert, KindNoGain
		d.reason("no significant, generalizing PRIMARY dev F1 gain and no latency-only improvement (>= %.0f%% p95 drop with identical TP/FP)", lat.MinDropPct)
	}
	return d, nil
}

func missingSide(b, c *SetRun) string {
	if b == nil {
		return "baseline"
	}
	return "candidate"
}

func contains(xs []string, x string) bool {
	for _, v := range xs {
		if v == x {
			return true
		}
	}
	return false
}

func total(c eval.Counts) int { return c.TP + c.FP + c.TN + c.FN }

func sameDataset(set string, b, c *SetRun) error {
	if total(b.Counts) != total(c.Counts) || b.Counts.TP+b.Counts.FN != c.Counts.TP+c.Counts.FN {
		return fmt.Errorf("set %s: baseline and candidate evaluated different datasets (%d/%d samples, %d/%d positives)",
			set, total(b.Counts), total(c.Counts), b.Counts.TP+b.Counts.FN, c.Counts.TP+c.Counts.FN)
	}
	return nil
}

func newDelta(set string, role Role, holdout bool, b, c *SetRun) SetDelta {
	return SetDelta{
		Set: set, Role: role, Holdout: holdout,
		Base: b.Counts, Cand: c.Counts,
		BaseF1: b.Counts.F1(), CandF1: c.Counts.F1(),
		BaseP: b.Counts.Precision(), CandP: c.Counts.Precision(),
		BaseR: b.Counts.Recall(), CandR: c.Counts.Recall(),
		BaseP95Ms: b.P95Ms, CandP95Ms: c.P95Ms,
	}
}

// oneTPStep is the F1 change of turning one FN of c into a TP, or +Inf when
// c has no FN (the holdout cannot gain a TP, so the ratio is not resolvable).
func oneTPStep(c eval.Counts) float64 {
	if c.FN == 0 {
		return math.Inf(1)
	}
	n := c
	n.TP++
	n.FN--
	return n.F1() - c.F1()
}

// Pair codes: bit 0 = label, bit 1 = baseline predicted, bit 2 = candidate predicted.
func code(label int, bp, cp bool) uint8 {
	c := uint8(label & 1)
	if bp {
		c |= 2
	}
	if cp {
		c |= 4
	}
	return c
}

// pairCodes pairs the per-sample predictions of two runs of one set. It uses
// predictions-<set>.jsonl when both runs have it (checking index, label and
// text hash), else reconstructs the paired table from the fp/fn files and
// counts. It returns nil codes when neither source is available in both runs.
func pairCodes(b, c *SetRun) ([]uint8, string, error) {
	if b.Pred != nil && c.Pred != nil {
		if len(b.Pred) != len(c.Pred) || len(b.Pred) != total(b.Counts) {
			return nil, "", fmt.Errorf("predictions have %d/%d rows for %d samples", len(b.Pred), len(c.Pred), total(b.Counts))
		}
		codes := make([]uint8, len(b.Pred))
		for i := range b.Pred {
			bp, cp := b.Pred[i], c.Pred[i]
			if bp.Index != cp.Index || bp.Label != cp.Label || bp.TextSHA256 != cp.TextSHA256 {
				return nil, "", fmt.Errorf("prediction row %d differs between runs (index, label or text hash)", i)
			}
			codes[i] = code(bp.Label, bp.Predicted, cp.Predicted)
		}
		return codes, "predictions", nil
	}
	if b.FP == nil || b.FN == nil || c.FP == nil || c.FN == nil {
		return nil, "", nil
	}
	var codes []uint8
	add := func(n int, cd uint8) {
		for ; n > 0; n-- {
			codes = append(codes, cd)
		}
	}
	// Positives: a text in an fn file was missed by that run.
	// Negatives: a text in an fp file was flagged by that run.
	for label, files := range map[int][2]map[string]int{1: {b.FN, c.FN}, 0: {b.FP, c.FP}} {
		var n int
		if label == 1 {
			n = b.Counts.TP + b.Counts.FN
		} else {
			n = b.Counts.FP + b.Counts.TN
		}
		hit := label == 1 // the prediction of a sample not listed in the file
		keys := map[string]bool{}
		for k := range files[0] {
			keys[k] = true
		}
		for k := range files[1] {
			keys[k] = true
		}
		listed := 0
		for k := range keys {
			nb, nc := files[0][k], files[1][k]
			both := min(nb, nc)
			add(both, code(label, !hit, !hit))
			add(nb-both, code(label, !hit, hit))
			add(nc-both, code(label, hit, !hit))
			listed += max(nb, nc)
		}
		if listed > n {
			return nil, "", fmt.Errorf("fp/fn files list more samples than the set has")
		}
		add(n-listed, code(label, hit, hit))
	}
	// Map iteration order does not matter: the bootstrap below only sees
	// codes through a sorted copy.
	sort.Slice(codes, func(i, j int) bool { return codes[i] < codes[j] })
	return codes, "fp_fn", nil
}

func f1(tp, fp, fn int) float64 { return eval.Counts{TP: tp, FP: fp, FN: fn}.F1() }

// pairedBootstrap resamples the paired samples with replacement and returns
// the percentile confidence interval of F1(candidate) - F1(baseline).
func pairedBootstrap(codes []uint8, resamples int, seed uint64, confidence float64) Bootstrap {
	score := func(sample func(int) uint8, n int) float64 {
		var btp, bfp, bfn, ctp, cfp, cfn int
		for i := 0; i < n; i++ {
			cd := sample(i)
			pos, bp, cp := cd&1 == 1, cd&2 != 0, cd&4 != 0
			switch {
			case pos && bp:
				btp++
			case pos:
				bfn++
			case bp:
				bfp++
			}
			switch {
			case pos && cp:
				ctp++
			case pos:
				cfn++
			case cp:
				cfp++
			}
		}
		return f1(ctp, cfp, cfn) - f1(btp, bfp, bfn)
	}
	n := len(codes)
	out := Bootstrap{Observed: score(func(i int) uint8 { return codes[i] }, n)}
	if n == 0 {
		return out
	}
	rng := rand.New(rand.NewPCG(seed, seed^0x9e3779b97f4a7c15))
	diffs := make([]float64, resamples)
	for r := range diffs {
		diffs[r] = score(func(int) uint8 { return codes[rng.IntN(n)] }, n)
	}
	sort.Float64s(diffs)
	alpha := (1 - confidence) / 2
	lo := int(math.Floor(alpha * float64(resamples)))
	hi := int(math.Ceil((1-alpha)*float64(resamples))) - 1
	out.Lower, out.Upper = diffs[lo], diffs[min(hi, resamples-1)]
	return out
}
