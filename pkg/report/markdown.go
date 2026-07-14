package report

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/user/safeanalyze/pkg/config"
)

// WriteMarkdown writes a human-readable Markdown report to path.
func WriteMarkdown(report *Report, cfg config.OutputConfig, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	var b strings.Builder
	writeMDHeader(&b, report)
	writeMDSummary(&b, report)
	writeMDFindings(&b, report)
	writeMDErrors(&b, report)
	if cfg.IncludeFileTree {
		writeMDFileTree(&b, report)
	}

	if err := os.WriteFile(path, []byte(b.String()), 0644); err != nil {
		return fmt.Errorf("writing markdown file: %w", err)
	}
	return nil
}

func writeMDHeader(b *strings.Builder, report *Report) {
	fmt.Fprintf(b, "# safeanalyze Report: %s\n\n", report.Target)
	fmt.Fprintf(b, "- **Started:** %s\n", report.StartedAt.Format("2006-01-02 15:04:05 MST"))
	fmt.Fprintf(b, "- **Completed:** %s\n", report.CompletedAt.Format("2006-01-02 15:04:05 MST"))
	fmt.Fprintf(b, "- **Total findings:** %d\n\n", len(report.Findings))
}

func writeMDSummary(b *strings.Builder, report *Report) {
	fmt.Fprintln(b, "## Summary")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "| Metric | Value |\n")
	fmt.Fprintf(b, "| --- | --- |\n")
	fmt.Fprintf(b, "| Files scanned | %d |\n", report.Summary.FilesScanned)
	fmt.Fprintf(b, "| Files sanitized | %d |\n", report.Summary.FilesSanitized)
	fmt.Fprintf(b, "| Bytes before | %d |\n", report.Summary.BytesBefore)
	fmt.Fprintf(b, "| Bytes after | %d |\n", report.Summary.BytesAfter)
	fmt.Fprintf(b, "| Total findings | %d |\n", len(report.Findings))
	fmt.Fprintf(b, "| Errors | %d |\n", len(report.Errors))
	fmt.Fprintln(b)

	if len(report.Summary.CountsBySeverity) > 0 {
		fmt.Fprintln(b, "### Findings by severity")
		fmt.Fprintln(b)
		fmt.Fprintf(b, "| Severity | Count |\n")
		fmt.Fprintf(b, "| --- | --- |\n")
		for _, sev := range severityOrder() {
			if count, ok := report.Summary.CountsBySeverity[sev]; ok {
				fmt.Fprintf(b, "| %s | %d |\n", strings.Title(sev), count)
			}
		}
		fmt.Fprintln(b)
	}

	if len(report.Summary.CountsBySource) > 0 {
		fmt.Fprintln(b, "### Findings by source")
		fmt.Fprintln(b)
		fmt.Fprintf(b, "| Source | Count |\n")
		fmt.Fprintf(b, "| --- | --- |\n")
		sources := sortedKeys(report.Summary.CountsBySource)
		for _, src := range sources {
			fmt.Fprintf(b, "| %s | %d |\n", src, report.Summary.CountsBySource[src])
		}
		fmt.Fprintln(b)
	}
}

func writeMDFindings(b *strings.Builder, report *Report) {
	fmt.Fprintln(b, "## Findings")
	fmt.Fprintln(b)

	if len(report.Findings) == 0 {
		fmt.Fprintln(b, "No findings detected.")
		fmt.Fprintln(b)
		return
	}

	bySeverity := groupBySeverity(report.Findings)
	for _, sev := range severityOrder() {
		group := bySeverity[sev]
		if len(group) == 0 {
			continue
		}
		fmt.Fprintf(b, "### %s (%d)\n\n", strings.Title(sev), len(group))
		fmt.Fprintf(b, "| Rule | File | Line | Column | Message | Source | Confidence |\n")
		fmt.Fprintf(b, "| --- | --- | --- | --- | --- | --- | --- |\n")
		for _, f := range group {
			fmt.Fprintf(b, "| %s | %s | %d | %d | %s | %s | %s |\n",
				escapeMD(f.RuleID), escapeMD(f.File), f.Line, f.Column,
				escapeMD(f.Message), escapeMD(f.Source), escapeMD(f.Confidence))
		}
		fmt.Fprintln(b)
	}
}

func writeMDErrors(b *strings.Builder, report *Report) {
	fmt.Fprintln(b, "## Errors")
	fmt.Fprintln(b)

	if len(report.Errors) == 0 {
		fmt.Fprintln(b, "No errors recorded.")
		fmt.Fprintln(b)
		return
	}

	for _, e := range report.Errors {
		fmt.Fprintf(b, "- **[%s]** %s\n", escapeMD(e.Source), escapeMD(e.Message))
	}
	fmt.Fprintln(b)
}

func writeMDFileTree(b *strings.Builder, report *Report) {
	files := make(map[string]struct{})
	for _, f := range report.Findings {
		files[f.File] = struct{}{}
	}
	if len(files) == 0 {
		return
	}

	fmt.Fprintln(b, "## Affected files")
	fmt.Fprintln(b)
	paths := sortedKeysFromSet(files)
	for _, p := range paths {
		fmt.Fprintf(b, "- `%s`\n", p)
	}
	fmt.Fprintln(b)
}

func groupBySeverity(findings []Finding) map[string][]Finding {
	m := make(map[string][]Finding)
	for _, f := range findings {
		m[f.Severity] = append(m[f.Severity], f)
	}
	return m
}

func severityOrder() []string {
	return []string{SeverityCritical, SeverityHigh, SeverityMedium, SeverityLow, SeverityInfo}
}

func sortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func sortedKeysFromSet(s map[string]struct{}) []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func escapeMD(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}
