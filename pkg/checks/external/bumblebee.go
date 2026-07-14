package external

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/user/safeanalyze/pkg/report"
)

// bumblebeeStage runs the Bumblebee scanner in deep profile mode.
type bumblebeeStage struct{}

// NewBumblebeeStage creates a Bumblebee pipeline stage.
func NewBumblebeeStage() *bumblebeeStage { return &bumblebeeStage{} }

// Name returns the stage name.
func (s *bumblebeeStage) Name() string { return "bumblebee" }

// Dependencies returns no dependencies.
func (s *bumblebeeStage) Dependencies() []string { return nil }

// Run executes bumblebee and parses NDJSON findings.
func (s *bumblebeeStage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := prepareReport(input, target)
	binary := resolveBinary("bumblebee")
	if binary == "" {
		return missingBinary(out, s.Name()), nil
	}
	args := []string{"scan", "--profile", "deep", "--root", target, "--findings-only"}
	cmd := exec.CommandContext(ctx, binary, args...)
	output, err := cmd.Output()
	if err != nil {
		out.AddError(s.Name(), fmt.Sprintf("bumblebee execution failed: %v", err))
		return out, nil
	}
	return parseBumblebeeOutput(out, output)
}

func parseBumblebeeOutput(out *report.Report, data []byte) (*report.Report, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Buffer(make([]byte, 1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		var f struct {
			Rule     string `json:"rule"`
			Title    string `json:"title"`
			Message  string `json:"message"`
			File     string `json:"file"`
			Path     string `json:"path"`
			Line     int    `json:"line"`
			Column   int    `json:"column"`
			Severity string `json:"severity"`
		}
		if err := json.Unmarshal(line, &f); err != nil {
			out.AddError("bumblebee", fmt.Sprintf("parsing NDJSON line: %v", err))
			continue
		}
		file := f.File
		if file == "" {
			file = f.Path
		}
		msg := f.Message
		if msg == "" {
			msg = f.Title
		}
		out.AddFinding(report.Finding{
			RuleID:      coalesce(f.Rule, "bumblebee-finding"),
			Title:       coalesce(f.Title, f.Rule, "Bumblebee finding"),
			Description: msg,
			Severity:    mapSeverity(f.Severity),
			Category:    "external",
			File:        file,
			Line:        f.Line,
			Column:      f.Column,
			Message:     msg,
			Source:      "bumblebee",
			Confidence:  report.ConfidenceHeuristic,
		})
	}
	if err := scanner.Err(); err != nil {
		out.AddError("bumblebee", fmt.Sprintf("reading NDJSON: %v", err))
	}
	return out, nil
}
