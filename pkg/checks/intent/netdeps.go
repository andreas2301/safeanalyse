package intent

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/user/safeanalyze/pkg/checks/yara"
	"github.com/user/safeanalyze/pkg/filerole"
	"github.com/user/safeanalyze/pkg/report"
)

// NetDepsOptions configures the network dependency stage.
type NetDepsOptions struct {
	// MaxPackages caps the number of dependencies fetched; the rest are
	// recorded as skipped in the report, never silently dropped.
	MaxPackages int
	// Timeout bounds the whole stage.
	Timeout time.Duration
	// GoBinary is the go command used to download modules ("go" if empty).
	GoBinary string
}

// NetDepsStage downloads the direct dependencies that are not vendored in
// the repository (Go modules, npm packages) and runs the intent analyzers on
// their source. Nothing is built, installed or executed: npm tarballs are
// read in memory, Go modules are fetched with `go mod download`, which only
// downloads and verifies checksums.
type NetDepsStage struct {
	opts     NetDepsOptions
	excluded []string
}

// NewNetDepsStage returns the network dependency stage.
func NewNetDepsStage(opts NetDepsOptions, excluded []string) *NetDepsStage {
	if opts.MaxPackages <= 0 {
		opts.MaxPackages = 300
	}
	if opts.Timeout <= 0 {
		opts.Timeout = 10 * time.Minute
	}
	if opts.GoBinary == "" {
		opts.GoBinary = "go"
	}
	return &NetDepsStage{opts: opts, excluded: excluded}
}

// Name implements pipeline.Stage.
func (s *NetDepsStage) Name() string { return "netdeps" }

// Dependencies implements pipeline.Stage.
func (s *NetDepsStage) Dependencies() []string { return nil }

// depAnalyzers run on dependency source. depsource, ci and agentsurface are
// left out: a dependency's go.mod replace directives, CI and agent files do
// not affect consumers.
var depAnalyzers = []func([]*File) []report.Finding{analyzeEgress, analyzeAutorun, analyzeChains, analyzeInstallScripts}

// Run implements pipeline.Stage.
func (s *NetDepsStage) Run(ctx context.Context, target string, _ *report.Report) (*report.Report, error) {
	rep := report.NewReport(target)
	ctx, cancel := context.WithTimeout(ctx, s.opts.Timeout)
	defer cancel()

	files, err := loadCached(target, s.excluded)
	if err != nil {
		return nil, err
	}
	budget := s.opts.MaxPackages
	var deps []*depSource
	goDeps, skipped := goModuleDeps(files, budget)
	budget -= len(goDeps)
	npmDeps, skippedNPM := npmDeps(files, budget)
	skipped += skippedNPM
	if skipped > 0 {
		rep.AddError(s.Name(), fmt.Sprintf("%d dependencies not fetched: max_packages=%d reached", skipped, s.opts.MaxPackages))
	}

	if len(goDeps) > 0 {
		got, err := s.fetchGo(ctx, goDeps)
		if err != nil {
			rep.AddError(s.Name(), err.Error())
		}
		deps = append(deps, got...)
	}
	if len(npmDeps) > 0 {
		got, errs := fetchNPM(ctx, npmDeps)
		for _, e := range errs {
			rep.AddError(s.Name(), e)
		}
		deps = append(deps, got...)
	}

	scanned := 0
	for _, d := range deps {
		for _, f := range d.findings {
			rep.AddFinding(f)
		}
		if d.files == nil {
			continue
		}
		scanned += len(d.files)
		for _, analyze := range depAnalyzers {
			for _, f := range analyze(d.files) {
				if f.Attributes == nil {
					f.Attributes = map[string]string{}
				}
				f.Attributes["dependency"] = d.id
				// Endpoint inventory of third-party code is noise unless it
				// is exfil-, IP- or telemetry-class.
				if f.Tier == report.TierB && (f.RuleID == "egress_unknown" || f.RuleID == "egress_api" || f.RuleID == "egress_infra") {
					f.Tier = report.TierC
				}
				f.Message = "[dependency " + d.id + "] " + f.Message
				rep.AddFinding(f)
			}
		}
		for _, f := range d.files {
			inner := strings.TrimPrefix(f.Rel, d.prefix)
			if !strings.Contains(inner, "/") && (strings.HasPrefix(f.Base(), "license") || strings.HasPrefix(f.Base(), "copying") || strings.HasPrefix(f.Base(), "licence")) {
				lf := licenseFinding(f, f.Content, "License of dependency "+d.id)
				if lf.Tier != report.TierC {
					rep.AddFinding(lf)
				}
			}
		}
	}
	rep.Metadata["netdeps_fetched"] = fmt.Sprint(len(deps))
	rep.Summary.FilesScanned = scanned
	rep.Finish()
	return rep, nil
}

