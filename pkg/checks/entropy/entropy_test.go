package entropy

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestShannonEntropy(t *testing.T) {
	if e := shannonEntropy("aaaaaaaaaa"); e != 0 {
		t.Errorf("expected 0 entropy for constant string, got %f", e)
	}
	if e := shannonEntropy("aB3$xY9!pQ2@mN7#kL5%wZ8^"); e < 4.0 {
		t.Errorf("expected high entropy (>4.0) for random-looking string, got %f", e)
	}
}

func TestEntropyStageDetectsHighEntropyBlob(t *testing.T) {
	dir := t.TempDir()
	// Synthetic base64-like high-entropy string (>40 chars).
	blob := "const payload = \"QmFzZTY0RW5jb2RlZFBheWxvYWRXaXRoSGlnaEVudHJvcHlBbmRSYW5kb21CeXRlczEyMzQ1Njc4OTA=\"\n"
	if err := os.WriteFile(filepath.Join(dir, "app.js"), []byte(blob), 0644); err != nil {
		t.Fatal(err)
	}

	st := NewStage(4.5, 1<<20, nil)
	rep, err := st.Run(context.Background(), dir, nil)
	if err != nil {
		t.Fatalf("Stage.Run failed: %v", err)
	}
	if rep.Summary.FilesScanned != 1 {
		t.Errorf("expected 1 file scanned, got %d", rep.Summary.FilesScanned)
	}
	if len(rep.Findings) == 0 {
		t.Errorf("expected at least one high-entropy finding")
	}
}
