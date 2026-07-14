package ml

import (
	"testing"
	"unicode/utf8"

	"github.com/knights-analytics/hugot/pipelines"
)

func TestChunkTextBoundaries(t *testing.T) {
	cases := []struct {
		name    string
		text    string
		maxSize int
		overlap int
	}{
		{"short", "hello world", 100, 10},
		{"long ascii", "line1\nline2\nline3\nline4\nline5\n", 10, 2},
		{"multibyte", "こんにちは世界、このテキストは長いです。", 12, 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			chunks := chunkText(tc.text, tc.maxSize, tc.overlap)
			if len(chunks) == 0 {
				t.Fatal("expected at least one chunk")
			}
			for i, c := range chunks {
				if !utf8.ValidString(c.text) {
					t.Fatalf("chunk %d is not valid UTF-8", i)
				}
				if c.line < 1 {
					t.Fatalf("chunk %d line %d < 1", i, c.line)
				}
			}
		})
	}
}

func TestInjectionProbability(t *testing.T) {
	inj := injectionProbability([]pipelines.ClassificationOutput{{Label: "INJECTION", Score: 0.9}})
	if inj < 0.89 || inj > 0.91 {
		t.Fatalf("expected injection probability ~0.9, got %f", inj)
	}
	safe := injectionProbability([]pipelines.ClassificationOutput{{Label: "SAFE", Score: 0.9}})
	if safe < 0.09 || safe > 0.11 {
		t.Fatalf("expected injection probability ~0.1, got %f", safe)
	}
}

// Ensure the helper type is not used; we reference pipelines.ClassificationOutput directly.
