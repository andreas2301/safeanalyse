package yara

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
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
			Severity:    "critical",
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
			Severity:    "high",
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
			Severity:    "high",
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
			Severity:    "medium",
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
			Severity:    "medium",
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
			Severity:    "high",
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
			Severity:    "critical",
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
