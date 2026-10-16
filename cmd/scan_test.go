package cmd

import (
	"strings"
	"testing"

	"github.com/user/safeanalyze/pkg/report"
)

func TestBlockingFindingsScopedToSource(t *testing.T) {
	rep := report.NewReport("t")
	rep.AddFinding(report.Finding{RuleID: "trufflehog_secret", Severity: report.SeverityHigh, Source: "trufflehog", File: "a.go", Line: 1})
	rep.AddFinding(report.Finding{RuleID: "prompt_injection_comment", Severity: report.SeverityHigh, Source: "yara", File: "b.md", Line: 2})

	if got := blockingFindings(rep, "hiddenchars", report.SeverityHigh); len(got) != 0 {
		t.Fatalf("expected no hiddenchars blocking findings, got %d", len(got))
	}

	rep.AddFinding(report.Finding{RuleID: "hidden_char_zero_width", Severity: report.SeverityHigh, Source: "hiddenchars", File: "c.md", Line: 3})
	got := blockingFindings(rep, "hiddenchars", report.SeverityHigh)
	if len(got) != 1 || got[0].RuleID != "hidden_char_zero_width" {
		t.Fatalf("expected the hiddenchars finding only, got %+v", got)
	}
	if msg := describeFindings(got); msg != "hidden_char_zero_width at c.md:3" {
		t.Errorf("unexpected description %q", msg)
	}
}

func TestDescribeFindingsTruncates(t *testing.T) {
	var fs []report.Finding
	for i := 1; i <= 5; i++ {
		fs = append(fs, report.Finding{RuleID: "r", File: "f", Line: i})
	}
	if msg := describeFindings(fs); !strings.HasSuffix(msg, "(+2 more)") {
		t.Errorf("expected truncation suffix, got %q", msg)
	}
}
