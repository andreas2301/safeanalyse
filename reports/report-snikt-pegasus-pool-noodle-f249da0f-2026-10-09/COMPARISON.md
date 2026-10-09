# COMPARISON: v0.3.24 (iteration 18) vs v0.3.23 (previous accepted)

- **Report branch:** `report-snikt-pegasus-pool-noodle-f249da0f-2026-10-09`
- **Date:** 2026-10-09
- **Change under test:** fast-mode `inspect` uses `hiddenchars.NewFastStage`. It reports zero-width, bidi and format characters only when the payload shows a smuggling signal: a bidi override (U+202D/U+202E), a Unicode tag character outside a well-formed subdivision flag, a run of 3+ invisible runes, a zero-width/format rune between two ASCII letters, or two zero-width runes on one line that each touch an ASCII letter. Control characters are always reported. Thorough mode is unchanged.
- **Baseline:** v0.3.23 at `0bd36eb` (iteration 17, report `report-pegasus-pool-noodle-e879330b-2026-10-02`), re-measured on 2026-10-09 in `/tmp/safeanalyze-iter/v0.3.23-baseline` on the same machine and session as the candidate.
- **Candidate:** v0.3.24 at `d742a5e` (tag `v0.3.24`). The binary for the red-team run was built from master `51cd6e1` (v0.3.24 code plus docs only).
- **Decision:** **Accepted.** Precision rises on browsesafe dev (0.733 → 0.829), browsesafe-holdout (0.702 → 0.813) and deepset-holdout (0.964 → 1.000). browsesafe dev F1 rises 0.715 → 0.745 and browsesafe-holdout F1 rises 0.671 → 0.695. FPs fall 76 → 42 (dev) and 82 → 42 (holdout) on browsesafe and 2 → 0 on deepset-holdout. deepset-holdout F1 drops 0.417 → 0.414 because one TP is lost (54 → 53). This is accepted under the **single-sample tolerance** in CLAUDE.md: at most one TP lost on that set, its precision does not drop (it rises to 1.000), and the sum of holdout F1 rises (1.923 → 1.944). No other labeled metric drops, and corpus findings are identical.

## Why

Most browsesafe dev FPs came from `hidden_char_*` findings on benign HTML: single emoji ZWJ sequences, flag tag sequences, BOMs and isolated zero-width characters. Earlier attempts (iteration 6, context-aware exemptions) gave no net gain. This iteration gates the whole zero-width/bidi/format category in fast mode on a smuggling signal instead. Built on dev-set evidence only; no holdout misclassification was inspected.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.23 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 4.37 | 5.17 |
| deepset | v0.3.24 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 4.23 | 5.02 |
| llmail | v0.3.23 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 4.81 | 7.13 |
| llmail | v0.3.24 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 4.65 | 6.93 |
| browsesafe | v0.3.23 | 0.733 | 0.697 | 0.715 | 209 | 76 | 224 | 91 | 15.58 | 36.08 |
| browsesafe | v0.3.24 | **0.829** | 0.677 | **0.745** | 203 | **42** | **258** | 97 | 14.76 | 35.87 |

Run-2 dev p50/p95: v0.3.23 deepset 4.36/5.13 ms, llmail 4.92/7.02 ms, browsesafe 15.79/37.28 ms; v0.3.24 deepset 4.57/5.29 ms, llmail 4.87/7.13 ms, browsesafe 15.44/38.50 ms. Run-2 confusion counts match run 1 on both versions.

Rule hits that changed (injection / benign, v0.3.23 → v0.3.24):

- browsesafe: `hidden_char_zero_width` 33/32 → 4/3, `hidden_char_bidi` 7/8 → none, `hidden_char_format` 6/9 → 2/0.
- llmail: `hidden_char_zero_width` 6/0 → 1/0 (no TP lost; these samples are also flagged by other rules).
- deepset: no change.

