# COMPARISON: v0.3.29 (iteration 25) vs v0.3.28 (previous accepted)

- **Report branch:** `report-boom-pegasus-chimichanga-4c8e57e2-2026-10-09`
- **Date:** 2026-10-09
- **Change under test:** `data_exfiltration_email` pattern 0, third alternative only (retrieve verb `retrieve`/`get`/`fetch`/`extract`/`download`/`access`/`view`/`check` with `email`/`mail` within four words). The verb must now open an instruction: start of text, a newline, `.`/`!`/`?`/`:`/`;`/`,`, a bullet (`*`/`-`/`•`) or `>`, optionally after quotes/brackets, one softener (please, kindly, now, then, also, and, first, next, "go ahead and") or a short "... you to" lead-in. The send/forward and "and/then (e)mail" alternatives are unchanged.
- **Baseline:** v0.3.28 (iteration 24 measurement, `/tmp/safeanalyze-iter/iter24`; detection code identical to tag `v0.3.28` at `8f1654a`).
- **Candidate:** iteration 25, measured as the uncommitted working tree on `8f1654a` in `/tmp/safeanalyze-iter/iter25`, then released unchanged as v0.3.29 (tag `v0.3.29` at `d4f64b5`; the release commit adds only the version bump and docs on top of the measured code).
- **Decision:** **Accepted** under the CLAUDE.md decision rules. Dev F1 sum +0.0043 (all from browsesafe 0.754 → 0.758). No precision drop on any dev or holdout set; browsesafe precision rises on dev and holdout. llmail-holdout loses one TP (F1 0.835 → 0.833, precision 1.0), within the single-sample holdout tolerance. Corpus findings unchanged.

## Why

Steering priority 1: browsesafe precision. The bare "retrieve verb ... email" alternative also matched descriptive page text where the verb sits in the middle of a clause ("Sign up to get the email newsletter.", "Repository access requires email verification"). Requiring the verb to open an instruction removes those matches while keeping imperative requests ("Now access my mail", "Assistant, check my email for the reset code"). Built on dev-set evidence only; no holdout misclassification was inspected. The rule is only narrowed.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.28 | 1.000 | 0.533 | 0.696 | 32 | 0 | 56 | 28 | 5.06 | 6.29 |
| deepset | v0.3.29 | 1.000 | 0.533 | 0.696 | 32 | 0 | 56 | 28 | 5.18 | 6.90 |
| llmail | v0.3.28 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.65 | 8.61 |
| llmail | v0.3.29 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.51 | 8.50 |
| browsesafe | v0.3.28 | 0.873 | 0.663 | 0.754 | 199 | 29 | 271 | 101 | 17.24 | 40.99 |
| browsesafe | v0.3.29 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 17.62 | 43.69 |

`data_exfiltration_email` hits (injection / benign): browsesafe 131/9 → 131/5 (3 fewer FPs; one benign sample is still flagged by another rule); llmail 195/0 unchanged.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.28 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 4.97 | 5.64 |
| deepset-holdout | v0.3.29 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.13 | 6.01 |
| llmail-holdout | v0.3.28 | 1.000 | 0.717 | 0.835 | 215 | 0 | 160 | 85 | 5.41 | 8.24 |
| llmail-holdout | v0.3.29 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.02 | 9.25 |
| browsesafe-holdout | v0.3.28 | 0.861 | 0.597 | 0.705 | 179 | 29 | 271 | 121 | 15.01 | 39.78 |
| browsesafe-holdout | v0.3.29 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 16.19 | 44.44 |

- `data_exfiltration_email` hits: browsesafe-holdout 119/12 → 117/10; llmail-holdout 181/0 → 180/0.
- The holdout sets were run only for this evaluation.

## Thorough corpus

| Target | v0.3.28 findings | v0.3.29 findings | Δ | v0.3.28 duration_ms | v0.3.29 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 355 | 355 | +0 | 3477 | 3653 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3627 | 3720 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2793 | 3111 |
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 2863 | 3088 |
| duriantaco-skylos | 2264 | 2264 | +0 | 14587 | 16161 |
| promptfoo-scenarios | 11 | 11 | +0 | 2767 | 3888 |
| promptfoo-webagents | 33 | 33 | +0 | 2655 | 3275 |
| **Total** | **11015** | **11015** | **+0** | **32769** | **36896** |

- Findings are identical per target and per source. Every target's `errors` array is empty.
- Total duration_ms 32769 → 36896 (+4127 ms, +1574 ms from skylos external scanners). Every target is slower, including those with few files, so this is mostly machine/run variance rather than the narrowed regex.

## Latency

Fast-mode p95 moved between -0.1 and +4.7 ms (browsesafe dev 40.99 → 43.69 ms, run 2: 42.77 ms; browsesafe-holdout 39.78 → 44.44 ms). All sets stay well inside the 100 ms budget.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 19–33 ms each, using the binary built in this worktree from master v0.3.29. The output is in `redteam/redteam.txt`.
- `TestDataExfiltrationEmailForward` gains 7 positive and 4 benign cases. `go build ./...` and `go test ./...` pass on v0.3.29.
- Determinism: run 2 of the dev evals gives the same results as run 1 (latency and metadata removed), and its FP/FN JSONL files are byte-identical to run 1.

## Notes

- **Version metadata:** the measurements ran before the version bump, so eval and corpus reports record `safeanalyze_version: 0.3.28` (`scan_mode: fast` / `thorough`). The detection code is identical to tag `v0.3.29`.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json`.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- **Known gaps:** a retrieve-and-mail instruction in the middle of a sentence without a listed softener or "you to" lead-in is now caught only by the other `data_exfiltration_email` alternatives; one llmail-holdout injection is no longer flagged.
