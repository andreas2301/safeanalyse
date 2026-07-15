# safeanalyze Report: /tmp/safeanalyze-doc-webagents.html

- **Started:** 2026-07-15 01:53:17 UTC
- **Duration:** 1558 ms
- **Total findings:** 31

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 1 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 31 |
| Errors | 1 |

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
| prompt_injection_comment | . | 28 | 3254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 56 | 1073 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 58 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 71 | 187 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 81 | 354 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 88 | 68 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 88 | 94 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 88 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 88 | 199 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 88 | 247 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 89 | 75 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 89 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 90 | 2392 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 91 | 38 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 92 | 2550 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 119 | 2432 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | . | 119 | 2504 | Comment containing prompt injection keywords | yara | deterministic |

### High (14)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| hidden_char_zero_width | . | 32 | 204 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 52 | 254 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 54 | 219 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 59 | 224 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 64 | 244 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 69 | 342 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 72 | 249 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 74 | 239 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 80 | 259 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 83 | 219 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 88 | 274 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 93 | 234 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 103 | 184 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 112 | 199 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |

## Errors

- **[semgrep]** skipped: target contains 1 files, below 50-file threshold

## Affected files

- `.`

