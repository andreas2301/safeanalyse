// Package fastinspect provides a zero-disk, precompiled in-memory fast
// prompt-injection inspector shared by inspect, icap, and eval.
package fastinspect

import (
	"sync"

	"github.com/user/safeanalyze/pkg/checks/hiddenchars"
	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/report"
)

const virtualFilename = "payload.txt"

// Inspector runs the fast prompt-injection check suite (built-in YARA rules +
// high-signal hidden-character detection) purely in memory using precompiled
// engines. An Inspector is immutable after construction and safe for
// concurrent use by multiple goroutines.
type Inspector struct {
	yara   *yara.Engine
	hidden *hiddenchars.Detector
}

var (
	defaultOnce      sync.Once
	defaultInspector *Inspector
)

// New compiles the fast-path detection engines and returns a new Inspector.
func New() *Inspector {
	return &Inspector{
		yara:   yara.NewEngine(),
		hidden: hiddenchars.NewDetector([]string{"zero_width", "bidi", "control"}),
	}
}

// Default returns a process-wide singleton Inspector whose regexes are compiled
// once on first use.
func Default() *Inspector {
	defaultOnce.Do(func() {
		defaultInspector = New()
	})
	return defaultInspector
}

// Scan inspects payload in memory and returns a deterministic Report matching
// the legacy yara + hiddenchars fast-stage pipeline.
func (ins *Inspector) Scan(payload string) *report.Report {
	rep := report.NewReport(virtualFilename)
	if yara.IsBinaryContent([]byte(payload)) {
		rep.Finish()
		return rep
	}

	for _, m := range ins.yara.ScanFile(payload, virtualFilename) {
		rep.AddFinding(report.Finding{
			RuleID:      m.Rule,
			Title:       m.Rule,
			Description: m.Description,
			Severity:    m.Severity,
			Category:    "pattern_match",
			File:        m.File,
			Line:        m.Line,
			Column:      m.Column,
			Match:       m.Match,
			Message:     m.Description,
			Source:      "yara",
			Confidence:  report.ConfidenceDeterministic,
		})
	}

	for _, f := range ins.hidden.ScanContent(payload, virtualFilename, true) {
		rep.AddFinding(f)
	}

	report.SortFindings(rep.Findings)
	rep.Summary.FilesScanned = 1
	rep.Finish()
	return rep
}
