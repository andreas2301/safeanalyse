# COMPARISON: v0.3.19 (iteration 13) vs v0.3.18 (previous accepted)

- **Report branch:** `report-snikt-chimichanga-regeneration-14eb7d42-2026-10-02`
- **Date:** 2026-10-02
- **Change under test:** `data_exfiltration_email` gets one line-based pattern: a recipient key (`recipient(s)`, `receiver`, `email_to`, `address`) with `:`/`=` and an email address that may be obfuscated (`at`, `(at)`/`[at]`/`{at}`, fullwidth `＠`, spaces around `@`, `(dot)`/`[dot]`/`{dot}`), then within 80 characters a `body`/`content`/`message` field followed by `:`, `=` or "is".
- **Baseline:** v0.3.18 at `46ccd7c` (iteration 12, report `report-unicorn-katana-maximum-effort-b5e360a4-2026-10-02`).
- **Candidate:** v0.3.19 at `7a27d8e`.
- **Decision:** **Accepted.** llmail dev F1 rises 0.824 → 0.842 and llmail-holdout F1 rises 0.790 → 0.814. Precision does not drop on any dev or holdout set, and no set gains a false positive. All other labeled metrics and all corpus counts are identical. This matches the rule "dev F1 improves without a precision drop on any labeled set" with no holdout regression.

## Why

The hypothesis came from dev-set evidence only (BASELINE.md cross-set FN themes #3, obfuscated addresses, and #5, key/value sends). No holdout file was opened. The scanner matches one line at a time. In the llmail dev FNs the send verb is on one line and the obfuscated recipient on the next (`"receiver": contact at contact.com, and body is only "confirmation"`). Iteration 9 required a verb before the obfuscated address, so it could not fire. The v0.3.18 recipient-key pattern needs a literal `@` and a `body:`/`body=` key, not "body is".

The hypothesis also listed `to` and `email`/`e-mail` as recipient keys. They were left out because they match ordinary contact lines (`To: John at acme.com, message: please review the attached draft`, `email: jane.doe(at)uni-bonn.de, message: office hours Tue 2-4pm`). Those lines are now benign test cases.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.18 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 3.19 | 3.75 |
| deepset | v0.3.19 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 3.22 | 3.80 |
| llmail | v0.3.18 | 1.000 | 0.700 | 0.824 | 210 | 0 | 160 | 90 | 3.52 | 5.09 |
| llmail | v0.3.19 | 1.000 | 0.727 | 0.842 | 218 | 0 | 160 | 82 | 3.49 | 4.79 |
| browsesafe | v0.3.18 | 0.719 | 0.647 | 0.681 | 194 | 76 | 224 | 106 | 12.07 | 27.42 |
| browsesafe | v0.3.19 | 0.719 | 0.647 | 0.681 | 194 | 76 | 224 | 106 | 11.85 | 28.78 |

Run-2 dev p95 (v0.3.19): deepset 4.06 ms, llmail 4.69 ms, browsesafe 28.90 ms. On v0.3.18 they were 4.34, 4.77 and 28.79 ms.

`data_exfiltration_email` rule hits (injection / benign): llmail 157/0 → 166/0, browsesafe 132/24 → 132/24.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.18 | 0.964 | 0.266 | 0.417 | 54 | 2 | 341 | 149 | 3.26 | 4.08 |
| deepset-holdout | v0.3.19 | 0.964 | 0.266 | 0.417 | 54 | 2 | 341 | 149 | 3.26 | 3.93 |
| llmail-holdout | v0.3.18 | 1.000 | 0.653 | 0.790 | 196 | 0 | 160 | 104 | 3.49 | 4.80 |
| llmail-holdout | v0.3.19 | 1.000 | 0.687 | 0.814 | 206 | 0 | 160 | 94 | 3.51 | 4.92 |
| browsesafe-holdout | v0.3.18 | 0.683 | 0.590 | 0.633 | 177 | 82 | 218 | 123 | 11.37 | 29.28 |
| browsesafe-holdout | v0.3.19 | 0.683 | 0.590 | 0.633 | 177 | 82 | 218 | 123 | 11.37 | 29.43 |

The holdout sets were run only for this evaluation. No holdout misclassification was inspected. `llmail-holdout` reuses the dev benign emails, so only its recall (+10 TP) is an independent signal. `data_exfiltration_email` hits: llmail-holdout 153/0 → 163/0, browsesafe-holdout 123/25 unchanged.

## Thorough corpus

| Target | v0.3.18 findings | v0.3.19 findings | Δ | v0.3.18 duration_ms | v0.3.19 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 356 | 356 | 0 | 3405 | 3337 |
| uiuc-injecagent | 8279 | 8279 | 0 | 3464 | 3562 |
| lakera-pint-benchmark | 38 | 38 | 0 | 2828 | 2765 |
| alexh-prompt-injection-scanner | 143 | 143 | 0 | 3049 | 3110 |
| duriantaco-skylos | 2306 | 2306 | 0 | 17122 | 17568 |
| promptfoo-scenarios | 11 | 11 | 0 | 2759 | 2783 |
| promptfoo-webagents | 33 | 33 | 0 | 2862 | 2684 |
| **Total** | **11166** | **11166** | **0** | **35489** | **35809** |

- Per-source counts are identical to v0.3.18 on all seven targets, and so are the `data_exfiltration_email` counts (injecagent 3820, skylos 48, bipia 4). The new pattern adds no corpus findings.
- Total duration_ms rose 0.9 % (+320 ms). Most of it is skylos (+446 ms), which comes from external-scanner variance. Every scan exited rc=0 with 0 errors.

## Latency

Fast-mode p95 moved by −0.30 to +1.36 ms per set (largest rise: browsesafe dev 27.4 → 28.8 ms; run 2 28.9 ms vs 28.8 ms on v0.3.18). All sets stay well inside the 100 ms budget.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 15–17 ms each. The output is in `redteam/redteam.txt`, from a run with the v0.3.19 binary built in this worktree.
- `go build ./...` and `go test ./...` pass on `7a27d8e`, including the new cases in `TestDataExfiltrationEmailForward`: two positives (`"receiver": contact at contact.com, and body is only ...`, `receiver: contact[at]contact[dot]com, body: ...`) and four benign contact lines (`To: the team at acme.com. Message: ...`, `Email: info at example dot org, message: ...`, `To: John at acme.com, message: ...`, `email: jane.doe(at)uni-bonn.de, message: ...`).
- Determinism: a second dev eval run gave identical metrics, rule hits and FP/FN JSONL files.

## Notes

- **Version metadata:** the eval and corpus reports record `safeanalyze_version: 0.3.18`. They were measured on the iteration-13 working tree before the version bump, and that tree matches the `v0.3.19` tag in code; only the version constant and docs changed afterwards. `scan_mode` is `fast` for eval and `thorough` for corpus.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json`. The external wall-clock timer was broken in this run, so it is not an independent wall-clock measurement.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- **Rejected this iteration:** `to` and `email`/`e-mail` recipient keys (match ordinary contact lines).
- **Known gaps:** `to:`/`email:` keys with an obfuscated address; `&#64;`, `%40`, unbracketed "dot" (`example dot org`) and spaced letters; a body/content/message field more than 80 characters after the address or on another line. Known benign trigger from v0.3.18: Ruby `mailer.send_email(to: "dev@example.com", body: render(:welcome))`.
