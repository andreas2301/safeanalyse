# safeanalyze Report: /tmp/safeanalyze-doc-scenarios.html

- **Started:** 2026-07-15 01:26:21 UTC
- **Duration:** 2765 ms
- **Total findings:** 4

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 1 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 4 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 1 |
| High | 3 |

### Findings by source

| Source | Count |
| --- | --- |
| hiddenchars | 3 |
| yara | 1 |

## Findings

### Critical (1)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| prompt_injection_comment | . | 28 | 3197 | Comment containing prompt injection keywords | yara | deterministic |

### High (3)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| hidden_char_zero_width | . | 31 | 189 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 39 | 219 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 50 | 249 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |

## Errors

No errors recorded.

## Affected files

- `.`

