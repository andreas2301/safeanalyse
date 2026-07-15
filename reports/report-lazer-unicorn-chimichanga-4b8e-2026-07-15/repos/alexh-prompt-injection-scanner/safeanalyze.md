# safeanalyze Report: /tmp/safeanalyze-prompt-injection-scanner

- **Started:** 2026-07-15 00:06:11 UTC
- **Completed:** 2026-07-15 00:06:11 UTC
- **Total findings:** 5

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 18 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 5 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 3 |
| High | 2 |

### Findings by source

| Source | Count |
| --- | --- |
| yara | 5 |

## Findings

### Critical (3)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| prompt_injection_comment | README.md | 404 | 19 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 405 | 19 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | prompt_injection_scanner/rules.py | 997 | 55 | Comment containing prompt injection keywords | yara | deterministic |

### High (2)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| obfuscated_javascript | tests/test_cli.py | 185 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | tests/test_reporter.py | 376 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |

## Errors

No errors recorded.

## Affected files

- `README.md`
- `prompt_injection_scanner/rules.py`
- `tests/test_cli.py`
- `tests/test_reporter.py`

