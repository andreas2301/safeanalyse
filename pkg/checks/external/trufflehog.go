package external

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/user/safeanalyze/pkg/report"
)

// trufflehogStage runs TruffleHog filesystem scanner.
type trufflehogStage struct{}

// NewTrufflehogStage creates a TruffleHog pipeline stage.
func NewTrufflehogStage() *trufflehogStage { return &trufflehogStage{} }

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
	args := []string{"filesystem", target, "--json"}
	cmd := exec.CommandContext(ctx, binary, args...)
	output, err := cmd.Output()
	if err != nil {
		out.AddError(s.Name(), fmt.Sprintf("trufflehog execution failed: %v", err))
		return out, nil
	}
	return parseTrufflehogOutput(out, output)
}

func parseTrufflehogOutput(out *report.Report, data []byte) (*report.Report, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Buffer(make([]byte, 1024), 1024*1024)
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
