package intent

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/user/safeanalyze/pkg/report"
)

var officialNPMHosts = map[string]bool{"registry.npmjs.org": true, "registry.yarnpkg.com": true, "registry.npmjs.com": true}

var goModReplace = regexp.MustCompile(`^\s*(?:replace\s+)?([^\s()]+)(?:\s+v[^\s]+)?\s+=>\s+([^\s]+)(?:\s+(v[^\s]+))?`)
var goModRequire = regexp.MustCompile(`^\s*(?:require\s+)?([a-zA-Z0-9][^\s()]*\.[^\s()]+)\s+(v[0-9][^\s]*)(\s*//\s*indirect)?`)

var goEnvRisk = regexp.MustCompile(`\b(GONOSUMDB|GONOSUMCHECK|GOINSECURE|GONOSUMDB|GOFLAGS\s*=\s*["']?[^"'\n]*-insecure|GOPROXY\s*=\s*["']?(direct|off|https?://(?:[^,"'\s]*)))`)
var goSumDBOff = regexp.MustCompile(`GOSUMDB\s*=\s*["']?off|GONOSUMCHECK\s*=\s*1|GOFLAGS\s*=\s*["']?[^"'\n]*-insecure`)

var npmrcRegistry = regexp.MustCompile(`(?im)^\s*(@[\w-]+:)?registry\s*=\s*(\S+)|npmRegistryServer:\s*["']?(\S+?)["']?\s*$`)
var npmrcInsecure = regexp.MustCompile(`(?im)^\s*strict-ssl\s*=\s*false|enableStrictSsl:\s*false|unsafe-perm\s*=\s*true`)

var pipIndex = regexp.MustCompile(`(?i)(--extra-index-url|--index-url|-i)\s+(\S+)|(extra-index-url|index-url)\s*=\s*(\S+)|--trusted-host\s+(\S+)|trusted-host\s*=\s*(\S+)`)
var pipDirectURL = regexp.MustCompile(`(?i)(git\+https?://|git\+ssh://|git\+git://|https?://\S+\.(tar\.gz|zip|whl))`)

var cargoSource = regexp.MustCompile(`(?m)^\s*(replace-with|registry)\s*=\s*"([^"]+)"|^\s*\[source\.([^\]]+)\]`)
var gemSource = regexp.MustCompile(`(?m)^\s*source\s+["']([^"']+)["']`)

func analyzeDepSource(files []*File) []report.Finding {
	var out []report.Finding
	for _, f := range files {
		base := f.Base()
		switch {
		case base == "go.mod":
			out = append(out, goMod(f, files)...)
		case base == "package-lock.json" || base == "npm-shrinkwrap.json":
			out = append(out, npmLock(f)...)
		case base == "yarn.lock" || base == "pnpm-lock.yaml":
			out = append(out, textLockHosts(f)...)
		case base == "package.json":
			out = append(out, npmManifest(f)...)
		case base == ".npmrc" || base == ".yarnrc" || base == ".yarnrc.yml":
			out = append(out, npmrc(f)...)
		case strings.HasPrefix(base, "requirements") && strings.HasSuffix(base, ".txt") || base == "pip.conf" || base == "pip.ini" || base == "pyproject.toml" || base == "setup.cfg" || base == "pipfile":
			out = append(out, pipSources(f)...)
		case strings.HasSuffix(strings.ToLower(f.Rel), ".cargo/config.toml") || strings.HasSuffix(strings.ToLower(f.Rel), ".cargo/config"):
			if line, l := firstMatchLine(f, cargoSource); line > 0 {
				out = append(out, finding("depsource", "dep_cargo_source_replacement", report.SeverityMedium, report.TierB, QDeps, f, line, l,
					"Cargo source replacement or alternative registry", nil))
			}
		case base == "gemfile":
			for i, l := range f.Lines() {
				if m := gemSource.FindStringSubmatch(l); m != nil && !strings.Contains(m[1], "rubygems.org") {
					out = append(out, finding("depsource", "dep_gem_source", report.SeverityMedium, report.TierB, QDeps, f, i+1, l,
						"Gem source other than rubygems.org", map[string]string{"host": m[1]}))
				}
			}
		}
		if f.Role != "lockfile" {
			if line, l := firstMatchLine(f, goSumDBOff); line > 0 {
				out = append(out, finding("depsource", "dep_go_checksum_disabled", report.SeverityHigh, report.TierA, QDeps, f, line, l,
					"Go module checksum verification is disabled or insecure fetching is allowed", nil))
			} else if line, l := firstMatchLine(f, goEnvRisk); line > 0 {
				out = append(out, finding("depsource", "dep_go_env_override", report.SeverityLow, report.TierB, QDeps, f, line, l,
					"Go module proxy / checksum database settings are overridden", nil))
			}
		}
	}
	return out
}

