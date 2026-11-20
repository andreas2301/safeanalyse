package intent

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/user/safeanalyze/pkg/filerole"
	"github.com/user/safeanalyze/pkg/report"
)

// npmLifecycle are package.json scripts npm runs without being asked:
// on install of the package or of a git dependency.
// npmInstallHooks are the ones that also run for a registry dependency.
var npmInstallHooks = map[string]bool{"preinstall": true, "install": true, "postinstall": true}

var npmLifecycle = []string{"preinstall", "install", "postinstall", "prepare", "preprepare", "postprepare", "prepublish", "prepack", "postpack", "dependencies"}

// goInitDanger matches calls that should not run at program start or import.
var goInitDanger = regexp.MustCompile(`exec\.Command|os\.StartProcess|syscall\.(Exec|ForkExec)|http\.(Get|Post|NewRequest|DefaultClient)|net\.Dial|\.Do\(req|os\.(WriteFile|Create|Remove|RemoveAll)\(|ioutil\.WriteFile|base64\.\w+\.DecodeString|plugin\.Open`)

var goGenerate = regexp.MustCompile(`^\s*//go:generate\s+(.*)`)
var goLinkname = regexp.MustCompile(`^\s*//go:linkname\b`)
var goCgo = regexp.MustCompile(`^\s*import\s+"C"`)
var goInitFunc = regexp.MustCompile(`^func\s+init\s*\(\s*\)\s*\{`)
var goGlobalCall = regexp.MustCompile(`^var\s+\w+\s*=\s*func\s*\(.*\)\s*\w*\s*\{`)

var pySetupDanger = regexp.MustCompile(`subprocess|os\.system|os\.popen|urllib|requests|socket|http\.client|exec\(|eval\(|base64|__import__|cmdclass`)
var pthImport = regexp.MustCompile(`^\s*import\s`)

var rustDanger = regexp.MustCompile(`Command::new|reqwest|TcpStream|ureq|std::net|curl`)

var devcontainerCmd = regexp.MustCompile(`"(initializeCommand|onCreateCommand|updateContentCommand|postCreateCommand|postStartCommand|postAttachCommand)"\s*:\s*(.*)`)

var vscodeFolderOpen = regexp.MustCompile(`"runOn"\s*:\s*"folderOpen"`)
var vscodeSettingsExec = regexp.MustCompile(`"([\w.]*(executablePath|interpreterPath|\.path|Path|shellArgs[\w.]*|\.env\.(linux|osx|windows)|terminal\.integrated\.profiles[\w.]*))"\s*:`)

var gradleDistribution = regexp.MustCompile(`distributionUrl\s*=\s*(\S+)`)

