package intent

import (
	"fmt"
	"strings"

	"github.com/user/safeanalyze/pkg/filerole"
	"github.com/user/safeanalyze/pkg/report"
)

// callSite matches code that opens an outbound connection, per language.
var callSite = compileMulti(
	// Go
	`http\.(Get|Post|PostForm|Head|NewRequest|NewRequestWithContext)\(`, `\.Do\(req`, `net\.Dial`, `DialContext\(`,
	`websocket\.(Dial|DefaultDialer)`, `smtp\.(SendMail|Dial)`, `grpc\.(Dial|NewClient)\(`, `net\.Lookup`,
	// JavaScript / TypeScript
	`\bfetch\(`, `XMLHttpRequest`, `navigator\.sendBeacon`, `new WebSocket\(`, `new EventSource\(`,
	`\baxios(\.\w+)?\(`, `\bgot(\.\w+)?\(`, `\bsuperagent\b`, `https?\.(request|get)\(`, `\$\.(ajax|get|post)\(`,
	`new Image\(\)\.src`, `dns\.(resolve|lookup)\w*\(`,
	// Python
	`requests\.(get|post|put|patch|delete|request|Session)\(`, `urllib\.request`, `urlopen\(`, `httpx\.`,
	`aiohttp\.`, `socket\.(socket|create_connection)\(`, `http\.client\.`, `smtplib\.`, `ftplib\.`,
	// Shell / PowerShell
	`\bcurl\b`, `\bwget\b`, `Invoke-WebRequest`, `Invoke-RestMethod`, `Net\.WebClient`, `/dev/tcp/`, `\bnc\s+-`,
	// Java / C# / Ruby / PHP / Rust
	`HttpClient`, `URLConnection`, `OkHttpClient`, `WebClient\(`, `Net::HTTP`, `curl_exec\(`, `file_get_contents\(\s*['"]https?:`,
	`reqwest::`, `TcpStream::connect`, `ureq::`,
)

// guardPattern matches configuration, environment or opt-in/opt-out checks
// that may gate an outbound call.
var guardPattern = compileMulti(`(?i)os\.(Getenv|LookupEnv)`, `(?i)process\.env\.`, `(?i)os\.environ`, `(?i)getenv\(|import\.meta\.env|\b(cfg|config|conf|settings|opts?|options)\.|flag\.(Bool|String)|opt[_-]?(in|out)|_(ENABLED|DISABLE[D]?)\b|DO_NOT_TRACK|NO_TELEMETRY|TELEMETRY|consent|enabled\b`)

// telemetrySDK matches imports or initialisation of analytics, error
// reporting and feature-flag SDKs.
var telemetrySDK = compileMulti(
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]@sentry/[\w-]+['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]sentry-sdk['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]posthog(-js|-node)?['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]mixpanel(-browser)?['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]@amplitude/[\w-]+['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]amplitude-js['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]analytics-node['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]@segment/[\w-]+['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]react-ga4?['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]@vercel/analytics['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]@datadog/browser-[\w-]+['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]dd-trace['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]newrelic['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]@bugsnag/[\w-]+['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]rollbar['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]launchdarkly[\w-]*['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]@launchdarkly/[\w-]+['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]logrocket['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]@fullstory/[\w-]+['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]hotjar['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]@statsig/[\w-]+['"]`,
	`(?i)(from\s+|require\(\s*|import\(\s*|import\s+)['"]statsig-node['"]`,
	`(?i)\bimport\s+(sentry_sdk|posthog|mixpanel|rollbar|bugsnag|newrelic|ddtrace|launchdarkly)\b`,
	`(?i)github\.com/(getsentry/sentry-go|posthog/posthog-go|DataDog/dd-trace-go|newrelic/go-agent|launchdarkly/go-server-sdk|segmentio/analytics-go|bugsnag/bugsnag-go|rollbar/rollbar-go)`,
)

