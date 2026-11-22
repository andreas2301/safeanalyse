package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/install"
	"github.com/user/safeanalyze/pkg/sandbox"
	"github.com/user/safeanalyze/pkg/version"
)

// DoctorCheck describes one environment or scanner readiness check.
type DoctorCheck struct {
	Name    string `json:"name"`
	Kind    string `json:"kind"`   // runtime, scanner, catalog, sandbox
	Status  string `json:"status"` // ok, missing, disabled
	Detail  string `json:"detail,omitempty"`
	Enabled bool   `json:"enabled"`
}

// DoctorReport is the structured output of `safeanalyze doctor`.
type DoctorReport struct {
	Version string        `json:"safeanalyze_version"`
	OS      string        `json:"os"`
	Arch    string        `json:"arch"`
	Checks  []DoctorCheck `json:"checks"`
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Verify toolchain, external scanners, catalogs, and sandbox readiness",
	Long: `Run pre-flight diagnostics for CI/CD and enterprise deployments.
Checks Git, Go, sandbox runtimes (docker/firejail), configured external
scanners (bumblebee, semgrep, trufflehog, gitleaks, plugins), and Bumblebee
exposure catalogs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		asJSON, _ := cmd.Flags().GetBool("json")
		rep := runDoctor()

		if asJSON {
			data, err := json.MarshalIndent(rep, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(data))
			return nil
		}

		fmt.Printf("safeanalyze doctor (v%s, %s/%s)\n\n", rep.Version, rep.OS, rep.Arch)
		for _, c := range rep.Checks {
			tag := "[OK]      "
			switch c.Status {
			case "missing":
				tag = "[MISSING] "
			case "disabled":
				tag = "[DISABLED]"
			}
			if c.Detail != "" {
				fmt.Printf("  %s %-26s (%s) %s\n", tag, c.Name, c.Kind, c.Detail)
			} else {
				fmt.Printf("  %s %-26s (%s)\n", tag, c.Name, c.Kind)
			}
		}
		return nil
	},
}

func runDoctor() DoctorReport {
	rep := DoctorReport{
		Version: version.Version,
		OS:      runtime.GOOS,
		Arch:    runtime.GOARCH,
	}

	for _, bin := range []string{"git", "go"} {
		if path, err := exec.LookPath(bin); err == nil {
			rep.Checks = append(rep.Checks, DoctorCheck{Name: bin, Kind: "runtime", Status: "ok", Detail: path, Enabled: true})
		} else {
			rep.Checks = append(rep.Checks, DoctorCheck{Name: bin, Kind: "runtime", Status: "missing", Detail: "not found on PATH", Enabled: true})
		}
	}

	sbMode := cfg.Sandbox.Mode
	if sbMode == "" {
		sbMode = "none"
	}
	sb := sandbox.NewRunner(sandbox.Config{Mode: sbMode})
	sbStatus := "ok"
	sbDetail := "mode=" + sbMode + ", detected=" + sandbox.AutoDetect()
	if sbMode == "none" {
		sbStatus = "disabled"
	} else if !sb.Available() {
		sbStatus = "missing"
	}
	rep.Checks = append(rep.Checks, DoctorCheck{Name: "sandbox", Kind: "sandbox", Status: sbStatus, Detail: sbDetail, Enabled: sbMode != "none"})

	for _, sc := range cfg.Scanners {
		binName := sc.Name
		path := lookupScanner(binName)
		status := "disabled"
		detail := "disabled in config"
		if sc.Enabled {
			if path != "" {
				status = "ok"
				detail = path
			} else {
				status = "missing"
				detail = "not found in PATH or ~/.safeanalyze/bin"
			}
		} else if path != "" {
			detail = "installed at " + path + " (disabled in config)"
		}
		rep.Checks = append(rep.Checks, DoctorCheck{
			Name:    sc.Name,
			Kind:    "scanner",
			Status:  status,
			Detail:  detail,
			Enabled: sc.Enabled,
		})

		if sc.Name == "bumblebee" {
			cat := sc.ExposureCatalog
			if cat == "" {
				cat = install.ResolveAsset("bumblebee", "threat_intel")
			}
			catStatus := "disabled"
			catDetail := "no catalog configured"
			if cat != "" {
				if _, err := os.Stat(cat); err == nil {
					catStatus = "ok"
					catDetail = cat
				} else if sc.Enabled {
					catStatus = "missing"
					catDetail = "configured path not found: " + cat
				}
			} else if sc.Enabled {
				catStatus = "missing"
				catDetail = "set exposure_catalog or run 'safeanalyze install bumblebee'"
			}
			rep.Checks = append(rep.Checks, DoctorCheck{
				Name:    "bumblebee-catalog",
				Kind:    "catalog",
				Status:  catStatus,
				Detail:  catDetail,
				Enabled: sc.Enabled,
			})
		}
	}

	return rep
}

func lookupScanner(name string) string {
	if p, err := exec.LookPath(name); err == nil {
		return p
	}
	return install.ResolveBinary(name)
}

func init() {
	doctorCmd.Flags().Bool("json", false, "output diagnostics as JSON")
}
