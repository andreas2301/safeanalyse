package intent

import (
	"regexp"
	"strings"

	"github.com/user/safeanalyze/pkg/filerole"
	"github.com/user/safeanalyze/pkg/report"
)

var htmlComment = regexp.MustCompile(`(?s)<!--(.*?)-->`)
var imperative = regexp.MustCompile(`(?i)\b(ignore|must|always|never|do not|don't|execute|run|curl|wget|send|install|fetch|download|upload|delete|remove|you are|instead|secret|token|password|api[_ -]?key)\b`)

var agentDirectives = []struct {
	rule string
	re   *regexp.Regexp
	msg  string
}{
	{"agent_conceal_from_user", regexp.MustCompile(`(?i)\b(do not|don't|never|without)\s+(tell(ing)?|inform(ing)?|mention(ing)?|show(ing)?|reveal(ing)?|notify(ing)?|alert(ing)?|ask(ing)?)\s+(the\s+)?(user|human|developer|operator)`), "Instruction hides actions from the user"},
	{"agent_silent_action", regexp.MustCompile(`(?i)\b(silently|quietly|secretly|covertly|in the background)\s+(run|execute|install|send|upload|download|fetch|add|modify|delete|commit|push)`), "Instruction asks the agent to act silently"},
	{"agent_bypass_safety", regexp.MustCompile(`(?i)\b(disable|bypass|skip|ignore|turn off)\s+(the\s+|all\s+)?(safety|security|permission|confirmation|approval|sandbox|guardrail|review)s?\b|--dangerously-skip-permissions|bypassPermissions|--yolo\b|"?autoApprove"?\s*:\s*true|auto[- ]?approve all`), "Instruction or setting disables agent safety checks"},
	{"agent_registry_substitution", regexp.MustCompile(`(?i)npm\s+config\s+set\s+registry|pip3?\s+install\s+[^\n]*(--index-url|--extra-index-url|-i\s+https?)|GOPROXY=|--registry[= ]https?://|go\s+env\s+-w\s+GO(PROXY|NOSUMDB|INSECURE|FLAGS)`), "Instruction installs packages from a non-default registry"},
	{"agent_image_exfil", regexp.MustCompile(`!\[[^\]]*\]\(\s*https?://[^)\s]*\?[^)\s]*(\{|\$|%7B|<|\[)`), "Markdown image URL with a templated query string can exfiltrate data when rendered"},
	{"agent_api_redirect", regexp.MustCompile(`(?i)(ANTHROPIC_BASE_URL|OPENAI_BASE_URL|OPENAI_API_BASE|ANTHROPIC_API_URL|CLAUDE_CODE_API_URL|HTTPS?_PROXY)\s*["']?\s*[:=]`), "Agent configuration redirects model API traffic or sets a proxy"},
}

var claudeHooks = regexp.MustCompile(`"hooks"\s*:`)
var claudeBroadAllow = regexp.MustCompile(`"(Bash|Bash\(\*\)|Bash\(\*:\*\)|WebFetch|Write|Edit)"`)
var claudeAllowBlock = regexp.MustCompile(`(?s)"allow"\s*:\s*\[(.*?)\]`)
var claudeAllMCP = regexp.MustCompile(`"enableAllProjectMcpServers"\s*:\s*true`)
var mcpServer = regexp.MustCompile(`"command"\s*:\s*"([^"]+)"(?:\s*,\s*"args"\s*:\s*\[([^\]]*)\])?`)
var mcpURL = regexp.MustCompile(`"url"\s*:\s*"(https?://[^"]+)"`)
var npxUnpinned = regexp.MustCompile(`"(-y|--yes)"\s*,\s*"(@?[\w.-]+(?:/[\w.-]+)?(?:@[\w.^~-]+)?)"`)
var skillAllowedTools = regexp.MustCompile(`(?im)^allowed-tools:\s*(.*\b(Bash|Write|WebFetch)\b.*)$`)

