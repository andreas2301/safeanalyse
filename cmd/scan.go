package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/checks/external"
	"github.com/user/safeanalyze/pkg/checks/intent"
	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/pipeline"
	"github.com/user/safeanalyze/pkg/report"
	"github.com/user/safeanalyze/pkg/triage"
	"github.com/user/safeanalyze/pkg/version"
)

// ExitError carries a specific process exit status.
type ExitError struct {
	Code int
	Err  error
}

func (e *ExitError) Error() string { return e.Err.Error() }

// ExitBlocked is the exit status of a scan blocked by Tier A findings.
const ExitBlocked = 2

var scanCmd = &cobra.Command{
	Use:   "scan <path>",
	Short: "Run security scanners and emit a structured report",
	Long: `Run all configured security checks against a path and produce a unified
report (SARIF, Markdown, HTML, JSON) in the configured output directory.

Thorough mode also runs the intent checks and answers fixed verdict questions
(call home, auto-run, exfiltration chains, remote code, dependency sources, CI,
AI-agent manipulation, privileges, hidden content, secrets, license). Findings
are ranked into tiers (A verdict-relevant, B capability inventory, C noise);
nothing is dropped. The scan exits with status 2 when Tier A findings exist
(disable with --no-fail or intent.fail_on_tier_a: false).

--network downloads the direct, non-vendored dependencies (Go modules,
npm packages from package-lock.json) and scans their source. Nothing is
installed or executed.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := args[0]
		verbose, _ := cmd.Flags().GetBool("verbose")
		outputDir, _ := cmd.Flags().GetString("output-dir")
		mode, _ := cmd.Flags().GetString("mode")
		network, _ := cmd.Flags().GetBool("network")
		noFail, _ := cmd.Flags().GetBool("no-fail")
		network = network || cfg.Intent.Network.Enabled

		if _, err := os.Stat(target); err != nil {
			return fmt.Errorf("target path does not exist: %s", target)
		}

		// Apply mode-aware exclusions (dependency paths excluded in fast mode).
		cfg.Sanitization.ExcludedPaths = cfg.Sanitization.EffectiveExcludedPaths(mode)

		var stageNames []string
		registry := pipeline.BuiltInRegistry(cfg)
		switch mode {
		case "fast":
			stageNames = []string{"yara", "hiddenchars"}
		case "thorough":
			if err := external.RegisterAll(registry, cfg); err != nil {
				return fmt.Errorf("registering external scanners: %w", err)
			}
			stageNames = pipeline.DefaultStages()
			if cfg.Intent.Enabled {
				stageNames = append(stageNames, intent.StageNames()...)
				if network {
					stageNames = append(stageNames, "netdeps")
				}
			}
			for _, name := range registry.Names() {
				if !contains(stageNames, name) && !contains(intent.StageNames(), name) && name != "netdeps" {
					stageNames = append(stageNames, name)
				}
			}
		default:
			return fmt.Errorf("unknown mode %q; use fast or thorough", mode)
		}

		stages, err := registry.Build(stageNames)
		if err != nil {
			return fmt.Errorf("building pipeline: %w", err)
		}

		engine := pipeline.NewEngine(stages, 4)
		ctx := context.Background()
		if verbose {
			color.Cyan("Running %s pipeline with %d stage(s)...\n", mode, len(stages))
		}

		rep, err := engine.Run(ctx, target)
		if err != nil {
			return fmt.Errorf("pipeline failed: %w", err)
		}
		rep.Metadata["safeanalyze_version"] = version.Version
		rep.Metadata["scan_mode"] = mode
		triage.Assign(rep)
		triage.Finalize(rep, stageNames, coverageNotes(mode, network))

		printReportSummary(rep)

		outCfg := cfg.Output
		if outputDir != "" {
			outCfg.OutDir = outputDir
		}
		if err := report.WriteAll(rep, outCfg); err != nil {
			return fmt.Errorf("writing reports: %w", err)
		}
		if verbose {
			color.Green("Reports written to %s\n", outCfg.OutDir)
		}

		if cfg.Intent.FailOnTierA {
			if n := rep.CountTier(report.TierA); n > 0 && !noFail {
				return &ExitError{Code: ExitBlocked, Err: fmt.Errorf("blocked: %d Tier A finding(s) (verdict-relevant); see the report verdicts. First: %s",
					n, describeFindings(tierFindings(rep, report.TierA)))}
			}
			return nil
		}
		// Legacy per-stage gates, used only when the tier gate is disabled.
		if cfg.YARA.FailOnFindings && !noFail {
			if blocking := blockingFindings(rep, "yara", report.SeverityCritical); len(blocking) > 0 {
				return fmt.Errorf("blocking YARA findings detected: %s", describeFindings(blocking))
			}
		}
		if cfg.HiddenChars.FailOnFindings && !noFail {
			if blocking := blockingFindings(rep, "hiddenchars", report.SeverityHigh); len(blocking) > 0 {
				return fmt.Errorf("blocking hidden character findings detected: %s", describeFindings(blocking))
			}
		}

		return nil
	},
}

// blockingFindings returns the findings emitted by the named stage at the
// given severity, so a stage's blocking gate is not tripped by other stages.
func blockingFindings(rep *report.Report, source, severity string) []report.Finding {
	var out []report.Finding
	for _, f := range rep.Findings {
		if f.Source == source && f.Severity == severity {
			out = append(out, f)
		}
	}
	return out
}

// tierFindings returns the findings of one tier, most severe first.
func tierFindings(rep *report.Report, tier string) []report.Finding {
	var out []report.Finding
	for _, f := range rep.Findings {
		if f.Tier == tier {
			out = append(out, f)
		}
	}
	report.SortFindingsBySeverity(out)
	return out
}

// coverageNotes records the known limits of a run in the report.
func coverageNotes(mode string, network bool) []report.CoverageItem {
	notes := []report.CoverageItem{
		{Check: "dynamic-analysis", Status: "not_available", Note: "nothing is executed; behaviour is inferred statically"},
		{Check: "deobfuscation", Status: "not_available", Note: "encoded payloads are flagged but not decoded and re-scanned"},
		{Check: "reachability", Status: "not_available", Note: "no call-graph analysis: findings are not proven reachable, chains are line-proximity heuristics"},
		{Check: "intent-excluded-paths", Status: "info", Note: "not scanned by intent checks: " + strings.Join(cfg.Intent.ExcludedPaths, ", ")},
	}
	if mode == "fast" {
		notes = append(notes, report.CoverageItem{Check: "intent", Status: "not_run", Note: "fast mode runs only yara and hiddenchars"})
	}
	if !network {
		notes = append(notes, report.CoverageItem{Check: "netdeps", Status: "not_run", Note: "non-vendored dependencies were not downloaded; use --network"})
	} else {
		notes = append(notes, report.CoverageItem{Check: "netdeps-scope", Status: "info", Note: "direct Go and npm dependencies only; transitive and Python dependencies are not fetched"})
	}
	return notes
}

// describeFindings names the first few blocking findings for the error message.
func describeFindings(findings []report.Finding) string {
	const max = 3
	parts := make([]string, 0, max)
	for i, f := range findings {
		if i == max {
			break
		}
		parts = append(parts, fmt.Sprintf("%s at %s:%d", f.RuleID, f.File, f.Line))
	}
	s := strings.Join(parts, ", ")
	if len(findings) > max {
		s += fmt.Sprintf(" (+%d more)", len(findings)-max)
	}
	return s
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func printReportSummary(rep *report.Report) {
	fmt.Println()
	color.Cyan("========================================")
	color.Cyan("         SCAN SUMMARY")
	color.Cyan("========================================")
	fmt.Printf("  Target:        %s\n", rep.Target)
	fmt.Printf("  Files scanned: %d\n", rep.Summary.FilesScanned)
	fmt.Printf("  Findings:      %d (Tier A %d, Tier B %d, Tier C %d)\n", len(rep.Findings),
		rep.CountTier(report.TierA), rep.CountTier(report.TierB), rep.CountTier(report.TierC))
	if len(rep.Errors) > 0 {
		color.Yellow("  Errors:        %d\n", len(rep.Errors))
		for _, e := range rep.Errors {
			fmt.Printf("    [%s] %s\n", e.Source, e.Message)
		}
	}

	if len(rep.Verdicts) > 0 {
		color.Cyan("\nVerdicts:\n")
		for _, v := range rep.Verdicts {
			answer := strings.ToUpper(v.Answer)
			switch v.Answer {
			case report.AnswerYes:
				answer = color.RedString(answer)
			case report.AnswerUnclear:
				answer = color.YellowString(answer)
			case report.AnswerNo:
				answer = color.GreenString(answer)
			}
			fmt.Printf("  %-12s %-9s %s (A %d, B %d)\n", v.ID, answer, v.Question, v.TierA, v.TierB)
		}
	}

	tierA := tierFindings(rep, report.TierA)
	if len(tierA) > 0 {
		color.Red("\nTier A findings:\n")
		for _, f := range tierA {
			colorFn := yara.SeverityColor(f.Severity)
			fmt.Printf("  [%s] %s:%d %s\n", colorFn(f.Severity), f.File, f.Line, f.RuleID)
			fmt.Printf("    %s\n", f.Message)
		}
	} else if len(rep.Findings) == 0 {
		color.Green("\nNo findings.\n")
	}
	if n := len(rep.Findings) - len(tierA); n > 0 {
		fmt.Printf("\n  %d Tier B/C findings are listed in the report files.\n", n)
	}
}

func init() {
	scanCmd.Flags().BoolP("verbose", "v", false, "show scanner output")
	scanCmd.Flags().StringP("output-dir", "o", "", "output directory for reports")
	scanCmd.Flags().StringP("mode", "m", "thorough", "scan mode: fast or thorough")
	scanCmd.Flags().Bool("network", false, "download direct, non-vendored dependencies and scan their source (thorough mode)")
	scanCmd.Flags().Bool("no-fail", false, "report only: do not exit non-zero on Tier A or other blocking findings")
}
