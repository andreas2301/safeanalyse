// Package pipeline orchestrates safeanalyze checks as a deterministic DAG of stages.
package pipeline

import (
	"context"

	"github.com/user/safeanalyze/pkg/report"
)

// Stage is a single step in the analysis pipeline.
type Stage interface {
	// Name returns the unique name of this stage.
	Name() string
	// Dependencies returns the names of stages that must complete before this one runs.
	Dependencies() []string
	// Run executes the stage against the target path. input may be nil for the first stage.
	// A stage should append its findings/errors to a copy or to input and return the result.
	Run(ctx context.Context, target string, input *report.Report) (*report.Report, error)
}
