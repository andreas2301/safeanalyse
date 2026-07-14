package external

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/user/safeanalyze/pkg/report"
)

// semgrepStage runs Semgrep with security-audit and prompt-injection rules.
type semgrepStage struct{}

// NewSemgrepStage creates a Semgrep pipeline stage.
func NewSemgrepStage() *semgrepStage { return &semgrepStage{} }

// Name returns the stage name.
func (s *semgrepStage) Name() string { return "semgrep" }

// Dependencies returns no dependencies.
func (s *semgrepStage) Dependencies() []string { return nil }

// Run executes semgrep and parses JSON output into findings.
func (s *semgrepStage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := prepareReport(input, target)
	binary := resolveBinary("semgrep")
	if binary == "" {
		return missingBinary(out, s.Name()), nil
	}
	args := []string{
		"--config=p/security-audit",
		"--config=https://raw.githubusercontent.com/foundationmachines/semgrep-ai-security/main/rules/prompt-injection.yml",
		target,
		"--json",
	}
	cmd := exec.CommandContext(ctx, binary, args...)
	output, err := cmd.Output()
	if err != nil {
		out.AddError(s.Name(), fmt.Sprintf("semgrep execution failed: %v", err))
		return out, nil
	}
	return parseSemgrepOutput(out, output)
}

func parseSemgrepOutput(out *report.Report, data []byte) (*report.Report, error) {
	var result struct {
		Results []struct {
			CheckID string `json:"check_id"`
			Path    string `json:"path"`
			Start   struct {
				Line int `json:"line"`
				Col  int `json:"col"`
			} `json:"start"`
			Extra struct {
				Message  string `json:"message"`
				Severity string `json:"severity"`
			} `json:"extra"`
		} `json:"results"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		out.AddError("semgrep", fmt.Sprintf("parsing JSON: %v", err))
		return out, nil
	}
	for _, r := range result.Results {
		out.AddFinding(report.Finding{
			RuleID:      r.CheckID,
			Title:       r.CheckID,
			Description: r.Extra.Message,
			Severity:    mapSemgrepSeverity(r.Extra.Severity),
			Category:    "external",
			File:        r.Path,
			Line:        r.Start.Line,
			Column:      r.Start.Col,
			Message:     r.Extra.Message,
			Source:      "semgrep",
			Confidence:  report.ConfidenceDeterministic,
		})
	}
	return out, nil
}

func mapSemgrepSeverity(s string) string {
	switch s {
	case "ERROR":
		return report.SeverityHigh
	case "WARNING":
		return report.SeverityMedium
	case "INFO":
		return report.SeverityInfo
	default:
		return report.SeverityMedium
	}
}
