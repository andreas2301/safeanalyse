package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/diff"
)

var diffCmd = &cobra.Command{
	Use:   "diff <original> <sanitized>",
	Short: "Show what was removed during sanitization",
	Long: `Compare an original file or directory against its sanitized version
and display a colored diff of what was removed (comments, non-ASCII chars, etc.).

Examples:
  safeanalyze diff ./repo/main.go ./safeanalyze-out/sanitized/main.go
  safeanalyze diff ./repo ./safeanalyze-out/sanitized --unified`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		origPath := args[0]
		sanPath := args[1]
		unified, _ := cmd.Flags().GetBool("unified")
		contextLines, _ := cmd.Flags().GetInt("context")

		origInfo, err := os.Stat(origPath)
		if err != nil {
			return fmt.Errorf("original path not found: %w", err)
		}
		sanInfo, err := os.Stat(sanPath)
		if err != nil {
			return fmt.Errorf("sanitized path not found: %w", err)
		}

		if origInfo.IsDir() != sanInfo.IsDir() {
			return fmt.Errorf("both paths must be files or both directories")
		}

		if !origInfo.IsDir() {
			// Single file diff
			return diffFile(origPath, sanPath, unified, contextLines)
		}

		// Directory diff
		return diffDir(origPath, sanPath, unified, contextLines)
	},
}

func diffFile(origPath, sanPath string, unified bool, contextLines int) error {
	origData, err := os.ReadFile(origPath)
	if err != nil {
		return err
	}
	sanData, err := os.ReadFile(sanPath)
	if err != nil {
		return err
	}

	changes := diff.Diff(string(origData), string(sanData))
	added, removed, unchanged := diff.Stats(changes)

	fmt.Printf("\n%s\n", color.CyanString("=== %s ===", filepath.Base(origPath)))
	fmt.Printf("Lines: %d added, %d removed, %d unchanged\n\n", added, removed, unchanged)

	if unified {
		fmt.Println(diff.UnifiedDiff(changes, origPath, sanPath))
	} else {
		fmt.Println(diff.Format(changes, contextLines))
	}

	return nil
}

func diffDir(origDir, sanDir string, unified bool, contextLines int) error {
	var totalFiles int
	var filesWithChanges int

	err := filepath.WalkDir(origDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(origDir, path)
		if err != nil {
			return nil
		}

		sanFile := filepath.Join(sanDir, rel)
		if _, err := os.Stat(sanFile); os.IsNotExist(err) {
			color.Yellow("  [REMOVED] %s (not present in sanitized output)\n", rel)
			filesWithChanges++
			totalFiles++
			return nil
		}

		origData, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		sanData, err := os.ReadFile(sanFile)
		if err != nil {
			return nil
		}

		if string(origData) == string(sanData) {
			return nil // identical
		}

		filesWithChanges++
		totalFiles++

		changes := diff.Diff(string(origData), string(sanData))
		added, removed, _ := diff.Stats(changes)

		fmt.Printf("\n%s  %s\n", color.CyanString("▶"), color.WhiteString(rel))
		fmt.Printf("  %s %d added, %d removed\n", color.HiBlackString("│"), added, removed)

		if unified {
			udiff := diff.UnifiedDiff(changes, rel, rel)
			// Indent unified diff
			for _, line := range strings.Split(udiff, "\n") {
				if line == "" {
					continue
				}
				fmt.Printf("  %s %s\n", color.HiBlackString("│"), line)
			}
		} else {
			for _, c := range changes {
				if c.Type == diff.Equal {
					continue
				}
				prefix := color.HiBlackString("│")
				if c.Type == diff.Added {
					fmt.Printf("  %s %s\n", prefix, color.GreenString("+ %s", c.Text))
				} else {
					fmt.Printf("  %s %s\n", prefix, color.RedString("- %s", c.Text))
				}
			}
		}

		return nil
	})

	fmt.Printf("\n%s\n", color.CyanString("========================================"))
	fmt.Printf("Files analyzed: %d\n", totalFiles)
	fmt.Printf("Files with changes: %d\n", filesWithChanges)

	return err
}

func init() {
	diffCmd.Flags().BoolP("unified", "u", false, "show unified diff format")
	diffCmd.Flags().IntP("context", "C", 0, "show N lines of context around changes")
	rootCmd.AddCommand(diffCmd)
}