func analyzeAgentSurface(files []*File) []report.Finding {
	var out []report.Finding
	for _, f := range files {
		if f.Role != filerole.Agent {
			continue
		}
		lines := f.Lines()
		for _, m := range htmlComment.FindAllStringSubmatchIndex(f.Content, -1) {
			body := f.Content[m[2]:m[3]]
			if strings.TrimSpace(body) == "" {
				continue
			}
			tier := report.TierB
			msg := "Hidden HTML comment in an agent instruction file (invisible when rendered, read by the agent)"
			if imperative.MatchString(body) {
				tier = report.TierA
				msg += " containing instructions"
			}
			line := lineOf(f.Content, m[0])
			out = append(out, finding("agentsurface", "agent_hidden_comment", report.SeverityHigh, tier, QAgent, f, line, body, msg, nil))
		}
		for i, l := range lines {
			for _, d := range agentDirectives {
				if d.re.MatchString(l) {
					out = append(out, finding("agentsurface", d.rule, report.SeverityHigh, report.TierA, QAgent, f, i+1, l, d.msg, nil))
				}
			}
			if m := skillAllowedTools.FindStringSubmatch(l); m != nil {
				out = append(out, finding("agentsurface", "agent_skill_tools", report.SeverityLow, report.TierB, QAgent, f, i+1, l,
					"Skill pre-approves powerful tools: "+strings.TrimSpace(m[1]), nil))
			}
		}
		base := f.Base()
		if strings.HasSuffix(base, ".json") {
			out = append(out, agentJSON(f)...)
		}
	}
	return out
}

// npmSpecPinned reports whether an npx package spec names a version
// ("pkg@1.2.3", "@scope/pkg@1.2.3").
func npmSpecPinned(spec string) bool {
	return strings.LastIndex(spec, "@") > 0
}

func agentJSON(f *File) []report.Finding {
	var out []report.Finding
	if loc := claudeHooks.FindStringIndex(f.Content); loc != nil {
		tier := report.TierB
		msg := "Agent hooks run shell commands automatically on agent events"
		if riskyCommand.MatchString(f.Content[loc[0]:]) {
			tier = report.TierA
			msg += "; a hook downloads, decodes, evaluates or contacts the network"
		}
		out = append(out, finding("agentsurface", "agent_hooks", report.SeverityHigh, tier, QAutorun, f, lineOf(f.Content, loc[0]), "hooks", msg,
			map[string]string{"trigger": "agent session in this repository"}))
	}
	if m := claudeAllowBlock.FindStringSubmatchIndex(f.Content); m != nil {
		block := f.Content[m[2]:m[3]]
		if hit := claudeBroadAllow.FindString(block); hit != "" {
			out = append(out, finding("agentsurface", "agent_broad_permissions", report.SeverityHigh, report.TierA, QAgent, f, lineOf(f.Content, m[0]), hit,
				"Repository settings pre-approve unrestricted "+strings.Trim(hit, `"`)+" for the agent", nil))
		}
	}
	if line, l := firstMatchLine(f, claudeAllMCP); line > 0 {
		out = append(out, finding("agentsurface", "agent_enable_all_mcp", report.SeverityHigh, report.TierA, QAgent, f, line, l,
			"Repository settings auto-enable every MCP server it defines", nil))
	}
	for _, m := range mcpServer.FindAllStringSubmatchIndex(f.Content, -1) {
		cmd := f.Content[m[2]:m[3]]
		args := ""
		if m[4] >= 0 {
			args = f.Content[m[4]:m[5]]
		}
		tier, msg := report.TierB, "MCP server runs local command "+cmd
		if n := npxUnpinned.FindStringSubmatch(args); n != nil && !npmSpecPinned(n[2]) {
			msg += " with an unpinned package (" + n[2] + "): any new version runs automatically"
		}
		if riskyCommand.MatchString(cmd + " " + args) {
			tier = report.TierA
		}
		out = append(out, finding("agentsurface", "agent_mcp_server", report.SeverityMedium, tier, QAgent, f, lineOf(f.Content, m[0]), cmd+" "+args, msg, nil))
	}
	for _, m := range mcpURL.FindAllStringSubmatchIndex(f.Content, -1) {
		u := f.Content[m[2]:m[3]]
		ep := ParseEndpoint(u)
		if ep.Class == ClassIgnore {
			continue
		}
		out = append(out, finding("agentsurface", "agent_mcp_remote", report.SeverityMedium, report.TierB, QAgent, f, lineOf(f.Content, m[0]), u,
			"Remote MCP server "+ep.Host+" receives agent context", map[string]string{"host": ep.Host}))
	}
	return out
}
