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
	_ pipeline.Stage = (*genericPluginStage)(nil)
)

// RegisterAll registers enabled external scanner stages and declarative plugins
// with the provided registry. A scanner is registered only if it appears in
// cfg.Scanners with Enabled=true.
func RegisterAll(r *pipeline.Registry, cfg *config.Config) error {
	for _, s := range cfg.Scanners {
		if !s.Enabled {
			continue
		}
		sc := s
		timeout := effectiveTimeout(sc.TimeoutSeconds)
		switch sc.Name {
		case "semgrep":
			r.Register("semgrep", func() pipeline.Stage {
				return NewSemgrepStageWithOptions(cfg.Sanitization.ExcludedPaths, sc.RemoteRules, timeout)
			})
		case "bumblebee":
			r.Register("bumblebee", func() pipeline.Stage {
				return NewBumblebeeStageWithOptions(sc.ExposureCatalog, timeout)
			})
		case "prompt-injection-scanner":
			r.Register("prompt-injection-scanner", func() pipeline.Stage {
				return NewPromptInjectionScannerStageWithTimeout(timeout)
			})
		case "gitleaks":
			r.Register("gitleaks", func() pipeline.Stage {
				return NewGitleaksStageWithTimeout(timeout)
			})
		case "trufflehog":
			// Offline-first default: disable live verification unless Verify is
			// explicitly opted in (and NoVerification is not forced true).
			noVerify := !sc.Verify || sc.NoVerification
			r.Register("trufflehog", func() pipeline.Stage {
				return NewTrufflehogStageWithTimeout(noVerify, timeout)
			})
		default:
			if sc.Command != "" && sc.Format != "" {
				r.Register(sc.Name, func() pipeline.Stage {
					return NewGenericPluginStage(sc)
				})
			}
		}
	}
	return nil
}
