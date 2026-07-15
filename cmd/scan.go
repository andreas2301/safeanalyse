package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/checks/external"
	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/pipeline"
	"github.com/user/safeanalyze/pkg/report"
	"github.com/user/safeanalyze/pkg/version"
)

var scanCmd = &cobra.Command{
	Use:   "scan <path>",
	Short: "Run security scanners and emit a structured report",
	Long: `Run all configured security checks against a path and produce a unified
report (SARIF, Markdown, HTML) in the configured output directory.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := args[0]
		verbose, _ := cmd.Flags().GetBool("verbose")
		outputDir, _ := cmd.Flags().GetString("output-dir")
		mode, _ := cmd.Flags().GetString("mode")

		if _, err := os.Stat(target); err != nil {
			return fmt.Errorf("target path does not exist: %s", target)
		}

		// Apply mode-aware exclusions (dependency paths excluded in fast mode).
		cfg.Sanitization.ExcludedPaths = cfg.Sanitization.EffectiveExcludedPaths(mode)

		var stageNames []string
		switch mode {
		case "fast":
			stageNames = []string{"yara", "hiddenchars"}
		case "thorough":
			registry := pipeline.BuiltInRegistry(cfg)
			if err := external.RegisterAll(registry, cfg); err != nil {
				return fmt.Errorf("registering external scanners: %w", err)
			}
			stageNames = pipeline.DefaultStages()
			for _, name := range registry.Names() {
				if !contains(stageNames, name) {
					stageNames = append(stageNames, name)
				}
			}
		default:
			return fmt.Errorf("unknown mode %q; use fast or thorough", mode)
		}

		registry := pipeline.BuiltInRegistry(cfg)
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

		if cfg.YARA.FailOnFindings && rep.HasSeverity(report.SeverityCritical) {
			return fmt.Errorf("blocking YARA findings detected")
		}
		if cfg.HiddenChars.FailOnFindings && rep.HasSeverity(report.SeverityHigh) {
			return fmt.Errorf("blocking hidden character findings detected")
		}

		return nil
	},
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
	fmt.Printf("  Findings:      %d\n", len(rep.Findings))
	for sev, count := range rep.Summary.CountsBySeverity {
		fmt.Printf("    %s: %d\n", sev, count)
	}
	if len(rep.Errors) > 0 {
		color.Yellow("  Errors:        %d\n", len(rep.Errors))
		for _, e := range rep.Errors {
			fmt.Printf("    [%s] %s\n", e.Source, e.Message)
		}
	}

	if len(rep.Findings) > 0 {
		color.Red("\nFindings:\n")
		for _, f := range rep.Findings {
			colorFn := yara.SeverityColor(f.Severity)
			fmt.Printf("  [%s] %s:%d:%d %s\n", colorFn(f.Severity), f.File, f.Line, f.Column, f.RuleID)
			fmt.Printf("    %s\n", f.Message)
		}
	} else {
		color.Green("\nNo findings.\n")
	}
}

func init() {
	scanCmd.Flags().BoolP("verbose", "v", false, "show scanner output")
	scanCmd.Flags().StringP("output-dir", "o", "", "output directory for reports")
	scanCmd.Flags().StringP("mode", "m", "thorough", "scan mode: fast or thorough")
}
