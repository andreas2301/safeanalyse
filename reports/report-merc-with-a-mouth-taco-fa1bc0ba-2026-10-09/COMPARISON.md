# COMPARISON: v0.3.25 (iteration 21) vs v0.3.24 (previous accepted)

- **Report branch:** `report-merc-with-a-mouth-taco-fa1bc0ba-2026-10-09`
- **Date:** 2026-10-09
- **Change under test:** `data_exfiltration_email` patterns 0 and 4 are narrowed. Pattern 0 (retrieve verb + mail) now needs an explicit `send`/`forward` within 60 characters, "and|then (e)mail", or `email`/`mail` within four words of the verb. Pattern 4 (retrieve verb + data noun) needs the noun within four words of the verb instead of within 80 characters.
- **Baseline:** v0.3.24 at master `51cd6e1`, measured on 2026-10-09 in `/tmp/safeanalyze-iter/v0.3.24-baseline`.
- **Candidate:** iteration 21, measured as the uncommitted working tree on `51cd6e1` in `/tmp/safeanalyze-iter/iter21`, then released unchanged as v0.3.25 at `ea43ecf` (tag `v0.3.25`; the release commit adds only the version bump and docs).
- **Decision:** **Accepted.** browsesafe precision rises on dev (0.829 → 0.873) and holdout (0.813 → 0.861). browsesafe F1 rises on dev (0.745 → 0.754) and holdout (0.695 → 0.705). Dev F1 sum +0.0088. No precision or F1 drops on any dev or holdout set. The corpus drop (−151) is all `data_exfiltration_email` and matches the FP reduction on the labeled sets. This is a retest of iteration 20, which had been rejected only by an undocumented 0.01 gain threshold that is not in the CLAUDE.md decision rules.

## Why

Most remaining browsesafe dev FPs from `data_exfiltration_email` were navigation and footer text where a retrieve verb (get, view, check, access) and a mail or data word fell within the old 60/80-character window ("Get started ... Contact us by email", "View our privacy policy. Questions? Email support."). The rule now needs the words to be close or an explicit send verb. Built on dev-set evidence only; no holdout misclassification was inspected.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.24 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 5.01 | 6.57 |
| deepset | v0.3.25 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 4.33 | 4.80 |
| llmail | v0.3.24 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 4.64 | 6.92 |
| llmail | v0.3.25 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.05 | 7.97 |
| browsesafe | v0.3.24 | 0.829 | 0.677 | 0.745 | 203 | 42 | 258 | 97 | 15.85 | 37.26 |
| browsesafe | v0.3.25 | 0.873 | 0.663 | 0.754 | 199 | 29 | 271 | 101 | 15.24 | 36.95 |

`data_exfiltration_email` hits (injection / benign, v0.3.24 → v0.3.25): browsesafe 140/24 → 131/9; llmail 197/0 → 195/0 (no TP lost; those samples are also flagged by other rules); deepset none.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.24 | 1.000 | 0.261 | 0.414 | 53 | 0 | 343 | 150 | 4.37 | 5.13 |
| deepset-holdout | v0.3.25 | 1.000 | 0.261 | 0.414 | 53 | 0 | 343 | 150 | 4.36 | 5.15 |
| llmail-holdout | v0.3.24 | 1.000 | 0.717 | 0.835 | 215 | 0 | 160 | 85 | 4.75 | 6.98 |
| llmail-holdout | v0.3.25 | 1.000 | 0.717 | 0.835 | 215 | 0 | 160 | 85 | 4.83 | 7.34 |
| browsesafe-holdout | v0.3.24 | 0.812 | 0.607 | 0.695 | 182 | 42 | 258 | 118 | 14.53 | 38.53 |
| browsesafe-holdout | v0.3.25 | 0.861 | 0.597 | 0.705 | 179 | 29 | 271 | 121 | 14.47 | 40.66 |

- Holdout F1 sum: 0.414 + 0.835 + 0.695 = 1.944 (v0.3.24) → 0.414 + 0.835 + 0.705 = 1.954 (v0.3.25).
- `data_exfiltration_email` hits: browsesafe-holdout 129/25 → 119/12; llmail-holdout 183/0 → 181/0 (no TP lost).
- The holdout sets were run only for this evaluation.

## Thorough corpus

| Target | v0.3.24 findings | v0.3.25 findings | Δ | v0.3.24 duration_ms | v0.3.25 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 356 | 355 | -1 | 3244 | 3422 |
| uiuc-injecagent | 8279 | 8171 | -108 | 3602 | 3598 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2747 | 2697 |
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 3038 | 2961 |
| duriantaco-skylos | 2306 | 2264 | -42 | 14401 | 17080 |
| promptfoo-scenarios | 11 | 11 | +0 | 2851 | 2816 |
| promptfoo-webagents | 33 | 33 | +0 | 2837 | 2853 |
| **Total** | **11166** | **11015** | **-151** | **32720** | **35427** |

- All 151 removed findings are `data_exfiltration_email`: uiuc-injecagent 3820 → 3712, duriantaco-skylos 48 → 6, microsoft-bipia 4 → 3. Every other rule and source has the same count on every target.
- skylos is a code repository; its 42 removed hits were code where a retrieve verb and a mail/file word were near each other by chance. InjecAgent attack lines that ask to retrieve data and email it are still flagged.
- Total duration_ms rose 32720 → 35427 (+2707 ms). skylos alone is +2679 ms (external-scanner variance); the other targets moved by −77 to +178 ms. Every target's `errors` array is empty.

## Latency

Fast-mode p95 changed by −1.8 to +2.1 ms across the six sets (browsesafe dev 37.3 → 37.0 ms, browsesafe-holdout 38.5 → 40.7 ms). Run-2 dev p95: deepset 4.89, llmail 7.37, browsesafe 38.54 ms. All sets stay well inside the 100 ms budget.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 17–19 ms each, using the binary built in this worktree from master `ea43ecf` (v0.3.25). The output is in `redteam/redteam.txt`.
- New cases in `TestDataExfiltrationEmailForward` cover both patterns, with positive and benign examples. `go build ./...` and `go test ./...` pass on `ea43ecf`.
- Determinism: run 2 of the dev evals gives the same confusion counts as run 1, and its FP/FN JSONL files are byte-identical to run 1.

## Notes

- **Version metadata:** the measurements ran before the version bump, so eval and corpus reports record `safeanalyze_version: 0.3.24` (`scan_mode: fast` / `thorough`). The detection code is identical to tag `v0.3.25`.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json`.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- **Known gaps:** a retrieve verb with the mail or data noun more than four words away and no explicit send verb ("retrieve all of the customer's stored billing and shipping addresses") no longer matches patterns 0 or 4. Other `data_exfiltration_email` patterns may still catch such text.