type depRef struct {
	ecosystem, name, version, resolved, integrity string
}

type depSource struct {
	id, prefix string
	files      []*File
	findings   []report.Finding
}

// goModuleDeps lists the direct requirements of every go.mod, with replace
// directives applied (the replacement is what gets built).
func goModuleDeps(files []*File, budget int) ([]depRef, int) {
	seen := map[string]bool{}
	var deps []depRef
	skipped := 0
	for _, f := range files {
		if f.Base() != "go.mod" || f.Role == filerole.Vendored || f.Role == filerole.Test {
			continue
		}
		replace := map[string][2]string{}
		var reqs [][2]string
		inReplace, inRequire := false, false
		for _, l := range f.Lines() {
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
				if m := goModReplace.FindStringSubmatch(t); m != nil && m[3] != "" {
					replace[m[1]] = [2]string{m[2], m[3]}
				} else if m != nil {
					replace[m[1]] = [2]string{"", ""} // local path: nothing to fetch
				}
				continue
			}
			if inRequire || strings.HasPrefix(t, "require ") {
				if m := goModRequire.FindStringSubmatch(t); m != nil && m[3] == "" {
					reqs = append(reqs, [2]string{m[1], m[2]})
				}
			}
		}
		for _, r := range reqs {
			mod, ver := r[0], r[1]
			if rep, ok := replace[mod]; ok {
				if rep[0] == "" {
					continue
				}
				mod, ver = rep[0], rep[1]
			}
			key := mod + "@" + ver
			if seen[key] {
				continue
			}
			seen[key] = true
			if len(deps) >= budget {
				skipped++
				continue
			}
			deps = append(deps, depRef{ecosystem: "go", name: mod, version: ver})
		}
	}
	return deps, skipped
}

func (s *NetDepsStage) fetchGo(ctx context.Context, deps []depRef) ([]*depSource, error) {
	goBin, err := exec.LookPath(s.opts.GoBinary)
	if err != nil {
		return nil, fmt.Errorf("go toolchain not found (%s): %d Go modules not fetched", s.opts.GoBinary, len(deps))
	}
	tmp, err := os.MkdirTemp("", "safeanalyze-gomod-")
	if err != nil {
		return nil, err
	}
	env := append(cleanGoEnv(os.Environ()), "GOMODCACHE="+tmp+"/mod", "GOWORK=off", "GOFLAGS=-mod=mod", "GOTOOLCHAIN=local", "GOPATH="+tmp+"/gopath")
	defer func() {
		clean := exec.Command(goBin, "clean", "-modcache")
		clean.Env, clean.Dir = env, tmp
		_ = clean.Run()
		_ = os.RemoveAll(tmp)
	}()

	args := []string{"mod", "download", "-json"}
	for _, d := range deps {
		args = append(args, d.name+"@"+d.version)
	}
	cmd := exec.CommandContext(ctx, goBin, args...)
	cmd.Env, cmd.Dir = env, tmp // outside the target: its go.mod/go.work are not used
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	runErr := cmd.Run()

	var out []*depSource
	var errs []string
	dec := json.NewDecoder(&stdout)
	for {
		var m struct {
			Path, Version, Dir, Error string
		}
		if err := dec.Decode(&m); err != nil {
			break
		}
		id := m.Path + "@" + m.Version
		if m.Error != "" || m.Dir == "" {
			errs = append(errs, id+": "+m.Error)
			continue
		}
		prefix := "deps/go/" + id + "/"
		var files []*File
		_ = yara.WalkDirSorted(m.Dir, s.excluded, func(_ string, content []byte, rel string) error {
			if len(content) <= maxFileBytes && goBuildInput(rel) {
				files = append(files, depFile(prefix, rel, content))
			}
			return nil
		})
		out = append(out, &depSource{id: id, prefix: prefix, files: files})
	}
	if runErr != nil && len(out) == 0 {
		errs = append(errs, "go mod download: "+runErr.Error()+": "+strings.TrimSpace(stderr.String()))
	}
	if len(errs) > 0 {
		return out, errors.New("Go modules not fetched: " + strings.Join(errs, "; "))
	}
	return out, nil
}

