package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/icap"
	"github.com/user/safeanalyze/pkg/report"
	"github.com/user/safeanalyze/pkg/version"
)

var icapCmd = &cobra.Command{
	Use:   "icap",
	Short: "ICAP RESPMOD server that scans HTTP response bodies (Squid)",
	Long: `Run an ICAP (RFC 3507) RESPMOD service that scans HTTP response bodies with
the fast check suite (yara + hiddenchars, same as 'inspect --body' and 'eval').
Each body is scanned as one payload. Clean or non-text responses get
"204 No Content"; a response with a finding at or above --block-severity is
replaced by an HTTP 403 naming the rule IDs (the payload is never echoed).

gzip/deflate bodies are decoded for scanning (output capped at
--max-body-bytes). Bodies larger than --max-body-bytes are scanned on their
first N bytes. At most --max-concurrent-scans scans run at once; a request
that waits for a free slot spends its --timeout budget waiting. On read
errors, decode errors, a scan or slot wait exceeding --timeout, or (for
clients that do not send "Allow: 204") a body larger than --max-body-bytes,
the response is allowed (--fail-open, default) or blocked (--fail-closed).
One JSON log line per request is written to stderr.

Squid configuration:
  icap_enable on
  icap_service safeanalyze_resp respmod_precache bypass=on icap://127.0.0.1:1344/safeanalyze
  adaptation_access safeanalyze_resp allow all
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		listen, _ := cmd.Flags().GetString("listen")
		service, _ := cmd.Flags().GetString("service")
		maxBody, _ := cmd.Flags().GetInt("max-body-bytes")
		timeout, _ := cmd.Flags().GetDuration("timeout")
		failOpen, _ := cmd.Flags().GetBool("fail-open")
		failClosed, _ := cmd.Flags().GetBool("fail-closed")
		blockSev, _ := cmd.Flags().GetString("block-severity")
		maxScans, _ := cmd.Flags().GetInt("max-concurrent-scans")
		if maxScans < 1 {
			return fmt.Errorf("--max-concurrent-scans must be >= 1")
		}

		if failOpen && failClosed {
			return fmt.Errorf("--fail-open and --fail-closed are mutually exclusive")
		}
		blockSev = strings.ToLower(blockSev)
		if report.SeverityRank(blockSev) == 0 {
			return fmt.Errorf("invalid --block-severity %q (critical, high, medium or low)", blockSev)
		}

		logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
		srv, err := icap.NewServer(icap.Config{
			Service:            service,
			MaxBodyBytes:       maxBody,
			Timeout:            timeout,
			FailOpen:           !failClosed,
			BlockSeverity:      blockSev,
			MaxConcurrentScans: maxScans,
			ISTag:              "safeanalyze-" + version.Version,
			Scan:               icapScan,
			Logger:             logger,
		})
		if err != nil {
			return err
		}
		ln, err := net.Listen("tcp", listen)
		if err != nil {
			return fmt.Errorf("listening on %s: %w", listen, err)
		}
		logger.Info("icap server listening", "addr", ln.Addr().String(), "service", service,
			"max_body_bytes", maxBody, "timeout", timeout.String(), "fail_open", !failClosed,
			"block_severity", blockSev, "max_concurrent_scans", maxScans,
			"safeanalyze_version", version.Version)

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()
		go func() {
			<-ctx.Done()
			srv.Close()
		}()
		return srv.Serve(ln)
	},
}

// icapScan runs the same fast-mode detection path as inspect and eval.
func icapScan(_ context.Context, payload []byte) ([]report.Finding, error) {
	rep, err := inspectPayload(string(payload))
	if err != nil {
		return nil, err
	}
	return rep.Findings, nil
}

func init() {
	icapCmd.Flags().String("listen", "127.0.0.1:1344", "address to listen on")
	icapCmd.Flags().String("service", "/safeanalyze", "ICAP service path")
	icapCmd.Flags().Int("max-body-bytes", 2<<20, "scan at most this many (decoded) body bytes per response")
	icapCmd.Flags().Duration("timeout", 2*time.Second, "scan time budget per response (includes waiting for a scan slot)")
	icapCmd.Flags().Int("max-concurrent-scans", runtime.NumCPU()*2, "maximum scans in flight, including timed-out scans still running")
	icapCmd.Flags().Bool("fail-open", false, "allow responses on errors/timeouts (default)")
	icapCmd.Flags().Bool("fail-closed", false, "block responses on errors/timeouts")
	icapCmd.Flags().String("block-severity", "high", "minimum finding severity that blocks (critical|high|medium|low)")
	rootCmd.AddCommand(icapCmd)
}