// selfUpdate matches self-update and version-check libraries or code.
var selfUpdate = compileMulti(
	`(?i)github\.com/(rhysd/go-github-selfupdate|creativeprojects/go-selfupdate|minio/selfupdate|inconshreveable/go-update|sanbornm/go-selfupdate)`,
	`(?i)electron-updater`,
	`(?i)autoUpdater\.`,
	`(?i)update-notifier`,
	`(?i)/releases/latest`,
)

const guardWindow = 20

func analyzeEgress(files []*File) []report.Finding {
	var out []report.Finding
	for _, f := range files {
		if f.Role == filerole.Lockfile {
			continue // registry URLs in lockfiles are handled by depsource
		}
		hasURL := urlPattern.MatchString(f.Content)
		sdkLoc := telemetrySDK.FindStringIndex(f.Content)
		updateLoc := selfUpdate.FindStringIndex(f.Content)
		if !hasURL && sdkLoc == nil && updateLoc == nil {
			continue // whole-file prefilter: nothing to report in this file
		}
		lines := f.Lines()
		guards := guardIndex(lines)
		calls := 0
		if sites := callSite.forText(f.Content); len(sites) > 0 {
			for _, l := range lines {
				if sites.MatchString(l) {
					calls++
				}
			}
		}
		code := filerole.IsCode(f.Role) || f.Role == filerole.Config || f.Role == filerole.Test

		if loc := sdkLoc; loc != nil && f.Role != filerole.Docs {
			line := lineOf(f.Content, loc[0])
			guarded := guardedAround(guards, line)
			tier, how := report.TierA, "no opt-in/opt-out guard found nearby: likely on by default"
			if guarded {
				tier, how = report.TierB, "an env/config guard is nearby: check whether it is opt-in"
			}
			out = append(out, finding("egress", "telemetry_sdk", report.SeverityHigh, tier, QCallHome, f, line,
				f.Content[loc[0]:loc[1]], "Telemetry/analytics SDK in use; "+how,
				map[string]string{"class": ClassTelemetry, "guard": guardString(guarded)}))
		}
		if loc := updateLoc; loc != nil && filerole.IsCode(f.Role) {
			line := lineOf(f.Content, loc[0])
			guarded := guardedAround(guards, line)
			tier := report.TierB
			if !guarded {
				tier = report.TierA
			}
			out = append(out, finding("egress", "update_check", report.SeverityMedium, tier, QCallHome, f, line,
				f.Content[loc[0]:loc[1]], "Self-update or version check contacts a remote server ("+guardString(guarded)+")",
				map[string]string{"class": "update", "guard": guardString(guarded)}))
		}

		seen := map[string]bool{}
		for i, l := range lines {
			if !hasURL {
				break
			}
			for _, raw := range urlPattern.FindAllString(l, -1) {
				ep := ParseEndpoint(raw)
				if ep.Class == ClassIgnore || ep.Class == ClassReference {
					continue
				}
				key := ep.Class + "|" + ep.Host
				if seen[key] {
					continue
				}
				seen[key] = true
				if fnd, ok := endpointFinding(f, i+1, l, ep, calls, code, guards); ok {
					out = append(out, fnd)
				}
			}
		}
	}
	return out
}

