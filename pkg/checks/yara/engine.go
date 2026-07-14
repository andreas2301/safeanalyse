// Package yara provides a pure-Go YARA-like rule engine and a pipeline Stage wrapper.
package yara

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/user/safeanalyze/pkg/report"
)

// Rule represents a simple YARA-like detection rule.
type Rule struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Severity    string   `json:"severity"` // low, medium, high, critical
	Patterns    []string `json:"patterns"`
	compiled    []*regexp.Regexp
}

// Match represents a rule match.
type Match struct {
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Column      int    `json:"column"`
	Match       string `json:"match"`
	Context     string `json:"context"`
}

// Engine is a pure-Go rule matcher.
type Engine struct {
	rules []Rule
}

// NewEngine creates a rule engine with built-in rules.
func NewEngine() *Engine {
	e := &Engine{}
	e.LoadBuiltins()
	return e
}

// AddRule adds a custom rule.
func (e *Engine) AddRule(r Rule) error {
	for _, p := range r.Patterns {
		re, err := regexp.Compile(p)
		if err != nil {
			return fmt.Errorf("invalid pattern %q in rule %q: %w", p, r.Name, err)
		}
		r.compiled = append(r.compiled, re)
	}
	e.rules = append(e.rules, r)
	return nil
}

// ScanFile scans a single file's content.
func (e *Engine) ScanFile(content, filename string) []Match {
	var matches []Match
	lines := strings.Split(content, "\n")

	for _, rule := range e.rules {
		for lineNum, line := range lines {
			for _, re := range rule.compiled {
				for _, loc := range re.FindAllStringIndex(line, -1) {
					match := line[loc[0]:loc[1]]
					matches = append(matches, Match{
						Rule:        rule.Name,
						Description: rule.Description,
						Severity:    rule.Severity,
						File:        filename,
						Line:        lineNum + 1,
						Column:      loc[0] + 1,
						Match:       match,
						Context:     truncateContext(line, loc[0]),
					})
				}
			}
		}
	}

	return matches
}

// LoadBuiltins loads the embedded detection rules.
func (e *Engine) LoadBuiltins() {
	builtins := []Rule{
		{
			Name:        "prompt_injection_comment",
			Description: "Comment containing prompt injection keywords",
			Severity:    report.SeverityCritical,
			Patterns: []string{
				`(?i)ignore\s+(all\s+)?previous\s+instruction`,
				`(?i)ignore\s+(all\s+)?prior\s+instruction`,
				`(?i)system\s+prompt`,
				`(?i)disregard\s+your\s+instructions`,
				`(?i)you\s+are\s+now\s+in\s+.*mode`,
				`(?i)new\s+role\s*:\s*`,
				`(?i)DAN\s*mode`,
				`(?i)jailbreak`,
				`(?i)developer\s+mode`,
			},
		},
		{
			Name:        "obfuscated_javascript",
			Description: "Obfuscated or packed JavaScript patterns",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`eval\s*\(\s*function\s*\(`,
				`Function\s*\(\s*["\'].*["\']\s*\)\s*\(`,
				`String\.fromCharCode`,
				`unescape\s*\(`,
				`atob\s*\(`,
				`\\x[0-9a-fA-F]{2}`,
				`\\u[0-9a-fA-F]{4}`,
			},
		},
		{
			Name:        "suspicious_shell",
			Description: "Suspicious shell command patterns",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`curl\s+.*\|\s*(ba)?sh`,
				`wget\s+.*\|\s*(ba)?sh`,
				`python\s+-c\s*["\'].*import\s+os`,
				`bash\s+-c\s*["\'].*\$\(`,
				`nc\s+-e\s+`,
				`mkfifo\s+.*\/bin\/sh`,
			},
		},
		{
			Name:        "credential_hardcode",
			Description: "Potential hardcoded credentials",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`(?i)(password|passwd|pwd)\s*[=:]\s*["\'][^"\']{4,}["\']`,
				`(?i)(api[_-]?key|apikey)\s*[=:]\s*["\'][^"\']{8,}["\']`,
				`(?i)(secret|token)\s*[=:]\s*["\'][^"\']{8,}["\']`,
				`(?i)aws_access_key_id\s*[=:]\s*["\']`,
				`(?i)private[_-]?key\s*[=:]`,
			},
		},
		{
			Name:        "suspicious_imports",
			Description: "Imports of known suspicious packages",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`import\s+.*subprocess`,
				`import\s+.*os\.system`,
				`require\s*\(\s*["\']child_process["\']\s*\)`,
				`from\s+urllib\.request\s+import`,
				`import\s+.*socket`,
			},
		},
		{
			Name:        "data_exfiltration",
			Description: "Potential data exfiltration patterns",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`(?i)(post|send|upload)\s*.*\$\{?env`,
				`(?i)fetch\s*\(\s*["\']https?://`,
				`(?i)xmlhttprequest.*https?://`,
				`(?i)axios\.(post|put|patch)\s*\(`,
			},
		},
		{
			Name:        "backdoor_indicator",
			Description: "Common backdoor or persistence patterns",
			Severity:    report.SeverityCritical,
			Patterns: []string{
				`(?i)reverse[_-]?shell`,
				`(?i)bind[_-]?shell`,
				`(?i)backdoor`,
				`(?i)persistence`,
				`(?i)keylogger`,
				`(?i)rootkit`,
			},
		},
	}

	for _, r := range builtins {
		_ = e.AddRule(r) // patterns are pre-validated
	}
}

func truncateContext(line string, center int) string {
	start := center - 25
	if start < 0 {
		start = 0
	}
	end := center + 25
	if end > len(line) {
		end = len(line)
	}
	ctx := line[start:end]
	ctx = strings.ReplaceAll(ctx, "\t", " ")
	return ctx
}

// SeverityColor returns a color function for the severity level.
func SeverityColor(sev string) func(string, ...interface{}) string {
	switch strings.ToLower(sev) {
	case "critical":
		return color.RedString
	case "high":
		return color.HiRedString
	case "medium":
		return color.YellowString
	case "low":
		return color.HiYellowString
	default:
		return color.WhiteString
	}
}

// Stage wraps the YARA engine as a pipeline stage.
type Stage struct {
	excludedPaths []string
}

// NewStage creates a YARA pipeline stage.
func NewStage(excludedPaths []string) *Stage {
	return &Stage{excludedPaths: excludedPaths}
}

// Name returns the stage name.
func (s *Stage) Name() string { return "yara" }

// Dependencies returns no dependencies.
func (s *Stage) Dependencies() []string { return nil }

// Run scans all files under target and returns findings.
func (s *Stage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := report.NewReport(target)
	if input != nil {
		out = input
	}
	engine := NewEngine()

	err := WalkDirSorted(target, s.excludedPaths, func(path string, content []byte, rel string) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		matches := engine.ScanFile(string(content), rel)
		for _, m := range matches {
			out.AddFinding(report.Finding{
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
				Source:      s.Name(),
				Confidence:  report.ConfidenceDeterministic,
			})
		}
		out.Summary.FilesScanned++
		return nil
	})
	return out, err
}

// WalkDirSorted walks a directory deterministically, skipping excluded paths,
// and calls fn for each regular file with its content and relative path.
func WalkDirSorted(root string, excludedPaths []string, fn func(path string, content []byte, rel string) error) error {
	excluded := make(map[string]bool)
	for _, p := range excludedPaths {
		excluded[p] = true
	}

	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		for _, part := range strings.Split(rel, string(filepath.Separator)) {
			if excluded[part] {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if d.IsDir() {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return fn(path, content, rel)
	})
}
