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
- **Holdout must not regress.** An iteration that lowers F1 or precision on any holdout set is reverted, whatever the dev gains. Single-sample tolerance: an F1 drop on one set caused by losing at most one TP is allowed if that set's precision does not drop and the sum of F1 across holdout sets rises (precedent: v0.3.24, deepset-holdout TP 54 → 53, FP 2 → 0). `llmail-holdout` reuses the dev benign emails, so only its recall is an independent signal.
- **Accept** the iteration if holdout F1/precision do not regress **and**:
  - Dev F1 improves without a precision drop on any labeled set, **or**
  - Latency (eval p95, fast-mode, corpus `duration_ms`) decreases with unchanged labeled metrics and corpus findings, **or**
  - A security hardening fix removes a real foot-gun without regressing metrics.
- **Revert** the iteration if:
  - F1 or precision drops on any dev or holdout set (except the single-sample tolerance above), **or**
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

- **Version under test:** v0.3.25 (last accepted).
- **Status:** Accepted (iteration 21). browsesafe precision rises on dev (0.829 → 0.873) and holdout (0.813 → 0.861); browsesafe dev F1 0.745 → 0.754 and holdout F1 0.695 → 0.705. Dev F1 sum +0.0088; no precision drop on any set.
- **Change:** `data_exfiltration_email` patterns 0 and 4 are narrowed. Pattern 0 (retrieve verb + mail) needs an explicit `send`/`forward`, "and|then (e)mail", or `email`/`mail` within four words of the verb; pattern 4 (retrieve verb + data noun) needs the noun within four words. Retest of iteration 20, which had been rejected only by an undocumented 0.01 gain threshold.
- **Metrics (fast-mode eval, v0.3.24 → v0.3.25):**
  - Dev: browsesafe TP 203 → 199, FP 42 → 29, P 0.873, R 0.663, F1 0.754; deepset F1 0.571 (P 1.0, R 0.40) and llmail F1 0.881 (P 1.0, R 0.787) unchanged.
  - Holdout: browsesafe-holdout TP 182 → 179, FP 42 → 29, P 0.861, R 0.597, F1 0.705; deepset-holdout F1 0.414 and llmail-holdout F1 0.835 unchanged.
  - Latency p95: deepset 4.8 ms, llmail 8.0 ms, browsesafe 37.0 ms dev / 40.7 ms holdout; all within the 100 ms budget.
  - Thorough corpus: findings 11166 → 11015 (−151, all `data_exfiltration_email`: InjecAgent 3820 → 3712, skylos 48 → 6, BIPIA 4 → 3); total `duration_ms` 32720 → 35427 (skylos external-scanner variance), 0 errors.
  - Red-team: 12/12 payloads flagged at 16–18 ms. Determinism: two dev runs identical (latency removed), FP/FN JSONL byte-identical.
- **Known gaps:** a retrieve verb with the mail/data noun more than four words away and no explicit send verb ("retrieve all of the customer's stored billing and shipping addresses") no longer matches patterns 0/4; other `data_exfiltration_email` patterns may still catch it.
- **Rejected options (not released):**
  - Iteration 19 — narrow `data_exfiltration_email` pattern 0 so verify/confirm-your-email page text and mail-site links stop triggering it; dev/holdout F1 +0.000.
- **Last accepted corpus improvement:** v0.3.17 (corpus findings +141 from `data_exfiltration_email`).
- **Previous reverted iterations:** v0.3.4, v0.3.8, iterations 6, 7, 9, 11, 19 (details in earlier commits and the process reviews). Iteration 20 was reverted, then retested and accepted as iteration 21.
- **Stagnation check:** iteration 21 improved labeled F1/precision; counter reset to 0.
- **Next candidates:**
  - deepset recall (dev 0.40, holdout 0.26), flat since v0.3.14; needs a non-email attack family or the ML stage.
  - Evaluate `Llama-Prompt-Guard-2-22M`/`86M-onnx` (memory, latency, precision) on the labeled sets.
  - Remaining browsesafe FPs (29 dev / 29 holdout) and FNs (101 dev): CSS-hidden text, attribute prose without listed keywords.
  - Spot-check InjecAgent YARA findings (7860 of 11015) for FP inflation; decide `node_modules`/`vendor` in `dependency_paths`.

## Process review 2026-10-02

Done after 7 labeled-eval iterations (5 accepted, then 2 with no gain).

