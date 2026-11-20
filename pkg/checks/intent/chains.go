package intent

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/user/safeanalyze/pkg/filerole"
	"github.com/user/safeanalyze/pkg/report"
)

// Behaviour signals. A single signal is a capability; a chain of signals
// close to each other is behaviour.
const (
	SigNet       = "NET"
	SigSecret    = "READ_SECRET"
	SigEnvEnum   = "ENV_ENUM"
	SigDecode    = "DECODE"
	SigEval      = "EVAL"
	SigSpawn     = "SPAWN"
	SigWriteExec = "WRITE_EXEC"
	SigPersist   = "PERSIST"
)

var signalPatterns = map[string]multiPattern{
	SigNet: callSite,
	SigSecret: compileMulti(`(?i)\.ssh/`, `(?i)\bid_(rsa|ed25519|ecdsa|dsa)\b`, `(?i)\.aws/credentials`, `(?i)\.azure/`, `(?i)\.config/gcloud`,
		`(?i)\.kube/config`, `(?i)\.docker/config\.json`, `(?i)\.git-credentials`, `(?i)\.netrc\b`, `(?i)\.pypirc`, `(?i)\.npmrc`,
		`Login Data`, `Local State`, `cookies\.sqlite`, `logins\.json`, `key[34]\.db`, `(?i)keychain`, `/etc/shadow`, `wallet\.dat`,
		`\.bitcoin/`, `\.ethereum/keystore`, `(?i)metamask`, `Exodus`, `(?i)discord[^\n]{0,40}(Local Storage|leveldb)`,
		`Telegram Desktop`, `\btdata\b`, `\.bash_history`, `\.zsh_history`, `AWS_SECRET_ACCESS_KEY`, `169\.254\.169\.254`,
		`metadata\.google\.internal`, `/proc/self/environ`, `(chrome|browser)\.cookies\.getAll`),
	SigEnvEnum: compileMulti(`os\.Environ\(\)`, `JSON\.stringify\(\s*process\.env\s*\)`, `Object\.(keys|entries|values)\(\s*process\.env\s*\)`,
		`os\.environ\.(items|copy|keys|values)\(\)`, `dict\(os\.environ\)`, `json\.dumps\(\s*(dict\()?os\.environ`, `\bprintenv\b`,
		`System\.getenv\(\)\s*[;)]`, `Get-ChildItem\s+env:`, `\benv\s*\|`),
	SigDecode: compileMulti(`base64\.(StdEncoding|URLEncoding|RawStdEncoding|RawURLEncoding)\.Decode`, `hex\.DecodeString`, `\batob\(`,
		`Buffer\.from\([^)]*['"](base64|hex)['"]`, `b64decode`, `a85decode`, `bytes\.fromhex`, `codecs\.decode`, `zlib\.decompress`,
		`gzip\.decompress`, `lzma\.decompress`, `String\.fromCharCode`, `\bbase64\s+(-d|--decode)\b`, `xxd\s+-r`, `FromBase64String`,
		`marshal\.loads`, `unhexlify`),
	SigEval: compileMulti(`(^|[^.\w])eval\(`, `new Function\(`, `(^|[^.\w])Function\(\s*['"]`, `vm\.(runInNewContext|runInThisContext|runInContext|Script)\b`,
		`(^|[^.\w])exec\(\s*[^)]`, `compile\([^)]*['"]exec['"]`, `pickle\.loads`, `Invoke-Expression`, `(^|\s)iex\s`, `\|\s*(ba|z)?sh\b`,
		`\b(ba|z)?sh\s+-c\b`, `plugin\.Open\(`, `__import__\(`, `importlib\.import_module\(`, `setTimeout\(\s*['"]`),
	SigSpawn: compileMulti(`exec\.Command(Context)?\(`, `os\.StartProcess`, `syscall\.(Exec|ForkExec)`, `child_process`, `\bspawn(Sync)?\(`,
		`\bexecSync\(`, `\bexecFile(Sync)?\(`, `subprocess\.`, `os\.system\(`, `os\.popen\(`, `Runtime\.getRuntime\(\)\.exec`,
		`ProcessBuilder\(`, `Start-Process`, `ShellExecute`, `\bpopen\(`, `Command::new\(`),
	SigWriteExec: compileMulti(`os\.Chmod\([^)]*0o?[0-7]?[1357][0-7]{2}\)`, `chmod\s+(\+x|u\+x|a\+x|[0-7]?[1357][0-7]{2})\b`,
		`fs\.chmod(Sync)?\([^)]*0o?[0-7]?[1357][0-7]{2}`, `os\.chmod\([^)]*(0o?[0-7]?[1357][0-7]{2}|S_IEXEC|S_IXUSR)`, `\bS_IXUSR\b`, `set_permissions\(`),
	SigPersist: compileMulti(`(?i)\bcrontab\b`, `(?i)/etc/cron`, `(?i)systemctl\s+(--user\s+)?enable`, `(?i)/etc/systemd/system`,
		`(?i)\.config/systemd/user`, `(?i)LaunchAgents`, `(?i)LaunchDaemons`, `(?i)\.bashrc`, `(?i)\.bash_profile`, `(?i)\.zshrc`,
		`(?i)\.zprofile`, `(?i)CurrentVersion\\{1,2}Run`, `(?i)\bschtasks\b`, `(?i)authorized_keys`, `(?i)/etc/rc\.local`,
		`(?i)ld\.so\.preload`, `(?i)\bLD_PRELOAD\b`, `(?i)\.config/autostart`, `(?i)Startup\\{1,2}`),
}