func endpointFinding(f *File, line int, text string, ep Endpoint, calls int, code bool, guards []int) (report.Finding, bool) {
	guarded := guardedAround(guards, line)
	// A URL in a comment documents; a URL in code without any call site
	// nearby is data (icon lists, docs bundles). Both stay visible below A.
	comment := !CodeLine(text)
	live := code && !comment && callNear(text, strings.Index(text, ep.Host), calls, f.Role)
	attrs := map[string]string{"host": ep.Host, "class": ep.Class, "guard": guardString(guarded), "call_sites": fmt.Sprint(calls)}
	sev, tier := report.SeverityLow, report.TierC
	msg := fmt.Sprintf("Hardcoded endpoint %s (%s)", ep.Host, ep.Class)
	switch ep.Class {
	case ClassExfil:
		sev, tier = report.SeverityCritical, report.TierB
		msg = fmt.Sprintf("Endpoint %s is a typical exfiltration/callback/tunnel channel", ep.Host)
		// A hardcoded webhook/bot token is attacker-owned infrastructure even
		// without a call site in this file (it may be read from here).
		if code && !comment && (live || exfilURL.MatchString(text)) {
			tier = report.TierA
		}
	case ClassRawIP:
		sev, tier = report.SeverityHigh, report.TierB
		msg = fmt.Sprintf("Hardcoded public IP address %s", ep.Host)
		switch {
		case InProse(text, strings.Index(text, ep.Host)):
			msg += " in description/prose text (not contacted directly; an AI agent reading it may use it as a target)"
		case live:
			tier = report.TierA
		}
	case ClassRecon:
		sev, tier = report.SeverityHigh, report.TierB
		msg = fmt.Sprintf("Public-IP/geolocation lookup service %s (host reconnaissance)", ep.Host)
	case ClassTelemetry:
		sev, tier = report.SeverityHigh, report.TierB
		if live && !guarded {
			tier = report.TierA
		}
		msg = fmt.Sprintf("Telemetry/analytics endpoint %s (%s)", ep.Host, guardString(guarded))
	case ClassAPI:
		sev, tier = report.SeverityMedium, report.TierC
		if !comment && code && (calls > 0 || f.Role == filerole.CI || f.Role == filerole.Build || f.Role == filerole.Script) {
			tier = report.TierB
		}
		msg = fmt.Sprintf("Third-party API %s receives data (%s)", ep.Host, guardString(guarded))
	case ClassInfra:
		sev, tier = report.SeverityInfo, report.TierC
		if isMirrorRegistry(ep.Host) && code {
			sev, tier = report.SeverityMedium, report.TierB
			msg = fmt.Sprintf("Third-party registry mirror %s", ep.Host)
		}
	case ClassUnknown:
		if !comment && code && f.Role != filerole.Test && (calls > 0 || f.Role == filerole.Build || f.Role == filerole.Script || f.Role == filerole.CI) {
			sev, tier = report.SeverityMedium, report.TierB
			msg = fmt.Sprintf("Hardcoded endpoint %s in a file with %d outbound call site(s) (%s)", ep.Host, calls, guardString(guarded))
		}
	}
	return finding("egress", "egress_"+ep.Class, sev, tier, QCallHome, f, line, text, msg, attrs), true
}

// guardIndex returns prefix counts of guard lines, so guardedAround answers
// in constant time. Long (minified) lines are checked on their first 4 KB.
func guardIndex(lines []string) []int {
	prefix := make([]int, len(lines)+1)
	for i, l := range lines {
		if len(l) > 4096 {
			l = l[:4096]
		}
		prefix[i+1] = prefix[i]
		if guardPattern.MatchString(l) {
			prefix[i+1]++
		}
	}
	return prefix
}

// guardedAround reports whether a guard line lies within guardWindow lines
// of the 1-based line.
func guardedAround(prefix []int, line int) bool {
	lo, hi := line-1-guardWindow, line+guardWindow
	if lo < 0 {
		lo = 0
	}
	if hi > len(prefix)-1 {
		hi = len(prefix) - 1
	}
	return prefix[hi]-prefix[lo] > 0
}

func guardString(guarded bool) string {
	if guarded {
		return "env/config guard nearby"
	}
	return "no guard found"
}

// callNear reports whether an outbound call site can use the endpoint at
// byte offset pos of line: anywhere in a normal file with call sites, but in
// a long (minified or data) line only within 300 bytes of the URL. Build,
// script and CI files fetch URLs with curl/wget, so any URL there counts.
func callNear(line string, pos, calls int, role string) bool {
	switch role {
	case filerole.Build, filerole.Script, filerole.CI:
		return true
	}
	if calls == 0 {
		return false
	}
	if len(line) <= 1000 || pos < 0 {
		return true
	}
	lo, hi := max(0, pos-300), min(len(line), pos+300)
	return callSite.MatchString(line[lo:hi])
}