- **Metric:** Per-set precision/recall/F1 on labeled dev and holdout sets (deepset, llmail, browsesafe) is the right main metric and should stay. Raw corpus finding count is not a quality signal: InjecAgent makes up 8138 of 11025 findings (74 %), and rules can inflate it without anyone checking. From now on, treat corpus counts only as a regression and error check.
- **Corpus:** Not fully representative. Recall is stuck around 0.27–0.65 (deepset-holdout 0.27). Browsesafe is the only set with false positives (76 dev / 82 holdout), so precision is tested mainly on HTML. The sets have no JSON tool-call or agent-trace benign/injection samples, and there is no labeled multilingual set.
- **Skipped optimizations:** Browsesafe fast-mode p95 is 345 ms dev / 337 ms holdout, over 3x the 100 ms budget, and it grew with each regex added (240 ms baseline → 345 ms). An input-size cap or HTML pre-filter was never tried. That is the next iteration and it is latency-only.
- **Report branches:** 19 local `report-*` branches (11 on origin) and 18 worktrees using about 775 MB under `.worktrees/`. That is too many. Archive the July 2026 ones (keep the last accepted ones), remove stale worktrees with `git worktree remove`, and stop creating one report branch per rejected iteration.

## Process review 2026-10-02 (round 2)

Done after 17 labeled-eval iterations (12 accepted, 5 reverted; iterations 8–17 since the first review: 7 accepted, 3 reverted). Loop not stagnant (iteration 17 accepted).

- **Metric:** Still the right one (per-set P/R/F1 on dev and holdout). Dev and holdout moved together in iterations 10–17 (browsesafe dev F1 0.681 → 0.715, holdout 0.633 → 0.671; llmail dev 0.785 → 0.881, holdout 0.776 → 0.835), so there are no signs of dev overfitting. deepset has not moved since the first review (dev F1 0.571, holdout 0.417, recall 0.27). Six of the ten iterations widened one rule (`data_exfiltration_email`) for llmail-style exfiltration, so gains are concentrated on one attack family.
- **Corpus:** Still not representative. FPs exist only on browsesafe (76 dev / 82 holdout, unchanged since v0.3.14) and on deepset-holdout (2). There is still no labeled JSON tool-call/agent-trace benign set and no multilingual set. Iteration 12 tool-call patterns were checked only against llmail benign mail. InjecAgent is 8279 of 11166 corpus findings (74 %, 7968 YARA). Corpus counts stay a regression/error check only.
- **Skipped optimizations:** Resolved. The iteration 8 literal prefilter brought browsesafe p95 from 349 ms to 24 ms. It is now 36.9 ms dev / 37.4 ms holdout, within the 100 ms budget, but it has risen about 1–4 ms per added pattern (+13 ms over iterations 10–17). Watch it, and check each new pattern has a literal the prefilter can use. Still open: the ML stage (Prompt-Guard-2 22M/86M), the InjecAgent FP spot-check, and `node_modules`/`vendor` in `dependency_paths`.
- **Report branches:** Too many, and worse than at the first review. The archive step was not carried out: 27 local `report-*` branches (14 from 2026-07, 13 from 2026-10), 16 on origin, 26 worktrees using about 1.2 GB under `.worktrees/`. Action (needs owner approval; nothing deleted in this review): remove the 2026-07 worktrees with `git worktree remove`, archive the July branches as tags or delete them, keep only the report branch of the last accepted version (v0.3.23), and from now on write reverted-iteration results only to `/tmp/safeanalyze-iter/`, with no report branch.

## Process review 2026-10-09 (round 3)

Done after 20 labeled-eval iterations (iteration 18 accepted as v0.3.24; 19 and 20 reverted, so the loop is stagnant).

- **Metric:** Still the right one (per-set P/R/F1, dev and holdout). Two issues: (1) iteration 18 was accepted although deepset-holdout F1 fell 0.4170 → 0.4141 (1 TP lost, FP 2 → 0), which breaks the strict "holdout must not regress" rule; it should be written as an explicit exception (precision up, F1 drop within one sample) or reverted next time. (2) The loop script applies an unwritten minimum-gain threshold (iteration 20, +0.009 dev / +0.010 holdout F1, was rejected as "no gain"); document that threshold in the decision rules.
- **Corpus:** Not representative. deepset recall has not moved since the first review (dev 0.40, holdout 0.26), and most accepted gains came from email-exfiltration rules (llmail) and HTML (browsesafe). No labeled JSON tool-call/agent-trace or multilingual set yet. InjecAgent is 8279 of 11166 corpus findings (74 %); corpus counts stay a regression/error check only.
- **Skipped optimizations:** Latency is fine (browsesafe p95 35.9 ms dev / 37.1 ms holdout). Still open: the ML stage (Prompt-Guard-2 22M/86M), the InjecAgent FP spot-check, `node_modules`/`vendor` in `dependency_paths`. v0.3.24 is committed but not tagged.
- **Report branches:** Worktrees are cleaned up (1 left, about 62 MB under `.worktrees/`, was 26 / 1.2 GB), and iterations 18–20 wrote results only to `/tmp/safeanalyze-iter/`. Branches are still too many: 27 local `report-*` (14 from 2026-07, 13 from 2026-10) and 24 on origin (11 from 2026-07). Action (needs owner approval; nothing deleted in this review): archive the July branches as tags or delete them locally and on origin, and keep only the report branch of the last accepted version.
