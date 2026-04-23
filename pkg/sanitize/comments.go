package sanitize

import (
	"bytes"
	"strings"

	"github.com/user/safeanalyze/pkg/aststrip"
)

// CommentStyle defines how comments look in a language.
type CommentStyle struct {
	Single     string   // e.g., "//"
	MultiStart string   // e.g., "/*"
	MultiEnd   string   // e.g., "*/"
	Nested     bool     // true if multi-line comments can nest
}

// LanguageComments maps file extensions to their comment styles.
var LanguageComments = map[string]CommentStyle{
	".go":    {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".js":    {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".ts":    {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".jsx":   {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".tsx":   {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".java":  {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".c":     {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".cpp":   {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".h":     {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".hpp":   {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".rs":    {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".swift": {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".kt":    {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".scala": {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".php":   {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".cs":    {Single: "//", MultiStart: "/*", MultiEnd: "*/"},
	".py":    {Single: "#", MultiStart: `"""`, MultiEnd: `"""`},
	".rb":    {Single: "#", MultiStart: "=begin", MultiEnd: "=end"},
	".sh":    {Single: "#", MultiStart: "", MultiEnd: ""},
	".bash":  {Single: "#", MultiStart: "", MultiEnd: ""},
	".zsh":   {Single: "#", MultiStart: "", MultiEnd: ""},
	".yaml":  {Single: "#", MultiStart: "", MultiEnd: ""},
	".yml":   {Single: "#", MultiStart: "", MultiEnd: ""},
	".json":  {Single: "", MultiStart: "", MultiEnd: ""}, // no comments in strict JSON
	".toml":  {Single: "#", MultiStart: "", MultiEnd: ""},
	".sql":   {Single: "--", MultiStart: "/*", MultiEnd: "*/"},
	".html":  {Single: "", MultiStart: "<!--", MultiEnd: "-->"},
	".xml":   {Single: "", MultiStart: "<!--", MultiEnd: "-->"},
	".css":   {Single: "", MultiStart: "/*", MultiEnd: "*/"},
	".lua":   {Single: "--", MultiStart: "--[[", MultiEnd: "]]"},
	".r":     {Single: "#", MultiStart: "", MultiEnd: ""},
}

// StripComments removes comments from source code for a given extension.
// It attempts AST-based stripping first (for supported languages), then falls
// back to regex-based stripping.
func StripComments(content []byte, ext string) []byte {
	// Try AST-based stripping first (bulletproof for supported languages)
	if aststrip.CanASTStrip(ext) {
		return aststrip.Strip(content, ext)
	}

	// Fallback to regex-based stripping
	style, ok := LanguageComments[strings.ToLower(ext)]
	if !ok || (style.Single == "" && style.MultiStart == "") {
		return content // unknown language, pass through
	}

	if style.MultiStart == "" {
		return stripSingleLineComments(content, style.Single)
	}

	return stripMultiLineComments(content, style.Single, style.MultiStart, style.MultiEnd)
}

func stripSingleLineComments(content []byte, marker string) []byte {
	lines := bytes.Split(content, []byte("\n"))
	markerBytes := []byte(marker)
	var out bytes.Buffer

	for i, line := range lines {
		if idx := bytes.Index(line, markerBytes); idx >= 0 {
			// Check if marker is inside a string (naive check)
			if !isInsideString(line[:idx]) {
				line = bytes.TrimRight(line[:idx], " \t")
			}
		}
		if len(line) > 0 {
			out.Write(line)
		}
		if i < len(lines)-1 {
			out.WriteByte('\n')
		}
	}
	return out.Bytes()
}

func stripMultiLineComments(content []byte, single, multiStart, multiEnd string) []byte {
	var out bytes.Buffer
	inMulti := false
	lines := bytes.Split(content, []byte("\n"))
	ms := []byte(multiStart)
	me := []byte(multiEnd)
	singleB := []byte(single)

	for _, line := range lines {
		if inMulti {
			if idx := bytes.Index(line, me); idx >= 0 {
				line = line[idx+len(me):]
				inMulti = false
				// Fall through to process rest of line
			} else {
				continue // entire line is inside multi-line comment
			}
		}

		// Process the line (now known not to be fully inside multi-line comment)
		var processed bytes.Buffer
		i := 0
		for i < len(line) {
			// Check for multi-line comment start
			if !inMulti && i+len(ms) <= len(line) && bytes.Equal(line[i:i+len(ms)], ms) {
				inMulti = true
				i += len(ms)
				// Check if it ends on same line
				if idx := bytes.Index(line[i:], me); idx >= 0 {
					i += idx + len(me)
					inMulti = false
				} else {
					break
				}
				continue
			}

			// Check for single-line comment
			if !inMulti && i+len(singleB) <= len(line) && bytes.Equal(line[i:i+len(singleB)], singleB) {
				if !isInsideString(line[:i]) {
					break
				}
			}

			processed.WriteByte(line[i])
			i++
		}

		trimmed := bytes.TrimRight(processed.Bytes(), " \t")
		if len(trimmed) > 0 {
			out.Write(trimmed)
		}
		out.WriteByte('\n')
	}

	result := out.Bytes()
	if len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1]
	}
	return result
}

// isInsideString is a naive heuristic: odd number of unescaped quotes before position.
func isInsideString(prefix []byte) bool {
	inString := false
	escape := false
	for _, b := range prefix {
		if escape {
			escape = false
			continue
		}
		if b == '\\' {
			escape = true
			continue
		}
		if b == '"' || b == '\'' || b == '`' {
			inString = !inString
		}
	}
	return inString
}
