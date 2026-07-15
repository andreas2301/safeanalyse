# safeanalyze Test Reports

**Branch:** `report-juggernaut-chimichanga-beaver-01ddce8b-2026-07-15`
**Device:** juggernaut-chimichanga-beaver-01ddce8b
**Date:** 2026-07-15
**Tool version:** v0.3.0

This directory contains scan reports produced by `safeanalyze` against a collection of public prompt-injection test repositories, datasets, and documentation pages.

## Scanned targets

### Repositories

| Target | Source | Duration | Findings | Status |
|--------|--------|----------|----------|--------|
| microsoft-bipia | https://github.com/microsoft/BIPIA | 2.66 s | 324 | scanned |
| uiuc-injecagent | https://github.com/uiuc-kang-lab/InjecAgent | 5.19 s | 2620 | scanned |
| lakera-pint-benchmark | https://github.com/lakeraai/pint-benchmark | 0.26 s | 35 | scanned |
| alexh-prompt-injection-scanner | https://github.com/alexh-scrt/prompt-injection-scanner | 0.30 s | 14 | scanned |
| duriantaco-skylos | https://github.com/duriantaco/skylos | 14.45 s | 1128 | scanned |

### Datasets

| Target | Source | Status |
|--------|--------|--------|
| MAlmasabi/Indirect-Prompt-Injection-BIPIA-GPT | https://huggingface.co/datasets/MAlmasabi/Indirect-Prompt-Injection-BIPIA-GPT | gated / inaccessible |
| deepkeep/Prompt_Injection_v2 | https://huggingface.co/datasets/deepkeep/Prompt_Injection_v2 | gated / inaccessible |

### Documentation

| Target | Source | Duration | Findings | Status |
|--------|--------|----------|----------|--------|
| promptfoo-scenarios | https://www.promptfoo.dev/docs/configuration/scenarios/ | 0.08 s | 4 | scanned |
| promptfoo-webagents | https://www.promptfoo.dev/blog/indirect-prompt-injection-web-agents/ | 0.08 s | 31 | scanned |

## Output formats

Each successful scan produces:

- `safeanalyze.json` — full machine-readable report
- `safeanalyze.md` — human-readable Markdown summary
- `safeanalyze.html` — self-contained dashboard
- `safeanalyze.sarif` — SARIF v2.1.0 for IDE/security-tool consumption
- `duration.txt` — wall-clock scan duration in milliseconds

## Configuration used

The scans were run with `--mode thorough` and a test configuration that enables deterministic checks (YARA, entropy, hidden chars), disables the stochastic ONNX classifier, and enables the `alexh-scrt/prompt-injection-scanner` external bridge.

See `/tmp/safeanalyze-test-config-v030.yaml` on the scanning host for the exact config.

## Notes

- Two Hugging Face datasets could not be accessed because they are gated/private.
- Report durations are recorded as `duration_ms` in `safeanalyze.json`.
- Binary files (images, PDFs, archives, fonts, compiled artifacts) are skipped by the text-based checks as of v0.2.6.
- Detected credential-like values are redacted in report output as of v0.2.7.
- Indirect prompt-injection, tool-injection, delimiter-breakout, encoded-injection, markdown-injection, and instruction-to-include-code rules were added in v0.2.8/v0.2.9/v0.3.0.
