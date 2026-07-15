# safeanalyze Test Reports

**Branch:** `report-lazer-unicorn-chimichanga-4b8e-2026-07-15`  
**Device:** lazer-unicorn-chimichanga-4b8e  
**Date:** 2026-07-15  
**Tool version:** v0.2.1

This directory contains scan reports produced by `safeanalyze` against a collection of public prompt-injection test repositories, datasets, and documentation pages.

## Scanned targets

### Repositories

| Target | Source | Duration | Status |
|--------|--------|----------|--------|
| microsoft-bipia | https://github.com/microsoft/BIPIA | 1.51 s | scanned |
| uiuc-injecagent | https://github.com/uiuc-kang-lab/InjecAgent | 5.35 s | scanned |
| lakera-pint-benchmark | https://github.com/lakeraai/pint-benchmark | 0.80 s | scanned |
| alexh-prompt-injection-scanner | https://github.com/alexh-scrt/prompt-injection-scanner | 0.20 s | scanned |
| duriantaco-skylos | https://github.com/duriantaco/skylos | 86.23 s | scanned |

### Datasets

| Target | Source | Status |
|--------|--------|--------|
| MAlmasabi/Indirect-Prompt-Injection-BIPIA-GPT | https://huggingface.co/datasets/MAlmasabi/Indirect-Prompt-Injection-BIPIA-GPT | gated / inaccessible |
| deepkeep/Prompt_Injection_v2 | https://huggingface.co/datasets/deepkeep/Prompt_Injection_v2 | gated / inaccessible |

### Documentation

| Target | Source | Duration | Status |
|--------|--------|----------|--------|
| promptfoo scenarios | https://www.promptfoo.dev/docs/configuration/scenarios/ | 0.06 s | scanned |
| promptfoo web agents | https://www.promptfoo.dev/blog/indirect-prompt-injection-web-agents/ | 0.05 s | scanned |

## Output formats

Each successful scan produces:

- `safeanalyze.json` — full machine-readable report
- `safeanalyze.md` — human-readable Markdown summary
- `safeanalyze.html` — self-contained dashboard
- `safeanalyze.sarif` — SARIF v2.1.0 for IDE/security-tool consumption
- `duration.txt` — wall-clock scan duration in seconds

## Configuration used

The scans were run with `--mode thorough` and a test configuration that:

- Enables built-in deterministic checks (YARA, entropy, hidden chars)
- Disables `fail_on_findings` so reports are always written
- Disables external scanners and ML by default (no model/download available in this run)
- Outputs all four report formats
- Caps reports at 10,000 findings to prevent oversized files

See `/tmp/safeanalyze-test-config.yaml` on the scanning host for the exact config.

## Notes

- External scanners (Semgrep, Bumblebee, prompt-injection-scanner, Gitleaks, TruffleHog) were not installed for this run, so they did not contribute findings.
- The ML prompt-injection classifier requires a downloaded ONNX model (`~/.safeanalyze/models/...`) which was not present; it was disabled for this run.
- Two Hugging Face datasets could not be accessed because they are gated/private.
- Compared to v0.2.0 baseline, `duriantaco-skylos` scan time dropped from 575.78 s to 86.23 s after capping entropy findings and skipping very long strings.
