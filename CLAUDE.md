# CLAUDE.md — safeanalyze contributor guide

## Project purpose

`safeanalyze` is a Go CLI security scanner that checks untrusted code before it is fed to AI assistants. It supports two operational modes:

- **Fast mode** — sub-100 ms deterministic checks for reverse-proxy/Squid integration.
- **Thorough mode** — full deterministic + stochastic + external-scanner suite for repository review.

## Development conventions

- Keep changes minimal and surgical. Every changed line must trace to a functional requirement.
- Match existing Go style, error wrapping, and package layout.
- Do not add speculative abstractions or unrelated cleanups.
- All new checks must implement `pkg/pipeline.Stage`.
- All new report formats must consume `pkg/report.Report`.
- Tests run with `go test ./...`. Builds must pass with `go build ./...`.
- Commits describe the functional change; do not include review notes or process metadata.

## Release process

1. Bump `pkg/version/version.go` on every meaningful improvement.
2. Update `CHANGELOG.md` with functional and non-functional changes.
3. Update `README.md` so the header, feature list, and config example match the new version.
4. Commit improvements to `master`.
5. Tag the commit: `git tag vX.Y.Z && git push origin master vX.Y.Z`.
6. Reports generated from a versioned build must record `safeanalyze_version`, `scan_mode`, and `duration_ms` in their metadata.
7. Report branches are named `report-<devicename>-<date>`. Device names are random Deadpool-style UUID phrases (e.g., `maximum-effort-chimichanga-9f7d`).

## Autoresearch improvement loop

