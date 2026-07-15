# safeanalyze Test Reports

**Branch:** `report-spicy-beaver-sassy-3edea82c-2026-07-15`
**Device:** spicy-beaver-sassy-3edea82c
**Date:** 2026-07-15
**Tool version:** v0.3.4

This directory contains scan reports produced by `safeanalyze` against a collection of public prompt-injection test repositories, datasets, and documentation pages.

## Scanned targets

### Repositories

| Target | Source | Duration | Findings | Status |
|--------|--------|----------|----------|--------|
| microsoft-bipia | https://github.com/microsoft/BIPIA | 4.92 s | 331 | scanned |
| uiuc-injecagent | https://github.com/uiuc-kang-lab/InjecAgent | 9.98 s | 6924 | scanned |
| lakera-pint-benchmark | https://github.com/lakeraai/pint-benchmark | 2.72 s | 35 | scanned |
| alexh-prompt-injection-scanner | https://github.com/alexh-scrt/prompt-injection-scanner | 2.86 s | 142 | scanned |
| duriantaco-skylos | https://github.com/duriantaco/skylos | 27.43 s | 1371 | scanned |

### Datasets

| Target | Source | Status |
|--------|--------|--------|
| MAlmasabi/Indirect-Prompt-Injection-BIPIA-GPT | https://huggingface.co/datasets/MAlmasabi/Indirect-Prompt-Injection-BIPIA-GPT | gated / inaccessible |
| deepkeep/Prompt_Injection_v2 | https://huggingface.co/datasets/deepkeep/Prompt_Injection_v2 | gated / inaccessible |

### Documentation

| Target | Source | Duration | Findings | Status |
|--------|--------|----------|----------|--------|
| promptfoo-scenarios | https://www.promptfoo.dev/docs/configuration/scenarios/ | 2.82 s | 11 | scanned |
| promptfoo-webagents | https://www.promptfoo.dev/blog/indirect-prompt-injection-web-agents/ | 2.65 s | 31 | scanned |

## Output formats

Each successful scan produces:

- `safeanalyze.json` — full machine-readable report
- `safeanalyze.md` — human-readable Markdown summary
- `safeanalyze.html` — self-contained dashboard
- `safeanalyze.sarif` — SARIF v2.1.0 for IDE/security-tool consumption
- `duration.txt` — wall-clock scan duration in milliseconds

## Configuration used

The scans were run with `--mode thorough` and a test configuration that enables deterministic checks (YARA, entropy, hidden chars), disables the stochastic ONNX classifier, and enables the `alexh-scrt/prompt-injection-scanner`, Semgrep (`p/security-audit`), and TruffleHog external bridges.

See `/tmp/safeanalyze-test-config-v034.yaml` on the scanning host for the exact config.

## Notes

- Two Hugging Face datasets could not be accessed because they are gated/private.
- Report durations are recorded as `duration_ms` in `safeanalyze.json`.
- Binary files (images, PDFs, archives, fonts, compiled artifacts) are skipped by the text-based checks as of v0.2.6.
- Detected credential-like values are redacted in report output as of v0.2.7.
- Indirect prompt-injection, tool-injection, delimiter-breakout, encoded-injection, markdown-injection, and instruction-to-include-code rules were added in v0.2.8/v0.2.9/v0.3.0.
- Semgrep and TruffleHog were enabled by default in thorough mode in v0.3.1.
- v0.3.2 added `data_exfiltration_email`, `account_access_request`, `output_constraint`, `system_boundary`, and `template_injection` rules to catch social-engineering and system-boundary prompt-injection patterns.
