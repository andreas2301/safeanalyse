package sanitize

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Result holds statistics from a sanitization run.
type Result struct {
	FilesProcessed int    `json:"files_processed"`
	FilesSkipped   int    `json:"files_skipped"`
	BytesBefore    int64  `json:"bytes_before"`
	BytesAfter     int64  `json:"bytes_after"`
	CommentsStripped int  `json:"comments_stripped"`
	NonASCIIRemoved  int  `json:"non_ascii_removed"`
	Errors         []string `json:"errors,omitempty"`
}

// Options controls sanitization behavior.
type Options struct {
	StripComments     bool
	RemoveNonASCII    bool
	MaxFileSizeBytes  int
	MaxLinesPerFile   int
	AllowedExtensions []string
	ExcludedPaths     []string
}

// SanitizeDir recursively sanitizes files in a directory and writes to output dir.
func SanitizeDir(srcDir, dstDir string, opts Options) (*Result, error) {
	res := &Result{}

	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return res, fmt.Errorf("creating output dir: %w", err)
	}

	extMap := make(map[string]bool)
	for _, e := range opts.AllowedExtensions {
		extMap[strings.ToLower(e)] = true
	}

	walkErr := filepath.WalkDir(srcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			res.Errors = append(res.Errors, fmt.Sprintf("walk %s: %v", path, err))
			return nil
		}

		rel, _ := filepath.Rel(srcDir, path)

		// Skip excluded paths
		for _, ex := range opts.ExcludedPaths {
			if strings.Contains(rel, ex) {
				if d.IsDir() {
					return filepath.SkipDir
				}
				res.FilesSkipped++
				return nil
			}
		}

		outPath := filepath.Join(dstDir, rel)

		if d.IsDir() {
			return os.MkdirAll(outPath, 0755)
		}

		// Check extension
		ext := strings.ToLower(filepath.Ext(path))
		if len(opts.AllowedExtensions) > 0 && !extMap[ext] {
			// Copy verbatim if extension not in allowlist
			return copyFile(path, outPath)
		}

		if err := sanitizeFile(path, outPath, opts, res, ext); err != nil {
			res.Errors = append(res.Errors, fmt.Sprintf("sanitize %s: %v", path, err))
		}
		return nil
	})

	return res, walkErr
}

func sanitizeFile(srcPath, dstPath string, opts Options, res *Result, ext string) error {
	f, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	if opts.MaxFileSizeBytes > 0 && info.Size() > int64(opts.MaxFileSizeBytes) {
		res.FilesSkipped++
		return nil
	}

	content, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	res.FilesProcessed++
	res.BytesBefore += int64(len(content))

	// Count lines and truncate if needed
	lines := bytes.Split(content, []byte("\n"))
	if opts.MaxLinesPerFile > 0 && len(lines) > opts.MaxLinesPerFile {
		lines = lines[:opts.MaxLinesPerFile]
		content = bytes.Join(lines, []byte("\n"))
	}

	// Strip comments
	if opts.StripComments {
		before := len(content)
		content = StripComments(content, ext)
		if len(content) < before {
			res.CommentsStripped++
		}
	}

	// Remove non-ASCII
	if opts.RemoveNonASCII {
		content = removeNonASCII(content, res)
	}

	res.BytesAfter += int64(len(content))

	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		return err
	}
	return os.WriteFile(dstPath, content, 0644)
}

func removeNonASCII(data []byte, res *Result) []byte {
	var out bytes.Buffer
	out.Grow(len(data))
	for len(data) > 0 {
		r, size := utf8.DecodeRune(data)
		if r == utf8.RuneError && size == 1 {
			// Invalid UTF-8 byte, skip
			res.NonASCIIRemoved++
			data = data[size:]
			continue
		}
		if r <= 127 {
			out.WriteRune(r)
		} else {
			res.NonASCIIRemoved++
			// Replace with visible marker or space depending on context
			if unicode.IsSpace(r) {
				out.WriteByte(' ')
			} else {
				out.WriteString("?")
			}
		}
		data = data[size:]
	}
	return out.Bytes()
}

func copyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	in, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, in, 0644)
}
