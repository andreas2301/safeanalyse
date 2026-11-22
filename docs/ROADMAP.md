# safeanalyze — Enterprise Roadmap (Goal / State / Delta)

> **Status:** Adopted from the 6-Role Warroom Council review ([`docs/WARROOM_COUNCIL.md`](WARROOM_COUNCIL.md)), November 2026.
> **Scope:** Independent evaluation of Track 1 (Pre-AI Repository & Supply-Chain Threat Scanner Orchestrator) and Track 2 (Inline AI Proxy / Fast-Path Tripwire & ICAP Service), unified under one single-binary, offline-first Go platform.

---

## Strategic Decision Summary (Build vs. Buy & Where to Use)

| Track | Strategic Role | Build vs. Buy Verdict | Where to Deploy |
|---|---|---|---|
| **Track 1: Pre-AI Repo & Supply-Chain Orchestrator** (`scan --mode thorough`, Intent Analysis, External Scanners, Triage/Verdicts, SARIF/CI Gate) | **Core differentiator** — cross-layer code + agent-instruction + supply-chain orchestration with offline audit trail | **HYBRID (Build Control Plane + Intent/Triage; Integrate OSS Specialist Scanners)** — Do not rebuild secret scanning, SAST, or advisory catalogs. Orchestrate Perplexity Bumblebee, Semgrep, TruffleHog, Gitleaks, and any SARIF/JSON/NDJSON plugin via a declarative contract. | Pre-AI repo/skill/MCP intake gate; CI/PR security gate (GitHub Code Scanning SARIF); air-gapped or zero-egress enterprise security reviews. |
| **Track 2: Inline Fast-Path & ICAP Proxy** (`inspect`, `icap` RESPMOD, ≤ 100 ms deterministic tripwire) | **L0 tripwire** — cheap, deterministic, zero-GPU first filter for Unicode smuggling, chat-template forgery, and literal overrides | **KEEP & HARDEN RUNTIME; DO NOT INVEST IN MORE REGEX RULES** — Migrate semantic/paraphrase/multilingual detection to an optional local classifier (Llama Prompt Guard 2) or sidecar judge (`safeanalyze-llm`). | Squid ICAP RESPMOD in **log-only/shadow mode** for HTML & JSON (avoiding base-rate alert fatigue) and **selective blocking** on plain-text / Unicode smuggling; L0 pre-filter ahead of ML/LLM/architectural isolation (CaMeL / least-privilege tools). |

---

## Roadmap Branches

The Warroom Council identified two cohesive, independently mergeable enterprise tracks:

1. **Branch A — `feat/enterprise-threat-orchestrator` (Track 1: Control Plane, Security Hardening & Plugin Orchestration)**
   Fixes every P0 security/correctness defect, makes Perplexity Bumblebee and all external scanners production-grade, introduces a zero-code declarative external plugin architecture (SARIF / safeanalyze JSON / NDJSON), adds `merge` and `doctor` commands, adds Go CI, and commits offline synthetic & intent test suites.
2. **Branch B — `feat/enterprise-fastpath-icap` (Track 2: In-Memory Fast Path & Production-Ready ICAP Service)**
   Eliminates per-payload regex recompilation and world-readable `/tmp` payload writes via an in-memory compiled `FastInspector`, fixes streaming Squid `inspect` stdin, adds ICAP `MaxConnections` protection, `--log-only` & content-type blocking policy (solving HTML/JSON base-rate collapse), and adds `/healthz`, `/readyz`, and Prometheus `/metrics` endpoints.

---

## Goal / State / Delta Matrix

### Initiative 1 — Security & Correctness Baseline (Branch A)

- **Goal:** `safeanalyze` is safe to run on hostile, adversarial repositories with zero data races, zero host-file exfiltration, zero secret leaks in reports, and a 100% green `go test -race ./...` suite on fresh clones.
- **State (v0.4.0 on `master`):**
  - `go test ./...` fails on fresh clones: `.gitignore` ignores `.vscode/`, silently excluding `testdata/intent/malicious/vscode-folder-open/.vscode/tasks.json`.
  - Secret leakage in capped reports: `pkg/report/writer.go:29-41` calls `report.Redact()` instead of `reportToWrite.Redact()` when `len(Findings) > max_findings`; `Redact()` only scans `Match` (not `Message`, `Description`, `Attributes`, or `Verdict.Evidence`), while `gitleaks.go:78` embeds raw secrets in `Message`.
  - Symlink & OOM exposure: `yara.WalkDirSorted`, `sanitize.SanitizeDir`, and `utils.CopyDir` follow symlinks and read unbounded files (`os.ReadFile` without size checks); `SanitizeDir` copies non-allowlisted files (`.md`, `CLAUDE.md`, `.json`, symlinks) verbatim into AI ingest output.
  - `netdeps.go` SSRF & archive bomb: fetches arbitrary `https://` URLs from untrusted `package-lock.json` (including private/loopback/link-local IPs and redirects), lacks total decompressed bytes and file-count caps on tarballs, and omits `--` before `go mod download` args.
  - `cmd/clone.go` is unregistered on `rootCmd` and invokes `git clone` without disabling hooks, fsmonitor, symlinks, and LFS smudge filters.
  - Concurrency defects: `intent.File.Lines()` races across concurrent intent stages; `pipeline.Engine.Run()` leaks waiting goroutines on context cancellation (`cond.Wait()` never wakes on `ctx.Done()`).
