# safeanalyze — 6-Role Warroom Council Assessment

> **Date:** November 2026
> **Baseline Audited:** `andreas2301/safeanalyse` (`v0.4.0`, commit `a10be35`)
> **Mandate:** Evaluate whether to build or buy, where `safeanalyze` fits, what architectural/engineering/security/detection/operational gaps prevent enterprise-grade adoption, whether to expand evaluation sets, and how to improve the overall approach.

---

## 1. Council Composition & Methodologies Applied

To avoid single-discipline blind spots, four core roles were augmented with two domain specialists directly matched to the repository's bottlenecks (base-rate collapse / eval governance, and CI/ICAP operational readiness):

| Role | Authoritative Methodologies Used | Core Question Owned |
|---|---|---|
| **Enterprise Architect (EA)** | TOGAF ADM (Capability Mapping & Gap Analysis), Wardley Mapping (Value Chain Evolution), Geoffrey Moore Core vs. Context, Competitive Half-Life & TCO, Gartner TIME | *"Build vs. Buy? Where does our solution create durable value vs. commodity waste?"* |
| **Solution Architect (SA)** | C4 Model, ATAM Quality Attribute Scenarios, Enterprise Integration Patterns (Plugin/Sidecar Contracts), ADRs | *"How should components, external scanners, and deployment topologies compose cleanly?"* |
| **Principal Engineer (PE)** | Static/Dynamic Correctness & Concurrency Analysis (`-race`), Performance Profiling, Modularity & Idiomatic Go Review, Testability | *"What bugs, races, latency bottlenecks, and debt block reliability and maintainability?"* |
| **Security Engineer (SecEng)** | STRIDE (Scanner-as-Target Threat Model), MITRE ATLAS, OWASP Supply-Chain Defense | *"How can a hostile repository or HTTP response compromise, escape, DoS, or leak data from safeanalyze itself?"* |
| **ML / Detection & Evaluation Engineer** | Base-Rate Calibration, Contamination-Free Eval Governance, Synthetic Benchmark Design | *"Where is deterministic detection effective vs. mathematically doomed, and how should we evaluate?"* |
| **Platform / DevSecOps / SRE Engineer** | Google SRE Production Readiness Review (PRR), Observability & Rollout Safety, Supply-Chain Lifecycle | *"Can teams actually deploy, monitor, debug, and trust this in CI and production proxies without alert fatigue?"* |

---

## 2. Independent Role Findings (Codebase-Verified)

### 2.1 Enterprise Architect (EA)

**Capability Decomposition (Wardley Evolution × Core vs. Context):**

| Capability | Wardley Stage | Core vs. Context | Decision |
|---|---|---|---|
| Secret scanning (`TruffleHog`, `Gitleaks`) | Commodity / Product | Context | **Integrate (OSS)** — never rebuild detector databases |
| General SAST (`Semgrep` / `Opengrep`) | Product | Context | **Integrate (OSS)** — mature AST/taint across 30+ languages |
| Compromised-package advisory inventory (`Perplexity Bumblebee`, `OSV-Scanner`, `Trivy`) | Product | Context | **Integrate (OSS)** — Apache 2.0 single binary + community threat-intel catalogs |
| Semantic/multilingual/paraphrase prompt-injection detection (`Prompt Guard 2`, LLM judge) | Custom → Product | Context (model) / Core (integration) | **Integrate via local ONNX/sidecar plugin** — do not chase with regexes |
| **Pre-AI Intake Orchestrator + Cross-Layer Intent Analysis + 11-Question Tiered Verdicts** | **Genesis / Custom** | **Core Differentiator** | **BUILD & OWN** — MalSkillBench (June 2026, arXiv:2606.07131v3) showed across 12 tools that supply-chain scanners and prompt-injection detectors each see only half of an agent/repo attack and naive unions over-trigger. Unifying code intent, agent instruction surfaces, and external scanners into ranked Tier A/B/C verdicts inside one offline Go binary is a genuine unmet enterprise need. |
| **Fast Deterministic L0 Tripwire** (Unicode smuggling, chat-template forgery, literal overrides) | Product | Context-enabler | **KEEP & HARDEN RUNTIME** — zero-GPU, <10 ms, catches Unicode smuggling ML models normalize away; stop investing engineering hours in narrow regexes. |

