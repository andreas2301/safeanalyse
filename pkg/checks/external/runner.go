// Package external wraps third-party security scanners as pipeline stages.
package external

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/user/safeanalyze/pkg/install"
	"github.com/user/safeanalyze/pkg/report"
)

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
	switch strings.ToLower(s) {
	case "critical", "crit", "severe":
		return report.SeverityCritical
	case "high", "error":
		return report.SeverityHigh
	case "medium", "warn", "warning":
		return report.SeverityMedium
	case "low":
		return report.SeverityLow
	default:
		return report.SeverityInfo
	}
}
