# safeanalyze v0.3.7

A Go CLI tool that sanitizes and scans untrusted code repositories **before** feeding them to AI assistants. Implements defense-in-depth inspired by [Zones of Distrust](https://github.com/bluvibytes/zone-of-distrust).

## Why?

Prompt injection via malicious code is real. A repo can contain:
- Hidden comments telling an LLM to "ignore all prior instructions"
- Zero-width Unicode characters hiding payload instructions
- Bidirectional text overrides that reorder displayed code
- High-entropy encoded payloads (base64, hex)
- Secrets or malware mixed with legitimate source

**safeanalyze** runs a security pipeline so AI assistants never see raw, unverified code.

## What's new in v0.3.7

- **Parallel YARA file scanning** — scans files concurrently with a worker pool, cutting thorough-mode wall-clock time on large repos (e.g., `skylos` ~24 s → ~9 s, `InjecAgent` ~9 s → ~2.4 s) while keeping findings stable.

## What's new in v0.3.6

- **Hardened `safeanalyze clone`** — validates the destination directory and uses a `--` separator when invoking `git clone`, preventing option-injection via the URL or directory argument.
- **Cleaner default exclusions** — added common Python cache/build directories to `excluded_paths` and documented that `dependency_paths` are scanned only in thorough mode.

## What's new in v0.3.5

- **Semgrep file-count gate** — skips Semgrep on targets with fewer than 50 files, cutting fixed startup latency on small benchmark repos without losing coverage on larger codebases.
- **Reverted v0.3.4 encoded-prompt-injection expansion** — the extra base64/hex fragments added latency but no new detections on the test corpus.

## What's new in v0.3.3

- **Narrowed `template_injection` rule** — focuses on `{{...}}` and `${jndi:...}` to avoid false positives from ordinary `${variable}`, `<%...%>`, and `#{...}` interpolations in code and tests.

## What's new in v0.3.2

- **Semgrep + TruffleHog by default in thorough mode** — expanded external-scanner coverage.
- **ML opt-in** — stochastic ONNX classifier disabled by default until a sub-2 GB model is validated.

## What's new in v0.3.0

- **ML stage guardrails** — configurable file-size and timeout limits prevent the stochastic stage from hanging on oversized inputs.
- **Indirect prompt-injection rules** — detect user-comment/email/web-content injections, tool/function-call payloads, and delimiter breakouts.
- **Broader "ignore" matching** — catches "ignore every instruction" and similar variants.
- **Red-team and premortem scripts** — `scripts/redteam.sh` and `scripts/premortem.sh` support the autoresearch improvement loop.

## What's new in v0.2.7

- **Secret redaction in reports** — detected credential-like values are masked in JSON/SARIF/Markdown/HTML output so reports can be committed to public branches without leaking keys from test fixtures.

## What's new in v0.2.6

- **Binary-file skipping** — images, PDFs, archives, fonts, and compiled artifacts are skipped by the text-based checks, eliminating noise from control bytes in binary data.

## What's new in v0.2.5

- **External scanners actually run in thorough mode** — fixed a registry reuse bug so enabled third-party scanners are executed.
- **Clone URL validation** — rejects git option injection and shell metacharacters in `safeanalyze clone`.
- **ONNX model install reuse** — downloads the model and tokenizer artifacts the ML stage expects.
- **Severity-aware report capping** — `output.max_findings` keeps critical/high findings first.
- **Expanded prompt-injection YARA rules** — detects `strictly adhere to...`, `you are now`, `pretend you are`, `do not mention`, `I am the developer`, and plural instruction overrides.
- **Version-pinned, parallel scanner installer** — external scanners are cloned and built concurrently at known-good refs.
- **Real `alexh-scrt/prompt-injection-scanner` bridge** — parses the scanner's actual JSON schema.
- **Report duration in milliseconds** — `duration_ms` replaces `completed_at`.
- **Dependency-path config** — `node_modules` and `vendor` are scanned only in thorough mode by default.

## Pipeline

```
Untrusted Repo
    |
    v
[Clone / Read]       <-- git clone wrapper with auto-cleanup
    |
    v
[Deterministic Checks]  <-- YARA rules, entropy, hidden chars (parallel)
    |
    v
[Stochastic ML Check]   <-- ONNX prompt-injection classifier (optional)
    |
    v
[External Scanners]     <-- semgrep, bumblebee, prompt-injection-scanner,
                              gitleaks, trufflehog (optional, source-installed)
    |
    v
[Sanitization]       <-- AST-aware comment stripping, non-ASCII removal,
                           size limits, extension filtering
    |
    v
[Diff Review]        <-- colored diff showing exactly what was removed
    |
    v
[Reporting]          <-- SARIF + Markdown + HTML dashboard + JSON
    |
    v
[Sandbox Launch]     <-- Docker or Firejail isolated AI session
```

## Install

```bash
git clone https://github.com/andreas2301/safeanalyse.git
cd safeanalyse
go build -o safeanalyze .
```

Or with `go install`:
```bash
go install github.com/andreas2301/safeanalyse@latest
```

Install optional scanner/model dependencies from source:
```bash
./safeanalyze install --all
./safeanalyze install model
```

## Quick Start

```bash
# Version
./safeanalyze --version

# Create a default config
./safeanalyze init

# Full thorough pipeline on a repo
./safeanalyze scan ./my-suspicious-repo --mode thorough

# Fast scan (YARA + hidden chars only, ~0.6 ms per typical payload)
./safeanalyze scan ./my-suspicious-repo --mode fast

# Inspect a single payload for Squid/reverse-proxy integration
./safeanalyze inspect --body < request.txt
./safeanalyze inspect --squid  # OK/ERR format

# Sanitize only
./safeanalyze sanitize ./my-suspicious-repo

# Full ingest pipeline
./safeanalyze ingest ./my-suspicious-repo

# Clone + analyze a remote repo
./safeanalyze clone https://github.com/user/repo.git
```

## Commands

| Command | Description |
|---------|-------------|
| `init [path]` | Create a `safeanalyze.yaml` config file |
| `install --all` | Clone/build external scanners and download ML model |
| `install model` | Download the prompt-injection ONNX model |
| `inspect` | Fast stdin/file inspection for reverse proxies (Squid helper) |
| `scan <path>` | Run security checks and emit SARIF/Markdown/HTML/JSON reports |
| `sanitize <src> [dst]` | Strip comments (AST-aware), remove non-ASCII, enforce limits |
| `ingest <path>` | Full pipeline: scan → sanitize → format for AI |
| `diff <orig> <sanitized>` | Show colored diff of what sanitization removed |
| `clone <url> [dir]` | Clone repo, run ingest, auto-delete raw clone |

## Scan modes

### Fast mode (`--mode fast`)

- **Checks:** built-in YARA rules + hidden Unicode characters only.
- **Latency:** ~1 ms per typical HTTP request payload.
- **Use case:** inline reverse-proxy inspection (Squid external ACL helper), request/response filtering.
- **No:** ML model, external scanners, entropy analysis, file walking beyond the given payload.

### Thorough mode (`--mode thorough`)

- **Checks:** fast checks + entropy analysis + stochastic ONNX prompt-injection classifier + optional external scanners.
- **Latency:** seconds to minutes depending on repository size.
- **Use case:** pre-ingestion security review of an entire repository.

## Features

### 1. Built-in YARA-like Rule Engine

Pure-Go regex rule engine with embedded detection patterns:

| Rule | Severity | Detects |
|------|----------|---------|
| `prompt_injection_comment` | critical | "ignore previous instructions", "system prompt", "DAN mode", "jailbreak", "override your safety" |
| `obfuscated_javascript` | high | eval(Function(...)), String.fromCharCode, atob, hex escapes |
| `suspicious_shell` | high | curl \| bash, wget \| bash, netcat reverse shells |
| `credential_hardcode` | medium | password=, api_key=, secret=, AWS keys |
| `suspicious_imports` | medium | subprocess, child_process, urllib requests |
| `data_exfiltration` | high | fetch to external URLs, axios post, XMLHttpRequest |
| `data_exfiltration_email` | high | "retrieve ... and email to ...", exfiltration via email |
| `account_access_request` | medium | "access my account", "retrieve my payment history" |
| `output_constraint` | medium | "output only", "do not mention warnings", "no disclaimer" |
| `system_boundary` | critical | `<system>`, `[system]`, `system_instruction` markers |
| `template_injection` | medium | `{{...}}` templates, `${jndi:...}` (GitHub Actions, Jinja, Log4j-style) |
| `indirect_prompt_injection` | high | user-comment/email/web-content injections, delimiter breakouts |
| `encoded_prompt_injection` | high | base64/hex/URL-encoded injection keywords |
| `backdoor_indicator` | critical | reverse_shell, bind_shell, keylogger, rootkit |

### 2. Entropy Analysis

Detects high-entropy strings that may be encoded secrets or payloads:
- **Shannon entropy** scoring (configurable threshold)
- **Base64 blob detection** with validation
- **Hex blob detection**
- String-literal-aware scanning
- Skips files larger than `entropy.max_file_size_bytes`

### 3. Hidden Unicode Character Detection

| Category | Characters |
|----------|------------|
| **Zero-Width** | U+200B (ZWSP), U+200C (ZWNJ), U+200D (ZWJ), U+FEFF (BOM), U+2060 (WJ) |
| **Bidi Overrides** | U+202A-E, U+2066-69 |
| **Control** | C0/C1 control chars (excluding tab/newline) |
| **Whitespace** | Unusual spaces (nbsp, em-space, etc.) |
| **Format** | Unicode format chars (Cf category) |

### 4. Stochastic Prompt-Injection Classifier

Optional ONNX classifier (`protectai/deberta-v3-base-prompt-injection`, ~738 MB) using [hugot](https://github.com/knights-analytics/hugot). Runs after sanitization so it classifies what the AI will actually see.

### 5. External Scanner Bridges

| Scanner | Source | Purpose |
|---------|--------|---------|
| Semgrep | `semgrep/semgrep` | SAST + prompt-injection rules |
| Bumblebee | `perplexityai/bumblebee` | Package/extension inventory |
| prompt-injection-scanner | `alexh-scrt/prompt-injection-scanner` | GitHub Actions prompt injection |
| Gitleaks | `gitleaks/gitleaks` | Secret detection |
| TruffleHog | `trufflesecurity/trufflehog` | Secret detection + verification |

All can be installed automatically via `safeanalyze install`.

### 6. Reporting

Every scan produces a unified `Report` with findings and writes:
- `safeanalyze.sarif` — SARIF v2.1.0 for security tooling
- `safeanalyze.md` — human-readable Markdown summary
- `safeanalyze.html` — self-contained dashboard
- `safeanalyze.json` — full machine-readable report

Reports include `safeanalyze_version`, `scan_mode`, and `duration_ms` metadata.

### 7. AST-Aware Comment Stripping

For **Go files**, uses `go/ast` to precisely remove comments without touching string literals. For other languages, uses an enhanced regex fallback.

### 8. Sandbox Launch

Launch Claude or another AI assistant in an isolated environment after ingestion:
```bash
# Docker (cross-platform)
safeanalyze ingest ./repo --sandbox

# Firejail (Linux only)
safeanalyze ingest ./repo --sandbox
```

## Configuration

`safeanalyze.yaml`:

```yaml
scanners:
  - name: semgrep
    command: "semgrep --config=p/security-audit {path} --json"
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
    - target
    - build
    - dist
    - .venv
  dependency_paths:
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
  max_file_size_bytes: 5242880
  fail_on_findings: false

yara:
  enabled: true
  fail_on_findings: true

ml:
  enabled: false
  threshold: 0.5
  batch_size: 4

output:
  formats:
    - markdown
    - html
    - sarif
    - json
  out_dir: ./safeanalyze-out
  max_findings: 10000

sandbox:
  mode: none              # none, docker, firejail
  docker_image: alpine:latest
  firejail_profile: default
```

## Squid Integration

`safeanalyze inspect` is designed to be used as a Squid `external_acl_type` helper:

```squid
external_acl_type prompt_injection_check %SRC %DST %METHOD %URI /usr/local/bin/safeanalyze inspect --squid
acl prompt_injection_detected external prompt_injection_check
http_access deny prompt_injection_detected
```

The helper reads payloads from stdin, runs the fast check suite, and returns `OK` or `ERR <rule>`.

## Architecture

```
cmd/              Cobra CLI commands
pkg/
  checks/         Pluggable pipeline stages
    yara/         Regex rule engine
    entropy/      Shannon entropy + base64/hex detection
    hiddenchars/  Unicode suspicious character detection
    ml/           ONNX prompt-injection classifier
    external/     External scanner bridges
  pipeline/       DAG orchestration engine
  report/         SARIF / Markdown / HTML / JSON writers
  install/        Source-based scanner/model installer
  sanitize/       Comment stripper + ASCII enforcer + limits
  sandbox/        Cross-platform sandbox abstraction
  ingest/         Markdown/JSON/plain formatter for AI consumption
  config/         YAML configuration loading
  version/        Release version constant
  utils/          Filesystem helpers
```

## Versioning

Releases are tagged with `vMAJOR.MINOR.PATCH`. The current version is embedded in the binary (`safeanalyze --version`) and in every scan report's metadata.

## License

MPL
