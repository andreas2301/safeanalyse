package intent

import (
	"regexp"
	"strings"

	"github.com/user/safeanalyze/pkg/checks/yara"
)

// multiPattern is an alternation kept as separate regexes, each with its own
// required-literal prefilter. Go's regexp engine is slow on large
// alternations; checking cheap literals first skips most lines without
// changing what matches.
type multiPattern []pattern

type pattern struct {
	re *regexp.Regexp
	pf [][]string
}

func compileMulti(parts ...string) multiPattern {
	m := make(multiPattern, 0, len(parts))
	for _, p := range parts {
		m = append(m, pattern{re: regexp.MustCompile(p), pf: yara.PatternPrefilter(p)})
	}
	return m
}

func (p pattern) possible(folded string) bool {
	for _, clause := range p.pf {
		ok := false
		for _, lit := range clause {
			if strings.Contains(folded, lit) {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

// forText returns the patterns that can match somewhere in text.
func (m multiPattern) forText(text string) multiPattern {
	folded := yara.FoldString(text)
	var out multiPattern
	for _, p := range m {
		if p.possible(folded) && p.re.MatchString(text) {
			out = append(out, p)
		}
	}
	return out
}

// find returns the leftmost match location in line, or nil.
func (m multiPattern) find(line string) []int {
	if len(m) == 0 {
		return nil
	}
	folded := yara.FoldString(line)
	var best []int
	for _, p := range m {
		if !p.possible(folded) {
			continue
		}
		if loc := p.re.FindStringIndex(line); loc != nil && (best == nil || loc[0] < best[0]) {
			best = loc
		}
	}
	return best
}

// MatchString reports whether any pattern matches s.
func (m multiPattern) MatchString(s string) bool { return m.find(s) != nil }

// FindStringIndex returns the leftmost match location in s, or nil.
func (m multiPattern) FindStringIndex(s string) []int { return m.find(s) }
