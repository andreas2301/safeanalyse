# COMPARISON: v0.3.32 (iteration 29) vs v0.3.31 (previous accepted)

- **Report branch:** `report-deadpool-taco-unicorn-e5f37f32-2026-10-09`
- **Date:** 2026-10-09
- **Change under test:** two new patterns in the existing `output_constraint` rule. (a) Forced fallback answers with a quoted reply: English "if you/u don't/dont/do not know (it/that/this/the answer)(,) (just/simply/then) say/answer/respond/reply/output (with) (only)(:) \"Y\"" (a reply starting with "I " or "I'", such as "I don't know", is exempt), German "wenn du (es/das/die Antwort) nicht weißt, (dann) sag/sage/antworte/schreib(e) (einfach) (mit) \"Y\"", Spanish "si no (lo/la respuesta) sabes, di/responde/contesta/escribe \"Y\"". (b) Refusal suppression: "you are/you're/you supposed to/must/have to (always) answer/respond/reply(,) at all times/no matter what/always" followed by punctuation or end of text. No new rule name.
- **Baseline:** v0.3.31 (iteration 28 measurement, `/tmp/safeanalyze-iter/iter28`; detection code identical to tag `v0.3.31` at `1a4fb5a`).
- **Candidate:** iteration 29, measured as the uncommitted working tree on `1a4fb5a` in `/tmp/safeanalyze-iter/iter29`, then released unchanged as v0.3.32 (tag `v0.3.32` at `3c7b2b8`; the release commit adds only the version bump and docs on top of the measured code).
- **Decision:** **Accepted** under the CLAUDE.md decision rules. Dev F1 sum +0.0462 (all from deepset 0.800 → 0.846). No precision drop on any dev or holdout set; all holdout sets unchanged. Corpus findings unchanged.

## Why

Steering priority 2: deepset recall, refusal/output suppression. Dev-set FNs included instructions that force a fixed fallback answer ("if you don't know, say ...") and that forbid refusing ("you are supposed to answer at all times"). Requiring a quoted reply and excluding replies that start with "I " / "I'" keeps honest RAG/FAQ prompts ("If you don't know the answer, just say \"I don't know\".", "If you don't know the answer, say so.") unflagged. Built on dev-set evidence only; no holdout misclassification was inspected.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.31 | 1.000 | 0.667 | 0.800 | 40 | 0 | 56 | 20 | 5.27 | 5.80 |
| deepset | v0.3.32 | 1.000 | 0.733 | 0.846 | 44 | 0 | 56 | 16 | 5.26 | 6.13 |
| llmail | v0.3.31 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.88 | 9.21 |
| llmail | v0.3.32 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 6.06 | 9.33 |
| browsesafe | v0.3.31 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 19.00 | 46.00 |
| browsesafe | v0.3.32 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 19.20 | 48.33 |

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.31 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.37 | 6.25 |
| deepset-holdout | v0.3.32 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.51 | 6.82 |
| llmail-holdout | v0.3.31 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.08 | 9.10 |
| llmail-holdout | v0.3.32 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.23 | 9.52 |
| browsesafe-holdout | v0.3.31 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 17.40 | 47.12 |
| browsesafe-holdout | v0.3.32 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 17.25 | 45.49 |

- The holdout sets were run only for this evaluation; no holdout metric moved.

## Thorough corpus

| Target | v0.3.31 findings | v0.3.32 findings | Δ | v0.3.31 duration_ms | v0.3.32 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 355 | 355 | +0 | 3354 | 3364 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3565 | 3455 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2823 | 2644 |
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 2986 | 3091 |
| duriantaco-skylos | 2264 | 2264 | +0 | 16771 | 15964 |
| promptfoo-scenarios | 11 | 11 | +0 | 2867 | 2816 |
| promptfoo-webagents | 33 | 33 | +0 | 2711 | 2983 |
| **Total** | **11015** | **11015** | **+0** | **35077** | **34317** |

- Findings are identical per target and per source. Every target's `errors` array is empty and no scan wrote ERROR.txt.
- Total duration_ms 35077 → 34317 (-760 ms), mostly skylos (-807 ms, external scanners); run-to-run variance, not the change.

## Latency

Fast-mode p95 moved between -1.63 and +2.33 ms (browsesafe dev 46.00 → 48.33 ms, run 2: 51.02 ms; browsesafe-holdout 47.12 → 45.49 ms; deepset 5.80 → 6.13 ms). All sets stay well inside the 100 ms budget.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads, using the binary built in this worktree from master v0.3.32. In this worktree run each payload took 31–41 ms (the iteration 29 measurement run gave 19–21 ms; the difference is machine load, as the detection code is identical). Both are within the 100 ms budget. The output is in `redteam/redteam.txt`.
- Added `TestOutputConstraintRefusalSuppression` (11 positive, 8 benign cases). `go build ./...` and `go test ./...` pass on v0.3.32.
- Determinism: run 2 of the dev evals gives the same results as run 1 (latency and metadata removed), and its FP/FN JSONL files are byte-identical to run 1.

## Notes

- **Version metadata:** the measurements ran before the version bump, so eval and corpus reports record `safeanalyze_version: 0.3.31` (`scan_mode: fast` / `thorough`). The detection code is identical to tag `v0.3.32`.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json` (the scan itself did not write it).
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- **Known gaps:** fallback answers without a quoted reply, other conditions ("if you can't answer", "if unsure"), French/Italian/Portuguese forms, and refusal suppression without answer/respond/reply ("never refuse") are not matched. The gain is on deepset dev only; no holdout set moved.
