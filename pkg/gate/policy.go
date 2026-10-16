// Package gate implements the autoresearch keep/revert decision rules
// (CLAUDE.md "Set taxonomy" and "Decision rules") on two directories of
// `safeanalyze eval --json` results.
package gate

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Role is how a labeled set takes part in the decision.
type Role string

// Set roles from CLAUDE.md "Set taxonomy".
const (
	RolePrimary    Role = "PRIMARY"
	RoleRecallOnly Role = "RECALL-ONLY"
	RoleGuardOnly  Role = "GUARD-ONLY"
	RoleReportOnly Role = "REPORT-ONLY"
	RoleSealed     Role = "SEALED"
)

// Policy holds the set roles and thresholds of the decision rules.
type Policy struct {
	Roles struct {
		Primary    []string `yaml:"primary"`
		RecallOnly []string `yaml:"recall_only"`
		GuardOnly  []string `yaml:"guard_only"`
		ReportOnly []string `yaml:"report_only"`
	} `yaml:"roles"`
	HoldoutSuffix string `yaml:"holdout_suffix"`
	SealedSuffix  string `yaml:"sealed_suffix"`
	Tolerance     struct {
		MaxTPLost int `yaml:"max_tp_lost"`
		MaxSets   int `yaml:"max_sets"`
	} `yaml:"tolerance"`
	Generalization struct {
		MinHoldoutRatio float64 `yaml:"min_holdout_ratio"`
	} `yaml:"generalization"`
	Significance struct {
		Resamples  int     `yaml:"resamples"`
		Seed       uint64  `yaml:"seed"`
		Confidence float64 `yaml:"confidence"`
	} `yaml:"significance"`
	Latency struct {
		BudgetP95Ms  float64 `yaml:"budget_p95_ms"`
		MaxRisePct   float64 `yaml:"max_rise_pct"`
		MinDropPct   float64 `yaml:"min_drop_pct"`
		NoiseFloorMs float64 `yaml:"noise_floor_ms"`
	} `yaml:"latency"`
}

// DefaultPolicy returns the policy committed in testdata/eval/policy.yaml.
func DefaultPolicy() *Policy {
	p := &Policy{HoldoutSuffix: "-holdout", SealedSuffix: "-sealed"}
	p.Roles.Primary = []string{"llmail", "browsesafe", "toolresp"}
	p.Roles.RecallOnly = []string{"multilingual"}
	p.Roles.GuardOnly = []string{"deepset", "nemotron", "spanish"}
	p.Roles.ReportOnly = []string{"promptprotection-stress", "hardbenign"}
	p.Tolerance.MaxTPLost = 1
	p.Tolerance.MaxSets = 1
	p.Generalization.MinHoldoutRatio = 0.5
	p.Significance.Resamples = 10000
	p.Significance.Seed = 1337
	p.Significance.Confidence = 0.95
	p.Latency.BudgetP95Ms = 100
	p.Latency.MaxRisePct = 15
	p.Latency.MinDropPct = 10
	p.Latency.NoiseFloorMs = 2
	return p
}

// LoadPolicy reads and validates a policy YAML file. Unknown keys are errors.
func LoadPolicy(path string) (*Policy, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading policy: %w", err)
	}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	p := &Policy{}
	if err := dec.Decode(p); err != nil {
		return nil, fmt.Errorf("parsing policy %s: %w", path, err)
	}
	if err := p.Validate(); err != nil {
		return nil, fmt.Errorf("policy %s: %w", path, err)
	}
	return p, nil
}

// Validate checks that roles are disjoint and thresholds are in range.
func (p *Policy) Validate() error {
	if len(p.Roles.Primary) == 0 {
		return fmt.Errorf("roles.primary is empty")
	}
	if p.HoldoutSuffix == "" || p.SealedSuffix == "" || p.HoldoutSuffix == p.SealedSuffix {
		return fmt.Errorf("holdout_suffix and sealed_suffix must be set and differ")
	}
	seen := map[string]Role{}
	for role, sets := range p.roleSets() {
		for _, s := range sets {
			if prev, ok := seen[s]; ok {
				return fmt.Errorf("set %q has two roles (%s, %s)", s, prev, role)
			}
			seen[s] = role
		}
	}
	switch {
	case p.Tolerance.MaxTPLost < 0 || p.Tolerance.MaxSets < 0:
		return fmt.Errorf("tolerance values must be >= 0")
	case p.Generalization.MinHoldoutRatio < 0:
		return fmt.Errorf("generalization.min_holdout_ratio must be >= 0")
	case p.Significance.Resamples < 100:
		return fmt.Errorf("significance.resamples must be >= 100")
	case p.Significance.Confidence <= 0 || p.Significance.Confidence >= 1:
		return fmt.Errorf("significance.confidence must be in (0, 1)")
	case p.Latency.BudgetP95Ms <= 0 || p.Latency.MaxRisePct < 0 || p.Latency.MinDropPct <= 0 || p.Latency.NoiseFloorMs < 0:
		return fmt.Errorf("latency thresholds out of range")
	}
	return nil
}

func (p *Policy) roleSets() map[Role][]string {
	return map[Role][]string{
		RolePrimary:    p.Roles.Primary,
		RoleRecallOnly: p.Roles.RecallOnly,
		RoleGuardOnly:  p.Roles.GuardOnly,
		RoleReportOnly: p.Roles.ReportOnly,
	}
}
