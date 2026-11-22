// Package triage ranks findings into tiers, answers the fixed verdict
// questions and records what was (not) checked. It never removes findings:
// a false positive is preferred over a missed malicious behaviour, so noise is
// handled by ranking, not by dropping.
package triage

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/user/safeanalyze/pkg/checks/intent"
	"github.com/user/safeanalyze/pkg/filerole"
	"github.com/user/safeanalyze/pkg/report"
)

// Questions are the fixed verdict questions, in report order, with the
// stages that can answer them.
var Questions = []struct {
	ID, Text string
	Stages   []string
}{
	{intent.QCallHome, "Does the code contact external servers (call home) without an explicit opt-in?", []string{"egress"}},
	{intent.QAutorun, "Does code run automatically on install, build, import, IDE/devcontainer open or agent start?", []string{"autorun", "agentsurface"}},
	{intent.QExfil, "Is there a data-exfiltration chain (secrets or environment → network)?", []string{"chains"}},
	{intent.QRemoteExec, "Does it download, decode or evaluate code at runtime or install time?", []string{"chains", "installscripts"}},
	{intent.QDeps, "Do dependencies come from unexpected sources (forks, mirrors, git/URL, disabled checksums)?", []string{"depsource", "bumblebee"}},
	{intent.QCI, "Can CI leak secrets or run untrusted contributor code?", []string{"ci"}},
	{intent.QAgent, "Can the repository manipulate an AI agent that opens it (prompt injection, agent config)?", []string{"agentsurface", "yara", "llm"}},
	{intent.QPrivileges, "Does it request broad privileges or install persistence?", []string{"extensions", "chains"}},
	{intent.QHidden, "Is there hidden or obfuscated content (invisible Unicode, encoded payloads)?", []string{"hiddenchars", "entropy", "yara"}},
	{intent.QSecrets, "Does it contain hardcoded secrets?", []string{"yara", "trufflehog", "gitleaks"}},
	{intent.QLicense, "Does the license impose obligations or restrictions on your use?", []string{"license"}},
}

var promptInjectionRules = map[string]bool{
	"prompt_injection_comment": true, "multilingual_prompt_injection": true, "task_switch_announcement": true,
	"chat_template_boundary": true, "system_boundary": true, "indirect_prompt_injection": true,
	"delimiter_breakout": true, "encoded_prompt_injection": true, "output_constraint": true,
	"data_exfiltration_email": true, "hidden_attribute_email": true, "account_access_request": true,
}

// Assign sets FileRole on every finding and a Tier and Question on findings
// whose stage did not set them.
func Assign(rep *report.Report) {
	lines := lineReader(rep.Target)
	for i := range rep.Findings {
		f := &rep.Findings[i]
		if f.FileRole == "" && f.File != "" {
			f.FileRole = filerole.Classify(f.File)
		}
		if f.Tier != "" {
			continue
		}
		f.Tier, f.Question = legacyTier(f)
		// A keyword in a comment or in natural-language text cannot execute:
		// keep it visible as Tier B instead of verdict-relevant.
		if f.Tier == report.TierA && keywordRules[f.RuleID] && f.FileRole != filerole.Agent {
			if l := lines(f.File, f.Line); l != "" && (!intent.CodeLine(l) || intent.InProse(l, matchOffset(l, f))) {
				f.Tier = report.TierB
				f.Message += " (in a comment or prose, not code)"
			}
		}
	}
}

// matchOffset returns the byte offset of a finding's match in its line.
func matchOffset(line string, f *report.Finding) int {
	if f.Column > 0 {
		return f.Column - 1
	}
	return strings.Index(strings.ToLower(line), strings.ToLower(f.Match))
}

// keywordRules are pattern rules that fire on a keyword whose meaning
// depends on whether the line is code.
var keywordRules = map[string]bool{"backdoor_indicator": true, "suspicious_shell": true}

