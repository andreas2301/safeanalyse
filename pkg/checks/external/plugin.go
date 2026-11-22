package external

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/user/safeanalyze/pkg/config"
	"github.com/user/safeanalyze/pkg/report"
)

// Supported declarative plugin output formats.
const (
	FormatSARIF           = "sarif"
	FormatSafeanalyzeJSON = "safeanalyze-json"
	FormatNDJSON          = "ndjson"
)

// genericPluginStage executes any user-configured external scanner declared in
// safeanalyze.yaml (e.g. osv-scanner, trivy, guarddog, safeanalyze-llm) without
// shell execution, normalizing SARIF 2.1.0, safeanalyze-json, or NDJSON output
// into unified report.Finding records.
type genericPluginStage struct {
	name     string
	command  string
	format   string
	question string
	timeout  time.Duration
}

// NewGenericPluginStage constructs a declarative external plugin stage.
func NewGenericPluginStage(sc config.ScannerConfig) *genericPluginStage {
	return &genericPluginStage{
		name:     sc.Name,
		command:  sc.Command,
		format:   strings.ToLower(strings.TrimSpace(sc.Format)),
		question: strings.TrimSpace(sc.Question),
		timeout:  effectiveTimeout(sc.TimeoutSeconds),
	}
}

// Name implements pipeline.Stage.
func (s *genericPluginStage) Name() string { return s.name }

// Dependencies implements pipeline.Stage.
func (s *genericPluginStage) Dependencies() []string { return nil }

// Run executes the configured command safely (direct argv, no shell) and parses
// its stdout according to s.format.
func (s *genericPluginStage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := prepareReport(input, target)
	tokens, err := splitArgs(s.command)
	if err != nil {
		out.AddError(s.name, fmt.Sprintf("invalid plugin command: %v", err))
		return out, nil
	}
	if len(tokens) == 0 {
		out.AddError(s.name, "empty plugin command")
		return out, nil
	}

	binary := resolveBinary(tokens[0])
	if binary == "" {
		return missingBinary(out, s.name), nil
	}
	args := make([]string, 0, len(tokens)-1)
	for _, tok := range tokens[1:] {
		args = append(args, strings.ReplaceAll(tok, "{path}", target))
	}

	output, stderr, exitCode, runErr := runCommand(ctx, s.timeout, binary, args)
	if runErr != nil {
		// Many security scanners (osv-scanner, trivy, gitleaks, semgrep) exit 1
		// when findings exist; attempt to parse stdout if non-empty.
		if exitCode == 1 && len(bytes.TrimSpace(output)) > 0 {
			return s.parseOutput(out, output)
		}
		msg := fmt.Sprintf("plugin %s failed: %v", s.name, runErr)
		if stderr != "" {
			msg += ": " + stderr
		}
		out.AddError(s.name, msg)
		return out, nil
	}
	if len(bytes.TrimSpace(output)) == 0 {
		return out, nil
	}
	return s.parseOutput(out, output)
}

func (s *genericPluginStage) parseOutput(out *report.Report, data []byte) (*report.Report, error) {
	switch s.format {
	case FormatSARIF:
		return parseSARIFPluginOutput(out, s.name, s.question, data)
	case FormatSafeanalyzeJSON:
		return parseSafeanalyzeJSONPluginOutput(out, s.name, s.question, data)
	case FormatNDJSON:
		return parseNDJSONPluginOutput(out, s.name, s.question, data)
	default:
		out.AddError(s.name, fmt.Sprintf("unsupported plugin format %q (expected sarif, safeanalyze-json, or ndjson)", s.format))
		return out, nil
	}
}

