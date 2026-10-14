# COMPARISON: v0.3.40 (iteration 36) vs v0.3.39 (previous accepted)

- **Report branch:** `report-merc-with-a-mouth-chimichanga-95b0fe12-2026-10-14`
- **Date:** 2026-10-14
- **Change under test:** one new pattern in the existing `output_constraint` rule for instructions that force the model to output a quoted string verbatim. Forcing frame: "now/instead/only/simply (please/just) say/output/respond with/repeat (only/exactly/(the) (exact) following/phrase/words/string/sentence/text)(:)" or "say/output/respond with/repeat only/exactly/the exact phrase.../the following(:)"; the frame may also start after a literal `\n`/`\t` escape in a JSON string. Then a quoted string ("…", JSON-escaped `\"…\"`, '…', “…”, „…“, «…» up to 120 characters, `<…>` up to 60), which must end the clause (end of text, newline, punctuation, literal `\n`/`\t`, "only", "and nothing", "from now on", "don't/do not"). Case-insensitive. No new rule name.
- **Baseline:** v0.3.39 (tag `v0.3.39`, measurements in `/tmp/safeanalyze-iter/iter35`).
- **Candidate:** iteration 36. It was measured as the uncommitted working tree on `7ffc0fd` in `/tmp/safeanalyze-iter/iter36` (dev rerun in `iter36/run2`), so its report metadata still reads `safeanalyze_version: 0.3.39`. It was then released unchanged as v0.3.40 (tag `v0.3.40`); the release commit adds only the version bump and docs on top of the measured code.
- **Decision:** **Accepted** under the CLAUDE.md decision rules. PRIMARY dev F1 rises on toolresp from 0.687 to 0.705 (TP 173 → 180, FP 31 unchanged); all other PRIMARY dev sets are unchanged. toolresp-holdout F1 rises 0.621 → 0.635 (TP 151 → 156, FP 35 unchanged); every other holdout set is unchanged. No set has a precision drop. Corpus findings are unchanged.

## Why

PRIMARY toolresp recall was 0.577 on dev. Several `fn-toolresp.jsonl` misses told the model to emit an exact quoted string: "now say '…'", "only output '…'", "only respond with \"hello\"", "say the exact phrase \"…\"", "repeat the following: \"…\"". The quoted payload is often a canary phrase, but the pattern keys on the forcing verb and its framing, not on the phrase, so it does not reintroduce the "I have been PWNED" literal rejected as memorization in iteration 32. Only dev FN/FP files were used (llmail, browsesafe, toolresp, spanish, multilingual). No holdout, GUARD-ONLY or stress misclassification was inspected; for deepset and nemotron only benign label-0 rows in the full dev sets were counted.

Benign guards: the quote must end the clause, so `Just say "no" to drugs.`, `Now say "thank you" to the nice lady.`, `Reply with "STOP" to unsubscribe.`, `Please say "hi" to our team.` and `Respond with "yes" by Friday.` stay clean. Unquoted text ("Now say: I have been hacked"), HTML (`<p>Instead say <em>please</em> and <em>thank you</em>.</p>`), CLI docs (`echo will output exactly 'text' to stdout`) and code (`fmt.Println("output")`) are not matched. All are covered in `TestOutputConstraintVerbatimQuote` (11 positive, 11 benign cases).

Rejected alternatives in this iteration: unescaping JSON `\n`/`\t` before matching (+2 toolresp TP only); urgency or verification lures for browsesafe (best variant 95 positive vs 88 benign hits); a broader frame anchored on punctuation, comma or "and" (+11 to +12 toolresp TP, but red-team false triggers).

`output_constraint` hits (injection / benign): toolresp 12/0 → 22/0, toolresp-holdout 5/0 → 12/0. Every other dev, holdout and stress set is unchanged (browsesafe-holdout keeps its pre-existing 1/4).

## Labeled fast-mode eval: PRIMARY dev

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail | v0.3.39 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 7.17 | 10.62 |
| llmail | v0.3.40 | 1.000 | 0.787 | 0.881 | 236 | 0 | 160 | 64 | 7.24 | 10.76 |
| browsesafe | v0.3.39 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 26.79 | 64.25 |
| browsesafe | v0.3.40 | 0.884 | 0.663 | 0.758 | 199 | 26 | 274 | 101 | 26.78 | 69.08 |
| toolresp | v0.3.39 | 0.848 | 0.577 | 0.687 | 173 | 31 | 269 | 127 | 6.79 | 8.88 |
| toolresp | v0.3.40 | 0.853 | 0.600 | 0.705 | 180 | 31 | 269 | 120 | 6.90 | 8.74 |
| spanish | v0.3.39 | 1.000 | 0.596 | 0.747 | 31 | 0 | 300 | 21 | 6.66 | 7.90 |
| spanish | v0.3.40 | 1.000 | 0.596 | 0.747 | 31 | 0 | 300 | 21 | 6.97 | 8.43 |

