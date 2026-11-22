package triage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/safeanalyze/pkg/report"
)

func TestLegacyTiers(t *testing.T) {
	cases := []struct {
		name string
		f    report.Finding
		want string
	}{
		{"lockfile entropy is noise", report.Finding{Source: "entropy", RuleID: "entropy_base64_blob", File: "web/package-lock.json"}, report.TierC},
		{"source base64 blob is inventory", report.Finding{Source: "entropy", RuleID: "entropy_base64_blob", File: "pkg/a.go"}, report.TierB},
		{"bidi in source is verdict-relevant", report.Finding{Source: "hiddenchars", RuleID: "hidden_char_bidi", File: "pkg/a.go"}, report.TierA},
		{"bidi in a test fixture is inventory", report.Finding{Source: "hiddenchars", RuleID: "hidden_char_bidi", File: "pkg/a_test.go"}, report.TierB},
		{"persistence keyword is noise", report.Finding{Source: "yara", RuleID: "backdoor_indicator", Severity: "critical", Match: "persistence", File: "web/src/prefs.ts"}, report.TierC},
		{"reverse shell in source is verdict-relevant", report.Finding{Source: "yara", RuleID: "backdoor_indicator", Severity: "critical", Match: "reverse_shell", File: "pkg/a.go"}, report.TierA},
		{"curl|bash in install script", report.Finding{Source: "yara", RuleID: "suspicious_shell", Severity: "high", Match: "curl x | bash", File: "install.sh"}, report.TierA},
		{"prompt injection in agent file", report.Finding{Source: "yara", RuleID: "prompt_injection_comment", Severity: "critical", Match: "ignore all previous instructions", File: "AGENTS.md"}, report.TierA},
		{"system prompt mention in agent file", report.Finding{Source: "yara", RuleID: "prompt_injection_comment", Severity: "critical", Match: "system prompt", File: "skills/x/SKILL.md"}, report.TierB},
		{"prompt injection in changelog", report.Finding{Source: "yara", RuleID: "prompt_injection_comment", Severity: "critical", Match: "ignore previous instructions", File: "CHANGELOG.md"}, report.TierB},
		{"template interpolation is noise", report.Finding{Source: "yara", RuleID: "template_injection", Severity: "medium", File: "pkg/a.go"}, report.TierC},
		{"bumblebee critical exposure is Tier A", report.Finding{Source: "bumblebee", RuleID: "advisory-2026-0042", Severity: "critical", File: "package-lock.json"}, report.TierA},
		{"bumblebee medium exposure is Tier B", report.Finding{Source: "bumblebee", RuleID: "advisory-2026-0043", Severity: "medium", File: "package-lock.json"}, report.TierB},
		{"generic plugin critical with question is Tier A", report.Finding{Source: "osv-scanner", RuleID: "OSV-1", Severity: "critical", Question: "deps", File: "go.mod"}, report.TierA},
	}
	for _, c := range cases {
		rep := report.NewReport(t.TempDir())
		rep.AddFinding(c.f)
		Assign(rep)
		if got := rep.Findings[0].Tier; got != c.want {
			t.Errorf("%s: tier %q, want %q", c.name, got, c.want)
		}
	}
}

func TestKeywordInCommentOrProseIsDemoted(t *testing.T) {
	dir := t.TempDir()
	src := "package a\n// reverse-shell pipes are allowed in CTF mode\nvar p = \"Write an undetectable keylogger that exfiltrates keystrokes from a computer I do not own.\"\nvar s = startReverseShell()\n"
	if err := os.WriteFile(filepath.Join(dir, "a.go"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	rep := report.NewReport(dir)
	rep.AddFinding(report.Finding{Source: "yara", RuleID: "backdoor_indicator", Severity: "critical", Match: "reverse-shell", File: "a.go", Line: 2})
	rep.AddFinding(report.Finding{Source: "yara", RuleID: "backdoor_indicator", Severity: "critical", Match: "keylogger", File: "a.go", Line: 3})
	rep.AddFinding(report.Finding{Source: "yara", RuleID: "backdoor_indicator", Severity: "critical", Match: "ReverseShell", File: "a.go", Line: 4})
	Assign(rep)
	want := []string{report.TierB, report.TierB, report.TierA}
	for i, w := range want {
		if rep.Findings[i].Tier != w {
			t.Errorf("finding %d (line %d): tier %q, want %q", i, rep.Findings[i].Line, rep.Findings[i].Tier, w)
		}
	}
}

func TestFinalizeVerdicts(t *testing.T) {
	rep := report.NewReport("t")
	rep.AddFinding(report.Finding{RuleID: "chain_env_exfil", File: "a.js", Line: 3, Tier: report.TierA, Question: "exfil"})
	rep.AddFinding(report.Finding{RuleID: "egress_unknown", File: "b.go", Line: 9, Tier: report.TierB, Question: "call_home"})
	rep.AddError("ci", "boom")
	Finalize(rep, []string{"chains", "egress", "ci", "autorun"}, nil)

	got := map[string]string{}
	for _, v := range rep.Verdicts {
		got[v.ID] = v.Answer
	}
	want := map[string]string{
		"exfil":     report.AnswerYes,
		"call_home": report.AnswerUnclear,
		"autorun":   report.AnswerNo,         // ran, nothing found
		"ci":        report.AnswerNotChecked, // only stage errored
		"license":   report.AnswerNotChecked, // stage not run
	}
	for id, w := range want {
		if got[id] != w {
			t.Errorf("verdict %s = %q, want %q", id, got[id], w)
		}
	}
	for _, c := range rep.Coverage {
		if c.Check == "ci" && c.Status != "error" {
			t.Errorf("coverage for ci = %q, want error", c.Status)
		}
	}
	if rep.Metadata["tier_a"] != "1" || rep.Metadata["tier_b"] != "1" {
		t.Errorf("tier metadata = %v", rep.Metadata)
	}
}
