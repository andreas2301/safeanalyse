package intent

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/user/safeanalyze/pkg/report"
)

// highRiskPermissions grant access to user data or browser control.
var highRiskPermissions = map[string]string{
	"debugger": "attach the debugger to any tab (read/modify all traffic)", "cookies": "read cookies",
	"webRequest": "observe requests", "webRequestBlocking": "modify requests", "proxy": "route traffic",
	"nativeMessaging": "talk to native programs", "history": "read browsing history", "clipboardRead": "read the clipboard",
	"management": "manage other extensions", "downloads": "manage downloads", "privacy": "change privacy settings",
	"tabCapture": "capture tab content", "desktopCapture": "capture the screen", "declarativeNetRequestWithHostAccess": "rewrite requests",
	"scripting": "inject scripts", "identity": "access account tokens", "webNavigation": "observe navigation",
}

var allHostPatterns = map[string]bool{"<all_urls>": true, "*://*/*": true, "http://*/*": true, "https://*/*": true, "*://*/": true}

func analyzeExtensions(files []*File) []report.Finding {
	var out []report.Finding
	for _, f := range files {
		if f.Base() != "manifest.json" || !strings.Contains(f.Content, "manifest_version") {
			continue
		}
		var m struct {
			Permissions         []any  `json:"permissions"`
			OptionalPermissions []any  `json:"optional_permissions"`
			HostPermissions     []any  `json:"host_permissions"`
			UpdateURL           string `json:"update_url"`
			ContentScripts      []struct {
				Matches []string `json:"matches"`
			} `json:"content_scripts"`
			ExternallyConnectable struct {
				Matches []string `json:"matches"`
			} `json:"externally_connectable"`
			CSP any `json:"content_security_policy"`
		}
		if err := json.Unmarshal([]byte(f.Content), &m); err != nil {
			continue
		}
		allHosts := false
		var risky []string
		for _, p := range append(append(append([]any{}, m.Permissions...), m.OptionalPermissions...), m.HostPermissions...) {
			s, _ := p.(string)
			if allHostPatterns[s] {
				allHosts = true
			}
			if why, ok := highRiskPermissions[s]; ok {
				risky = append(risky, s+" ("+why+")")
			}
		}
		for _, cs := range m.ContentScripts {
			for _, match := range cs.Matches {
				if allHostPatterns[match] {
					allHosts = true
				}
			}
		}
		sort.Strings(risky)
		if allHosts || len(risky) > 0 {
			tier, sev := report.TierB, report.SeverityMedium
			if allHosts && len(risky) > 0 {
				tier, sev = report.TierA, report.SeverityHigh
			}
			scope := "selected sites"
			if allHosts {
				scope = "ALL sites"
			}
			out = append(out, finding("extensions", "extension_privileges", sev, tier, QPrivileges, f, lineOfKey(f, `"permissions"`), strings.Join(risky, ", "),
				fmt.Sprintf("Browser extension has access to %s with permissions: %s", scope, strings.Join(risky, ", ")), nil))
		}
		if m.UpdateURL != "" && !strings.Contains(m.UpdateURL, "clients2.google.com") && !strings.Contains(m.UpdateURL, "edge.microsoft.com") {
			out = append(out, finding("extensions", "extension_self_hosted_updates", report.SeverityHigh, report.TierA, QCallHome, f, lineOfKey(f, `"update_url"`), m.UpdateURL,
				"Extension updates itself from a non-store server: its code can change after review", map[string]string{"host": m.UpdateURL}))
		}
		for _, match := range m.ExternallyConnectable.Matches {
			if allHostPatterns[match] || strings.HasPrefix(match, "*://*.") {
				out = append(out, finding("extensions", "extension_externally_connectable", report.SeverityMedium, report.TierB, QPrivileges, f, lineOfKey(f, `"externally_connectable"`), match,
					"Any website can message this extension", nil))
				break
			}
		}
		if csp, _ := json.Marshal(m.CSP); strings.Contains(string(csp), "unsafe-eval") {
			out = append(out, finding("extensions", "extension_unsafe_eval", report.SeverityMedium, report.TierB, QRemoteExec, f, lineOfKey(f, `"content_security_policy"`), "unsafe-eval",
				"Extension CSP allows eval of strings", nil))
		}
	}
	return out
}