- **Delta:**
  1. Un-ignore `testdata/intent/**/.vscode/` in `.gitignore` and commit malicious + benign `vscode-folder-open/.vscode/tasks.json` fixtures.
  2. Redact `reportToWrite` across `Match`, `Message`, `Description`, `Attributes`, and `Verdict.Evidence` with expanded secret patterns (PEM, OpenAI, Anthropic, Slack, AWS, GitHub, GitLab, secret-category findings).
  3. Skip symlinks, non-regular files, and oversized files (>16 MiB default) in `WalkDirSorted`, `SanitizeDir`, and `CopyDir`; stop copying non-allowlisted files verbatim in `SanitizeDir`.
  4. Enforce SSRF host/IP blocking (loopback, RFC1918, link-local, CGNAT, ULA, redirect validation), total tar decompression/entry limits, and `go mod download --` in `netdeps.go`.
  5. Register `cloneCmd` on `rootCmd` and harden `git clone` with `-c core.hooksPath=/dev/null -c core.fsmonitor=false -c core.symlinks=false -c filter.lfs.smudge=`.
  6. Synchronize `intent.File.Lines()` with `sync.Once`, bound/invalidate the file cache, and broadcast `cond` on `ctx.Done()` in `pipeline.Engine.Run()`.

### Initiative 2 — External Scanner & Bumblebee Control Plane (Branch A)