func analyzeAutorun(files []*File) []report.Finding {
	var out []report.Finding
	for _, f := range files {
		if f.Role == filerole.Vendored && f.Base() != "package.json" {
			continue
		}
		base := f.Base()
		rel := strings.ToLower(f.Rel)
		switch {
		case base == "package.json":
			out = append(out, npmScripts(f)...)
		case base == "setup.py":
			out = append(out, pythonSetup(f)...)
		case strings.HasSuffix(base, ".pth"):
			if line, l := firstMatchLine(f, pthImport); line > 0 {
				out = append(out, finding("autorun", "autorun_python_pth", report.SeverityCritical, report.TierA, QAutorun, f, line, l,
					".pth file with an import line runs code at every Python interpreter start", map[string]string{"trigger": "python start"}))
			}
		case strings.HasSuffix(base, ".go"):
			out = append(out, goAutorun(f)...)
		case base == "build.rs":
			tier := report.TierB
			if rustDanger.MatchString(f.Content) {
				tier = report.TierA
			}
			out = append(out, finding("autorun", "autorun_cargo_build_script", report.SeverityHigh, tier, QAutorun, f, 1, "build.rs",
				"Cargo build script runs on every build"+dangerSuffix(tier), map[string]string{"trigger": "cargo build"}))
		case strings.Contains(rel, ".devcontainer/") && strings.HasSuffix(base, ".json"):
			out = append(out, devcontainer(f)...)
		case strings.HasSuffix(rel, ".vscode/tasks.json"):
			if line, l := firstMatchLine(f, vscodeFolderOpen); line > 0 {
				out = append(out, finding("autorun", "autorun_vscode_folder_open", report.SeverityCritical, report.TierA, QAutorun, f, line, l,
					"VS Code task runs automatically when the folder is opened", map[string]string{"trigger": "open folder in VS Code"}))
			}
		case strings.HasSuffix(rel, ".vscode/settings.json"):
			if line, l := firstMatchLine(f, vscodeSettingsExec); line > 0 {
				out = append(out, finding("autorun", "autorun_vscode_settings_exec", report.SeverityMedium, report.TierB, QAutorun, f, line, l,
					"Workspace settings point an executable, interpreter or terminal environment at a repository-controlled value", map[string]string{"trigger": "open folder in VS Code"}))
			}
		case strings.Contains(rel, ".husky/") || strings.Contains(rel, ".githooks/") || base == "lefthook.yml" || base == ".pre-commit-config.yaml":
			tier := report.TierB
			if riskyCommand.MatchString(f.Content) && base != ".pre-commit-config.yaml" {
				tier = report.TierA
			}
			out = append(out, finding("autorun", "autorun_git_hook", report.SeverityMedium, tier, QAutorun, f, 1, f.Rel,
				"Git hook configuration runs commands on commit/push once installed"+dangerSuffix(tier), map[string]string{"trigger": "git commit/push"}))
		case base == ".envrc":
			tier := report.TierB
			if riskyCommand.MatchString(f.Content) {
				tier = report.TierA
			}
			out = append(out, finding("autorun", "autorun_direnv", report.SeverityMedium, tier, QAutorun, f, 1, ".envrc",
				"direnv runs .envrc when entering the directory (after `direnv allow`)"+dangerSuffix(tier), map[string]string{"trigger": "cd into directory (direnv)"}))
		case base == "gradle-wrapper.properties":
			if line, l := firstMatchLine(f, gradleDistribution); line > 0 && !strings.Contains(l, "services.gradle.org") {
				out = append(out, finding("autorun", "autorun_gradle_wrapper_source", report.SeverityHigh, report.TierA, QAutorun, f, line, l,
					"Gradle wrapper downloads its distribution from a non-official host", map[string]string{"trigger": "./gradlew"}))
			}
		}
	}
	return out
}

func dangerSuffix(tier string) string {
	if tier == report.TierA {
		return "; it downloads, decodes, evaluates or contacts the network"
	}
	return ""
}

func npmScripts(f *File) []report.Finding {
	var pkg struct {
		Scripts        map[string]string `json:"scripts"`
		ActivationEvts []string          `json:"activationEvents"`
		Engines        map[string]string `json:"engines"`
	}
	if err := json.Unmarshal([]byte(f.Content), &pkg); err != nil {
		return nil
	}
	var out []report.Finding
	for _, name := range npmLifecycle {
		cmd, ok := pkg.Scripts[name]
		if !ok || f.Dep && !npmInstallHooks[name] {
			continue
		}
		tier, sev := report.TierB, report.SeverityMedium
		if riskyCommand.MatchString(cmd) || strings.Contains(cmd, "node ") && !strings.Contains(cmd, "node_modules/.bin") {
			tier, sev = report.TierA, report.SeverityHigh
		}
		line := lineOfKey(f, `"`+name+`"`)
		out = append(out, finding("autorun", "autorun_npm_lifecycle", sev, tier, QAutorun, f, line, fmt.Sprintf("%q: %q", name, cmd),
			fmt.Sprintf("npm runs the %q script automatically on install", name)+dangerSuffix(tier), map[string]string{"trigger": "npm install"}))
	}
	if _, isVSCode := pkg.Engines["vscode"]; isVSCode {
		for _, ev := range pkg.ActivationEvts {
			if ev == "*" || ev == "onStartupFinished" {
				out = append(out, finding("autorun", "autorun_vscode_extension_startup", report.SeverityMedium, report.TierB, QAutorun, f, lineOfKey(f, `"activationEvents"`), ev,
					"VS Code extension activates on every editor start", map[string]string{"trigger": "editor start"}))
			}
		}
	}
	return out
}

