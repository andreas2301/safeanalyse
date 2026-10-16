package gate

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/user/safeanalyze/pkg/eval"
)

// mkSet builds a SetRun whose per-sample predictions are in a canonical
// order: positives first (the first tp predicted), then negatives (the first
// fp predicted). Two runs built this way pair sample by sample, so a run with
// more TP is a strict superset of one with fewer.
func mkSet(tp, fp, tn, fn int, p95 float64) *SetRun {
	s := &SetRun{Counts: eval.Counts{TP: tp, FP: fp, TN: tn, FN: fn}, P95Ms: p95}
	for i := 0; i < tp+fn; i++ {
		s.Pred = append(s.Pred, eval.Prediction{Index: i, Label: 1, Predicted: i < tp, TextSHA256: fmt.Sprint("p", i)})
	}
	for i := 0; i < fp+tn; i++ {
		s.Pred = append(s.Pred, eval.Prediction{Index: tp + fn + i, Label: 0, Predicted: i < fp, TextSHA256: fmt.Sprint("n", i)})
	}
	return s
}

// baseRun mirrors the v0.3.44 metrics (TP/FP/TN/FN and p95) of every gated
// set plus the report-only stress set.
func baseRun() Run {
	return Run{
		"llmail":                  mkSet(240, 0, 160, 60, 12.7),
		"llmail-holdout":          mkSet(216, 0, 160, 84, 12.5),
		"browsesafe":              mkSet(199, 26, 274, 101, 44.1),
		"browsesafe-holdout":      mkSet(179, 27, 273, 121, 48.2),
		"toolresp":                mkSet(189, 31, 269, 111, 11.7),
		"toolresp-holdout":        mkSet(157, 35, 265, 143, 10.7),
		"multilingual":            mkSet(21, 0, 0, 819, 10.9),
		"multilingual-holdout":    mkSet(20, 0, 0, 805, 9.7),
		"deepset":                 mkSet(46, 0, 56, 14, 10.2),
		"deepset-holdout":         mkSet(62, 0, 343, 141, 9.9),
		"nemotron":                mkSet(34, 0, 300, 266, 10.1),
		"nemotron-holdout":        mkSet(35, 0, 300, 265, 10.6),
		"spanish":                 mkSet(37, 0, 300, 15, 9.8),
		"spanish-holdout":         mkSet(23, 0, 371, 30, 10.1),
		"promptprotection-stress": mkSet(30, 30, 166, 150, 9.4),
	}
}

func clone(r Run) Run {
	out := Run{}
	for k, v := range r {
		c := mkSet(v.Counts.TP, v.Counts.FP, v.Counts.TN, v.Counts.FN, v.P95Ms)
		out[k] = c
	}
	return out
}

