package intent

import (
	"net"
	"net/url"
	"regexp"
	"strings"
)

// Host classes used by the egress check.
const (
	ClassIgnore    = "ignore"    // localhost, private ranges, documentation hosts, placeholders
	ClassReference = "reference" // schema/namespace/license URLs, not contacted at runtime
	ClassExfil     = "exfil"     // channels typical for exfiltration, tunnelling or callbacks
	ClassRecon     = "recon"     // public-IP / geo lookup services
	ClassTelemetry = "telemetry" // analytics, error reporting, feature-flag services
	ClassAPI       = "api"       // third-party APIs (LLM, search, messaging) that receive data
	ClassInfra     = "infra"     // package registries, code hosting, CDNs
	ClassRawIP     = "raw_ip"    // a literal public IP address
	ClassUnknown   = "unknown"
)

// urlPattern extracts URLs from text.
var urlPattern = regexp.MustCompile(`(?i)\b(?:https?|wss?|ftp)://[^\s"'<>\x60()\[\]{}\\|^,;]+`)

// exfilURL matches URLs that embed a credential-bearing callback channel: a
// hardcoded Discord/Slack webhook or Telegram bot token is attacker-owned
// infrastructure, not a configurable integration.
var exfilURL = regexp.MustCompile(`(?i)(discord(app)?\.com/api/webhooks/\d+|hooks\.slack\.com/(services|workflows)/[A-Z0-9]+|api\.telegram\.org/bot\d{6,}:[\w-]{20,}|webhook\.office\.com/webhookb2/)`)

var exfilHosts = []string{
	"pastebin.com", "paste.ee", "hastebin.com", "ghostbin.", "rentry.co", "dpaste.", "termbin.com",
	"transfer.sh", "file.io", "0x0.st", "anonfiles.", "gofile.io", "catbox.moe", "temp.sh", "bashupload.com",
	"ngrok.io", "ngrok-free.app", "ngrok.app", "ngrok.dev", "trycloudflare.com", "serveo.net", "loca.lt",
	"localtunnel.me", "localhost.run", "pagekite.me", "bore.pub", "playit.gg",
	"requestbin.", "pipedream.net", "webhook.site", "beeceptor.com", "hookbin.com", "requestcatcher.com",
	"interact.sh", "interactsh.com", "oast.pro", "oast.live", "oast.site", "oast.online", "oast.fun", "oast.me",
	"burpcollaborator.net", "oastify.com", "canarytokens.com", "dnslog.cn", "ceye.io", "dnslog.link",
	"duckdns.org", "ddns.net", "hopto.org", "zapto.org", "sytes.net", "no-ip.", "dynu.net", "freedns.afraid.org",
	"portmap.io", "tunnelto.dev", ".onion",
}

var reconHosts = []string{
	"ipinfo.io", "ip-api.com", "ifconfig.me", "ifconfig.co", "api.ipify.org", "icanhazip.com", "checkip.amazonaws.com",
	"ipapi.co", "myexternalip.com", "wtfismyip.com", "ident.me", "ipecho.net",
}

var telemetryHosts = []string{
	"google-analytics.com", "googletagmanager.com", "analytics.google.com", "doubleclick.net",
	"sentry.io", "segment.io", "segment.com", "mixpanel.com", "amplitude.com", "posthog.com", "plausible.io",
	"umami.is", "datadoghq.com", "datadoghq.eu", "browser-intake-", "newrelic.com", "nr-data.net", "bugsnag.com",
	"rollbar.com", "honeycomb.io", "heapanalytics.com", "heap.io", "hotjar.com", "fullstory.com", "clarity.ms",
	"logrocket.", "launchdarkly.com", "statsig.com", "usefathom.com", "matomo.cloud", "app-measurement.com",
	"scarf.sh", "telemetry.", "stats.", "metrics.", "beacon.",
}

var apiHosts = []string{
	"api.openai.com", "api.anthropic.com", "generativelanguage.googleapis.com", "api.deepseek.com", "api.mistral.ai",
	"api.groq.com", "api.together.xyz", "openrouter.ai", "api.cohere.ai", "api.cohere.com", "dashscope.", "api.x.ai",
	"api.telegram.org", "discord.com", "discordapp.com", "hooks.slack.com", "slack.com", "api.twilio.com",
	"api.sendgrid.com", "oapi.dingtalk.com", "qyapi.weixin.qq.com", "open.feishu.cn", "api.search.brave.com",
	"api.tavily.com", "html.duckduckgo.com", "api.keenable.ai", "googleapis.com", "amazonaws.com", "azure.com",
	"api.github.com", "api.gitlab.com", "hooks.zapier.com", "api.pushover.net", "ntfy.sh",
}

var infraHosts = []string{
	"github.com", "githubusercontent.com", "gitlab.com", "bitbucket.org", "codeberg.org", "gitee.com",
	"registry.npmjs.org", "npmjs.com", "registry.yarnpkg.com", "pypi.org", "pythonhosted.org", "proxy.golang.org",
	"sum.golang.org", "pkg.go.dev", "golang.org", "go.dev", "crates.io", "rubygems.org", "repo.maven.apache.org",
	"repo1.maven.org", "packagist.org", "nuget.org", "cdn.jsdelivr.net", "unpkg.com", "cdnjs.cloudflare.com",
	"docker.io", "docker.com", "ghcr.io", "quay.io", "gcr.io", "deb.debian.org", "archive.ubuntu.com",
	"dl-cdn.alpinelinux.org", "fonts.googleapis.com", "fonts.gstatic.com", "nodejs.org", "python.org",
	"npmmirror.com", "goproxy.cn", "goproxy.io", "mirrors.aliyun.com", "tuna.tsinghua.edu.cn",
}

