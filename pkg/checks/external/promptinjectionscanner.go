package external

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/user/safeanalyze/pkg/report"
)

// promptInjectionScannerStage runs the prompt-injection-scanner.
type promptInjectionScannerStage struct{}

// NewPromptInjectionScannerStage creates a prompt-injection-scanner pipeline stage.
func NewPromptInjectionScannerStage() *promptInjectionScannerStage {
	return &promptInjectionScannerStage{}
}

// Name returns the stage name.
func (s *promptInjectionScannerStage) Name() string { return "prompt-injection-scanner" }

// Dependencies returns no dependencies.
func (s *promptInjectionScannerStage) Dependencies() []string { return nil }

// Run executes prompt-injection-scanner and parses JSON findings.
func (s *promptInjectionScannerStage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := prepareReport(input, target)
	binary := resolveBinary("prompt-injection-scanner")
	if binary == "" {
		return missingBinary(out, s.Name()), nil
	}
	args := []string{target, "--format", "json", "--min-severity", "HIGH", "--no-fail"}
	cmd := exec.CommandContext(ctx, binary, args...)
	output, err := cmd.Output()
	if err != nil {
		out.AddError(s.Name(), fmt.Sprintf("prompt-injection-scanner execution failed: %v", err))
		return out, nil
	}
	return parsePromptInjectionScannerOutput(out, output)
}

func parsePromptInjectionScannerOutput(out *report.Report, data []byte) (*report.Report, error) {
	var result struct {
		Findings []struct {
			RuleID            string `json:"rule_id"`
			RuleName          string `json:"rule_name"`
			Severity          string `json:"severity"`
			FilePath          string `json:"file_path"`
			LineNumber        int    `json:"line_number"`
			MatchedExpression string `json:"matched_expression"`
			Description       string `json:"description"`
		} `json:"findings"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		out.AddError("prompt-injection-scanner", fmt.Sprintf("parsing JSON: %v", err))
		return out, nil
	}
	for _, r := range result.Findings {
		msg := r.Description
		if msg == "" {
			msg = r.RuleName
		}
		out.AddFinding(report.Finding{
			RuleID:      coalesce(r.RuleID, "prompt-injection"),
			Title:       coalesce(r.RuleName, r.RuleID, "Prompt injection"),
			Description: msg,
			Severity:    mapSeverity(r.Severity),
			Category:    "prompt_injection",
			File:        r.FilePath,
			Line:        r.LineNumber,
			Column:      0,
			Match:       r.MatchedExpression,
			Message:     msg,
			Source:      "prompt-injection-scanner",
			Confidence:  report.ConfidenceDeterministic,
		})
	}
	return out, nil
}
