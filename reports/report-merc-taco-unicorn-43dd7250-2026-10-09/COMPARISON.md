# COMPARISON: v0.3.30 (iteration 26) vs v0.3.29 (previous accepted)

- **Report branch:** `report-merc-taco-unicorn-43dd7250-2026-10-09`
- **Date:** 2026-10-09
- **Change under test:** three new patterns in the existing `output_constraint` rule for output-hijack game rules: (a) "let's play a game where/in which you (must/will/...) (only) answer/respond/reply/say/pretend/act"; (b) "answer/respond to/reply to every/each/any question/message/prompt/query/request ... with" a reply that starts with a quote character, only after "you must/will/shall/have to/need to" or "from now on"; (c) call-and-response "if/when/whenever I say X, you say "Y"" in English, Spanish ("si yo digo X, tú dices "Y"") and German ("wenn ich X sage, sagst du "Y""), with the reply starting with a quote character. No new rule name; `data_exfiltration_email` untouched.
- **Baseline:** v0.3.29 (iteration 25 measurement, `/tmp/safeanalyze-iter/iter25`; detection code identical to tag `v0.3.29` at `d4f64b5`).
- **Candidate:** iteration 26, measured as the uncommitted working tree on `d4f64b5` in `/tmp/safeanalyze-iter/iter26`, then released unchanged as v0.3.30 (tag `v0.3.30` at `4b26989`; the release commit adds only the version bump and docs on top of the measured code).
- **Decision:** **Accepted** under the CLAUDE.md decision rules. Dev F1 sum +0.0543 (all from deepset 0.696 → 0.750). No precision drop on any dev or holdout set; all holdout sets unchanged. Corpus findings unchanged.

## Why

Steering priority 2: deepset recall (refusal/output suppression). The patterns cover classic jailbreak shapes that dictate the model's replies through a "game" or a fixed call-and-response, not sample strings. Exam/API instructions ("Answer each question with a complete sentence.", "Respond to every request with a JSON body ...") and unquoted replies ("When I say jump, you say how high") stay unflagged. Built on dev-set evidence only; no holdout misclassification was inspected.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.29 | 1.000 | 0.533 | 0.696 | 32 | 0 | 56 | 28 | 5.18 | 6.90 |
| deepset | v0.3.30 | 1.000 | 0.600 | 0.750 | 36 | 0 | 56 | 24 | 4.90 | 5.58 |
| llmail | v0.3.29 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.51 | 8.50 |
| llmail | v0.3.30 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.66 | 8.51 |
| browsesafe | v0.3.29 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 17.62 | 43.69 |
| browsesafe | v0.3.30 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 18.12 | 45.80 |

`output_constraint` hits (injection / benign): deepset 0/0 → 4/0; llmail 1/0 unchanged; none on browsesafe dev.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.29 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.13 | 6.01 |
| deepset-holdout | v0.3.30 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.29 | 6.24 |
| llmail-holdout | v0.3.29 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.02 | 9.25 |
| llmail-holdout | v0.3.30 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 5.74 | 8.38 |
| browsesafe-holdout | v0.3.29 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 16.19 | 44.44 |
| browsesafe-holdout | v0.3.30 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 17.05 | 46.10 |

- `output_constraint` hits unchanged on every holdout set (deepset-holdout 1/0, browsesafe-holdout 1/4); the new patterns add no holdout hit.
- The holdout sets were run only for this evaluation.

## Thorough corpus

| Target | v0.3.29 findings | v0.3.30 findings | Δ | v0.3.29 duration_ms | v0.3.30 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 355 | 355 | +0 | 3653 | 3565 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3720 | 3457 |
| lakera-pint-benchmark | 38 | 38 | +0 | 3111 | 3017 |
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 3088 | 3156 |
| duriantaco-skylos | 2264 | 2264 | +0 | 16161 | 18241 |
| promptfoo-scenarios | 11 | 11 | +0 | 3888 | 3054 |
| promptfoo-webagents | 33 | 33 | +0 | 3275 | 2727 |
| **Total** | **11015** | **11015** | **+0** | **36896** | **37217** |

- Findings are identical per target and per source. Every target's `errors` array is empty.
- Total duration_ms 36896 → 37217 (+321 ms). skylos is +2080 ms (external scanners) while five of the other six targets are faster, so this is run variance rather than the new regexes.

## Latency

Fast-mode p95 moved between -1.3 and +2.1 ms (browsesafe dev 43.69 → 45.80 ms, run 2: 45.18 ms; browsesafe-holdout 44.44 → 46.10 ms). All sets stay well inside the 100 ms budget.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 18–19 ms each, using the binary built in this worktree from master v0.3.30. The output is in `redteam/redteam.txt`.
- New `TestOutputConstraintGameRules` (7 positive, 11 benign cases). `go build ./...` and `go test ./...` pass on v0.3.30.
- Determinism: run 2 of the dev evals gives the same results as run 1 (latency and metadata removed), and its FP/FN JSONL files are byte-identical to run 1.

## Notes

- **Version metadata:** the measurements ran before the version bump, so eval and corpus reports record `safeanalyze_version: 0.3.29` (`scan_mode: fast` / `thorough`). The detection code is identical to tag `v0.3.30`.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json` (the scan itself did not write it).
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- **Known gaps:** game rules with an unquoted reply, answer-every-question instructions without a "you must"/"from now on" lead-in, and call-and-response rules in other languages are not matched. The gain is on deepset dev only; no holdout set moved.
