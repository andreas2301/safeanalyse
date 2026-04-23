package aststrip

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// StripGoComments uses the Go AST to precisely remove comments while preserving
// string literals that happen to contain comment-like text.
func StripGoComments(src []byte) []byte {
	fset := token.NewFileSet()
	var file *ast.File
	file, err := parser.ParseFile(fset, "input.go", src, parser.ParseComments)
	f := file
	if err != nil {
		// Fallback: return source unchanged if parse fails
		return src
	}

	// Collect comment positions
	var comments []struct{ start, end int }
	for _, cg := range f.Comments {
		for _, c := range cg.List {
			pos := fset.Position(c.Pos())
			end := fset.Position(c.End())
			comments = append(comments, struct{ start, end int }{
				start: pos.Offset,
				end:   end.Offset,
			})
		}
	}

	if len(comments) == 0 {
		return src
	}

	// Reconstruct source without comments
	var out strings.Builder
	last := 0
	for _, c := range comments {
		if c.start > last {
			out.Write(src[last:c.start])
		}
		last = c.end
	}
	if last < len(src) {
		out.Write(src[last:])
	}

	result := []byte(out.String())
	// Clean up extra blank lines left by comment removal
	return normalizeBlankLines(result)
}

// normalizeBlankLines collapses 3+ consecutive newlines to 2.
func normalizeBlankLines(src []byte) []byte {
	var out strings.Builder
	consecutiveNewlines := 0
	for _, b := range src {
		if b == '\n' {
			consecutiveNewlines++
			if consecutiveNewlines <= 2 {
				out.WriteByte(b)
			}
		} else {
			consecutiveNewlines = 0
			out.WriteByte(b)
		}
	}
	return []byte(out.String())
}

// CanASTStrip returns true if we have an AST-based stripper for this extension.
func CanASTStrip(ext string) bool {
	switch strings.ToLower(ext) {
	case ".go":
		return true
	default:
		return false
	}
}

// Strip dispatches to the appropriate AST stripper or falls back to regex.
func Strip(src []byte, ext string) []byte {
	switch strings.ToLower(ext) {
	case ".go":
		return StripGoComments(src)
	default:
		return src // caller should use regex fallback
	}
}
