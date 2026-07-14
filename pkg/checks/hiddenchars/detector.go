// Package hiddenchars provides detection of invisible/suspicious Unicode characters
// and a pipeline Stage wrapper.
package hiddenchars

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/report"
)

// Detector scans files for invisible/suspicious Unicode characters.
type Detector struct {
	categories map[string]bool
}

// NewDetector creates a detector filtering by category names.
func NewDetector(categories []string) *Detector {
	d := &Detector{categories: make(map[string]bool)}
	for _, c := range categories {
		d.categories[strings.ToLower(c)] = true
	}
	if d.categories["all"] {
		d.categories["zero_width"] = true
		d.categories["bidi"] = true
		d.categories["control"] = true
		d.categories["whitespace"] = true
		d.categories["homoglyph"] = true
	}
	return d
}

// IsSuspicious determines if a rune matches the configured categories.
func (d *Detector) IsSuspicious(r rune) (category, name string, ok bool) {
	if d.categories["zero_width"] {
		switch r {
		case '\u200B':
			return "zero_width", "ZERO WIDTH SPACE", true
		case '\u200C':
			return "zero_width", "ZERO WIDTH NON-JOINER", true
		case '\u200D':
			return "zero_width", "ZERO WIDTH JOINER", true
		case '\uFEFF':
			return "zero_width", "BYTE ORDER MARK", true
		case '\u2060':
			return "zero_width", "WORD JOINER", true
		case '\u180E':
			return "zero_width", "MONGOLIAN VOWEL SEPARATOR", true
		}
	}

	if d.categories["bidi"] {
		switch r {
		case '\u202A':
			return "bidi", "LEFT-TO-RIGHT EMBEDDING", true
		case '\u202B':
			return "bidi", "RIGHT-TO-LEFT EMBEDDING", true
		case '\u202C':
			return "bidi", "POP DIRECTIONAL FORMATTING", true
		case '\u202D':
			return "bidi", "LEFT-TO-RIGHT OVERRIDE", true
		case '\u202E':
			return "bidi", "RIGHT-TO-LEFT OVERRIDE", true
		case '\u2066':
			return "bidi", "LEFT-TO-RIGHT ISOLATE", true
		case '\u2067':
			return "bidi", "RIGHT-TO-LEFT ISOLATE", true
		case '\u2068':
			return "bidi", "FIRST STRONG ISOLATE", true
		case '\u2069':
			return "bidi", "POP DIRECTIONAL ISOLATE", true
		}
	}

	if d.categories["control"] {
		if unicode.IsControl(r) && r != '\t' && r != '\n' && r != '\r' {
			name := fmt.Sprintf("CONTROL (U+%04X)", r)
			if r < 0x20 {
				name = fmt.Sprintf("C0 CONTROL (U+%04X)", r)
			} else if r >= 0x7F && r <= 0x9F {
				name = fmt.Sprintf("C1 CONTROL (U+%04X)", r)
			}
			return "control", name, true
		}
	}

	if d.categories["whitespace"] {
		switch r {
		case '\u00A0', '\u1680', '\u2000', '\u2001', '\u2002', '\u2003',
			'\u2004', '\u2005', '\u2006', '\u2007', '\u2008', '\u2009',
			'\u200A', '\u202F', '\u205F', '\u3000':
			return "whitespace", fmt.Sprintf("UNUSUAL WHITESPACE (U+%04X)", r), true
		}
	}

	if d.categories["zero_width"] || d.categories["control"] {
		if unicode.Is(unicode.Cf, r) {
			return "format", fmt.Sprintf("FORMAT CHAR (U+%04X)", r), true
		}
	}

	return "", "", false
}

// ScanFile scans a single file for suspicious characters.
func (d *Detector) ScanFile(path string) ([]Finding, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var findings []Finding
	lines := strings.Split(string(data), "\n")

	for lineIdx, line := range lines {
		col := 0
		for _, r := range line {
			if cat, name, ok := d.IsSuspicious(r); ok {
				findings = append(findings, Finding{
					File:      path,
					Line:      lineIdx + 1,
					Column:    col + 1,
					Rune:      r,
					CodePoint: fmt.Sprintf("U+%04X", r),
					Category:  cat,
					Name:      name,
					Context:   truncateContext(line, col),
				})
			}
			col += utf8.RuneLen(r)
		}
	}

	return findings, nil
}

// Finding represents a detected suspicious character.
type Finding struct {
	File      string `json:"file"`
	Line      int    `json:"line"`
	Column    int    `json:"column"`
	Rune      rune   `json:"rune"`
	CodePoint string `json:"code_point"`
	Category  string `json:"category"`
	Name      string `json:"name"`
	Context   string `json:"context"`
}

func truncateContext(line string, center int) string {
	start := center - 20
	if start < 0 {
		start = 0
	}
	end := center + 20
	if end > len(line) {
		end = len(line)
	}
	ctx := line[start:end]
	ctx = strings.ReplaceAll(ctx, "\t", " ")
	ctx = strings.ReplaceAll(ctx, "\n", " ")
	return ctx
}

// Stage wraps the hidden char detector as a pipeline stage.
type Stage struct {
	categories    []string
	excludedPaths []string
}

// NewStage creates a hidden-char pipeline stage.
func NewStage(categories, excludedPaths []string) *Stage {
	return &Stage{categories: categories, excludedPaths: excludedPaths}
}

// Name returns the stage name.
func (s *Stage) Name() string { return "hiddenchars" }

// Dependencies returns no dependencies.
func (s *Stage) Dependencies() []string { return nil }

// Run scans all files under target and returns findings.
func (s *Stage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := report.NewReport(target)
	if input != nil {
		out = input
	}
	det := NewDetector(s.categories)

	err := yara.WalkDirSorted(target, s.excludedPaths, func(path string, content []byte, rel string) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		lines := strings.Split(string(content), "\n")
		for lineIdx, line := range lines {
			col := 0
			for _, r := range line {
				if cat, name, ok := det.IsSuspicious(r); ok {
					sev := report.SeverityHigh
					if cat == "whitespace" {
						sev = report.SeverityLow
					}
					out.AddFinding(report.Finding{
						RuleID:     "hidden_char_" + cat,
						Title:      "Hidden Unicode character: " + name,
						Severity:   sev,
						Category:   "hidden_char",
						File:       rel,
						Line:       lineIdx + 1,
						Column:     col + 1,
						Match:      fmt.Sprintf("U+%04X", r),
						Message:    fmt.Sprintf("%s (%s) in %s", name, cat, rel),
						Source:     s.Name(),
						Confidence: report.ConfidenceDeterministic,
					})
				}
				col += utf8.RuneLen(r)
			}
		}
		out.Summary.FilesScanned++
		return nil
	})
	return out, err
}

// walkDir is retained for non-pipeline consumers.
func walkDir(root string, exclude []string, fn func(string) error) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			for _, ex := range exclude {
				if strings.Contains(path, ex) {
					return filepath.SkipDir
				}
			}
			return nil
		}
		return fn(path)
	})
}

// ScanDir recursively scans a directory for suspicious characters.
func (d *Detector) ScanDir(root string, exclude []string) ([]Finding, error) {
	var all []Finding
	err := walkDir(root, exclude, func(path string) error {
		f, err := d.ScanFile(path)
		if err != nil {
			return err
		}
		all = append(all, f...)
		return nil
	})
	return all, err
}
