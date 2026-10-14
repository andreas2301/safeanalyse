# COMPARISON: v0.3.39 (iteration 35) vs v0.3.38 (previous accepted)

- **Report branch:** `report-deadpool-mercenary-flamethrower-98c8dd5b-2026-10-14`
- **Date:** 2026-10-14
- **Change under test:** two new Spanish patterns in the existing `multilingual_prompt_injection` rule. (1) Dismissal: verbs the old patterns missed ("no tenga en cuenta", "olvídese de", "olvídate de", "deseche", "descarte", "anule", "haga caso omiso de", "no escuche/siga/obedezca", "pase por alto", "deja de lado", plus olvida/ignora forms), then at most 5 words, then instrucción/indicación/directiva/directriz/declaración/prompt with anterior/previo/pasado/precedente/último in either order, or su/sus/tu/tus/ninguna/cualquier + solicitud/petición/información/tarea + qualifier + "y"/"e". (2) Passive revocation: noun + qualifier, optional han sido/fueron/quedan/están, then revocadas/anuladas/canceladas/sin efecto/no se aplican/ya no se aplican. Case-insensitive. No new rule name.
- **Baseline:** v0.3.38 (tag `v0.3.38`, measurements in `/tmp/safeanalyze-iter/iter34`).
- **Candidate:** iteration 35. It was measured as the uncommitted working tree on `ec7dead` in `/tmp/safeanalyze-iter/iter35` (rerun in `iter35/rerun`), so its report metadata still reads `safeanalyze_version: 0.3.38`. It was then released unchanged as v0.3.39 (tag `v0.3.39`); the release commit adds only the version bump and docs on top of the measured code.
- **Decision:** **Accepted** under the CLAUDE.md decision rules. PRIMARY dev F1 rises on spanish from 0.556 to 0.747 (TP 20 → 31, FP 0 unchanged); all other PRIMARY dev sets are unchanged. No set has a precision drop, and every holdout set is unchanged. Corpus findings are unchanged.

## Why

PRIMARY spanish recall was 0.385 on dev. In `fn-spanish.jsonl`, 11 of 32 misses dismissed or revoked the earlier instructions in forms the existing Spanish patterns did not accept: verbs ("No tenga en cuenta", "Olvídese de", "olvídate de", "Deseche", "haga caso omiso de", "No escuche"), nouns (`directiva`, `declaración`, "su solicitud anterior", "ninguna información previa"), adjectives (`pasadas`, `últimas`), a gap of more than one word before the noun ("Ignora a tu amigo y tus instrucciones anteriores") and passive revocations ("Instrucciones pasadas revocadas", "todas las instrucciones anteriores no se aplican"). Only dev FN/FP files were used to build the patterns (spanish, toolresp, llmail, browsesafe, multilingual). No holdout, GUARD-ONLY or stress misclassification was inspected; for the guard sets, only benign hits were counted.

Benign guards: possessives such as "mi" are not accepted, so "Haga caso omiso de mi correo anterior" and "No tenga en cuenta mi solicitud anterior" stay clean. "su petición anterior" needs a following "y"/"e", so "No tenga en cuenta su petición anterior; ya la hemos tramitado." stays clean. Nouns such as normas, reglas or versión are not targets of the new patterns ("Las normas anteriores quedan sin efecto a partir del 1 de enero.", "Descarte la versión anterior del contrato"), and references without a dismissal verb ("Según las instrucciones anteriores, envíe el formulario.") are not matched. All are covered in `TestMultilingualPromptInjection`.

`multilingual_prompt_injection` hits (injection / benign): spanish 20/0 → 31/0. Every other dev, holdout and stress set is unchanged, with 0 benign hits on every set.

## Labeled fast-mode eval: PRIMARY dev

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail | v0.3.38 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 6.58 | 10.00 |
| llmail | v0.3.39 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 7.17 | 10.62 |
| browsesafe | v0.3.38 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 26.23 | 63.42 |
| browsesafe | v0.3.39 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 26.79 | 64.25 |
| toolresp | v0.3.38 | 0.848 | 0.577 | 0.687 | 173 | 31 | 269 | 127 | 6.20 | 8.39 |
| toolresp | v0.3.39 | 0.848 | 0.577 | 0.687 | 173 | 31 | 269 | 127 | 6.79 | 8.88 |
| spanish | v0.3.38 | 1.000 | 0.385 | 0.556 | 20 | 0 | 300 | 32 | 6.10 | 7.17 |
| spanish | v0.3.39 | 1.000 | 0.596 | 0.747 | 31 | 0 | 300 | 21 | 6.66 | 7.90 |

