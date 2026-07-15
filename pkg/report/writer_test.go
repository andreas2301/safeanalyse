package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/user/safeanalyze/pkg/config"
)

func sampleReport() *Report {
	r := NewReport("test-target")
	r.StartedAt = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	r.DurationMs = 60000
	r.Summary.FilesScanned = 10
	r.Summary.FilesSanitized = 2
	r.Summary.BytesBefore = 1024
	r.Summary.BytesAfter = 512
	r.AddFinding(Finding{
		RuleID:      "TEST-001",
		Title:       "Test Finding",
		Description: "A description",
		Severity:    SeverityHigh,
		Category:    "test",
		File:        "src/main.go",
		Line:        42,
		Column:      7,
		Match:       "secret=12345",
		Message:     "Possible secret exposed",
		Source:      "mock",
		Confidence:  ConfidenceDeterministic,
		Remediation: "Remove the secret",
	})
	r.AddFinding(Finding{
		RuleID:   "TEST-002",
		Title:    "Low Priority",
		Severity: SeverityLow,
		Category: "style",
		File:     "src/util.go",
		Line:     10,
		Column:   1,
		Message:  "Minor issue",
		Source:   "mock",
	})
	r.AddError("scanner", "scanner not installed")
	return r
}

func TestWriteAll_JSON(t *testing.T) {
	dir := t.TempDir()
	cfg := config.OutputConfig{Formats: []string{"json"}, OutDir: dir}

	if err := WriteAll(sampleReport(), cfg); err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	path := filepath.Join(dir, "safeanalyze.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading JSON output: %v", err)
	}

	var parsed Report
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("JSON output is not valid: %v", err)
	}
	if parsed.Target != "test-target" {
		t.Errorf("expected target test-target, got %s", parsed.Target)
	}
	if len(parsed.Findings) != 2 {
		t.Errorf("expected 2 findings, got %d", len(parsed.Findings))
	}
}

func TestWriteAll_Markdown(t *testing.T) {
	dir := t.TempDir()
	cfg := config.OutputConfig{Formats: []string{"markdown"}, OutDir: dir, IncludeFileTree: true}

	if err := WriteAll(sampleReport(), cfg); err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "safeanalyze.md"))
	if err != nil {
		t.Fatalf("reading markdown output: %v", err)
	}
	md := string(data)

	for _, want := range []string{
		"# safeanalyze Report: test-target",
		"TEST-001",
		"Possible secret exposed",
		"Affected files",
		"src/main.go",
		"[scanner]",
	} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown output missing expected content: %q", want)
		}
	}
}

func TestWriteAll_SARIF(t *testing.T) {
	dir := t.TempDir()
	cfg := config.OutputConfig{Formats: []string{"sarif"}, OutDir: dir}

	if err := WriteAll(sampleReport(), cfg); err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "safeanalyze.sarif"))
	if err != nil {
		t.Fatalf("reading SARIF output: %v", err)
	}

	var sarif map[string]interface{}
	if err := json.Unmarshal(data, &sarif); err != nil {
		t.Fatalf("SARIF output is not valid JSON: %v", err)
	}

	if sarif["version"] != "2.1.0" {
		t.Errorf("expected SARIF version 2.1.0, got %v", sarif["version"])
	}

	runs, ok := sarif["runs"].([]interface{})
	if !ok || len(runs) == 0 {
		t.Fatalf("SARIF has no runs")
	}
	run := runs[0].(map[string]interface{})
	results, ok := run["results"].([]interface{})
	if !ok || len(results) != 2 {
		t.Errorf("expected 2 SARIF results, got %d", len(results))
	}
}

func TestWriteAll_MultipleFormats(t *testing.T) {
	dir := t.TempDir()
	cfg := config.OutputConfig{Formats: []string{"json", "markdown", "html", "sarif"}, OutDir: dir}

	if err := WriteAll(sampleReport(), cfg); err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	for _, name := range []string{"safeanalyze.json", "safeanalyze.md", "safeanalyze.html", "safeanalyze.sarif"} {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err != nil {
			t.Errorf("expected output file %s: %v", name, err)
		}
	}
}

func TestOutputConfig_EffectiveFormats(t *testing.T) {
	if got := (config.OutputConfig{Formats: []string{"json", "html"}}).EffectiveFormats(); len(got) != 2 || got[0] != "json" || got[1] != "html" {
		t.Errorf("expected [json html], got %v", got)
	}
	if got := (config.OutputConfig{Format: "sarif"}).EffectiveFormats(); len(got) != 1 || got[0] != "sarif" {
		t.Errorf("expected [sarif], got %v", got)
	}
	if got := (config.OutputConfig{}).EffectiveFormats(); len(got) != 1 || got[0] != "markdown" {
		t.Errorf("expected [markdown], got %v", got)
	}
}

func TestWriteAll_UnsupportedFormat(t *testing.T) {
	dir := t.TempDir()
	cfg := config.OutputConfig{Formats: []string{"plain"}, OutDir: dir}

	err := WriteAll(sampleReport(), cfg)
	if err == nil {
		t.Fatal("expected error for unsupported format")
	}
	if !strings.Contains(err.Error(), "unsupported output format") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestWriteAll_CapPreservesSeverity(t *testing.T) {
	dir := t.TempDir()
	cfg := config.OutputConfig{Formats: []string{"json"}, OutDir: dir, MaxFindings: 3}

	r := NewReport("cap-test")
	for i := 0; i < 10; i++ {
		r.AddFinding(Finding{RuleID: "low", Severity: SeverityLow, File: fmt.Sprintf("f%d", i), Line: i, Source: "test"})
	}
	r.AddFinding(Finding{RuleID: "critical", Severity: SeverityCritical, File: "critical.go", Line: 1, Source: "test"})

	if err := WriteAll(r, cfg); err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	data, _ := os.ReadFile(filepath.Join(dir, "safeanalyze.json"))
	var parsed Report
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("JSON unmarshal failed: %v", err)
	}
	if len(parsed.Findings) != 3 {
		t.Fatalf("expected 3 capped findings, got %d", len(parsed.Findings))
	}
	seenCritical := false
	for _, f := range parsed.Findings {
		if f.RuleID == "critical" {
			seenCritical = true
		}
		if f.Severity == SeverityLow && f.File == "f9" {
			t.Errorf("low-severity finding should have been dropped in favor of higher severity")
		}
	}
	if !seenCritical {
		t.Errorf("critical finding was dropped by cap")
	}
}