func goMod(f *File, files []*File) []report.Finding {
	var out []report.Finding
	sum := ""
	dir := strings.TrimSuffix(f.Rel, "go.mod")
	for _, g := range files {
		if g.Rel == dir+"go.sum" {
			sum = g.Content
		}
	}
	inReplace, inRequire := false, false
	var requires []string
	replaced := map[string]string{} // module -> "path version" of its replacement, or "" for a local path
	for i, l := range f.Lines() {
		t := strings.TrimSpace(l)
		switch {
		case strings.HasPrefix(t, "replace ("):
			inReplace = true
			continue
		case strings.HasPrefix(t, "require ("):
			inRequire = true
			continue
		case t == ")":
			inReplace, inRequire = false, false
			continue
		}
		if inReplace || strings.HasPrefix(t, "replace ") {
			m := goModReplace.FindStringSubmatch(t)
			if m == nil {
				continue
			}
			from, to := m[1], m[2]
			replaced[from] = strings.TrimSpace(to + " " + m[3])
			if strings.HasPrefix(to, ".") || strings.HasPrefix(to, "/") {
				out = append(out, finding("depsource", "dep_go_replace_local", report.SeverityMedium, report.TierB, QDeps, f, i+1, l,
					fmt.Sprintf("Module %s is replaced by local directory %s", from, to), map[string]string{"module": from}))
			} else if owner(from) != owner(to) {
				out = append(out, finding("depsource", "dep_go_replace_fork", report.SeverityHigh, report.TierB, QDeps, f, i+1, l,
					fmt.Sprintf("Module %s is replaced by %s (different owner): the code built is not the upstream module", from, to),
					map[string]string{"module": from, "replacement": to}))
			}
			continue
		}
		if inRequire || strings.HasPrefix(t, "require ") {
			if m := goModRequire.FindStringSubmatch(t); m != nil {
				requires = append(requires, m[1]+" "+m[2])
			}
		}
	}
	if sum == "" && len(requires) > 0 {
		out = append(out, finding("depsource", "dep_go_sum_missing", report.SeverityHigh, report.TierB, QDeps, f, 1, "go.mod",
			"go.mod has requirements but no go.sum: dependency content is not pinned by checksum", nil))
	} else {
		for _, r := range requires {
			// A replaced module is verified under its replacement's path.
			if rep, ok := replaced[strings.Fields(r)[0]]; ok {
				if !strings.Contains(rep, " ") {
					continue // local directory: not in go.sum
				}
				r = rep
			}
			if !strings.Contains(sum, r+" ") && !strings.Contains(sum, r+"/go.mod ") {
				out = append(out, finding("depsource", "dep_go_sum_entry_missing", report.SeverityMedium, report.TierB, QDeps, f, lineOfKey(f, strings.Fields(r)[0]), r,
					"Required module has no go.sum entry", nil))
			}
		}
	}
	return out
}

// owner returns the first two path elements of a module path (host/owner).
func owner(mod string) string {
	parts := strings.Split(strings.ToLower(mod), "/")
	if len(parts) >= 2 {
		return parts[0] + "/" + parts[1]
	}
	return strings.ToLower(mod)
}

