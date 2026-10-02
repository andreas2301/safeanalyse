# safeanalyze Report: /tmp/safeanalyze-doc-scenarios.html

- **Started:** 2026-10-02 21:49:29 UTC
- **Duration:** 2755 ms
- **Total findings:** 11

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 1 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 11 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 1 |
| High | 3 |
| Medium | 7 |

### Findings by source

| Source | Count |
| --- | --- |
| hiddenchars | 3 |
| yara | 8 |

## Findings

### Critical (1)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| prompt_injection_comment | . | 30 | 2707 | Comment containing prompt injection keywords | yara | deterministic |

### High (3)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| hidden_char_zero_width | . | 33 | 189 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 41 | 219 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |
| hidden_char_zero_width | . | 52 | 249 | ZERO WIDTH SPACE (zero_width) in . | hiddenchars | deterministic |

### Medium (7)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| template_injection | . | 35 | 497 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | . | 35 | 511 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | . | 35 | 710 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | . | 35 | 724 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | . | 38 | 8481 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | . | 38 | 10855 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | . | 38 | 13324 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |

## Errors

No errors recorded.

## Affected files

- `.`

