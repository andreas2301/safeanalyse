package external

import (
	"testing"

	"github.com/user/safeanalyze/pkg/report"
)

func TestParseSemgrepOutput(t *testing.T) {
	data := []byte(`{"results": [{"check_id": "rule1", "path": "test.go", "start": {"line": 5, "col": 10}, "end": {"line": 5, "col": 20}, "extra": {"message": "injection", "severity": "ERROR"}}], "errors": []}`)
	out := report.NewReport("target")
	out, err := parseSemgrepOutput(out, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(out.Findings))
	}
	f := out.Findings[0]
	if f.RuleID != "rule1" || f.File != "test.go" || f.Line != 5 || f.Column != 10 {
		t.Fatalf("unexpected finding: %+v", f)
	}
	if f.Source != "semgrep" || f.Confidence != report.ConfidenceDeterministic {
		t.Fatalf("unexpected source/confidence: %s/%s", f.Source, f.Confidence)
	}
	if f.Severity != report.SeverityHigh {
		t.Fatalf("expected severity high, got %s", f.Severity)
	}
}

func TestParseBumblebeeOutput(t *testing.T) {
	data := []byte(`{"rule": "r1", "file": "a.go", "line": 1, "column": 2, "severity": "high", "message": "m1"}
{"title": "t2", "path": "b.go", "severity": "medium"}
`)
	out := report.NewReport("target")
	out, err := parseBumblebeeOutput(out, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Findings) != 2 {
		t.Fatalf("expected 2 findings, got %d", len(out.Findings))
	}
	if out.Findings[0].Source != "bumblebee" {
		t.Fatalf("unexpected source: %s", out.Findings[0].Source)
	}
}

func TestParsePromptInjectionScannerOutput(t *testing.T) {
	data := []byte(`[{"rule": "pi", "file": "test.go", "line": 3, "column": 4, "severity": "HIGH", "match": "match"}]`)
	out := report.NewReport("target")
	out, err := parsePromptInjectionScannerOutput(out, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(out.Findings))
	}
	f := out.Findings[0]
	if f.Source != "prompt-injection-scanner" || f.Severity != report.SeverityHigh {
		t.Fatalf("unexpected finding: %+v", f)
	}
}

func TestParsePromptInjectionScannerOutputWrapper(t *testing.T) {
	data := []byte(`{"findings": [{"rule": "pi2", "path": "x.go", "line": 7, "severity": "CRITICAL", "message": "wrapped"}]}`)
	out := report.NewReport("target")
	out, err := parsePromptInjectionScannerOutput(out, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(out.Findings))
	}
	if out.Findings[0].Severity != report.SeverityCritical {
		t.Fatalf("expected critical, got %s", out.Findings[0].Severity)
	}
}

func TestParseGitleaksOutput(t *testing.T) {
	data := []byte(`[{"RuleID": "aws", "Description": "AWS key", "StartLine": 7, "StartColumn": 8, "Match": "AKIA...", "Secret": "SECRET", "File": "test.go", "Tags": []}]`)
	out := report.NewReport("target")
	out, err := parseGitleaksOutput(out, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(out.Findings))
	}
	f := out.Findings[0]
	if f.RuleID != "aws" || f.Match != "SECRET" || f.Source != "gitleaks" {
		t.Fatalf("unexpected finding: %+v", f)
	}
}

func TestParseTrufflehogOutput(t *testing.T) {
	data := []byte(`{"DetectorName": "AWS", "Verified": true, "Raw": "secret", "SourceMetadata": {"Data": {"Filesystem": {"file": "test.go", "line": 9}}}}
`)
	out := report.NewReport("target")
	out, err := parseTrufflehogOutput(out, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(out.Findings))
	}
	f := out.Findings[0]
	if f.RuleID != "aws" || f.Severity != report.SeverityCritical || f.Source != "trufflehog" {
		t.Fatalf("unexpected finding: %+v", f)
	}
}

func TestMapSeverity(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"HIGH", report.SeverityHigh},
		{"error", report.SeverityHigh},
		{"medium", report.SeverityMedium},
		{"warn", report.SeverityMedium},
		{"low", report.SeverityLow},
		{"", report.SeverityInfo},
		{"unknown", report.SeverityInfo},
	}
	for _, c := range cases {
		if got := mapSeverity(c.in); got != c.want {
			t.Errorf("mapSeverity(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCoalesce(t *testing.T) {
	if got := coalesce("", "b", "c"); got != "b" {
		t.Errorf("coalesce = %q, want b", got)
	}
	if got := coalesce("", "", ""); got != "" {
		t.Errorf("coalesce = %q, want empty", got)
	}
}
