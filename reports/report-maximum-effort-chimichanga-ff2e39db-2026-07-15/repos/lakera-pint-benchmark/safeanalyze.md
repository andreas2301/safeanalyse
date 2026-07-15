# safeanalyze Report: /tmp/safeanalyze-pint-benchmark

- **Started:** 2026-07-15 00:45:24 UTC
- **Duration:** 164 ms
- **Total findings:** 27

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 32 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 27 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 25 |
| Medium | 2 |

### Findings by source

| Source | Count |
| --- | --- |
| yara | 27 |

## Findings

### Critical (25)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| prompt_injection_comment | DETAILS.md | 33 | 75 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | DETAILS.md | 35 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 11 | 129 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 39 | 4 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 39 | 76 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 109 | 4 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 109 | 34 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 109 | 96 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 138 | 43 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 174 | 59 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | benchmark/data/example-dataset.yaml | 16 | 10 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | benchmark/data/example-dataset.yaml | 21 | 26 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | benchmark/data/example-dataset.yaml | 21 | 110 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | benchmark/data/example-dataset.yaml | 21 | 164 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | benchmark/data/example-dataset.yaml | 21 | 415 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | benchmark/data/example-dataset.yaml | 22 | 14 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | examples/hugging-face/epivolis/hyperion.md | 8 | 120 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | results/aporia_guardrails.md | 14 | 1 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | results/azure_content_safety.md | 14 | 1 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | results/bedrock_guardrails.md | 14 | 1 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | results/lakera_guard.md | 14 | 1 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | results/llama_prompt_guard.md | 14 | 1 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | results/llama_prompt_guard_2.md | 14 | 1 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | results/model_armor.md | 14 | 1 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | results/protect_ai.md | 14 | 1 | Comment containing prompt injection keywords | yara | deterministic |

### Medium (2)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| credential_hardcode | benchmark/example.env | 7 | 14 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | benchmark/example.env | 13 | 14 | Potential hardcoded credentials | yara | deterministic |

## Errors

No errors recorded.

## Affected files

- `DETAILS.md`
- `README.md`
- `benchmark/data/example-dataset.yaml`
- `benchmark/example.env`
- `examples/hugging-face/epivolis/hyperion.md`
- `results/aporia_guardrails.md`
- `results/azure_content_safety.md`
- `results/bedrock_guardrails.md`
- `results/lakera_guard.md`
- `results/llama_prompt_guard.md`
- `results/llama_prompt_guard_2.md`
- `results/model_armor.md`
- `results/protect_ai.md`

