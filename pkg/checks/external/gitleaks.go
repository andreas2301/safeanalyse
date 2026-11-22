package external

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/user/safeanalyze/pkg/report"
)

// gitleaksStage runs Gitleaks against a directory.
type gitleaksStage struct {
	timeout time.Duration
}

// NewGitleaksStage creates a Gitleaks pipeline stage.
func NewGitleaksStage() *gitleaksStage { return &gitleaksStage{timeout: defaultScannerTimeout} }

// NewGitleaksStageWithTimeout creates a Gitleaks stage with a custom timeout.
func NewGitleaksStageWithTimeout(timeout time.Duration) *gitleaksStage {
	if timeout <= 0 {
		timeout = defaultScannerTimeout
	}
	return &gitleaksStage{timeout: timeout}
}

// Name returns the stage name.
func (s *gitleaksStage) Name() string { return "gitleaks" }

// Dependencies returns no dependencies.
func (s *gitleaksStage) Dependencies() []string { return nil }

// Run executes gitleaks and parses JSON findings. Gitleaks exits with status 1
// when leaks are found, so exit code 1 with valid JSON stdout is treated as a
// normal findings result rather than an execution failure.
func (s *gitleaksStage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := prepareReport(input, target)
	binary := resolveBinary("gitleaks")
	if binary == "" {
		return missingBinary(out, s.Name()), nil
	}
	args := []string{"dir", "--report-format", "json", "--report-path", "-", "--", target}
	output, stderr, exitCode, err := runCommand(ctx, s.timeout, binary, args)
	if err != nil {
		// Exit code 1 means leaks detected; parse stdout if present.
		if exitCode == 1 && len(bytes.TrimSpace(output)) > 0 {
			return parseGitleaksOutput(out, output)
		}
		msg := fmt.Sprintf("gitleaks execution failed: %v", err)
		if stderr != "" {
			msg += ": " + stderr
		}
		out.AddError(s.Name(), msg)
		return out, nil
	}
	if len(bytes.TrimSpace(output)) == 0 {
		return out, nil
	}
	return parseGitleaksOutput(out, output)
}

func parseGitleaksOutput(out *report.Report, data []byte) (*report.Report, error) {
	var findings []struct {
		RuleID      string   `json:"RuleID"`
		Description string   `json:"Description"`
		StartLine   int      `json:"StartLine"`
		StartColumn int      `json:"StartColumn"`
		Match       string   `json:"Match"`
		Secret      string   `json:"Secret"`
		File        string   `json:"File"`
		Tags        []string `json:"Tags"`
	}
	if err := json.Unmarshal(data, &findings); err != nil {
		out.AddError("gitleaks", fmt.Sprintf("parsing JSON: %v", err))
		return out, nil
	}
	for _, f := range findings {
		sev := report.SeverityHigh
		for _, tag := range f.Tags {
			if tag == "verified" {
				sev = report.SeverityCritical
				break
			}
		}
		match := f.Secret
		if match == "" {
			match = f.Match
		}
		// Never embed the raw secret in Message; Match is redacted before write.
		msg := coalesce(f.Description, f.RuleID, "Potential secret detected")
		out.AddFinding(report.Finding{
			RuleID:      f.RuleID,
			Title:       coalesce(f.RuleID, "secret"),
			Description: f.Description,
			Severity:    sev,
			Category:    "secret",
			File:        f.File,
			Line:        f.StartLine,
			Column:      f.StartColumn,
			Match:       match,
			Message:     msg,
			Source:      "gitleaks",
			Confidence:  report.ConfidenceDeterministic,
		})
	}
	return out, nil
}
