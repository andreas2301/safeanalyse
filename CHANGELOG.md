# Changelog

All notable functional and non-functional changes to `safeanalyze` are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.3.8] — 2026-07-15

### Functional

- **ML stage robustness for smaller models:**
  - `injectionProbability` now searches for the `INJECTION` label across all classification outputs instead of assuming the first output is the injection logit. This makes the stage compatible with models that return labels in either order.
  - Reduced `maxChunkChars` from 1800 to 1000 so smaller models with a 512-token context window no longer fail on long subword-heavy chunks.

### Non-functional

- The stochastic ML stage remains disabled by default until a sub-2 GB model is validated that improves detection on the test corpus without excessive false positives.
- Default corpus scan (ML disabled) remains unchanged: 8845 findings.

## [0.3.7] — 2026-07-15

### Functional

- **Parallel YARA file scanning:** The YARA pipeline stage now scans files concurrently using a worker pool sized to `runtime.NumCPU()`. File collection still walks the target deterministically; report output is sorted before writing so results remain stable.

### Non-functional

- Reduced thorough-scan wall-clock time on large repositories:
  - `repos/duriantaco-skylos`: ~24 s → ~9 s
  - `repos/uiuc-injecagent`: ~9 s → ~2.4 s
  - `repos/microsoft-bipia`: ~4.2 s → ~3 s
- Detection coverage unchanged on the test corpus.

## [0.3.6] — 2026-07-15

### Functional

- **Hardened `safeanalyze clone` against option injection:** The directory argument is now validated (no leading `-`, no shell/control characters) and `git clone` is invoked with a `--` separator so the URL and directory are always treated as positional arguments.
- **Expanded default `excluded_paths`:** Added common Python cache/build directories (`.tox`, `.pytest_cache`, `.mypy_cache`) to the default exclusion list and clarified that `dependency_paths` are scanned only in thorough mode.

### Non-functional

- Added unit tests for `validateCloneDir` and the hardened `git clone` invocation.

## [0.3.5] — 2026-07-15

### Functional

- **Semgrep file-count gate:** `semgrep` is now skipped in thorough mode when the target contains fewer than 50 regular files. This avoids Semgrep's high fixed startup cost on small benchmark repositories while preserving coverage on larger codebases.
- **Reverted v0.3.4 encoded-prompt-injection expansion:** The pre-computed base64/hex fragment additions from v0.3.4 added latency without improving detection on the test corpus, so they have been removed.

### Non-functional

- Reduced thorough-scan latency for small repositories (e.g., `lakeraai/pint-benchmark`, `alexh-scrt/prompt-injection-scanner`) with no expected change in findings.

## [0.3.4] — 2026-07-15

### Functional

- **Expanded `encoded_prompt_injection` rule:** Added pre-computed base64 and hex fragments for common prompt-injection phrases and their variants, including "ignore previous instructions", "ignore all previous instructions", "ignore every instruction", "system prompt", "developer mode", "jailbreak", "DAN mode", "disregard your instructions", "strictly adhere to...", "you are now", and "pretend you are". URL-encoded variants were also extended.

### Non-functional

- Fast mode remains deterministic and within the sub-100 ms budget because the encoded-pattern checks are pure regex additions.

## [0.3.3] — 2026-07-15

### Functional

- **Narrowed `template_injection` rule:** Removed overly broad `${...}`, `<%...%>`, and `#{...}` patterns that matched ordinary code variables, test strings, and shell/JS interpolations. The rule now focuses on `{{...}}` (Handlebars/Mustache/Jinja/GitHub Actions) and `${jndi:...}` (Log4j-style JNDI injection), which are the template syntaxes most commonly exploited in prompt-injection and CI/CD attacks.

### Non-functional

- Reduced false-positive volume in repositories with heavy templated test code (e.g., `alexh-scrt/prompt-injection-scanner`, `duriantaco/skylos`) while preserving true-positive detections in GitHub Actions workflows.

## [0.3.2] — 2026-07-15

### Functional

- **Expanded prompt-injection YARA rules:** Added deterministic detection for social-engineering exfiltration, account access, output constraints, system-boundary markers, and template/variable interpolation:
  - `data_exfiltration_email` flags requests to retrieve data and email it to an address (e.g., InjecAgent attacker instructions).
  - `account_access_request` flags requests to access user accounts, payment methods, histories, etc.
  - `output_constraint` flags instructions that suppress warnings, disclaimers, or constrain output format.
  - `system_boundary` flags `<system>`, `[system]`, `system_instruction`, and similar system-prompt boundary markers.
  - `template_injection` flags `{{...}}`, `${...}`, `<%...%>`, `#{...}`, and `${jndi:` interpolations that may inject instructions.
  - `prompt_injection_comment` extended with privilege-escalation patterns: "override your safety", "disable filters", "bypass guidelines", "you are in admin mode".

### Non-functional

- Updated unit tests for the expanded YARA rule set.
- Verified fast-mode latency remains within the sub-100 ms budget.

## [0.3.1] — 2026-07-15

### Functional

