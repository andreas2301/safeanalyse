# safeanalyze v0.3.7 corpus evaluation

**Device:** boom-zany-sarcasm-cd4dee48  
**Date:** 2026-07-15  
**Version:** v0.3.7  
**Mode:** thorough

## Red-team result

`8/8` adversarial payloads flagged by `safeanalyze inspect` (fast mode).

## Corpus findings vs. v0.3.5 baseline

| Target | v0.3.5 findings | v0.3.5 duration (ms) | v0.3.7 findings | v0.3.7 duration (ms) | Δ findings | Δ duration (ms) |
|---|---:|---:|---:|---:|---:|---:|
| docs/promptfoo-scenarios | 11 | 1652 | 11 | 1667 | 0 | +15 |
| docs/promptfoo-webagents | 31 | 1571 | 31 | 1676 | 0 | +105 |
| repos/lakera-pint-benchmark | 35 | 1778 | 35 | 1617 | 0 | −161 |
| repos/alexh-prompt-injection-scanner | 142 | 2495 | 142 | 2383 | 0 | −112 |
| repos/microsoft-bipia | 331 | 4247 | 331 | 3235 | 0 | −1012 |
| repos/uiuc-injecagent | 6924 | 9286 | 6924 | 2372 | 0 | −6914 |
| repos/duriantaco-skylos | 1371 | 24298 | 1371 | 8925 | 0 | −15373 |
| **Total** | **8845** | — | **8845** | — | **0** | **−23352** |

## Interpretation

- Detection coverage is unchanged (8845 total findings).
- Parallel YARA file scanning cuts wall-clock time dramatically on repositories with many files or long lines:
  - `uiuc-injecagent`: ~9.3 s → ~2.4 s
  - `duriantaco-skylos`: ~24.3 s → ~8.9 s
  - `microsoft-bipia`: ~4.2 s → ~3.2 s
- Small targets show normal variance because their scan time is dominated by scanner startup and report writing, not YARA walking.

## Notes

- Semgrep was skipped for the same small targets as in v0.3.5 (file count below 50).
- Report output remains deterministic because `pkg/report` sorts findings before writing.
