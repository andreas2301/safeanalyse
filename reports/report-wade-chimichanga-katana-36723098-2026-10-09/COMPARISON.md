# COMPARISON: v0.3.31 (iteration 28) vs v0.3.30 (previous accepted)

- **Report branch:** `report-wade-chimichanga-katana-36723098-2026-10-09`
- **Date:** 2026-10-09
- **Change under test:** one new pattern in the existing `prompt_injection_comment` rule for English "ignore/forget/disregard what/whatever I/we/you said/wrote/told you/mentioned/asked/typed/discussed before/earlier/previously/above/so far/until now" overrides and the short form "... what's/what is/was/came/comes before/above/...". It also matches the misspellings "igmre" (`ig??re`), "ingore", "waht" and "wat". The verb must open an instruction (start of text, after `\n . ! ? : ; , > " ' * -`, or after please/now/just/and/then/so/also), a time word is required, and the short form also needs a sentence end, "and" or a capitalized next word. No new rule name.
- **Baseline:** v0.3.30 (iteration 26 measurement, `/tmp/safeanalyze-iter/iter26`; detection code identical to tag `v0.3.30` at `4b26989`).
- **Candidate:** iteration 28, measured as the uncommitted working tree on `4b26989` in `/tmp/safeanalyze-iter/iter28`, then released unchanged as v0.3.31 (tag `v0.3.31` at `1a4fb5a`; the release commit adds only the version bump and docs on top of the measured code).
- **Decision:** **Accepted** under the CLAUDE.md decision rules. Dev F1 sum +0.0500 (all from deepset 0.750 → 0.800). No precision drop on any dev or holdout set; all holdout sets unchanged. Corpus findings unchanged.

## Why

Steering priority 2: deepset recall. The German, Spanish, French and Portuguese/Italian override patterns already cover "what I said before" ("alles, was ich davor ...", "lo que te dije", "ce que je t'ai dit"); English only covered "everything said before". The instruction-start and time-word guards keep narrative and conditional text ("I'll never forget what you said before you left", "You can ignore what I said before if the build passes", "Ignore what's above the threshold in the chart.") unflagged. Built on dev-set evidence only; no holdout misclassification was inspected.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.30 | 1.000 | 0.600 | 0.750 | 36 | 0 | 56 | 24 | 4.90 | 5.58 |
| deepset | v0.3.31 | 1.000 | 0.667 | 0.800 | 40 | 0 | 56 | 20 | 5.27 | 5.80 |
| llmail | v0.3.30 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.66 | 8.51 |
| llmail | v0.3.31 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.88 | 9.21 |
| browsesafe | v0.3.30 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 18.12 | 45.80 |
| browsesafe | v0.3.31 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 19.00 | 46.00 |

`prompt_injection_comment` hits (injection / benign): deepset 10/0 → 14/0; llmail 17/0 and browsesafe 128/1 unchanged.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.30 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.29 | 6.24 |
| deepset-holdout | v0.3.31 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.37 | 6.25 |
| llmail-holdout | v0.3.30 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 5.74 | 8.38 |
| llmail-holdout | v0.3.31 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.08 | 9.10 |
| browsesafe-holdout | v0.3.30 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 17.05 | 46.10 |
| browsesafe-holdout | v0.3.31 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 17.40 | 47.12 |

- `prompt_injection_comment` hits unchanged on every holdout set (deepset-holdout 28/0, llmail-holdout 9/0, browsesafe-holdout 104/2); the new pattern adds no holdout hit.
- The holdout sets were run only for this evaluation.

## Thorough corpus

| Target | v0.3.30 findings | v0.3.31 findings | Δ | v0.3.30 duration_ms | v0.3.31 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 355 | 355 | +0 | 3565 | 3354 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3457 | 3565 |
| lakera-pint-benchmark | 38 | 38 | +0 | 3017 | 2823 |
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 3156 | 2986 |
| duriantaco-skylos | 2264 | 2264 | +0 | 18241 | 16771 |
| promptfoo-scenarios | 11 | 11 | +0 | 3054 | 2867 |
| promptfoo-webagents | 33 | 33 | +0 | 2727 | 2711 |
| **Total** | **11015** | **11015** | **+0** | **37217** | **35077** |

- Findings are identical per target and per source. Every target's `errors` array is empty.
- Total duration_ms 37217 → 35077 (-2140 ms), mostly skylos (-1470 ms, external scanners); run variance, not the change.

## Latency

Fast-mode p95 moved between +0.01 and +1.02 ms (browsesafe dev 45.80 → 46.00 ms, run 2: 47.20 ms; browsesafe-holdout 46.10 → 47.12 ms; llmail 8.51 → 9.21 ms). All sets stay well inside the 100 ms budget.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 19–22 ms each, using the binary built in this worktree from master v0.3.31. The output is in `redteam/redteam.txt`.
- `TestPromptInjectionCommentOverrideVariants` gains 8 positive and 6 benign cases; `TestPrefilterEquivalence` gains 6 inputs. `go build ./...` and `go test ./...` pass on v0.3.31.
- Determinism: run 2 of the dev evals gives the same results as run 1 (latency and metadata removed), and its FP/FN JSONL files are byte-identical to run 1.

## Notes

- **Version metadata:** the measurements ran before the version bump, so eval and corpus reports record `safeanalyze_version: 0.3.30` (`scan_mode: fast` / `thorough`). The detection code is identical to tag `v0.3.31`.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json` (the scan itself did not write it).
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- Iteration 27 (gate bare `get` in `account_access_request` to instruction starts) was reverted before this iteration: browsesafe-holdout F1 0.708 → 0.702. It has no report branch.
- **Known gaps:** overrides without a time word ("ignore what I said"), other verbs ("never mind what I said", "scratch that") and non-instruction-start positions are not matched. The gain is on deepset dev only; no holdout set moved.
