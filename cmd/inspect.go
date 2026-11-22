package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/fastinspect"
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

Note: Squid passes only the format fields (client, host, method, URL) to an
external_acl helper, never the response body, so this setup scans request
metadata only. To scan HTTP response bodies, run 'safeanalyze icap' and
attach it to Squid as an ICAP RESPMOD service.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		body, _ := cmd.Flags().GetBool("body")
		squid, _ := cmd.Flags().GetBool("squid")
		verbose, _ := cmd.Flags().GetBool("verbose")

		out := bufio.NewWriter(os.Stdout)
		defer out.Flush()

		anyFinding := false
		handlePayload := func(idx int, payload string) error {
			if strings.TrimSpace(payload) == "" {
				return nil
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
					fmt.Fprintln(out, "OK")
				} else {
					fmt.Fprintf(out, "ERR %s\n", rep.Findings[0].RuleID)
				}
			} else if verbose {
				if clean {
					fmt.Fprintf(out, "payload %d: clean (%.2f ms)\n", idx, float64(elapsed.Microseconds())/1000)
				} else {
					fmt.Fprintf(out, "payload %d: %d finding(s) (%.2f ms)\n", idx, len(rep.Findings), float64(elapsed.Microseconds())/1000)
					for _, f := range rep.Findings {
						fmt.Fprintf(out, "  [%s] %s: %s\n", f.Severity, f.RuleID, f.Message)
					}
				}
			} else {
				if !clean {
					fmt.Fprintf(out, "payload %d: [%s] %s\n", idx, rep.Findings[0].Severity, rep.Findings[0].Message)
				}
			}
			if err := out.Flush(); err != nil {
				return fmt.Errorf("flushing output: %w", err)
			}

			if !clean {
				anyFinding = true
			}
			return nil
		}

		if body {
			var data []byte
			var err error
			if len(args) > 0 {
				data, err = os.ReadFile(args[0])
				if err != nil {
					return fmt.Errorf("reading input file: %w", err)
				}
			} else {
				data, err = io.ReadAll(os.Stdin)
				if err != nil {
					return fmt.Errorf("reading stdin: %w", err)
				}
			}
			if err := handlePayload(1, string(data)); err != nil {
				return err
			}
		} else {
			var r io.Reader
			if len(args) > 0 {
				f, err := os.Open(args[0])
				if err != nil {
					return fmt.Errorf("reading input file: %w", err)
				}
				defer f.Close()
				r = f
			} else {
				r = os.Stdin
			}

			scanner := bufio.NewScanner(r)
			scanner.Buffer(make([]byte, 0, 64*1024), 4<<20)
			lineNum := 0
			for scanner.Scan() {
				lineNum++
				if err := handlePayload(lineNum, scanner.Text()); err != nil {
					return err
				}
			}
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("scanning input: %w", err)
			}
		}

		if anyFinding {
			os.Exit(1)
		}
		return nil
	},
}

func inspectPayload(payload string) (*report.Report, error) {
	return fastinspect.Default().Scan(payload), nil
}

func init() {
	inspectCmd.Flags().Bool("body", false, "read entire input as one payload instead of line-by-line")
	inspectCmd.Flags().Bool("squid", false, "output Squid external_acl format (OK/ERR)")
	inspectCmd.Flags().BoolP("verbose", "v", false, "print per-payload results")
}
