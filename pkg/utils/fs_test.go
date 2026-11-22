package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyDirSkipsSymlinks(t *testing.T) {
	src := t.TempDir()
	dst := t.TempDir()
	outside := filepath.Join(t.TempDir(), "outside.txt")
	if err := os.WriteFile(outside, []byte("secret"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(src, "ok.txt"), []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(src, "link.txt")); err != nil {
		t.Skipf("symlink unsupported: %v", err)
	}

	if err := CopyDir(src, dst); err != nil {
		t.Fatalf("CopyDir failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dst, "ok.txt")); err != nil {
		t.Errorf("expected ok.txt copied: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dst, "link.txt")); !os.IsNotExist(err) {
		t.Errorf("expected symlink link.txt to be skipped")
	}
}

func TestFormatBytes(t *testing.T) {
	if got := FormatBytes(512); got != "512 B" {
		t.Errorf("got %q, want 512 B", got)
	}
	if got := FormatBytes(2048); got != "2.0 KB" {
		t.Errorf("got %q, want 2.0 KB", got)
	}
}
