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
	args := []string{target, "--format", "json", "--min-severity", "HIGH"}
	cmd := exec.CommandContext(ctx, binary, args...)
	output, err := cmd.Output()
	if err != nil {
		out.AddError(s.Name(), fmt.Sprintf("prompt-injection-scanner execution failed: %v", err))
		return out, nil
	}
	return parsePromptInjectionScannerOutput(out, output)
}

func parsePromptInjectionScannerOutput(out *report.Report, data []byte) (*report.Report, error) {
	var results []struct {
		Rule     string `json:"rule"`
		Title    string `json:"title"`
		Message  string `json:"message"`
		File     string `json:"file"`
		Path     string `json:"path"`
		Line     int    `json:"line"`
		Column   int    `json:"column"`
		Severity string `json:"severity"`
		Match    string `json:"match"`
	}
	if err := json.Unmarshal(data, &results); err != nil {
		// Try single object wrapper with a "findings" field.
		var wrapper struct {
			Findings []struct {
				Rule     string `json:"rule"`
				Title    string `json:"title"`
				Message  string `json:"message"`
				File     string `json:"file"`
				Path     string `json:"path"`
				Line     int    `json:"line"`
				Column   int    `json:"column"`
				Severity string `json:"severity"`
				Match    string `json:"match"`
			} `json:"findings"`
		}
		if err2 := json.Unmarshal(data, &wrapper); err2 == nil && len(wrapper.Findings) > 0 {
			for _, r := range wrapper.Findings {
				results = append(results, struct {
					Rule     string `json:"rule"`
					Title    string `json:"title"`
					Message  string `json:"message"`
					File     string `json:"file"`
					Path     string `json:"path"`
					Line     int    `json:"line"`
					Column   int    `json:"column"`
					Severity string `json:"severity"`
					Match    string `json:"match"`
				}(r))
			}
		} else {
			out.AddError("prompt-injection-scanner", fmt.Sprintf("parsing JSON: %v", err))
			return out, nil
		}
	}
	for _, r := range results {
		file := r.File
		if file == "" {
			file = r.Path
		}
		msg := r.Message
		if msg == "" {
			msg = r.Title
		}
		out.AddFinding(report.Finding{
			RuleID:      coalesce(r.Rule, "prompt-injection"),
			Title:       coalesce(r.Title, r.Rule, "Prompt injection"),
			Description: msg,
			Severity:    mapSeverity(r.Severity),
			Category:    "prompt_injection",
			File:        file,
			Line:        r.Line,
			Column:      r.Column,
			Match:       r.Match,
			Message:     msg,
			Source:      "prompt-injection-scanner",
			Confidence:  report.ConfidenceDeterministic,
		})
	}
	return out, nil
}
