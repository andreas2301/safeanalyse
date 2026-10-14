# COMPARISON: v0.3.44 (iteration 40) vs v0.3.43 (previous accepted)

- **Report branch:** `report-maximum-effort-taco-pegasus-abfca521-2026-10-14`
- **Date:** 2026-10-14
- **Change under test:** one new pattern in the existing `prompt_injection_comment` rule for a line that tells the model to obey whatever text comes next: anchor (text start, `.`/`?`/`!`/`:` + whitespace, or a literal `\n`/`\t` escape) + up to three non-word characters + optional please/now/just + "follow" + "all"/"every" + optional "the"/"of the" + instruction(s)/directive(s)/command(s) + in/from/within/contained in/given in + "the following" + optional text/message/email/content/input/prompt/block + `:`. Each word tolerates single spaces between letters. Case-insensitive. No new rule name.
- **Baseline:** v0.3.43 (tag `v0.3.43`, measurements in `/tmp/safeanalyze-iter/iter39`).
- **Candidate:** iteration 40. It was measured as the uncommitted working tree on `0dfee62` in `/tmp/safeanalyze-iter/iter40` (dev rerun in `iter40/run2`), so its report metadata still reads `safeanalyze_version: 0.3.43`. It was then released unchanged as v0.3.44 (tag `v0.3.44`); the release commit adds only the version bump and docs on top of the measured code.
- **Decision:** **Accepted** under the CLAUDE.md decision rules. PRIMARY dev F1 rises on llmail from 0.881 to 0.889 (TP 236 → 240, FP 0 unchanged); all other PRIMARY dev sets are unchanged. llmail-holdout F1 rises 0.833 → 0.837 (TP 214 → 216, FP 0); every other holdout set is unchanged. No set has a precision drop. Corpus findings are unchanged.

## Why

PRIMARY llmail recall was 0.787 on dev. Four of the 64 `fn-llmail.jsonl` dev misses open with a hand-off line such as "Follow all instructions in the following:", and three of them space out letters inside words ("Fol lo W all  ins tructIonS  in the following:", "Follow  al l instruct i ons in the  following :", "FollO w aLl   instrUc tions in the follOwing:"). Only dev FN/FP files were used (llmail, browsesafe, toolresp, spanish, multilingual). No holdout, GUARD-ONLY or stress misclassification was inspected; for deepset and nemotron only benign label-0 rows on the dev sets were counted. The pattern keys on the hand-off frame, not on any dataset string.

Benign guards: the phrase must start the text, a sentence or a clause after a colon, must use "all"/"every", and "the following" (with an optional noun) must be followed directly by a colon. "Please follow all instructions in the following section to install the driver.", "Please follow all instructions in the following section:", "Applicants should follow all the instructions given in the following: (a) attach a CV" and `<p>Follow all the instructions in the following text:</p>` stay clean. All are covered in `TestPromptInjectionCommentOverrideVariants` (4 positive, 5 benign new cases).

`prompt_injection_comment` hits (injection / benign): llmail 17/0 → 38/0 (+4 TP; the rest were already caught by other rules), llmail-holdout 9/0 → 30/0 (+2 TP). Every other dev, holdout and stress set is unchanged, and no other rule's hit counts changed.

## Labeled fast-mode eval: PRIMARY dev

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail | v0.3.43 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 8.63 | 12.60 |
| llmail | v0.3.44 | 1.000 | 0.800 | 0.889 | 240 | 0 | 160 | 60 | 8.80 | 13.20 |
| browsesafe | v0.3.43 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 29.19 | 74.13 |
| browsesafe | v0.3.44 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 29.83 | 74.31 |
| toolresp | v0.3.43 | 0.859 | 0.630 | 0.727 | 189 | 31 | 269 | 111 | 8.06 | 10.30 |
| toolresp | v0.3.44 | 0.859 | 0.630 | 0.727 | 189 | 31 | 269 | 111 | 8.32 | 10.70 |
| spanish | v0.3.43 | 1.000 | 0.712 | 0.831 | 37 | 0 | 300 | 15 | 8.42 | 9.80 |
| spanish | v0.3.44 | 1.000 | 0.712 | 0.831 | 37 | 0 | 300 | 15 | 8.29 | 9.51 |

