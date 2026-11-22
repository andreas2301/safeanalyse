package external

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/user/safeanalyze/pkg/config"
	"github.com/user/safeanalyze/pkg/pipeline"
	"github.com/user/safeanalyze/pkg/report"
)

func TestParseBumblebeeRealSchemaNDJSON(t *testing.T) {
	// Mirrors the real Perplexity Bumblebee v0.1.2 NDJSON output:
	// package record (skip), finding record (keep), scan_summary record (skip).
	data := []byte(`{"record_type":"package","schema_version":"0.2.0","scanner_name":"bumblebee","ecosystem":"npm","package_name":"lodash","version":"4.17.21","source_file":"package-lock.json"}
{"record_type":"finding","schema_version":"0.2.0","scanner_name":"bumblebee","finding_type":"package_exposure","severity":"critical","catalog_id":"advisory-2026-0042","catalog_name":"evil-pkg 1.2.3 (compromised release)","ecosystem":"npm","package_name":"evil-pkg","normalized_name":"evil-pkg","version":"1.2.3","source_type":"npm-lockfile","source_file":"web/package-lock.json","confidence":"high","evidence":"exact name+version match (version=1.2.3)"}
{"record_type":"scan_summary","schema_version":"0.2.0","scanner_name":"bumblebee","findings_count":1}
`)
	out := report.NewReport("target")
	out, err := parseBumblebeeOutput(out, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Findings) != 1 {
		t.Fatalf("expected 1 finding (package & scan_summary skipped), got %d", len(out.Findings))
	}
	f := out.Findings[0]
	if f.RuleID != "advisory-2026-0042" {
		t.Errorf("RuleID = %q, want advisory-2026-0042", f.RuleID)
	}
	if f.Severity != report.SeverityCritical {
		t.Errorf("Severity = %q, want critical", f.Severity)
	}
	if f.File != "web/package-lock.json" {
		t.Errorf("File = %q, want web/package-lock.json", f.File)
	}
	if f.Category != "supply_chain" || f.Source != "bumblebee" {
		t.Errorf("unexpected category/source: %s/%s", f.Category, f.Source)
	}
	if f.Attributes["ecosystem"] != "npm" || f.Attributes["package"] != "evil-pkg" || f.Attributes["version"] != "1.2.3" {
		t.Errorf("unexpected attributes: %+v", f.Attributes)
	}
	if !strings.Contains(f.Message, "exact name+version match") {
		t.Errorf("expected evidence in message, got %q", f.Message)
	}
}

func TestGitleaksDoesNotEmbedRawSecretInMessage(t *testing.T) {
	rawSecret := "tok_" + "supersecretvalue_123456"
	data := []byte(`[{"RuleID":"custom-token","Description":"Custom API token","StartLine":3,"StartColumn":1,"Match":"` + rawSecret + `","Secret":"` + rawSecret + `","File":"config.env","Tags":[]}]`)
	out := report.NewReport("target")
	out, err := parseGitleaksOutput(out, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(out.Findings))
	}
	if strings.Contains(out.Findings[0].Message, rawSecret) {
		t.Errorf("gitleaks embedded raw secret in Message: %q", out.Findings[0].Message)
	}
}

func TestSplitArgs(t *testing.T) {
	got, err := splitArgs(`osv-scanner scan --format sarif --label "my repo" {path}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"osv-scanner", "scan", "--format", "sarif", "--label", "my repo", "{path}"}
	if len(got) != len(want) {
		t.Fatalf("len=%d (%v), want %d (%v)", len(got), got, len(want), want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("token[%d] = %q, want %q", i, got[i], want[i])
		}
	}
	if _, err := splitArgs(`bad "unterminated`); err == nil {
		t.Error("expected error for unterminated quote")
	}
}