// set replaces the counts of one set (keeping its sample totals) in run.
func set(r Run, name string, tp, fp int) {
	old := r[name]
	pos, neg := old.Counts.TP+old.Counts.FN, old.Counts.FP+old.Counts.TN
	r[name] = mkSet(tp, fp, neg-fp, pos-tp, old.P95Ms)
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name     string
		base     func(Run)
		cand     func(Run)
		decision string
		kind     string
		reason   string // substring of some reason
		note     string // substring of some note
	}{
		{
			name:     "identical runs: no gain",
			cand:     func(Run) {},
			decision: Revert, kind: KindNoGain, reason: "no significant",
		},
		{
			name: "significant generalizing PRIMARY dev gain",
			cand: func(r Run) {
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
			},
			decision: Accept, kind: KindGain, reason: "llmail dev F1",
		},
		{
			name: "generalization ratio: holdout gain below half the dev gain",
			cand: func(r Run) {
				set(r, "llmail", 260, 0)         // dev F1 +0.0397
				set(r, "llmail-holdout", 220, 0) // holdout F1 +0.0090 < 0.0198
			},
			decision: Revert, kind: KindRegression, reason: "does not generalize",
		},
		{
			name: "generalization ratio: holdout gain exactly at the threshold passes",
			cand: func(r Run) {
				set(r, "toolresp", 229, 31)         // dev F1 0.7269 -> 0.8031 (+0.0762)
				set(r, "toolresp-holdout", 182, 35) // holdout 0.6382 -> 0.7000 (+0.0618)
			},
			decision: Accept, kind: KindGain,
		},
		{
			name: "tiny dev gain: ratio not resolvable, gain not significant",
			cand: func(r Run) {
				set(r, "browsesafe", 200, 26)
			},
			decision: Revert, kind: KindNoGain, note: "not significant",
		},
		{
			name: "dev gain on two PRIMARY sets, one does not generalize",
			cand: func(r Run) {
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
				set(r, "toolresp", 220, 31)
			},
			decision: Revert, kind: KindRegression, reason: "toolresp: dev F1 gain",
		},
		{
			name: "PRIMARY dev precision drop despite gain",
			cand: func(r Run) {
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
				set(r, "browsesafe", 199, 27)
			},
			decision: Revert, kind: KindRegression, reason: "browsesafe (PRIMARY): precision dropped",
		},
		{
			name: "PRIMARY holdout F1 drop of two TP",
			cand: func(r Run) {
				set(r, "browsesafe", 230, 26)
				set(r, "browsesafe-holdout", 205, 27)
				set(r, "llmail-holdout", 214, 0)
			},
			decision: Revert, kind: KindRegression, reason: "llmail-holdout (PRIMARY): F1 dropped",
		},
		{
			name: "one TP lost with FP > 0 lowers precision: no tolerance",
			cand: func(r Run) {
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
				set(r, "browsesafe-holdout", 178, 27)
			},
			decision: Revert, kind: KindRegression, reason: "browsesafe-holdout (PRIMARY): precision dropped",
		},
		{
			name: "GUARD-ONLY F1 drop of two TP",
			cand: func(r Run) {
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
				set(r, "nemotron", 32, 0)
			},
			decision: Revert, kind: KindRegression, reason: "nemotron (GUARD-ONLY): F1 dropped",
		},
		{
			name: "GUARD-ONLY precision drop",
			cand: func(r Run) {
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
				set(r, "spanish-holdout", 23, 1)
			},
			decision: Revert, kind: KindRegression, reason: "spanish-holdout (GUARD-ONLY): precision dropped",
		},
		{
			name: "v0.3.24 precedent: one TP lost, precision up, PRIMARY holdout sum rises",
			base: func(r Run) { set(r, "deepset-holdout", 54, 2) },
			cand: func(r Run) {
				set(r, "deepset-holdout", 53, 0) // F1 0.4170 -> 0.4141, P 0.964 -> 1.0
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
			},
			decision: Accept, kind: KindGain, note: "deepset-holdout (GUARD-ONLY): F1 0.4170 -> 0.4141 tolerated",
		},
		{
			name: "tolerance on a PRIMARY dev set",
			cand: func(r Run) {
				set(r, "llmail", 239, 0)
				set(r, "browsesafe", 230, 26)
				set(r, "browsesafe-holdout", 205, 27)
			},
			decision: Accept, kind: KindGain, note: "llmail (PRIMARY): F1 0.8889 -> 0.8868 tolerated",
		},
		{
			name: "tolerance needs the PRIMARY holdout F1 sum to rise",
			cand: func(r Run) {
				set(r, "deepset-holdout", 61, 0)
				set(r, "llmail", 260, 0) // dev-only gain; holdout sum unchanged
			},
			decision: Revert, kind: KindRegression, reason: "PRIMARY holdout F1 sum did not rise",
		},
		{
			name: "tolerance covers at most one set",
			cand: func(r Run) {
				set(r, "deepset-holdout", 61, 0)
				set(r, "spanish", 36, 0)
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
			},
			decision: Revert, kind: KindRegression, reason: "2 sets need the single-sample tolerance",
		},
		{
			name: "tolerance does not cover a precision drop",
			cand: func(r Run) {
				set(r, "browsesafe", 198, 27)
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
			},
			decision: Revert, kind: KindRegression, reason: "browsesafe (PRIMARY): precision dropped",
		},
		{
			name: "RECALL-ONLY loses one TP: no tolerance",
			cand: func(r Run) {
				set(r, "multilingual-holdout", 19, 0)
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
			},
			decision: Revert, kind: KindRegression, reason: "multilingual-holdout (RECALL-ONLY): recall dropped",
		},
		{
			name: "RECALL-ONLY gain alone does not count",
			cand: func(r Run) {
				set(r, "multilingual", 40, 0)
				set(r, "multilingual-holdout", 40, 0)
			},
			decision: Revert, kind: KindNoGain,
		},
		{
			name: "GUARD-ONLY gain alone does not count",
			cand: func(r Run) {
				set(r, "deepset", 56, 0)
				set(r, "deepset-holdout", 120, 0)
			},
			decision: Revert, kind: KindNoGain,
		},
		{
			name: "REPORT-ONLY and sealed regressions are ignored",
			base: func(r Run) { r["llmail-sealed"] = mkSet(200, 0, 160, 100, 12) },
			cand: func(r Run) {
				set(r, "promptprotection-stress", 10, 90)
				r["llmail-sealed"] = mkSet(100, 50, 110, 200, 12)
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
			},
			decision: Accept, kind: KindGain,
		},
		{
			name:     "latency-only accept: counts identical, p95 drops >= 10%",
			base:     func(r Run) { r["browsesafe"].P95Ms = 70.72 },
			cand:     func(r Run) { r["browsesafe"].P95Ms = 44.09 },
			decision: Accept, kind: KindLatency, reason: "browsesafe p95 70.72 -> 44.09",
		},
		{
			name:     "latency drop under 10% is no gain",
			cand:     func(r Run) { r["browsesafe"].P95Ms = 40.5 },
			decision: Revert, kind: KindNoGain,
		},
		{
			name:     "latency drop under the noise floor is no gain",
			cand:     func(r Run) { r["llmail"].P95Ms = 11.0 },
			decision: Revert, kind: KindNoGain,
		},
		{
			name: "latency drop with changed counts is not latency-only",
			cand: func(r Run) {
				r["browsesafe"].P95Ms = 20
				set(r, "browsesafe-holdout", 180, 27)
			},
			decision: Revert, kind: KindNoGain,
		},
		{
			name:     "latency rise > 15% without a gain",
			cand:     func(r Run) { r["browsesafe"].P95Ms = 55 },
			decision: Revert, kind: KindRegression, reason: "p95 latency rose",
		},
		{
			name:     "latency rise > 15% under the noise floor is ignored",
			cand:     func(r Run) { r["deepset"].P95Ms = 12 },
			decision: Revert, kind: KindNoGain,
		},
		{
			name: "latency rise > 15% with a gain is accepted",
			cand: func(r Run) {
				r["browsesafe"].P95Ms = 60
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
			},
			decision: Accept, kind: KindGain, note: "accepted with the F1 gain",
		},
		{
			name: "latency over the budget",
			cand: func(r Run) {
				r["browsesafe"].P95Ms = 101
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
			},
			decision: Revert, kind: KindRegression, reason: "over the 100 ms budget",
		},
		{
			name: "holdout latency is not gated",
			cand: func(r Run) {
				r["browsesafe-holdout"].P95Ms = 150
			},
			decision: Revert, kind: KindNoGain,
		},
		{
			name: "no per-sample predictions: gain not counted",
			cand: func(r Run) {
				set(r, "llmail", 260, 0)
				set(r, "llmail-holdout", 230, 0)
				r["llmail"].Pred = nil
			},
			decision: Revert, kind: KindNoGain, note: "no per-sample predictions",
		},
	}
	pol := DefaultPolicy()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := baseRun()
			if tt.base != nil {
				tt.base(base)
			}
			cand := clone(base)
			tt.cand(cand)
			d, err := Evaluate(pol, base, cand)
			if err != nil {
				t.Fatal(err)
			}
			if d.Decision != tt.decision || d.Kind != tt.kind {
				t.Fatalf("decision = %s (%s), want %s (%s)\nreasons: %q\nnotes: %q", d.Decision, d.Kind, tt.decision, tt.kind, d.Reasons, d.Notes)
			}
			if tt.reason != "" && !anyContains(d.Reasons, tt.reason) {
				t.Errorf("no reason contains %q: %q", tt.reason, d.Reasons)
			}
			if tt.note != "" && !anyContains(d.Notes, tt.note) {
				t.Errorf("no note contains %q: %q", tt.note, d.Notes)
			}
		})
	}
}

