package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// WriteSARIF writes the report as a SARIF v2.1.0 JSON file to path.
// It uses a minimal hand-rolled sarifLog that is valid against the SARIF 2.1.0
// schema and includes tool driver metadata with one rule per unique RuleID.
func WriteSARIF(report *Report, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	rules := buildSARIFRules(report.Findings)
	ruleIndices := make(map[string]int, len(rules))
	for i, r := range rules {
		ruleIndices[r.ID] = i
	}

	results := make([]sarifResult, 0, len(report.Findings))
	for _, f := range report.Findings {
		results = append(results, findingToResult(f, ruleIndices))
	}

	log := sarifLog{
		Version: "2.1.0",
		Schema:  "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
		Runs: []sarifRun{
			{
				Tool: sarifTool{
					Driver: sarifDriver{
						Name:           "safeanalyze",
						InformationURI: "https://github.com/user/safeanalyze",
						Rules:          rules,
					},
				},
				Results: results,
				Invocations: []sarifInvocation{
					{
						ExecutionSuccessful: true,
						StartTimeUTC:        report.StartedAt.Format(timeRFC3339Millis),
						EndTimeUTC:          report.CompletedAt.Format(timeRFC3339Millis),
					},
				},
			},
		},
	}

	data, err := json.MarshalIndent(log, "", "  ")
	if err != nil {
		return fmt.Errorf("encoding SARIF: %w", err)
	}
	data = append(data, '\n')

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("writing SARIF file: %w", err)
	}
	return nil
}

const timeRFC3339Millis = "2006-01-02T15:04:05.000Z07:00"

type sarifLog struct {
	Version string     `json:"version"`
	Schema  string     `json:"$schema"`
	Runs    []sarifRun `json:"runs"`
}

type sarifRun struct {
	Tool        sarifTool         `json:"tool"`
	Results     []sarifResult     `json:"results"`
	Invocations []sarifInvocation `json:"invocations"`
}

type sarifTool struct {
	Driver sarifDriver `json:"driver"`
}

type sarifDriver struct {
	Name           string      `json:"name"`
	InformationURI string      `json:"informationUri"`
	Rules          []sarifRule `json:"rules"`
}

type sarifRule struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	ShortDescription sarifMessage      `json:"shortDescription"`
	FullDescription  sarifMessage      `json:"fullDescription,omitempty"`
	Help             sarifMessage      `json:"help,omitempty"`
	Properties       map[string]string `json:"properties,omitempty"`
}

type sarifMessage struct {
	Text string `json:"text"`
}

type sarifResult struct {
	RuleID     string           `json:"ruleId"`
	RuleIndex  int              `json:"ruleIndex"`
	Level      string           `json:"level"`
	Message    sarifMessage     `json:"message"`
	Locations  []sarifLocation  `json:"locations"`
	Properties sarifResultProps `json:"properties"`
}

type sarifLocation struct {
	PhysicalLocation sarifPhysicalLocation `json:"physicalLocation"`
}

type sarifPhysicalLocation struct {
	ArtifactLocation sarifArtifactLocation `json:"artifactLocation"`
	Region           sarifRegion           `json:"region"`
}

type sarifArtifactLocation struct {
	URI string `json:"uri"`
}

type sarifRegion struct {
	StartLine   int          `json:"startLine"`
	StartColumn int          `json:"startColumn"`
	Snippet     sarifMessage `json:"snippet,omitempty"`
}

type sarifResultProps struct {
	Category    string `json:"category,omitempty"`
	Confidence  string `json:"confidence,omitempty"`
	Source      string `json:"source,omitempty"`
	Remediation string `json:"remediation,omitempty"`
}

type sarifInvocation struct {
	ExecutionSuccessful bool   `json:"executionSuccessful"`
	StartTimeUTC        string `json:"startTimeUtc"`
	EndTimeUTC          string `json:"endTimeUtc"`
}

func buildSARIFRules(findings []Finding) []sarifRule {
	seen := make(map[string]struct{})
	rules := make([]sarifRule, 0)
	for _, f := range findings {
		if _, ok := seen[f.RuleID]; ok || f.RuleID == "" {
			continue
		}
		seen[f.RuleID] = struct{}{}
		rules = append(rules, sarifRule{
			ID:               f.RuleID,
			Name:             f.Title,
			ShortDescription: sarifMessage{Text: f.Title},
			FullDescription:  sarifMessage{Text: f.Description},
			Help:             sarifMessage{Text: f.Remediation},
			Properties: map[string]string{
				"category": f.Category,
			},
		})
	}
	sort.Slice(rules, func(i, j int) bool { return rules[i].ID < rules[j].ID })
	return rules
}

func findingToResult(f Finding, ruleIndices map[string]int) sarifResult {
	idx := -1
	if i, ok := ruleIndices[f.RuleID]; ok {
		idx = i
	}

	msg := f.Message
	if msg == "" {
		msg = f.Title
	}

	line := f.Line
	if line <= 0 {
		line = 1
	}
	column := f.Column
	if column <= 0 {
		column = 1
	}

	return sarifResult{
		RuleID:    f.RuleID,
		RuleIndex: idx,
		Level:     severityToSARIFLevel(f.Severity),
		Message:   sarifMessage{Text: msg},
		Locations: []sarifLocation{
			{
				PhysicalLocation: sarifPhysicalLocation{
					ArtifactLocation: sarifArtifactLocation{URI: f.File},
					Region: sarifRegion{
						StartLine:   line,
						StartColumn: column,
						Snippet:     sarifMessage{Text: f.Match},
					},
				},
			},
		},
		Properties: sarifResultProps{
			Category:    f.Category,
			Confidence:  f.Confidence,
			Source:      f.Source,
			Remediation: f.Remediation,
		},
	}
}

func severityToSARIFLevel(sev string) string {
	switch sev {
	case SeverityCritical, SeverityHigh:
		return "error"
	case SeverityMedium:
		return "warning"
	case SeverityLow, SeverityInfo:
		return "note"
	default:
		return "warning"
	}
}