func TestGenericPluginParsers(t *testing.T) {
	sarifData := []byte(`{
		"version": "2.1.0",
		"runs": [{
			"results": [{
				"ruleId": "OSV-2026-001",
				"level": "error",
				"message": {"text": "Vulnerable dependency"},
				"locations": [{"physicalLocation": {"artifactLocation": {"uri": "go.mod"}, "region": {"startLine": 12, "startColumn": 1}}}]
			}]
		}]
	}`)
	out, err := parseSARIFPluginOutput(report.NewReport("t"), "osv-scanner", "deps", sarifData)
	if err != nil || len(out.Findings) != 1 {
		t.Fatalf("sarif parse failed: err=%v findings=%d", err, len(out.Findings))
	}
	if out.Findings[0].RuleID != "OSV-2026-001" || out.Findings[0].Question != "deps" || out.Findings[0].Severity != report.SeverityHigh {
		t.Errorf("unexpected sarif finding: %+v", out.Findings[0])
	}

	saJSON := []byte(`{
		"target": "t",
		"findings": [{"rule_id": "llm.promptguard2", "title": "Injection", "severity": "high", "file": "AGENTS.md", "line": 4, "message": "score=0.99", "source": "llm", "confidence": "ml"}],
		"metadata": {"safeanalyze_version": "0.0.0-override", "llm_model": "prompt-guard-2-22m"}
	}`)
	out2, err := parseSafeanalyzeJSONPluginOutput(report.NewReport("t"), "safeanalyze-llm", "agent", saJSON)
	if err != nil || len(out2.Findings) != 1 {
		t.Fatalf("safeanalyze-json parse failed: err=%v findings=%d", err, len(out2.Findings))
	}
	if out2.Metadata["llm_model"] != "prompt-guard-2-22m" {
		t.Errorf("expected llm_model metadata preserved, got %v", out2.Metadata)
	}
	if _, exists := out2.Metadata["safeanalyze_version"]; exists {
		t.Errorf("companion must not overwrite safeanalyze_version metadata")
	}

	ndjson := []byte("{\"rule\":\"CUSTOM-1\",\"severity\":\"critical\",\"file\":\"a.py\",\"line\":7,\"message\":\"bad call\"}\n")
	out3, err := parseNDJSONPluginOutput(report.NewReport("t"), "custom", "remote_exec", ndjson)
	if err != nil || len(out3.Findings) != 1 {
		t.Fatalf("ndjson parse failed: err=%v findings=%d", err, len(out3.Findings))
	}
	if out3.Findings[0].Severity != report.SeverityCritical || out3.Findings[0].Question != "remote_exec" {
		t.Errorf("unexpected ndjson finding: %+v", out3.Findings[0])
	}
}

func TestGenericPluginLiveExecutionAndRegistration(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX shell script fixture")
	}
	dir := t.TempDir()
	script := filepath.Join(dir, "fake-scanner.sh")
	body := `#!/bin/sh
echo '{"version":"2.1.0","runs":[{"results":[{"ruleId":"FAKE-1","level":"error","message":{"text":"found issue in '"$1"'"},"locations":[{"physicalLocation":{"artifactLocation":{"uri":"app.go"},"region":{"startLine":2}}}]}]}]}'
exit 1
`
	if err := os.WriteFile(script, []byte(body), 0755); err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{
		Scanners: []config.ScannerConfig{
			{
				Name:           "fake-sarif",
				Command:        script + " {path}",
				Format:         "sarif",
				Question:       "deps",
				Enabled:        true,
				TimeoutSeconds: 10,
			},
		},
	}
	reg := pipeline.NewRegistry()
	if err := RegisterAll(reg, cfg); err != nil {
		t.Fatalf("RegisterAll failed: %v", err)
	}
	stages, err := reg.Build([]string{"fake-sarif"})
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	rep, err := stages[0].Run(context.Background(), "/tmp/demo", nil)
	if err != nil {
		t.Fatalf("stage Run failed: %v", err)
	}
	if len(rep.Errors) != 0 {
		t.Fatalf("expected 0 errors on exit-1 SARIF findings, got %+v", rep.Errors)
	}
	if len(rep.Findings) != 1 || rep.Findings[0].RuleID != "FAKE-1" {
		t.Fatalf("unexpected findings: %+v", rep.Findings)
	}
}
