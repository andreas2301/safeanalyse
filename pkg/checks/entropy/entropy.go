// Package entropy provides high-entropy string detection and a pipeline Stage wrapper.
package entropy

import (
	"context"
	"encoding/base64"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/report"
)

// Detector scans for suspicious encoded/entropy-heavy strings.
type Detector struct {
	minEntropy float64
	minLength  int
	maxLength  int
}

// NewDetector creates an entropy scanner with sensible defaults.
func NewDetector() *Detector {
	return &Detector{
		minEntropy: 4.5,
		minLength:  20,
		maxLength:  4096,
	}
}

// WithThreshold sets custom entropy threshold.
func (d *Detector) WithThreshold(t float64) *Detector {
	d.minEntropy = t
	return d
}

// ScanString analyzes a single string and returns findings.
func (d *Detector) ScanString(s, file string, line, col int) []Finding {
	var findings []Finding

	if b64 := findBase64(s); b64 != "" && len(b64) >= d.minLength {
		ent := shannonEntropy(b64)
		if ent >= d.minEntropy {
			findings = append(findings, Finding{
				File:    file,
				Line:    line,
				Column:  col,
				Type:    "base64_blob",
				Value:   truncate(b64, 60),
				Entropy: ent,
				Length:  len(b64),
				Context: truncate(s, 80),
			})
		}
	}

	if hex := findHex(s); hex != "" && len(hex) >= d.minLength {
		ent := shannonEntropy(hex)
		if ent >= d.minEntropy {
			findings = append(findings, Finding{
				File:    file,
				Line:    line,
				Column:  col,
				Type:    "hex_blob",
				Value:   truncate(hex, 60),
				Entropy: ent,
				Length:  len(hex),
				Context: truncate(s, 80),
			})
		}
	}

	if len(s) >= d.minLength && len(s) <= d.maxLength {
		ent := shannonEntropy(s)
		if ent >= d.minEntropy+0.5 {
			if !isBenignHighEntropy(s) {
				findings = append(findings, Finding{
					File:    file,
					Line:    line,
					Column:  col,
					Type:    "high_entropy",
					Value:   truncate(s, 60),
					Entropy: ent,
					Length:  len(s),
					Context: truncate(s, 80),
				})
			}
		}
	}

	return findings
}

// ScanLine analyzes a line of text for entropy anomalies.
func (d *Detector) ScanLine(line, file string, lineNum int) []Finding {
	var findings []Finding
	for _, tok := range tokenizeStrings(line) {
		findings = append(findings, d.ScanString(tok.value, file, lineNum, tok.start)...)
	}
	return findings
}

// ScanContent analyzes full file content.
func (d *Detector) ScanContent(content, file string) []Finding {
	var findings []Finding
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		findings = append(findings, d.ScanLine(line, file, i+1)...)
	}
	return findings
}

// Finding represents a high-entropy or suspicious string detection.
type Finding struct {
	File    string  `json:"file"`
	Line    int     `json:"line"`
	Column  int     `json:"column"`
	Type    string  `json:"type"`
	Value   string  `json:"value"`
	Entropy float64 `json:"entropy"`
	Length  int     `json:"length"`
	Context string  `json:"context"`
}

type token struct {
	value string
	start int
}

func tokenizeStrings(line string) []token {
	var tokens []token
	inString := false
	var quoteChar rune
	var current strings.Builder
	start := 0
	escape := false

	for i, r := range line {
		if escape {
			current.WriteRune(r)
			escape = false
			continue
		}
		if r == '\\' {
			escape = true
			current.WriteRune(r)
			continue
		}
		if !inString {
			if r == '"' || r == '\'' || r == '`' {
				inString = true
				quoteChar = r
				start = i
				current.Reset()
			}
		} else {
			if r == quoteChar {
				inString = false
				tokens = append(tokens, token{value: current.String(), start: start})
			} else {
				current.WriteRune(r)
			}
		}
	}
	return tokens
}

