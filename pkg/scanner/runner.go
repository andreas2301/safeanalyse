package scanner

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/user/safeanalyze/pkg/config"
)

// Result holds the output of a scanner run.
type Result struct {
	Name       string `json:"name"`
	Command    string `json:"command"`
	ExitCode   int    `json:"exit_code"`
	Stdout     string `json:"stdout"`
	Stderr     string `json:"stderr"`
	Duration   time.Duration `json:"duration"`
	Findings   bool   `json:"findings"`
	Skipped    bool   `json:"skipped"`
	Error      string `json:"error,omitempty"`
}

// Runner executes external security scanners.
type Runner struct {
	scanners []config.ScannerConfig
}

// NewRunner creates a scanner runner from config.
func NewRunner(cfg []config.ScannerConfig) *Runner {
	return &Runner{scanners: cfg}
}

// RunAll executes all enabled scanners against the given path.
func (r *Runner) RunAll(path string, verbose bool) ([]Result, error) {
	var results []Result
	var hasError bool

	for _, s := range r.scanners {
		if !s.Enabled {
			results = append(results, Result{Name: s.Name, Skipped: true})
			continue
		}

		res := r.runOne(s, path, verbose)
		results = append(results, res)

		if res.Error != "" {
			hasError = true
		}
	}

	if hasError {
		return results, fmt.Errorf("one or more scanners failed")
	}
	return results, nil
}

func (r *Runner) runOne(s config.ScannerConfig, path string, verbose bool) Result {
	res := Result{Name: s.Name, Command: s.Command}

	cmdStr := strings.ReplaceAll(s.Command, "{path}", path)
	parts := strings.Fields(cmdStr)
	if len(parts) == 0 {
		res.Error = "empty command"
		return res
	}

	if verbose {
		color.Yellow("[scanner] running %s: %s\n", s.Name, cmdStr)
	}

	start := time.Now()
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	res.Duration = time.Since(start)
	res.ExitCode = cmd.ProcessState.ExitCode()
	res.Stdout = string(out)

	if err != nil {
		// Non-zero exit doesn't always mean failure for scanners
		if cmd.ProcessState != nil && cmd.ProcessState.ExitCode() != 0 {
			res.Findings = true
		} else {
			res.Error = err.Error()
		}
	}

	// Heuristic: if output contains findings keywords, mark as findings
	lower := strings.ToLower(res.Stdout)
	if strings.Contains(lower, "finding") || strings.Contains(lower, "vulnerability") ||
		strings.Contains(lower, "secret") || strings.Contains(lower, "issue") ||
		strings.Contains(lower, "match") || strings.Contains(lower, "result") {
		res.Findings = true
	}

	if verbose {
		if res.Findings {
			color.Red("[scanner] %s: findings detected (exit %d, %s)\n", s.Name, res.ExitCode, res.Duration)
		} else if res.Error != "" {
			color.Red("[scanner] %s: error: %s\n", s.Name, res.Error)
		} else {
			color.Green("[scanner] %s: clean (exit %d, %s)\n", s.Name, res.ExitCode, res.Duration)
		}
	}

	return res
}

// PrintSummary prints a human-readable summary of scanner results.
func PrintSummary(results []Result) {
	fmt.Println()
	color.Cyan("========================================")
	color.Cyan("         SCANNER SUMMARY")
	color.Cyan("========================================")

	var findings, errors, clean int
	for _, r := range results {
		if r.Skipped {
			fmt.Printf("  %s: %s\n", r.Name, color.YellowString("SKIPPED"))
			continue
		}
		if r.Error != "" {
			errors++
			fmt.Printf("  %s: %s (%s)\n", r.Name, color.RedString("ERROR"), r.Error)
		} else if r.Findings {
			findings++
			fmt.Printf("  %s: %s\n", r.Name, color.RedString("FINDINGS"))
		} else {
			clean++
			fmt.Printf("  %s: %s\n", r.Name, color.GreenString("CLEAN"))
		}
	}

	fmt.Println()
	if errors > 0 {
		color.Red("Result: %d error(s), %d finding(s), %d clean\n", errors, findings, clean)
	} else if findings > 0 {
		color.Yellow("Result: %d finding(s), %d clean\n", findings, clean)
	} else {
		color.Green("Result: all scanners clean\n")
	}
}
