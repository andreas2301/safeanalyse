package report

import (
	"cmp"
	"slices"
)

// Merge combines multiple reports into a single deterministic report.
// Findings are sorted by (file, line, column, rule_id, source) to ensure
// stable output regardless of concurrent execution order.
func Merge(target string, reports []*Report) *Report {
	merged := NewReport(target)
	if len(reports) == 0 {
		return merged
	}

	// Use the earliest start time across all stage reports.
	for _, r := range reports {
		if r == nil {
			continue
		}
		if r.StartedAt.Before(merged.StartedAt) {
			merged.StartedAt = r.StartedAt
		}
		for k, v := range r.Metadata {
			if _, exists := merged.Metadata[k]; !exists {
				merged.Metadata[k] = v
			}
		}
		merged.Summary.FilesScanned = max(merged.Summary.FilesScanned, r.Summary.FilesScanned)
		merged.Summary.FilesSanitized = max(merged.Summary.FilesSanitized, r.Summary.FilesSanitized)
		merged.Summary.BytesBefore = max(merged.Summary.BytesBefore, r.Summary.BytesBefore)
		merged.Summary.BytesAfter = max(merged.Summary.BytesAfter, r.Summary.BytesAfter)
		for _, f := range r.Findings {
			merged.AddFinding(f)
		}
		for _, e := range r.Errors {
			merged.AddError(e.Source, e.Message)
		}
	}

	SortFindings(merged.Findings)
	return merged
}

// SortFindingsBySeverity sorts findings by severity (highest first) and then
// by the standard deterministic tuple for stable output.
func SortFindingsBySeverity(findings []Finding) {
	slices.SortStableFunc(findings, func(a, b Finding) int {
		if c := cmp.Compare(SeverityRank(b.Severity), SeverityRank(a.Severity)); c != 0 {
			return c
		}
		if c := cmp.Compare(a.File, b.File); c != 0 {
			return c
		}
		if c := cmp.Compare(a.Line, b.Line); c != 0 {
			return c
		}
		if c := cmp.Compare(a.Column, b.Column); c != 0 {
			return c
		}
		if c := cmp.Compare(a.RuleID, b.RuleID); c != 0 {
			return c
		}
		return cmp.Compare(a.Source, b.Source)
	})
}

// SortFindings sorts findings deterministically in place.
func SortFindings(findings []Finding) {
	slices.SortStableFunc(findings, func(a, b Finding) int {
		if c := cmp.Compare(a.File, b.File); c != 0 {
			return c
		}
		if c := cmp.Compare(a.Line, b.Line); c != 0 {
			return c
		}
		if c := cmp.Compare(a.Column, b.Column); c != 0 {
			return c
		}
		if c := cmp.Compare(a.RuleID, b.RuleID); c != 0 {
			return c
		}
		return cmp.Compare(a.Source, b.Source)
	})
}
