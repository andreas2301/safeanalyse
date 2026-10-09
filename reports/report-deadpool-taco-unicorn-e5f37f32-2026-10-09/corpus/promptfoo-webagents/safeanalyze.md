# safeanalyze Report: /tmp/safeanalyze-doc-webagents.html

- **Started:** 2026-10-09 21:45:23 UTC
- **Duration:** 2983 ms
- **Total findings:** 33

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 1 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 33 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 19 |
| High | 14 |

### Findings by source

| Source | Count |
| --- | --- |
| hiddenchars | 14 |
| yara | 19 |

## Findings

### Critical (19)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| prompt_injection_comment | . | 30 | 2764 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 58 | 1073 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 60 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 63 | 2907 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 73 | 187 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 83 | 243 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 83 | 354 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 90 | 68 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 90 | 94 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 90 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 90 | 199 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 90 | 247 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 91 | 75 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 91 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 92 | 2392 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 93 | 38 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 94 | 2550 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 121 | 2432 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 121 | 2504 | Comment containing prompt injection keywords | yara | deterministic |

### High (14)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| hidden_char_zero_width | . | 34 | 204 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 54 | 254 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 56 | 219 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 61 | 224 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 66 | 244 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 71 | 342 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 74 | 249 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 76 | 239 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 82 | 259 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 85 | 219 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 90 | 274 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 95 | 234 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 105 | 184 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 114 | 199 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |

## Errors

No errors recorded.

## Affected files

- `.`

