package external

import (
	"github.com/user/safeanalyze/pkg/config"
	"github.com/user/safeanalyze/pkg/pipeline"
)

// Compile-time interface checks.
var (
	_ pipeline.Stage = (*semgrepStage)(nil)
	_ pipeline.Stage = (*bumblebeeStage)(nil)
	_ pipeline.Stage = (*promptInjectionScannerStage)(nil)
	_ pipeline.Stage = (*gitleaksStage)(nil)
	_ pipeline.Stage = (*trufflehogStage)(nil)
)

// RegisterAll registers enabled external scanner stages with the provided registry.
// A scanner is registered only if it appears in cfg.Scanners with Enabled=true.
func RegisterAll(r *pipeline.Registry, cfg *config.Config) error {
	enabled := make(map[string]bool)
	for _, s := range cfg.Scanners {
		if s.Enabled {
			enabled[s.Name] = true
		}
	}

	factories := map[string]func() pipeline.Stage{
		"semgrep":                  func() pipeline.Stage { return NewSemgrepStage(cfg.Sanitization.ExcludedPaths) },
		"bumblebee":                func() pipeline.Stage { return NewBumblebeeStage() },
		"prompt-injection-scanner": func() pipeline.Stage { return NewPromptInjectionScannerStage() },
		"gitleaks":                 func() pipeline.Stage { return NewGitleaksStage() },
		"trufflehog":               func() pipeline.Stage { return NewTrufflehogStage() },
	}

	for name, factory := range factories {
		if enabled[name] {
			r.Register(name, factory)
		}
	}
	return nil
}
