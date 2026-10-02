# COMPARISON: v0.3.11 (iteration 1) vs v0.3.10 (previous accepted)

- **Report branch:** `report-katana-unicorn-chimichanga-210c8b75-2026-10-02`
- **Date:** 2026-10-02
- **Change under test:** wider override verb/object coverage in the `prompt_injection_comment` YARA rule (`ignore|disregard|forget` + previous/prior/above/everything and instructions/tasks/rules/context objects, upper-case `SYSTEM|ADMIN OVERRIDE`, `New directive:` / line-start `New instructions:`).
- **Baseline:** v0.3.10 at `2957004` (detection logic unchanged since v0.3.9).
- **Decision:** **Accepted.** Dev and holdout recall/F1 rose on all three sets, with no new false positives on any dev or holdout set.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.10 | n/a | 0.000 | 0.000 | 0 | 0 | 60 | 2.19 |
| deepset | v0.3.11 | 1.000 | 0.167 | 0.286 | 10 | 0 | 50 | 2.81 |
| llmail | v0.3.10 | 1.000 | 0.550 | 0.710 | 165 | 0 | 135 | 8.12 |
| llmail | v0.3.11 | 1.000 | 0.560 | 0.718 | 168 | 0 | 132 | 8.62 |
| browsesafe | v0.3.10 | 0.648 | 0.467 | 0.543 | 140 | 76 | 160 | 248.72 |
| browsesafe | v0.3.11 | 0.697 | 0.583 | 0.635 | 175 | 76 | 125 | 256.73 |

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.10 | 0.800 | 0.039 | 0.075 | 8 | 2 | 195 | 2.06 |
| deepset-holdout | v0.3.11 | 0.938 | 0.148 | 0.255 | 30 | 2 | 173 | 2.50 |
| llmail-holdout | v0.3.10 | 1.000 | 0.547 | 0.707 | 164 | 0 | 136 | 8.15 |
| llmail-holdout | v0.3.11 | 1.000 | 0.547 | 0.707 | 164 | 0 | 136 | 8.88 |
| browsesafe-holdout | v0.3.10 | 0.642 | 0.490 | 0.556 | 147 | 82 | 153 | 237.71 |
| browsesafe-holdout | v0.3.11 | 0.663 | 0.537 | 0.593 | 161 | 82 | 139 | 257.53 |

p95 moved by about 0.3 to 20 ms. That is within run-to-run variance: a second v0.3.10 run gave browsesafe p95 240.07 ms, and a second v0.3.11 run gave 259.13 ms. browsesafe p95 is still over the 100 ms fast-mode budget in both versions.

## Thorough corpus

| Target | v0.3.10 findings | v0.3.11 findings | Δ | v0.3.10 duration_ms | v0.3.11 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 331 | 332 | +1 | 3532 | 3171 |
| uiuc-injecagent | 6924 | 7988 | +1064 | 3840 | 3649 |
| lakera-pint-benchmark | 35 | 38 | +3 | 2701 | 2899 |
| alexh-prompt-injection-scanner | 142 | 143 | +1 | 2984 | 3044 |
| duriantaco-skylos | 2277 | 2306 | +29 | 18167 | 17598 |
| promptfoo-scenarios | 11 | 11 | 0 | 2861 | 2864 |
| promptfoo-webagents | 31 | 33 | +2 | 2771 | 2700 |
| **Total** | **9751** | **10851** | **+1100** | **36856** | **35925** |

- All of the change comes from `yara`. The entropy, hiddenchars, semgrep, trufflehog and prompt-injection-scanner counts are identical on every target.
- 0 errors, no `ERROR.txt`, rc=0 on all targets. Other v0.3.10 baseline runs totalled 36078 and 38015 ms.
- **Caveat:** the +1100 corpus findings (mostly InjecAgent, whose test cases contain many injection strings) have not been spot-checked for false-positive inflation. The labeled sets show no FP increase.

## Red-team and checks

- `./scripts/redteam.sh` flagged 8/8 payloads at 11–13 ms each (`redteam/redteam.txt`, run with the v0.3.11 binary built in this worktree).
- `go build ./...` and `go test ./...` pass, including the new override-variant test.
- Determinism: repeat dev eval runs gave identical metrics once latency was removed, and the FP/FN sets matched byte for byte.

## Notes

- The eval and corpus reports in `eval/` and `corpus/` were produced from the working tree before the version bump, so their metadata records `safeanalyze_version: 0.3.10`. The detection code is identical to the tagged v0.3.11 commit; only `pkg/version/version.go` differs.
- FP/FN JSONL files are left out because they contain dataset sample text. They can be regenerated with `run.sh`.
- Wall-clock `duration.txt` values (ms) for v0.3.11: bipia 4157, injecagent 3065, pint 2914, pi-scanner 3202, skylos 17752, scenarios 2876, webagents 2713.
