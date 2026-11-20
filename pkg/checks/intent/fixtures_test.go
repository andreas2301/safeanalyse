package intent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/user/safeanalyze/pkg/report"
)

// runAll runs every intent analyzer on one fixture directory.
func runAll(t *testing.T, dir string) []report.Finding {
	t.Helper()
	files, err := LoadFiles(dir, nil)
	if err != nil {
		t.Fatalf("loading %s: %v", dir, err)
	}
	var out []report.Finding
	for _, a := range Analyzers() {
		out = append(out, a.Run(files)...)
	}
	return out
}

// TestMaliciousFixturesReachTierA is the recall gate: every malicious fixture
// must produce a Tier A finding for each question listed in its EXPECT file.
func TestMaliciousFixturesReachTierA(t *testing.T) {
	dirs, _ := filepath.Glob("../../../testdata/intent/malicious/*")
	if len(dirs) == 0 {
		t.Fatal("no malicious fixtures found")
	}
	for _, dir := range dirs {
		dir := dir
		t.Run(filepath.Base(dir), func(t *testing.T) {
			raw, err := os.ReadFile(filepath.Join(dir, "EXPECT"))
			if err != nil {
				t.Fatalf("missing EXPECT: %v", err)
			}
			got := map[string]bool{}
			findings := runAll(t, dir)
			for _, f := range findings {
				if f.Tier == report.TierA {
					got[f.Question] = true
				}
			}
			for _, q := range strings.Fields(string(raw)) {
				if !got[q] {
					t.Errorf("no Tier A finding for question %q; findings: %s", q, describe(findings))
				}
			}
		})
	}
}

// TestBenignTwinsStayBelowTierA is the precision guard: the benign twin of
// each malicious fixture must not produce Tier A findings.
func TestBenignTwinsStayBelowTierA(t *testing.T) {
	dirs, _ := filepath.Glob("../../../testdata/intent/benign/*")
	if len(dirs) == 0 {
		t.Fatal("no benign fixtures found")
	}
	for _, dir := range dirs {
		dir := dir
		t.Run(filepath.Base(dir), func(t *testing.T) {
			var a []report.Finding
			for _, f := range runAll(t, dir) {
				if f.Tier == report.TierA {
					a = append(a, f)
				}
			}
			if len(a) > 0 {
				t.Errorf("benign fixture has Tier A findings: %s", describe(a))
			}
		})
	}
}

func describe(fs []report.Finding) string {
	var parts []string
	for _, f := range fs {
		parts = append(parts, f.Tier+":"+f.Question+":"+f.RuleID+"@"+f.File)
	}
	return strings.Join(parts, ", ")
}