// referenceHosts are identifiers in schemas, XML namespaces and licenses.
var referenceHosts = []string{
	"www.w3.org", "w3.org", "json-schema.org", "schema.org", "schemas.", "xmlns.", "purl.org", "opensource.org",
	"spdx.org", "www.apache.org", "apache.org/licenses", "gnu.org", "creativecommons.org", "semver.org",
	"keepachangelog.com", "conventionalcommits.org", "ietf.org", "rfc-editor.org", "wikipedia.org",
	"img.shields.io", "shields.io", "badge.fury.io", "goreportcard.com", "codecov.io", "contributor-covenant.org",
	"arxiv.org", "doi.org", "stackoverflow.com", "developer.mozilla.org", "x.com", "twitter.com", "youtube.com",
	"linkedin.com", "medium.com", "reddit.com", "patreon.com", "opencollective.com", "tidelift.com", "feross.org",
	"buymeacoffee.com", "ko-fi.com", "github.com/sponsors",
}

// mirrorRegistryHosts are third-party registry mirrors: legitimate, but
// outside the official registry's integrity and malware takedown process.
var mirrorRegistryHosts = []string{"npmmirror.com", "goproxy.cn", "goproxy.io", "mirrors.aliyun.com", "tuna.tsinghua.edu.cn", "mirrors.cloud.tencent.com", "repo.huaweicloud.com"}

// Endpoint is one URL found in a file.
type Endpoint struct {
	URL   string
	Host  string
	Class string
}

// ParseEndpoint normalizes and classifies a raw URL.
func ParseEndpoint(raw string) Endpoint {
	raw = strings.TrimRight(raw, ".:")
	ep := Endpoint{URL: raw}
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		ep.Class = ClassIgnore
		return ep
	}
	host := strings.ToLower(u.Hostname())
	ep.Host = host
	ep.Class = classifyHost(host, strings.ToLower(raw))
	return ep
}

func classifyHost(host, lowerURL string) string {
	if host == "" || strings.ContainsAny(host, "${}%<>*") || !strings.Contains(host, ".") && net.ParseIP(host) == nil {
		return ClassIgnore
	}
	if exfilURL.MatchString(lowerURL) {
		return ClassExfil
	}
	if ip := net.ParseIP(host); ip != nil {
		if ip.IsLoopback() || ip.IsPrivate() || ip.IsUnspecified() || ip.IsLinkLocalUnicast() || isDocIP(ip) {
			if ip.String() == "169.254.169.254" {
				return ClassExfil // cloud metadata endpoint: credential theft target
			}
			return ClassIgnore
		}
		return ClassRawIP
	}
	if host == "localhost" || isDocHost(host) {
		return ClassIgnore
	}
	switch {
	case hostIn(host, lowerURL, referenceHosts):
		return ClassReference
	case hostIn(host, lowerURL, exfilHosts):
		return ClassExfil
	case hostIn(host, lowerURL, reconHosts):
		return ClassRecon
	case hostIn(host, lowerURL, apiHosts):
		return ClassAPI
	case hostIn(host, lowerURL, infraHosts):
		return ClassInfra
	case hostIn(host, lowerURL, telemetryHosts):
		return ClassTelemetry
	}
	return ClassUnknown
}

// hostIn matches host against entries: an entry with a leading or trailing
// dot or a dash is a substring marker, one with a slash matches the URL, any
// other entry matches the host or a parent domain.
func hostIn(host, lowerURL string, entries []string) bool {
	for _, e := range entries {
		switch {
		case strings.Contains(e, "/"):
			if strings.Contains(lowerURL, e) {
				return true
			}
		case strings.HasPrefix(e, "."):
			if strings.HasSuffix(host, e) {
				return true
			}
		case strings.HasSuffix(e, ".") || strings.HasSuffix(e, "-"):
			if strings.HasPrefix(host, e) || strings.Contains(host, "."+e) {
				return true
			}
		default:
			if host == e || strings.HasSuffix(host, "."+e) {
				return true
			}
		}
	}
	return false
}

func isDocHost(host string) bool {
	for _, suffix := range []string{".example", ".test", ".invalid", ".local", ".localhost", ".internal", ".lan", ".home.arpa", ".tld"} {
		if strings.HasSuffix(host, suffix) {
			return true
		}
	}
	for _, d := range []string{"example.com", "example.org", "example.net", "example.edu"} {
		if host == d || strings.HasSuffix(host, "."+d) {
			return true
		}
	}
	return false
}

var docNets = func() []*net.IPNet {
	var nets []*net.IPNet
	for _, c := range []string{"192.0.2.0/24", "198.51.100.0/24", "203.0.113.0/24", "100.64.0.0/10", "198.18.0.0/15", "2001:db8::/32"} {
		_, n, _ := net.ParseCIDR(c)
		nets = append(nets, n)
	}
	return nets
}()

func isDocIP(ip net.IP) bool {
	for _, n := range docNets {
		if n.Contains(ip) {
			return true
		}
	}
	return false
}

// isMirrorRegistry reports whether host is a third-party registry mirror.
func isMirrorRegistry(host string) bool {
	return hostIn(host, "", mirrorRegistryHosts)
}
