// Package install manages source-based installation of external scanners
// and ML models used by safeanalyze.
package install

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// ScannerMeta holds metadata for an installable scanner.
type ScannerMeta struct {
	Name     string
	Repo     string
	Language string // "go" or "python"
	Binary   string // binary name after installation
}

// KnownScanners is the registry of supported external scanners.
// Repository URLs for bumblebee and prompt-injection-scanner are placeholders
// pointing to the foundationmachines organization; update them when official
// repositories are published.
var KnownScanners = map[string]ScannerMeta{
	"semgrep": {
		Name:     "semgrep",
		Repo:     "https://github.com/semgrep/semgrep.git",
		Language: "python",
		Binary:   "semgrep",
	},
	"bumblebee": {
		Name:     "bumblebee",
		Repo:     "https://github.com/foundationmachines/bumblebee.git",
		Language: "go",
		Binary:   "bumblebee",
	},
	"prompt-injection-scanner": {
		Name:     "prompt-injection-scanner",
		Repo:     "https://github.com/foundationmachines/prompt-injection-scanner.git",
		Language: "python",
		Binary:   "prompt-injection-scanner",
	},
	"gitleaks": {
		Name:     "gitleaks",
		Repo:     "https://github.com/gitleaks/gitleaks.git",
		Language: "go",
		Binary:   "gitleaks",
	},
	"trufflehog": {
		Name:     "trufflehog",
		Repo:     "https://github.com/trufflesecurity/trufflehog.git",
		Language: "go",
		Binary:   "trufflehog",
	},
}

// safeanalyzeHome returns ~/.safeanalyze (or equivalent on Windows).
func safeanalyzeHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".safeanalyze"), nil
}

// toolDir returns the clone directory for a scanner.
func toolDir(name string) (string, error) {
	home, err := safeanalyzeHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "tools", name), nil
}

// binDir returns the safeanalyze binary directory.
func binDir() (string, error) {
	home, err := safeanalyzeHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "bin"), nil
}

// venvDir returns the virtualenv directory for a Python scanner.
func venvDir(name string) (string, error) {
	home, err := safeanalyzeHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "venv", name), nil
}

// modelDir returns the models directory.
func modelDir() (string, error) {
	home, err := safeanalyzeHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "models"), nil
}

// Install clones and installs the named scanner.
func Install(name string) error {
	meta, ok := KnownScanners[name]
	if !ok {
		return fmt.Errorf("unknown scanner %q", name)
	}
	if err := ensureDirs(meta.Language); err != nil {
		return err
	}
	toolPath, err := toolDir(name)
	if err != nil {
		return err
	}
	if err := gitClone(meta.Repo, toolPath); err != nil {
		return err
	}
	switch meta.Language {
	case "go":
		return installGo(name, toolPath)
	case "python":
		return installPython(name, toolPath)
	default:
		return fmt.Errorf("unsupported language %q for %s", meta.Language, name)
	}
}

// InstallAll installs all known scanners.
func InstallAll() error {
	for name := range KnownScanners {
		if err := Install(name); err != nil {
			return err
		}
	}
	return nil
}

// InstallModel downloads the protectai prompt-injection ONNX model.
func InstallModel() error {
	dir, err := modelDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating model dir: %w", err)
	}

	const url = "https://huggingface.co/protectai/deberta-v3-base-prompt-injection/resolve/main/onnx/model.onnx"
	dst := filepath.Join(dir, "deberta-v3-base-prompt-injection.onnx")
	if err := downloadFile(url, dst); err != nil {
		// Fallback to huggingface-cli if available.
		if hfErr := installModelViaHFCLI(dir); hfErr == nil {
			return nil
		}
		return err
	}
	return nil
}

