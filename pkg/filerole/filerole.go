// Package filerole classifies repository files by the role they play (source,
// test, docs, lockfile, CI, build, agent instructions, ...). Triage uses the
// role to rank findings: the same pattern means something different in a
// lockfile, a test fixture and an install script.
package filerole

import (
	"path"
	"strings"
)

// Roles, checked in the order Classify lists them.
const (
	Lockfile  = "lockfile"
	Vendored  = "vendored"
	Agent     = "agent"
	CI        = "ci"
	Test      = "test"
	Generated = "generated"
	Build     = "build"
	Docs      = "docs"
	Script    = "script"
	Config    = "config"
	Source    = "source"
	Other     = "other"
)

var lockfiles = map[string]bool{
	"package-lock.json": true, "npm-shrinkwrap.json": true, "yarn.lock": true,
	"pnpm-lock.yaml": true, "bun.lockb": true, "go.sum": true, "cargo.lock": true,
	"poetry.lock": true, "pipfile.lock": true, "composer.lock": true,
	"gemfile.lock": true, "uv.lock": true, "pdm.lock": true, "packages.lock.json": true,
}

// agentFiles are instruction or configuration files read by AI coding agents.
var agentFiles = map[string]bool{
	"agents.md": true, "claude.md": true, "claude.local.md": true, "gemini.md": true,
	"skill.md": true, ".cursorrules": true, ".windsurfrules": true, ".clinerules": true,
	"copilot-instructions.md": true, ".mcp.json": true, "mcp.json": true,
	".aider.conf.yml": true, "conventions.md": true,
}

var agentDirs = []string{".claude/", ".cursor/", ".windsurf/", ".clinerules/", ".roo/", ".continue/", ".github/instructions/", ".github/prompts/"}

var buildFiles = map[string]bool{
	"makefile": true, "gnumakefile": true, "cmakelists.txt": true, "build.rs": true,
	"setup.py": true, "setup.cfg": true, "pyproject.toml": true, "package.json": true,
	"go.mod": true, "cargo.toml": true, "gemfile": true, "rakefile": true, "build.gradle": true,
	"build.gradle.kts": true, "settings.gradle": true, "pom.xml": true, "justfile": true,
	"taskfile.yml": true, "magefile.go": true, ".envrc": true, ".npmrc": true, ".yarnrc": true,
	".yarnrc.yml": true, "pip.conf": true, "lefthook.yml": true, ".pre-commit-config.yaml": true,
	"vcpkg.json": true, "meson.build": true, "configure": true, "configure.ac": true,
}

var buildDirs = []string{".devcontainer/", ".husky/", ".githooks/", "gradle/wrapper/", ".mvn/", ".cargo/"}

var testDirs = []string{"test/", "tests/", "__tests__/", "testdata/", "fixtures/", "__fixtures__/", "__mocks__/", "mock/", "mocks/", "spec/", "e2e/"}

var docExts = map[string]bool{".md": true, ".rst": true, ".adoc": true, ".txt": true, ".markdown": true}

var scriptExts = map[string]bool{".sh": true, ".bash": true, ".zsh": true, ".fish": true, ".ps1": true, ".psm1": true, ".bat": true, ".cmd": true}

var configExts = map[string]bool{".json": true, ".yaml": true, ".yml": true, ".toml": true, ".ini": true, ".cfg": true, ".conf": true, ".env": true, ".xml": true, ".plist": true}

var sourceExts = map[string]bool{
	".go": true, ".py": true, ".js": true, ".mjs": true, ".cjs": true, ".ts": true, ".mts": true,
	".cts": true, ".jsx": true, ".tsx": true, ".rs": true, ".java": true, ".kt": true, ".kts": true,
	".scala": true, ".c": true, ".cc": true, ".cpp": true, ".h": true, ".hpp": true, ".cs": true,
	".rb": true, ".php": true, ".swift": true, ".m": true, ".lua": true, ".pl": true, ".pm": true,
	".r": true, ".dart": true, ".ex": true, ".exs": true, ".erl": true, ".hs": true, ".vue": true,
	".svelte": true, ".astro": true, ".html": true, ".htm": true, ".sql": true, ".s": true, ".asm": true,
	".pth": true, ".vbs": true, ".groovy": true, ".nim": true, ".zig": true, ".v": true, ".jl": true,
}