func anyContains(xs []string, sub string) bool {
	for _, x := range xs {
		if strings.Contains(x, sub) {
			return true
		}
	}
	return false
}

func TestEvaluateErrors(t *testing.T) {
	pol := DefaultPolicy()
	base := baseRun()
	cand := clone(base)
	delete(cand, "spanish-holdout")
	if _, err := Evaluate(pol, base, cand); err == nil || !strings.Contains(err.Error(), "spanish-holdout") {
		t.Fatalf("missing set: err = %v", err)
	}
	cand = clone(base)
	cand["toolresp"] = mkSet(189, 31, 268, 111, 11)
	if _, err := Evaluate(pol, base, cand); err == nil || !strings.Contains(err.Error(), "different datasets") {
		t.Fatalf("dataset mismatch: err = %v", err)
	}
	// Report-only sets may be missing.
	cand = clone(base)
	delete(cand, "promptprotection-stress")
	if _, err := Evaluate(pol, base, cand); err != nil {
		t.Fatalf("missing report-only set: %v", err)
	}
	// Predictions that do not pair (different text hash) are an error.
	cand = clone(base)
	set(cand, "llmail", 260, 0)
	set(cand, "llmail-holdout", 230, 0)
	cand["llmail"].Pred[0].TextSHA256 = "other"
	if _, err := Evaluate(pol, base, cand); err == nil || !strings.Contains(err.Error(), "differs between runs") {
		t.Fatalf("prediction mismatch: err = %v", err)
	}
}

