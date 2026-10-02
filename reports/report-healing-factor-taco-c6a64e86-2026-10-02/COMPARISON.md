# COMPARISON: v0.3.15 (iteration 5) vs v0.3.14 (previous accepted)

- **Report branch:** `report-healing-factor-taco-c6a64e86-2026-10-02`
- **Date:** 2026-10-02
- **Change under test:** the `data_exfiltration_email` YARA rule gains two patterns. They match a forward / transmit / relay verb (including -ed / -ing forms) aimed at an email address, but only when a sensitive-data noun such as credentials, conversation history, records, logs or session appears nearby. The patterns cover active phrasing ("forward your current conversation history ... to records@...") and passive phrasing ("responses are transmitted to our verification partner at compliance@...").
- **Baseline:** v0.3.14 at `1f67e3a` (report `report-merc-mouth-chimichanga-3499c54e-2026-10-02`).
- **Candidate:** v0.3.15 at `487c17f`.
- **Decision:** **Accepted.** browsesafe recall and F1 went up on both dev and holdout. deepset and llmail did not change. No dev or holdout set gained a false positive. The hypothesis comes from BASELINE.md cross-set FN theme #5 (exfiltration phrasing) and was built from the dev sets only. The holdout sets were run only for this evaluation.

## Labeled fast-mode eval: dev sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset | v0.3.14 | 1.000 | 0.400 | 0.571 | 24 | 0 | 36 | 3.91 |
| deepset | v0.3.15 | 1.000 | 0.400 | 0.571 | 24 | 0 | 36 | 4.22 |
| llmail | v0.3.14 | 1.000 | 0.647 | 0.785 | 194 | 0 | 106 | 10.88 |
| llmail | v0.3.15 | 1.000 | 0.647 | 0.785 | 194 | 0 | 106 | 11.78 |
| browsesafe | v0.3.14 | 0.700 | 0.590 | 0.640 | 177 | 76 | 123 | 316.46 |
| browsesafe | v0.3.15 | 0.719 | 0.647 | 0.681 | 194 | 76 | 106 | 345.49 |

Mean dev F1 rose from 0.666 to 0.679. On browsesafe dev, `data_exfiltration_email` hits on injection samples rose from 89 to 130, and benign hits stayed at 24. Most of the new injection hits were already caught by another rule, so the net gain is +17 TP. llmail rule hits did not change (134 injection / 0 benign).

## Labeled fast-mode eval: holdout sets

| Set | Version | P | R | F1 | TP | FP | FN | p95 ms |
|---|---|---|---|---|---|---|---|---|
| deepset-holdout | v0.3.14 | 0.964 | 0.266 | 0.417 | 54 | 2 | 149 | 3.32 |
| deepset-holdout | v0.3.15 | 0.964 | 0.266 | 0.417 | 54 | 2 | 149 | 3.84 |
| llmail-holdout | v0.3.14 | 1.000 | 0.633 | 0.776 | 190 | 0 | 110 | 10.79 |
| llmail-holdout | v0.3.15 | 1.000 | 0.633 | 0.776 | 190 | 0 | 110 | 12.56 |
| browsesafe-holdout | v0.3.14 | 0.664 | 0.540 | 0.596 | 162 | 82 | 138 | 309.69 |
| browsesafe-holdout | v0.3.15 | 0.683 | 0.590 | 0.633 | 177 | 82 | 123 | 337.15 |

Mean holdout F1 rose from 0.596 to 0.609. On browsesafe-holdout, `data_exfiltration_email` hits rose from 91 to 122 on injection samples and stayed at 25 on benign samples (+15 TP). TN did not change on any set: deepset 56, llmail 160, browsesafe 224, deepset-holdout 341, llmail-holdout 160, browsesafe-holdout 218.

