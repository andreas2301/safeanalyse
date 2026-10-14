# COMPARISON: v0.3.36 (iteration 32) vs v0.3.35 (previous accepted)

- **Report branch:** `report-bad-touch-unicorn-tacos-8fcf811b-2026-10-14`
- **Date:** 2026-10-14
- **Change under test:** one new pattern in the existing `suspicious_shell` rule, after the `mkfifo` pattern: a label colon, at most 3 characters (no backtick, quote or newline), an opening backtick, then within 120 characters on the same line `wh[o0][a4]m[i1]`, `/etc/passwd` or `/etc/shadow`, `sh -c`, `system(`, `wscript.sh` (WScript.Shell) or `nc -l/-e/-z`; case-insensitive. No new rule name.
- **Baseline:** v0.3.35 (tag `v0.3.35`, measurements in `/tmp/safeanalyze-iter/iter31`).
- **Candidate:** iteration 32, measured as the uncommitted working tree on `b8c270a` in `/tmp/safeanalyze-iter/iter32` (so its report metadata still reads `safeanalyze_version: 0.3.35`), then released unchanged as v0.3.36 (tag `v0.3.36`; the release commit adds only the version bump and docs on top of the measured code).
- **Decision:** **Accepted** under the CLAUDE.md decision rules. PRIMARY dev F1 gain from toolresp 0.603 → 0.643 (TP 143 → 157, FP 31 unchanged, precision 0.822 → 0.835); all other PRIMARY dev sets unchanged. No precision drop on any set; every holdout set non-regressing (toolresp-holdout F1 0.577 → 0.607). Corpus findings unchanged.

## Why

PRIMARY toolresp recall was 0.477 on dev. In `fn-toolresp.jsonl`, about 16 of 157 misses were code-execution payloads hidden in JSON tool-response fields behind a label and a backtick span, often in leetspeak (``run: `sh -c 'whoami ...; curl ...'` ``, ``Run: `fortr4n -e 'C411 5YS73M("who4mi")'` ``, ``Bypass: `print(open('/etc/passwd').read())` ``). The old `suspicious_shell` patterns covered only `curl|sh`, `python -c ... import os`, `bash -c ...$(`, `nc -e` and `mkfifo`, with no leetspeak and no notion of a backtick-delimited payload. Built on dev FN/FP files only (toolresp, llmail, browsesafe, spanish, multilingual); no holdout, GUARD-ONLY or stress misclassification was inspected. The HackAPrompt canary "I have been PWNED" (largest single FN theme) was left out on purpose as dataset memorization.

Benign guards: the span must follow a label colon, so documentation such as "Type `cat /etc/shadow` now." or "Start a listener with `nc -l 8080` on the host." stays clean, as does a JSON value ``{"cmd":"`whoami`"}`` (quote between colon and backtick).

`suspicious_shell` hits (injection / benign): toolresp 0/0 → 29/0, toolresp-holdout 0/0 → 28/0, promptprotection-stress 4/0 → 4/0; no other set has hits.

## Labeled fast-mode eval: PRIMARY dev

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail | v0.3.35 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 6.19 | 9.38 |
| llmail | v0.3.36 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 6.53 | 9.62 |
| browsesafe | v0.3.35 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 19.87 | 49.73 |
| browsesafe | v0.3.36 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 20.65 | 50.40 |
| toolresp | v0.3.35 | 0.822 | 0.477 | 0.603 | 143 | 31 | 269 | 157 | 5.74 | 7.34 |
| toolresp | v0.3.36 | 0.835 | 0.523 | 0.643 | 157 | 31 | 269 | 143 | 6.11 | 7.68 |
| spanish | v0.3.35 | 1.000 | 0.385 | 0.556 | 20 | 0 | 300 | 32 | 5.87 | 6.90 |
| spanish | v0.3.36 | 1.000 | 0.385 | 0.556 | 20 | 0 | 300 | 32 | 6.17 | 7.44 |

