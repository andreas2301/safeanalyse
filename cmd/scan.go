package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/entropy"
	"github.com/user/safeanalyze/pkg/hiddenchars"
	"github.com/user/safeanalyze/pkg/scanner"
	"github.com/user/safeanalyze/pkg/yara"
)

var scanCmd = &cobra.Command{
	Use:   "scan <path>",
	Short: "Run security scanners, YARA rules, entropy analysis, and hidden-character detection",
	Long:  `Run all configured external scanners, built-in YARA-like rules, entropy analysis, and detect invisible Unicode characters in the repository.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := args[0]
		verbose, _ := cmd.Flags().GetBool("verbose")

		// Verify target exists
		if _, err := os.Stat(target); err != nil {
			return fmt.Errorf("target path does not exist: %s", target)
		}

		var hasBlockingIssue bool

		// Phase 1: External scanners
		color.Cyan("\n[Phase 1] Running external scanners...\n")
		runner := scanner.NewRunner(cfg.Scanners)
		results, _ := runner.RunAll(target, verbose)
		scanner.PrintSummary(results)

		// Phase 2: Built-in YARA rules
		if cfg.YARA.Enabled {
			color.Cyan("\n[Phase 2] Running built-in YARA-like rules...\n")
			engine := yara.NewEngine()
			var allMatches []yara.Match

			walkErr := filepath.WalkDir(target, func(path string, d os.DirEntry, err error) error {
				if err != nil || d.IsDir() {
					return nil
				}
				content, err := os.ReadFile(path)
				if err != nil {
					return nil
				}
				rel, _ := filepath.Rel(target, path)
				matches := engine.ScanFile(string(content), rel)
				allMatches = append(allMatches, matches...)
				return nil
			})
			if walkErr != nil {
				color.Yellow("Warning: YARA walk error: %v\n", walkErr)
			}

			if len(allMatches) == 0 {
				color.Green("No YARA rule matches.\n")
			} else {
				color.Red("Found %d YARA rule match(es):\n", len(allMatches))
				for _, m := range allMatches {
					colorFn := yara.SeverityColor(m.Severity)
					fmt.Printf("  [%s] %s:%d:%d %q\n", colorFn(m.Severity), m.File, m.Line, m.Column, m.Match)
					fmt.Printf("    Rule: %s (%s)\n", m.Rule, m.Description)
				}
				if cfg.YARA.FailOnFindings {
					hasBlockingIssue = true
				}
			}
		}

		// Phase 3: Entropy analysis
		if cfg.Entropy.Enabled {
			color.Cyan("\n[Phase 3] Running entropy analysis...\n")
			det := entropy.NewDetector().WithThreshold(cfg.Entropy.Threshold)
			var allFindings []entropy.Finding

			walkErr := filepath.WalkDir(target, func(path string, d os.DirEntry, err error) error {
				if err != nil || d.IsDir() {
					return nil
				}
				content, err := os.ReadFile(path)
				if err != nil {
					return nil
				}
				rel, _ := filepath.Rel(target, path)
				findings := det.ScanContent(string(content), rel)
				allFindings = append(allFindings, findings...)
				return nil
			})
			if walkErr != nil {
				color.Yellow("Warning: entropy walk error: %v\n", walkErr)
			}

			if len(allFindings) == 0 {
				color.Green("No high-entropy strings detected.\n")
			} else {
				color.Yellow("Found %d suspicious string(s):\n", len(allFindings))
				for _, f := range allFindings {
					fmt.Printf("  %s:%d [%s] entropy=%.2f len=%d %q\n", f.File, f.Line, f.Type, f.Entropy, f.Length, f.Value)
				}
				if cfg.Entropy.FailOnFindings {
					hasBlockingIssue = true
				}
			}
		}

		// Phase 4: Hidden character detection
		if cfg.HiddenChars.Enabled {
			color.Cyan("\n[Phase 4] Scanning for hidden Unicode characters...\n")
			det := hiddenchars.NewDetector(cfg.HiddenChars.Categories)
			findings, err := det.ScanDir(target, cfg.Sanitization.ExcludedPaths)
			if err != nil {
				return fmt.Errorf("hidden char scan failed: %w", err)
			}

			if len(findings) == 0 {
				color.Green("No hidden characters detected.\n")
			} else {
				color.Red("Found %d hidden character occurrence(s):\n", len(findings))
				for _, f := range findings {
					fmt.Printf("  %s:%d:%d %s (%s) [%s] context: %q\n",
						f.File, f.Line, f.Column, f.CodePoint, f.Name, f.Category, f.Context)
				}
				if cfg.HiddenChars.FailOnFindings {
					hasBlockingIssue = true
				}
			}
		}

		fmt.Println()
		if hasBlockingIssue {
			color.Red("Scan completed with BLOCKING issues. Review output before ingestion.\n")
			return fmt.Errorf("blocking findings detected")
		}
		color.Green("Scan completed successfully.\n")
		return nil
	},
}

func init() {
	scanCmd.Flags().BoolP("verbose", "v", false, "show scanner output")
}