func pythonSetup(f *File) []report.Finding {
	tier, sev, msg := report.TierB, report.SeverityMedium, "setup.py runs arbitrary code on install/build"
	line, l := firstMatchLine(f, pySetupDanger)
	if line > 0 {
		tier, sev, msg = report.TierA, report.SeverityHigh, "setup.py runs on install and spawns processes, decodes, evaluates or contacts the network"
	} else {
		line, l = 1, "setup.py"
	}
	return []report.Finding{finding("autorun", "autorun_python_setup", sev, tier, QAutorun, f, line, l, msg, map[string]string{"trigger": "pip install"})}
}

// goAutorun reports init() functions and package-level initialisers that run
// dangerous calls at import time, go:generate directives, linkname and cgo.
func goAutorun(f *File) []report.Finding {
	var out []report.Finding
	lines := f.Lines()
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		if m := goGenerate.FindStringSubmatch(l); m != nil {
			tier := report.TierB
			if riskyCommand.MatchString(m[1]) {
				tier = report.TierA
			}
			out = append(out, finding("autorun", "autorun_go_generate", report.SeverityMedium, tier, QAutorun, f, i+1, l,
				"go:generate runs this command on `go generate`"+dangerSuffix(tier), map[string]string{"trigger": "go generate"}))
			continue
		}
		if goLinkname.MatchString(l) {
			out = append(out, finding("autorun", "go_linkname", report.SeverityLow, report.TierC, QHidden, f, i+1, l,
				"go:linkname reaches into unexported runtime/package symbols", nil))
			continue
		}
		if goCgo.MatchString(l) {
			out = append(out, finding("autorun", "go_cgo", report.SeverityLow, report.TierC, QAutorun, f, i+1, l,
				"cgo: C code is compiled and linked on build", map[string]string{"trigger": "go build"}))
			continue
		}
		isInit := goInitFunc.MatchString(l)
		if !isInit && !goGlobalCall.MatchString(l) {
			continue
		}
		end := blockEnd(lines, i)
		body := strings.Join(lines[i:end+1], "\n")
		if !isInit && !strings.HasSuffix(strings.TrimSpace(lines[end]), "}()") {
			i = end
			continue // a function value, not an immediately invoked initialiser
		}
		if loc := goInitDanger.FindStringIndex(body); loc != nil {
			what := "init()"
			if !isInit {
				what = "package-level initialiser"
			}
			hit := lineOf(body, loc[0]) + i
			out = append(out, finding("autorun", "autorun_go_init", report.SeverityCritical, report.TierA, QAutorun, f, hit, lines[hit-1],
				what+" runs process, network, file or decode operations at program start / package import", map[string]string{"trigger": "program start / import"}))
		}
		i = end
	}
	return out
}

// blockEnd returns the index of the line that closes the brace block opened
// on line start. Braces inside strings are not special-cased; this is a
// recall-oriented approximation.
func blockEnd(lines []string, start int) int {
	depth := 0
	for i := start; i < len(lines); i++ {
		depth += strings.Count(lines[i], "{") - strings.Count(lines[i], "}")
		if depth <= 0 && i > start || depth <= 0 && strings.Contains(lines[i], "}") {
			return i
		}
	}
	return len(lines) - 1
}

func devcontainer(f *File) []report.Finding {
	var out []report.Finding
	for i, l := range f.Lines() {
		m := devcontainerCmd.FindStringSubmatch(l)
		if m == nil {
			continue
		}
		tier, sev := report.TierB, report.SeverityMedium
		msg := fmt.Sprintf("Dev container %s runs automatically", m[1])
		if m[1] == "initializeCommand" {
			tier, sev = report.TierA, report.SeverityHigh
			msg = "Dev container initializeCommand runs on the HOST before the container starts"
		} else if riskyCommand.MatchString(m[2]) {
			tier, sev = report.TierA, report.SeverityHigh
		}
		out = append(out, finding("autorun", "autorun_devcontainer", sev, tier, QAutorun, f, i+1, l, msg+dangerSuffix(tier),
			map[string]string{"trigger": "open in dev container"}))
	}
	return out
}

func lineOfKey(f *File, key string) int {
	if i := strings.Index(f.Content, key); i >= 0 {
		return lineOf(f.Content, i)
	}
	return 1
}
