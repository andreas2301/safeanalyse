package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
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
		maxConns, _ := cmd.Flags().GetInt("max-connections")
		logOnly, _ := cmd.Flags().GetBool("log-only")
		blockCTRaw, _ := cmd.Flags().GetString("block-content-types")
		adminListen, _ := cmd.Flags().GetString("admin-listen")
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
		blockContentTypes := parseContentTypeList(blockCTRaw)

		logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
		srv, err := icap.NewServer(icap.Config{
			Service:            service,
			MaxBodyBytes:       maxBody,
			Timeout:            timeout,
			FailOpen:           !failClosed,
			BlockSeverity:      blockSev,
			MaxConcurrentScans: maxScans,
			MaxConnections:     maxConns,
			LogOnly:            logOnly,
			BlockContentTypes:  blockContentTypes,
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

		var adminSrv *http.Server
		var adminAddr string
		if adminListen != "" {
			adminLn, err := net.Listen("tcp", adminListen)
			if err != nil {
				ln.Close()
				return fmt.Errorf("listening admin on %s: %w", adminListen, err)
			}
			adminAddr = adminLn.Addr().String()
			adminSrv = &http.Server{
				Handler:           srv.Metrics().Handler(srv.Ready),
				ReadHeaderTimeout: 5 * time.Second,
			}
			logger.Info("icap admin server listening", "addr", adminAddr)
			go func() {
				if err := adminSrv.Serve(adminLn); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("icap admin server error", "error", err.Error())
				}
			}()
			defer func() {
				shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				_ = adminSrv.Shutdown(shutdownCtx)
			}()
		}

		logger.Info("icap server listening", "addr", ln.Addr().String(), "service", service,
			"max_body_bytes", maxBody, "timeout", timeout.String(), "fail_open", !failClosed,
			"block_severity", blockSev, "max_concurrent_scans", maxScans,
			"max_connections", maxConns, "log_only", logOnly,
			"block_content_types", blockContentTypes, "admin_addr", adminAddr,
			"safeanalyze_version", version.Version)

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()
		go func() {
			<-ctx.Done()
			srv.Close()
			if adminSrv != nil {
				shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				_ = adminSrv.Shutdown(shutdownCtx)
			}
		}()
		return srv.Serve(ln)
	},
}

// parseContentTypeList splits a comma-separated list of media types into
// trimmed non-empty entries.
func parseContentTypeList(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	var out []string
	for _, part := range strings.Split(raw, ",") {
		if s := strings.TrimSpace(part); s != "" {
			out = append(out, s)
		}
	}
	return out
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
	icapCmd.Flags().Int("max-connections", 256, "maximum concurrent TCP connections (<0 for unlimited)")
	icapCmd.Flags().Bool("log-only", false, "scan and log findings without blocking responses (shadow mode)")
	icapCmd.Flags().String("block-content-types", "", "comma-separated media types eligible to block, e.g. text/plain,message/rfc822 (empty blocks all scannable types)")
	icapCmd.Flags().String("admin-listen", "", "optional HTTP listen address for /healthz, /readyz, /metrics, e.g. 127.0.0.1:9102")
	icapCmd.Flags().Bool("fail-open", false, "allow responses on errors/timeouts (default)")
	icapCmd.Flags().Bool("fail-closed", false, "block responses on errors/timeouts")
	icapCmd.Flags().String("block-severity", "high", "minimum finding severity that blocks (critical|high|medium|low)")
	rootCmd.AddCommand(icapCmd)
}
