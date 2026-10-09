# COMPARISON: v0.3.28 (iteration 24) vs v0.3.27 (previous accepted)

- **Report branch:** `report-chimichanga-wade-unicorn-c4a7fa08-2026-10-09`
- **Date:** 2026-10-09
- **Change under test:** two new patterns in the existing `persona_hijack` rule for role-persistence / stay-in-character instructions in English and German: "stay/remain/keep in character", "always/forever remain in their/your roles/characters/personas", "never/without/don't/must not (even) break (out of) (their/your) character/role/persona", "verharren/bleib* stets/immer/jederzeit in ihren/deiner Rolle(n)/Charakter*/Figur*" and "aus ihren/deinen Rollen/Charakteren auszubrechen/herauszufallen".
- **Baseline:** v0.3.27 (iteration 23 measurement, `/tmp/safeanalyze-iter/iter23`; detection code identical to tag `v0.3.27` at `f1d6849`).
- **Candidate:** iteration 24, measured as the uncommitted working tree on `f1d6849` in `/tmp/safeanalyze-iter/iter24`, then released unchanged as v0.3.28 (tag `v0.3.28` at `8f1654a`; the release commit adds only the version bump and docs on top of the measured code).
- **Decision:** **Accepted** under the CLAUDE.md decision rules. Dev F1 sum +0.0440 (all from deepset 0.652 → 0.696). Holdout: deepset-holdout F1 0.432 → 0.468, the others unchanged. No precision drop and no new FP on any dev or holdout set. Corpus findings unchanged.

## Why

Steering priority 2: persona/role-play, deepset recall. Role-play injections tell the model to keep its assigned persona ("always remain in their roles without even breaking out of their characters", "never break character", "verharren stets in ihren Rollen"). Each form needs a following sentence end, punctuation or a continuation word, so job/role descriptions ("She remains in the role of CEO", "Er bleibt in seiner Rolle als Trainer") and "character limit/encoding/set" text stay unflagged. Built on dev-set evidence only; no holdout misclassification was inspected. `data_exfiltration_email` is not touched and no rule name is added.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.27 | 1.000 | 0.483 | 0.652 | 29 | 0 | 56 | 31 | 4.61 | 5.44 |
| deepset | v0.3.28 | 1.000 | 0.533 | 0.696 | 32 | 0 | 56 | 28 | 5.06 | 6.29 |
| llmail | v0.3.27 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.04 | 7.82 |
| llmail | v0.3.28 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 5.64 | 8.61 |
| browsesafe | v0.3.27 | 0.873 | 0.663 | 0.754 | 199 | 29 | 271 | 101 | 15.73 | 39.09 |
| browsesafe | v0.3.28 | 0.873 | 0.663 | 0.754 | 199 | 29 | 271 | 101 | 17.24 | 40.99 |

`persona_hijack` hits (injection / benign): deepset 5/0 → 8/0 (3 new TPs); none on llmail or browsesafe dev.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.27 | 1.000 | 0.276 | 0.432 | 56 | 0 | 343 | 147 | 4.59 | 5.57 |
| deepset-holdout | v0.3.28 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 4.97 | 5.64 |
| llmail-holdout | v0.3.27 | 1.000 | 0.717 | 0.835 | 215 | 0 | 160 | 85 | 5.17 | 7.98 |
| llmail-holdout | v0.3.28 | 1.000 | 0.717 | 0.835 | 215 | 0 | 160 | 85 | 5.41 | 8.24 |
| browsesafe-holdout | v0.3.27 | 0.861 | 0.597 | 0.705 | 179 | 29 | 271 | 121 | 15.68 | 42.60 |
| browsesafe-holdout | v0.3.28 | 0.861 | 0.597 | 0.705 | 179 | 29 | 271 | 121 | 15.01 | 39.78 |

- `persona_hijack` hits: deepset-holdout 9/0 → 15/0; browsesafe-holdout 1/0 (unchanged).
- The holdout sets were run only for this evaluation.

## Thorough corpus

| Target | v0.3.27 findings | v0.3.28 findings | Δ | v0.3.27 duration_ms | v0.3.28 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 355 | 355 | +0 | 3502 | 3477 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3366 | 3627 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2680 | 2793 |
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 2972 | 2863 |
| duriantaco-skylos | 2264 | 2264 | +0 | 17841 | 14587 |
| promptfoo-scenarios | 11 | 11 | +0 | 2815 | 2767 |
| promptfoo-webagents | 33 | 33 | +0 | 2784 | 2655 |
| **Total** | **11015** | **11015** | **+0** | **35960** | **32769** |

- Findings are identical per target and per source (the new patterns produce no corpus hits). Every target's `errors` array is empty.
- Total duration_ms 35960 → 32769 (-3191 ms, -3254 ms from skylos external scanners), within run-to-run variance.

## Latency

Fast-mode p95 moved between -2.8 and +1.9 ms (deepset 5.44 → 6.29 ms, browsesafe dev 39.09 → 40.99 ms, run 2: 43.64 ms; browsesafe-holdout 42.60 → 39.78 ms). All sets stay well inside the 100 ms budget.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 18–23 ms each, using the binary built in this worktree from master v0.3.28. The output is in `redteam/redteam.txt`.
- `TestPersonaHijack` gains 10 positive and 13 benign cases (en/de). `go build ./...` and `go test ./...` pass on v0.3.28.
- Determinism: run 2 of the dev evals gives the same results as run 1 (latency and metadata removed), and its FP/FN JSONL files are byte-identical to run 1.

## Notes

- **Version metadata:** the measurements ran before the version bump, so eval and corpus reports record `safeanalyze_version: 0.3.27` (`scan_mode: fast` / `thorough`). The detection code is identical to tag `v0.3.28`.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json`.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- **Known gaps:** role persistence without the listed verbs ("don't drop the act", "keep playing", "stay as DAN"), possessives other than their/your (en) and ihr/dein/euer (de), languages other than en/de, and a following word outside the continuation list are not matched.
