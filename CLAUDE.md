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
- Labeled eval sets fetched with `./scripts/fetch_eval.sh` (pinned, sha256-verified, seeded; not committed). Dev: `testdata/eval/{deepset,llmail,browsesafe,toolresp,nemotron,multilingual,spanish}.jsonl`. Holdout: the same names with `-holdout`. Report-only stress set: `testdata/eval/promptprotection-stress.jsonl`. See "Set taxonomy" below for how each set is used.
- The corpus scan config sets `no_verification: true` on the `trufflehog` scanner so results do not depend on live credential verification.
- Enough disk space and memory for the test corpus. Note: the default DeBERTa ONNX model uses ~11 GB RAM at load/inference time, so the stochastic ML stage is currently disabled by default until a sub-2 GB model is validated.
- A stable, otherwise idle machine for duration comparisons (run-to-run variance should be <10 % for large targets).

### Iteration steps

1. **Measure baseline** — run `./scripts/fetch_eval.sh`, then `safeanalyze eval <set>.jsonl --json eval-<set>.json --fp-out fp-<set>.jsonl --fn-out fn-<set>.jsonl` on every dev **and** holdout set and on the report-only stress set. Record precision, recall, F1, TP/FP/TN/FN and p50/p95 latency. Also run `safeanalyze scan --mode thorough` against the pinned test corpus (trufflehog `no_verification: true`) and record findings + durations from `safeanalyze.json` (`duration_ms`) and `duration.txt`.
2. **Premortem** — run `./scripts/premortem.sh` and ask: "If this improvement lands and the tool still fails in production, what most likely broke?" Document the biggest risks (false positives, latency blow-out, missing variants).
3. **Analyze gaps** — inspect the `fn-*.jsonl` / `fp-*.jsonl` of the **PRIMARY dev** sets (llmail, browsesafe, toolresp) and the multilingual dev set, plus the corpus findings, against known prompt-injection patterns. Never inspect holdout misclassifications, GUARD-ONLY (deepset, nemotron, spanish) misclassifications or promptprotection-stress (REPORT-ONLY) misclassifications to design rules.
4. **Hypothesize** — pick one concrete improvement: a new YARA rule, a tuned entropy threshold, a smaller/faster model, a file-size limit, or parallelism.
5. **Implement** — make the smallest change that tests the hypothesis. Do not combine multiple unrelated changes in one iteration.
6. **Red-team** — run `./scripts/redteam.sh`. Try to evade the new check with rephrased, encoded, or multi-language injections. If it is trivially bypassed, revert or harden.
7. **Evaluate** — re-run `./scripts/fetch_eval.sh` (output sha256s must match `testdata/eval/SOURCES.md`) and `safeanalyze eval` on every dev and holdout set and on the report-only stress set (`promptprotection-stress.jsonl`; recorded, never gating), then re-run the same pinned test corpus. Compare:
   - Precision, recall and F1 per labeled set (dev and holdout), grouped by the set taxonomy below; stress-set metrics are recorded but do not gate.
   - Total and per-target finding counts.
   - Wall-clock duration per target (`duration_ms`).
   - Fast-mode latency (`./scripts/redteam.sh`).
   - Any new errors or scanner skips.
8. **Keep or revert** — apply the decision rules below. If accepted, commit and bump the version. If not, revert and try another hypothesis.
9. **Repeat** until two consecutive iterations show no measurable improvement.

### Set taxonomy

Every labeled set has one role. Each role applies to the dev set and its `-holdout` twin unless stated otherwise. Sources, licenses and sampling rules are in `testdata/eval/SOURCES.md`.

| Role | Sets | Gate | Hypothesis source? | Counts toward acceptance? |
|---|---|---|---|---|
| **PRIMARY** | llmail, browsesafe, toolresp | F1 and precision must not drop on dev or holdout | Yes, dev only | Yes (dev F1 gains) |
| **RECALL-ONLY** | multilingual | Recall must not drop on dev or holdout (positives only; precision/F1 are meaningless) | Yes, dev only | No |
| **GUARD-ONLY** | deepset, nemotron, spanish | F1 and precision must not drop on dev or holdout | **Never** | **No** — gains are ignored |
| **REPORT-ONLY** | promptprotection-stress | None; record P/R/F1/FP in the iteration notes | **Never** | No |