// cleanGoEnv drops settings that would weaken checksum verification.
func cleanGoEnv(env []string) []string {
	var out []string
	for _, e := range env {
		k := e[:max(strings.Index(e, "="), 0)]
		switch k {
		case "GOFLAGS", "GONOSUMDB", "GONOSUMCHECK", "GOINSECURE", "GOMODCACHE", "GOWORK", "GOTOOLCHAIN", "GOPATH":
			continue
		}
		if k == "GOSUMDB" && strings.Contains(e, "off") {
			continue
		}
		out = append(out, e)
	}
	return out
}

// goBuildInput reports whether a file of a Go module can affect a consumer's
// build: Go, assembly and cgo sources, plus the root license. CI workflows,
// dev containers and Dockerfiles of a dependency never run for its users.
func goBuildInput(rel string) bool {
	rel = strings.ReplaceAll(rel, "\\", "/")
	switch strings.ToLower(path.Ext(rel)) {
	case ".go", ".s", ".c", ".h", ".cc", ".cpp", ".hpp", ".m":
		return !strings.Contains(rel, "testdata/")
	}
	base := strings.ToLower(path.Base(rel))
	return !strings.Contains(rel, "/") && (strings.HasPrefix(base, "license") || strings.HasPrefix(base, "copying") || strings.HasPrefix(base, "licence"))
}

func depFile(prefix, inner string, content []byte) *File {
	inner = strings.ReplaceAll(inner, "\\", "/")
	f := NewFile(prefix+inner, string(content))
	f.Dep = true
	f.Role = filerole.Classify(inner)
	if f.Role == filerole.Lockfile {
		f.Role = filerole.Other
	}
	return f
}

// npmDeps lists the direct dependencies recorded in each package-lock.json.
func npmDeps(files []*File, budget int) ([]depRef, int) {
	seen := map[string]bool{}
	var deps []depRef
	skipped := 0
	for _, f := range files {
		if f.Base() != "package-lock.json" || f.Role == filerole.Vendored {
			continue
		}
		var lock struct {
			Packages map[string]struct {
				lockPackage
				Dependencies         map[string]string `json:"dependencies"`
				DevDependencies      map[string]string `json:"devDependencies"`
				OptionalDependencies map[string]string `json:"optionalDependencies"`
			} `json:"packages"`
		}
		if json.Unmarshal([]byte(f.Content), &lock) != nil {
			continue
		}
		root := lock.Packages[""]
		var names []string
		for _, m := range []map[string]string{root.Dependencies, root.DevDependencies, root.OptionalDependencies} {
			for n := range m {
				names = append(names, n)
			}
		}
		sort.Strings(names)
		for _, n := range names {
			p, ok := lock.Packages["node_modules/"+n]
			if !ok || p.Resolved == "" || !strings.HasPrefix(p.Resolved, "https://") {
				continue
			}
			key := n + "@" + p.Version
			if seen[key] {
				continue
			}
			seen[key] = true
			if len(deps) >= budget {
				skipped++
				continue
			}
			deps = append(deps, depRef{ecosystem: "npm", name: n, version: p.Version, resolved: p.Resolved, integrity: p.Integrity})
		}
	}
	return deps, skipped
}

