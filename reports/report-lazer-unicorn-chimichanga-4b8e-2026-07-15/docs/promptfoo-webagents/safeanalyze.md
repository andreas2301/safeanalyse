# safeanalyze Report: /tmp/safeanalyze-doc-webagents

- **Started:** 2026-07-15 00:07:37 UTC
- **Completed:** 2026-07-15 00:07:37 UTC
- **Total findings:** 31

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 1 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 31 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 17 |
| High | 14 |

### Findings by source

| Source | Count |
| --- | --- |
| hiddenchars | 14 |
| yara | 17 |

## Findings

### Critical (17)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| prompt_injection_comment | index.html | 28 | 3254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 56 | 1073 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 58 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 71 | 187 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 81 | 354 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 88 | 68 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 88 | 94 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 88 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 88 | 199 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 88 | 247 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 89 | 75 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 89 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 90 | 2392 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 91 | 38 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 92 | 2550 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 119 | 2432 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | index.html | 119 | 2504 | Comment containing prompt injection keywords | yara | deterministic |

### High (14)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| hidden_char_zero_width | index.html | 32 | 204 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 52 | 254 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 54 | 219 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 59 | 224 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 64 | 244 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 69 | 342 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 72 | 249 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 74 | 239 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 80 | 259 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 83 | 219 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 88 | 274 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 93 | 234 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 103 | 184 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |
| hidden_char_zero_width | index.html | 112 | 199 | ZERO WIDTH SPACE (zero_width) in index.html | hiddenchars | deterministic |

## Errors

No errors recorded.

## Affected files

- `index.html`

