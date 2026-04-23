package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CopyDir recursively copies a directory tree.
func CopyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		outPath := filepath.Join(dst, rel)

		if d.IsDir() {
			return os.MkdirAll(outPath, 0755)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(outPath, data, 0644)
	})
}

// TempDir creates a temporary directory with a given prefix.
func TempDir(prefix string) (string, error) {
	return os.MkdirTemp("", prefix+"-*")
}

// FileSize returns the size of a file in bytes.
func FileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// FormatBytes returns a human-readable byte string.
func FormatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
