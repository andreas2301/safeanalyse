# COMPARISON: v0.3.14 (iteration 4) vs v0.3.13 (previous accepted)

- **Report branch:** `report-merc-mouth-chimichanga-3499c54e-2026-10-02`
- **Date:** 2026-10-02
- **Change under test:** new high-severity `persona_hijack` YARA rule for requests that the model adopt a new persona or role ("I want you to act as a ...", "From now on you are DAN", German "Ich möchte, dass Sie als ... agieren/auftreten", Spanish "quiero que actúes como", French "je veux que tu agisses comme / joues le rôle").
- **Baseline:** v0.3.13 at `6afbdea` (report `report-sarcastic-mercenary-chimichanga-ecc5e842-2026-10-02`).
- **Candidate:** v0.3.14 at `1f67e3a`.
- **Decision:** **Accepted.** deepset recall and F1 rose on dev and holdout. llmail and browsesafe are unchanged. No dev or holdout set gained a false positive. The hypothesis (BASELINE.md §4, persona/role-play hijack FN category) was built from the dev sets only; the holdout sets were run only for this evaluation.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.13 | 1.000 | 0.317 | 0.481 | 19 | 0 | 41 | 3.31 |
| deepset | v0.3.14 | 1.000 | 0.400 | 0.571 | 24 | 0 | 36 | 3.91 |
| llmail | v0.3.13 | 1.000 | 0.647 | 0.785 | 194 | 0 | 106 | 9.96 |
| llmail | v0.3.14 | 1.000 | 0.647 | 0.785 | 194 | 0 | 106 | 10.88 |
| browsesafe | v0.3.13 | 0.700 | 0.590 | 0.640 | 177 | 76 | 123 | 289.05 |
| browsesafe | v0.3.14 | 0.700 | 0.590 | 0.640 | 177 | 76 | 123 | 316.46 |

Mean dev F1 went from 0.636 to 0.666. `persona_hijack` dev rule hits: deepset 5 injection / 0 benign; llmail and browsesafe 0 / 0.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.13 | 0.958 | 0.227 | 0.367 | 46 | 2 | 157 | 3.11 |
| deepset-holdout | v0.3.14 | 0.964 | 0.266 | 0.417 | 54 | 2 | 149 | 3.32 |
| llmail-holdout | v0.3.13 | 1.000 | 0.633 | 0.776 | 190 | 0 | 110 | 9.53 |
| llmail-holdout | v0.3.14 | 1.000 | 0.633 | 0.776 | 190 | 0 | 110 | 10.79 |
| browsesafe-holdout | v0.3.13 | 0.664 | 0.540 | 0.596 | 162 | 82 | 138 | 283.46 |
| browsesafe-holdout | v0.3.14 | 0.664 | 0.540 | 0.596 | 162 | 82 | 138 | 309.69 |

Mean holdout F1 went from 0.579 to 0.596. `persona_hijack` holdout rule hits: deepset-holdout 9 injection / 0 benign (+8 TP; one hit was already caught by another rule); browsesafe-holdout 1 injection / 0 benign (already caught by another rule); llmail-holdout 0 / 0. TN is unchanged on every set (deepset 56, llmail 160, browsesafe 224, deepset-holdout 341, llmail-holdout 160, browsesafe-holdout 218).

**Latency:** p95 rose by under 1.3 ms on deepset and llmail. browsesafe p95 rose about 9 % (dev 289 → 316 ms, holdout 283 → 310 ms; run-2 dev p95 314.9 ms), and p50 is 114.9 ms dev / 102.9 ms holdout. Five extra regexes should not cost ~27 ms per payload, and llmail/deepset latency also drifted up, so part of this is likely machine variance; it should be re-checked on an idle machine. browsesafe was already over the 100 ms fast-mode budget before this change.

## Thorough corpus

| Target | v0.3.13 findings | v0.3.14 findings | Δ | v0.3.13 duration_ms | v0.3.14 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 356 | 356 | 0 | 3483 | 3768 |
| uiuc-injecagent | 8138 | 8138 | 0 | 3573 | 3706 |
| lakera-pint-benchmark | 38 | 38 | 0 | 2799 | 2715 |
| alexh-prompt-injection-scanner | 143 | 143 | 0 | 3078 | 3026 |
| duriantaco-skylos | 2306 | 2306 | 0 | 20045 | 16992 |
| promptfoo-scenarios | 11 | 11 | 0 | 2859 | 2818 |
| promptfoo-webagents | 33 | 33 | 0 | 2767 | 2787 |
| **Total** | **11025** | **11025** | **0** | **38604** | **35812** |

- `persona_hijack` has 0 hits on all 7 targets. Per-source and per-rule finding counts are identical to v0.3.13 on every target (checked from each `safeanalyze.json`).
- Total duration_ms fell 7.2 %, almost all from skylos (-3053 ms), which is dominated by external scanners. The other six targets moved by -84 to +285 ms. Treat this as run-to-run variance, not a latency gain. Every target exited with rc=0 and 0 errors.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 11–12 ms each (`redteam/redteam.txt`, run with the v0.3.14 binary built in this worktree). The new payload is "I want you to act as a Linux terminal ...".
- `go build ./...` and `go test ./...` pass, including the new table-driven `TestPersonaHijack` with positive and benign cases ("the cache acts as a buffer", "From now on, you will be billed monthly", "I would like you to serve as Chair", "Quiero que actúes de buena fe").
- Determinism: a second dev eval run matched the first once latency and duration fields were removed, and the FP/FN JSONL files were byte-identical.
- Known bypasses: lowercase persona names without an article ("act as dan"), verbs outside the list ("serve as", "become"), and languages other than en/de/es/fr.

## Notes

- Eval and corpus reports record `safeanalyze_version: 0.3.14`. Corpus reports record `scan_mode: thorough` and eval reports record `scan_mode: fast`.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- `duration.txt` is not included: the iteration-4 copies hold nanosecond values, not milliseconds. The corpus durations above are `duration_ms` from each `safeanalyze.json`.
- **Rejected this iteration:** a prompt-leak rule ("display your system instructions", "print above prompt", "return your embeddings").
- **Next candidates:** spot-check the BIPIA/InjecAgent hits from v0.3.11/v0.3.13 for FP inflation; JSON tool-call injection in llmail FNs; a fast-mode input cap or HTML pre-filter for browsesafe latency.