func shannonEntropy(s string) float64 {
	if len(s) == 0 {
		return 0
	}
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	var entropy float64
	length := float64(len(s))
	for _, count := range freq {
		p := float64(count) / length
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}
	return entropy
}

var base64Regex = regexp.MustCompile(`[A-Za-z0-9+/]{40,}={0,2}`)

func findBase64(s string) string {
	matches := base64Regex.FindAllString(s, -1)
	var longest string
	for _, m := range matches {
		if len(m) > len(longest) {
			if _, err := base64.StdEncoding.DecodeString(padBase64(m)); err == nil {
				longest = m
			}
		}
	}
	return longest
}

func padBase64(s string) string {
	switch len(s) % 4 {
	case 2:
		return s + "=="
	case 3:
		return s + "="
	}
	return s
}

var hexRegex = regexp.MustCompile(`[0-9a-fA-F]{32,}`)

func findHex(s string) string {
	matches := hexRegex.FindAllString(s, -1)
	var longest string
	for _, m := range matches {
		if len(m) > len(longest) {
			longest = m
		}
	}
	return longest
}

func isBenignHighEntropy(s string) bool {
	if matched, _ := regexp.MatchString(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`, s); matched {
		return true
	}
	lower := strings.ToLower(s)
	if strings.Contains(lower, "http") || strings.Contains(lower, "https") || strings.Contains(lower, "www") {
		return true
	}
	if allSameChar(s) {
		return true
	}
	return false
}

func allSameChar(s string) bool {
	if len(s) == 0 {
		return true
	}
	first := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != first {
			return false
		}
	}
	return true
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}

// Summary returns a human-readable summary of findings.
func Summary(findings []Finding) string {
	counts := make(map[string]int)
	for _, f := range findings {
		counts[f.Type]++
	}
	var parts []string
	for t, c := range counts {
		parts = append(parts, fmt.Sprintf("%s: %d", t, c))
	}
	return strings.Join(parts, ", ")
}

// Stage wraps the entropy detector as a pipeline stage.
type Stage struct {
	threshold        float64
	maxFileSizeBytes int
	excludedPaths    []string
}

// NewStage creates an entropy pipeline stage.
func NewStage(threshold float64, maxFileSizeBytes int, excludedPaths []string) *Stage {
	return &Stage{threshold: threshold, maxFileSizeBytes: maxFileSizeBytes, excludedPaths: excludedPaths}
}

// Name returns the stage name.
func (s *Stage) Name() string { return "entropy" }

// Dependencies returns no dependencies.
func (s *Stage) Dependencies() []string { return nil }

// Run scans all files under target and returns findings.
func (s *Stage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := report.NewReport(target)
	if input != nil {
		out = input
	}
	det := NewDetector().WithThreshold(s.threshold)

	err := yara.WalkDirSorted(target, s.excludedPaths, func(path string, content []byte, rel string) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if s.maxFileSizeBytes > 0 && len(content) > s.maxFileSizeBytes {
			out.AddError(s.Name(), fmt.Sprintf("skipped %s: size %d exceeds entropy max %d", rel, len(content), s.maxFileSizeBytes))
			return nil
		}
		findings := det.ScanContent(string(content), rel)
		for _, f := range findings {
			sev := report.SeverityLow
			if f.Type == "base64_blob" || f.Type == "hex_blob" {
				sev = report.SeverityMedium
			}
			out.AddFinding(report.Finding{
				RuleID:     "entropy_" + f.Type,
				Title:      "High-entropy " + f.Type,
				Severity:   sev,
				Category:   "entropy",
				File:       f.File,
				Line:       f.Line,
				Column:     f.Column,
				Match:      f.Value,
				Message:    fmt.Sprintf("entropy=%.2f len=%d type=%s", f.Entropy, f.Length, f.Type),
				Source:     s.Name(),
				Confidence: report.ConfidenceHeuristic,
			})
		}
		out.Summary.FilesScanned++
		return nil
	})
	return out, err
}
