package entropy

import (
	"encoding/base64"
	"fmt"
	"math"
	"regexp"
	"strings"
)

// Finding represents a high-entropy or suspicious string detection.
type Finding struct {
	File       string  `json:"file"`
	Line       int     `json:"line"`
	Column     int     `json:"column"`
	Type       string  `json:"type"`       // "high_entropy", "base64_blob", "hex_blob"
	Value      string  `json:"value"`
	Entropy    float64 `json:"entropy"`
	Length     int     `json:"length"`
	Context    string  `json:"context"`
}

// Detector scans for suspicious encoded/entropy-heavy strings.
type Detector struct {
	minEntropy   float64
	minLength    int
	maxLength    int
}

// NewDetector creates an entropy scanner with sensible defaults.
func NewDetector() *Detector {
	return &Detector{
		minEntropy: 4.5,  // Shannon entropy threshold (0-8 for bytes)
		minLength:  20,   // Minimum string length to analyze
		maxLength:  4096, // Maximum string length
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

	// Check for base64 blobs
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

	// Check for hex blobs
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

	// Check for high-entropy strings that aren't obviously base64/hex
	if len(s) >= d.minLength && len(s) <= d.maxLength {
		ent := shannonEntropy(s)
		if ent >= d.minEntropy+0.5 { // Higher threshold for general strings
			// Avoid flagging obvious non-secrets
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

	// Tokenize by quotes
	tokens := tokenizeStrings(line)
	for _, tok := range tokens {
		f := d.ScanString(tok.value, file, lineNum, tok.start)
		findings = append(findings, f...)
	}

	return findings
}

// ScanContent analyzes full file content.
func (d *Detector) ScanContent(content, file string) []Finding {
	var findings []Finding
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		f := d.ScanLine(line, file, i+1)
		findings = append(findings, f...)
	}
	return findings
}

type token struct {
	value string
	start int
}

// tokenizeStrings extracts quoted strings from a line.
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

// shannonEntropy calculates Shannon entropy of a string in bits per byte.
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
			// Validate it's actually decodable base64
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

// isBenignHighEntropy filters out obvious false positives.
func isBenignHighEntropy(s string) bool {
	// UUIDs
	if matched, _ := regexp.MatchString(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`, s); matched {
		return true
	}
	// Common variable names or paths
	lower := strings.ToLower(s)
	if strings.Contains(lower, "http") || strings.Contains(lower, "https") || strings.Contains(lower, "www") {
		return true
	}
	// Simple repetition patterns
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
