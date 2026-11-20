package intent

import (
	"regexp"
	"strings"

	"github.com/user/safeanalyze/pkg/filerole"
	"github.com/user/safeanalyze/pkg/report"
)

var ciUses = regexp.MustCompile(`^\s*(?:-\s*)?uses:\s*["']?([^@\s"']+)@([^\s"'#]+)`)
var ciDockerUses = regexp.MustCompile(`^\s*(?:-\s*)?uses:\s*["']?docker://([^\s"']+)`)
var ciRun = regexp.MustCompile(`^(\s*)(?:-\s*)?run:\s*(.*)$`)
var sha40 = regexp.MustCompile(`^[0-9a-f]{40}$`)

// untrustedContext are GitHub expressions an outside contributor controls.
var untrustedContext = regexp.MustCompile(`\$\{\{\s*(github\.event\.(issue|pull_request|comment|review|review_comment|discussion|discussion_comment|head_commit|commits|pages|workflow_run|release)\.[\w.\[\]*]*(title|body|message|name|ref|label|email|head_branch|default_branch|description)|github\.head_ref|inputs\.[\w-]+)\s*\}\}`)

var ciPRTarget = regexp.MustCompile(`(?m)^\s*(on:.*pull_request_target|pull_request_target\s*:|-\s*pull_request_target\b)`)
var ciHeadCheckout = regexp.MustCompile(`ref:\s*\$\{\{\s*(github\.event\.pull_request\.head\.(sha|ref)|github\.head_ref)`)
var ciWriteAll = regexp.MustCompile(`^\s*permissions:\s*write-all`)
var ciUnsafeCommands = regexp.MustCompile(`ACTIONS_ALLOW_UNSECURE_COMMANDS`)
var ciSecretUse = regexp.MustCompile(`\$\{\{\s*secrets\.(\w+)|\$\{?(\w*(TOKEN|SECRET|API_KEY|APIKEY|PASSWORD|PASS|KEY))\b`)
var ciSelfHosted = regexp.MustCompile(`runs-on:.*self-hosted`)
var ciPullRequest = regexp.MustCompile(`(?m)^\s*(on:.*\bpull_request\b|pull_request\s*:|-\s*pull_request\b)`)

func analyzeCI(files []*File) []report.Finding {
	var out []report.Finding
	for _, f := range files {
		if f.Role != filerole.CI {
			continue
		}
		yaml := strings.HasSuffix(f.Base(), ".yml") || strings.HasSuffix(f.Base(), ".yaml")
		if !yaml {
			out = append(out, ciScriptEgress(f)...)
			continue
		}
		lines := f.Lines()
		prTarget := ciPRTarget.MatchString(f.Content)
		if prTarget {
			if line, l := firstMatchLine(f, ciHeadCheckout); line > 0 {
				out = append(out, finding("ci", "ci_pr_target_head_checkout", report.SeverityCritical, report.TierA, QCI, f, line, l,
					"pull_request_target checks out the contributor's code with repository secrets and write token (\"pwn request\")", nil))
			} else {
				line, l := firstMatchLine(f, regexp.MustCompile(`pull_request_target`))
				out = append(out, finding("ci", "ci_pr_target", report.SeverityMedium, report.TierB, QCI, f, line, l,
					"pull_request_target runs with secrets on events from forks; review what it executes", nil))
			}
		}
		if ciSelfHosted.MatchString(f.Content) && (ciPullRequest.MatchString(f.Content) || prTarget) {
			line, l := firstMatchLine(f, ciSelfHosted)
			out = append(out, finding("ci", "ci_self_hosted_pr", report.SeverityHigh, report.TierB, QCI, f, line, l,
				"Self-hosted runner used for pull-request events: fork code can run on your machine", nil))
		}
		runIndent := -1
		for i, l := range lines {
			if m := ciRun.FindStringSubmatch(l); m != nil {
				runIndent = len(m[1])
				out = append(out, ciRunLine(f, i, l)...)
				continue
			}
			trimmed := strings.TrimSpace(l)
			indent := len(l) - len(strings.TrimLeft(l, " "))
			if runIndent >= 0 && trimmed != "" && indent > runIndent {
				out = append(out, ciRunLine(f, i, l)...)
				continue
			}
			if trimmed != "" {
				runIndent = -1
			}
			if m := ciUses.FindStringSubmatch(l); m != nil {
				action, ref := m[1], m[2]
				if strings.HasPrefix(action, "./") || sha40.MatchString(ref) {
					continue
				}
				tier := report.TierB
				if strings.HasPrefix(action, "actions/") || strings.HasPrefix(action, "github/") {
					tier = report.TierC
				}
				out = append(out, finding("ci", "ci_unpinned_action", report.SeverityMedium, tier, QCI, f, i+1, l,
					"Action "+action+" is referenced by mutable ref "+ref+" instead of a commit SHA", map[string]string{"action": action}))
			} else if m := ciDockerUses.FindStringSubmatch(l); m != nil && !strings.Contains(m[1], "@sha256:") {
				out = append(out, finding("ci", "ci_unpinned_action", report.SeverityMedium, report.TierB, QCI, f, i+1, l,
					"Docker action image is not pinned by digest", nil))
			}
			if ciWriteAll.MatchString(l) {
				out = append(out, finding("ci", "ci_write_all", report.SeverityMedium, report.TierB, QCI, f, i+1, l, "Workflow grants write-all permissions", nil))
			}
			if ciUnsafeCommands.MatchString(l) {
				out = append(out, finding("ci", "ci_unsafe_commands", report.SeverityHigh, report.TierA, QCI, f, i+1, l, "Deprecated unsafe workflow commands are re-enabled", nil))
			}
		}
	}
	return out
}

// ciRunLine checks one line of a run: block.
func ciRunLine(f *File, i int, l string) []report.Finding {
	var out []report.Finding
	if m := untrustedContext.FindString(l); m != "" {
		out = append(out, finding("ci", "ci_script_injection", report.SeverityCritical, report.TierA, QCI, f, i+1, l,
			"Attacker-controlled expression "+m+" is interpolated into a shell script (script injection)", nil))
	}
	if ciSecretUse.MatchString(l) {
		for _, raw := range urlPattern.FindAllString(l, -1) {
			ep := ParseEndpoint(raw)
			if ep.Class == ClassIgnore || ep.Host == "api.github.com" || strings.HasSuffix(ep.Host, "github.com") {
				continue
			}
			out = append(out, finding("ci", "ci_secret_egress", report.SeverityHigh, report.TierB, QCI, f, i+1, l,
				"A secret is sent to third-party host "+ep.Host+" from CI", map[string]string{"host": ep.Host}))
		}
	}
	return out
}

// ciScriptEgress checks helper scripts under .github/ for secrets sent to
// third-party hosts.
func ciScriptEgress(f *File) []report.Finding {
	if !ciSecretUse.MatchString(f.Content) {
		return nil
	}
	var out []report.Finding
	seen := map[string]bool{}
	for i, l := range f.Lines() {
		for _, raw := range urlPattern.FindAllString(l, -1) {
			ep := ParseEndpoint(raw)
			if ep.Class == ClassIgnore || ep.Class == ClassReference || strings.HasSuffix(ep.Host, "github.com") || seen[ep.Host] {
				continue
			}
			seen[ep.Host] = true
			out = append(out, finding("ci", "ci_secret_egress", report.SeverityMedium, report.TierB, QCI, f, i+1, l,
				"CI script that handles secrets/tokens contacts third-party host "+ep.Host, map[string]string{"host": ep.Host}))
		}
	}
	return out
}