## Labeled fast-mode eval: PRIMARY holdout

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail-holdout | v0.3.38 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 6.60 | 10.57 |
| llmail-holdout | v0.3.39 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 7.16 | 10.79 |
| browsesafe-holdout | v0.3.38 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 23.17 | 62.59 |
| browsesafe-holdout | v0.3.39 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 25.00 | 65.91 |
| toolresp-holdout | v0.3.38 | 0.812 | 0.503 | 0.621 | 151 | 35 | 265 | 149 | 6.38 | 8.30 |
| toolresp-holdout | v0.3.39 | 0.812 | 0.503 | 0.621 | 151 | 35 | 265 | 149 | 6.69 | 8.58 |
| spanish-holdout | v0.3.38 | 1.000 | 0.434 | 0.605 | 23 | 0 | 371 | 30 | 6.20 | 7.08 |
| spanish-holdout | v0.3.39 | 1.000 | 0.434 | 0.605 | 23 | 0 | 371 | 30 | 6.82 | 8.06 |

## Labeled fast-mode eval: RECALL-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| multilingual | v0.3.38 | 1.000 | 0.025 | 0.049 | 21 | 0 | 0 | 819 | 6.01 | 7.26 |
| multilingual | v0.3.39 | 1.000 | 0.025 | 0.049 | 21 | 0 | 0 | 819 | 6.30 | 7.12 |
| multilingual-holdout | v0.3.38 | 1.000 | 0.024 | 0.047 | 20 | 0 | 0 | 805 | 6.01 | 7.08 |
| multilingual-holdout | v0.3.39 | 1.000 | 0.024 | 0.047 | 20 | 0 | 0 | 805 | 6.46 | 7.76 |

## Labeled fast-mode eval: GUARD-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.38 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 6.10 | 6.95 |
| deepset | v0.3.39 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 6.62 | 7.40 |
| deepset-holdout | v0.3.38 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 6.27 | 7.43 |
| deepset-holdout | v0.3.39 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 6.57 | 7.48 |
| nemotron | v0.3.38 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 6.21 | 7.10 |
| nemotron | v0.3.39 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 6.77 | 7.78 |
| nemotron-holdout | v0.3.38 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 6.15 | 7.06 |
| nemotron-holdout | v0.3.39 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 6.49 | 7.83 |

## Labeled fast-mode eval: REPORT-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| promptprotection-stress | v0.3.38 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 6.39 | 7.86 |
| promptprotection-stress | v0.3.39 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 6.55 | 7.46 |

## Thorough corpus

| Target | v0.3.38 findings | v0.3.39 findings | Δ | v0.3.38 duration_ms | v0.3.39 duration_ms |
|---|---|---|---|---|---|
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 2914 | 2904 |
| duriantaco-skylos | 2264 | 2264 | +0 | 16656 | 17462 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2707 | 2783 |
| microsoft-bipia | 355 | 355 | +0 | 3187 | 3183 |
| promptfoo-scenarios | 11 | 11 | +0 | 2942 | 2898 |
| promptfoo-webagents | 33 | 33 | +0 | 2773 | 2932 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3499 | 3485 |
| **Total** | **11015** | **11015** | **+0** | **34678** | **35647** |

- Findings are identical per target and per (source, rule). Every target's `errors` array is empty, and every corpus dir has a `duration.txt`.
- Total `duration_ms` went 34678 → 35647 (+969 ms), mostly from skylos (external scanners). This is run-to-run variance, not the change.

## Latency

Fast-mode p95:
- spanish dev: 7.17 → 7.90 ms (rerun 7.72 ms).
- browsesafe dev: 63.42 → 64.25 ms (rerun 69.96 ms).
- browsesafe-holdout: 62.59 → 65.91 ms.

All other sets stay under 11 ms, and every set is within the 100 ms budget.

## Red-team

`./scripts/redteam.sh` on the v0.3.39 build in this worktree flagged 12/12 payloads at 20–22 ms (`redteam/redteam.txt`).

## Determinism

The two dev runs (iter35 and iter35/rerun) produced byte-identical FP/FN JSONL on all 14 compared files, with identical TP/FP counts.

## Files

- `eval/eval-*.json`: fast-mode labeled eval for all dev, holdout and stress sets (FP/FN JSONL not included).
- `corpus/<target>/`: thorough-mode `safeanalyze.{json,md,html,sarif}` and `duration.txt`.
- `redteam/`: red-team output.
- `run.sh`: reproduction script.