// toFPFN converts canonical predictions into the fp/fn multisets eval writes.
func toFPFN(s *SetRun) *SetRun {
	out := &SetRun{Counts: s.Counts, P95Ms: s.P95Ms, FP: map[string]int{}, FN: map[string]int{}}
	for _, p := range s.Pred {
		switch {
		case p.Label == 1 && !p.Predicted:
			out.FN[p.TextSHA256]++
		case p.Label == 0 && p.Predicted:
			out.FP[p.TextSHA256]++
		}
	}
	return out
}

func TestPairCodesFallbackMatchesPredictions(t *testing.T) {
	b := mkSet(190, 30, 270, 110, 0)
	c := mkSet(205, 28, 272, 95, 0)
	// Make the candidate also lose two positives the baseline caught.
	c.Pred[0].Predicted, c.Pred[1].Predicted = false, false
	c.Pred[205].Predicted, c.Pred[206].Predicted = true, true
	fromPred, src, err := pairCodes(b, c)
	if err != nil || src != "predictions" {
		t.Fatalf("predictions: %v %s", err, src)
	}
	fromFiles, src, err := pairCodes(toFPFN(b), toFPFN(c))
	if err != nil || src != "fp_fn" {
		t.Fatalf("fp/fn: %v %s", err, src)
	}
	hist := func(codes []uint8) [8]int {
		var h [8]int
		for _, c := range codes {
			h[c]++
		}
		return h
	}
	if hist(fromPred) != hist(fromFiles) {
		t.Fatalf("paired tables differ: %v vs %v", hist(fromPred), hist(fromFiles))
	}
	p1 := pairedBootstrap(fromPred, 2000, 7, 0.95)
	p2 := pairedBootstrap(fromPred, 2000, 7, 0.95)
	if p1 != p2 {
		t.Fatalf("bootstrap not deterministic: %+v vs %+v", p1, p2)
	}
	if !(p1.Lower <= p1.Observed && p1.Observed <= p1.Upper) {
		t.Fatalf("observed outside CI: %+v", p1)
	}
	if want := c.Counts.F1() - b.Counts.F1(); p1.Observed-want > 1e-12 || want-p1.Observed > 1e-12 {
		t.Fatalf("observed = %v, want %v", p1.Observed, want)
	}
}

