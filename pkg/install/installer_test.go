package install

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveBinary(t *testing.T) {
	home, err := safeanalyzeHome()
	if err != nil {
		t.Skipf("cannot determine home: %v", err)
	}
	bin := filepath.Join(home, "bin")
	if err := os.MkdirAll(bin, 0755); err != nil {
		t.Skipf("cannot create bin dir: %v", err)
	}

	fake := filepath.Join(bin, "gitleaks")
	f, err := os.Create(fake)
	if err != nil {
		t.Fatalf("creating fake binary: %v", err)
	}
	f.Close()
	defer os.Remove(fake)

	if got := ResolveBinary("gitleaks"); got != fake {
		t.Fatalf("ResolveBinary(gitleaks) = %q, want %q", got, fake)
	}
	if got := ResolveBinary("nonexistent"); got != "" {
		t.Fatalf("ResolveBinary(nonexistent) = %q, want empty", got)
	}
}

func TestKnownScannersContainsExpected(t *testing.T) {
	expected := []string{"semgrep", "bumblebee", "prompt-injection-scanner", "gitleaks", "trufflehog"}
	for _, name := range expected {
		if _, ok := KnownScanners[name]; !ok {
			t.Errorf("KnownScanners missing %q", name)
		}
	}
}