// lineReader returns a cached reader of line n (1-based) of a file under
// target, or "" when it cannot be read.
func lineReader(target string) func(file string, n int) string {
	cache := map[string][]string{}
	return func(file string, n int) string {
		if file == "" || n <= 0 {
			return ""
		}
		ls, ok := cache[file]
		if !ok {
			path := file
			if !filepath.IsAbs(path) {
				path = filepath.Join(target, file)
			}
			if data, err := os.ReadFile(path); err == nil && len(data) < 16<<20 {
				ls = strings.Split(string(data), "\n")
			}
			cache[file] = ls
		}
		if n > len(ls) {
			return ""
		}
		return ls[n-1]
	}
}

// legacyTier ranks findings of the pattern, entropy, hidden-character and
// external stages by rule and file role.
func legacyTier(f *report.Finding) (string, string) {
	role := f.FileRole
	test := role == filerole.Test
	noise := role == filerole.Lockfile || role == filerole.Generated || test
	switch f.Source {
	case "entropy":
		if f.RuleID == "entropy_base64_blob" && !noise && role != filerole.Docs && role != filerole.Vendored {
			return report.TierB, intent.QHidden
		}
		return report.TierC, intent.QHidden
	case "hiddenchars":
		invisible := f.RuleID == "hidden_char_bidi" || f.RuleID == "hidden_char_zero_width"
		switch {
		case invisible && !test:
			return report.TierA, intent.QHidden
		case invisible || !test:
			return report.TierB, intent.QHidden // tests run on `go test`/CI: keep visible
		}
		return report.TierC, intent.QHidden
	case "yara", "ml-prompt-injection", "prompt-injection-scanner", "llm":
		return yaraTier(f, role)
	case "trufflehog", "gitleaks":
		if test || role == filerole.Docs {
			return report.TierC, intent.QSecrets
		}
		if strings.Contains(strings.ToLower(f.Message), "verified") && !strings.Contains(strings.ToLower(f.Message), "unverified") {
			return report.TierA, intent.QSecrets
		}
		return report.TierB, intent.QSecrets
	case "bumblebee":
		if test {
			return report.TierC, intent.QDeps
		}
		if report.SeverityRank(f.Severity) >= report.SeverityRank(report.SeverityHigh) {
			return report.TierA, intent.QDeps
		}
		return report.TierB, intent.QDeps
	case "semgrep":
		if test {
			return report.TierC, ""
		}
		if report.SeverityRank(f.Severity) >= report.SeverityRank(report.SeverityCritical) {
			return report.TierA, intent.QRemoteExec
		}
		if report.SeverityRank(f.Severity) >= report.SeverityRank(report.SeverityHigh) {
			return report.TierB, intent.QRemoteExec
		}
		return report.TierC, ""
	}
	q := f.Question
	if test {
		return report.TierC, q
	}
	if q != "" && report.SeverityRank(f.Severity) >= report.SeverityRank(report.SeverityCritical) {
		return report.TierA, q
	}
	if report.SeverityRank(f.Severity) >= report.SeverityRank(report.SeverityHigh) {
		return report.TierB, q
	}
	return report.TierC, q
}

// weakPromptMatch reports matches that merely name a concept ("system
// prompt") instead of carrying an instruction.
func weakPromptMatch(match string) bool {
	m := strings.TrimSpace(match)
	return m == "system prompt" || m == "system prompts" || m == "the system prompt"
}

