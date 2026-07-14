package ml

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// downloadFiles describes the files required to run the prompt-injection
// classifier. The ONNX model lives under the repo's onnx/ directory; the
// tokenizer and model config files live at the repository root.
var downloadFiles = []struct {
	local  string
	remote string
}{
	{local: "model.onnx", remote: "onnx/model.onnx"},
	{local: "config.json", remote: "config.json"},
	{local: "tokenizer.json", remote: "tokenizer.json"},
	{local: "tokenizer_config.json", remote: "tokenizer_config.json"},
	{local: "special_tokens_map.json", remote: "special_tokens_map.json"},
	{local: "added_tokens.json", remote: "added_tokens.json"},
	{local: "spm.model", remote: "spm.model"},
}

// DownloadModel downloads the protectai/deberta-v3-base-prompt-injection ONNX
// model and its tokenizer artifacts to the given directory. The files are
// placed in a subdirectory named after the model (deberta-v3-base-prompt-injection).
//
// Pass an empty string to use the default ~/.safeanalyze/models location.
func DownloadModel(dir string) error {
	if dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("resolving home directory: %w", err)
		}
		dir = filepath.Join(home, ".safeanalyze", "models")
	}

	modelDir := filepath.Join(dir, "deberta-v3-base-prompt-injection")
	if err := os.MkdirAll(modelDir, 0755); err != nil {
		return fmt.Errorf("creating model directory: %w", err)
	}

	client := &http.Client{
		// No overall timeout: the ONNX model is ~740 MB and may take a while
		// on slow links. The underlying transport will detect stuck connections.
	}

	for _, f := range downloadFiles {
		dest := filepath.Join(modelDir, f.local)
		if info, err := os.Stat(dest); err == nil && info.Size() > 0 {
			continue
		}

		url := fmt.Sprintf("https://huggingface.co/%s/resolve/main/%s", modelRepo, f.remote)
		if err := downloadFile(client, url, dest); err != nil {
			return fmt.Errorf("downloading %s: %w", f.remote, err)
		}
	}

	return nil
}

func downloadFile(client *http.Client, url, dest string) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %s", resp.Status)
	}

	tmp := dest + ".tmp"
	out, err := os.Create(tmp)
	if err != nil {
		return err
	}

	_, copyErr := io.Copy(out, resp.Body)
	closeErr := out.Close()
	if copyErr != nil {
		os.Remove(tmp)
		return copyErr
	}
	if closeErr != nil {
		os.Remove(tmp)
		return closeErr
	}

	return os.Rename(tmp, dest)
}
