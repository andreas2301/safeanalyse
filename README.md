# safeanalyze

A Go CLI tool that sanitizes and scans untrusted code repositories **before** feeding them to Claude or other AI assistants. Implements defense-in-depth inspired by [Zones of Distrust](https://github.com/bluvibytes/zone-of-distrust).

## Why?

Prompt injection via malicious code is real. A repo can contain:
- Hidden comments telling Claude to "ignore all prior instructions"
- Zero-width Unicode characters hiding payload instructions
- Bidirectional text overrides that reorder displayed code
- High-entropy encoded payloads (base64, hex)
- Secrets or malware mixed with legitimate source

**safeanalyze** runs a security pipeline so Claude never sees raw, unverified code.

## Pipeline

```
Untrusted Repo
    |
    v
[Clone / Read]       <-- git clone wrapper with auto-cleanup
    |
    v
[External Scanners]  <-- trufflehog, semgrep, yara (optional)
    |
    v
[Built-in YARA]      <-- prompt injection, backdoor, credential rules
    |
    v
[Entropy Analysis]   <-- high-entropy strings, base64/hex blobs
    |
    v
[Hidden Char Scan]   <-- zero-width, bidi overrides, control chars
    |
    v
[Sanitization]       <-- AST-aware comment stripping, non-ASCII removal,
                           size limits, extension filtering
    |
    v
[Diff Review]        <-- colored diff showing exactly what was removed
    |
    v
[Format for AI]      <-- markdown/JSON/plain, bounded chunks
    |
    v
[Sandbox Launch]     <-- Docker or Firejail isolated Claude session
```

## Install

```bash
git clone https://github.com/youruser/safeanalyze.git
cd safeanalyze
go build -o safeanalyze .
```

Or with `go install`:
```bash
go install github.com/youruser/safeanalyze@latest
```

## Quick Start

```bash
# Create a default config
./safeanalyze init

# Full pipeline on a repo
./safeanalyze ingest ./my-suspicious-repo

# Or step by step
./safeanalyze scan ./my-suspicious-repo          # all scanners
./safeanalyze sanitize ./my-suspicious-repo      # sanitize only
./safeanalyze diff ./my-suspicious-repo ./safeanalyze-out/sanitized

# Clone + analyze a remote repo
./safeanalyze clone https://github.com/user/repo.git

# Launch Claude in sandbox after ingestion
./safeanalyze ingest ./my-suspicious-repo --sandbox
```

## Commands

| Command | Description |
|---------|-------------|
| `init [path]` | Create a `safeanalyze.yaml` config file |
| `scan <path>` | Run external scanners, YARA rules, entropy analysis, hidden chars |
| `sanitize <src> [dst]` | Strip comments (AST-aware), remove non-ASCII, enforce limits |
| `ingest <path>` | Full pipeline: scan → sanitize → format for AI |
| `diff <orig> <sanitized>` | Show colored diff of what sanitization removed |
| `clone <url> [dir]` | Clone repo, run ingest, auto-delete raw clone |

## Features

### 1. AST-Aware Comment Stripping

For **Go files**, uses `go/ast` to precisely remove comments without touching string literals containing `//` or `/*`. For other languages, uses an enhanced regex engine with string-literal detection.

Supported: Go, Python, JavaScript/TypeScript, Rust, Java, C/C++, Ruby, PHP, Swift, Kotlin, Scala, Shell, SQL, HTML, CSS, Lua, and more.

### 2. Built-in YARA-like Rule Engine

Pure-Go rule engine with embedded detection patterns:

| Rule | Severity | Detects |
|------|----------|---------|
| `prompt_injection_comment` | critical | "ignore previous instructions", "system prompt", "DAN mode", "jailbreak" |
| `obfuscated_javascript` | high | eval(Function(...)), String.fromCharCode, atob, hex escapes |
| `suspicious_shell` | high | curl \| bash, wget \| bash, netcat reverse shells |
| `credential_hardcode` | medium | password=, api_key=, secret=, AWS keys |
| `suspicious_imports` | medium | subprocess, child_process, urllib requests |
| `data_exfiltration` | high | fetch to external URLs, axios post, XMLHttpRequest |
| `backdoor_indicator` | critical | reverse_shell, bind_shell, keylogger, rootkit |

### 3. Entropy Analysis

Detects high-entropy strings that may be encoded secrets or payloads:
- **Shannon entropy** scoring (configurable threshold)
- **Base64 blob detection** with validation
- **Hex blob detection**
- String-literal-aware scanning (only checks quoted strings)
- Filters false positives (UUIDs, URLs, repeated characters)

### 4. Hidden Unicode Character Detection

| Category | Characters |
|----------|------------|
| **Zero-Width** | U+200B (ZWSP), U+200C (ZWNJ), U+200D (ZWJ), U+FEFF (BOM), U+2060 (WJ) |
| **Bidi Overrides** | U+202A-E (LTR/RTL embed/override/pop), U+2066-69 (isolates) |
| **Control** | C0/C1 control chars (excluding tab/newline) |
| **Whitespace** | Unusual spaces (nbsp, em-space, etc.) |
| **Format** | Unicode format chars (Cf category) |

### 5. Diff Mode

Compare original vs sanitized with colored output:
```bash
# Side-by-side changes
safeanalyze diff ./repo ./safeanalyze-out/sanitized

# Unified diff format
safeanalyze diff ./repo ./safeanalyze-out/sanitized --unified
```

### 6. Git Clone Wrapper

Clone, analyze, and optionally delete the raw repo:
```bash
# Clone, ingest, delete raw
safeanalyze clone https://github.com/user/repo.git

# Clone, ingest, keep raw
safeanalyze clone https://github.com/user/repo.git --keep-raw
```

### 7. Sandbox Launch

Launch Claude in an isolated environment after ingestion:
```bash
# Docker (Windows, macOS, Linux)
safeanalyze ingest ./repo --sandbox

# Firejail (Linux only)
safeanalyze ingest ./repo --sandbox  # auto-detects firejail on Linux
```

## Configuration

`safeanalyze.yaml`:

```yaml
scanners:
  - name: trufflehog
    command: "trufflehog filesystem {path} --json"
    enabled: true
    fail_on_findings: true
  - name: semgrep
    command: "semgrep --config=auto {path} --json"
    enabled: false
    fail_on_findings: false

sanitization:
  strip_comments: true
  remove_non_ascii: true
  max_file_size_bytes: 50000
  max_lines_per_file: 500
  allowed_extensions:
    - .go
    - .py
    - .js
    - .ts
    - .rs
  excluded_paths:
    - .git
    - node_modules
    - vendor

hidden_chars:
  enabled: true
  categories:
    - zero_width
    - bidi
    - control
  fail_on_findings: true

entropy:
  enabled: true
  threshold: 4.5
  min_length: 20
  fail_on_findings: false

yara:
  enabled: true
  fail_on_findings: true

output:
  format: markdown        # markdown, json, plain
  single_file: false
  include_file_tree: true
  out_dir: ./safeanalyze-out

sandbox:
  mode: none              # none, docker, firejail
  docker_image: alpine:latest
  firejail_profile: default
```

## External Scanner Dependencies (optional)

| Tool | Install | Purpose |
|------|---------|---------|
| `trufflehog` | `brew install trufflesecurity/trufflehog/trufflehog` | Secret detection |
| `semgrep` | `pip install semgrep` | Static analysis |
| `git` | System package | Clone wrapper |

Built-in YARA, entropy, and hidden-char detection require no external tools.

## Sandbox Recommendation

Even after sanitization, run Claude in isolation:

```bash
# Linux (firejail)
firejail --net=none --private --read-only=/home/user claude

# Docker (cross-platform)
docker run --rm -it --network none --read-only \
  -v $(pwd)/safeanalyze-out/ingest:/ingest:ro \
  claude-sandbox

# Or let safeanalyze launch it:
safeanalyze ingest ./repo --sandbox
```

## Architecture

```
cmd/           Cobra CLI commands
pkg/
  aststrip/    Go AST-based comment stripping + regex fallback
  config/      YAML configuration loading
  diff/        Line-by-line diff engine (LCS-based)
  entropy/     Shannon entropy + base64/hex detection
  hiddenchars/ Unicode suspicious character detection
  ingest/      Markdown/JSON/plain formatter for AI consumption
  sandbox/     Cross-platform sandbox abstraction (Docker/Firejail)
  sanitize/    Comment stripper + ASCII enforcer + size limits
  scanner/     External tool orchestrator
  yara/        Pure-Go YARA-like rule engine
  utils/       Filesystem helpers
```

## License

MPL