- **Goal:** `safeanalyze` reliably orchestrates Perplexity Bumblebee, Semgrep, TruffleHog, Gitleaks, prompt-injection-scanner, and *any* user-configured external threat scanner (OSV-Scanner, Trivy, GuardDog, Cisco Skill Scanner, `safeanalyze-llm` sidecar) offline-by-default, with bounded timeouts and unified Tier A/B/C verdict integration.
- **State (v0.4.0 on `master`):**
  - `pkg/checks/external/bumblebee.go` passes `--findings-only` without `--exposure-catalog` (rejected by real Bumblebee v0.1.2) and parses non-existent JSON fields (`rule`/`title`/`file`) instead of Bumblebee's real NDJSON schema (`record_type`, `catalog_id`, `catalog_name`, `ecosystem`, `package_name`, `version`, `source_file`, `evidence`, `confidence`).
  - `pkg/checks/external/gitleaks.go` discards all findings when Gitleaks exits with code `1` (Gitleaks' normal exit code when leaks are detected).
  - `pkg/checks/external/semgrep.go` fetches an unpinned remote `main`-branch GitHub rule URL on every scan, ignores `excludedPaths`, and lacks `--`.
  - `pkg/checks/external/trufflehog.go` performs live network credential verification by default (`no_verification: false`).
  - `pkg/checks/external/registry.go` ignores `safeanalyze.yaml` `command`, `timeout`, and custom scanner entries; adding a new scanner requires modifying Go source code.
  - `pkg/triage/triage.go` has no verdict mapping for `bumblebee`, `semgrep`, or custom external plugins (even `critical` findings stay Tier B with no verdict question).
  - No CLI command to merge companion/sidecar `safeanalyze.json` reports (blocking `docs/LLM_SCANNER_SPEC.md`).
- **Delta:**
  1. Fix Bumblebee stage: auto-discover `~/.safeanalyze/tools/bumblebee/threat_intel` or honor `exposure_catalog` config, pass `--exposure-catalog` + `--max-duration`, parse real Bumblebee NDJSON `finding` records while skipping `package`/`scan_summary`, and map to `QDeps` (Tier A for `critical`/`high`).
  2. Fix Gitleaks exit-code-1 handling and stop embedding raw secrets in `Message`; make Semgrep offline-first (no remote `main` URL by default, honor `excludedPaths` + `--`); make TruffleHog offline-first by default (`no_verification` defaults true unless `verify: true`); add per-scanner `timeout_seconds` across all external stages.
  3. Add a declarative **Generic External Plugin Stage** (`format: sarif | safeanalyze-json | ndjson`) with safe argv tokenization (no shell execution), `{path}` substitution, timeout enforcement, and configurable `question`/`tier` triage integration.
  4. Add `safeanalyze merge <report1.json> <report2.json> ... -o <out-dir>` to merge companion/sidecar reports, re-run triage, and emit unified SARIF/JSON/HTML/Markdown reports.
  5. Upgrade SARIF 2.1.0 writer (`driver.version`, repository URI, `partialFingerprints`, valid `ruleIndex`, accurate `executionSuccessful`).

### Initiative 3 — Platform / DevSecOps / SRE Readiness & Synthetic Eval Suites (Branch A)

- **Goal:** Zero-surprise enterprise CI/CD and operator ergonomics, backed by automated Go CI and offline synthetic/intent regression suites (without modifying external fetched eval sets).
- **State (v0.4.0 on `master`):**
  - `.github/workflows/` only checks Conventional Commits; no Go build/test/vet CI exists.
  - No pre-flight diagnostics command to verify installed scanners, catalogs, sandboxes, and config.
  - 9 Go packages have zero unit tests (`config`, `sanitize`, `sandbox`, `utils`, `entropy`, `aststrip`, `diff`, `ingest`, `scanner`).
  - External eval sets require network fetching; no committed, self-contained synthetic prompt-injection suite exists for offline CI regression checks, and modern 2026 MCP/agent-skill mixed attacks lack dedicated intent fixtures.
- **Delta:**
  1. Add `safeanalyze doctor [--json]` verifying config validity, Go/git/sandbox availability, external scanner binaries & versions, and Bumblebee exposure catalogs.
  2. Add `.github/workflows/ci.yml` running `go vet ./...`, `go test -race ./...`, binary build, and offline synthetic suite check.
  3. Commit `testdata/synthetic/prompt_injection_suite.jsonl` (balanced adversarial + hard-benign synthetic regression set; external fetched datasets remain completely untouched).
  4. Expand `testdata/intent/` with fixed `vscode-folder-open` (malicious + benign twin) plus `mcp-unpinned-rce` and `skill-mixed-attack` (malicious + benign twins).
  5. Add unit tests for `pkg/config`, `pkg/sanitize`, `pkg/sandbox`, `pkg/utils`, and `pkg/checks/entropy`.

### Initiative 4 — In-Memory Fast Path & Production-Ready ICAP Service (Branch B)

- **Goal:** Track 2 (`inspect`, `icap`, `eval`) executes purely in memory with precompiled rules, zero `/tmp` disk leakage, streaming Squid compatibility, DoS-bounded connections, base-rate-safe shadow/content-type policies, and SRE observability (`/healthz`, `/readyz`, `/metrics`).
- **State (v0.4.0 on `master`):**
  - `cmd/inspect.go:117-135` (shared by `icap` and `eval`) recompiles all 114+ YARA regexes and writes every HTTP response body to world-readable `/tmp/safeanalyze-inspect-*/payload.txt` (`0644`) on every single request.
  - `cmd/inspect.go:51` calls `os.ReadFile("/dev/stdin")`, blocking until EOF and breaking streaming Squid `external_acl_type` helpers.
  - `pkg/icap/server.go` accepts unbounded concurrent TCP connections (FD/goroutine exhaustion risk), lacks a `--log-only` shadow rollout mode and `--block-content-types` policy (needed to avoid 9–12% FPR base-rate collapse on HTML/JSON), and exposes no health or Prometheus metrics endpoint.
- **Delta:**
  1. Introduce `pkg/fastinspect`: immutable compiled `yara.Engine` + `hiddenchars.Detector` scanning `[]byte` in memory with zero temp files and zero per-request regex compilation; wire `inspect`, `icap`, and `eval` to it with byte-identical detection parity tests and latency benchmarks.
  2. Fix `cmd/inspect.go` to stream `os.Stdin` line-by-line via `bufio.Scanner` with immediate stdout flush in Squid mode.
  3. Harden `pkg/icap/server.go` & `cmd/icap.go`: add `MaxConnections` (`--max-connections`, default 256), `--log-only` shadow mode, `--block-content-types` selective blocking policy, and `--admin-listen` HTTP server exposing `/healthz`, `/readyz`, and Prometheus `/metrics` (zero external dependencies).
  4. Add race-checked concurrency, connection-cap, shadow-mode, content-type policy, metrics, and benchmark tests plus `.github/workflows/ci.yml` and synthetic suite verification.

---

## Sequencing & Exit Criteria

| Milestone | Branch | Verification Gate |
|---|---|---|
| **M1: Roadmap & Council Record** | Both (first commit) | `docs/ROADMAP.md` and `docs/WARROOM_COUNCIL.md` committed with Conventional Commits |
| **M2: Security & Concurrency Hardening** | `feat/enterprise-threat-orchestrator` | `go test -race ./...` 100% green; redaction/SSRF/symlink/tar-bomb/race regression tests pass |
| **M3: Bumblebee + Declarative Plugins + Merge + SARIF** | `feat/enterprise-threat-orchestrator` | Bumblebee real-schema NDJSON tests, Gitleaks exit-1 tests, generic SARIF/JSON/NDJSON plugin tests, `merge` & `doctor` tests pass |
| **M4: Synthetic Suites, Intent Fixtures & CI** | `feat/enterprise-threat-orchestrator` | `testdata/intent` recall & precision gates pass; synthetic suite runs offline via `safeanalyze eval`; `.github/workflows/ci.yml` valid |
| **M5: In-Memory Fast Path & Enterprise ICAP** | `feat/enterprise-fastpath-icap` | `pkg/fastinspect` parity + benchmark proves latency drop & zero `/tmp` writes; ICAP connection-cap, `--log-only`, `--block-content-types`, `/healthz`, `/metrics` pass `-race` |