// envForChild excludes the common, benign pattern of passing the environment
// to a child process.
var envForChild = regexp.MustCompile(`Env\s*[:=]\s*append\(\s*os\.Environ\(\)|env\s*:\s*\{?\s*\.\.\.process\.env|env=os\.environ|\.\.\.process\.env`)

type chainRule struct {
	rule, question, msg string
	a                   string
	b                   []string
}

var chainRules = []chainRule{
	{"chain_secret_exfil", QExfil, "reads credentials/secrets and sends data over the network", SigSecret, []string{SigNet}},
	{"chain_env_exfil", QExfil, "collects the whole environment and sends data over the network", SigEnvEnum, []string{SigNet}},
	{"chain_decode_exec", QRemoteExec, "decodes data and evaluates or executes it (staged loader)", SigDecode, []string{SigEval, SigSpawn}},
	{"chain_download_exec", QRemoteExec, "downloads data and makes it executable or evaluates it (dropper)", SigNet, []string{SigWriteExec, SigEval}},
	{"chain_persistence", QPrivileges, "installs persistence and contacts the network or spawns processes", SigPersist, []string{SigNet, SigSpawn}},
}

// chainWindow is the line distance within which two signals form a chain.
// Co-occurrence instead of data flow keeps recall high.
const chainWindow = 30

func analyzeChains(files []*File) []report.Finding {
	var out []report.Finding
	for _, f := range files {
		if !filerole.IsCode(f.Role) && f.Role != filerole.Test {
			continue
		}
		hits := signalHits(f)
		if len(hits) == 0 {
			continue
		}
		for _, cr := range chainRules {
			a := hits[cr.a]
			if len(a) == 0 {
				continue
			}
			for _, sig := range cr.b {
				la, lb, ok := nearest(a, hits[sig], chainWindow)
				if !ok {
					continue
				}
				first := min(la, lb)
				out = append(out, finding("chains", cr.rule, report.SeverityCritical, report.TierA, cr.question, f, first, f.Lines()[first-1],
					fmt.Sprintf("Code %s (%s line %d + %s line %d)", cr.msg, cr.a, la, sig, lb),
					map[string]string{"signals": cr.a + "+" + sig, "lines": fmt.Sprintf("%d,%d", la, lb)}))
				break
			}
		}
		if l := hits[SigPersist]; len(l) > 0 {
			out = append(out, finding("chains", "capability_persistence", report.SeverityMedium, report.TierB, QPrivileges, f, l[0], f.Lines()[l[0]-1],
				"Code references a persistence location (cron, systemd, shell rc, autostart, authorized_keys, LD_PRELOAD)", nil))
		}
		if l := hits[SigSecret]; len(l) > 0 {
			out = append(out, finding("chains", "capability_secret_access", report.SeverityMedium, report.TierB, QExfil, f, l[0], f.Lines()[l[0]-1],
				"Code references a credential store or secret location", nil))
		}
		var caps []string
		for _, sig := range []string{SigNet, SigSpawn, SigEval, SigDecode, SigEnvEnum, SigWriteExec} {
			if len(hits[sig]) > 0 {
				caps = append(caps, fmt.Sprintf("%s×%d", sig, len(hits[sig])))
			}
		}
		if len(caps) > 0 {
			sort.Strings(caps)
			out = append(out, finding("chains", "capability_inventory", report.SeverityInfo, report.TierC, "", f, 1, strings.Join(caps, " "),
				"Capabilities used in this file: "+strings.Join(caps, ", "), map[string]string{"capabilities": strings.Join(caps, " ")}))
		}
	}
	return out
}

