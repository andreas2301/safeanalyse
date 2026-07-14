package pipeline

import (
	"context"
	"fmt"
	"testing"

	"github.com/user/safeanalyze/pkg/report"
)

// mockStage is a deterministic test stage.
type mockStage struct {
	name         string
	deps         []string
	findings     []report.Finding
	shouldErr    bool
}

func (m *mockStage) Name() string { return m.name }
func (m *mockStage) Dependencies() []string { return m.deps }
func (m *mockStage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	if m.shouldErr {
		return nil, fmt.Errorf("stage %s error", m.name)
	}
	out := report.NewReport(target)
	if input != nil {
		out = input
	}
	for _, f := range m.findings {
		out.AddFinding(f)
	}
	out.Summary.FilesScanned = 1
	return out, nil
}

func TestEngineTopologicalOrder(t *testing.T) {
	stages := []Stage{
		&mockStage{name: "a", findings: []report.Finding{{RuleID: "a", File: "x.go", Line: 1, Source: "a"}}},
		&mockStage{name: "b", deps: []string{"a"}, findings: []report.Finding{{RuleID: "b", File: "x.go", Line: 2, Source: "b"}}},
		&mockStage{name: "c", deps: []string{"a"}, findings: []report.Finding{{RuleID: "c", File: "x.go", Line: 3, Source: "c"}}},
		&mockStage{name: "d", deps: []string{"b", "c"}, findings: []report.Finding{{RuleID: "d", File: "x.go", Line: 4, Source: "d"}}},
	}

	engine := NewEngine(stages, 4)
	rep, err := engine.Run(context.Background(), "target")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rep.Findings) != 4 {
		t.Fatalf("expected 4 findings, got %d", len(rep.Findings))
	}
	// Findings should be deterministically sorted by file/line/column/rule/source.
	expected := []string{"a", "b", "c", "d"}
	for i, f := range rep.Findings {
		if f.RuleID != expected[i] {
			t.Fatalf("finding %d: expected rule %s, got %s", i, expected[i], f.RuleID)
		}
	}
}

func TestEngineDetectsCycle(t *testing.T) {
	stages := []Stage{
		&mockStage{name: "a", deps: []string{"b"}},
		&mockStage{name: "b", deps: []string{"a"}},
	}
	engine := NewEngine(stages, 4)
	_, err := engine.Run(context.Background(), "target")
	if err == nil {
		t.Fatal("expected cycle error")
	}
}

func TestEngineDetectsUnknownDependency(t *testing.T) {
	stages := []Stage{
		&mockStage{name: "a", deps: []string{"missing"}},
	}
	engine := NewEngine(stages, 4)
	_, err := engine.Run(context.Background(), "target")
	if err == nil {
		t.Fatal("expected unknown dependency error")
	}
}

func TestEnginePropagatesStageError(t *testing.T) {
	stages := []Stage{
		&mockStage{name: "a"},
		&mockStage{name: "b", deps: []string{"a"}, shouldErr: true},
	}
	engine := NewEngine(stages, 4)
	_, err := engine.Run(context.Background(), "target")
	if err == nil {
		t.Fatal("expected stage error")
	}
}
