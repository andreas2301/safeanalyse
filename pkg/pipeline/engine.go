package pipeline

import (
	"context"
	"fmt"
	"sync"

	"github.com/user/safeanalyze/pkg/report"
)

// Engine executes pipeline stages according to their dependency DAG.
type Engine struct {
	stages        []Stage
	maxParallelism int
}

// NewEngine creates an engine from a list of stages. Stages are topologically
// sorted on first Run.
func NewEngine(stages []Stage, maxParallelism int) *Engine {
	if maxParallelism <= 0 {
		maxParallelism = 4
	}
	return &Engine{
		stages:         stages,
		maxParallelism: maxParallelism,
	}
}

// Run executes all stages against target and returns a merged report.
// Context cancellation aborts running and pending stages.
func (e *Engine) Run(ctx context.Context, target string) (*report.Report, error) {
	ordered, err := e.topologicalSort()
	if err != nil {
		return nil, err
	}

	// Reports indexed by stage name.
	reports := make(map[string]*report.Report)
	var mu sync.Mutex

	// Track completed stages.
	completed := make(map[string]bool)
	cond := sync.NewCond(&mu)

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	// Worker semaphore to limit parallelism.
	sem := make(chan struct{}, e.maxParallelism)

	for _, stage := range ordered {
		stage := stage
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Wait for dependencies.
			mu.Lock()
			for _, dep := range stage.Dependencies() {
				for !completed[dep] {
					if ctx.Err() != nil {
						mu.Unlock()
						return
					}
					cond.Wait()
				}
			}
			mu.Unlock()

			// Build input report from dependencies. Stages receive an empty report
			// by default; dependencies only enforce execution order. If a stage
			// needs prior findings, that can be added later via a context mechanism.
			input := report.NewReport(target)
			mu.Lock()
			for _, dep := range stage.Dependencies() {
				if r, ok := reports[dep]; ok {
					// Carry forward metadata only, not findings, to avoid duplication.
					for k, v := range r.Metadata {
						if _, exists := input.Metadata[k]; !exists {
							input.Metadata[k] = v
						}
					}
					if r.Summary.FilesScanned > input.Summary.FilesScanned {
						input.Summary.FilesScanned = r.Summary.FilesScanned
					}
				}
			}
			mu.Unlock()

			// Acquire slot.
			select {
			case sem <- struct{}{}:
			case <-ctx.Done():
				return
			}

			out, runErr := stage.Run(ctx, target, input)
			<-sem

			if runErr != nil {
				select {
				case errCh <- fmt.Errorf("stage %q failed: %w", stage.Name(), runErr):
				default:
				}
			}

			mu.Lock()
			if out != nil {
				reports[stage.Name()] = out
			} else {
				reports[stage.Name()] = report.NewReport(target)
			}
			completed[stage.Name()] = true
			cond.Broadcast()
			mu.Unlock()
		}()
	}

	// Wait for completion or cancellation.
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	// Merge all stage reports deterministically.
	var all []*report.Report
	for _, stage := range ordered {
		if r, ok := reports[stage.Name()]; ok {
			all = append(all, r)
		}
	}
	merged := report.Merge(target, all)
	merged.Finish()
	return merged, nil
}

// topologicalSort returns stages ordered so dependencies come first.
func (e *Engine) topologicalSort() ([]Stage, error) {
	byName := make(map[string]Stage)
	for _, s := range e.stages {
		if _, exists := byName[s.Name()]; exists {
			return nil, fmt.Errorf("duplicate stage name %q", s.Name())
		}
		byName[s.Name()] = s
	}

	// Kahn's algorithm.
	inDegree := make(map[string]int)
	for name := range byName {
		inDegree[name] = 0
	}
	dependents := make(map[string][]string)
	for name, s := range byName {
		for _, dep := range s.Dependencies() {
			if _, ok := byName[dep]; !ok {
				return nil, fmt.Errorf("stage %q depends on unknown stage %q", name, dep)
			}
			inDegree[name]++
			dependents[dep] = append(dependents[dep], name)
		}
	}

	// Seed queue with stages that have zero dependencies, preserving registration order.
	var queue []string
	for _, s := range e.stages {
		if inDegree[s.Name()] == 0 {
			queue = append(queue, s.Name())
		}
	}

	var ordered []Stage
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		ordered = append(ordered, byName[name])
		for _, dep := range dependents[name] {
			inDegree[dep]--
			if inDegree[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	if len(ordered) != len(e.stages) {
		return nil, fmt.Errorf("cycle detected in pipeline dependencies")
	}
	return ordered, nil
}