**EA Build-vs-Buy Verdict:**
- **Do NOT replace everything with a commercial SaaS** (e.g., Lakera Guard, Socket.dev, Cisco AI Defense cloud) if offline/zero-egress code privacy is a hard constraint — commercial SaaS sends untrusted/proprietary code outside your trust boundary and incurs compounding per-seat/per-call TCO.
- **Do NOT build commodity detectors in-house.** Focus 100% of custom engineering on the **orchestrator control plane, intent/triage correlation, declarative plugin integration (Bumblebee, Semgrep, TruffleHog, Gitleaks, SARIF/JSON sidecars), and runtime hardening**.

---

### 2.2 Solution Architect (SA)

**C4 & Integration Contract Audit:**
1. **Split-Brain External Scanner Architecture:** `safeanalyze.yaml` defines `scanners[].command` and `scanners[].fail_on_findings`, but `pkg/checks/external/registry.go` ignores them and hardcodes 5 adapters, while `cmd/ingest.go` invokes a separate legacy `pkg/scanner/runner.go` that flags any output containing the word `"result"`.
2. **Broken Perplexity Bumblebee Contract (`pkg/checks/external/bumblebee.go`):**
   - Invokes `bumblebee scan --profile deep --root <target> --findings-only` **without `--exposure-catalog`** — real Bumblebee v0.1.2 rejects `--findings-only` without a catalog.
   - Unmarshals fabricated JSON fields (`rule`, `title`, `message`, `file`) instead of Bumblebee's actual NDJSON finding schema (`record_type`, `finding_type`, `catalog_id`, `catalog_name`, `ecosystem`, `package_name`, `version`, `source_file`, `evidence`, `confidence`) and fails to ignore `package`/`scan_summary` records.
3. **Orphaned External Findings in Triage (`pkg/triage/triage.go`):** `bumblebee` and `semgrep` are unmapped in `Questions` and `legacyTier()` — even a `critical` compromised-package exposure match from Bumblebee cannot reach Tier A or answer `QDeps`.
4. **Missing Plugin & Report-Merge Contracts:** Plugging in OSV-Scanner, Trivy, GuardDog, Cisco Skill Scanner, or `safeanalyze-llm` (`docs/LLM_SCANNER_SPEC.md`) currently requires editing Go source code and lacks a `safeanalyze merge` CLI subcommand.
5. **Fast-Path Architectural Coupling (`cmd/inspect.go:117-135`):** `inspect`, `icap`, and `eval` write every payload to a temp directory on disk and construct a new `pipeline.Engine` + `yara.Stage` per payload just to reuse file-based stages.

---

### 2.3 Principal Engineer (PE)

**Verified Defects & Debt:**
1. **Broken `master` Test Suite:** `go test ./...` fails on fresh clone (`TestMaliciousFixturesReachTierA/vscode-folder-open`) because `.gitignore` ignores `.vscode/`, silently dropping the fixture's `.vscode/tasks.json`.
2. **Data Race & Unbounded Cache (`pkg/checks/intent/common.go:51-56,136`):** `File.Lines()` mutates `f.lines` without synchronization across concurrent intent stages sharing `[]*File`; global `cache` map never evicts or checks staleness.
3. **Goroutine Leak on Cancellation (`pkg/pipeline/engine.go:70-76`):** `cond.Wait()` is never woken by `cond.Broadcast()` when `ctx` is canceled.
4. **Per-Request Regex Recompilation (`cmd/inspect.go:119`, `pkg/checks/yara/engine.go:50`):** All 114+ regexes are recompiled on every payload (~5 ms floor wasted on every ICAP/eval/inspect call).
5. **Gitleaks Exit-Code Bug (`pkg/checks/external/gitleaks.go:34`):** Gitleaks returns exit status `1` when leaks are found; `cmd.Output()` treats non-zero exit as a hard error and drops all findings.
6. **Unregistered Command (`cmd/clone.go`, `cmd/root.go:58`):** `cloneCmd` is implemented and tested, yet never registered via `rootCmd.AddCommand`.
7. **Non-Streaming Stdin (`cmd/inspect.go:51`):** `os.ReadFile("/dev/stdin")` blocks until EOF, breaking interactive Squid `external_acl_type` helpers and non-POSIX systems.
8. **SARIF 2.1.0 Spec Issues (`pkg/report/sarif.go`):** Missing `driver.version`, placeholder `informationUri` (`github.com/user/safeanalyze`), missing `partialFingerprints`, `RuleIndex: -1` on empty `RuleID`, hardcoded `ExecutionSuccessful: true`.
9. **Test Coverage & CI Gaps:** 9 packages have zero test files (`config`, `sanitize`, `sandbox`, `utils`, `entropy`, `aststrip`, `diff`, `ingest`, `scanner`) and `.github/workflows/` has no Go test/vet workflow.