- deepset was moved to GUARD-ONLY by user decision on 2026-10-14 because it was overfitted: iterations 26, 28, 29 and 30 moved deepset dev only, and dev recall reached 0.77 against 0.31 on the holdout. Its FP/FN files must not be read to design rules, and a deepset gain alone never justifies accepting an iteration.
- nemotron was made GUARD-ONLY by user decision on 2026-10-14: 98 % of its positives use one of two injection wrappers (`[PRIORITY MESSAGE]` and `⚠️ URGENT: Before completing any other tasks`), so a single rule for a wrapper would move recall without generalizing. The same restrictions as deepset apply.
- spanish was made GUARD-ONLY by user decision on 2026-10-16: with 52/53 positives (mostly machine-translated) it overfitted in round 4 (dev F1 0.400 → 0.831 vs holdout 0.464 → 0.605). The same restrictions as deepset apply.
- `llmail-holdout` reuses the dev benign emails, so only its recall is an independent signal. multilingual, nemotron, toolresp and most of spanish are synthetic or machine-translated; read their caveats in SOURCES.md before drawing conclusions.

### Decision rules

- **Primary metrics** are F1 and precision on the PRIMARY sets (see "Set taxonomy"). Raw corpus finding counts are secondary: the corpus is unlabeled, so a higher count can mean more false positives.
- **Holdout must not regress.** An iteration that lowers F1 or precision on any PRIMARY or GUARD-ONLY holdout set, or recall on the RECALL-ONLY holdout set, is reverted, whatever the dev gains. Single-sample tolerance: an F1 drop on one PRIMARY or GUARD-ONLY set caused by losing at most one TP is allowed if that set's precision does not drop and the sum of F1 across the **PRIMARY holdout sets** rises (precedent: v0.3.24, deepset-holdout TP 54 → 53, FP 2 → 0). Apart from that tolerance, GUARD-ONLY sets allow no F1 or precision drop at all. The tolerance does not apply to the RECALL-ONLY sets: their recall must not drop by even one TP. `llmail-holdout` reuses the dev benign emails, so only its recall is an independent signal.
- **Accept** the iteration if holdout F1/precision do not regress **and**:
  - Dev F1 improves on at least one PRIMARY set without a precision drop on any PRIMARY or GUARD-ONLY set and without a recall drop on the RECALL-ONLY sets, **or**
  - Latency (eval p95, fast-mode, corpus `duration_ms`) decreases with unchanged labeled metrics and corpus findings, **or**
  - A security hardening fix removes a real foot-gun without regressing metrics.
- **Revert** the iteration if:
  - F1 or precision drops on any PRIMARY or GUARD-ONLY dev or holdout set (except the single-sample tolerance above), **or**
  - Recall drops on the RECALL-ONLY dev or holdout set, **or**
  - Corpus finding count drops without a labeled-set explanation (e.g. removed false positives), **or**
  - Latency increases without a labeled-metric gain, **or**
  - The change introduces non-deterministic output or new errors.
- **Stagnation** — stop the loop after two consecutive accepted/reverted iterations produce no improvement in PRIMARY F1/precision or latency.

### Process review

After every five iterations (or immediately after two consecutive no-improvement iterations), review the process itself:

- Are we measuring the right metric? Are the labeled dev/holdout sets still representative, and is the holdout still untouched by rule design?
- Is the test corpus still representative? Are there other public prompt-injection benchmarks or URLs we should include?
- Are there obvious optimizations we skipped (e.g., faster file walking, smaller model, capping noisy rules)?
- Are report branches becoming too large? Should old report branches be archived?

Document the outcome of the review in CLAUDE.md or the next commit message.

### Report branches

- One report branch per loop **round** (user decision 2026-10-16), created after the round's last accepted iteration and comparing the round's final version against the version the round started from. Per-iteration results stay in `/tmp/safeanalyze-iter/<iter>/` only. Existing per-iteration branches are kept.
- Create the report branch as a Git worktree from `master`:
  ```bash
  git worktree add -B report-<devicename>-<date> .worktrees/report-<devicename>-<date> master
  ```
