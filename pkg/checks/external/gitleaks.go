package external

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/user/safeanalyze/pkg/report"
)

// gitleaksStage runs Gitleaks against a directory.
type gitleaksStage struct{}

// NewGitleaksStage creates a Gitleaks pipeline stage.
func NewGitleaksStage() *gitleaksStage { return &gitleaksStage{} }

// Name returns the stage name.
func (s *gitleaksStage) Name() string { return "gitleaks" }

// Dependencies returns no dependencies.
func (s *gitleaksStage) Dependencies() []string { return nil }

// Run executes gitleaks and parses JSON findings.
func (s *gitleaksStage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := prepareReport(input, target)
	binary := resolveBinary("gitleaks")
	if binary == "" {
		return missingBinary(out, s.Name()), nil
	}
	args := []string{"dir", target, "--report-format", "json", "--report-path", "-"}
	cmd := exec.CommandContext(ctx, binary, args...)
	output, err := cmd.Output()
	if err != nil {
		out.AddError(s.Name(), fmt.Sprintf("gitleaks execution failed: %v", err))
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
			Message:     fmt.Sprintf("%s: %s", f.Description, match),
			Source:      "gitleaks",
			Confidence:  report.ConfidenceDeterministic,
		})
	}
	return out, nil
}