func npmLock(f *File) []report.Finding {
	var lock struct {
		Packages     map[string]lockPackage `json:"packages"`
		Dependencies map[string]lockPackage `json:"dependencies"`
	}
	if err := json.Unmarshal([]byte(f.Content), &lock); err != nil {
		return nil
	}
	pkgs := lock.Packages
	if len(pkgs) == 0 {
		pkgs = lock.Dependencies
	}
	type hostStat struct {
		count   int
		example string
	}
	hosts := map[string]*hostStat{}
	var out []report.Finding
	noIntegrity := 0
	for name, p := range pkgs {
		if name == "" || p.Resolved == "" {
			continue
		}
		r := p.Resolved
		switch {
		case strings.HasPrefix(r, "git+") || strings.HasPrefix(r, "git:") || strings.HasPrefix(r, "github:"):
			out = append(out, finding("depsource", "dep_npm_git_source", report.SeverityMedium, report.TierB, QDeps, f, lineOfKey(f, r), r,
				fmt.Sprintf("%s is installed from a git source (no registry malware takedown, mutable unless pinned to a commit)", name), nil))
			continue
		case strings.HasPrefix(r, "file:") || strings.HasPrefix(r, "link:"):
			continue
		case strings.HasPrefix(r, "http://"):
			out = append(out, finding("depsource", "dep_npm_plain_http", report.SeverityHigh, report.TierA, QDeps, f, lineOfKey(f, r), r,
				fmt.Sprintf("%s is downloaded over plain HTTP", name), nil))
		}
		if p.Integrity == "" {
			noIntegrity++
		}
		u, err := url.Parse(r)
		if err != nil || u.Host == "" {
			continue
		}
		host := strings.ToLower(u.Hostname())
		if officialNPMHosts[host] {
			continue
		}
		hs := hosts[host]
		if hs == nil {
			hs = &hostStat{example: r}
			hosts[host] = hs
		}
		hs.count++
	}
	names := make([]string, 0, len(hosts))
	for h := range hosts {
		names = append(names, h)
	}
	sort.Strings(names)
	for _, h := range names {
		hs := hosts[h]
		what := "a non-official registry"
		if isMirrorRegistry(h) {
			what = "a third-party registry mirror"
		}
		out = append(out, finding("depsource", "dep_npm_registry", report.SeverityHigh, report.TierB, QDeps, f, lineOfKey(f, hs.example), hs.example,
			fmt.Sprintf("%d package(s) are resolved from %s (%s) instead of registry.npmjs.org", hs.count, h, what), map[string]string{"host": h, "count": fmt.Sprint(hs.count)}))
	}
	if noIntegrity > 0 {
		out = append(out, finding("depsource", "dep_npm_missing_integrity", report.SeverityMedium, report.TierB, QDeps, f, 1, f.Rel,
			fmt.Sprintf("%d remote package(s) have no integrity hash", noIntegrity), nil))
	}
	return out
}

type lockPackage struct {
	Version   string `json:"version"`
	Resolved  string `json:"resolved"`
	Integrity string `json:"integrity"`
	License   any    `json:"license"`
	Dev       bool   `json:"dev"`
}

var lockResolved = regexp.MustCompile(`(?:resolved|tarball):?\s+["']?(https?://[^"'\s,}]+)`)

func textLockHosts(f *File) []report.Finding {
	counts := map[string]int{}
	examples := map[string]int{}
	for i, l := range f.Lines() {
		m := lockResolved.FindStringSubmatch(l)
		if m == nil {
			continue
		}
		u, err := url.Parse(m[1])
		if err != nil {
			continue
		}
		h := strings.ToLower(u.Hostname())
		if officialNPMHosts[h] {
			continue
		}
		if counts[h] == 0 {
			examples[h] = i + 1
		}
		counts[h]++
	}
	var out []report.Finding
	for h, n := range counts {
		out = append(out, finding("depsource", "dep_npm_registry", report.SeverityHigh, report.TierB, QDeps, f, examples[h], f.Lines()[examples[h]-1],
			fmt.Sprintf("%d package(s) are resolved from %s instead of the official registry", n, h), map[string]string{"host": h, "count": fmt.Sprint(n)}))
	}
	return out
}

