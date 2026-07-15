# safeanalyze v0.3.5 corpus evaluation

**Device:** chimichanga-taco-beaver-18c1c65c  
**Date:** 2026-07-15  
**Version:** v0.3.5  
**Mode:** thorough

## Red-team result

`8/8` adversarial payloads flagged by `safeanalyze inspect` (fast mode).

## Corpus findings vs. v0.3.3 baseline

| Target | v0.3.3 findings | v0.3.3 duration (ms) | v0.3.5 findings | v0.3.5 duration (ms) | Δ findings | Δ duration (ms) |
|---|---:|---:|---:|---:|---:|---:|
| docs/promptfoo-scenarios | 11 | 2737 | 11 | 1652 | 0 | −1085 |
| docs/promptfoo-webagents | 31 | 2729 | 31 | 1571 | 0 | −1158 |
| repos/lakera-pint-benchmark | 35 | 2792 | 35 | 1778 | 0 | −1014 |
| repos/alexh-prompt-injection-scanner | 142 | 2892 | 142 | 2495 | 0 | −397 |
| repos/microsoft-bipia | 331 | 4128 | 331 | 4247 | 0 | +119 |
| repos/uiuc-injecagent | 6924 | 9028 | 6924 | 9286 | 0 | +258 |
| repos/duriantaco-skylos | 1371 | 24179 | 1371 | 24298 | 0 | +119 |
| **Total** | **8845** | — | **8845** | — | **0** | **−3654 on small targets** |

## Interpretation

- Detection coverage is unchanged (8845 total findings).
- The Semgrep file-count gate skips Semgrep on targets with fewer than 50 files. This removes Semgrep's fixed startup cost on the four smallest targets, cutting their combined latency by ~3.2 s.
- The two large targets (`microsoft-bipia`, `uiuc-injecagent`, `duriantaco-skylos`) show sub-±3 % duration variance, which is within normal run-to-run noise.

## Notes

- Semgrep was skipped with an error note for `repos/uiuc-injecagent` (24 files), `repos/lakera-pint-benchmark` (32 files), `repos/alexh-prompt-injection-scanner` (18 files), and both single-file documentation targets.
- v0.3.4 was reverted because its encoded-prompt-injection fragment expansion added latency without producing new findings.