- **Semgrep and TruffleHog enabled by default in thorough mode:** Default config now runs Semgrep (`p/security-audit`) and TruffleHog alongside the built-in checks and `prompt-injection-scanner`.
- **ML disabled by default:** The stochastic ONNX classifier is opt-in until a model that fits the ~2 GB RAM budget is validated.
- **TruffleHog no longer fails the pipeline by default:** Secret findings are reported rather than blocking, matching the behavior of other external scanners.

### Non-functional

- Updated example `safeanalyze.yaml` to list all external scanners with their default enabled flags.

## [0.3.0] — 2026-07-15

### Functional

- **Instruction-to-include-code detection:** New YARA rule `instruction_to_include_code` flags natural-language directives telling an LLM to include, merge, or execute a provided code snippet. This catches the BIPIA `code_attack` indirect-injection pattern without hard-coding dataset phrases.

### Non-functional

- Red-team payload set expanded to include a delimiter-breakout and a code-snippet-instruction example.

## [0.2.9] — 2026-07-15

### Functional

- **Encoded prompt-injection detection:** New YARA rule `encoded_prompt_injection` catches base64, hex, URL-encoded, and Unicode-escape variants of injection keywords.
- **Markdown/HTML injection detection:** New YARA rule `suspicious_markdown_injection` flags links, images, and comments that may carry injected instructions.

### Non-functional

- Expanded YARA builtin rule set continues to be exercised by `scripts/redteam.sh`.

## [0.2.8] — 2026-07-15

### Functional

- **ML stage guardrails:** Added `ml.max_file_size_bytes` (default 1 MB) and `ml.timeout_seconds` (default 5 minutes) so the stochastic stage cannot hang indefinitely or ingest multi-megabyte lockfiles. ML remains optional and is disabled by default until a model that fits the ~2 GB RAM budget is validated.
- **Indirect prompt-injection rules:** New YARA rules `indirect_prompt_injection`, `llm_tool_injection`, and `delimiter_breakout` catch user-comment/email/web-content injections, tool/function-call payloads, and delimiter-breakout attempts.
- **Broader "ignore" pattern:** The `prompt_injection_comment` rule now matches "ignore every instruction", "ignore all prior instructions", and variants without requiring the literal word "previous".
- **Red-team and premortem scripts:** Added `scripts/redteam.sh` (adversarial payload sweep against `safeanalyze inspect`) and `scripts/premortem.sh` (failure-mode questionnaire) to support the autoresearch improvement loop.

### Non-functional

- Verified fast-mode latency remains ~10 ms per adversarial payload, well under the 100 ms reverse-proxy budget.
- Documented that the current default DeBERTa model consumes more than the requested ~2 GB RAM; smaller compatible models need further validation.

## [0.2.7] — 2026-07-15

### Functional

- **Report secret redaction:** Before writing JSON/SARIF/Markdown/HTML output, potential secret values in finding `match` fields are redacted. This prevents GitHub push-protection blocks when test fixtures contain fake API keys, and keeps public report branches safe.

### Non-functional

- Added `pkg/report` tests verifying that Stripe keys, GitLab PATs, and quoted credential values are redacted while non-secret matches remain readable.

## [0.2.6] — 2026-07-15

### Functional

- **Binary-file skipping:** The file walker used by YARA, entropy, hidden-char, and ML stages now skips binary files (by extension and by content heuristic). This eliminates thousands of false-positive hidden-char findings in images, PDFs, archives, and compiled artifacts.

### Non-functional

- Added unit tests for binary detection (`IsBinaryContent`) and binary skipping in `WalkDirSorted`.

## [0.2.5] — 2026-07-15

### Functional

- **External scanners in thorough mode:** Fixed `cmd/scan.go` to use the same pipeline registry that registers enabled external scanners, so Semgrep, Bumblebee, prompt-injection-scanner, Gitleaks, and TruffleHog actually run when enabled.
- **Clone URL validation:** `safeanalyze clone` now rejects git option injection (`-u`, `--upload-pack`) and shell/control characters. Only http(s), ssh, git, and scp-style URLs are accepted.
- **ONNX model install reuse:** `safeanalyze install model` now delegates to `pkg/checks/ml.DownloadModel`, avoiding duplicated file lists and downloading tokenizer/config artifacts.

### Non-functional

- Added `cmd/clone_test.go` covering valid and malicious clone URLs.
- Updated `CLAUDE.md` autoresearch loop with explicit premortem and red-team steps and security-hardening notes.

## [0.2.4] — 2026-07-15

### Functional

- **ONNX model download path:** `safeanalyze install model` now saves `model.onnx` inside `~/.safeanalyze/models/deberta-v3-base-prompt-injection/`, matching the path expected by the ML classifier stage.

### Non-functional

- README header and feature list synced to v0.2.4.

## [0.2.3] — 2026-07-15

### Functional

- **Real prompt-injection-scanner bridge:** Parser now matches the actual JSON schema emitted by `alexh-scrt/prompt-injection-scanner` (`rule_id`, `rule_name`, `file_path`, `line_number`, `matched_expression`).
- **Installer build fixes:** Python scanners install a compatible `setuptools==68.2.2` build environment; editable installs fall back to regular installs when the legacy backend is unavailable.
- **Correct scanner pins:** TruffleHog pinned to existing tag `v3.95.9`.

