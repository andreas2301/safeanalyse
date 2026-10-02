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
- Corpus targets checked out at the pinned upstream SHAs/sha256s in `testdata/eval/SOURCES.md` ("Thorough-mode corpus manifest"; re-baselined 2026-10-02). Do not compare against runs on other revisions.
- Labeled eval sets fetched with `./scripts/fetch_eval.sh` (pinned, sha256-verified, seeded; not committed). Dev: `testdata/eval/{deepset,llmail,browsesafe}.jsonl`. Holdout: `testdata/eval/{deepset,llmail,browsesafe}-holdout.jsonl`.
- The corpus scan config sets `no_verification: true` on the `trufflehog` scanner so results do not depend on live credential verification.
- Enough disk space and memory for the test corpus. Note: the default DeBERTa ONNX model uses ~11 GB RAM at load/inference time, so the stochastic ML stage is currently disabled by default until a sub-2 GB model is validated.
- A stable, otherwise idle machine for duration comparisons (run-to-run variance should be <10 % for large targets).

### Iteration steps

1. **Measure baseline** — run `./scripts/fetch_eval.sh`, then `safeanalyze eval <set>.jsonl --json eval-<set>.json --fp-out fp-<set>.jsonl --fn-out fn-<set>.jsonl` on every dev **and** holdout set. Record precision, recall, F1, TP/FP/TN/FN and p50/p95 latency. Also run `safeanalyze scan --mode thorough` against the pinned test corpus (trufflehog `no_verification: true`) and record findings + durations from `safeanalyze.json` (`duration_ms`) and `duration.txt`.
2. **Premortem** — run `./scripts/premortem.sh` and ask: "If this improvement lands and the tool still fails in production, what most likely broke?" Document the biggest risks (false positives, latency blow-out, missing variants).
3. **Analyze gaps** — inspect the **dev** sets' `fn-*.jsonl` / `fp-*.jsonl` and the corpus findings against known prompt-injection patterns. Never inspect holdout misclassifications to design rules.
4. **Hypothesize** — pick one concrete improvement: a new YARA rule, a tuned entropy threshold, a smaller/faster model, a file-size limit, or parallelism.
5. **Implement** — make the smallest change that tests the hypothesis. Do not combine multiple unrelated changes in one iteration.
6. **Red-team** — run `./scripts/redteam.sh`. Try to evade the new check with rephrased, encoded, or multi-language injections. If it is trivially bypassed, revert or harden.
7. **Evaluate** — re-run `./scripts/fetch_eval.sh` (output sha256s must match `testdata/eval/SOURCES.md`) and `safeanalyze eval` on every dev and holdout set, then re-run the same pinned test corpus. Compare:
   - Precision, recall and F1 per labeled set (dev and holdout).
   - Total and per-target finding counts.
   - Wall-clock duration per target (`duration_ms`).
   - Fast-mode latency (`./scripts/redteam.sh`).
   - Any new errors or scanner skips.
8. **Keep or revert** — apply the decision rules below. If accepted, commit and bump the version. If not, revert and try another hypothesis.
9. **Repeat** until two consecutive iterations show no measurable improvement.

### Decision rules

- **Primary metrics** are F1 and precision on the labeled sets. Raw corpus finding counts are secondary: the corpus is unlabeled, so a higher count can mean more false positives.
- **Holdout must not regress.** An iteration that lowers F1 or precision on any holdout set is reverted, whatever the dev gains. `llmail-holdout` reuses the dev benign emails, so only its recall is an independent signal.
- **Accept** the iteration if holdout F1/precision do not regress **and**:
  - Dev F1 improves without a precision drop on any labeled set, **or**
  - Latency (eval p95, fast-mode, corpus `duration_ms`) decreases with unchanged labeled metrics and corpus findings, **or**
  - A security hardening fix removes a real foot-gun without regressing metrics.
- **Revert** the iteration if:
  - F1 or precision drops on any dev or holdout set, **or**
  - Corpus finding count drops without a labeled-set explanation (e.g. removed false positives), **or**
  - Latency increases without a labeled-metric gain, **or**
  - The change introduces non-deterministic output or new errors.
- **Stagnation** — stop the loop after two consecutive accepted/reverted iterations produce no improvement in labeled F1/precision or latency.

### Process review

After every five iterations (or immediately after two consecutive no-improvement iterations), review the process itself:

- Are we measuring the right metric? Are the labeled dev/holdout sets still representative, and is the holdout still untouched by rule design?
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

