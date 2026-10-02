package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/eval"
	"github.com/user/safeanalyze/pkg/report"
	"github.com/user/safeanalyze/pkg/version"
)

var evalCmd = &cobra.Command{
	Use:   "eval <file.jsonl>",
	Short: "Measure fast-mode precision/recall on a labeled JSONL dataset",
	Long: `Run the fast inspect check suite (yara + hiddenchars) on every sample of a
labeled JSONL dataset and report TP/FP/TN/FN, precision, recall, F1,
per-rule hit counts and per-sample latency.

Each line is {"text": "...", "label": 1|0, "source": "..."} where label 1 is
an injection and 0 is benign. Each sample is scanned as one payload, like
'inspect --body'. A sample is predicted positive when it has >= 1 finding.
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		jsonOut, _ := cmd.Flags().GetString("json")
		fpOut, _ := cmd.Flags().GetString("fp-out")
		fnOut, _ := cmd.Flags().GetString("fn-out")

		samples, err := eval.LoadFile(args[0])
		if err != nil {
			return err
		}
		if len(samples) == 0 {
			return fmt.Errorf("dataset %s contains no samples", args[0])
		}

		start := time.Now()
		res, err := eval.Run(samples, func(text string) ([]report.Finding, error) {
			rep, err := inspectPayload(text)
			if err != nil {
				return nil, err
			}
			return rep.Findings, nil
		})
		if err != nil {
			return err
		}
		durationMs := time.Since(start).Milliseconds()

		printEvalResult(args[0], res, durationMs)

		if jsonOut != "" {
			doc := struct {
				Metadata map[string]any `json:"metadata"`
				*eval.Result
			}{
				Metadata: map[string]any{
					"safeanalyze_version": version.Version,
					"scan_mode":           "fast",
					"duration_ms":         durationMs,
					"dataset":             args[0],
				},
				Result: res,
			}
			data, err := json.MarshalIndent(doc, "", "  ")
			if err != nil {
				return fmt.Errorf("encoding eval result: %w", err)
			}
			if err := os.WriteFile(jsonOut, append(data, '\n'), 0644); err != nil {
				return fmt.Errorf("writing eval result: %w", err)
			}
		}
		if fpOut != "" {
			if err := writeOutcomes(fpOut, res.FalsePositives); err != nil {
				return err
			}
		}
		if fnOut != "" {
			if err := writeOutcomes(fnOut, res.FalseNegatives); err != nil {
				return err
			}
		}
		return nil
	},
}

func printEvalResult(dataset string, res *eval.Result, durationMs int64) {
	fmt.Printf("dataset:   %s (%d samples, %d ms)\n", dataset, res.Samples, durationMs)
	fmt.Printf("confusion: TP=%d FP=%d TN=%d FN=%d\n", res.TP, res.FP, res.TN, res.FN)
	fmt.Printf("metrics:   precision=%.3f recall=%.3f f1=%.3f\n", res.Precision, res.Recall, res.F1)
	fmt.Printf("latency:   p50=%.2f ms p95=%.2f ms\n", res.LatencyP50Ms, res.LatencyP95Ms)

	sources := make([]string, 0, len(res.BySource))
	for s := range res.BySource {
		sources = append(sources, s)
	}
	sort.Strings(sources)
	fmt.Println("by source:")
	for _, s := range sources {
		m := res.BySource[s]
		fmt.Printf("  %-40s n=%-4d TP=%d FP=%d TN=%d FN=%d precision=%.3f recall=%.3f f1=%.3f\n",
			s, m.Samples, m.TP, m.FP, m.TN, m.FN, m.Precision, m.Recall, m.F1)
	}

	rules := make([]string, 0, len(res.RuleHits))
	for r := range res.RuleHits {
		rules = append(rules, r)
	}
	sort.Strings(rules)
	fmt.Println("rule hits (samples):")
	if len(rules) == 0 {
		fmt.Println("  none")
	}
	for _, r := range rules {
		h := res.RuleHits[r]
		fmt.Printf("  %-40s injection=%d benign=%d\n", r, h.Injection, h.Benign)
	}
}

// writeOutcomes writes misclassified samples as JSONL for gap analysis.
func writeOutcomes(path string, outcomes []eval.Outcome) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("creating %s: %w", path, err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	for _, o := range outcomes {
		if err := enc.Encode(o); err != nil {
			return fmt.Errorf("writing %s: %w", path, err)
		}
	}
	if err := w.Flush(); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}
	return nil
}

func init() {
	evalCmd.Flags().String("json", "", "write metrics and metadata as JSON to this file")
	evalCmd.Flags().String("fp-out", "", "write false-positive samples as JSONL to this file")
	evalCmd.Flags().String("fn-out", "", "write false-negative samples as JSONL to this file")
}
