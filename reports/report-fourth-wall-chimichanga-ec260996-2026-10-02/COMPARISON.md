# COMPARISON: v0.3.12 (iteration 2) vs v0.3.11 (previous accepted)

- **Report branch:** `report-fourth-wall-chimichanga-ec260996-2026-10-02`
- **Date:** 2026-10-02
- **Change under test:** new critical `multilingual_prompt_injection` YARA rule for German, Spanish, French, Portuguese and Italian forms of "ignore/forget all previous instructions". Bare `alles`/`todo`/`tout` objects only match at a clause end or with a conversational follow-up.
- **Baseline:** v0.3.11 at `6663e60` (report `report-katana-unicorn-chimichanga-210c8b75-2026-10-02`).
- **Candidate:** v0.3.12 at `06f4792`.
- **Decision:** **Accepted.** Recall and F1 rose on deepset (dev and holdout) and browsesafe (dev and holdout). llmail is unchanged. No dev or holdout set gained a false positive. The hypothesis was built from the dev sets only; the holdout sets were run only for this evaluation.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.11 | 1.000 | 0.167 | 0.286 | 10 | 0 | 50 | 2.81 |
| deepset | v0.3.12 | 1.000 | 0.317 | 0.481 | 19 | 0 | 41 | 3.09 |
| llmail | v0.3.11 | 1.000 | 0.560 | 0.718 | 168 | 0 | 132 | 8.62 |
| llmail | v0.3.12 | 1.000 | 0.560 | 0.718 | 168 | 0 | 132 | 9.59 |
| browsesafe | v0.3.11 | 0.697 | 0.583 | 0.635 | 175 | 76 | 125 | 256.73 |
| browsesafe | v0.3.12 | 0.700 | 0.590 | 0.640 | 177 | 76 | 123 | 300.29 |

Mean dev F1 went from 0.546 to 0.613.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.11 | 0.938 | 0.148 | 0.255 | 30 | 2 | 173 | 2.50 |
| deepset-holdout | v0.3.12 | 0.958 | 0.227 | 0.367 | 46 | 2 | 157 | 2.95 |
| llmail-holdout | v0.3.11 | 1.000 | 0.547 | 0.707 | 164 | 0 | 136 | 8.88 |
| llmail-holdout | v0.3.12 | 1.000 | 0.547 | 0.707 | 164 | 0 | 136 | 9.59 |
| browsesafe-holdout | v0.3.11 | 0.663 | 0.537 | 0.593 | 161 | 82 | 139 | 257.53 |
| browsesafe-holdout | v0.3.12 | 0.664 | 0.540 | 0.596 | 162 | 82 | 138 | 286.75 |

Mean holdout F1 went from 0.518 to 0.556.

**Latency:** browsesafe p95 rose from about 257 ms to 287–300 ms (the run-2 dev p95 was 294.95 ms). That is more than the run-to-run variance seen in v0.3.11, which was about 3 ms, so the five extra regexes probably cost about 30–40 ms per large HTML payload. browsesafe was already over the 100 ms fast-mode budget before this change. deepset and llmail p95 rose by under 1 ms and stay under 10 ms. The change was accepted because it gained coverage, which the decision rules allow even when latency goes up.

## Thorough corpus

| Target | v0.3.11 findings | v0.3.12 findings | Δ | v0.3.11 duration_ms | v0.3.12 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 332 | 332 | 0 | 3171 | 3448 |
| uiuc-injecagent | 7988 | 7988 | 0 | 3649 | 3817 |
| lakera-pint-benchmark | 38 | 38 | 0 | 2899 | 2733 |
| alexh-prompt-injection-scanner | 143 | 143 | 0 | 3044 | 2948 |
| duriantaco-skylos | 2306 | 2306 | 0 | 17598 | 18106 |
| promptfoo-scenarios | 11 | 11 | 0 | 2864 | 2908 |
| promptfoo-webagents | 33 | 33 | 0 | 2700 | 2686 |
| **Total** | **10851** | **10851** | **0** | **35925** | **36646** |

- The per-source counts are identical on every target. The new rule fires nowhere in the corpus, which is all English.
- Total duration_ms rose 2.0 %, which is within the observed baseline spread of 35925 to 38015 ms. Every target exited with rc=0, 0 errors and no `ERROR.txt`.

## Red-team and checks

- `./scripts/redteam.sh` flagged 8/8 payloads at 11 ms each (`redteam/redteam.txt`, run with the v0.3.12 binary built in this worktree).
- `go build ./...` and `go test ./...` pass, including the new table-driven multilingual test with positive and benign cases.
- Determinism: a second dev eval run matched the first once latency and duration fields were removed, and the FP/FN JSONL files were byte-identical.
- Known bypasses: homoglyphs, fullwidth letters, letter spacing and base64 (the English rule has the same gaps); periphrastic forms ("Haz caso omiso de", "Ne tenez pas compte de"); and languages outside de/es/fr/pt/it.

## Notes

- All eval and corpus reports here record `safeanalyze_version: 0.3.12`. Corpus reports record `scan_mode: thorough` and eval reports record `scan_mode: fast`.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- No `duration.txt` was written for this run. The corpus durations above are `duration_ms` from each `safeanalyze.json`.
- **Next candidate:** chat-template and role-boundary tokens. The simulation gave llmail +23 TP with 0 benign hits.