---

### 2.4 Security Engineer (SecEng — STRIDE Scanner-as-Target)

| STRIDE Threat | Location | Verified Vulnerability | Severity |
|---|---|---|---|
| **Information Disclosure** | `pkg/report/writer.go:29-41`, `model.go:46-51`, `gitleaks.go:78` | When findings exceed `max_findings`, `WriteAll` redacts `report` instead of the copied `reportToWrite` -> **unredacted secrets written to disk**. `Redact()` ignores `Message` (where Gitleaks injects raw secrets), `Description`, `Attributes`, `Evidence`, and misses PEM/OpenAI/Anthropic/Slack keys. | **Critical** |
| **Information Disclosure / DoS** | `pkg/checks/yara/engine.go:783-820`, `pkg/sanitize/sanitizer.go:75-78,169`, `pkg/utils/fs.go:26` | Symlinks are followed into host filesystem (`/etc/passwd`, `~/.ssh`, `/dev/zero`); broken symlinks abort the whole scan; no file-size cap before `os.ReadFile`; `SanitizeDir` copies non-allowlisted files (`CLAUDE.md`, `.md`, `.json`, symlinks) **verbatim** into AI ingest output. | **High** |
| **Tampering / Supply Chain** | `pkg/checks/external/semgrep.go:36`, `trufflehog.go:40`, `cmd/clone.go:57` | Semgrep pulls unpinned remote GitHub `main` rule URL on every scan; TruffleHog performs live network verification by default; `git clone` lacks `-c core.hooksPath=/dev/null -c core.fsmonitor=false -c core.symlinks=false -c filter.lfs.smudge=`; missing `--` before `target`/`go mod download` args. | **High** |
| **SSRF / DoS** | `pkg/checks/intent/netdeps.go:242,408-468` | `fetchTarball` fetches arbitrary `https://` URLs from untrusted `package-lock.json` (allowing loopback/RFC1918/link-local SSRF & redirects) and decompresses tarballs without total decompressed byte or entry count caps (gzip/tar bomb OOM). | **High** |
| **Information Disclosure / DoS** | `cmd/inspect.go:131`, `pkg/icap/server.go:110` | ICAP writes every HTTP response body to world-readable `/tmp/safeanalyze-inspect-*/payload.txt` (`0644`); `icap.Server.Serve()` accepts unbounded concurrent connections. | **High** |

---

### 2.5 ML / Detection & Evaluation Engineer

1. **Keep External Eval Sets Untouched:** `testdata/eval/{llmail,browsesafe,toolresp,deepset,nemotron,spanish,multilingual}*` are external pinned dependencies and must remain untouched to preserve historical comparability and avoid data contamination.
2. **Add Committed Synthetic & Intent Regression Suites:**
   - Commit an offline synthetic prompt-injection & hard-benign suite (`testdata/synthetic/prompt_injection_suite.jsonl`) covering overrides, smuggling, exfiltration, obfuscation, paraphrases, and realistic security/engineering benign prose — runnable in CI with zero network fetches.
   - Expand `testdata/intent/` with fixed `vscode-folder-open` (malicious + benign twin) and 2026 AI-agent supply-chain scenarios (`mcp-unpinned-rce`, `skill-mixed-attack` + benign twins).
