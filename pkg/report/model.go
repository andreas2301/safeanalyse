// Package report defines the unified finding and report model used by every
// safeanalyze check and output writer.
package report

import (
	"regexp"
	"strings"
	"time"
)

// secretTokenPatterns redact common secret-like values from report output so
// that generated reports can be committed to public branches without leaking
// credentials found in test data.
var secretTokenPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)(sk_|rk_|pk_)(live|test)_[A-Za-z0-9_\-]{10,}`),
	regexp.MustCompile(`(?i)sk-ant-[A-Za-z0-9_\-]{16,}`),
	regexp.MustCompile(`(?i)sk-(proj-|svcacct-)?[A-Za-z0-9_\-]{20,}`),
	regexp.MustCompile(`(?i)glpat-[A-Za-z0-9_\-]{10,}`),
	regexp.MustCompile(`(?i)gh[pousr]_[A-Za-z0-9]{30,}`),
	regexp.MustCompile(`(?i)github_pat_[A-Za-z0-9_]{20,}`),
	regexp.MustCompile(`(?i)xox[baprs]-[A-Za-z0-9\-]{10,}`),
	regexp.MustCompile(`(?i)(AKIA|ASIA)[A-Z0-9]{16}`),
	regexp.MustCompile(`(?s)-----BEGIN [A-Z0-9 ]*PRIVATE KEY-----.*?-----END [A-Z0-9 ]*PRIVATE KEY-----`),
	regexp.MustCompile(`-----BEGIN [A-Z0-9 ]*PRIVATE KEY-----`),
}

// quotedString matches a double-quoted string. It is used for conservative
// redaction of credential-like findings.
var quotedString = regexp.MustCompile(`"[^"]{4,}"`)

// redactTokens replaces known secret token patterns in arbitrary text.
func redactTokens(s string) string {
	for _, re := range secretTokenPatterns {
		s = re.ReplaceAllString(s, "[REDACTED]")
	}
	return s
}

// redactMatch sanitizes a finding match string for public report output.
func redactMatch(match, ruleID, source, category string) string {
	match = redactTokens(match)

	// For credential findings, also redact any quoted value so fake keys in
	// test fixtures are not pushed verbatim.
	if ruleID == "credential_hardcode" {
		match = quotedString.ReplaceAllString(match, `"[REDACTED]"`)
	}
	if source == "entropy" && strings.Contains(strings.ToLower(match), "pat") {
		match = quotedString.ReplaceAllString(match, `"[REDACTED]"`)
	}
	// External secret scanners (trufflehog, gitleaks) may detect provider-specific
	// tokens that do not match generic prefix regexes; never emit raw secrets.
	if category == "secret" && match != "" && !strings.Contains(match, "[REDACTED]") {
		return "[REDACTED]"
	}

	return match
}

// Redact sanitizes finding fields and verdict evidence in place so reports can
// be stored or shared without leaking credentials found in scanned targets.
func (r *Report) Redact() {
	for i := range r.Findings {
		f := &r.Findings[i]
		f.Match = redactMatch(f.Match, f.RuleID, f.Source, f.Category)
		f.Message = redactTokens(f.Message)
		f.Description = redactTokens(f.Description)
		for k, v := range f.Attributes {
			f.Attributes[k] = redactTokens(v)
		}
	}
	for i := range r.Verdicts {
		for j, ev := range r.Verdicts[i].Evidence {
			r.Verdicts[i].Evidence[j] = redactTokens(ev)
		}
	}
}
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
	// Tier ranks a finding for triage: A is verdict-relevant, B is capability
	// inventory, C is hygiene/noise. Nothing is dropped; tiers only order output.
	Tier string `json:"tier,omitempty"`
	// FileRole is the role of File (source, test, docs, lockfile, ci, ...).
	FileRole string `json:"file_role,omitempty"`
	// Question is the verdict question this finding answers (see Verdict).
	Question string `json:"question,omitempty"`
	// Attributes carries finding-specific evidence such as host, trigger or
	// dependency.
	Attributes map[string]string `json:"attributes,omitempty"`
}

// Triage tiers.
const (
	TierA = "A"
	TierB = "B"
	TierC = "C"
)

// Verdict answers one fixed question about the target.
type Verdict struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Answer   string   `json:"answer"` // yes, unclear, no, not_checked
	TierA    int      `json:"tier_a"`
	TierB    int      `json:"tier_b"`
	Evidence []string `json:"evidence,omitempty"`
}

// Verdict answers.
const (
	AnswerYes        = "yes"
	AnswerUnclear    = "unclear"
	AnswerNo         = "no"
	AnswerNotChecked = "not_checked"
)

// CoverageItem states whether a check ran, so a clean result is not mistaken
// for full coverage.
type CoverageItem struct {
	Check  string `json:"check"`
	Status string `json:"status"` // ran, error, not_run, not_available
	Note   string `json:"note,omitempty"`
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
	Verdicts   []Verdict         `json:"verdicts,omitempty"`
	Coverage   []CoverageItem    `json:"coverage,omitempty"`
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

// CountTier returns the number of findings in the given tier.
func (r *Report) CountTier(tier string) int {
	n := 0
	for _, f := range r.Findings {
		if f.Tier == tier {
			n++
		}
	}
	return n
}

// HasSeverity returns true if the report contains any finding of the given severity.
func (r *Report) HasSeverity(sev string) bool {
	return r.Summary.CountsBySeverity[sev] > 0
}

// Finish records the total scan duration in milliseconds.
func (r *Report) Finish() {
	r.DurationMs = time.Since(r.StartedAt).Milliseconds()
}
