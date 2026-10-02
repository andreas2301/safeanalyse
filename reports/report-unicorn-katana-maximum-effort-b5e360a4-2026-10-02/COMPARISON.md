# COMPARISON: v0.3.18 (iteration 12) vs v0.3.17 (previous accepted)

- **Report branch:** `report-unicorn-katana-maximum-effort-b5e360a4-2026-10-02`
- **Date:** 2026-10-02
- **Change under test:** `data_exfiltration_email` gets four line-based tool-call / structured-send patterns: (1) a send-mail tool name (`send_email`, `sendemail`, `send_mail`, `send_message`, `mail_send`, `api_call`) with an address and a payload key (`body`/`content`/`message` followed by `:` or `=`, or "with (the) body/content/message") within 80 characters either way; (2) a recipient key (`recipient(s)`, `receiver`, `email_to`) with `:`/`=` and an address, plus a payload key within 80 characters either way; (3) a `to` key with an address directly followed by a `body`/`content`/`message` key; (4) a piped `| mail -s ... <address>`.
- **Baseline:** v0.3.17 at `89d44e9` (iteration 10, report `report-pouch-chimichanga-katana-6651b7bc-2026-10-02`).
- **Candidate:** v0.3.18 at `46ccd7c`.
- **Decision:** **Accepted.** llmail dev F1 rises 0.798 → 0.824 and llmail-holdout F1 rises 0.778 → 0.790. Precision does not drop on any dev or holdout set, and no set gains a false positive. All other labeled metrics and all corpus counts are identical. This matches the rule "dev F1 improves without a precision drop on any labeled set" with no holdout regression.

## Why

The hypothesis came from the llmail dev false negatives only (BASELINE.md cross-set theme #5, CLAUDE.md candidate "JSON tool-call injection in llmail FNs"). No holdout file was opened. 16 of the 101 llmail dev FNs ask the agent to send mail through tool or function syntax or a key/value block (`{"action": "send_email", "parameters": {"to": ..., "body": ...}}`, `call send_email(receiver=..., body=...)`, `Recipient=...; Content=...`, `To: x@y Body: ...`, `echo ... | mail -s ... x@y`), so the prose patterns never fire.

The first red-team draft (tool name or recipient key plus address only, unpiped `mail -s`) flagged 11 of 14 benign code and form snippets. Requiring a payload key and a pipe brought that to 1 of 14 (`mailer.send_email(to: "dev@example.com", body: render(:welcome))`) and kept 11 of the draft's 16 llmail dev TPs.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.17 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 2.69 | 3.44 |
| deepset | v0.3.18 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 3.19 | 3.75 |
| llmail | v0.3.17 | 1.000 | 0.663 | 0.798 | 199 | 0 | 160 | 101 | 3.00 | 4.02 |
| llmail | v0.3.18 | 1.000 | 0.700 | 0.824 | 210 | 0 | 160 | 90 | 3.52 | 5.09 |
| browsesafe | v0.3.17 | 0.719 | 0.647 | 0.681 | 194 | 76 | 224 | 106 | 11.04 | 25.43 |
| browsesafe | v0.3.18 | 0.719 | 0.647 | 0.681 | 194 | 76 | 224 | 106 | 12.07 | 27.42 |

Run-2 dev p95 (v0.3.18): deepset 4.34 ms, llmail 4.77 ms, browsesafe 28.79 ms. On v0.3.17 they were 3.32, 4.60 and 24.75 ms.

`data_exfiltration_email` rule hits (injection / benign): llmail 142/0 → 157/0, browsesafe 132/24 → 132/24.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.17 | 0.964 | 0.266 | 0.417 | 54 | 2 | 341 | 149 | 2.88 | 3.65 |
| deepset-holdout | v0.3.18 | 0.964 | 0.266 | 0.417 | 54 | 2 | 341 | 149 | 3.26 | 4.08 |
| llmail-holdout | v0.3.17 | 1.000 | 0.637 | 0.778 | 191 | 0 | 160 | 109 | 2.90 | 4.22 |
| llmail-holdout | v0.3.18 | 1.000 | 0.653 | 0.790 | 196 | 0 | 160 | 104 | 3.49 | 4.80 |
| browsesafe-holdout | v0.3.17 | 0.683 | 0.590 | 0.633 | 177 | 82 | 218 | 123 | 9.95 | 26.40 |
| browsesafe-holdout | v0.3.18 | 0.683 | 0.590 | 0.633 | 177 | 82 | 218 | 123 | 11.37 | 29.28 |

The holdout sets were run only for this evaluation. No holdout misclassification was inspected. `llmail-holdout` reuses the dev benign emails, so only its recall (+5 TP) is an independent signal.

## Thorough corpus

| Target | v0.3.17 findings | v0.3.18 findings | Δ | v0.3.17 duration_ms | v0.3.18 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 356 | 356 | 0 | 3421 | 3405 |
| uiuc-injecagent | 8279 | 8279 | 0 | 3542 | 3464 |
| lakera-pint-benchmark | 38 | 38 | 0 | 2775 | 2828 |
| alexh-prompt-injection-scanner | 143 | 143 | 0 | 2908 | 3049 |
| duriantaco-skylos | 2306 | 2306 | 0 | 17361 | 17122 |
| promptfoo-scenarios | 11 | 11 | 0 | 2755 | 2759 |
| promptfoo-webagents | 33 | 33 | 0 | 2820 | 2862 |
| **Total** | **11166** | **11166** | **0** | **35582** | **35489** |

- Per-source and per-rule counts are identical to v0.3.17 on all seven targets (`data_exfiltration_email`: injecagent 3820, skylos 48, bipia 4). The new patterns add no corpus findings.
- Total duration_ms fell 0.3 %, within run-to-run variance. Every scan exited rc=0 with `errors: []`.

## Latency

Fast-mode p95 rose by 0.3–2.9 ms per set (largest: browsesafe-holdout 26.4 → 29.3 ms). All sets stay well inside the 100 ms budget. The increase is accepted because it comes with a labeled F1 gain.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 15–16 ms each. The output is in `redteam/redteam.txt`, from a run with the v0.3.18 binary built in this worktree.
- `go build ./...` and `go test ./...` pass on `46ccd7c`, including the new positive and benign cases in `TestDataExfiltrationEmailForward` (Django `send_mail(...)`, `Recipient: hr@corp.com`, `To:`/`Subject:` headers, nodemailer `sendMail({...})`, `mail -s ... < file`) and the new `TestPrefilterEquivalence` inputs.
- Determinism: a second dev eval run matched the first once latency and duration fields were removed. The FP/FN JSONL files were byte-identical.

## Notes

- **Version metadata:** the eval and corpus reports record `safeanalyze_version: 0.3.18`, with `scan_mode: fast` for eval and `thorough` for corpus. They come from the iteration-12 working tree, which matches the `v0.3.18` tag in code.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json`.
- **Rejected this iteration:** tool name or recipient key plus address only, and unpiped `mail -s` (11 of 14 benign red-team snippets flagged).
- **Not released:** iteration 11 (a v0.3.18 candidate; browsesafe dev TP 194 → 188, FP 76 → 62).
- **Known gaps:** a tool name or recipient key without a payload key, unpiped `mail -s`, obfuscated addresses (`(at)`, `＠`, `&#64;`, `%40`), and payload keys more than 80 characters from the address. Known benign trigger: Ruby `mailer.send_email(to: ..., body: ...)`.