func yaraTier(f *report.Finding, role string) (string, string) {
	test := role == filerole.Test || role == filerole.Lockfile
	code := role == filerole.Source || role == filerole.Build || role == filerole.Script || role == filerole.CI || role == filerole.Generated || role == filerole.Vendored
	match := strings.ToLower(f.Match)
	switch {
	case promptInjectionRules[f.RuleID] || f.Source != "yara":
		strong := report.SeverityRank(f.Severity) >= report.SeverityRank(report.SeverityHigh) && !weakPromptMatch(match)
		switch {
		case role == filerole.Agent && strong:
			return report.TierA, intent.QAgent
		case role == filerole.Agent:
			return report.TierB, intent.QAgent
		case test:
			return report.TierC, intent.QAgent
		case report.SeverityRank(f.Severity) >= report.SeverityRank(report.SeverityHigh):
			return report.TierB, intent.QAgent
		}
		return report.TierC, intent.QAgent
	case f.RuleID == "backdoor_indicator":
		switch {
		case strings.Contains(match, "persistence"):
			return report.TierC, intent.QPrivileges
		case test:
			return report.TierC, intent.QPrivileges
		case strings.Contains(match, "backdoor"):
			return report.TierB, intent.QPrivileges
		case code || role == filerole.Agent:
			return report.TierA, intent.QPrivileges
		}
		return report.TierB, intent.QPrivileges
	case f.RuleID == "suspicious_shell":
		switch {
		case test:
			return report.TierC, intent.QRemoteExec
		case code && role != filerole.Source || role == filerole.Agent:
			return report.TierA, intent.QRemoteExec
		}
		return report.TierB, intent.QRemoteExec
	case f.RuleID == "obfuscated_javascript":
		if test || role == filerole.Generated || role == filerole.Vendored || role == filerole.Docs {
			return report.TierC, intent.QHidden
		}
		return report.TierB, intent.QHidden
	case f.RuleID == "data_exfiltration":
		if test || role == filerole.Docs {
			return report.TierC, intent.QCallHome
		}
		return report.TierB, intent.QCallHome
	case f.RuleID == "credential_hardcode":
		if test || role == filerole.Docs {
			return report.TierC, intent.QSecrets
		}
		return report.TierB, intent.QSecrets
	}
	// template_injection, suspicious_imports and any other pattern rule:
	// inventory-level noise unless critical outside tests.
	if !test && f.Severity == report.SeverityCritical {
		return report.TierB, ""
	}
	return report.TierC, ""
}

// Finalize computes verdicts and coverage. ran lists the stages that were part
// of the run; notes lists known limits to record in the coverage table.
func Finalize(rep *report.Report, ran []string, notes []report.CoverageItem) {
	ranSet := map[string]bool{}
	for _, s := range ran {
		ranSet[s] = true
	}
	stageErr := map[string]string{}
	for _, e := range rep.Errors {
		if _, ok := stageErr[e.Source]; !ok {
			stageErr[e.Source] = e.Message
		}
	}

	rep.Verdicts = nil
	for _, q := range Questions {
		v := report.Verdict{ID: q.ID, Question: q.Text}
		var evA, evB []string
		for _, f := range rep.Findings {
			if f.Question != q.ID {
				continue
			}
			switch f.Tier {
			case report.TierA:
				v.TierA++
				evA = append(evA, evidence(f))
			case report.TierB:
				v.TierB++
				evB = append(evB, evidence(f))
			}
		}
		checked := false
		for _, s := range q.Stages {
			if ranSet[s] && stageErr[s] == "" {
				checked = true
			}
		}
		switch {
		case v.TierA > 0:
			v.Answer = report.AnswerYes
			v.Evidence = capList(evA, 10)
		case v.TierB > 0:
			v.Answer = report.AnswerUnclear
			v.Evidence = capList(evB, 10)
		case checked:
			v.Answer = report.AnswerNo
		default:
			v.Answer = report.AnswerNotChecked
		}
		rep.Verdicts = append(rep.Verdicts, v)
	}

	rep.Coverage = nil
	names := append([]string(nil), ran...)
	sort.Strings(names)
	for _, s := range names {
		item := report.CoverageItem{Check: s, Status: "ran"}
		if msg, ok := stageErr[s]; ok {
			item.Status, item.Note = "error", msg
		}
		rep.Coverage = append(rep.Coverage, item)
	}
	rep.Coverage = append(rep.Coverage, notes...)
	rep.Metadata["tier_a"] = fmt.Sprint(rep.CountTier(report.TierA))
	rep.Metadata["tier_b"] = fmt.Sprint(rep.CountTier(report.TierB))
	rep.Metadata["tier_c"] = fmt.Sprint(rep.CountTier(report.TierC))
}

func evidence(f report.Finding) string {
	return fmt.Sprintf("%s:%d %s", f.File, f.Line, f.RuleID)
}

func capList(s []string, n int) []string {
	if len(s) <= n {
		return s
	}
	return append(s[:n:n], fmt.Sprintf("… and %d more", len(s)-n))
}
