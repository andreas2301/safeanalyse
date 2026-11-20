package report

import (
	"fmt"
	"sort"
)

// EgressEntry is one host of the call-home map.
type EgressEntry struct {
	Host      string
	Class     string
	Tier      string
	Guard     string
	Locations []string
}

// EgressMap groups egress findings of tiers A and B by host, most relevant
// first.
func EgressMap(findings []Finding) []EgressEntry {
	byHost := map[string]*EgressEntry{}
	for _, f := range findings {
		if f.Source != "egress" || f.Tier == TierC || f.Tier == "" {
			continue
		}
		host := f.Attributes["host"]
		if host == "" {
			host = f.RuleID + ": " + f.Match
		}
		e, ok := byHost[host]
		if !ok {
			e = &EgressEntry{Host: host, Class: f.Attributes["class"], Tier: f.Tier, Guard: f.Attributes["guard"]}
			byHost[host] = e
		}
		if f.Tier < e.Tier {
			e.Tier = f.Tier
		}
		if f.Attributes["guard"] != e.Guard {
			e.Guard = "mixed"
		}
		e.Locations = append(e.Locations, fmt.Sprintf("%s:%d", f.File, f.Line))
	}
	out := make([]EgressEntry, 0, len(byHost))
	for _, e := range byHost {
		out = append(out, *e)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Tier != out[j].Tier {
			return out[i].Tier < out[j].Tier
		}
		return out[i].Host < out[j].Host
	})
	return out
}

// RuleCount is the number of findings of one rule in one file role.
type RuleCount struct {
	RuleID   string
	FileRole string
	Count    int
}

// TierSummary counts the findings of a tier by rule and file role, largest
// group first.
func TierSummary(findings []Finding, tier string) []RuleCount {
	counts := map[[2]string]int{}
	for _, f := range findings {
		if f.Tier == tier {
			counts[[2]string{f.RuleID, f.FileRole}]++
		}
	}
	out := make([]RuleCount, 0, len(counts))
	for k, n := range counts {
		out = append(out, RuleCount{RuleID: k[0], FileRole: k[1], Count: n})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count != out[j].Count {
			return out[i].Count > out[j].Count
		}
		if out[i].RuleID != out[j].RuleID {
			return out[i].RuleID < out[j].RuleID
		}
		return out[i].FileRole < out[j].FileRole
	})
	return out
}

// FindingsInTier returns the findings of a tier, most severe first. An empty
// tier selects findings without a tier (reports from older pipelines).
func FindingsInTier(findings []Finding, tier string) []Finding {
	var out []Finding
	for _, f := range findings {
		if f.Tier == tier {
			out = append(out, f)
		}
	}
	SortFindingsBySeverity(out)
	return out
}
