package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/checks/hiddenchars"
	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/pipeline"
	"github.com/user/safeanalyze/pkg/report"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Fast stdin/file inspection for reverse-proxy integration (Squid)",
	Long: `Read one or more payloads from stdin or a file and run the fast
prompt-injection check suite. Designed for Squid external_acl helpers or
similar reverse-proxy integrations.

Reads line-by-line by default. With --body, reads the entire input as one
payload. Exits with code 1 if any finding is detected, code 0 if clean.

Example Squid external_acl helper usage:
  external_acl_type prompt_injection_check %SRC %DST %METHOD %URI /usr/local/bin/safeanalyze inspect --squid
  acl prompt_injection_detected external prompt_injection_check
  http_access deny prompt_injection_detected
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		body, _ := cmd.Flags().GetBool("body")
		squid, _ := cmd.Flags().GetBool("squid")
		verbose, _ := cmd.Flags().GetBool("verbose")

		var input string
		if len(args) > 0 {
			data, err := os.ReadFile(args[0])
			if err != nil {
				return fmt.Errorf("reading input file: %w", err)
			}
			input = string(data)
		} else {
			data, err := os.ReadFile("/dev/stdin")
			if err != nil {
				return fmt.Errorf("reading stdin: %w", err)
			}
			input = string(data)
		}

		var payloads []string
		if body {
			payloads = []string{input}
		} else {
			scanner := bufio.NewScanner(strings.NewReader(input))
			for scanner.Scan() {
				payloads = append(payloads, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("scanning input: %w", err)
			}
		}

		anyFinding := false
		for i, payload := range payloads {
			if strings.TrimSpace(payload) == "" {
				continue
			}
			start := time.Now()
			rep, err := inspectPayload(payload)
			elapsed := time.Since(start)
			if err != nil {
				return err
			}

			clean := len(rep.Findings) == 0
			if squid {
				if clean {
					fmt.Println("OK")
				} else {
					fmt.Printf("ERR %s\n", rep.Findings[0].RuleID)
				}
			} else if verbose {
				if clean {
					fmt.Printf("payload %d: clean (%.2f ms)\n", i+1, float64(elapsed.Microseconds())/1000)
				} else {
					fmt.Printf("payload %d: %d finding(s) (%.2f ms)\n", i+1, len(rep.Findings), float64(elapsed.Microseconds())/1000)
					for _, f := range rep.Findings {
						fmt.Printf("  [%s] %s: %s\n", f.Severity, f.RuleID, f.Message)
					}
				}
			} else {
				if !clean {
					fmt.Printf("payload %d: [%s] %s\n", i+1, rep.Findings[0].Severity, rep.Findings[0].Message)
				}
			}

			if !clean {
				anyFinding = true
			}
		}

		if anyFinding {
			os.Exit(1)
		}
		return nil
	},
}

func inspectPayload(payload string) (*report.Report, error) {
	stages := []pipeline.Stage{
		yara.NewStage(nil),
		hiddenchars.NewStage([]string{"zero_width", "bidi", "control"}, nil),
	}

	engine := pipeline.NewEngine(stages, 4)
	// Write payload to a temp file so existing Stage implementations can scan it.
	tmpDir, err := os.MkdirTemp("", "safeanalyze-inspect-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)
	payloadPath := tmpDir + "/payload.txt"
	if err := os.WriteFile(payloadPath, []byte(payload), 0644); err != nil {
		return nil, err
	}
	return engine.Run(context.Background(), tmpDir)
}

func init() {
	inspectCmd.Flags().Bool("body", false, "read entire input as one payload instead of line-by-line")
	inspectCmd.Flags().Bool("squid", false, "output Squid external_acl format (OK/ERR)")
	inspectCmd.Flags().BoolP("verbose", "v", false, "print per-payload results")
}