func npmManifest(f *File) []report.Finding {
	var pkg map[string]json.RawMessage
	if err := json.Unmarshal([]byte(f.Content), &pkg); err != nil {
		return nil
	}
	var out []report.Finding
	for _, section := range []string{"dependencies", "devDependencies", "optionalDependencies", "peerDependencies"} {
		var deps map[string]string
		if raw, ok := pkg[section]; !ok || json.Unmarshal(raw, &deps) != nil {
			continue
		}
		for name, spec := range deps {
			var rule, msg string
			tier := report.TierB
			switch {
			case strings.HasPrefix(spec, "npm:"):
				alias := strings.TrimPrefix(spec, "npm:")
				if i := strings.LastIndex(alias, "@"); i > 0 {
					alias = alias[:i]
				}
				if alias == name {
					continue
				}
				rule, msg = "dep_npm_alias", fmt.Sprintf("%q is an alias that installs a different package (%s)", name, spec)
			case strings.HasPrefix(spec, "git") || strings.HasPrefix(spec, "github:") || strings.Contains(spec, "/") && !strings.HasPrefix(spec, "file:") && !strings.HasPrefix(spec, "link:") && !strings.HasPrefix(spec, "workspace:"):
				rule, msg = "dep_npm_url_source", fmt.Sprintf("%q is installed from %s instead of the registry", name, spec)
				if strings.HasPrefix(spec, "http://") {
					tier = report.TierA
				}
			default:
				continue
			}
			out = append(out, finding("depsource", rule, report.SeverityMedium, tier, QDeps, f, lineOfKey(f, `"`+name+`"`), name+": "+spec, msg, nil))
		}
	}
	return out
}

func npmrc(f *File) []report.Finding {
	var out []report.Finding
	for i, l := range f.Lines() {
		if m := npmrcRegistry.FindStringSubmatch(l); m != nil {
			reg := m[2]
			if reg == "" {
				reg = m[3]
			}
			u, err := url.Parse(reg)
			if err == nil && officialNPMHosts[strings.ToLower(u.Hostname())] {
				continue
			}
			out = append(out, finding("depsource", "dep_npm_registry_config", report.SeverityHigh, report.TierB, QDeps, f, i+1, l,
				"Package manager is configured to use a non-official registry", map[string]string{"host": reg}))
		}
		if npmrcInsecure.MatchString(l) {
			out = append(out, finding("depsource", "dep_npm_insecure_config", report.SeverityHigh, report.TierA, QDeps, f, i+1, l,
				"Package manager TLS verification is disabled or scripts run with elevated rights", nil))
		}
	}
	return out
}

func pipSources(f *File) []report.Finding {
	var out []report.Finding
	for i, l := range f.Lines() {
		if m := pipIndex.FindStringSubmatch(l); m != nil {
			tier, sev := report.TierB, report.SeverityMedium
			msg := "pip uses a custom package index"
			lower := strings.ToLower(l)
			switch {
			case strings.Contains(lower, "extra-index-url"):
				tier, sev, msg = report.TierA, report.SeverityHigh, "pip --extra-index-url mixes a second index with PyPI: dependency-confusion risk"
			case strings.Contains(lower, "trusted-host"):
				tier, sev, msg = report.TierA, report.SeverityHigh, "pip --trusted-host disables TLS verification for a host"
			}
			out = append(out, finding("depsource", "dep_pip_index", sev, tier, QDeps, f, i+1, l, msg, nil))
			continue
		}
		if pipDirectURL.MatchString(l) {
			out = append(out, finding("depsource", "dep_pip_direct_url", report.SeverityMedium, report.TierB, QDeps, f, i+1, l,
				"Python dependency installed from a git or archive URL instead of PyPI", nil))
		}
	}
	return out
}
