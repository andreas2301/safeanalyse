package external

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/user/safeanalyze/pkg/report"
)

// trufflehogStage runs TruffleHog filesystem scanner.
type trufflehogStage struct {
	noVerification bool
	timeout        time.Duration
}

// NewTrufflehogStage creates a TruffleHog pipeline stage. When noVerification
// is true, trufflehog is run with --no-verification so results do not depend on
// network access to credential providers.
func NewTrufflehogStage(noVerification bool) *trufflehogStage {
	return &trufflehogStage{noVerification: noVerification, timeout: defaultScannerTimeout}
}

// NewTrufflehogStageWithTimeout creates a TruffleHog stage with custom timeout.
func NewTrufflehogStageWithTimeout(noVerification bool, timeout time.Duration) *trufflehogStage {
	if timeout <= 0 {
		timeout = defaultScannerTimeout
	}
	return &trufflehogStage{noVerification: noVerification, timeout: timeout}
}

// Name returns the stage name.
func (s *trufflehogStage) Name() string { return "trufflehog" }

// Dependencies returns no dependencies.
func (s *trufflehogStage) Dependencies() []string { return nil }

// Run executes trufflehog and parses NDJSON findings.
func (s *trufflehogStage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := prepareReport(input, target)
	binary := resolveBinary("trufflehog")
	if binary == "" {
		return missingBinary(out, s.Name()), nil
	}
	args := []string{"filesystem", "--json"}
	if s.noVerification {
		args = append(args, "--no-verification")
	}
	args = append(args, "--", target)

	output, stderr, _, err := runCommand(ctx, s.timeout, binary, args)
	if err != nil {
		if len(bytes.TrimSpace(output)) > 0 {
			before := len(out.Findings)
			parsed, _ := parseTrufflehogOutput(out, output)
			if len(parsed.Findings) > before {
				return parsed, nil
			}
		}
		msg := fmt.Sprintf("trufflehog execution failed: %v", err)
		if stderr != "" {
			msg += ": " + stderr
		}
		out.AddError(s.Name(), msg)
		return out, nil
	}
	return parseTrufflehogOutput(out, output)
}

func parseTrufflehogOutput(out *report.Report, data []byte) (*report.Report, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Buffer(make([]byte, 1024), 4*1024*1024)
	for scanner.Scan() {
		scanLine := scanner.Bytes()
		if len(bytes.TrimSpace(scanLine)) == 0 {
			continue
		}
		var f struct {
			DetectorName   string `json:"DetectorName"`
			Verified       bool   `json:"Verified"`
			Raw            string `json:"Raw"`
			Redacted       string `json:"Redacted"`
			SourceMetadata struct {
				Data struct {
					Filesystem struct {
						File string `json:"file"`
						Line int    `json:"line"`
					} `json:"Filesystem"`
				} `json:"Data"`
			} `json:"SourceMetadata"`
		}
		if err := json.Unmarshal(scanLine, &f); err != nil {
			out.AddError("trufflehog", fmt.Sprintf("parsing JSON line: %v", err))
			continue
		}
		file := f.SourceMetadata.Data.Filesystem.File
		line := f.SourceMetadata.Data.Filesystem.Line
		match := f.Raw
		if match == "" {
			match = f.Redacted
		}
		msg := fmt.Sprintf("Potential %s secret detected", f.DetectorName)
		if f.Verified {
			msg += " (verified)"
		}
		sev := report.SeverityHigh
		if f.Verified {
			sev = report.SeverityCritical
		}
		out.AddFinding(report.Finding{
			RuleID:      strings.ToLower(f.DetectorName),
			Title:       f.DetectorName,
			Description: msg,
			Severity:    sev,
			Category:    "secret",
			File:        file,
			Line:        line,
			Column:      0,
			Match:       match,
			Message:     msg,
			Source:      "trufflehog",
			Confidence:  report.ConfidenceDeterministic,
		})
	}
	if err := scanner.Err(); err != nil {
		out.AddError("trufflehog", fmt.Sprintf("reading NDJSON: %v", err))
	}
	return out, nil
}