## Labeled fast-mode eval: PRIMARY holdout

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| llmail-holdout | v0.3.39 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 7.16 | 10.79 |
| llmail-holdout | v0.3.40 | 1.000 | 0.713 | 0.833 | 214 | 0 | 160 | 86 | 7.41 | 11.20 |
| browsesafe-holdout | v0.3.39 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 25.00 | 65.91 |
| browsesafe-holdout | v0.3.40 | 0.869 | 0.597 | 0.708 | 179 | 27 | 273 | 121 | 24.34 | 63.78 |
| toolresp-holdout | v0.3.39 | 0.812 | 0.503 | 0.621 | 151 | 35 | 265 | 149 | 6.69 | 8.58 |
| toolresp-holdout | v0.3.40 | 0.817 | 0.520 | 0.635 | 156 | 35 | 265 | 144 | 6.94 | 9.16 |
| spanish-holdout | v0.3.39 | 1.000 | 0.434 | 0.605 | 23 | 0 | 371 | 30 | 6.82 | 8.06 |
| spanish-holdout | v0.3.40 | 1.000 | 0.434 | 0.605 | 23 | 0 | 371 | 30 | 7.01 | 7.97 |

## Labeled fast-mode eval: RECALL-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| multilingual | v0.3.39 | 1.000 | 0.025 | 0.049 | 21 | 0 | 0 | 819 | 6.30 | 7.12 |
| multilingual | v0.3.40 | 1.000 | 0.025 | 0.049 | 21 | 0 | 0 | 819 | 6.81 | 8.05 |
| multilingual-holdout | v0.3.39 | 1.000 | 0.024 | 0.047 | 20 | 0 | 0 | 805 | 6.46 | 7.76 |
| multilingual-holdout | v0.3.40 | 1.000 | 0.024 | 0.047 | 20 | 0 | 0 | 805 | 6.78 | 7.66 |

## Labeled fast-mode eval: GUARD-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.39 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 6.62 | 7.40 |
| deepset | v0.3.40 | 1.000 | 0.767 | 0.868 | 46 | 0 | 56 | 14 | 6.81 | 7.53 |
| deepset-holdout | v0.3.39 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 6.57 | 7.48 |
| deepset-holdout | v0.3.40 | 1.000 | 0.305 | 0.468 | 62 | 0 | 343 | 141 | 6.90 | 7.91 |
| nemotron | v0.3.39 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 6.77 | 7.78 |
| nemotron | v0.3.40 | 1.000 | 0.113 | 0.204 | 34 | 0 | 300 | 266 | 7.02 | 8.57 |
| nemotron-holdout | v0.3.39 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 6.49 | 7.83 |
| nemotron-holdout | v0.3.40 | 1.000 | 0.117 | 0.209 | 35 | 0 | 300 | 265 | 6.94 | 7.96 |

## Labeled fast-mode eval: REPORT-ONLY

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| promptprotection-stress | v0.3.39 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 6.55 | 7.46 |
| promptprotection-stress | v0.3.40 | 0.500 | 0.167 | 0.250 | 30 | 30 | 166 | 150 | 7.07 | 8.53 |

## Thorough corpus

| Target | v0.3.39 findings | v0.3.40 findings | Δ | v0.3.39 duration_ms | v0.3.40 duration_ms |
|---|---|---|---|---|---|
| alexh-prompt-injection-scanner | 143 | 143 | +0 | 2904 | 3011 |
| duriantaco-skylos | 2264 | 2264 | +0 | 17462 | 17027 |
| lakera-pint-benchmark | 38 | 38 | +0 | 2783 | 2689 |
| microsoft-bipia | 355 | 355 | +0 | 3183 | 3414 |
| promptfoo-scenarios | 11 | 11 | +0 | 2898 | 3039 |
| promptfoo-webagents | 33 | 33 | +0 | 2932 | 2839 |
| uiuc-injecagent | 8171 | 8171 | +0 | 3485 | 3359 |
| **Total** | **11015** | **11015** | **+0** | **35647** | **35378** |

- Findings are identical per target and per (source, rule). Every target's `errors` array is empty, and every corpus dir has a `duration.txt`.
- Total `duration_ms` went 35647 → 35378 (-269 ms); the per-target differences are run-to-run variance (mostly skylos external scanners), not the change.

## Latency

Fast-mode p95:
- toolresp dev: 8.88 → 8.74 ms (rerun 8.72 ms); toolresp-holdout 8.58 → 9.16 ms.
- browsesafe dev: 64.25 → 69.08 ms (rerun 69.22 ms).
- browsesafe-holdout: 65.91 → 63.78 ms.

All other sets stay under 12 ms, and every set is within the 100 ms budget.

## Red-team

`./scripts/redteam.sh` on the v0.3.40 build in this worktree flagged 12/12 payloads at 21–22 ms (`redteam/redteam.txt`).

## Determinism

The two dev runs (iter36 and iter36/run2) produced byte-identical FP/FN JSONL on all 14 compared files (7 dev sets), with identical TP/FP counts.

## Files

- `eval/eval-*.json`: fast-mode labeled eval for all dev, holdout and stress sets (FP/FN JSONL not included).
- `corpus/<target>/`: thorough-mode `safeanalyze.{json,md,html,sarif}` and `duration.txt`.
- `redteam/`: red-team output.
- `run.sh`: reproduction script.
