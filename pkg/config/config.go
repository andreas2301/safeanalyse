package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the top-level configuration for safeanalyze.
type Config struct {
	Scanners      []ScannerConfig      `yaml:"scanners"`
	Sanitization  SanitizationConfig   `yaml:"sanitization"`
	HiddenChars   HiddenCharsConfig    `yaml:"hidden_chars"`
	Entropy       EntropyConfig        `yaml:"entropy"`
	YARA          YARAConfig           `yaml:"yara"`
	Output        OutputConfig         `yaml:"output"`
	Sandbox       SandboxConfig        `yaml:"sandbox"`
}

// ScannerConfig defines an external tool to run.
type ScannerConfig struct {
	Name           string `yaml:"name"`
	Command        string `yaml:"command"`
	Enabled        bool   `yaml:"enabled"`
	FailOnFindings bool   `yaml:"fail_on_findings"`
}

// SanitizationConfig controls how files are cleaned before AI ingestion.
type SanitizationConfig struct {
	StripComments    bool     `yaml:"strip_comments"`
	RemoveNonASCII   bool     `yaml:"remove_non_ascii"`
	MaxFileSizeBytes int      `yaml:"max_file_size_bytes"`
	MaxLinesPerFile  int      `yaml:"max_lines_per_file"`
	AllowedExtensions []string `yaml:"allowed_extensions"`
	ExcludedPaths    []string `yaml:"excluded_paths"`
}

// HiddenCharsConfig controls invisible character detection.
type HiddenCharsConfig struct {
	Enabled        bool     `yaml:"enabled"`
	Categories     []string `yaml:"categories"`
	FailOnFindings bool     `yaml:"fail_on_findings"`
}

// EntropyConfig controls entropy analysis.
type EntropyConfig struct {
	Enabled        bool    `yaml:"enabled"`
	Threshold      float64 `yaml:"threshold"`
	MinLength      int     `yaml:"min_length"`
	FailOnFindings bool    `yaml:"fail_on_findings"`
}

// YARAConfig controls the built-in YARA-like rule engine.
type YARAConfig struct {
	Enabled        bool `yaml:"enabled"`
	FailOnFindings bool `yaml:"fail_on_findings"`
}

// OutputConfig controls how results are formatted.
type OutputConfig struct {
	Format          string `yaml:"format"` // markdown, json, plain
	SingleFile      bool   `yaml:"single_file"`
	IncludeFileTree bool   `yaml:"include_file_tree"`
	OutDir          string `yaml:"out_dir"`
}

// SandboxConfig controls isolation settings.
type SandboxConfig struct {
	Mode       string   `yaml:"mode"` // none, firejail, docker
	DockerImage string  `yaml:"docker_image"`
	FirejailProfile string `yaml:"firejail_profile"`
}

// DefaultConfig returns a sensible default configuration.
func DefaultConfig() *Config {
	return &Config{
		Scanners: []ScannerConfig{
			{
				Name:           "trufflehog",
				Command:        "trufflehog filesystem {path} --json",
				Enabled:        true,
				FailOnFindings: true,
			},
			{
				Name:           "semgrep",
				Command:        "semgrep --config=auto {path} --json",
				Enabled:        false,
				FailOnFindings: false,
			},
		},
		Sanitization: SanitizationConfig{
			StripComments:     true,
			RemoveNonASCII:    true,
			MaxFileSizeBytes:  50000,
			MaxLinesPerFile:   500,
			AllowedExtensions: []string{".go", ".py", ".js", ".ts", ".jsx", ".tsx", ".rs", ".java", ".c", ".cpp", ".h", ".rb", ".php", ".swift", ".kt"},
			ExcludedPaths:     []string{".git", "node_modules", "vendor", "target", "build", "dist", ".venv", "__pycache__", ".idea", ".vscode"},
		},
		HiddenChars: HiddenCharsConfig{
			Enabled:        true,
			Categories:     []string{"zero_width", "bidi", "control"},
			FailOnFindings: true,
		},
		Entropy: EntropyConfig{
			Enabled:        true,
			Threshold:      4.5,
			MinLength:      20,
			FailOnFindings: false,
		},
		YARA: YARAConfig{
			Enabled:        true,
			FailOnFindings: true,
		},
		Output: OutputConfig{
			Format:          "markdown",
			SingleFile:      false,
			IncludeFileTree: true,
			OutDir:          "./safeanalyze-out",
		},
		Sandbox: SandboxConfig{
			Mode:            "none",
			DockerImage:     "alpine:latest",
			FirejailProfile: "default",
		},
	}
}

// Load reads configuration from a YAML file.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}
	cfg := DefaultConfig()
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}
	return cfg, nil
}

// WriteExample writes a default config file to the given path.
func WriteExample(path string) error {
	cfg := DefaultConfig()
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	header := []byte("# safeanalyze configuration file\n# See https://github.com/user/safeanalyze for docs\n\n")
	return os.WriteFile(path, append(header, data...), 0644)
}