func installModelViaHFCLI(dir string) error {
	if _, err := exec.LookPath("huggingface-cli"); err != nil {
		return err
	}
	cmd := exec.Command("huggingface-cli", "download", "protectai/deberta-v3-base-prompt-injection", "--local-dir", dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ResolveBinary returns the path to an installed binary or empty string.
func ResolveBinary(name string) string {
	bd, err := binDir()
	if err != nil {
		return ""
	}
	candidate := filepath.Join(bd, name)
	if runtime.GOOS == "windows" {
		candidate += ".exe"
	}
	if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
		return candidate
	}

	// For Python packages installed in venv, check venv bin.
	meta, ok := KnownScanners[name]
	if ok && meta.Language == "python" {
		vd, err := venvDir(name)
		if err == nil {
			pyBin := filepath.Join(vd, "bin", name)
			if runtime.GOOS == "windows" {
				pyBin = filepath.Join(vd, "Scripts", name+".exe")
			}
			if info, err := os.Stat(pyBin); err == nil && !info.IsDir() {
				return pyBin
			}
		}
	}
	return ""
}

func ensureDirs(lang string) error {
	dirs, err := dirsForLang(lang)
	if err != nil {
		return err
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("creating dir %s: %w", d, err)
		}
	}
	return nil
}

func dirsForLang(lang string) ([]string, error) {
	home, err := safeanalyzeHome()
	if err != nil {
		return nil, err
	}
	dirs := []string{filepath.Join(home, "bin"), filepath.Join(home, "tools")}
	if lang == "python" {
		dirs = append(dirs, filepath.Join(home, "venv"))
	}
	return dirs, nil
}

func gitClone(repo, dst string) error {
	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("directory %s already exists; remove it to reinstall", dst)
	}
	cmd := exec.Command("git", "clone", "--depth", "1", repo, dst)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cloning %s: %w", repo, err)
	}
	return nil
}

func installGo(name, toolPath string) error {
	bd, err := binDir()
	if err != nil {
		return err
	}
	outPath := filepath.Join(bd, name)
	if runtime.GOOS == "windows" {
		outPath += ".exe"
	}

	candidates := []string{"./cmd/" + name, "."}
	for _, c := range candidates {
		if ok, _ := isGoMainPackage(toolPath, c); ok {
			cmd := exec.Command("go", "build", "-o", outPath, c)
			cmd.Dir = toolPath
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("building %s: %w", name, err)
			}
			return nil
		}
	}

	// Fallback to go install ./...
	cmd := exec.Command("go", "install", "./...")
	cmd.Dir = toolPath
	cmd.Env = append(os.Environ(), "GOBIN="+bd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go install %s: %w", name, err)
	}
	return nil
}

func isGoMainPackage(dir, pattern string) (bool, error) {
	cmd := exec.Command("go", "list", "-f", "{{.Name}}", pattern)
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(string(out)) == "main", nil
}

func installPython(name, toolPath string) error {
	vd, err := venvDir(name)
	if err != nil {
		return err
	}
	py := "python3"
	if runtime.GOOS == "windows" {
		py = "python"
	}

	cmd := exec.Command(py, "-m", "venv", vd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("creating venv for %s: %w", name, err)
	}

	pip := filepath.Join(vd, "bin", "pip")
	if runtime.GOOS == "windows" {
		pip = filepath.Join(vd, "Scripts", "pip.exe")
	}

	cmd = exec.Command(pip, "install", "-e", ".")
	cmd.Dir = toolPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pip install %s: %w", name, err)
	}

	if err := linkPythonBinary(name, vd); err != nil {
		return err
	}
	return nil
}

func linkPythonBinary(name, vd string) error {
	bd, err := binDir()
	if err != nil {
		return err
	}
	src := filepath.Join(vd, "bin", name)
	dst := filepath.Join(bd, name)
	if runtime.GOOS == "windows" {
		src = filepath.Join(vd, "Scripts", name+".exe")
		dst += ".exe"
	}
	if _, err := os.Stat(src); err != nil {
		return nil
	}
	if err := os.Remove(dst); err != nil && !os.IsNotExist(err) {
		return err
	}
	return os.Symlink(src, dst)
}

func downloadFile(url, dst string) error {
	tmp := dst + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer os.Remove(tmp)

	resp, err := http.Get(url)
	if err != nil {
		_ = f.Close()
		return fmt.Errorf("downloading model: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		_ = f.Close()
		return fmt.Errorf("downloading model: status %d", resp.StatusCode)
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		_ = f.Close()
		return fmt.Errorf("writing model: %w", err)
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, dst)
}
