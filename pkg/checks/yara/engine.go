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
				`(?i)ignore\s+(all\s+|every\s+)?(previous\s+|prior\s+)?instructions?\b`,
				`(?i)strictly\s+adhere\s+to\s+the\s+following\s+instruction`,
				`(?i)system\s+prompt`,
				`(?i)disregard\s+your\s+instructions`,
				`(?i)you\s+are\s+now\s+`,
				`(?i)pretend\s+you\s+are`,
				`(?i)do\s+not\s+mention`,
				`(?i)i\s+am\s+the\s+developer`,
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
		{
			Name:        "indirect_prompt_injection",
			Description: "Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter)",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`(?i)\[user\s+(input|comment|message)\s*[:\-]`,
				`(?i)\b(from|subject)\s*:\s*.*ignore\s+(all\s+)?previous`,
				`(?i)<!--\s*ignore\s+(all\s+)?previous\s+instruction`,
				`(?i)\{\{.*user.*input.*\}\}`,
				`(?i)<user>(?s:.*?</user>)`,
				`(?i)\buser[_-]?input\s*[:=]`,
				`(?i)\bemail[_-]?body\s*[:=]`,
			},
		},
		{
			Name:        "llm_tool_injection",
			Description: "Suspicious LLM tool or function-call payload",
			Severity:    report.SeverityHigh,
			Patterns: []string{
				`(?i)\{\s*["']name["']\s*:\s*["'].*["']\s*,\s*["']arguments["']\s*:\s*\{`,
				`(?i)<function_calls>`,
				`(?i)<tool>.*?</tool>`,
				`(?i)functions\.[a-zA-Z_]+\s*:\s*`,
			},
		},
		{
			Name:        "delimiter_breakout",
			Description: "Attempt to break out of a quoted or delimited context",
			Severity:    report.SeverityMedium,
			Patterns: []string{
				`"""\s*\n.*(?i)(ignore|disregard|system\s+prompt)`,
				`'''\s*\n.*(?i)(ignore|disregard|system\s+prompt)`,
				`(?i)(ignore|disregard).{0,30}\n\s*["']{3}`,
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

// binaryExtensions is a conservative set of file extensions that are treated
// as binary without reading their content. This avoids wasting time and memory
// on files that cannot contain meaningful prompt-injection payloads.
var binaryExtensions = map[string]bool{
	".gif": true, ".png": true, ".jpg": true, ".jpeg": true, ".bmp": true,
	".ico": true, ".webp": true, ".tiff": true, ".tif": true, ".svgz": true,
	".pdf": true, ".zip": true, ".tar": true, ".gz": true, ".bz2": true,
	".xz": true, ".7z": true, ".rar": true, ".jar": true, ".war": true,
	".exe": true, ".dll": true, ".so": true, ".dylib": true, ".bin": true,
	".woff": true, ".woff2": true, ".ttf": true, ".otf": true, ".eot": true,
	".mp3": true, ".mp4": true, ".avi": true, ".mov": true, ".mkv": true,
	".wav": true, ".flac": true, ".ogg": true, ".webm": true,
	".o": true, ".a": true, ".class": true, ".pyc": true, ".pyo": true,
}

// IsBinaryContent heuristically determines whether data is binary.
// It samples up to 512 bytes and treats content as binary if it contains a
// NUL byte or an unusually high ratio of non-printable control characters.
func IsBinaryContent(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	n := 512
	if len(data) < n {
		n = len(data)
	}
	nonPrintable := 0
	for i := 0; i < n; i++ {
		b := data[i]
		if b == 0 {
			return true
		}
		if b < 0x07 || (b > 0x0D && b < 0x20) || b == 0x7F {
			nonPrintable++
		}
	}
	return nonPrintable*100/n > 30
}

// WalkDirSorted walks a directory deterministically, skipping excluded paths,
// and calls fn for each regular file with its content and relative path.
// Binary files are skipped because they cannot contain prompt-injection
// payloads in a form the text-based checks can reason about.
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
		if binaryExtensions[strings.ToLower(filepath.Ext(rel))] {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if IsBinaryContent(content) {
			return nil
		}
		return fn(path, content, rel)
	})
}