**Latency:** browsesafe p95 went up about 9 %, from 316 to 345 ms on dev and from 310 to 337 ms on holdout. The run-2 dev p95 was 343.9 ms. browsesafe p50 is 124.3 ms on dev and 113.4 ms on holdout. deepset and llmail p95 stayed under 13 ms. The two new patterns use bounded `.{0,80}` windows, and browsesafe samples are long HTML pages, so part of this cost is real. Over iterations 3 to 5, browsesafe p95 has climbed from 289 to 345 ms. It was already over the 100 ms fast-mode budget, so the next iteration should look at a fast-mode input cap or an HTML pre-filter.

## Thorough corpus

| Target | v0.3.14 findings | v0.3.15 findings | Δ | v0.3.14 duration_ms | v0.3.15 duration_ms |
|---|---|---|---|---|---|
| microsoft-bipia | 356 | 356 | 0 | 3768 | 3795 |
| uiuc-injecagent | 8138 | 8138 | 0 | 3706 | 3793 |
| lakera-pint-benchmark | 38 | 38 | 0 | 2715 | 2729 |
| alexh-prompt-injection-scanner | 143 | 143 | 0 | 3026 | 3027 |
| duriantaco-skylos | 2306 | 2306 | 0 | 16992 | 19662 |
| promptfoo-scenarios | 11 | 11 | 0 | 2818 | 2881 |
| promptfoo-webagents | 33 | 33 | 0 | 2787 | 2723 |
| **Total** | **11025** | **11025** | **0** | **35812** | **38610** |

- Every target has the same finding counts per source as v0.3.14. `data_exfiltration_email` also has the same counts: bipia 4, injecagent 3679, skylos 48, all others 0. Every count was checked against each `safeanalyze.json`.
- Total duration_ms rose 7.8 %. Skylos accounts for +2670 ms of that, and skylos time is dominated by external scanners (Semgrep, TruffleHog). The v0.3.13 skylos run took 20045 ms. The other six targets moved by -64 to +87 ms. This is run-to-run variance, not a regression caused by the rule. Every target exited with rc=0 and 0 errors.

## Red-team and checks

- `./scripts/redteam.sh` flagged 12/12 payloads at 12–13 ms each. The output is in `redteam/redteam.txt`, from a run with the v0.3.15 binary built in this worktree.
- `go build ./...` and `go test -count=1 ./...` pass. That includes the new table-driven `TestDataExfiltrationEmailForward`. Its benign cases include "please forward it to postmaster@...", "messages ... are forwarded to support@...", "SMTP relay to mail@...", "forward your resume to jobs@..." and "Forwarded message from a@x to b@y".
- Determinism: a second dev eval run matched the first once the metadata and latency fields were removed. The FP/FN JSONL files were byte-identical between the two runs.
- Known bypasses: `fwd`, spaced-out letters, obfuscated addresses (`(at)`, `[at]`, fullwidth `＠`, `&#64;`, `%40`), Cyrillic homoglyph verbs, "towards", other verbs (pass along, share with, deliver, route, CC), non-English verbs, and noun lists longer than 80 characters before "to".
- Known FP risk: business mail that forwards a sensitive noun triggers the rule, for example "expense reports should be forwarded to finance@..." or "forward the build logs to devops@...".

## Notes

- Eval and corpus reports record `safeanalyze_version: 0.3.15`. Corpus reports record `scan_mode: thorough`, and eval reports record `scan_mode: fast`.
- FP/FN JSONL files and scan logs are not included because they contain dataset sample text. `run.sh` regenerates them.
- Each `duration.txt` holds the corpus target's `duration_ms` from its `safeanalyze.json`, in milliseconds. The first wall-clock capture gave bad values, so these files were regenerated from the JSON.
- **Rejected this iteration:**
  - A verb-only `forward|transmit|relay ... to <email>` pattern, which false-triggered on 8/15 benign red-team texts.
  - Adding the `send`, `sent` or `report(ed)` verbs.
  - Context-aware hidden-char exemptions. The best browsesafe F1 was 0.6413, which is no net gain.
  - Urgency/verification lure rules, which hit 243 benign vs 247 injection samples on browsesafe.
