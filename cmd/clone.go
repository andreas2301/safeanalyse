package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/utils"
)

var cloneCmd = &cobra.Command{
	Use:   "clone <url> [dir]",
	Short: "Clone a repository, then run the full safeanalyze pipeline",
	Long: `Clone a git repository into a temporary directory, run the full
scan + sanitize + ingest pipeline, and optionally delete the raw clone.

Examples:
  safeanalyze clone https://github.com/user/repo.git
  safeanalyze clone https://github.com/user/repo.git ./my-analysis
  safeanalyze clone https://github.com/user/repo.git --keep-raw`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		if err := validateGitURL(url); err != nil {
			return fmt.Errorf("invalid clone URL: %w", err)
		}

		keepRaw, _ := cmd.Flags().GetBool("keep-raw")
		runIngest, _ := cmd.Flags().GetBool("ingest")

		// Determine clone directory
		var cloneDir string
		if len(args) > 1 {
			cloneDir = args[1]
		} else {
			// Derive name from URL
			base := filepath.Base(url)
			base = strings.TrimSuffix(base, ".git")
			var err error
			cloneDir, err = utils.TempDir("safeanalyze-clone-" + base)
			if err != nil {
				return fmt.Errorf("creating temp dir: %w", err)
			}
		}

		color.Cyan("\n[CLONE] %s -> %s\n", url, cloneDir)

		// Run git clone
		gitCmd := exec.Command("git", "clone", "--depth", "1", url, cloneDir)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		if err := gitCmd.Run(); err != nil {
			return fmt.Errorf("git clone failed: %w", err)
		}

		color.Green("Clone complete.\n")

		if runIngest {
			color.Cyan("\n[PIPELINE] Running ingest on cloned repository...\n")
			outDir := filepath.Join(filepath.Dir(cloneDir), "safeanalyze-out")

			// Find our own binary path
			self, err := os.Executable()
			if err != nil {
				self = "safeanalyze"
			}

			args := []string{"ingest", cloneDir}
			if cfgFile != "" {
				args = append(args, "--config", cfgFile)
			}
			ingestCmd := exec.Command(self, args...)
			ingestCmd.Stdout = os.Stdout
			ingestCmd.Stderr = os.Stderr
			ingestCmd.Env = append(os.Environ(), "SAFEANALYZE_OUT_DIR="+outDir)
			if err := ingestCmd.Run(); err != nil {
				return fmt.Errorf("ingest failed: %w", err)
			}
		}

		if !keepRaw {
			color.Yellow("\n[CLEANUP] Removing raw clone directory...\n")
			if err := os.RemoveAll(cloneDir); err != nil {
				color.Yellow("Warning: failed to remove clone dir: %v\n", err)
			} else {
				color.Green("Raw clone removed.\n")
			}
		} else {
			color.Yellow("Raw clone kept at: %s\n", cloneDir)
		}

		return nil
	},
}

func init() {
	cloneCmd.Flags().Bool("keep-raw", false, "keep the raw cloned repository after analysis")
	cloneCmd.Flags().Bool("ingest", true, "automatically run ingest after clone")
}

// validateGitURL rejects URLs that could be interpreted as git options or
// contain shell/control characters. Permitted forms are http(s)://, ssh://,
// git://, and scp-style user@host:path URLs.
func validateGitURL(url string) error {
	if url == "" {
		return fmt.Errorf("URL is empty")
	}
	if strings.HasPrefix(url, "-") {
		return fmt.Errorf("URL must not start with '-'")
	}
	forbidden := []string{"\n", "\r", "\t", ";", "|", "&", "$", "`", "\\", "<", ">"}
	for _, c := range forbidden {
		if strings.Contains(url, c) {
			return fmt.Errorf("URL contains forbidden character %q", c)
		}
	}
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") ||
		strings.HasPrefix(url, "ssh://") || strings.HasPrefix(url, "git://") {
		return nil
	}
	// scp-style: [user@]host:path
	if idx := strings.Index(url, "@"); idx > 0 && strings.Contains(url[idx:], ":") {
		return nil
	}
	return fmt.Errorf("URL must use http(s), ssh, git, or scp-style format")
}
