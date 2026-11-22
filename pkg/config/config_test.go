package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAndWriteExample(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "safeanalyze.yaml")
	if err := WriteExample(path); err != nil {
		t.Fatalf("WriteExample failed: %v", err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if !cfg.YARA.Enabled || !cfg.Intent.Enabled || !cfg.Intent.FailOnTierA {
		t.Errorf("expected default security checks enabled: %+v", cfg)
	}
	if cfg.ML.Enabled {
		t.Errorf("expected ML disabled by default")
	}
}

func TestEffectiveExcludedPaths(t *testing.T) {
	s := SanitizationConfig{
		ExcludedPaths:   []string{".git"},
		DependencyPaths: []string{"node_modules", "vendor"},
	}
	fast := s.EffectiveExcludedPaths("fast")
	if len(fast) != 3 {
		t.Fatalf("fast mode expected 3 excluded paths, got %v", fast)
	}
	thorough := s.EffectiveExcludedPaths("thorough")
	if len(thorough) != 1 || thorough[0] != ".git" {
		t.Fatalf("thorough mode expected [.git], got %v", thorough)
	}
}

func TestOutputConfigHelpers(t *testing.T) {
	if got := (OutputConfig{MaxFindings: -5}).EffectiveMaxFindings(); got != 0 {
		t.Errorf("negative max_findings should normalize to 0, got %d", got)
	}
	if got := (OutputConfig{MaxFindings: 100}).EffectiveMaxFindings(); got != 100 {
		t.Errorf("expected 100, got %d", got)
	}
}

func TestLoadPluginYAMLFields(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "cfg.yaml")
	raw := `scanners:
  - name: osv-scanner
    command: "osv-scanner scan --format sarif {path}"
    format: sarif
    question: deps
    enabled: true
    timeout_seconds: 120
`
	if err := os.WriteFile(p, []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(p)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if len(cfg.Scanners) != 1 {
		t.Fatalf("expected 1 scanner, got %d", len(cfg.Scanners))
	}
	sc := cfg.Scanners[0]
	if sc.Format != "sarif" || sc.Question != "deps" || sc.TimeoutSeconds != 120 {
		t.Errorf("unexpected scanner config: %+v", sc)
	}
}
