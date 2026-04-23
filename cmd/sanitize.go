package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/sanitize"
	"github.com/user/safeanalyze/pkg/utils"
)

var sanitizeCmd = &cobra.Command{
	Use:   "sanitize <src> [dst]",
	Short: "Sanitize a repository for AI ingestion",
	Long: `Copy a repository, strip comments, remove non-ASCII characters,
enforce size limits, and output a cleaned version ready for AI analysis.`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		src := args[0]
		dst := "./safeanalyze-out/sanitized"
		if len(args) > 1 {
			dst = args[1]
		}

		if _, err := os.Stat(src); err != nil {
			return fmt.Errorf("source path does not exist: %s", src)
		}

		color.Cyan("\n[SANITIZE] %s -> %s\n", src, dst)

		opts := sanitize.Options{
			StripComments:     cfg.Sanitization.StripComments,
			RemoveNonASCII:    cfg.Sanitization.RemoveNonASCII,
			MaxFileSizeBytes:  cfg.Sanitization.MaxFileSizeBytes,
			MaxLinesPerFile:   cfg.Sanitization.MaxLinesPerFile,
			AllowedExtensions: cfg.Sanitization.AllowedExtensions,
			ExcludedPaths:     cfg.Sanitization.ExcludedPaths,
		}

		res, err := sanitize.SanitizeDir(src, dst, opts)
		if err != nil {
			return fmt.Errorf("sanitization failed: %w", err)
		}

		fmt.Printf("  Files processed: %d\n", res.FilesProcessed)
		fmt.Printf("  Files skipped:   %d\n", res.FilesSkipped)
		fmt.Printf("  Bytes before:    %s\n", utils.FormatBytes(res.BytesBefore))
		fmt.Printf("  Bytes after:     %s\n", utils.FormatBytes(res.BytesAfter))
		fmt.Printf("  Comments stripped from: %d files\n", res.CommentsStripped)
		fmt.Printf("  Non-ASCII chars removed: %d\n", res.NonASCIIRemoved)
		if len(res.Errors) > 0 {
			color.Yellow("  Errors: %d\n", len(res.Errors))
			for _, e := range res.Errors[:min(5, len(res.Errors))] {
				fmt.Printf("    - %s\n", e)
			}
		}

		color.Green("\nSanitized output written to: %s\n", dst)
		return nil
	},
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