### Non-functional

- External scanner bridge tests updated to the real output schema.

## [0.2.2] — 2026-07-15

### Functional

- **Severity-aware report capping:** When `output.max_findings` is exceeded, findings are now capped by severity (critical first) so high-priority signals are never dropped because of noisy low-severity output.
- **Expanded prompt-injection rules:** Added detection for `strictly adhere to the following instruction`, `you are now`, `pretend you are`, `do not mention`, `I am the developer`, and the plural `ignore all previous instructions`.
- **Version-pinned scanner installer:** External scanners are cloned at known-good tags or commits and installed concurrently. Pins: Semgrep `v1.169.0`, Bumblebee `v0.1.2`, prompt-injection-scanner `33dd171b`, Gitleaks `v8.25.0`, TruffleHog `v3.105.0`.
- **Correct repository URLs:** Bumblebee now points to `perplexityai/bumblebee`; prompt-injection-scanner points to `alexh-scrt/prompt-injection-scanner`.
- **Report duration in milliseconds:** Reports expose `duration_ms` instead of `completed_at` so sub-second scan times are visible.
- **Dependency-path config in `safeanalyze.yaml`:** `node_modules` and `vendor` moved from `excluded_paths` to `dependency_paths`, matching the default behavior where they are skipped only in fast mode.

### Non-functional

- External scanner installation is now parallelized.
- Deterministic report sorting is preserved after severity-priority capping.

## [0.2.1] — 2026-07-15

### Functional

- **Scan modes:** `scan --mode fast` and `scan --mode thorough`. Fast mode runs only YARA + hidden-char checks for reverse-proxy use; thorough mode runs the full suite.
- **Dependency-path handling:** `node_modules` and `vendor` are no longer blanket-excluded. They are listed under `sanitization.dependency_paths` and are skipped only in fast mode; thorough mode scans them with normal limits.
- **Report finding cap:** New `output.max_findings` config (default 10,000) prevents multi-gigabyte reports on noisy targets.
- **Entropy limits:** Entropy analysis now caps findings per file at 1,000 and ignores strings longer than 1,000 bytes, eliminating lockfile blow-up.
- **Entropy file-size gate:** Entropy stage respects `entropy.max_file_size_bytes` (default 5 MB) and skips oversized files with an error record.

### Non-functional

- Reduced `duriantaco/skylos` thorough-scan time from ~576 s to ~86 s.
- Reduced `uiuc-kang-lab/InjecAgent` thorough-scan time from ~15 s to ~5 s.
- All report files now fit within GitHub's 100 MB limit.
- Reports include `safeanalyze_version` and `scan_mode` metadata.

## [0.2.0] — 2026-07-15

### Functional

- **Pluggable pipeline architecture:** New `pkg/pipeline` DAG scheduler with parallel independent stages and deterministic ordering.
- **Unified report model:** New `pkg/report` with `Report`, `Finding`, and `Summary` types consumed by all checks and writers.
- **External scanner bridges:** Added wrappers for Semgrep, Perplexity Bumblebee, `prompt-injection-scanner`, Gitleaks, and TruffleHog in `pkg/checks/external`.
- **Source-based installer:** New `safeanalyze install` command and `pkg/install` for cloning/building scanners and downloading the ONNX model.
- **Stochastic ML check:** Added ONNX prompt-injection classifier using `protectai/deberta-v3-base-prompt-injection` via `hugot`.
- **Multi-format reporting:** SARIF v2.1.0, Markdown summary, self-contained HTML dashboard, and JSON output.
- **Fast inspect command:** New `safeanalyze inspect` for Squid `external_acl` helpers, returning `OK`/`ERR`.
- **Versioning:** Added `pkg/version`, `--version` flag, and version metadata in reports.
- **Bug fixes:** Corrected `pkg/sandbox` `absPath()` to return the actual absolute path; removed duplicate `cloneCmd` registration.

### Non-functional

- Existing YARA, entropy, and hidden-char checks moved to `pkg/checks/*` and adapted to the `pipeline.Stage` interface.
- Deterministic file iteration and stable report sorting across concurrent execution.
- Added pipeline engine unit tests covering topological order, cycle detection, unknown dependencies, and error propagation.

## [0.1.0] — pre-refactor baseline

### Functional

- Sequential scanner pipeline: external scanners → YARA rules → entropy → hidden chars.
- Built-in YARA-like rule engine for prompt injection, obfuscation, shells, credentials, etc.
- Shannon entropy, base64/hex blob detection.
- Hidden Unicode character detection (zero-width, bidi, control, whitespace).
- Sanitization: AST-aware comment stripping, non-ASCII removal, size limits.
- Markdown/JSON/plain output formatting.
- Docker/Firejail sandbox wrappers.
- Git clone wrapper with optional cleanup.

### Non-functional

- Cobra-based CLI.
- YAML configuration.
- No report model, no parallel execution, no version metadata.