## Labeled fast-mode eval: PRIMARY holdout

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail-holdout | v0.3.43 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 8.66 | 13.12 |
| llmail-holdout | v0.3.44 | 1.000 | 0.720 | 0.837 | 216 | 0 | 160 | 84 | 8.59 | 12.79 |
| browsesafe-holdout | v0.3.43 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 26.78 | 72.21 |
| browsesafe-holdout | v0.3.44 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 27.24 | 73.56 |
| toolresp-holdout | v0.3.43 | 0.818 | 0.523 | 0.638 | 157 | 35 | 265 | 143 | 8.16 | 10.49 |
| toolresp-holdout | v0.3.44 | 0.818 | 0.523 | 0.638 | 157 | 35 | 265 | 143 | 8.13 | 10.49 |
| spanish-holdout | v0.3.43 | 1.000 | 0.434 | 0.605 | 23 | 0 | 371 | 30 | 8.06 | 9.25 |
| spanish-holdout | v0.3.44 | 1.000 | 0.434 | 0.605 | 23 | 0 | 371 | 30 | 8.28 | 9.94 |

## Labeled fast-mode eval: RECALL-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| multilingual | v0.3.43 | 1.000 | 0.025 | 0.049 | 21 | 0 | 0 | 819 | 7.89 | 9.39 |
| multilingual | v0.3.44 | 1.000 | 0.025 | 0.049 | 21 | 0 | 0 | 819 | 7.88 | 9.19 |
| multilingual-holdout | v0.3.43 | 1.000 | 0.024 | 0.047 | 20 | 0 | 0 | 805 | 7.84 | 9.10 |
| multilingual-holdout | v0.3.44 | 1.000 | 0.024 | 0.047 | 20 | 0 | 0 | 805 | 7.88 | 8.82 |

## Labeled fast-mode eval: GUARD-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.43 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 7.99 | 8.96 |
| deepset | v0.3.44 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 8.22 | 9.19 |
| deepset-holdout | v0.3.43 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 8.05 | 9.29 |
| deepset-holdout | v0.3.44 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 8.03 | 8.98 |
| nemotron | v0.3.43 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 8.10 | 9.35 |
| nemotron | v0.3.44 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 8.32 | 9.45 |
| nemotron-holdout | v0.3.43 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 8.18 | 9.29 |
| nemotron-holdout | v0.3.44 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 8.12 | 9.27 |

## Labeled fast-mode eval: REPORT-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| promptprotection-stress | v0.3.43 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 7.87 | 8.98 |
| promptprotection-stress | v0.3.44 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 8.35 | 9.81 |

## Thorough corpus

| Target | v0.3.43 findings | v0.3.44 findings | Δ | v0.3.43 duration_ms | v0.3.44 duration_ms |
|---|---|---|---|---|---|
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 2969 | 3081 |
| duriantaco-skylos | 2264 | 2264 | +0 | 16077 | 17154 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2855 | 2832 |
| microsoft-bipia | 355 | 355 | +0 | 3356 | 3354 |
| promptfoo-scenarios | 11 | 11 | +0 | 2879 | 2891 |
| promptfoo-webagents | 33 | 33 | +0 | 2892 | 2934 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3381 | 3481 |
| **Total** | **11015** | **11015** | **+0** | **34409** | **35727** |

- Findings are identical per target and per rule (`prompt_injection_comment` 3310 → 3310). Every target's `errors` array is empty, and every corpus dir has a `duration.txt`.
- Total `duration_ms` went 34409 → 35727 (+1318 ms), almost all from skylos (external scanners); this is run-to-run variance, not the change.

## Latency

Fast-mode p95:
- llmail dev: 12.60 → 13.20 ms (rerun 12.33 ms); llmail-holdout 13.12 → 12.79 ms.
- browsesafe dev: 74.13 → 74.32 ms (rerun 73.37 ms); browsesafe-holdout 72.21 → 73.56 ms.
- toolresp dev: 10.30 → 10.70 ms (rerun 10.24 ms).

All other sets stay under 11 ms, and every set is within the 100 ms budget.

## Red-team

`./scripts/redteam.sh` on the v0.3.44 build in this worktree flagged 12/12 payloads at 23–25 ms (`redteam/redteam.txt`).

## Determinism

The two dev runs (iter40 and iter40/run2) produced byte-identical FP/FN JSONL on all 14 compared files (7 dev sets).

## Files

- `eval/eval-*.json`: fast-mode labeled eval for all dev, holdout and stress sets (FP/FN JSONL not included).
- `corpus/<target>/`: thorough-mode `safeanalyze.{json,md,html,sarif}` and `duration.txt`.
- `redteam/`: red-team output.
- `run.sh`: reproduction script.
