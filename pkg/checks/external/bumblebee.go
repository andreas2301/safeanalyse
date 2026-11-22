package external

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/user/safeanalyze/pkg/install"
	"github.com/user/safeanalyze/pkg/report"
)

// bumblebeeStage runs Perplexity's Bumblebee supply-chain exposure scanner in
// deep profile mode against a target repository root.
type bumblebeeStage struct {
	exposureCatalog string
	timeout         time.Duration
}

// NewBumblebeeStage creates a Bumblebee pipeline stage with default options.
func NewBumblebeeStage() *bumblebeeStage {
	return &bumblebeeStage{timeout: defaultScannerTimeout}
}

// NewBumblebeeStageWithOptions creates a Bumblebee stage with an explicit
// exposure catalog path (file or directory) and timeout.
func NewBumblebeeStageWithOptions(exposureCatalog string, timeout time.Duration) *bumblebeeStage {
	if timeout <= 0 {
		timeout = defaultScannerTimeout
	}
	return &bumblebeeStage{
		exposureCatalog: exposureCatalog,
		timeout:         timeout,
	}
}

// Name returns the stage name.
func (s *bumblebeeStage) Name() string { return "bumblebee" }

// Dependencies returns no dependencies.
func (s *bumblebeeStage) Dependencies() []string { return nil }

// resolveCatalog finds an exposure catalog from explicit config or the
// installed Bumblebee checkout's threat_intel directory.
func (s *bumblebeeStage) resolveCatalog() string {
	if s.exposureCatalog != "" {
		if _, err := os.Stat(s.exposureCatalog); err == nil {
			return s.exposureCatalog
		}
		return ""
	}
	return install.ResolveAsset("bumblebee", "threat_intel")
}

// Run executes bumblebee and parses its NDJSON exposure findings.
func (s *bumblebeeStage) Run(ctx context.Context, target string, input *report.Report) (*report.Report, error) {
	out := prepareReport(input, target)
	binary := resolveBinary("bumblebee")
	if binary == "" {
		return missingBinary(out, s.Name()), nil
	}
	catalog := s.resolveCatalog()
	if catalog == "" {
		out.AddError(s.Name(), "bumblebee requires an exposure catalog: set exposure_catalog in config or run 'safeanalyze install bumblebee'")
		return out, nil
	}

	args := []string{
		"scan",
		"--profile", "deep",
		"--root", target,
		"--exposure-catalog", catalog,
		"--findings-only",
		"--max-duration", s.timeout.String(),
	}
	output, stderr, _, err := runCommand(ctx, s.timeout, binary, args)
	if err != nil {
		if len(bytes.TrimSpace(output)) == 0 {
			msg := fmt.Sprintf("bumblebee execution failed: %v", err)
			if stderr != "" {
				msg += ": " + stderr
			}
			out.AddError(s.Name(), msg)
			return out, nil
		}
		// Bumblebee may exit non-zero when exposure findings are detected;
		// attempt to parse any NDJSON emitted on stdout before recording error.
		before := len(out.Findings)
		parsed, _ := parseBumblebeeOutput(out, output)
		if len(parsed.Findings) == before {
			out.AddError(s.Name(), fmt.Sprintf("bumblebee execution failed: %v", err))
		}
		return parsed, nil
	}
	return parseBumblebeeOutput(out, output)
}

func parseBumblebeeOutput(out *report.Report, data []byte) (*report.Report, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Buffer(make([]byte, 1024), 4*1024*1024)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		var rec struct {
			// Real Bumblebee v0.1+ NDJSON schema fields
			RecordType     string `json:"record_type"`
			FindingType    string `json:"finding_type"`
			CatalogID      string `json:"catalog_id"`
			CatalogName    string `json:"catalog_name"`
			Ecosystem      string `json:"ecosystem"`
			PackageName    string `json:"package_name"`
			NormalizedName string `json:"normalized_name"`
			Version        string `json:"version"`
			ProjectPath    string `json:"project_path"`
			SourceType     string `json:"source_type"`
			SourceFile     string `json:"source_file"`
			Confidence     string `json:"confidence"`
			Evidence       string `json:"evidence"`

			// Shared / legacy simplified fields
			Rule     string `json:"rule"`
			Title    string `json:"title"`
			Message  string `json:"message"`
			File     string `json:"file"`
			Path     string `json:"path"`
			Line     int    `json:"line"`
			Column   int    `json:"column"`
			Severity string `json:"severity"`
		}
		if err := json.Unmarshal(line, &rec); err != nil {
			out.AddError("bumblebee", fmt.Sprintf("parsing NDJSON line: %v", err))
			continue
		}

		// Skip non-finding NDJSON records emitted by real Bumblebee runs.
		if rec.RecordType != "" && rec.RecordType != "finding" {
			continue
		}

		if rec.RecordType == "finding" {
			pkg := coalesce(rec.PackageName, rec.NormalizedName)
			pkgRef := pkg
			if pkgRef != "" && rec.Version != "" {
				pkgRef += "@" + rec.Version
			}
			ruleID := coalesce(rec.CatalogID, rec.FindingType, "bumblebee-exposure")
			title := coalesce(rec.CatalogName, pkgRef, ruleID)
			file := coalesce(rec.SourceFile, rec.ProjectPath, rec.File, rec.Path)
			msg := title
			if rec.Evidence != "" {
				msg = fmt.Sprintf("%s: %s", title, rec.Evidence)
			}
			conf := report.ConfidenceDeterministic
			if rec.Confidence == "low" || rec.Confidence == "medium" {
				conf = report.ConfidenceHeuristic
			}
			attrs := map[string]string{}
			if rec.Ecosystem != "" {
				attrs["ecosystem"] = rec.Ecosystem
			}
			if pkg != "" {
				attrs["package"] = pkg
			}
			if rec.Version != "" {
				attrs["version"] = rec.Version
			}
			if rec.CatalogID != "" {
				attrs["catalog_id"] = rec.CatalogID
			}
			if rec.SourceType != "" {
				attrs["source_type"] = rec.SourceType
			}
			out.AddFinding(report.Finding{
				RuleID:      ruleID,
				Title:       title,
				Description: msg,
				Severity:    mapSeverity(rec.Severity),
				Category:    "supply_chain",
				File:        file,
				Line:        rec.Line,
				Column:      rec.Column,
				Match:       pkgRef,
				Message:     msg,
				Source:      "bumblebee",
				Confidence:  conf,
				Attributes:  attrs,
			})
			continue
		}

		// Legacy simplified format fallback.
		file := coalesce(rec.File, rec.Path)
		msg := coalesce(rec.Message, rec.Title)
		out.AddFinding(report.Finding{
			RuleID:      coalesce(rec.Rule, "bumblebee-finding"),
			Title:       coalesce(rec.Title, rec.Rule, "Bumblebee finding"),
			Description: msg,
			Severity:    mapSeverity(rec.Severity),
			Category:    "supply_chain",
			File:        file,
			Line:        rec.Line,
			Column:      rec.Column,
			Message:     msg,
			Source:      "bumblebee",
			Confidence:  report.ConfidenceHeuristic,
		})
	}
	if err := scanner.Err(); err != nil {
		out.AddError("bumblebee", fmt.Sprintf("reading NDJSON: %v", err))
	}
	return out, nil
}
