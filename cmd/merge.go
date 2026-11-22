package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/report"
	"github.com/user/safeanalyze/pkg/triage"
	"github.com/user/safeanalyze/pkg/version"
)

var mergeCmd = &cobra.Command{
	Use:   "merge <report1.json> <report2.json> [more...]",
	Short: "Merge two or more safeanalyze JSON reports and re-run triage",
	Long: `Combine multiple safeanalyze JSON reports (for example, a core scan
report plus a companion ML/LLM scanner report) into one unified, deduplicated
report, re-run triage/verdicts, and emit all configured output formats
(JSON, SARIF, Markdown, HTML).

Exits with status 2 when Tier A findings exist unless --no-fail is set.`,
	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		outputDir, _ := cmd.Flags().GetString("output-dir")
		noFail, _ := cmd.Flags().GetBool("no-fail")
		target, _ := cmd.Flags().GetString("target")

		reports := make([]*report.Report, 0, len(args))
		stageSet := map[string]bool{}
		for _, p := range args {
			data, err := os.ReadFile(p)
			if err != nil {
				return fmt.Errorf("reading %s: %w", p, err)
			}
			var r report.Report
			if err := json.Unmarshal(data, &r); err != nil {
				return fmt.Errorf("parsing %s: %w", p, err)
			}
			if target == "" && r.Target != "" {
				target = r.Target
			}
			for _, f := range r.Findings {
				if f.Source != "" {
					stageSet[f.Source] = true
				}
			}
			for _, c := range r.Coverage {
				if c.Status == "ran" && c.Check != "" {
					stageSet[c.Check] = true
				}
			}
			reports = append(reports, &r)
		}

		merged := report.Merge(target, reports)
		if merged.Metadata == nil {
			merged.Metadata = map[string]string{}
		}
		if _, ok := merged.Metadata["safeanalyze_version"]; !ok {
			merged.Metadata["safeanalyze_version"] = version.Version
		}
		merged.Metadata["merged_reports"] = fmt.Sprint(len(reports))

		stages := make([]string, 0, len(stageSet))
		for s := range stageSet {
			stages = append(stages, s)
		}
		sort.Strings(stages)

		triage.Assign(merged)
		triage.Finalize(merged, stages, nil)
		printReportSummary(merged)

		outCfg := cfg.Output
		if outputDir != "" {
			outCfg.OutDir = outputDir
		}
		if err := report.WriteAll(merged, outCfg); err != nil {
			return fmt.Errorf("writing merged reports: %w", err)
		}

		if cfg.Intent.FailOnTierA && !noFail {
			if n := merged.CountTier(report.TierA); n > 0 {
				return &ExitError{
					Code: ExitBlocked,
					Err: fmt.Errorf("blocked: %d Tier A finding(s) after merge; first: %s",
						n, describeFindings(tierFindings(merged, report.TierA))),
				}
			}
		}
		return nil
	},
}

func init() {
	mergeCmd.Flags().StringP("output-dir", "o", "", "output directory for merged reports")
	mergeCmd.Flags().Bool("no-fail", false, "report only: do not exit non-zero on Tier A findings")
	mergeCmd.Flags().String("target", "", "override target label on the merged report")
}
