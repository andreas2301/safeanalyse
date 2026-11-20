package intent

import (
	"regexp"
	"strings"

	"github.com/user/safeanalyze/pkg/filerole"
	"github.com/user/safeanalyze/pkg/report"
)

// pipeToShell matches remote content piped or substituted into an
// interpreter: the downloaded code runs without any verification.
var pipeToShell = compileMulti(
	`(?i)\b(curl|wget|fetch)\b[^|\n]*\|\s*(sudo\s+(-\w+\s+)*)?(env\s+\S+\s+)*(ba|z|da|k|fi)?sh\b`,
	`(?i)\b(curl|wget)\b[^|\n]*\|\s*(sudo\s+)?(python[0-9.]*|perl|ruby|node|php)\b`,
	`(?i)\b(ba|z)?sh\s+(-c\s+)?["']?\$\(\s*(curl|wget)`,
	`(?i)\b(ba|z)?sh\s+<\(\s*(curl|wget)`,
	`(?i)\bsource\s+<\(\s*(curl|wget)`,
	`(?i)\b(iex|Invoke-Expression)\b[^\n]*(iwr|irm|Invoke-WebRequest|Invoke-RestMethod|DownloadString|Net\.WebClient)`,
	`(?i)(iwr|irm|Invoke-WebRequest|Invoke-RestMethod|DownloadString)[^\n]*\|\s*(iex|Invoke-Expression)\b`,
	`(?i)\beval\s+["']?\$\(\s*(curl|wget)`,
)

// download matches commands that fetch a file to disk.
var download = regexp.MustCompile(`(?i)\bcurl\b[^\n]*\s(-o|-O|--output|--remote-name)\b|\bwget\s|^\s*ADD\s+https?://|Invoke-WebRequest[^\n]*-OutFile|DownloadFile\(|\bgit\s+clone\b`)

// verification matches integrity checks of downloaded files.
var verification = regexp.MustCompile(`(?i)sha(1|224|256|384|512)sum|shasum|sha256:|--checksum|checksums?\.txt|gpg\s+--verify|cosign\s+verify|minisign|Get-FileHash|slsa-verifier|\bsha256\b|--branch\s+\S+|checkout\s+[0-9a-f]{7,40}`)

var dockerFrom = regexp.MustCompile(`(?i)^\s*FROM\s+(--platform=\S+\s+)?(\S+)`)

func analyzeInstallScripts(files []*File) []report.Finding {
	var out []report.Finding
	for _, f := range files {
		switch f.Role {
		case filerole.Lockfile, filerole.Vendored, filerole.Generated, filerole.Config, filerole.Other:
			continue
		}
		if f.Role == filerole.Source && !strings.HasSuffix(f.Base(), ".py") && !strings.HasSuffix(f.Base(), ".js") && !strings.HasSuffix(f.Base(), ".ts") && !strings.HasSuffix(f.Base(), ".go") {
			continue
		}
		lines := f.Lines()
		for i, l := range lines {
			if pipeToShell.MatchString(l) {
				tier, sev := report.TierA, report.SeverityCritical
				msg := "Remote script is piped into an interpreter: it runs unverified code from the network"
				if f.Role == filerole.Source {
					tier, sev = report.TierB, report.SeverityHigh
					msg += " (string inside source code)"
				}
				out = append(out, finding("installscripts", "remote_script_exec", sev, tier, QRemoteExec, f, i+1, l, msg, nil))
			}
		}
		if f.Role == filerole.Source || f.Role == filerole.Docs || f.Role == filerole.Test {
			continue
		}
		if line, l := firstMatchLine(f, download); line > 0 && !verification.MatchString(f.Content) {
			out = append(out, finding("installscripts", "download_without_verification", report.SeverityMedium, report.TierB, QRemoteExec, f, line, l,
				"Files are downloaded or cloned without a checksum, signature or pinned revision in this file", nil))
		}
		if strings.HasPrefix(f.Base(), "dockerfile") || strings.HasSuffix(f.Base(), ".dockerfile") {
			for i, l := range lines {
				m := dockerFrom.FindStringSubmatch(l)
				if m == nil || strings.Contains(m[2], "@sha256:") || strings.EqualFold(m[2], "scratch") || !strings.ContainsAny(m[2], ":/.") && !strings.Contains(m[2], "$") {
					continue // digest-pinned, scratch, or a build stage name
				}
				tier := report.TierC
				if strings.HasSuffix(m[2], ":latest") || !strings.Contains(m[2], ":") {
					tier = report.TierB
				}
				out = append(out, finding("installscripts", "docker_base_not_digest_pinned", report.SeverityLow, tier, QDeps, f, i+1, l,
					"Base image is referenced by a mutable tag instead of a digest", nil))
			}
		}
	}
	return out
}