const maxTarballBytes = 64 << 20

func fetchNPM(ctx context.Context, deps []depRef) ([]*depSource, []string) {
	client := &http.Client{Timeout: 2 * time.Minute}
	var (
		mu   sync.Mutex
		out  []*depSource
		errs []string
		wg   sync.WaitGroup
	)
	sem := make(chan struct{}, 8)
	for _, d := range deps {
		d := d
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			src, err := fetchTarball(ctx, client, d)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				errs = append(errs, d.name+"@"+d.version+": "+err.Error())
			}
			if src != nil {
				out = append(out, src)
			}
		}()
	}
	wg.Wait()
	sort.Slice(out, func(i, j int) bool { return out[i].id < out[j].id })
	sort.Strings(errs)
	return out, errs
}

func fetchTarball(ctx context.Context, client *http.Client, d depRef) (*depSource, error) {
	id := d.name + "@" + d.version
	src := &depSource{id: id, prefix: "deps/npm/" + id + "/"}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.resolved, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "safeanalyze-netdeps")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download: %s", resp.Status)
	}
	data, err := io.ReadAll(io.LimitReader(resp.Body, maxTarballBytes+1))
	if err != nil {
		return nil, err
	}
	if len(data) > maxTarballBytes {
		return nil, fmt.Errorf("tarball larger than %d bytes", maxTarballBytes)
	}
	lockFile := &File{Rel: "package-lock.json", Role: filerole.Lockfile}
	if ok, known := checkIntegrity(data, d.integrity); known && !ok {
		src.findings = append(src.findings, finding("netdeps", "dep_integrity_mismatch", report.SeverityCritical, report.TierA, QDeps, lockFile, 1, d.resolved,
			"Downloaded tarball of "+id+" does not match the lockfile integrity hash: the registry serves different content", map[string]string{"dependency": id}))
		return src, nil // do not analyse content that failed verification as if it were the dependency
	}
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	tr := tar.NewReader(gz)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return src, fmt.Errorf("reading tarball: %w", err)
		}
		if h.Typeflag != tar.TypeReg || h.Size > maxFileBytes {
			continue
		}
		name := path.Clean(strings.TrimPrefix(h.Name, "/"))
		if i := strings.Index(name, "/"); i >= 0 {
			name = name[i+1:] // strip "package/"
		}
		if name == "" || strings.HasPrefix(name, "..") {
			continue
		}
		content, err := io.ReadAll(io.LimitReader(tr, maxFileBytes))
		if err != nil || yara.IsBinaryContent(content) {
			continue
		}
		src.files = append(src.files, depFile(src.prefix, name, content))
	}
	sort.Slice(src.files, func(i, j int) bool { return src.files[i].Rel < src.files[j].Rel })
	return src, nil
}

// checkIntegrity verifies an SRI string (sha512-… or sha1-…). known is false
// when there is no usable hash.
func checkIntegrity(data []byte, sri string) (ok, known bool) {
	for _, part := range strings.Fields(sri) {
		algo, b64, found := strings.Cut(part, "-")
		if !found {
			continue
		}
		want, err := base64.StdEncoding.DecodeString(b64)
		if err != nil {
			continue
		}
		switch algo {
		case "sha512":
			sum := sha512.Sum512(data)
			return bytes.Equal(sum[:], want), true
		case "sha1":
			sum := sha1.Sum(data)
			return bytes.Equal(sum[:], want), true
		}
	}
	return false, false
}