## Labeled fast-mode eval: PRIMARY holdout

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail-holdout | v0.3.35 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.26 | 9.46 |
| llmail-holdout | v0.3.36 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.44 | 9.64 |
| browsesafe-holdout | v0.3.35 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 17.71 | 48.95 |
| browsesafe-holdout | v0.3.36 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 18.88 | 50.39 |
| toolresp-holdout | v0.3.35 | 0.795 | 0.453 | 0.577 | 136 | 35 | 265 | 164 | 5.91 | 7.74 |
| toolresp-holdout | v0.3.36 | 0.807 | 0.487 | 0.607 | 146 | 35 | 265 | 154 | 6.13 | 7.88 |
| spanish-holdout | v0.3.35 | 1.000 | 0.434 | 0.605 | 23 | 0 | 371 | 30 | 5.88 | 7.21 |
| spanish-holdout | v0.3.36 | 1.000 | 0.434 | 0.605 | 23 | 0 | 371 | 30 | 6.04 | 7.15 |

## Labeled fast-mode eval: RECALL-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| multilingual | v0.3.35 | 1.000 | 0.025 | 0.049 | 21 | 0 | 0 | 819 | 5.75 | 6.77 |
| multilingual | v0.3.36 | 1.000 | 0.025 | 0.049 | 21 | 0 | 0 | 819 | 5.84 | 6.71 |
| multilingual-holdout | v0.3.35 | 1.000 | 0.024 | 0.047 | 20 | 0 | 0 | 805 | 5.77 | 6.71 |
| multilingual-holdout | v0.3.36 | 1.000 | 0.024 | 0.047 | 20 | 0 | 0 | 805 | 5.86 | 6.78 |

## Labeled fast-mode eval: GUARD-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.35 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 5.87 | 7.00 |
| deepset | v0.3.36 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 6.05 | 6.91 |
| deepset-holdout | v0.3.35 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.89 | 6.84 |
| deepset-holdout | v0.3.36 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 6.04 | 6.81 |
| nemotron | v0.3.35 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 5.95 | 7.24 |
| nemotron | v0.3.36 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 6.20 | 7.23 |
| nemotron-holdout | v0.3.35 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 5.84 | 6.67 |
| nemotron-holdout | v0.3.36 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 6.03 | 6.82 |

## Labeled fast-mode eval: REPORT-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| promptprotection-stress | v0.3.35 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 5.66 | 6.52 |
| promptprotection-stress | v0.3.36 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 6.05 | 7.00 |

## Thorough corpus

| Target | v0.3.35 findings | v0.3.36 findings | Δ | v0.3.35 duration_ms | v0.3.36 duration_ms |
|---|---|---|---|---|---|
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 2973 | 3103 |
| duriantaco-skylos | 2264 | 2264 | +0 | 15596 | 13784 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2735 | 2892 |
| microsoft-bipia | 355 | 355 | +0 | 3146 | 3167 |
| promptfoo-scenarios | 11 | 11 | +0 | 3002 | 2882 |
| promptfoo-webagents | 33 | 33 | +0 | 3011 | 2952 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3921 | 3836 |
| **Total** | **11015** | **11015** | **+0** | **34384** | **32616** |

- Findings are identical per target and per (source, rule). Every target's `errors` array is empty, and every corpus dir carries `duration.txt`.
- Total duration_ms 34384 → 32616 (−1768 ms), mostly skylos (15596 → 13784, external scanners); run-to-run variance, not the change.

## Latency

Fast-mode p95: browsesafe dev 49.73 → 50.40 ms (run 2: 49.25 ms), browsesafe-holdout 48.95 → 50.39 ms; toolresp dev 7.34 → 7.68 ms. All other sets < 10 ms. All within the 100 ms budget.

## Red-team

`./scripts/redteam.sh` on the v0.3.36 build in this worktree: 12/12 payloads flagged at 19–22 ms (`redteam/redteam.txt`). The iteration-32 run before release is in `redteam/redteam-iter32.txt` (12/12, 20 ms).

## Determinism

Two dev runs (iter32 and iter32/run2) produced byte-identical FP/FN JSONL on all 14 compared files and identical TP/FP counts.

## Files

- `eval/eval-*.json` — fast-mode labeled eval for all dev, holdout and stress sets (FP/FN JSONL not included).
- `corpus/<target>/` — thorough-mode `safeanalyze.{json,md,html,sarif}` and `duration.txt`.
- `redteam/` — red-team output.
- `run.sh` — reproduction script.
