# COMPARISON: v0.3.35 (iteration 31) vs v0.3.34 (previous accepted)

- **Report branch:** `report-merc-with-a-mouth-taco-5a975cc8-2026-10-14`
- **Date:** 2026-10-14
- **Change under test:** one new pattern in the existing `multilingual_prompt_injection` rule, directly after the existing Spanish pattern, for Spanish "ignore/forget the previous ..." overrides: "(todo) lo dicho / lo anterior / lo previo / lo mencionado (antes/...)" before a sentence end or "y"; singular and extra nouns (instrucción, indicación, direcciones, regla, peticiones, declaraciones, directriz/directrices, pautas, restricciones) with anterior(es)/previa(s)/previo(s)/precedente(s); "cualquier petición/dirección anterior"; and "tus/sus instrucciones/pautas/reglas/... (+1 word)" before "y", `.`, `!` or end of line. No new rule name.
- **Baseline:** v0.3.34 extended-set baseline (`/tmp/safeanalyze-iter/v0.3.34-baseline`; tag `v0.3.34` detection code).
- **Candidate:** iteration 31, measured as the uncommitted working tree on `99ce686` in `/tmp/safeanalyze-iter/iter31` (so its report metadata still reads `safeanalyze_version: 0.3.34`), then released unchanged as v0.3.35 (tag `v0.3.35` at `b8c270a`; the release commit adds only the version bump and docs on top of the measured code).
- **Decision:** **Accepted** under the CLAUDE.md decision rules. PRIMARY dev F1 gain from spanish 0.400 → 0.556 (TP 13 → 20, FP 0); all other PRIMARY dev sets unchanged. No precision drop on any set; every holdout set non-regressing (spanish-holdout F1 0.464 → 0.605). RECALL-ONLY multilingual recall rises on dev and holdout. Corpus findings unchanged.

## Why

PRIMARY spanish recall was 0.25 on dev. In `fn-spanish.jsonl`, 12 of 39 misses were Spanish "ignore the previous instructions" phrasings the existing Spanish pattern did not cover (unlisted nouns, singular "LA INSTRUCCIÓN ANTERIOR", determiners cualquier/tus/sus/el, "todo lo dicho antes", bare "lo anterior"). The same forms appear in the multilingual dev misses. Built on dev FN/FP files only (spanish, multilingual, toolresp, llmail, browsesafe); no holdout, GUARD-ONLY or stress misclassification was inspected. Benign guards: "lo anterior" needs a sentence end or "y" (so "Olvida lo anterior, al final no vamos al cine." stays clean), and generic nouns (mensaje, tarea, pasos) are not listed.

Rejected alternative: dropping the `obfuscated_javascript` `\u[0-9a-fA-F]{4}` pattern would remove 29 of 31 toolresp dev FPs but cut toolresp dev TP 143 → 46 (105 injection samples also contain JSON `\u` escapes; a dataset artifact).

`multilingual_prompt_injection` hits (injection / benign): spanish 13/0 → 20/0, spanish-holdout 16/0 → 23/0, multilingual 11/0 → 17/0.

## Labeled fast-mode eval: PRIMARY dev

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail | v0.3.34 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 6.17 | 9.46 |
| llmail | v0.3.35 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 6.19 | 9.38 |
| browsesafe | v0.3.34 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 18.80 | 49.61 |
| browsesafe | v0.3.35 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 19.87 | 49.73 |
| toolresp | v0.3.34 | 0.822 | 0.477 | 0.603 | 143 | 31 | 269 | 157 | 5.51 | 7.29 |
| toolresp | v0.3.35 | 0.822 | 0.477 | 0.603 | 143 | 31 | 269 | 157 | 5.74 | 7.34 |
| spanish | v0.3.34 | 1.000 | 0.250 | 0.400 | 13 | 0 | 300 | 39 | 5.54 | 6.63 |
| spanish | v0.3.35 | 1.000 | 0.385 | 0.556 | 20 | 0 | 300 | 32 | 5.87 | 6.90 |