- Generate a random Deadpool-style device name for each run (e.g., `boom-zany-sarcasm-cd4dee48`).
- Build the binary in the worktree, run `./scripts/redteam.sh`, run the corpus scan, and write a `COMPARISON.md` against the version the round started from.
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

# Labeled eval (dev + holdout + report-only stress set)
./scripts/fetch_eval.sh
O=/tmp/safeanalyze-iter/<iter>; mkdir -p "$O"
for f in testdata/eval/*.jsonl; do n=$(basename "$f" .jsonl)
  ./safeanalyze eval "$f" --json "$O/eval-$n.json" --fp-out "$O/fp-$n.jsonl" --fn-out "$O/fn-$n.jsonl"
done

# Autoresearch helpers
./scripts/redteam.sh
./scripts/premortem.sh
```

## Current autoresearch iteration

- **Loop paused (user decision 2026-10-16)** after round 5 iteration 42, pending a strategy decision: the honest assessment (`docs/ASSESSMENT.md`) shows regex-only detection has a hard ceiling (stress recall 0.17, multilingual 0.02) and falling returns. Do not start new rule-adding rounds until a new strategy is confirmed. Iteration 42 (narrow `obfuscated_javascript` for JSON `\u` escapes) was reverted: toolresp FP 31 → 2 but TP 189 → 160 (dev), 157 → 132 (holdout), so F1 dropped — the precision/recall trade-off is worth revisiting.
- **v0.3.46** (outside the loop): `scan` blocking gates scoped to their own stage (external bug report). `hiddenchars` deliberately does not check `Mn` combining marks such as emoji variation selector U+FE0F.
- **Known gap:** the documented Squid `external_acl_type` helper only receives `%SRC %DST %METHOD %URI`, never response bodies; an ICAP RESPMOD mode is planned.

- **Version under test:** v0.3.45 (last accepted; iteration 41, first of round 5, latency-only). Previous: v0.3.44 (iteration 40); round 5 baseline measurements are in `/tmp/safeanalyze-iter/v0.3.44-r5-baseline/`, iteration 41 measurements in `/tmp/safeanalyze-iter/iter41/` (rerun in `iter41/run2/`).
- **Status:** Accepted (iteration 41, latency-only). Labeled metrics unchanged on every set (FP/FN JSONL byte-identical to the baseline), corpus findings identical, browsesafe p95 70.72 → 44.09 ms dev and 77.01 → 48.20 ms holdout.
- **Change:** `yara.Engine.ScanFile` replaces the per-pattern, per-line `strings.Contains` prefilter (`prefilterPass`) with a single-pass literal-to-line index (`lineIndex`): one walk over the folded text, keyed on each literal's first three bytes, builds a per-literal line bitset; each pattern's AND-of-ORs prefilter is evaluated on the bitsets and the regex runs only on candidate lines. Literals with `\n` never mark a line; a literal under three bytes marks every line. Hypothesis from a CPU profile of inspect over browsesafe dev (24.0 s CPU): prefilter 7.7 s (32 %, largest item), regex execution 8.4 s (3.15 s in leetspeak `encoded_prompt_injection[4]`, 0.89 s in `credential_hardcode[0]`, both without a prefilter), per-payload regex compile in `NewEngine` 3.3 s (about 5 ms per payload, the floor for llmail/toolresp), hiddenchars `IsSuspicious` 2.1 s (parallel, off the critical path). Adds `TestLineIndexMatchesPerLineContains`.
- **Metrics (fast-mode eval, v0.3.44 → v0.3.45):**
  - PRIMARY dev: llmail F1 0.889 (FP 0), browsesafe F1 0.758 (FP 26), toolresp F1 0.727 (FP 31), spanish F1 0.831 (FP 0), all unchanged.
  - PRIMARY holdout: llmail-holdout 0.837, browsesafe-holdout 0.708, toolresp-holdout 0.638, spanish-holdout 0.605, all unchanged.
  - RECALL-ONLY: multilingual recall 0.025 / holdout 0.024, unchanged.
  - GUARD-ONLY: deepset 0.868 / holdout 0.468, nemotron 0.204 / holdout 0.209, all unchanged.
  - REPORT-ONLY stress: promptprotection-stress P 0.500, R 0.167, F1 0.250, TP 30, FP 30, unchanged.
  - Latency p95: browsesafe dev 70.72 → 44.09 ms (rerun 44.47 ms, p50 28.59 → 18.33 ms), browsesafe-holdout 77.01 → 48.20 ms (p50 28.27 → 17.44 ms); llmail dev 12.73 → 12.06 ms, llmail-holdout 12.51 → 12.49 ms; toolresp dev 11.68 → 11.20 ms, toolresp-holdout 10.66 → 10.13 ms; all other sets < 11 ms.
  - Thorough corpus: findings 11015 → 11015 (identical per target), 0 errors; total `duration_ms` 36373 → 32789 (skylos 17913 → 14179, mostly external-scanner variance).
  - Determinism: iter41 and iter41/run2 FP/FN JSONL byte-identical to each other and to the baseline.
- **Known gaps:** the hand-off frame needs "follow" + "all"/"every" + a colon right after "the following"; other verbs (obey, execute, carry out, do), "below"/"after this" instead of "the following", missing colon, multi-space letter splits, and non-English forms are not matched. A role boundary forged on the next line (`</tool>`, `< user>` with a space) is still missed by `chat_template_boundary` because it scans line by line. Earlier gaps: single-word quoted replies ("reply with 'STOP'") are excluded on purpose from `output_constraint`, as are unquoted replies and non-English conditionals. toolresp still misses 111 of 300 dev positives (143 of 300 holdout); llmail misses 60 of 300 dev positives. 29 of 31 toolresp dev FPs still come from `obfuscated_javascript` on JSON `\u` escapes. spanish still misses 15 of 52 dev positives. browsesafe p95 is about 44 ms dev / 48 ms holdout after v0.3.45; the remaining hotspots are the unprefiltered leetspeak `encoded_prompt_injection[4]` and `credential_hardcode[0]` patterns and per-payload regex compile in `NewEngine` (about 5 ms).
- **Rejected options (not released):**
  - Iteration 36 alternatives — unescaping JSON `\n`/`\t` before matching (+2 toolresp TP only); urgency or verification lures for browsesafe (best variant 95 positive vs 88 benign hits, confirming the earlier rejection); a broader forced-output frame anchored on punctuation, comma or "and" (+11 to +12 toolresp TP, but false-triggered in red-team).
  - Iteration 32 alternative — match the HackAPrompt canary "I have been PWNED" (largest single toolresp FN theme); not pursued as dataset memorization.
  - Iteration 31 alternative — drop the `obfuscated_javascript` `\u[0-9a-fA-F]{4}` pattern to remove 29 of 31 toolresp dev FPs; it would cut toolresp dev TP 143 → 46 because 105 injection samples also contain JSON `\u` escapes (dataset artifact), so not pursued.
  - Iteration 27 — gate the bare `get` verb in `account_access_request` pattern 0 to instruction starts; browsesafe-holdout TP 179 → 177, reverted for a holdout regression.
  - Iteration 19 — narrow `data_exfiltration_email` pattern 0; dev/holdout F1 +0.000.
- **Last accepted corpus improvement:** v0.3.17 (corpus findings +141 from `data_exfiltration_email`).
- **Previous reverted iterations:** v0.3.4, v0.3.8, iterations 6, 7, 9, 11, 19, 27 (details in earlier commits and the process reviews). Iteration 20 was reverted, then retested and accepted as iteration 21.
- **Stagnation check:** iteration 41 reduced latency with unchanged labeled metrics; counter stays at 0. Round 5 so far (iteration 41): 1 accepted, 0 reverted. Round 4 (iterations 31–40): 10 accepted, 0 reverted.
- **Next candidates:**
  - deepset and nemotron are GUARD-ONLY (see "Set taxonomy"): do not target them; they only must not regress.
  - Round 4 review: pause spanish-dev-driven rules (dev/holdout recall gap 0.28, see review); latency-only done in iteration 41 (browsesafe p95 74 → 44 ms); further latency candidates: prefilter literals for leetspeak `encoded_prompt_injection[4]` and `credential_hardcode[0]`, compile rules once instead of per payload; otherwise toolresp precision (`obfuscated_javascript` on JSON `\u` escapes).
  - toolresp precision and recall: 29 of 31 dev FPs (34 of 35 holdout) come from `obfuscated_javascript` firing on JSON tool responses; recall 0.630 dev / 0.523 holdout after v0.3.44.
  - llmail recall (0.800 dev / 0.720 holdout after v0.3.44): forged role boundaries split across lines or written with spaces (`< user>`).
  - spanish recall (0.712 dev / 0.434 holdout after v0.3.41; holdout did not move in v0.3.39 or v0.3.41, so check whether the remaining dev themes generalize).
  - multilingual recall (0.025 dev / 0.024 holdout; RECALL-ONLY, positives only).
  - Evaluate `Llama-Prompt-Guard-2-22M`/`86M-onnx` (memory, latency, precision) on the labeled sets.
  - Remaining browsesafe FPs (26 dev / 27 holdout) and FNs (101 dev): CSS-hidden text, attribute prose without listed keywords.
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

Redo. The first round-3 run had loop-script bugs: it checked the holdout on a summed F1 instead of per set, it rejected iterations with an undocumented 0.01 minimum-gain threshold (iteration 20, later retested and accepted as iteration 21), and iteration 19 was never applied to the tree. This review covers the redo, iterations 21–30: 9 accepted (v0.3.25–v0.3.33), 1 reverted (27, browsesafe-holdout F1 0.708 → 0.702). The loop stopped at the MAXITER=30 cap, not at stagnation.

- **Metric:** Still the right one, checked per set (P/R/F1 on each dev and holdout set; no summed holdout check, no unwritten threshold). Every accepted iteration in the redo kept every holdout set's F1/precision unchanged or better. Dev F1: deepset 0.571 → 0.868, llmail 0.881, browsesafe 0.715 → 0.758 (FP 76 → 26). Holdout F1: deepset 0.468 (R 0.31), llmail 0.833, browsesafe 0.708. Risk: dev overfitting on deepset. Iterations 26, 28, 29 and 30 moved deepset dev only; dev recall 0.77 against holdout 0.31, a gap that keeps growing. Prefer attack families that generalize, and weigh a dev-only gain from one or two samples carefully.
- **Corpus:** Not representative. There is still no labeled JSON tool-call/agent-trace or multilingual set, and deepset-holdout is the weakest set. InjecAgent is 8171 of 11015 corpus findings (74 %). Corpus counts are only a regression and error check: 11166 → 11015 after iteration 21, then unchanged through iteration 30, with 0 errors.
- **Skipped optimizations:** Browsesafe p95 rose from 36.9 ms to 49.8 ms dev (48.4 ms holdout) over the redo, about +1.3 ms per accepted pattern. It is still within the 100 ms budget, but each new pattern needs a prefilter literal. Corpus scans do not write `duration.txt`, which breaks the report-contents contract; fix that next. Still open: the ML stage (Prompt-Guard-2 22M/86M), the InjecAgent FP spot-check, and `node_modules`/`vendor` in `dependency_paths`. v0.3.25–v0.3.33 are tagged.
- **Report branches:** Worse. There are 37 local `report-*` branches (14 from 2026-07, 23 from 2026-10) and 34 on origin (11 from 2026-07, 23 from 2026-10). There is 1 worktree, about 62 MB under `.worktrees/`. Action (needs owner approval; nothing was deleted in this review): archive or delete the July branches locally and on origin, keep only the report branch of the last accepted version, and write per-iteration results only to `/tmp/safeanalyze-iter/`.

## Process review 2026-10-14 (round 4)

Done after iterations 31–40 (v0.3.35–v0.3.44): 10 accepted, 0 reverted. The loop stopped at the MAXITER=40 cap, not at stagnation. Baseline is `/tmp/safeanalyze-iter/v0.3.34-r4-baseline/`; final is `iter40/`.

- **Metric:** Still the right one: per-set P/R/F1 with the role taxonomy (PRIMARY / RECALL-ONLY / GUARD-ONLY / REPORT-ONLY). Since deepset and nemotron became GUARD-ONLY, every gain landed on a PRIMARY set. No PRIMARY or GUARD-ONLY set lost precision. FP counts did not change on any set (browsesafe 26/27, toolresp 31/35, all others 0, stress 30). So round 4 was recall-only work, and precision went untested by any new FP source.
- **Dev vs holdout (F1, round start → end):**
  - llmail: dev 0.881 → 0.889 (+4 TP), holdout 0.833 → 0.837 (+2 TP). Moves together.
  - browsesafe: dev 0.758, holdout 0.708. Neither moved; no round-4 iteration targeted it.
  - toolresp: dev 0.603 → 0.727 (+0.124, TP 143 → 189), holdout 0.577 → 0.638 (+0.061, TP 136 → 157). The holdout gained about half of the dev gain. The gains generalize partly, and the dev/holdout recall gap widened from 0.02 to 0.11.
  - spanish: dev 0.400 → 0.831 (+0.431, TP 13 → 37 of 52), holdout 0.464 → 0.605 (+0.141, TP 16 → 23 of 53). The recall gap went from −0.05 to +0.28, and the holdout did not move in v0.3.39 or v0.3.41.
  - multilingual (RECALL-ONLY): recall 0.018 → 0.025 dev, 0.018 → 0.024 holdout. Moves together but stays negligible.
  - GUARD-ONLY and stress: unchanged.
- **Overfitting:** spanish shows **clear signs of dev overfitting**. It is small (52/53 positives, mostly machine-translated), and three iterations (31, 35, 37) pushed dev recall from 0.25 to 0.71 while the holdout reached only 0.43. toolresp shows moderate signs. Action: pause spanish-dev-driven rules until the set is larger or until a candidate also moves some independent Spanish signal. Before accepting a dev gain of 10 or more TP on a set with fewer than 100 positives, require that the target theme is general (no dataset strings) and record the dev:holdout TP-gain ratio in the iteration notes.
- **Sets representative?** Partly.
  - toolresp is the only JSON tool-response set. It is synthetic, and its 31 dev FPs are a dataset artifact (`obfuscated_javascript` on `\u` escapes, which also occur in 105 positives).
  - spanish is too small to gate on.
  - multilingual is positives-only machine translation, so it cannot test precision.
  - Precision is now tested only by browsesafe and toolresp benign rows. There is still no benign multilingual or agent-trace set.
  - Corpus counts stayed at 11015 with 0 errors, so they work only as a regression check. InjecAgent is still about 74 % of corpus findings.
- **Skipped optimizations:**
  - Latency is the main open risk. browsesafe p95 rose from 48.8 to 74.3 ms (+25.5 ms, about 2.5 ms per accepted pattern). Iteration 34 (leetspeak) alone added +13.5 ms. At this rate the 100 ms budget will be reached in about 10 iterations.
  - The next iteration should be latency-only: prefilter literals for the leetspeak and space-tolerant patterns, or a cap on regex input size for HTML.
  - Still open: the ML stage (Prompt-Guard-2 22M/86M), the InjecAgent FP spot-check, and `node_modules`/`vendor` in `dependency_paths`.
  - Resolved: `duration.txt` is now written into every report directory (v0.3.34). v0.3.35–v0.3.44 are tagged.
- **Report branches:**
  - The July 2026 archive is mostly done: 14 → 2 local branches. The 2 left are `report-boom-zany-sarcasm-cd4dee48-2026-07-15` and `report-nuclear-llama-tactical-216a-2026-07-15`.
  - October branches grew 23 → 33, so there are now 35 local and 35 on origin. There is 1 worktree, and `.worktrees/` holds about 4 KB.
  - The guidance against one branch per iteration is still not followed.
  - Action (needs owner approval; nothing was deleted in this review): delete or tag-archive the 2 July branches and the October branches older than the last accepted version, locally and on origin. After that, keep one report branch per review round, and keep per-iteration results in `/tmp/safeanalyze-iter/` only.

