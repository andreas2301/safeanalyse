# COMPARISON: v0.3.26 (iteration 22) vs v0.3.25 (previous accepted)

- **Report branch:** `report-deadpool-tacos-unicorn-a71419cc-2026-10-09`
- **Date:** 2026-10-09
- **Change under test:** new `task_switch_announcement` rule (high, en/de). It flags praise or an override ("Well done!", "Great.", "passed the first test", "ignore/forget", "Gut gemacht", "Vergiss") followed within 120 characters on the same line by an announcement of a new or next task/test ("Now comes a new task", "Here comes the second test", "Jetzt kommt der zweite Test", "Nun folgt ein neuer Test").
- **Baseline:** v0.3.25 (iteration 21 measurement, `/tmp/safeanalyze-iter/iter21`; detection code identical to tag `v0.3.25` at `ea43ecf`).
- **Candidate:** iteration 22, measured as the uncommitted working tree on `ea43ecf` in `/tmp/safeanalyze-iter/iter22`, then released unchanged as v0.3.26 (tag `v0.3.26`; the release commit adds only the version bump and docs on top of the measured code).
- **Decision:** **Accepted** under the CLAUDE.md decision rules. Dev F1 sum +0.0493 (all from deepset 0.571 → 0.621). Holdout F1/precision do not regress (deepset-holdout F1 0.414 → 0.432). No precision drop and no new FP on any dev or holdout set. Corpus findings unchanged.

## Why

Steering priority 2: deepset recall had been flat since v0.3.14. Injections in the deepset style praise the model for a fake first task and then announce the next one ("Well done! ... Now comes a new task."). The rule needs both the lead-in and the announcement, so course and exercise text ("Here comes the next test in our series ...", "Jetzt folgt die nächste Aufgabe: Berechnen Sie ...") stays unflagged. Built on dev-set evidence only; no holdout misclassification was inspected. `data_exfiltration_email` is not touched.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.25 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 4.33 | 4.80 |
| deepset | v0.3.26 | 1.000 | 0.450 | 0.621 | 27 | 0 | 56 | 33 | 4.46 | 5.21 |
| llmail | v0.3.25 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.05 | 7.97 |
| llmail | v0.3.26 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.05 | 8.05 |
| browsesafe | v0.3.25 | 0.873 | 0.663 | 0.754 | 199 | 29 | 271 | 101 | 15.24 | 36.95 |
| browsesafe | v0.3.26 | 0.873 | 0.663 | 0.754 | 199 | 29 | 271 | 101 | 16.71 | 39.84 |

`task_switch_announcement` hits (injection / benign): deepset 7/0 (3 of them new TPs); none on llmail or browsesafe.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.25 | 1.000 | 0.261 | 0.414 | 53 | 0 | 343 | 150 | 4.36 | 5.15 |
| deepset-holdout | v0.3.26 | 1.000 | 0.276 | 0.432 | 56 | 0 | 343 | 147 | 4.54 | 5.25 |
| llmail-holdout | v0.3.25 | 1.000 | 0.717 | 0.835 | 215 | 0 | 160 | 85 | 4.83 | 7.34 |
| llmail-holdout | v0.3.26 | 1.000 | 0.717 | 0.835 | 215 | 0 | 160 | 85 | 5.15 | 7.61 |
| browsesafe-holdout | v0.3.25 | 0.861 | 0.597 | 0.705 | 179 | 29 | 271 | 121 | 14.47 | 40.66 |
| browsesafe-holdout | v0.3.26 | 0.861 | 0.597 | 0.705 | 179 | 29 | 271 | 121 | 14.79 | 41.10 |

- `task_switch_announcement` hits: deepset-holdout 3/0 (all new TPs); none on llmail-holdout or browsesafe-holdout.
- The holdout sets were run only for this evaluation.

## Thorough corpus

| Target | v0.3.25 findings | v0.3.26 findings | Δ | v0.3.25 duration_ms | v0.3.26 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 355 | 355 | +0 | 3422 | 3387 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3598 | 3551 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2697 | 2856 |
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 2961 | 2917 |
| duriantaco-skylos | 2264 | 2264 | +0 | 17080 | 16927 |
| promptfoo-scenarios | 11 | 11 | +0 | 2816 | 2789 |
| promptfoo-webagents | 33 | 33 | +0 | 2853 | 2736 |
| **Total** | **11015** | **11015** | **+0** | **35427** | **35163** |

- Findings are identical per target and per source (the new rule produces no corpus hits). Every target's `errors` array is empty.
- Total duration_ms 35427 → 35163 (-264 ms), within run-to-run variance.

## Latency

Fast-mode p95 rose on every set by 0.1 to 2.9 ms (browsesafe dev 36.95 → 39.84 ms, run 2: 40.37 ms; browsesafe-holdout 40.67 → 41.10 ms). The decision rules accept this because labeled F1 improved. All sets stay well inside the 100 ms budget.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 17–20 ms each, using the binary built in this worktree from master v0.3.26. The output is in `redteam/redteam.txt`.
- `TestTaskSwitchAnnouncement` covers 8 positive and 8 benign cases (en/de); two samples were added to `TestPrefilterEquivalence`. `go build ./...` and `go test ./...` pass on v0.3.26.
- Determinism: run 2 of the dev evals gives the same results as run 1 (latency and metadata removed), and its FP/FN JSONL files are byte-identical to run 1.

## Notes

- **Version metadata:** the measurements ran before the version bump, so eval and corpus reports record `safeanalyze_version: 0.3.25` (`scan_mode: fast` / `thorough`). The detection code is identical to tag `v0.3.26`.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json`.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- **Known gaps:** an announcement with no praise/override lead-in on the same line, a lead-in more than 120 characters earlier, other languages, and other phrasings ("let's move on to the next task") are not matched.
