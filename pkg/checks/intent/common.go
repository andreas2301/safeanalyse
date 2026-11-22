// Package intent implements the "what does this repository actually do"
// checks: outbound endpoints (call home), code that runs automatically,
// dependency sources, install scripts, CI, browser extensions, behaviour
// chains, AI-agent instruction files and licenses.
//
// The checks favour recall: a pattern that may be malicious is reported and
// ranked with a triage tier instead of being dropped.
package intent

import (
	"context"
	"regexp"
	"strings"
	"sync"

	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/filerole"
	"github.com/user/safeanalyze/pkg/report"
)

// Verdict question IDs the intent checks answer.
const (
	QCallHome   = "call_home"
	QAutorun    = "autorun"
	QExfil      = "exfil"
	QRemoteExec = "remote_exec"
	QDeps       = "deps"
	QCI         = "ci"
	QAgent      = "agent"
	QPrivileges = "privileges"
	QLicense    = "license"
	QSecrets    = "secrets"
	QHidden     = "hidden"
)

// maxFileBytes bounds the files the intent checks read. Lockfiles of large
// web projects reach a few MB, so the limit is generous.
const maxFileBytes = 8 << 20

// File is one repository file as seen by the analyzers.
type File struct {
	Rel     string
	Role    string
	Content string
	// Dep marks source of a downloaded dependency (scan --network).
	Dep       bool
	linesOnce sync.Once
	lines     []string
}

// Lines returns the file content split into lines. Safe for concurrent callers.
func (f *File) Lines() []string {
	f.linesOnce.Do(func() {
		f.lines = strings.Split(f.Content, "\n")
	})
	return f.lines
}

// Base returns the lower-case base name of the file.
func (f *File) Base() string {
	rel := strings.ReplaceAll(f.Rel, "\\", "/")
	return strings.ToLower(rel[strings.LastIndex(rel, "/")+1:])
}

// Analyzer inspects a set of files and returns findings.
type Analyzer struct {
	Name string
	Run  func(files []*File) []report.Finding
}

// Analyzers returns every intent analyzer in a fixed order.
func Analyzers() []Analyzer {
	return []Analyzer{
		{"egress", analyzeEgress},
		{"autorun", analyzeAutorun},
		{"depsource", analyzeDepSource},
		{"installscripts", analyzeInstallScripts},
		{"ci", analyzeCI},
		{"extensions", analyzeExtensions},
		{"chains", analyzeChains},
		{"agentsurface", analyzeAgentSurface},
		{"license", analyzeLicense},
	}
}

// StageNames returns the pipeline stage names of the intent analyzers.
func StageNames() []string {
	var names []string
	for _, a := range Analyzers() {
		names = append(names, a.Name)
	}
	return names
}

// Stage wraps one analyzer as a pipeline stage.
type Stage struct {
	analyzer Analyzer
	excluded []string
}

// NewStage returns the pipeline stage for the named analyzer, or nil.
func NewStage(name string, excludedPaths []string) *Stage {
	for _, a := range Analyzers() {
		if a.Name == name {
			return &Stage{analyzer: a, excluded: excludedPaths}
		}
	}
	return nil
}

// Name implements pipeline.Stage.
func (s *Stage) Name() string { return s.analyzer.Name }

// Dependencies implements pipeline.Stage.
func (s *Stage) Dependencies() []string { return nil }

