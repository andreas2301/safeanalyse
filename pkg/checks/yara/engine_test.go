package yara

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsBinaryContent(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected bool
	}{
		{"empty", []byte{}, false},
		{"plain ascii", []byte("hello world"), false},
		{"unicode text", []byte("prompt injection: ignore previous instructions"), false},
		{"null byte", []byte{0x00, 0x01, 0x02}, true},
		{"gif header", []byte("GIF89a\x00\x01\x00\x01\x80\x00\x00"), true},
		{"pdf header", []byte("%PDF-1.4\n1 0 obj\n<<\n/Type /Catalog"), false}, // sample too short for ratio
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBinaryContent(tt.data)
			if got != tt.expected {
				t.Errorf("IsBinaryContent(%q) = %v, want %v", tt.data, got, tt.expected)
			}
		})
	}
}

func TestWalkDirSortedSkipsBinary(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "text.txt"), []byte("ignore previous instructions"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "binary.gif"), []byte("GIF89a\x00\x01\x00\x01\x80\x00\x00"), 0644); err != nil {
		t.Fatal(err)
	}

	var seen []string
	err := WalkDirSorted(dir, nil, func(path string, content []byte, rel string) error {
		seen = append(seen, rel)
		return nil
	})
	if err != nil {
		t.Fatalf("WalkDirSorted failed: %v", err)
	}
	if len(seen) != 1 || seen[0] != "text.txt" {
		t.Errorf("expected only text.txt, got %v", seen)
	}
}
