# COMPARISON: v0.3.13 (iteration 3) vs v0.3.12 (previous accepted)

- **Report branch:** `report-sarcastic-mercenary-chimichanga-ecc5e842-2026-10-02`
- **Date:** 2026-10-02
- **Change under test:** new high-severity `chat_template_boundary` YARA rule for fake chat-template special tokens (`<|im_start|>`, `<|im_end|>`, `<|endoftext|>`, `<|start user prompt|>`, fullwidth and look-alike pipes) and forged role-boundary tags (`</email><user>`, a bare `<user>`/`<human>`/`<assistant>` line).
- **Baseline:** v0.3.12 at `06f4792` (report `report-fourth-wall-chimichanga-ec260996-2026-10-02`).
- **Candidate:** v0.3.13 at `6afbdea`.
- **Decision:** **Accepted.** llmail recall and F1 rose on dev and holdout. deepset and browsesafe are unchanged. No dev or holdout set gained a false positive. The hypothesis (BASELINE.md cross-set FN theme #4) was built from the dev sets only; the holdout sets were run only for this evaluation.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.12 | 1.000 | 0.317 | 0.481 | 19 | 0 | 41 | 3.09 |
| deepset | v0.3.13 | 1.000 | 0.317 | 0.481 | 19 | 0 | 41 | 3.31 |
| llmail | v0.3.12 | 1.000 | 0.560 | 0.718 | 168 | 0 | 132 | 9.59 |
| llmail | v0.3.13 | 1.000 | 0.647 | 0.785 | 194 | 0 | 106 | 9.96 |
| browsesafe | v0.3.12 | 0.700 | 0.590 | 0.640 | 177 | 76 | 123 | 300.29 |
| browsesafe | v0.3.13 | 0.700 | 0.590 | 0.640 | 177 | 76 | 123 | 289.05 |

Mean dev F1 went from 0.613 to 0.636.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.12 | 0.958 | 0.227 | 0.367 | 46 | 2 | 157 | 2.95 |
| deepset-holdout | v0.3.13 | 0.958 | 0.227 | 0.367 | 46 | 2 | 157 | 3.11 |
| llmail-holdout | v0.3.12 | 1.000 | 0.547 | 0.707 | 164 | 0 | 136 | 9.59 |
| llmail-holdout | v0.3.13 | 1.000 | 0.633 | 0.776 | 190 | 0 | 110 | 9.53 |
| browsesafe-holdout | v0.3.12 | 0.664 | 0.540 | 0.596 | 162 | 82 | 138 | 286.75 |
| browsesafe-holdout | v0.3.13 | 0.664 | 0.540 | 0.596 | 162 | 82 | 138 | 283.46 |

Mean holdout F1 went from 0.556 to 0.579. TN is unchanged on every set (deepset 56, llmail 160, browsesafe 224, deepset-holdout 341, llmail-holdout 160, browsesafe-holdout 218).

**Latency:** p95 moved by under 0.4 ms on deepset and llmail, which is within run-to-run variance (run-2 dev p95: deepset 3.34, llmail 9.58, browsesafe 290.50 ms). browsesafe p95 went down by 3–11 ms, which is also within its variance. The three new regexes cost under 1 ms per payload. browsesafe is still over the 100 ms fast-mode budget, as it was before this change.

## Thorough corpus

| Target | v0.3.12 findings | v0.3.13 findings | Δ | v0.3.12 duration_ms | v0.3.13 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 332 | 356 | +24 | 3448 | 3483 |
| uiuc-injecagent | 7988 | 8138 | +150 | 3817 | 3573 |
| lakera-pint-benchmark | 38 | 38 | 0 | 2733 | 2799 |
| alexh-prompt-injection-scanner | 143 | 143 | 0 | 2948 | 3078 |
| duriantaco-skylos | 2306 | 2306 | 0 | 18106 | 20045 |
| promptfoo-scenarios | 11 | 11 | 0 | 2908 | 2859 |
| promptfoo-webagents | 33 | 33 | 0 | 2686 | 2767 |
| **Total** | **10851** | **11025** | **+174** | **36646** | **38604** |

- All 174 new findings come from `chat_template_boundary` (BIPIA +24, InjecAgent +150). Every other YARA rule count, and every non-YARA source count, is identical to v0.3.12 on every target.
- Total duration_ms rose 5.3 %. Almost all of it is skylos (+1939 ms), where the new rule adds no findings; the other six targets moved by -244 to +130 ms. Skylos is dominated by external scanners, so this is most likely run-to-run variance. Every target exited with rc=0 and 0 errors.
- The 174 new corpus hits have not been spot-checked for false positives yet.

## Red-team and checks

- `./scripts/redteam.sh` flagged 11/11 payloads at 11–12 ms each (`redteam/redteam.txt`, run with the v0.3.13 binary built in this worktree). Three of the payloads are new chat-template ones.
- `go build ./...` and `go test ./...` pass, including the new table-driven `TestChatTemplateBoundary` with positive and benign cases (Haskell `<|>`, bra-ket `<|x| , |y|>`, F# `<| y |>`, ordinary HTML, inline `<user>` placeholders).
- Determinism: a second dev eval run matched the first once latency and duration fields were removed, and the FP/FN JSONL files were byte-identical (`cmp`).
- Known bypasses: tokens split across lines, HTML-entity or base64-encoded tokens, role tags with attributes (`<user id=1>`), and an inline role tag that follows a closing tag not on the list.

## Notes

- The eval and corpus reports were generated from the iteration-3 working tree before the version bump, so they record `safeanalyze_version: 0.3.12`. The rule set is identical to the tagged v0.3.13 commit. Corpus reports record `scan_mode: thorough` and eval reports record `scan_mode: fast`.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- No `duration.txt` was written for this run. The corpus durations above are `duration_ms` from each `safeanalyze.json`.
- **Rejected this iteration:** hidden-char FP narrowing (browsesafe F1 0.6401 → 0.6396) and dropping hidden-char-only hits (browsesafe 0.6426 but llmail 0.7179 → 0.6928).
- **Next candidates:** spot-check the new BIPIA/InjecAgent hits for FP inflation; JSON tool-call injection in llmail FNs; a fast-mode input cap or HTML pre-filter for browsesafe latency.