func parseSARIFPluginOutput(out *report.Report, source, question string, data []byte) (*report.Report, error) {
	var doc struct {
		Runs []struct {
			Results []struct {
				RuleID  string `json:"ruleId"`
				Level   string `json:"level"`
				Message struct {
					Text string `json:"text"`
				} `json:"message"`
				Locations []struct {
					PhysicalLocation struct {
						ArtifactLocation struct {
							URI string `json:"uri"`
						} `json:"artifactLocation"`
						Region struct {
							StartLine   int `json:"startLine"`
							StartColumn int `json:"startColumn"`
							Snippet     struct {
								Text string `json:"text"`
							} `json:"snippet"`
						} `json:"region"`
					} `json:"physicalLocation"`
				} `json:"locations"`
			} `json:"results"`
		} `json:"runs"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		out.AddError(source, fmt.Sprintf("parsing SARIF: %v", err))
		return out, nil
	}
	for _, run := range doc.Runs {
		for _, r := range run.Results {
			var file, match string
			var line, col int
			if len(r.Locations) > 0 {
				pl := r.Locations[0].PhysicalLocation
				file = pl.ArtifactLocation.URI
				line = pl.Region.StartLine
				col = pl.Region.StartColumn
				match = pl.Region.Snippet.Text
			}
			out.AddFinding(report.Finding{
				RuleID:      coalesce(r.RuleID, source+"-finding"),
				Title:       coalesce(r.RuleID, source+" finding"),
				Description: r.Message.Text,
				Severity:    mapSeverity(r.Level),
				Category:    "external",
				File:        file,
				Line:        line,
				Column:      col,
				Match:       match,
				Message:     coalesce(r.Message.Text, r.RuleID),
				Source:      source,
				Confidence:  report.ConfidenceDeterministic,
				Question:    question,
			})
		}
	}
	return out, nil
}

func parseSafeanalyzeJSONPluginOutput(out *report.Report, source, question string, data []byte) (*report.Report, error) {
	var sub report.Report
	if err := json.Unmarshal(data, &sub); err != nil {
		out.AddError(source, fmt.Sprintf("parsing safeanalyze-json report: %v", err))
		return out, nil
	}
	for _, f := range sub.Findings {
		if f.Source == "" {
			f.Source = source
		}
		if f.Question == "" && question != "" {
			f.Question = question
		}
		out.AddFinding(f)
	}
	for _, e := range sub.Errors {
		src := e.Source
		if src == "" {
			src = source
		}
		out.AddError(src, e.Message)
	}
	for k, v := range sub.Metadata {
		if k == "safeanalyze_version" {
			continue
		}
		if _, exists := out.Metadata[k]; !exists {
			out.Metadata[k] = v
		}
	}
	return out, nil
}

func parseNDJSONPluginOutput(out *report.Report, source, question string, data []byte) (*report.Report, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Buffer(make([]byte, 1024), 4*1024*1024)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			continue
		}
		var rec struct {
			RuleID   string `json:"rule_id"`
			Rule     string `json:"rule"`
			Title    string `json:"title"`
			Message  string `json:"message"`
			Severity string `json:"severity"`
			Category string `json:"category"`
			File     string `json:"file"`
			Path     string `json:"path"`
			Line     int    `json:"line"`
			Column   int    `json:"column"`
			Match    string `json:"match"`
		}
		if err := json.Unmarshal(line, &rec); err != nil {
			out.AddError(source, fmt.Sprintf("parsing NDJSON line: %v", err))
			continue
		}
		ruleID := coalesce(rec.RuleID, rec.Rule, source+"-finding")
		out.AddFinding(report.Finding{
			RuleID:      ruleID,
			Title:       coalesce(rec.Title, ruleID),
			Description: rec.Message,
			Severity:    mapSeverity(rec.Severity),
			Category:    coalesce(rec.Category, "external"),
			File:        coalesce(rec.File, rec.Path),
			Line:        rec.Line,
			Column:      rec.Column,
			Match:       rec.Match,
			Message:     coalesce(rec.Message, rec.Title, ruleID),
			Source:      source,
			Confidence:  report.ConfidenceHeuristic,
			Question:    question,
		})
	}
	if err := scanner.Err(); err != nil {
		out.AddError(source, fmt.Sprintf("reading NDJSON: %v", err))
	}
	return out, nil
}
