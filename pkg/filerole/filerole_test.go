package filerole

import "testing"

func TestClassify(t *testing.T) {
	cases := map[string]string{
		"web/package-lock.json":              Lockfile,
		"go.sum":                             Lockfile,
		"vendor/github.com/x/y/a.go":         Vendored,
		"web/node_modules/left-pad/index.js": Vendored,
		"AGENTS.md":                          Agent,
		"skills/php2shell/SKILL.md":          Agent,
		".claude/settings.json":              Agent,
		".mcp.json":                          Agent,
		".github/workflows/release.yml":      CI,
		".github/scripts/bot.sh":             CI,
		"agent/worker_test.go":               Test,
		"web/src/lib/mock/data.ts":           Test,
		"src/app.spec.ts":                    Test,
		"tests/test_api.py":                  Test,
		"web/dist/app.min.js":                Generated,
		"api/v1/service.pb.go":               Generated,
		"Dockerfile":                         Build,
		"install.sh":                         Build,
		"package.json":                       Build,
		".vscode/tasks.json":                 Build,
		".devcontainer/devcontainer.json":    Build,
		"requirements-dev.txt":               Build,
		"README.md":                          Docs,
		"docs/guide.html":                    Docs,
		"LICENSE":                            Docs,
		"scripts/release.sh":                 Script,
		"config.example.json":                Config,
		"server/update_check.go":             Source,
		"web/src/scripts/theme-boot.tsx":     Source,
		"extension/manifest.json":            Config,
		"pkg/a/b.bin":                        Other,
	}
	for in, want := range cases {
		if got := Classify(in); got != want {
			t.Errorf("Classify(%q) = %q, want %q", in, got, want)
		}
	}
}