# Labeled eval (dev + holdout)
./scripts/fetch_eval.sh
for f in testdata/eval/*.jsonl; do ./safeanalyze eval "$f"; done

# Autoresearch helpers
./scripts/redteam.sh
./scripts/premortem.sh
```

## Current autoresearch iteration

- **Version under test:** v0.3.18 (last accepted).
- **Status:** Accepted (iteration 12). llmail dev F1 rises (0.798 → 0.824) and llmail-holdout F1 rises (0.778 → 0.790), with no FP change on any dev or holdout set. Every other labeled metric and every corpus count is identical to v0.3.17. Stagnation counter stays at 0.
- **Change:** `data_exfiltration_email` gains four line-based tool-call / structured-send patterns: (1) a send-mail tool name (`send_email`, `sendemail`, `send_mail`, `send_message`, `mail_send`, `api_call`) with an address and a payload key (`body`/`content`/`message` followed by `:` or `=`, or "with (the) body/content/message") within 80 characters either way; (2) a recipient key (`recipient(s)`, `receiver`, `email_to`) with `:`/`=` and an address, plus a payload key within 80 characters either way; (3) a `to` key with an address directly followed by a `body`/`content`/`message` key; (4) a piped `| mail -s ... <address>`. Hypothesis from BASELINE.md cross-set theme #5 and the "JSON tool-call injection in llmail FNs" candidate, built on dev sets only: 16 of the 101 llmail dev FNs send mail through tool/function syntax or key/value blocks. No holdout file was opened.
- **Metrics (fast-mode eval, v0.3.17 → v0.3.18):**
  - Dev: llmail TP 199 → 210, R 0.663 → 0.700, F1 0.798 → 0.824 (FP 0); deepset F1 0.571 (FP 0) and browsesafe F1 0.681 (FP 76) unchanged. `data_exfiltration_email` hits: llmail injection 142 → 157, browsesafe 132/24 (injection/benign) unchanged.
  - Holdout: llmail-holdout TP 191 → 196, R 0.637 → 0.653, F1 0.778 → 0.790 (FP 0); deepset-holdout F1 0.417 (FP 2) and browsesafe-holdout F1 0.633 (FP 82) unchanged.
  - Latency p95: dev deepset 3.44 → 3.75 ms (run 2: 4.34), llmail 4.02 → 5.09 ms (run 2: 4.77), browsesafe 25.4 → 27.4 ms (run 2: 28.8); holdout 3.65 → 4.08, 4.22 → 4.80, 26.4 → 29.3 ms. All within the 100 ms budget; the rise is accepted because of the labeled gain.
  - Thorough corpus: findings 11166 → 11166, identical per target, per source and per rule. Total `duration_ms` 35582 → 35489, 0 errors.
  - Red-team: 12/12 payloads flagged at 15–16 ms.
  - Determinism: two dev eval runs match once latency fields are removed; FP/FN JSONL files are byte-identical.
- **Known gaps:** a tool name or recipient key with an address but no payload key (`send_mail(..., ['to@example.com'])`, `Recipient: x@y` alone) is not matched by design; unpiped `mail -s` is not matched; obfuscated addresses (`(at)`, `＠`, `&#64;`, `%40`) and payload keys more than 80 characters from the address are missed. Known benign trigger from red-team: Ruby `mailer.send_email(to: "dev@example.com", body: render(:welcome))` (1 of 14 benign code/form snippets).
- **Rejected options this iteration (not released):**
  - Tool name or recipient key plus address only, and unpiped `mail -s` (11 of 14 benign red-team snippets flagged). Requiring a payload key and a pipe cost 5 of the draft's 16 llmail dev TPs.
- **Last accepted corpus improvement:** v0.3.17 (corpus findings +141 from `data_exfiltration_email`, all InjecAgent attacker strings in a spot check).
- **Previous reverted iterations:**
  - v0.3.4 — encoded-prompt-injection fragment expansion added latency but no new detections.
  - v0.3.8 (parallel entropy/hiddenchars) — did not improve latency, reverted before release.
  - Iteration 6 (v0.3.16 candidate) — zero-width chars reported only inside Latin text or runs of 3+; no gain.
  - Iteration 7 (v0.3.16 candidate) — `prompt_leak_request` rule (en/de); no gain.
  - Iteration 9 (v0.3.17 candidate) — labeled metrics identical to v0.3.16; not released.
  - Iteration 11 (v0.3.18 candidate) — not released. Measured browsesafe dev TP 194 → 188 / FP 76 → 62 (`account_access_request` and `data_exfiltration_email` benign hits down); llmail and deepset unchanged.
- **Stagnation check:** 0 consecutive no-gain iterations (iteration 12 accepted for its llmail F1 gain).
- **Next candidates (from the dev-set FN analysis):**
  - Remaining llmail FNs with a plain email address: "notify ... mail to me", address far from "to", and other verbs.
  - browsesafe precision (76 dev FPs; `account_access_request`, hidden-char and `data_exfiltration_email` benign hits).
  - Spot-check the +1064 (v0.3.11) and +150 (v0.3.13) InjecAgent YARA findings and the +24 BIPIA hits for FP inflation.
  - Evaluate `Llama-Prompt-Guard-2-86M-onnx` / `22M-onnx` for memory/latency/precision on the labeled sets.
  - Decide whether `node_modules`/`vendor` should remain in `dependency_paths` for thorough mode.

## Process review 2026-10-02

Done after 7 labeled-eval iterations (5 accepted, then 2 with no gain).

- **Metric:** Per-set precision/recall/F1 on labeled dev and holdout sets (deepset, llmail, browsesafe) is the right main metric and should stay. Raw corpus finding count is not a quality signal: InjecAgent makes up 8138 of 11025 findings (74 %), and rules can inflate it without anyone checking. From now on, treat corpus counts only as a regression and error check.
- **Corpus:** Not fully representative. Recall is stuck around 0.27–0.65 (deepset-holdout 0.27). Browsesafe is the only set with false positives (76 dev / 82 holdout), so precision is tested mainly on HTML. The sets have no JSON tool-call or agent-trace benign/injection samples, and there is no labeled multilingual set.
- **Skipped optimizations:** Browsesafe fast-mode p95 is 345 ms dev / 337 ms holdout, over 3x the 100 ms budget, and it grew with each regex added (240 ms baseline → 345 ms). An input-size cap or HTML pre-filter was never tried. That is the next iteration and it is latency-only.
- **Report branches:** 19 local `report-*` branches (11 on origin) and 18 worktrees using about 775 MB under `.worktrees/`. That is too many. Archive the July 2026 ones (keep the last accepted ones), remove stale worktrees with `git worktree remove`, and stop creating one report branch per rejected iteration.