func TestPairedBootstrapSignificance(t *testing.T) {
	codes := func(gain, loss, same int) []uint8 {
		var out []uint8
		for i := 0; i < gain; i++ {
			out = append(out, code(1, false, true))
		}
		for i := 0; i < loss; i++ {
			out = append(out, code(1, true, false))
		}
		for i := 0; i < same; i++ {
			out = append(out, code(1, true, true), code(0, false, false))
		}
		return out
	}
	if bs := pairedBootstrap(codes(20, 0, 250), 10000, 1337, 0.95); bs.Lower <= 0 {
		t.Errorf("20 net TP gained: CI lower = %v, want > 0", bs.Lower)
	}
	if bs := pairedBootstrap(codes(1, 0, 250), 10000, 1337, 0.95); bs.Lower > 0 {
		t.Errorf("1 TP gained: CI lower = %v, want <= 0", bs.Lower)
	}
	if bs := pairedBootstrap(codes(12, 10, 250), 10000, 1337, 0.95); bs.Lower > 0 {
		t.Errorf("12 gained / 10 lost: CI lower = %v, want <= 0", bs.Lower)
	}
}

func TestLoadRun(t *testing.T) {
	dir := t.TempDir()
	write := func(name, body string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0644); err != nil {
			t.Fatal(err)
		}
	}
	doc, _ := json.Marshal(map[string]any{"tp": 1, "fp": 1, "tn": 0, "fn": 1, "latency_p95_ms": 3.5, "metadata": map[string]any{}})
	write("eval-x.json", string(doc))
	write("predictions-x.jsonl", `{"index":0,"label":1,"predicted":true,"text_sha256":"a"}`+"\n")
	write("fp-x.jsonl", `{"text":"benign","label":0,"rules":["r"]}`+"\n")
	write("fn-x.jsonl", `{"text":"missed","label":1}`+"\n")
	run, err := LoadRun(dir)
	if err != nil {
		t.Fatal(err)
	}
	x := run["x"]
	if x == nil || x.Counts != (eval.Counts{TP: 1, FP: 1, FN: 1}) || x.P95Ms != 3.5 || len(x.Pred) != 1 || len(x.FP) != 1 || len(x.FN) != 1 {
		t.Fatalf("run = %+v", x)
	}
	if _, err := LoadRun(t.TempDir()); err == nil {
		t.Fatal("empty dir: want error")
	}
}

func TestCommittedPolicyMatchesDefault(t *testing.T) {
	p, err := LoadPolicy("../../testdata/eval/policy.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(p, DefaultPolicy()) {
		t.Fatalf("policy.yaml = %+v\nDefaultPolicy = %+v", p, DefaultPolicy())
	}
}

func TestPolicyValidate(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Policy)
		want   string
	}{
		{"default ok", func(*Policy) {}, ""},
		{"two roles", func(p *Policy) { p.Roles.GuardOnly = append(p.Roles.GuardOnly, "llmail") }, "two roles"},
		{"no primary", func(p *Policy) { p.Roles.Primary = nil }, "primary is empty"},
		{"same suffix", func(p *Policy) { p.SealedSuffix = p.HoldoutSuffix }, "suffix"},
		{"confidence", func(p *Policy) { p.Significance.Confidence = 1 }, "confidence"},
		{"resamples", func(p *Policy) { p.Significance.Resamples = 10 }, "resamples"},
		{"latency", func(p *Policy) { p.Latency.MinDropPct = 0 }, "latency"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := DefaultPolicy()
			tt.mutate(p)
			err := p.Validate()
			if tt.want == "" && err != nil || tt.want != "" && (err == nil || !strings.Contains(err.Error(), tt.want)) {
				t.Fatalf("err = %v, want %q", err, tt.want)
			}
		})
	}
	dir := t.TempDir()
	path := filepath.Join(dir, "p.yaml")
	if err := os.WriteFile(path, []byte("roles:\n  primary: [a]\nunknown_key: 1\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadPolicy(path); err == nil {
		t.Fatal("unknown key: want error")
	}
}