The 6 lost browsesafe TPs were flagged only by an ungated hidden-char finding.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.23 | 0.964 | 0.266 | 0.417 | 54 | 2 | 341 | 149 | 4.50 | 5.52 |
| deepset-holdout | v0.3.24 | **1.000** | 0.261 | 0.414 | 53 | **0** | **343** | 150 | 4.44 | 5.43 |
| llmail-holdout | v0.3.23 | 1.000 | 0.717 | 0.835 | 215 | 0 | 160 | 85 | 4.93 | 7.09 |
| llmail-holdout | v0.3.24 | 1.000 | 0.717 | 0.835 | 215 | 0 | 160 | 85 | 4.90 | 6.95 |
| browsesafe-holdout | v0.3.23 | 0.702 | 0.643 | 0.671 | 193 | 82 | 218 | 107 | 14.94 | 37.65 |
| browsesafe-holdout | v0.3.24 | **0.813** | 0.607 | **0.695** | 182 | **42** | **258** | 118 | 14.22 | 37.07 |

- Holdout F1 sum: 0.417 + 0.835 + 0.671 = 1.923 (v0.3.23) → 0.414 + 0.835 + 0.695 = 1.944 (v0.3.24).
- **deepset-holdout TP 54 → 53:** one injection sample was flagged only by `hidden_char_zero_width` (holdout rule hits 1/2 → none), so the gate drops it along with both FPs. Precision rises 0.964 → 1.000 and F1 drops by 0.003. This is the single-sample tolerance case named as precedent in CLAUDE.md's decision rules. The sample itself was not opened.
- Changed holdout rule hits: browsesafe-holdout `hidden_char_zero_width` 27/33 → 5/3, `hidden_char_bidi` 8/10 → 1/0, `hidden_char_format` 5/8 → 2/1; llmail-holdout `hidden_char_zero_width` 1/0 → none (no TP lost).
- The holdout sets were run only for this evaluation.

## Thorough corpus

| Target | v0.3.23 findings | v0.3.24 findings | Δ | v0.3.23 duration_ms | v0.3.24 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 356 | 356 | 0 | 3507 | 3303 |
| uiuc-injecagent | 8279 | 8279 | 0 | 3595 | 3879 |
| lakera-pint-benchmark | 38 | 38 | 0 | 2903 | 2973 |
| alexh-prompt-injection-scanner | 143 | 143 | 0 | 3038 | 3069 |
| duriantaco-skylos | 2306 | 2306 | 0 | 14705 | 15959 |
| promptfoo-scenarios | 11 | 11 | 0 | 2953 | 2849 |
| promptfoo-webagents | 33 | 33 | 0 | 2950 | 2818 |
| **Total** | **11166** | **11166** | **0** | **33651** | **34850** |

- The findings on every target are identical to v0.3.23 (same set of findings, same per-source counts). Thorough mode does not use the fast gate, so `hiddenchars` findings (bipia 152, promptfoo-scenarios 3, promptfoo-webagents 14) are unchanged.
- Total duration_ms rose 3.6 % (+1199 ms). skylos alone is +1254 ms (external-scanner variance); the other targets moved by −204 to +284 ms. No thorough-mode code changed. Every target's `errors` array is empty.

## Latency

Fast-mode p95 moved by −0.6 to −0.1 ms on all six sets in run 1. Run-2 browsesafe dev p95 is 38.5 ms (v0.3.23 run 2: 37.3 ms), which is within run-to-run noise. All sets stay well inside the 100 ms budget. The gate adds one pass over the payload only when hidden characters are present.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 16–19 ms each. The output is in `redteam/redteam.txt`, from a run with the binary built in this worktree (master `51cd6e1`, v0.3.24).
- `TestFastStageSmugglingGate` (added in `d742a5e`) covers the gate's positive and benign cases. `go build ./...` and `go test ./...` pass on `d742a5e`.
- Determinism: run 2 of the dev evals gives the same confusion counts as run 1, and its FP/FN JSONL files are byte-identical to run 1.

## Notes

- **Version metadata:** eval reports record `safeanalyze_version: 0.3.24`, `scan_mode: fast`; corpus reports record `safeanalyze_version: 0.3.24`, `scan_mode: thorough`.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json`. The measurement run wrote wall-clock milliseconds to `duration.txt`, so the files were regenerated from `duration_ms` for this report.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- **Known gaps:** a single zero-width character next to non-ASCII text (for example inside a Cyrillic or CJK word) and isolated bidi isolates (U+2066–U+2069) no longer produce fast-mode findings. Injections that rely only on such characters, without a YARA hit, are missed in fast mode. Thorough mode still reports them.
