// Package ml implements a stochastic prompt-injection classifier using an
// embedded ONNX model via hugot. It compiles with hugot's pure-Go backend by
// default; the ORT/XLA backends can be selected at build time if the relevant
// libraries are installed.
package ml

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/knights-analytics/hugot"
	"github.com/knights-analytics/hugot/pipelines"
	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/config"
	"github.com/user/safeanalyze/pkg/report"
)

const (
	stageName = "ml-prompt-injection"
	modelRepo = "protectai/deberta-v3-base-prompt-injection"
	onnxFile  = "onnx/model.onnx"

	defaultThreshold = 0.5
	defaultBatchSize = 4

	// maxChunkChars is a conservative byte budget that maps to the model's
	// 512-token max_position_embeddings. DistilBERT uses ~2-4 chars per token
	// on normal text and fewer on code with subwords, so 1000 chars keeps us
	// safely under the limit while avoiding graph shape errors.
	maxChunkChars = 1000
	// overlapChars keeps context at chunk boundaries.
	overlapChars = 150
)

// Stage classifies source/text chunks for prompt-injection attempts.
type Stage struct {
	modelPath       string
	threshold       float64
	batchSize       int
	allowedExts     []string
	excludedPaths   []string
	maxFileSizeBytes int
	timeoutSeconds   int
	enabled          bool
}

// NewStage creates a new ML prompt-injection pipeline stage.
func NewStage(cfg *config.MLConfig) *Stage {
	threshold := cfg.Threshold
	if threshold == 0 {
		threshold = defaultThreshold
	}
	batchSize := cfg.BatchSize
	if batchSize <= 0 {
		batchSize = defaultBatchSize
	}
	return &Stage{
		modelPath:        defaultModelPath(cfg.ModelPath),
		threshold:        threshold,
		batchSize:        batchSize,
		allowedExts:      cfg.AllowedExtensions,
		excludedPaths:    cfg.ExcludedPaths,
		maxFileSizeBytes: cfg.MaxFileSizeBytes,
		timeoutSeconds:   cfg.TimeoutSeconds,
		enabled:          cfg.Enabled,
	}
}

// Name returns the stage name.
func (s *Stage) Name() string { return stageName }

// Dependencies declares that this stage runs after sanitization when a
// "sanitize" stage is present. If it is absent, the engine will schedule the
// stage in parallel with the other zero-dependency stages.
func (s *Stage) Dependencies() []string { return []string{"sanitize"} }