## Labeled fast-mode eval: PRIMARY holdout

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail-holdout | v0.3.34 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.25 | 9.65 |
| llmail-holdout | v0.3.35 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.26 | 9.46 |
| browsesafe-holdout | v0.3.34 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 17.80 | 49.33 |
| browsesafe-holdout | v0.3.35 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 17.71 | 48.95 |
| toolresp-holdout | v0.3.34 | 0.795 | 0.453 | 0.577 | 136 | 35 | 265 | 164 | 5.73 | 7.44 |
| toolresp-holdout | v0.3.35 | 0.795 | 0.453 | 0.577 | 136 | 35 | 265 | 164 | 5.91 | 7.74 |
| spanish-holdout | v0.3.34 | 1.000 | 0.302 | 0.464 | 16 | 0 | 371 | 37 | 5.55 | 6.39 |
| spanish-holdout | v0.3.35 | 1.000 | 0.434 | 0.605 | 23 | 0 | 371 | 30 | 5.88 | 7.21 |

## Labeled fast-mode eval: RECALL-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| multilingual | v0.3.34 | 1.000 | 0.018 | 0.035 | 15 | 0 | 0 | 825 | 5.52 | 6.54 |
| multilingual | v0.3.35 | 1.000 | 0.025 | 0.049 | 21 | 0 | 0 | 819 | 5.75 | 6.77 |
| multilingual-holdout | v0.3.34 | 1.000 | 0.018 | 0.036 | 15 | 0 | 0 | 810 | 5.54 | 6.30 |
| multilingual-holdout | v0.3.35 | 1.000 | 0.024 | 0.047 | 20 | 0 | 0 | 805 | 5.77 | 6.71 |

## Labeled fast-mode eval: GUARD-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.34 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 5.31 | 6.32 |
| deepset | v0.3.35 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 5.87 | 7.00 |
| deepset-holdout | v0.3.34 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.65 | 6.40 |
| deepset-holdout | v0.3.35 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 5.89 | 6.84 |
| nemotron | v0.3.34 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 5.67 | 6.85 |
| nemotron | v0.3.35 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 5.95 | 7.24 |
| nemotron-holdout | v0.3.34 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 5.59 | 6.49 |
| nemotron-holdout | v0.3.35 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 5.84 | 6.67 |

## Labeled fast-mode eval: REPORT-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| promptprotection-stress | v0.3.34 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 5.59 | 6.88 |
| promptprotection-stress | v0.3.35 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 5.66 | 6.52 |

## Thorough corpus

| Target | v0.3.34 findings | v0.3.35 findings | Δ | v0.3.34 duration_ms | v0.3.35 duration_ms |
|---|---|---|---|---|---|
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 3052 | 2973 |
| duriantaco-skylos | 2264 | 2264 | +0 | 14691 | 15596 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2804 | 2735 |
| microsoft-bipia | 355 | 355 | +0 | 3268 | 3146 |
| promptfoo-scenarios | 11 | 11 | +0 | 3018 | 3002 |
| promptfoo-webagents | 33 | 33 | +0 | 2880 | 3011 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3486 | 3921 |
| **Total** | **11015** | **11015** | **+0** | **33199** | **34384** |

- Findings are identical per target and per source. Every target's `errors` array is empty, and every corpus dir now carries `duration.txt` (written by `scan` since v0.3.34).
- Total duration_ms 33199 → 34384 (+1185 ms), mostly skylos (+905 ms, external scanners) and InjecAgent (+435 ms); run-to-run variance, not the change (fast-mode labeled latency is flat).

## Latency

Fast-mode p95 moved by less than 1 ms on every set (browsesafe dev 49.61 → 49.73 ms against the v0.3.34 baseline dir; the loop's own pre-run measured 48.84 ms; run 2: 48.70 ms). browsesafe-holdout 49.33 → 48.95 ms. All other sets < 10 ms. All within the 100 ms budget.

## Red-team

`./scripts/redteam.sh` on the v0.3.35 build in this worktree: 12/12 payloads flagged at 18–20 ms (`redteam/redteam.txt`). The iteration-31 run before release is in `redteam/redteam-iter31.txt` (12/12, 19–20 ms).

## Determinism

Two dev runs (iter31 and iter31/run2) produced byte-identical FP/FN JSONL on all 14 compared files and identical TP/FP counts.

## Files

- `eval/eval-*.json` — fast-mode labeled eval for all dev, holdout and stress sets (FP/FN JSONL not included).
- `corpus/<target>/` — thorough-mode `safeanalyze.{json,md,html,sarif}` and `duration.txt`.
- `redteam/` — red-team output.
- `run.sh` — reproduction script.
