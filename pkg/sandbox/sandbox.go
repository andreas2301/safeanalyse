package sandbox

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Config controls sandbox behavior.
type Config struct {
	Mode            string // "none", "docker", "firejail"
	DockerImage     string
	FirejailProfile string
	ReadOnlyPaths   []string
	NetworkDisabled bool
}

// Runner prepares and executes commands inside a sandbox.
type Runner struct {
	cfg Config
}

// NewRunner creates a sandbox runner.
func NewRunner(cfg Config) *Runner {
	return &Runner{cfg: cfg}
}

// Available checks if the configured sandbox mode is available on this system.
func (r *Runner) Available() bool {
	switch r.cfg.Mode {
	case "docker":
		_, err := exec.LookPath("docker")
		return err == nil
	case "firejail":
		_, err := exec.LookPath("firejail")
		return err == nil
	case "none":
		return true
	default:
		return false
	}
}

// Command builds a sandboxed command that wraps the given args.
// Returns nil if sandbox mode is "none".
func (r *Runner) Command(args ...string) *exec.Cmd {
	switch r.cfg.Mode {
	case "docker":
		return r.dockerCmd(args)
	case "firejail":
		return r.firejailCmd(args)
	default:
		return exec.Command(args[0], args[1:]...)
	}
}

func (r *Runner) dockerCmd(args []string) *exec.Cmd {
	if len(args) == 0 {
		return nil
	}

	dockerArgs := []string{
		"run", "--rm", "-it",
		"--network", "none",
		"--read-only",
		"--tmpfs", "/tmp:noexec,nosuid,size=100m",
		"--cap-drop", "ALL",
		"--security-opt", "no-new-privileges:true",
	}

	// Mount read-only volumes
	for _, p := range r.cfg.ReadOnlyPaths {
		abs, _ := absPath(p)
		dockerArgs = append(dockerArgs, "-v", fmt.Sprintf("%s:%s:ro", abs, abs))
	}

	dockerArgs = append(dockerArgs, r.cfg.DockerImage)
	dockerArgs = append(dockerArgs, args...)

	return exec.Command("docker", dockerArgs...)
}

func (r *Runner) firejailCmd(args []string) *exec.Cmd {
	if len(args) == 0 {
		return nil
	}

	fjArgs := []string{
		"--net=none",
		"--private",
		"--shell=none",
		"--nogroups",
		"--caps.drop=all",
	}

	if r.cfg.FirejailProfile != "" && r.cfg.FirejailProfile != "default" {
		fjArgs = append(fjArgs, "--profile="+r.cfg.FirejailProfile)
	}

	for _, p := range r.cfg.ReadOnlyPaths {
		abs, _ := absPath(p)
		fjArgs = append(fjArgs, "--read-only="+abs)
	}

	fjArgs = append(fjArgs, args...)
	return exec.Command("firejail", fjArgs...)
}

// LaunchClaude builds a command to launch Claude with the ingest directory mounted.
func (r *Runner) LaunchClaude(ingestDir string) (*exec.Cmd, error) {
	if !r.Available() {
		return nil, fmt.Errorf("sandbox mode %q is not available on this system", r.cfg.Mode)
	}

	absDir, err := absPath(ingestDir)
	if err != nil {
		return nil, err
	}

	switch r.cfg.Mode {
	case "docker":
		// Use a generic shell container; user runs claude inside
		return exec.Command("docker", "run", "--rm", "-it",
			"--network", "none",
			"--read-only",
			"--tmpfs", "/tmp:noexec,nosuid,size=100m",
			"--cap-drop", "ALL",
			"--security-opt", "no-new-privileges:true",
			"-v", fmt.Sprintf("%s:/ingest:ro", absDir),
			r.cfg.DockerImage,
			"/bin/sh", "-c",
			fmt.Sprintf("echo 'Ingest files available at /ingest'; echo 'Launch Claude here manually'; /bin/sh"),
		), nil

	case "firejail":
		return exec.Command("firejail", "--net=none", "--private", "--shell=none",
			"--read-only="+absDir,
			"claude",
		), nil

	default:
		return nil, fmt.Errorf("no sandbox command for mode %q", r.cfg.Mode)
	}
}

// AutoDetect selects the best available sandbox mode for the current OS.
func AutoDetect() string {
	switch runtime.GOOS {
	case "linux":
		if _, err := exec.LookPath("firejail"); err == nil {
			return "firejail"
		}
		if _, err := exec.LookPath("docker"); err == nil {
			return "docker"
		}
	case "windows", "darwin":
		if _, err := exec.LookPath("docker"); err == nil {
			return "docker"
		}
	}
	return "none"
}

func absPath(p string) (string, error) {
	if strings.HasPrefix(p, "~/") || strings.HasPrefix(p, "~\\") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		p = home + p[1:]
	}
	return os.Getwd()
}