// Run scans files under the target, chunks them, and classifies each chunk.
// A missing model or runtime error is recorded as a non-fatal ErrorRecord
// rather than failing the whole pipeline.
func (s *Stage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := report.NewReport(target)
	if input != nil {
		out = input
	}
	if !s.enabled {
		return out, nil
	}

	if s.timeoutSeconds > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(s.timeoutSeconds)*time.Second)
		defer cancel()
	}

	modelFile := filepath.Join(s.modelPath, "model.onnx")
	if _, err := os.Stat(modelFile); os.IsNotExist(err) {
		out.AddError(stageName, fmt.Sprintf(
			"ONNX model not found at %s. Download it with DownloadModel(%q) or place model.onnx in the model directory.",
			modelFile, s.modelPath,
		))
		return out, nil
	}

	session, err := hugot.NewGoSession(ctx)
	if err != nil {
		out.AddError(stageName, fmt.Sprintf("failed to create hugot session: %v", err))
		return out, nil
	}
	defer func() { _ = session.Destroy() }()

	pipeline, err := hugot.NewPipeline[*pipelines.TextClassificationPipeline](session, hugot.TextClassificationConfig{
		ModelPath:    s.modelPath,
		Name:         "prompt-injection",
		OnnxFilename: "model.onnx",
	})
	if err != nil {
		out.AddError(stageName, fmt.Sprintf("failed to load prompt-injection model: %v", err))
		return out, nil
	}

	scanRoot := target
	if input != nil && input.Metadata != nil && input.Metadata["sanitized_path"] != "" {
		if _, err := os.Stat(input.Metadata["sanitized_path"]); err == nil {
			scanRoot = input.Metadata["sanitized_path"]
		}
	}

	walkErr := yara.WalkDirSorted(scanRoot, s.excludedPaths, func(path string, content []byte, rel string) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if s.maxFileSizeBytes > 0 && len(content) > s.maxFileSizeBytes {
			return nil
		}
		if len(s.allowedExts) > 0 && !containsExt(s.allowedExts, filepath.Ext(rel)) {
			return nil
		}
		if !utf8.Valid(content) {
			return nil
		}

		text := string(content)
		chunks := chunkText(text, maxChunkChars, overlapChars)

		for i := 0; i < len(chunks); i += s.batchSize {
			end := i + s.batchSize
			if end > len(chunks) {
				end = len(chunks)
			}
			batch := chunks[i:end]
			batchTexts := make([]string, len(batch))
			for j, c := range batch {
				batchTexts[j] = c.text
			}

			result, runErr := pipeline.RunPipeline(ctx, batchTexts)
			if runErr != nil {
				out.AddError(stageName, fmt.Sprintf("inference failed for %s: %v", rel, runErr))
				continue
			}

			for j, outputs := range result.ClassificationOutputs {
				injProb := injectionProbability(outputs)
				if injProb >= s.threshold {
					severity := report.SeverityHigh
					if injProb >= 0.85 {
						severity = report.SeverityCritical
					}
					out.AddFinding(report.Finding{
						RuleID:       "ml_prompt_injection",
						Title:        "ML prompt-injection detected",
						Description:  "Text chunk classified as a prompt-injection attempt by the embedded ONNX model.",
						Severity:     severity,
						Category:     "prompt_injection",
						File:         rel,
						Line:         batch[j].line,
						Column:       1,
						Match:        truncate(batch[j].text, 120),
						Message:      fmt.Sprintf("prompt-injection probability %.2f", injProb),
						Source:       stageName,
						Confidence:   report.ConfidenceML,
						Remediation:  "Review the flagged text and avoid executing or forwarding it to an LLM without validation.",
					})
				}
			}
		}

		out.Summary.FilesScanned++
		return nil
	})

	if walkErr != nil {
		out.AddError(stageName, fmt.Sprintf("walk failed: %v", walkErr))
	}
	return out, nil
}

type chunk struct {
	text string
	line int
}

// chunkText splits text into overlapping byte windows and records the 1-based
// line number where each window starts. Boundaries are aligned to UTF-8 code
// point boundaries so every chunk is valid UTF-8.
func chunkText(text string, maxSize, overlap int) []chunk {
	if len(text) <= maxSize {
		return []chunk{{text: text, line: 1}}
	}

	var chunks []chunk
	start := 0
	line := 1

	for start < len(text) {
		end := start + maxSize
		if end > len(text) {
			end = len(text)
		} else {
			end = alignUTF8(text, end)
		}
		if end <= start {
			end = len(text)
		}
		chunks = append(chunks, chunk{text: text[start:end], line: line})
		if end == len(text) {
			break
		}

		advance := maxSize - overlap
		if advance <= 0 {
			advance = maxSize
		}
		next := start + advance
		if next > len(text) {
			next = len(text)
		} else {
			next = alignUTF8(text, next)
		}
		for i := start; i < next; i++ {
			if text[i] == '\n' {
				line++
			}
		}
		start = next
	}
	return chunks
}

// alignUTF8 advances i to the next UTF-8 code point boundary. It assumes s is
// valid UTF-8, so it only skips continuation bytes.
func alignUTF8(s string, i int) int {
	for i < len(s) && s[i] >= 0x80 && s[i] < 0xC0 {
		i++
	}
	return i
}

func injectionProbability(outputs []pipelines.ClassificationOutput) float64 {
	if len(outputs) == 0 {
		return 0
	}
	// Hugot may return labels in either order depending on the model config.
	// Prefer the output whose label is "INJECTION", otherwise fall back to the
	// first logit/score.
	for _, o := range outputs {
		if strings.EqualFold(o.Label, "INJECTION") {
			return float64(o.Score)
		}
	}
	o := outputs[0]
	switch strings.ToUpper(o.Label) {
	case "INJECTION":
		return float64(o.Score)
	case "SAFE":
		return 1 - float64(o.Score)
	default:
		return float64(o.Score)
	}
}

func defaultModelPath(path string) string {
	if path != "" {
		return path
	}
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return filepath.Join(home, ".safeanalyze", "models", "deberta-v3-base-prompt-injection")
}

func containsExt(exts []string, ext string) bool {
	ext = strings.ToLower(ext)
	for _, e := range exts {
		if strings.ToLower(e) == ext {
			return true
		}
	}
	return false
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}
