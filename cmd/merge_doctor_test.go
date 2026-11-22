package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/user/safeanalyze/pkg/config"
	"github.com/user/safeanalyze/pkg/report"
)

func TestMergeReportsCombinesAndGatesTierA(t *testing.T) {
	dir := t.TempDir()
	outDir := filepath.Join(dir, "merged")

	r1 := report.NewReport("demo-repo")
	r1.Metadata["safeanalyze_version"] = "0.4.0"
	r1.AddFinding(report.Finding{
		RuleID:   "egress_unknown",
		Title:    "Outbound host",
		Severity: report.SeverityMedium,
		File:     "main.go",
		Line:     10,
		Message:  "contacts example.org",
		Source:   "egress",
		Tier:     report.TierB,
		Question: "call_home",
	})

	r2 := report.NewReport("demo-repo")
	r2.Metadata["llm_model"] = "prompt-guard-2-22m"
	r2.AddFinding(report.Finding{
		RuleID:     "llm.promptguard2",
		Title:      "Prompt injection detected",
		Severity:   report.SeverityHigh,
		File:       "AGENTS.md",
		Line:       3,
		Message:    "score=0.99",
		Source:     "llm",
		Confidence: report.ConfidenceML,
	})

	p1 := filepath.Join(dir, "r1.json")
	p2 := filepath.Join(dir, "r2.json")
	for p, r := range map[string]*report.Report{p1: r1, p2: r2} {
		b, err := json.Marshal(r)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, b, 0644); err != nil {
			t.Fatal(err)
		}
	}

	cfg = config.DefaultConfig()
	cfg.Output.Formats = []string{"json", "sarif"}
	cfg.Output.OutDir = outDir

	_ = mergeCmd.Flags().Set("output-dir", outDir)
	_ = mergeCmd.Flags().Set("no-fail", "false")
	err := mergeCmd.RunE(mergeCmd, []string{p1, p2})
	var exitErr *ExitError
	if !errors.As(err, &exitErr) || exitErr.Code != ExitBlocked {
		t.Fatalf("expected ExitBlocked (2) on Tier A merged finding, got %v", err)
	}

	mergedRaw, err := os.ReadFile(filepath.Join(outDir, "safeanalyze.json"))
	if err != nil {
		t.Fatalf("reading merged json: %v", err)
	}
	var merged report.Report
	if err := json.Unmarshal(mergedRaw, &merged); err != nil {
		t.Fatalf("unmarshaling merged json: %v", err)
	}
	if len(merged.Findings) != 2 {
		t.Fatalf("expected 2 merged findings, got %d", len(merged.Findings))
	}
	if merged.Metadata["llm_model"] != "prompt-guard-2-22m" {
		t.Errorf("missing companion metadata: %v", merged.Metadata)
	}
	if _, err := os.Stat(filepath.Join(outDir, "safeanalyze.sarif")); err != nil {
		t.Errorf("expected merged SARIF report: %v", err)
	}
}

func TestRunDoctorReturnsStructuredChecks(t *testing.T) {
	cfg = config.DefaultConfig()
	cfg.Scanners = append(cfg.Scanners, config.ScannerConfig{Name: "bumblebee", Enabled: false})
	rep := runDoctor()
	if rep.Version == "" || rep.OS == "" {
		t.Fatalf("missing version/os metadata: %+v", rep)
	}
	if len(rep.Checks) == 0 {
		t.Fatalf("expected doctor checks")
	}
	seenGit := false
	seenCatalog := false
	for _, c := range rep.Checks {
		if c.Name == "git" {
			seenGit = true
		}
		if c.Name == "bumblebee-catalog" {
			seenCatalog = true
		}
	}
	if !seenGit || !seenCatalog {
		t.Errorf("expected git and bumblebee-catalog checks, got %+v", rep.Checks)
	}
}
