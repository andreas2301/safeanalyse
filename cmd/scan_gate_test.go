package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// runScan executes `scan` with only the built-in, offline stages enabled.
func runScan(t *testing.T, target string, extra ...string) (string, error) {
	t.Helper()
	out := t.TempDir()
	cfgPath := filepath.Join(t.TempDir(), "cfg.yaml")
	// No external scanners: the test must not depend on installed tools.
	if err := os.WriteFile(cfgPath, []byte("scanners: []\noutput:\n  formats: [markdown, json]\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	args := append([]string{"scan", target, "-c", cfgPath, "-o", out}, extra...)
	rootCmd.SetArgs(args)
	defer rootCmd.SetArgs(nil)
	// Flags keep their values between Execute calls; reset the ones we use.
	_ = scanCmd.Flags().Set("no-fail", "false")
	err := rootCmd.Execute()
	md, _ := os.ReadFile(filepath.Join(out, "safeanalyze.md"))
	return string(md), err
}

func TestScanBlocksOnTierA(t *testing.T) {
	md, err := runScan(t, "../testdata/intent/malicious/npm-postinstall-env-exfil")
	var exitErr *ExitError
	if !errors.As(err, &exitErr) || exitErr.Code != ExitBlocked {
		t.Fatalf("expected ExitError with code %d, got %v", ExitBlocked, err)
	}
	for _, want := range []string{"## Verdict", "## Call-home map", "## Auto-run surface", "discord.com", "## Coverage"} {
		if !strings.Contains(md, want) {
			t.Errorf("report misses %q", want)
		}
	}
}

func TestScanNoFailReportsOnly(t *testing.T) {
	if _, err := runScan(t, "../testdata/intent/malicious/npm-postinstall-env-exfil", "--no-fail"); err != nil {
		t.Fatalf("--no-fail must not fail the scan, got %v", err)
	}
}

func TestScanBenignTwinPasses(t *testing.T) {
	md, err := runScan(t, "../testdata/intent/benign/npm-postinstall-env-exfil")
	if err != nil {
		t.Fatalf("benign twin blocked: %v\n%s", err, md)
	}
}
