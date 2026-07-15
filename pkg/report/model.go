// Package report defines the unified finding and report model used by every
// safeanalyze check and output writer.
package report

import (
	"strings"
	"time"
)

// Severity levels.
const (
	SeverityCritical = "critical"
	SeverityHigh     = "high"
	SeverityMedium   = "medium"
	SeverityLow      = "low"
	SeverityInfo     = "info"
)

// Confidence levels.
const (
	ConfidenceDeterministic = "deterministic"
	ConfidenceHeuristic     = "heuristic"
	ConfidenceML            = "ml"
)

// Finding represents a single security or quality issue.
type Finding struct {
	RuleID      string `json:"rule_id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Severity    string `json:"severity"`
	Category    string `json:"category"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Column      int    `json:"column"`
	Match       string `json:"match,omitempty"`
	Message     string `json:"message"`
	Source      string `json:"source"`     // which check produced it
	Confidence  string `json:"confidence"` // deterministic, heuristic, ml
	Remediation string `json:"remediation,omitempty"`
}

// ErrorRecord captures a non-fatal error from a stage.
type ErrorRecord struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}

// Summary aggregates report statistics.
type Summary struct {
	FilesScanned     int            `json:"files_scanned"`
	FilesSanitized   int            `json:"files_sanitized"`
	BytesBefore      int64          `json:"bytes_before"`
	BytesAfter       int64          `json:"bytes_after"`
	CountsBySeverity map[string]int `json:"counts_by_severity"`
	CountsBySource   map[string]int `json:"counts_by_source"`
}

// Report is the unified output of the safeanalyze pipeline.
type Report struct {
	Target     string            `json:"target"`
	StartedAt  time.Time         `json:"started_at"`
	DurationMs int64             `json:"duration_ms"`
	Findings   []Finding         `json:"findings"`
	Errors     []ErrorRecord     `json:"errors"`
	Summary    Summary           `json:"summary"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// NewReport creates an empty report for a target.
func NewReport(target string) *Report {
	return &Report{
		Target:    target,
		StartedAt: time.Now().UTC(),
		Findings:  []Finding{},
		Errors:    []ErrorRecord{},
		Metadata:  map[string]string{},
		Summary: Summary{
			CountsBySeverity: map[string]int{},
			CountsBySource:   map[string]int{},
		},
	}
}

// SeverityRank returns a numeric ordering for severity levels. Higher values
// represent more severe findings.
func SeverityRank(sev string) int {
	switch strings.ToLower(sev) {
	case SeverityCritical:
		return 4
	case SeverityHigh:
		return 3
	case SeverityMedium:
		return 2
	case SeverityLow:
		return 1
	default:
		return 0
	}
}

// AddFinding appends a finding and updates summary counts.
func (r *Report) AddFinding(f Finding) {
	r.Findings = append(r.Findings, f)
	r.Summary.CountsBySeverity[f.Severity]++
	r.Summary.CountsBySource[f.Source]++
}

// AddError appends a non-fatal error record.
func (r *Report) AddError(source, message string) {
	r.Errors = append(r.Errors, ErrorRecord{Source: source, Message: message})
}

// HasSeverity returns true if the report contains any finding of the given severity.
func (r *Report) HasSeverity(sev string) bool {
	return r.Summary.CountsBySeverity[sev] > 0
}

// Finish records the total scan duration in milliseconds.
func (r *Report) Finish() {
	r.DurationMs = time.Since(r.StartedAt).Milliseconds()
}
