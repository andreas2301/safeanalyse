package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/hiddenchars"
	"github.com/user/safeanalyze/pkg/ingest"
	"github.com/user/safeanalyze/pkg/sanitize"
	"github.com/user/safeanalyze/pkg/scanner"
	"github.com/user/safeanalyze/pkg/sandbox"
	"github.com/user/safeanalyze/pkg/utils"
)

var ingestCmd = &cobra.Command{
	Use:   "ingest <path>",
	Short: "Full pipeline: scan + sanitize + format for AI ingestion",
	Long: `Run the complete safeanalyze pipeline:
  1. Security scanners (trufflehog, semgrep, etc.)
  2. Hidden Unicode character detection
  3. Sanitization (strip comments, remove non-ASCII, enforce limits)
  4. Format output (markdown, JSON, or plain text)
  5. Optional: launch sandboxed Claude session

This is the primary command for preparing an untrusted repo for Claude analysis.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		src := args[0]
		verbose, _ := cmd.Flags().GetBool("verbose")
		noScan, _ := cmd.Flags().GetBool("no-scan")
		launchSandbox, _ := cmd.Flags().GetBool("sandbox")

		if _, err := os.Stat(src); err != nil {
			return fmt.Errorf("source path does not exist: %s", src)
		}

		// Respect env override for output dir (used by clone command)
		if envOut := os.Getenv("SAFEANALYZE_OUT_DIR"); envOut != "" {
			cfg.Output.OutDir = envOut
		}

		outDir := cfg.Output.OutDir
		if outDir == "" {
			outDir = "./safeanalyze-out"
		}
		sanitizedDir := filepath.Join(outDir, "sanitized")

		// Phase 1: Scan (unless skipped)
		if !noScan {
			color.Cyan("\n========== PHASE 1: SECURITY SCAN ==========\n")
			runner := scanner.NewRunner(cfg.Scanners)
			results, _ := runner.RunAll(src, verbose)
			scanner.PrintSummary(results)

			if cfg.HiddenChars.Enabled {
				color.Cyan("\n--- Hidden Character Detection ---\n")
				det := hiddenchars.NewDetector(cfg.HiddenChars.Categories)
				findings, err := det.ScanDir(src, cfg.Sanitization.ExcludedPaths)
				if err != nil {
					return fmt.Errorf("hidden char scan: %w", err)
				}
				if len(findings) > 0 {
					color.Red("Found %d hidden character occurrence(s):\n", len(findings))
					for _, f := range findings {
						fmt.Printf("  %s:%d:%d %s (%s)\n", f.File, f.Line, f.Column, f.CodePoint, f.Name)
					}
					if cfg.HiddenChars.FailOnFindings {
						return fmt.Errorf("hidden characters detected; aborting")
					}
				} else {
					color.Green("No hidden characters detected.\n")
				}
			}
		}

		// Phase 2: Sanitize
		color.Cyan("\n========== PHASE 2: SANITIZATION ==========\n")
		opts := sanitize.Options{
			StripComments:     cfg.Sanitization.StripComments,
			RemoveNonASCII:    cfg.Sanitization.RemoveNonASCII,
			MaxFileSizeBytes:  cfg.Sanitization.MaxFileSizeBytes,
			MaxLinesPerFile:   cfg.Sanitization.MaxLinesPerFile,
			AllowedExtensions: cfg.Sanitization.AllowedExtensions,
			ExcludedPaths:     cfg.Sanitization.ExcludedPaths,
		}

		res, err := sanitize.SanitizeDir(src, sanitizedDir, opts)
		if err != nil {
			return fmt.Errorf("sanitization failed: %w", err)
		}

		fmt.Printf("  Files processed: %d\n", res.FilesProcessed)
		fmt.Printf("  Files skipped:   %d\n", res.FilesSkipped)
		fmt.Printf("  Bytes before:    %s\n", utils.FormatBytes(res.BytesBefore))
		fmt.Printf("  Bytes after:     %s\n", utils.FormatBytes(res.BytesAfter))
		fmt.Printf("  Comments stripped: %d files\n", res.CommentsStripped)
		fmt.Printf("  Non-ASCII removed: %d chars\n", res.NonASCIIRemoved)

		// Phase 3: Format for ingestion
		color.Cyan("\n========== PHASE 3: FORMAT FOR AI INGESTION ==========\n")
		formatter := ingest.NewFormatter(cfg.Output)
		records, err := formatter.Run(sanitizedDir)
		if err != nil {
			return fmt.Errorf("formatting failed: %w", err)
		}

		ingestDir := filepath.Join(outDir, "ingest")
		if err := formatter.WriteOutput(records, ingestDir); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}

		color.Green("\n========================================")
		color.Green("  INGESTION PACKAGE READY")
		color.Green("========================================")
		fmt.Printf("  Sanitized code:  %s\n", sanitizedDir)
		fmt.Printf("  AI ingest files: %s\n", ingestDir)
		fmt.Printf("  Format:          %s\n", cfg.Output.Format)
		fmt.Printf("  Total files:     %d\n", len(records))
		fmt.Println()

		// Phase 4: Optional sandbox launch
		if launchSandbox {
			color.Cyan("========== PHASE 4: SANDBOX LAUNCH ==========\n")
			sb := sandbox.NewRunner(sandbox.Config{
				Mode:            cfg.Sandbox.Mode,
				DockerImage:     cfg.Sandbox.DockerImage,
				FirejailProfile: cfg.Sandbox.FirejailProfile,
				NetworkDisabled: true,
				ReadOnlyPaths:   []string{ingestDir},
			})

			if !sb.Available() {
				color.Yellow("Sandbox mode %q is not available on this system.\n", cfg.Sandbox.Mode)
				color.Yellow("Available modes: docker=%v, firejail=%v\n",
					sandbox.NewRunner(sandbox.Config{Mode: "docker"}).Available(),
					sandbox.NewRunner(sandbox.Config{Mode: "firejail"}).Available())
				color.Yellow("Run 'safeanalyze ingest <path> --sandbox-mode docker' to change mode.\n")
			} else {
				cmd, err := sb.LaunchClaude(ingestDir)
				if err != nil {
					return fmt.Errorf("sandbox launch failed: %w", err)
				}
				color.Green("Launching Claude in sandbox...\n")
				color.HiBlack("Command: %s\n", cmd.String())
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					return fmt.Errorf("sandbox process exited: %w", err)
				}
			}
		}

		color.Yellow("Next steps:")
		fmt.Println("  1. Review the sanitized output in:", sanitizedDir)
		fmt.Println("  2. Feed the ingest files to Claude:", ingestDir)
		if !launchSandbox {
			fmt.Println("  3. Run with --sandbox to launch Claude in an isolated container")
		}
		fmt.Println()

		return nil
	},
}

func init() {
	ingestCmd.Flags().BoolP("verbose", "v", false, "show scanner output")
	ingestCmd.Flags().Bool("no-scan", false, "skip security scanners (sanitization only)")
	ingestCmd.Flags().Bool("sandbox", false, "launch Claude in sandbox after ingestion")
}
