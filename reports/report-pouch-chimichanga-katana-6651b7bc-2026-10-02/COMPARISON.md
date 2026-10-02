# COMPARISON: v0.3.17 (iteration 10) vs v0.3.16 (previous accepted)

- **Report branch:** `report-pouch-chimichanga-katana-6651b7bc-2026-10-02`
- **Date:** 2026-10-02
- **Change under test:** `data_exfiltration_email` gets one new pattern: a send / forward / dispatch verb (with -ing forms) followed by whitespace, then, within the same sentence, an output noun (summary, confirmation, keyword, body, content(s), output(s), result(s), reply/replies, response(s), transcript(s)), then "to" and an email address within 40 characters. No `<`, `>`, `=`, `.`, `;`, `?` or `!` may appear between the verb and "to".
- **Baseline:** v0.3.16 at `5940bda` (iteration 8, report `report-regenerating-mercenary-chimichanga-38ddfe85-2026-10-02`).
- **Candidate:** v0.3.17 at `89d44e9`.
- **Decision:** **Accepted.** llmail dev F1 rises 0.785 → 0.798 and llmail-holdout F1 rises 0.776 → 0.778. Precision does not drop on any dev or holdout set, and no set gains a false positive. This matches the rule "dev F1 improves without a precision drop on any labeled set" with no holdout regression.

## Why

The hypothesis came from the llmail dev false negatives only. No holdout file was opened. 66 of the 106 llmail dev FNs contain a plain email address. The existing pattern `send.{0,40}(to|at).{0,30}<email>` misses them because the verb is too far from the address ("forward the summary containing the keyword "confirmation" to contact@..."). A bare verb-then-address pattern added 30–35 browsesafe dev FPs from HTML labels and attributes (`Email: no-reply@...`, `class="sender-email">x@y`). Requiring whitespace after the verb, an explicit "to", and no `<`, `>` or `=` in between removed them.

The hypothesis expected +11 dev TPs (8 llmail, 3 browsesafe). The released pattern also requires an output noun, and it gave +5 llmail and 0 browsesafe.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.16 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 2.59 | 3.07 |
| deepset | v0.3.17 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 2.69 | 3.44 |
| llmail | v0.3.16 | 1.000 | 0.647 | 0.785 | 194 | 0 | 160 | 106 | 2.77 | 3.88 |
| llmail | v0.3.17 | 1.000 | 0.663 | 0.798 | 199 | 0 | 160 | 101 | 3.00 | 4.02 |
| browsesafe | v0.3.16 | 0.719 | 0.647 | 0.681 | 194 | 76 | 224 | 106 | 10.48 | 23.94 |
| browsesafe | v0.3.17 | 0.719 | 0.647 | 0.681 | 194 | 76 | 224 | 106 | 11.04 | 25.43 |

Run-2 dev p95 (v0.3.17): deepset 3.32 ms, llmail 4.60 ms, browsesafe 24.75 ms. On v0.3.16 they were 3.12, 4.02 and 25.89 ms.

`data_exfiltration_email` rule hits (injection / benign): llmail 134/0 → 142/0, browsesafe 130/24 → 132/24. The 24 benign browsesafe hits were already false positives at v0.3.16.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.16 | 0.964 | 0.266 | 0.417 | 54 | 2 | 341 | 149 | 2.64 | 3.43 |
| deepset-holdout | v0.3.17 | 0.964 | 0.266 | 0.417 | 54 | 2 | 341 | 149 | 2.88 | 3.65 |
| llmail-holdout | v0.3.16 | 1.000 | 0.633 | 0.776 | 190 | 0 | 160 | 110 | 2.77 | 3.87 |
| llmail-holdout | v0.3.17 | 1.000 | 0.637 | 0.778 | 191 | 0 | 160 | 109 | 2.90 | 4.22 |
| browsesafe-holdout | v0.3.16 | 0.683 | 0.590 | 0.633 | 177 | 82 | 218 | 123 | 10.19 | 27.31 |
| browsesafe-holdout | v0.3.17 | 0.683 | 0.590 | 0.633 | 177 | 82 | 218 | 123 | 9.95 | 26.40 |

The holdout sets were run only for this evaluation. No holdout misclassification was inspected. `llmail-holdout` reuses the dev benign emails, so only its recall (+1 TP) is an independent signal.

## Thorough corpus

| Target | v0.3.16 findings | v0.3.17 findings | Δ | v0.3.16 duration_ms | v0.3.17 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 356 | 356 | 0 | 3321 | 3421 |
| uiuc-injecagent | 8138 | 8279 | +141 | 3376 | 3542 |
| lakera-pint-benchmark | 38 | 38 | 0 | 2864 | 2775 |
| alexh-prompt-injection-scanner | 143 | 143 | 0 | 3028 | 2908 |
| duriantaco-skylos | 2306 | 2306 | 0 | 15980 | 17361 |
| promptfoo-scenarios | 11 | 11 | 0 | 2827 | 2755 |
| promptfoo-webagents | 33 | 33 | 0 | 2851 | 2820 |
| **Total** | **11025** | **11166** | **+141** | **34247** | **35582** |

- All 141 new findings are `data_exfiltration_email` on uiuc-injecagent (rule count 3679 → 3820; YARA 7827 → 7968). Files: `test_cases_ds_base.json` 68, `test_cases_ds_enhanced.json` 68, `attacker_simulated_responses.json` 3, `attacker_cases_ds.jsonl` 2. A spot check shows matches such as "send a summary to amy.watson@..." and "send the result to amy.watson@...". In `test_cases_ds_base.json` every such string is in an `Attacker Instruction` field or in a `Tool Response` that embeds it. They are true attacker payloads, not FP inflation.
- Per-source counts on the other six targets are identical to v0.3.16. `data_exfiltration_email` is unchanged on bipia (4) and skylos (48).
- Total duration_ms rose 3.9 %. Most of it is skylos (+1381 ms), which is dominated by external scanners and is variance. injecagent rose 166 ms. Every scan exited rc=0 with `errors: []`.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 14–16 ms each. The output is in `redteam/redteam.txt`, from a run with the v0.3.17 binary built in this worktree.
- `go build ./...` and `go test -count=1 ./...` pass on `89d44e9`. That includes the new positive and benign cases in `TestDataExfiltrationEmailForward`, for example HTML `sender-email` spans, `Email:` labels, support footers and "unsubscribe" instructions.
- Determinism: a second dev eval run matched the first once latency and duration fields were removed. The FP/FN JSONL files were byte-identical.

## Notes

- **Version metadata:** the eval and corpus reports in `eval/` and `corpus/` came from the iteration-10 working tree before the version bump. They record `safeanalyze_version: 0.3.16`, with `scan_mode: fast` for eval and `thorough` for corpus. The code under test is the same as the `v0.3.17` tag except for the version constant. `redteam/redteam.txt` comes from the tagged v0.3.17 binary. To get reports stamped 0.3.17, run `run.sh`.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json`.
- **Rejected this iteration:**
  - a bare verb-then-address pattern (30–35 new browsesafe dev FPs);
  - `e-?mail` as a verb (+1 llmail TP, and it would open a path to HTML-label FPs).
- **Known gaps:** other verbs ("notify using mail to me"), addresses more than 40 characters after "to", obfuscated addresses (`(at)`, `＠`, `&#64;`, `%40`), and requests split over several sentences.
