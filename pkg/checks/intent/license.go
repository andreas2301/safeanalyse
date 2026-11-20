package intent

import (
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/user/safeanalyze/pkg/report"
)

// licenseKinds classifies license texts or SPDX identifiers, strongest
// obligation first.
var licenseKinds = []struct {
	id, kind string
	re       *regexp.Regexp
}{
	{"AGPL", "network-copyleft", regexp.MustCompile(`(?i)GNU AFFERO GENERAL PUBLIC LICENSE|\bAGPL-?[0-9]`)},
	{"SSPL", "source-available", regexp.MustCompile(`(?i)Server Side Public License|\bSSPL\b`)},
	{"BUSL", "source-available", regexp.MustCompile(`(?i)Business Source License|\bBUSL-[0-9]`)},
	{"Elastic", "source-available", regexp.MustCompile(`(?i)Elastic License|\bElastic-2\.0\b`)},
	{"Commons-Clause", "source-available", regexp.MustCompile(`(?i)Commons Clause`)},
	{"PolyForm", "source-available", regexp.MustCompile(`(?i)PolyForm`)},
	{"LGPL", "weak-copyleft", regexp.MustCompile(`(?i)GNU LESSER GENERAL PUBLIC LICENSE|\bLGPL-?[0-9]`)},
	{"GPL", "copyleft", regexp.MustCompile(`(?i)GNU GENERAL PUBLIC LICENSE|\bGPL-?[0-9]`)},
	{"MPL", "weak-copyleft", regexp.MustCompile(`(?i)Mozilla Public License|\bMPL-2\.0\b`)},
	{"EPL", "weak-copyleft", regexp.MustCompile(`(?i)Eclipse Public License|\bEPL-[0-9]`)},
	{"CC-BY-NC", "non-commercial", regexp.MustCompile(`(?i)NonCommercial|CC-BY-NC`)},
	{"Apache", "permissive", regexp.MustCompile(`(?i)Apache License|\bApache-2\.0\b`)},
	{"MIT", "permissive", regexp.MustCompile(`(?i)\bMIT License\b|Permission is hereby granted, free of charge|^MIT$`)},
	{"BSD", "permissive", regexp.MustCompile(`(?i)BSD|Redistribution and use in source and binary forms`)},
	{"ISC", "permissive", regexp.MustCompile(`(?i)\bISC\b`)},
	{"Unlicense", "permissive", regexp.MustCompile(`(?i)Unlicense|CC0`)},
}

var licenseObligation = map[string]string{
	"network-copyleft": "using it in a network service or distributing it requires releasing the complete corresponding source of your application under the same license",
	"copyleft":         "distributing a combined work requires releasing its source under the GPL",
	"source-available": "not open source: use (e.g. offering it as a service or in production) is restricted",
	"non-commercial":   "commercial use is not allowed",
	"weak-copyleft":    "modifications to the licensed files must be shared",
}

// ClassifyLicense returns the license ID and kind of a license text or SPDX
// expression, or empty strings.
func ClassifyLicense(text string) (string, string) {
	// An SPDX "A OR B" expression lets the user pick: report the most
	// permissive alternative.
	if len(text) < 200 && strings.Contains(text, " OR ") {
		for _, alt := range strings.Split(strings.Trim(text, "()"), " OR ") {
			if id, kind := ClassifyLicense(strings.TrimSpace(alt)); kind == "permissive" {
				return id, kind
			}
		}
	}
	for _, k := range licenseKinds {
		if k.re.MatchString(text) {
			return k.id, k.kind
		}
	}
	return "", ""
}

func analyzeLicense(files []*File) []report.Finding {
	var out []report.Finding
	hasRootLicense := false
	for _, f := range files {
		base := f.Base()
		dir := path.Dir(f.Rel)
		isLicense := strings.HasPrefix(base, "license") || strings.HasPrefix(base, "copying") || strings.HasPrefix(base, "licence")
		if isLicense && (dir == "." || strings.HasPrefix(f.Rel, "deps/") && strings.Count(f.Rel, "/") <= 4) {
			hasRootLicense = hasRootLicense || dir == "."
			out = append(out, licenseFinding(f, f.Content, "Repository license"))
		}
		if base == "package-lock.json" {
			out = append(out, lockLicenses(f)...)
		}
	}
	if !hasRootLicense && len(files) > 0 && !strings.HasPrefix(files[0].Rel, "deps/") {
		f := &File{Rel: "LICENSE", Role: "docs"}
		out = append(out, finding("license", "license_missing", report.SeverityMedium, report.TierB, QLicense, f, 0, "",
			"No license file at the repository root: no usage rights are granted by default", nil))
	}
	return out
}

func licenseFinding(f *File, text, what string) report.Finding {
	id, kind := ClassifyLicense(text)
	if id == "" {
		return finding("license", "license_unknown", report.SeverityLow, report.TierB, QLicense, f, 1, "", what+" could not be classified; review it", nil)
	}
	tier, sev := report.TierC, report.SeverityInfo
	msg := fmt.Sprintf("%s: %s (%s)", what, id, kind)
	if ob, ok := licenseObligation[kind]; ok {
		tier, sev = report.TierB, report.SeverityMedium
		if kind == "weak-copyleft" {
			tier = report.TierC
		}
		msg += ": " + ob
	}
	return finding("license", "license_"+strings.ReplaceAll(kind, "-", "_"), sev, tier, QLicense, f, 1, id, msg, map[string]string{"license": id, "kind": kind})
}

func lockLicenses(f *File) []report.Finding {
	var lock struct {
		Packages map[string]lockPackage `json:"packages"`
	}
	if json.Unmarshal([]byte(f.Content), &lock) != nil {
		return nil
	}
	byKind := map[string][]string{}
	for name, p := range lock.Packages {
		if name == "" {
			continue
		}
		var lic string
		switch v := p.License.(type) {
		case string:
			lic = v
		case map[string]any:
			lic, _ = v["type"].(string)
		}
		if lic == "" {
			continue
		}
		id, kind := ClassifyLicense(lic)
		if _, strong := licenseObligation[kind]; !strong || kind == "weak-copyleft" {
			continue
		}
		byKind[kind] = append(byKind[kind], lockPackageName(name)+" ("+id+")")
	}
	var kinds []string
	for k := range byKind {
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)
	var out []report.Finding
	for _, k := range kinds {
		pkgs := byKind[k]
		sort.Strings(pkgs)
		ex := pkgs
		if len(ex) > 8 {
			ex = ex[:8]
		}
		out = append(out, finding("license", "license_dependency_"+strings.ReplaceAll(k, "-", "_"), report.SeverityMedium, report.TierB, QLicense, f, 1, strings.Join(ex, ", "),
			fmt.Sprintf("%d npm dependencies are %s: %s", len(pkgs), k, licenseObligation[k]), map[string]string{"kind": k, "count": fmt.Sprint(len(pkgs))}))
	}
	return out
}

// lockPackageName returns the package name of a package-lock "packages" key
// such as "node_modules/a/node_modules/@b/c".
func lockPackageName(key string) string {
	if i := strings.LastIndex(key, "node_modules/"); i >= 0 {
		return key[i+len("node_modules/"):]
	}
	return key
}