// Run implements pipeline.Stage.
func (s *Stage) Run(ctx context.Context, target string, _ *report.Report) (*report.Report, error) {
	rep := report.NewReport(target)
	files, err := loadCached(target, s.excluded)
	if err != nil {
		return nil, err
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	rep.Summary.FilesScanned = len(files)
	for _, f := range s.analyzer.Run(files) {
		rep.AddFinding(f)
	}
	rep.Finish()
	return rep, nil
}

const maxCachedTargets = 16

var (
	cacheMu sync.Mutex
	cache   = map[string]*cachedFiles{}
)

type cachedFiles struct {
	once  sync.Once
	files []*File
	err   error
}

// ClearCache drops all cached repository file snapshots.
func ClearCache() {
	cacheMu.Lock()
	cache = map[string]*cachedFiles{}
	cacheMu.Unlock()
}

// loadCached reads the target once for all intent stages of a run.
func loadCached(target string, excluded []string) ([]*File, error) {
	key := target + "\x00" + strings.Join(excluded, "\x00")
	cacheMu.Lock()
	c, ok := cache[key]
	if !ok {
		if len(cache) >= maxCachedTargets {
			cache = map[string]*cachedFiles{}
		}
		c = &cachedFiles{}
		cache[key] = c
	}
	cacheMu.Unlock()
	c.once.Do(func() { c.files, c.err = LoadFiles(target, excluded) })
	return c.files, c.err
}

// LoadFiles reads all text files under root, skipping excluded path parts.
func LoadFiles(root string, excluded []string) ([]*File, error) {
	var files []*File
	err := yara.WalkDirSorted(root, excluded, func(_ string, content []byte, rel string) error {
		if len(content) > maxFileBytes {
			return nil
		}
		files = append(files, NewFile(rel, string(content)))
		return nil
	})
	return files, err
}

// NewFile builds a File and classifies its role.
func NewFile(rel, content string) *File {
	rel = strings.ReplaceAll(rel, "\\", "/")
	return &File{Rel: rel, Role: filerole.Classify(rel), Content: content}
}

// finding builds an intent finding. The tier is lowered one step for tests
// and docs, which do not run as part of the product (they still show up).
func finding(source, rule, severity, tier, question string, f *File, line int, match, msg string, attrs map[string]string) report.Finding {
	if tier == report.TierA && (f.Role == filerole.Test || f.Role == filerole.Docs) {
		tier = report.TierB
		msg += " (in " + f.Role + " file)"
	}
	if len(match) > 200 {
		match = match[:200] + "…"
	}
	return report.Finding{
		RuleID:     rule,
		Title:      rule,
		Severity:   severity,
		Category:   "intent",
		File:       f.Rel,
		Line:       line,
		Column:     1,
		Match:      strings.TrimSpace(match),
		Message:    msg,
		Source:     source,
		Confidence: report.ConfidenceHeuristic,
		Tier:       tier,
		FileRole:   f.Role,
		Question:   question,
		Attributes: attrs,
	}
}

// firstMatchLine returns the 1-based line of the first match of re in f, or 0.
func firstMatchLine(f *File, re *regexp.Regexp) (int, string) {
	for i, l := range f.Lines() {
		if re.MatchString(l) {
			return i + 1, l
		}
	}
	return 0, ""
}

// lineOf returns the 1-based line number of byte offset off in s.
func lineOf(s string, off int) int {
	return strings.Count(s[:off], "\n") + 1
}

// netCommand matches shell or code that talks to the network.
var netCommand = regexp.MustCompile(`(?i)\b(curl|wget|nc|ncat|socat|telnet)\b|https?://|/dev/tcp/|Invoke-WebRequest|Invoke-RestMethod|\biwr\b|\birm\b|Net\.WebClient|urllib|requests\.|fetch\(|http\.Get|net\.Dial`)

// riskyCommand matches commands that download, decode or evaluate code.
var riskyCommand = regexp.MustCompile(`(?i)\b(curl|wget|nc|ncat|socat)\b|/dev/tcp/|Invoke-WebRequest|Invoke-RestMethod|\biwr\b|\birm\b|Net\.WebClient|\beval\b|base64\s+(-d|--decode)|\bnode\s+-e\b|\bpython[0-9.]*\s+-c\b|\bperl\s+-e\b|\bbash\s+-c\b|\bsh\s+-c\b|\|\s*(sudo\s+)?(ba|z|da)?sh\b|powershell|\bpwsh\b|-EncodedCommand|https?://`)
