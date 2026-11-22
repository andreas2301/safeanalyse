// Package external wraps third-party security scanners and declarative plugins
// as pipeline stages.
package external

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/user/safeanalyze/pkg/install"
	"github.com/user/safeanalyze/pkg/report"
)

// defaultScannerTimeout bounds external scanner execution when no per-scanner
// timeout_seconds is configured.
const defaultScannerTimeout = 5 * time.Minute

// effectiveTimeout returns d when positive, else defaultScannerTimeout.
func effectiveTimeout(seconds int) time.Duration {
	if seconds > 0 {
		return time.Duration(seconds) * time.Second
	}
	return defaultScannerTimeout
}

// runCommand executes binary with args under a bounded timeout and captures
// stdout and stderr separately so diagnostics on stderr never corrupt JSON
// output on stdout.
func runCommand(ctx context.Context, timeout time.Duration, binary string, args []string) (stdout []byte, stderr string, exitCode int, err error) {
	if timeout <= 0 {
		timeout = defaultScannerTimeout
	}
	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(runCtx, binary, args...)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	runErr := cmd.Run()
	stderr = strings.TrimSpace(errBuf.String())
	if runErr != nil {
		if errors.Is(runCtx.Err(), context.DeadlineExceeded) {
			return outBuf.Bytes(), stderr, -1, fmt.Errorf("timed out after %s", timeout)
		}
		var exitErr *exec.ExitError
		if errors.As(runErr, &exitErr) {
			return outBuf.Bytes(), stderr, exitErr.ExitCode(), runErr
		}
		return outBuf.Bytes(), stderr, -1, runErr
	}
	return outBuf.Bytes(), stderr, 0, nil
}

// splitArgs splits a command template into argv tokens supporting single and
// double quotes without invoking a shell (preventing shell metacharacter
// injection).
func splitArgs(s string) ([]string, error) {
	var args []string
	var cur strings.Builder
	inToken := false
	var quote rune
	escaped := false

	for _, r := range s {
		switch {
		case escaped:
			cur.WriteRune(r)
			escaped = false
			inToken = true
		case r == '\\' && quote != '\'':
			escaped = true
			inToken = true
		case quote != 0:
			if r == quote {
				quote = 0
			} else {
				cur.WriteRune(r)
			}
		case r == '\'' || r == '"':
			quote = r
			inToken = true
		case r == ' ' || r == '\t' || r == '\n' || r == '\r':
			if inToken {
				args = append(args, cur.String())
				cur.Reset()
				inToken = false
			}
		default:
			cur.WriteRune(r)
			inToken = true
		}
	}
	if escaped {
		return nil, fmt.Errorf("unterminated escape in command")
	}
	if quote != 0 {
		return nil, fmt.Errorf("unterminated quote %q in command", quote)
	}
	if inToken {
		args = append(args, cur.String())
	}
	return args, nil
}

// prepareReport returns a usable report, reusing input if provided.
func prepareReport(input *report.Report, target string) *report.Report {
	if input != nil {
		return input
	}
	return report.NewReport(target)
}

// resolveBinary finds the scanner binary on PATH or in the safeanalyze bin directory.
func resolveBinary(name string) string {
	if path, err := exec.LookPath(name); err == nil {
		return path
	}
	if installed := install.ResolveBinary(name); installed != "" {
		return installed
	}
	if runtime.GOOS == "windows" {
		if path, err := exec.LookPath(name + ".exe"); err == nil {
			return path
		}
	}
	return ""
}

// missingBinary records a non-fatal error and returns the report.
func missingBinary(out *report.Report, name string) *report.Report {
	out.AddError(name, fmt.Sprintf("scanner binary %q not found in PATH or safeanalyze bin", name))
	return out
}

// coalesce returns the first non-empty string.
func coalesce(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// mapSeverity normalizes external severity strings to report severity constants.
func mapSeverity(s string) string {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "critical", "crit", "severe":
		return report.SeverityCritical
	case "high", "error":
		return report.SeverityHigh
	case "medium", "moderate", "warn", "warning":
		return report.SeverityMedium
	case "low", "note":
		return report.SeverityLow
	default:
		return report.SeverityInfo
	}
}
