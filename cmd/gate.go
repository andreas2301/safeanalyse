package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/gate"
	"github.com/user/safeanalyze/pkg/version"
)

const defaultPolicyPath = "testdata/eval/policy.yaml"

var gateCmd = &cobra.Command{
	Use:   "gate --baseline <dir> --candidate <dir>",
	Short: "Decide ACCEPT/REVERT for an autoresearch iteration from two eval directories",
	Long: `Compare the labeled-eval results of a candidate against a baseline and apply
the keep/revert decision rules (CLAUDE.md "Set taxonomy" and "Decision rules").

Each directory holds eval-<set>.json files written by 'safeanalyze eval --json',
and ideally predictions-<set>.jsonl from 'eval --predictions' (fp-/fn-<set>.jsonl
are used as a fallback) for the paired bootstrap on the PRIMARY dev sets.
Sealed sets (<set>-sealed) are always ignored.

Exit status: 0 for ACCEPT, 3 for REVERT, 1 for errors.
`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseDir, _ := cmd.Flags().GetString("baseline")
		candDir, _ := cmd.Flags().GetString("candidate")
		policyPath, _ := cmd.Flags().GetString("policy")
		jsonOut, _ := cmd.Flags().GetString("json")
		if baseDir == "" || candDir == "" {
			return fmt.Errorf("--baseline and --candidate are required")
		}

		var pol *gate.Policy
		if _, err := os.Stat(policyPath); os.IsNotExist(err) && !cmd.Flags().Changed("policy") {
			fmt.Fprintf(os.Stderr, "note: %s not found, using the built-in default policy\n", policyPath)
			pol = gate.DefaultPolicy()
		} else if pol, err = gate.LoadPolicy(policyPath); err != nil {
			return err
		}
		base, err := gate.LoadRun(baseDir)
		if err != nil {
			return fmt.Errorf("baseline: %w", err)
		}
		cand, err := gate.LoadRun(candDir)
		if err != nil {
			return fmt.Errorf("candidate: %w", err)
		}
		dec, err := gate.Evaluate(pol, base, cand)
		if err != nil {
			return err
		}
		printDecision(dec)

		if jsonOut != "" {
			doc := struct {
				Metadata map[string]any `json:"metadata"`
				*gate.Decision
			}{
				Metadata: map[string]any{
					"safeanalyze_version": version.Version,
					"baseline":            baseDir,
					"candidate":           candDir,
					"policy":              policyPath,
				},
				Decision: dec,
			}
			data, err := json.MarshalIndent(doc, "", "  ")
			if err != nil {
				return fmt.Errorf("encoding gate result: %w", err)
			}
			if err := os.WriteFile(jsonOut, append(data, '\n'), 0644); err != nil {
				return fmt.Errorf("writing gate result: %w", err)
			}
		}
		if dec.Decision != gate.Accept {
			os.Exit(3)
		}
		return nil
	},
}

func printDecision(d *gate.Decision) {
	fmt.Printf("%-28s %-12s %-22s %-22s %-17s %s\n", "set", "role", "F1 base -> cand", "P base -> cand", "TP/FP base->cand", "p95 ms base -> cand")
	for _, s := range d.Sets {
		counts := fmt.Sprintf("%d/%d -> %d/%d", s.Base.TP, s.Base.FP, s.Cand.TP, s.Cand.FP)
		fmt.Printf("%-28s %-12s %.4f -> %.4f       %.4f -> %.4f       %-17s %7.2f -> %7.2f  %s\n",
			s.Set, s.Role, s.BaseF1, s.CandF1, s.BaseP, s.CandP, counts, s.BaseP95Ms, s.CandP95Ms, s.Status)
	}
	for _, s := range d.Ignored {
		fmt.Printf("ignored: %s\n", s)
	}
	for _, n := range d.Notes {
		fmt.Printf("note: %s\n", n)
	}
	fmt.Printf("decision: %s (%s)\n", d.Decision, d.Kind)
	for _, r := range d.Reasons {
		fmt.Printf("  - %s\n", r)
	}
}

func init() {
	gateCmd.Flags().String("baseline", "", "directory with the baseline eval-<set>.json files")
	gateCmd.Flags().String("candidate", "", "directory with the candidate eval-<set>.json files")
	gateCmd.Flags().String("policy", defaultPolicyPath, "decision policy YAML (set roles and thresholds)")
	gateCmd.Flags().String("json", "", "write the decision and per-set deltas as JSON to this file")
	rootCmd.AddCommand(gateCmd)
}
