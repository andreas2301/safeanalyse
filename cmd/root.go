package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/config"
	"github.com/user/safeanalyze/pkg/version"
)

var (
	cfgFile string
	cfg     *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "safeanalyze",
	Short: "Safely analyze untrusted code repositories before AI ingestion",
	Long: `safeanalyze sanitizes, scans, and prepares codebases for AI analysis.

It implements a defense-in-depth pipeline inspired by Zones of Distrust:
  1. Clone / read repository
  2. Run external security scanners (trufflehog, semgrep, etc.)
  3. Run built-in YARA-like rules for prompt injection & malware
  4. Detect hidden Unicode characters (zero-width, bidi overrides)
  5. Entropy analysis for encoded payloads
  6. Strip comments and remove non-ASCII characters (AST-aware)
  7. Output sanitized, bounded content ready for Claude/AI ingestion
  8. Launch sandboxed Claude session

Use 'safeanalyze init' to create a configuration file.
`,
	Version: version.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile != "" {
			var err error
			cfg, err = config.Load(cfgFile)
			if err != nil {
				return fmt.Errorf("loading config: %w", err)
			}
		} else {
			cfg = config.DefaultConfig()
		}
		return nil
	},
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file path")
	rootCmd.AddCommand(initCmd, installCmd, inspectCmd, scanCmd, sanitizeCmd, ingestCmd, diffCmd)
}