// commentLine matches lines that are entirely a comment.
var commentLine = regexp.MustCompile(`^\s*(//|#|/\*|\*|--|;|<!--|REM\b|')`)

// regexLiteral matches lines that define a detection pattern rather than run
// the code the pattern describes (security tools, linters, this scanner).
var regexLiteral = regexp.MustCompile(`\(\?[imsU:]|\[\^|\\[sdwbSDWB][*+?{}()|\]\\]|(^|[^\\])\\\.[a-z]`)

// CodeLine reports whether a line can execute: not a comment and not a
// regular-expression definition.
func CodeLine(l string) bool {
	return !commentLine.MatchString(l) && !regexLiteral.MatchString(l)
}

var stringLiteral = regexp.MustCompile(`"(?:[^"\\\n]|\\.)*"|'(?:[^'\\\n]|\\.)*'|` + "`[^`\n]*`")
var shellSyntax = regexp.MustCompile(`\|\s|\$\(|&&|\s>\s|\s-[a-zA-Z]\s`)
var plainWord = regexp.MustCompile(`^[(\[]?[A-Za-z][A-Za-z'-]*[,.;:!?)\]]*$`)

// InProse reports whether byte offset pos of line lies inside a string
// literal that reads as natural language (a description, message or prompt):
// at least six words, at least half of them plain words, and no shell
// syntax. Such text names a behaviour instead of performing it.
func InProse(line string, pos int) bool {
	for _, loc := range stringLiteral.FindAllStringIndex(line, -1) {
		if pos < loc[0] || pos >= loc[1] {
			continue
		}
		body := line[loc[0]+1 : loc[1]-1]
		if shellSyntax.MatchString(body) {
			return false // a command line, even a wordy one, can run
		}
		tokens := strings.Fields(body)
		if len(tokens) < 6 {
			return false
		}
		plain := 0
		for _, t := range tokens {
			if plainWord.MatchString(t) {
				plain++
			}
		}
		return plain*2 >= len(tokens)
	}
	return false
}

// signalHits returns, per signal, the 1-based lines where it occurs.
// Comment lines, regex definitions and matches inside prose strings are
// skipped: they cannot execute.
func signalHits(f *File) map[string][]int {
	hits := map[string][]int{}
	// Whole-file prefilter: only signals present somewhere in the file are
	// searched line by line.
	present := map[string]multiPattern{}
	for sig, m := range signalPatterns {
		if sub := m.forText(f.Content); len(sub) > 0 {
			present[sig] = sub
		}
	}
	if len(present) == 0 {
		return hits
	}
	for i, l := range f.Lines() {
		if len(l) > 2000 {
			l = l[:2000] // minified lines: bound regex cost
		}
		if !CodeLine(l) {
			continue
		}
		for sig, re := range present {
			loc := re.FindStringIndex(l)
			if loc == nil || InProse(l, loc[0]) {
				continue
			}
			if sig == SigEnvEnum && envForChild.MatchString(l) {
				continue
			}
			hits[sig] = append(hits[sig], i+1)
		}
	}
	return hits
}

// nearest returns the closest pair of lines from the ascending slices a and
// b, and whether they are within window of each other.
func nearest(a, b []int, window int) (int, int, bool) {
	best, ba, bb := window+1, 0, 0
	for i, j := 0, 0; i < len(a) && j < len(b); {
		d := a[i] - b[j]
		if d < 0 {
			d = -d
		}
		if d < best {
			best, ba, bb = d, a[i], b[j]
		}
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return ba, bb, best <= window
}