3. **Solve Base-Rate Collapse Operationally, Not With More Regexes:** At 9–12% FPR on HTML/JSON, blocking inline proxy traffic yields <1% precision at realistic prevalence. Add `--log-only` and `--block-content-types` to ICAP so teams block only high-precision content types while logging HTML/JSON telemetry or routing uncertain items to an optional local classifier/sidecar (`Prompt Guard 2 22M` / `safeanalyze-llm`).

---

### 2.6 Platform / DevSecOps / SRE Engineer

1. **Production Readiness Gaps:** Missing Go CI workflow, missing `safeanalyze doctor` pre-flight check, missing per-scanner execution timeouts, missing ICAP `MaxConnections` + `/healthz` + `/readyz` + Prometheus `/metrics`, and missing `--log-only` canary rollout mode.
2. **External Scanner Lifecycle:** `safeanalyze install bumblebee` already clones `threat_intel/` under `~/.safeanalyze/tools/bumblebee/threat_intel`, so `safeanalyze` should auto-resolve it and surface its presence/status in `doctor`.

---

## 3. Warroom Council Synthesis & Key Debates Resolved

1. **Debate — Single Repo vs. Two Products (EA vs. PE):**
   *Resolution:* Keep one single-binary Go codebase (minimizes packaging/operational TCO), but separate Track 1 and Track 2 cleanly in architecture (`pkg/checks/external` + `pkg/checks/intent` + `pkg/triage` vs. `pkg/fastinspect` + `pkg/icap`) and ship two focused, independently reviewable roadmap branches/PRs.
2. **Debate — Custom Plugin Shell Commands vs. Security (SA vs. SecEng):**
   *Resolution:* Never pass `safeanalyze.yaml` `command` strings to `sh -c`. Tokenize into argv safely (supporting quoted arguments without shell expansion), substitute `{path}` as a single argv element, enforce `--` on built-ins, and require explicit `timeout_seconds`.
3. **Debate — Should We Tune Regexes on the New Synthetic Set? (PE vs. ML-Eval):**
   *Resolution:* **No.** Keep regex detection semantics unchanged in this cycle so `safeanalyze gate` TP/FP metrics stay byte-identical while fast-path latency drops dramatically and Track 1 orchestration reaches enterprise grade. Use the synthetic set as a transparent capability/regression dashboard in CI.

---

## 4. Beyond Code: What Else Should You Do to Improve the Overall Approach?

1. **Adopt Threat-Intel Catalog Automation for Bumblebee:** Schedule a CI job (or cron) to mirror/pin vetted `perplexityai/bumblebee/threat_intel/*.json` catalogs into an internal artifact store and pass `--exposure-catalog` in CI scans — turning Bumblebee from a static binary into a live supply-chain response sensor.
2. **Pair Scanning with Architectural Isolation:** Follow CaMeL / least-privilege agent design (no write/network/secret tools for agents reading untrusted repos) — treat a clean `safeanalyze` report as *reduced risk*, never as *proof of safety*.
3. **Run One Pilot Spike on Local Semantic Sidecar:** Using the new declarative `safeanalyze-json` plugin contract + `safeanalyze merge`, benchmark `Llama-Prompt-Guard-2-22M` (ONNX) in a sidecar against the sealed eval sets before committing to GPU/LLM-judge spend.
4. **Enforce Branch Protection & GitHub Code Scanning:** Require `.github/workflows/ci.yml` + `osengineer-gate.yml` on `master`, and upload `safeanalyze.sarif` via `github/codeql-action/upload-sarif` so Tier A/B findings appear natively on PR diffs with deduplicated `partialFingerprints`.
5. **Archive Stale Report Branches:** Clean up the 35+ legacy `report-*` branches noted in `CLAUDE.md` round 4 process review (keeping one report branch per completed round or using CI artifacts).
