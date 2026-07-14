package external

import (
	"testing"

	"github.com/user/safeanalyze/pkg/config"
	"github.com/user/safeanalyze/pkg/pipeline"
)

func TestRegisterAll(t *testing.T) {
	cfg := &config.Config{
		Scanners: []config.ScannerConfig{
			{Name: "semgrep", Enabled: true},
			{Name: "gitleaks", Enabled: true},
			{Name: "trufflehog", Enabled: false},
		},
	}
	r := pipeline.NewRegistry()
	if err := RegisterAll(r, cfg); err != nil {
		t.Fatalf("RegisterAll failed: %v", err)
	}

	names := r.Names()
	if len(names) != 2 {
		t.Fatalf("expected 2 registered stages, got %d: %v", len(names), names)
	}
	for _, name := range names {
		if name != "semgrep" && name != "gitleaks" {
			t.Fatalf("unexpected stage %q", name)
		}
	}
}
