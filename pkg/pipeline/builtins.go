package pipeline

import (
	"github.com/user/safeanalyze/pkg/checks/entropy"
	"github.com/user/safeanalyze/pkg/checks/hiddenchars"
	"github.com/user/safeanalyze/pkg/checks/ml"
	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/config"
)

// BuiltInRegistry returns a registry populated with safeanalyze's built-in stages.
func BuiltInRegistry(cfg *config.Config) *Registry {
	r := NewRegistry()
	r.Register("yara", func() Stage {
		return yara.NewStage(cfg.Sanitization.ExcludedPaths)
	})
	r.Register("entropy", func() Stage {
		return entropy.NewStage(cfg.Entropy.Threshold, cfg.Entropy.MaxFileSizeBytes, cfg.Sanitization.ExcludedPaths)
	})
	r.Register("hiddenchars", func() Stage {
		return hiddenchars.NewStage(cfg.HiddenChars.Categories, cfg.Sanitization.ExcludedPaths)
	})
	r.Register("ml-prompt-injection", func() Stage {
		return ml.NewStage(&cfg.ML)
	})
	return r
}

// DefaultStages returns the default ordered list of built-in stage names.
func DefaultStages() []string {
	return []string{"yara", "entropy", "hiddenchars", "ml-prompt-injection"}
}