// Classify returns the role of the file at the slash- or OS-separated relative
// path rel.
func Classify(rel string) string {
	p := strings.ToLower(strings.ReplaceAll(rel, "\\", "/"))
	base := path.Base(p)
	ext := path.Ext(base)
	dir := path.Dir(p) + "/"
	if dir == "./" {
		dir = ""
	}

	switch {
	case lockfiles[base]:
		return Lockfile
	case hasDir(dir, "node_modules/", "vendor/", "third_party/", "third-party/", "bower_components/", "site-packages/"):
		return Vendored
	case agentFiles[base] || hasDir(dir, agentDirs...):
		return Agent
	case strings.HasPrefix(p, ".github/workflows/") || strings.HasPrefix(p, ".github/actions/") ||
		strings.HasPrefix(p, ".github/scripts/") || base == ".gitlab-ci.yml" || base == "jenkinsfile" ||
		base == "azure-pipelines.yml" || base == ".travis.yml" || base == "bitbucket-pipelines.yml" ||
		strings.HasPrefix(p, ".circleci/") || (base == "action.yml" || base == "action.yaml"):
		return CI
	case isTest(p, base, dir):
		return Test
	case docExts[ext] && !strings.HasPrefix(base, "requirements") && !strings.HasPrefix(base, "constraints"):
		return Docs
	case isGenerated(base, dir):
		return Generated
	case buildFiles[base] || strings.HasPrefix(base, "dockerfile") || strings.HasSuffix(base, ".dockerfile") ||
		strings.HasPrefix(base, "docker-compose") || strings.HasPrefix(base, "compose.y") ||
		strings.HasPrefix(base, "install") && (scriptExts[ext] || ext == "") ||
		hasDir(dir, buildDirs...) || ext == ".mk" ||
		(strings.HasPrefix(base, "requirements") || strings.HasPrefix(base, "constraints")) && ext == ".txt" ||
		(strings.HasSuffix(dir, ".vscode/") && (base == "tasks.json" || base == "launch.json" || base == "settings.json")):
		return Build
	case docExts[ext] || strings.HasPrefix(base, "license") || strings.HasPrefix(base, "copying") ||
		strings.HasPrefix(base, "readme") || strings.HasPrefix(base, "changelog") || strings.HasPrefix(dir, "docs/"):
		return Docs
	case scriptExts[ext]:
		return Script
	case configExts[ext]:
		return Config
	case sourceExts[ext]:
		return Source
	}
	return Other
}

func hasDir(dir string, names ...string) bool {
	for _, n := range names {
		if strings.HasPrefix(dir, n) || strings.Contains(dir, "/"+n) {
			return true
		}
	}
	return false
}

func isTest(p, base, dir string) bool {
	if strings.HasSuffix(base, "_test.go") || strings.HasPrefix(base, "test_") && strings.HasSuffix(base, ".py") ||
		strings.HasSuffix(base, "_test.py") || strings.HasSuffix(base, "_spec.rb") || strings.HasSuffix(base, "_test.rs") {
		return true
	}
	for _, marker := range []string{".test.", ".spec.", ".stories.", ".fixture."} {
		if strings.Contains(base, marker) {
			return true
		}
	}
	return hasDir(dir, testDirs...)
}

func isGenerated(base, dir string) bool {
	if strings.Contains(base, ".min.") || strings.HasSuffix(base, ".map") || strings.HasSuffix(base, ".pb.go") ||
		strings.HasSuffix(base, "_generated.go") || strings.HasSuffix(base, ".gen.go") || strings.HasSuffix(base, "_pb2.py") ||
		strings.HasSuffix(base, ".bundle.js") || strings.HasSuffix(base, ".chunk.js") {
		return true
	}
	return hasDir(dir, "dist/", "build/", "out/", ".next/", "target/", "public/build/")
}

// IsCode reports whether role holds code that may execute (as opposed to
// documentation, lockfiles or tests).
func IsCode(role string) bool {
	switch role {
	case Source, Build, Script, CI, Generated, Vendored, Agent:
		return true
	}
	return false
}
