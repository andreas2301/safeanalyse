// Package eval measures detection quality of a safeanalyze check suite against
// labeled prompt-injection samples stored as JSONL.
package eval

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/user/safeanalyze/pkg/report"
)

// Label values used in eval datasets.
const (
	LabelBenign    = 0
	LabelInjection = 1
)

// Sample is one labeled payload. Label is 1 for injection, 0 for benign.
type Sample struct {
	Text   string `json:"text"`
	Label  int    `json:"label"`
	Source string `json:"source,omitempty"`
}

// Load reads JSONL samples from r. Blank lines are skipped; every other line
// must be a JSON object with a "text" field and a "label" of 0 or 1.
func Load(r io.Reader) ([]Sample, error) {
	var samples []Sample
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 64*1024*1024)
	line := 0
	for scanner.Scan() {
		line++
		raw := strings.TrimSpace(scanner.Text())
		if raw == "" {
			continue
		}
		var row struct {
			Text   *string `json:"text"`
			Label  *int    `json:"label"`
			Source string  `json:"source"`
		}
		if err := json.Unmarshal([]byte(raw), &row); err != nil {
			return nil, fmt.Errorf("line %d: %w", line, err)
		}
		if row.Text == nil {
			return nil, fmt.Errorf("line %d: missing text", line)
		}
		if row.Label == nil {
			return nil, fmt.Errorf("line %d: missing label", line)
		}
		if *row.Label != LabelBenign && *row.Label != LabelInjection {
			return nil, fmt.Errorf("line %d: label must be 0 or 1, got %d", line, *row.Label)
		}
		samples = append(samples, Sample{Text: *row.Text, Label: *row.Label, Source: row.Source})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading samples: %w", err)
	}
	return samples, nil
}

// LoadFile reads JSONL samples from path.
func LoadFile(path string) ([]Sample, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening dataset: %w", err)
	}
	defer f.Close()
	return Load(f)
}

// Detector runs a check suite on a single payload and returns its findings.
type Detector func(text string) ([]report.Finding, error)

// Counts is a binary confusion matrix.
type Counts struct {
	TP int `json:"tp"`
	FP int `json:"fp"`
	TN int `json:"tn"`
	FN int `json:"fn"`
}

// Add records one prediction against its label.
func (c *Counts) Add(label int, predicted bool) {
	switch {
	case predicted && label == LabelInjection:
		c.TP++
	case predicted:
		c.FP++
	case label == LabelInjection:
		c.FN++
	default:
		c.TN++
	}
}

// Precision returns TP/(TP+FP), or 0 when nothing was predicted positive.
func (c Counts) Precision() float64 { return ratio(c.TP, c.TP+c.FP) }

// Recall returns TP/(TP+FN), or 0 when there are no positive samples.
func (c Counts) Recall() float64 { return ratio(c.TP, c.TP+c.FN) }

// F1 returns the harmonic mean of precision and recall.
func (c Counts) F1() float64 { return ratio(2*c.TP, 2*c.TP+c.FP+c.FN) }

func ratio(num, den int) float64 {
	if den == 0 {
		return 0
	}
	return float64(num) / float64(den)
}

// Metrics is a confusion matrix plus derived scores.
type Metrics struct {
	Counts
	Samples   int     `json:"samples"`
	Precision float64 `json:"precision"`
	Recall    float64 `json:"recall"`
	F1        float64 `json:"f1"`
}

func newMetrics(c Counts) Metrics {
	return Metrics{
		Counts:    c,
		Samples:   c.TP + c.FP + c.TN + c.FN,
		Precision: c.Precision(),
		Recall:    c.Recall(),
		F1:        c.F1(),
	}
}

// RuleHits counts the samples on which a rule fired, split by sample label.
type RuleHits struct {
	Injection int `json:"injection"`
	Benign    int `json:"benign"`
}

// Outcome is a sample together with the rules that fired on it.
type Outcome struct {
	Sample
	Rules []string `json:"rules,omitempty"`
}

// Result is the outcome of evaluating a detector on a dataset.
type Result struct {
	Metrics
	BySource       map[string]Metrics  `json:"by_source"`
	RuleHits       map[string]RuleHits `json:"rule_hits"`
	LatencyP50Ms   float64             `json:"latency_p50_ms"`
	LatencyP95Ms   float64             `json:"latency_p95_ms"`
	FalsePositives []Outcome           `json:"-"`
	FalseNegatives []Outcome           `json:"-"`
}

// Run evaluates detect on every sample. A sample is predicted positive when
// the detector returns at least one finding.
func Run(samples []Sample, detect Detector) (*Result, error) {
	res := &Result{
		BySource: map[string]Metrics{},
		RuleHits: map[string]RuleHits{},
	}
	var total Counts
	bySource := map[string]*Counts{}
	latencies := make([]time.Duration, 0, len(samples))

	for i, s := range samples {
		start := time.Now()
		findings, err := detect(s.Text)
		latencies = append(latencies, time.Since(start))
		if err != nil {
			return nil, fmt.Errorf("sample %d: %w", i+1, err)
		}

		predicted := len(findings) > 0
		total.Add(s.Label, predicted)
		sc, ok := bySource[s.Source]
		if !ok {
			sc = &Counts{}
			bySource[s.Source] = sc
		}
		sc.Add(s.Label, predicted)

		rules := uniqueRules(findings)
		for _, rule := range rules {
			h := res.RuleHits[rule]
			if s.Label == LabelInjection {
				h.Injection++
			} else {
				h.Benign++
			}
			res.RuleHits[rule] = h
		}

		switch {
		case predicted && s.Label == LabelBenign:
			res.FalsePositives = append(res.FalsePositives, Outcome{Sample: s, Rules: rules})
		case !predicted && s.Label == LabelInjection:
			res.FalseNegatives = append(res.FalseNegatives, Outcome{Sample: s})
		}
	}

	res.Metrics = newMetrics(total)
	for src, c := range bySource {
		res.BySource[src] = newMetrics(*c)
	}
	res.LatencyP50Ms = Percentile(latencies, 50)
	res.LatencyP95Ms = Percentile(latencies, 95)
	return res, nil
}

func uniqueRules(findings []report.Finding) []string {
	seen := map[string]bool{}
	var rules []string
	for _, f := range findings {
		if !seen[f.RuleID] {
			seen[f.RuleID] = true
			rules = append(rules, f.RuleID)
		}
	}
	sort.Strings(rules)
	return rules
}

// Percentile returns the nearest-rank p-th percentile of d in milliseconds,
// or 0 for an empty slice.
func Percentile(d []time.Duration, p float64) float64 {
	if len(d) == 0 {
		return 0
	}
	sorted := append([]time.Duration(nil), d...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	rank := int(math.Ceil(p / 100 * float64(len(sorted))))
	if rank < 1 {
		rank = 1
	}
	if rank > len(sorted) {
		rank = len(sorted)
	}
	return float64(sorted[rank-1].Microseconds()) / 1000
}
