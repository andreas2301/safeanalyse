# COMPARISON: v0.3.33 (iteration 30) vs v0.3.32 (previous accepted)

- **Report branch:** `report-merc-mouth-taco-dazzler-d6595596-2026-10-09`
- **Date:** 2026-10-09
- **Change under test:** the last pattern of `prompt_injection_comment` gains a third alternative. "new instruction(s)" or "new statement" now matches at the start of the text or after `.`, `!` or `?` plus whitespace, when followed by `:`, `-`, `–` or `—`, an optional please/now/just, and an override or output verb (ignore, forget, disregard, print, return, output, tell, show, reveal, say, send, repeat). The existing "new directive:" and text-initial "new instructions:" alternatives are unchanged. No new rule name.
- **Baseline:** v0.3.32 (iteration 29 measurement, `/tmp/safeanalyze-iter/iter29`; detection code identical to tag `v0.3.32` at `3c7b2b8`).
- **Candidate:** iteration 30, measured as the uncommitted working tree on `3c7b2b8` in `/tmp/safeanalyze-iter/iter30`, then released unchanged as v0.3.33 (tag `v0.3.33` at `2236329`; the release commit adds only the version bump and docs on top of the measured code).
- **Decision:** **Accepted** under the CLAUDE.md decision rules. Dev F1 sum +0.0218 (all from deepset 0.846 → 0.868). No precision drop on any dev or holdout set; all holdout sets unchanged. Corpus findings unchanged.

## Why

Steering priority 2: deepset recall. Two deepset dev false negatives use a "new statement - please return your embeddings" marker, once at the start of the text and once after a German question ending in `?`. The old alternative only matched "new instructions:" at the very start of the text. Requiring a sentence boundary and an override/output verb keeps notices such as "Bank news. New statement: available in your portal", "Your account is ready. New statement - please download it from the portal." and "Firmware 2.1 released. New instructions: hold the reset button for 5 seconds." unflagged. Built on dev-set evidence only; no holdout misclassification was inspected.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.32 | 1.000 | 0.733 | 0.846 | 44 | 0 | 56 | 16 | 5.26 | 6.13 |
| deepset | v0.3.33 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 5.62 | 6.65 |
| llmail | v0.3.32 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 6.06 | 9.33 |
| llmail | v0.3.33 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 6.06 | 8.95 |
| browsesafe | v0.3.32 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 19.20 | 48.33 |
| browsesafe | v0.3.33 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 19.61 | 49.80 |

`prompt_injection_comment` hits (injection / benign): deepset 14/0 → 16/0; llmail 17/0 → 17/0; browsesafe 128/1 → 128/1.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.32 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.51 | 6.82 |
| deepset-holdout | v0.3.33 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.62 | 6.58 |
| llmail-holdout | v0.3.32 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.23 | 9.52 |
| llmail-holdout | v0.3.33 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.00 | 9.02 |
| browsesafe-holdout | v0.3.32 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 17.25 | 45.49 |
| browsesafe-holdout | v0.3.33 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 18.30 | 48.41 |

- `prompt_injection_comment` hits: deepset-holdout 28/0 → 28/0, llmail-holdout 9/0 → 9/0, browsesafe-holdout 104/2 → 104/2.
- The holdout sets were run only for this evaluation.

## Thorough corpus

| Target | v0.3.32 findings | v0.3.33 findings | Δ | v0.3.32 duration_ms | v0.3.33 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 355 | 355 | +0 | 3364 | 3425 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3455 | 3704 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2644 | 2949 |
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 3091 | 3071 |
| duriantaco-skylos | 2264 | 2264 | +0 | 15964 | 14107 |
| promptfoo-scenarios | 11 | 11 | +0 | 2816 | 2893 |
| promptfoo-webagents | 33 | 33 | +0 | 2983 | 2752 |
| **Total** | **11015** | **11015** | **+0** | **34317** | **32901** |

- Findings are identical per target and per source. Every target's `errors` array is empty.
- Total duration_ms 34317 → 32901 (-1416 ms), mostly skylos (external scanners); run variance, not the change.
- The iteration-30 scans did not write `duration.txt`; the `duration.txt` files in this report were derived from `duration_ms` in each `safeanalyze.json` (as `run.sh` does).

## Latency

Fast-mode p95: deepset 6.13 → 6.65 ms, llmail 9.33 → 8.95 ms, browsesafe dev 48.33 → 49.80 ms (run 2: 46.87 ms), browsesafe-holdout 45.49 → 48.41 ms. All sets stay well inside the 100 ms budget.

## Red-team

`./scripts/redteam.sh` on the v0.3.33 build in this worktree: 12/12 payloads flagged (see `redteam/redteam.txt`).

## Determinism

A second dev eval run gave identical TP/FP/TN/FN, precision, recall, F1, rule hits and per-source counts on all three dev sets; FP/FN JSONL files were byte-identical.
