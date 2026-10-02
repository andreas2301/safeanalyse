# COMPARISON: v0.3.16 (iteration 8) vs v0.3.15 (previous accepted)

- **Report branch:** `report-regenerating-mercenary-chimichanga-38ddfe85-2026-10-02`
- **Date:** 2026-10-02
- **Change under test:** `yara.Engine.ScanFile` gets an exact required-literal prefilter. When a rule is added, each pattern is parsed with `regexp/syntax` into an AND of OR-clauses of literals of at least 3 bytes. The literals are case-folded with the same `unicode.SimpleFold` equivalence that `(?i)` uses, and the result is memoized per pattern string. A pattern runs only when its clauses pass on the folded document and then on the folded line. A pattern with no derivable literals always runs.
- **Baseline:** v0.3.15 at `487c17f` (report `report-healing-factor-taco-c6a64e86-2026-10-02`), re-measured on master `d75bd72` in the same session.
- **Candidate:** v0.3.16 at `5940bda`.
- **Decision:** **Accepted (latency-only).** Every labeled metric (TP/FP/TN/FN, P/R/F1 and per-rule hits) is identical to v0.3.15 on all dev and holdout sets. Corpus findings are identical per target and per source. Fast-mode p95 dropped on every set, and browsesafe is now under the 100 ms budget on both dev and holdout. This matches the decision rule "latency decreases with unchanged labeled metrics and corpus findings".

## Why

A CPU profile of `inspectPayload` over browsesafe dev (600 docs) showed that regex execution was 95 % of CPU: `tryBacktrack` was 68 % and `unicode.SimpleFold` 12 %. Regex compile was only 0.6 %, so caching the engine would not help. The cost was spread over all 107 patterns, about 398k lines × 107 separate `(?i)` backtracking calls. A document-level-only prefilter reached 129 ms p95. Adding the per-line check is what brings p95 under budget. In the prototype, old and new `ScanFile` were run side by side, and `reflect.DeepEqual` on the `[]Match` output of every dev sample found 0 differences on deepset, llmail and browsesafe.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.15 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 2.84 | 3.85 |
| deepset | v0.3.16 | 1.000 | 0.400 | 0.571 | 24 | 0 | 56 | 36 | 2.59 | 3.07 |
| llmail | v0.3.15 | 1.000 | 0.647 | 0.785 | 194 | 0 | 160 | 106 | 3.82 | 11.65 |
| llmail | v0.3.16 | 1.000 | 0.647 | 0.785 | 194 | 0 | 160 | 106 | 2.77 | 3.88 |
| browsesafe | v0.3.15 | 0.719 | 0.647 | 0.681 | 194 | 76 | 224 | 106 | 124.27 | 348.97 |
| browsesafe | v0.3.16 | 0.719 | 0.647 | 0.681 | 194 | 76 | 224 | 106 | 10.48 | 23.94 |

Run-2 dev p95 (v0.3.16): deepset 3.12 ms, llmail 4.02 ms, browsesafe 25.89 ms. On v0.3.15 the values were 4.03 ms, 12.03 ms and 336.93 ms.

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | TN | FN | p50 ms | p95 ms |
|---|---|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.15 | 0.964 | 0.266 | 0.417 | 54 | 2 | 341 | 149 | 2.83 | 3.80 |
| deepset-holdout | v0.3.16 | 0.964 | 0.266 | 0.417 | 54 | 2 | 341 | 149 | 2.64 | 3.43 |
| llmail-holdout | v0.3.15 | 1.000 | 0.633 | 0.776 | 190 | 0 | 160 | 110 | 3.96 | 12.12 |
| llmail-holdout | v0.3.16 | 1.000 | 0.633 | 0.776 | 190 | 0 | 160 | 110 | 2.77 | 3.87 |
| browsesafe-holdout | v0.3.15 | 0.683 | 0.590 | 0.633 | 177 | 82 | 218 | 123 | 114.11 | 340.59 |
| browsesafe-holdout | v0.3.16 | 0.683 | 0.590 | 0.633 | 177 | 82 | 218 | 123 | 10.19 | 27.31 |

The holdout sets were run only for this evaluation, and no holdout misclassification was inspected.

## Thorough corpus

| Target | v0.3.15 findings | v0.3.16 findings | Δ | v0.3.15 duration_ms | v0.3.16 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 356 | 356 | 0 | 3990 | 3321 |
| uiuc-injecagent | 8138 | 8138 | 0 | 3866 | 3376 |
| lakera-pint-benchmark | 38 | 38 | 0 | 2704 | 2864 |
| alexh-prompt-injection-scanner | 143 | 143 | 0 | 2905 | 3028 |
| duriantaco-skylos | 2306 | 2306 | 0 | 16835 | 15980 |
| promptfoo-scenarios | 11 | 11 | 0 | 2758 | 2827 |
| promptfoo-webagents | 33 | 33 | 0 | 2825 | 2851 |
| **Total** | **11025** | **11025** | **0** | **35883** | **34247** |

- Per-source counts are identical on every target. For example, bipia is entropy 57 / hiddenchars 152 / semgrep 1 / yara 146, and skylos is entropy 780 / semgrep 55 / trufflehog 53 / yara 1418.
- Total duration_ms fell 4.6 %. The YARA-heavy targets bipia (−669 ms) and injecagent (−490 ms) improved. Skylos (−855 ms) is dominated by external scanners, so most of its change is variance. The other targets moved by +26 to +160 ms. Every scan exited rc=0 with `errors: []`.
- The v0.3.15 column was re-measured in this session on master `d75bd72` (35883 ms total). The value recorded in the v0.3.15 report was 38610 ms.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 14–18 ms each. The output is in `redteam/redteam.txt`, from a run with the v0.3.16 binary built in this worktree. Per-payload time is dominated by process start-up.
- `go build ./...` and `go test -count=1 ./...` pass. That includes the new `TestPrefilterEquivalence`, which compares the prefiltered scan with an unfiltered scan on injection, benign, multilingual, chat-template, fold (`ſ`, Kelvin `K`), invalid-UTF-8, CRLF and multi-line inputs.
- Determinism: a second dev eval run matched the first once the latency and metadata fields were removed. The FP/FN JSONL files were byte-identical.

## Notes

- **Version metadata:** the eval and corpus reports in `eval/` and `corpus/` came from the iteration-8 working tree before the version bump. They record `safeanalyze_version: 0.3.15`, with `scan_mode: fast` for eval and `thorough` for corpus. The code under test is the same as the `v0.3.16` tag except for the version constant. `redteam/redteam.txt` comes from the tagged v0.3.16 binary. To get reports stamped 0.3.16, run `run.sh`.
- FP/FN JSONL files and scan logs are left out because they contain dataset sample text. `run.sh` regenerates them.
- `duration.txt` holds each target's `duration_ms` from its `safeanalyze.json`.
- **Rejected this iteration:**
  - a document-level-only prefilter (browsesafe p95 129 ms, still over budget);
  - engine caching or precompiling (compile is 0.6 % of CPU).
- **Risk:** the prefilter is exact only while `requiredLiterals` stays sound, meaning every clause is a necessary condition for a match. New regex constructs must either map to a sound clause or return nil (no constraint).