The project uses an iterative, measurement-driven improvement loop inspired by [karpathy/autoresearch](https://github.com/karpathy/autoresearch). The goal is to keep improving detection coverage, precision, and latency on the prompt-injection test corpus until two consecutive iterations show no measurable progress.

### Prerequisites

- Go toolchain installed (`go test ./...` and `go build .` must work).
- Git with push access to the repository.
- Test corpus available under `/tmp/`:
  - `safeanalyze-BIPIA`
  - `safeanalyze-InjecAgent`
  - `safeanalyze-pint-benchmark`
  - `safeanalyze-prompt-injection-scanner`
  - `safeanalyze-skylos`
  - `safeanalyze-doc-scenarios.html`
  - `safeanalyze-doc-webagents.html`
- External scanners installed via `./safeanalyze install --all` (Semgrep, TruffleHog, prompt-injection-scanner, etc.).
- Enough disk space and memory for the test corpus. Note: the default DeBERTa ONNX model uses ~11 GB RAM at load/inference time, so the stochastic ML stage is currently disabled by default until a sub-2 GB model is validated.
- A stable, otherwise idle machine for duration comparisons (run-to-run variance should be <10 % for large targets).

### Iteration steps

1. **Measure baseline** — run `safeanalyze scan --mode thorough` against the test corpus and record findings + durations from `safeanalyze.json` (`duration_ms`) and `duration.txt`.
2. **Premortem** — run `./scripts/premortem.sh` and ask: "If this improvement lands and the tool still fails in production, what most likely broke?" Document the biggest risks (false positives, latency blow-out, missing variants).
3. **Analyze gaps** — compare findings to known prompt-injection patterns. Look for false negatives and false positives.
4. **Hypothesize** — pick one concrete improvement: a new YARA rule, a tuned entropy threshold, a smaller/faster model, a file-size limit, or parallelism.
5. **Implement** — make the smallest change that tests the hypothesis. Do not combine multiple unrelated changes in one iteration.
6. **Red-team** — run `./scripts/redteam.sh`. Try to evade the new check with rephrased, encoded, or multi-language injections. If it is trivially bypassed, revert or harden.
7. **Evaluate** — re-run the same test corpus. Compare:
   - Total and per-target finding counts.
   - Wall-clock duration per target (`duration_ms`).
   - Fast-mode latency (`./scripts/redteam.sh`).
   - Any new errors or scanner skips.
8. **Keep or revert** — if coverage, precision, or latency improved, commit and bump the version. If not, revert and try another hypothesis.
9. **Repeat** until two consecutive iterations show no measurable improvement.

### Decision rules

- **Accept** the iteration if:
  - Finding count increases without obvious false-positive inflation, **or**
  - Latency decreases with unchanged findings, **or**
  - A security hardening fix removes a real foot-gun without regressing metrics.
- **Revert** the iteration if:
  - Finding count drops, **or**
  - Latency increases without a coverage gain, **or**
  - The change introduces non-deterministic output or new errors.
- **Stagnation** — stop the loop after two consecutive accepted/reverted iterations produce no improvement in detection coverage, precision, or latency.

### Process review

After every five iterations (or immediately after two consecutive no-improvement iterations), review the process itself:

- Are we measuring the right metric? Should we add precision/recall against labeled test data?
- Is the test corpus still representative? Are there other public prompt-injection benchmarks or URLs we should include?
- Are there obvious optimizations we skipped (e.g., faster file walking, smaller model, capping noisy rules)?
- Are report branches becoming too large? Should old report branches be archived?

Document the outcome of the review in CLAUDE.md or the next commit message.

### Report branches

- Create each report branch as a Git worktree from `master`:
  ```bash
  git worktree add -B report-<devicename>-<date> .worktrees/report-<devicename>-<date> master
  ```
- Generate a random Deadpool-style device name for each run (e.g., `boom-zany-sarcasm-cd4dee48`).
- Build the binary in the worktree, run `./scripts/redteam.sh`, run the corpus scan, and write a `COMPARISON.md` against the previous accepted version.
- Push only the report branch:
  ```bash
  git push -u origin report-<devicename>-<date>
  ```
- Keep tool improvements on `master`; never commit build artifacts or source-code experiments to report branches.

## Security hardening notes

- `safeanalyze clone` validates both the URL and the destination directory to prevent git option injection (`-u`, `--upload-pack`, shell metacharacters). It invokes `git clone` with a `--` separator so the URL and directory are always treated as positional arguments.
- External scanners run as separate processes with non-fatal error handling; a missing or crashing scanner must not abort the whole pipeline.
- Fast mode intentionally avoids ML inference, external scanners, and repository walking to keep latency deterministic.
- Potential secret values in finding `match` fields are redacted before reports are written, so report branches can be pushed publicly.

## Fast-mode latency budget

- Target: ≤ 100 ms per payload; ideally a few milliseconds.
- Fast mode runs only `yara` and `hiddenchars` stages on a single payload.
- No ML inference, no external scanners, no entropy analysis, no repository walk.

## Report contents

- Each report directory includes `safeanalyze.json`, `safeanalyze.md`, `safeanalyze.html`, `safeanalyze.sarif`, and `duration.txt`.
- Reports must include `safeanalyze_version` and `scan_mode` metadata.
- Durations are recorded in milliseconds (`duration_ms`), not wall-clock end times.

## Useful commands

```bash
# Build
go build -o safeanalyze .

# Test
go test ./...

# Fast inspect latency check
echo 'ignore all previous instructions' | ./safeanalyze inspect --verbose

# Thorough repo scan
./safeanalyze scan ./repo --mode thorough -o ./out

# Install dependencies
./safeanalyze install --all

# Autoresearch helpers
./scripts/redteam.sh
./scripts/premortem.sh
```

## Current autoresearch iteration

- **Version under test:** v0.3.9
- **Status:** Accepted (removes unsafe Semgrep skip).
- **Change:** Removed the 50-file minimum gate so Semgrep scans repositories of any size when enabled.
- **Last accepted corpus improvement:** v0.3.7 parallel YARA scanning (`report-boom-zany-sarcasm-cd4dee48-2026-07-15`).
- **Previous reverted iterations:**
  - v0.3.4 — encoded-prompt-injection fragment expansion added latency but no new detections.
  - v0.3.8 (parallel entropy/hiddenchars) — did not improve latency, reverted before release.
- **Stagnation check:** Two consecutive no-corpus-gain attempts have now occurred (parallel entropy/hiddenchars, small-model ML). Further progress likely requires either a better small prompt-injection model or a new detection-rule hypothesis backed by gap analysis.
- **Next candidates:**
  - Evaluate `Llama-Prompt-Guard-2-86M-onnx` and `22M-onnx` for memory/latency/precision; if suitable, enable a sub-2 GB model by default.
  - Add targeted YARA rules for chat-template boundary tokens or tool-output injection only if gap analysis shows missing true positives.
  - Decide whether `node_modules`/`vendor` should remain in `dependency_paths` for thorough mode.
