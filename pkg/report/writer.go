package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/user/safeanalyze/pkg/config"
)

// Writer persists a report to one or more output formats.
type Writer interface {
	Write(report *Report, cfg config.OutputConfig, dir string) error
}

// WriteAll writes all configured output formats to cfg.OutDir.
// Supported formats: sarif, markdown, html, json.
func WriteAll(report *Report, cfg config.OutputConfig) error {
	if cfg.OutDir == "" {
		cfg.OutDir = "./safeanalyze-out"
	}

	if err := os.MkdirAll(cfg.OutDir, 0755); err != nil {
		return fmt.Errorf("creating output directory %s: %w", cfg.OutDir, err)
	}

	reportToWrite := report
	if max := cfg.EffectiveMaxFindings(); max > 0 && len(report.Findings) > max {
		capped := *report
		capped.Findings = make([]Finding, len(report.Findings))
		copy(capped.Findings, report.Findings)
		SortFindingsBySeverity(capped.Findings)
		capped.Findings = capped.Findings[:max]
		SortFindings(capped.Findings)
		capped.Metadata["findings_capped"] = fmt.Sprintf("showing %d of %d", max, len(report.Findings))
		reportToWrite = &capped
	}

	report.Redact()

	formats := cfg.EffectiveFormats()
	var errs []string
	for _, f := range formats {
		if err := writeFormat(reportToWrite, cfg, f); err != nil {
			errs = append(errs, fmt.Sprintf("%s: %v", f, err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("output errors: %s", strings.Join(errs, "; "))
	}
	return nil
}

func writeFormat(report *Report, cfg config.OutputConfig, format string) error {
	switch strings.ToLower(strings.TrimSpace(format)) {
	case "sarif":
		return WriteSARIF(report, filepath.Join(cfg.OutDir, "safeanalyze.sarif"))
	case "markdown", "md":
		return WriteMarkdown(report, cfg, filepath.Join(cfg.OutDir, "safeanalyze.md"))
	case "html":
		return WriteHTML(report, cfg, filepath.Join(cfg.OutDir, "safeanalyze.html"))
	case "json":
		return writeJSON(report, filepath.Join(cfg.OutDir, "safeanalyze.json"))
	default:
		return fmt.Errorf("unsupported output format %q", format)
	}
}

func writeJSON(report *Report, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("encoding JSON report: %w", err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("writing JSON file: %w", err)
	}
	return nil
}
