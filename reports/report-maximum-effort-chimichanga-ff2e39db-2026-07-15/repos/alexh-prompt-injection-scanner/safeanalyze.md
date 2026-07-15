# safeanalyze Report: /tmp/safeanalyze-prompt-injection-scanner

- **Started:** 2026-07-15 00:45:24 UTC
- **Duration:** 204 ms
- **Total findings:** 12

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 18 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 12 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 5 |
| High | 7 |

### Findings by source

| Source | Count |
| --- | --- |
| prompt-injection-scanner | 7 |
| yara | 5 |

## Findings

### Critical (5)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| PI010 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The ``run`` step embeds `${{ github.event.comment.body }}`, which originates from user-submitted issue or comment content. Any GitHub user can craft issue bodies or comments containing prompt injection payloads to manipulate AI agents triggered by this workflow. | prompt-injection-scanner | deterministic |
| PI010 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The ``run`` step embeds `${{ github.event.issue.body }}`, which originates from user-submitted issue or comment content. Any GitHub user can craft issue bodies or comments containing prompt injection payloads to manipulate AI agents triggered by this workflow. | prompt-injection-scanner | deterministic |
| prompt_injection_comment | README.md | 404 | 19 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 405 | 19 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | prompt_injection_scanner/rules.py | 997 | 55 | Comment containing prompt injection keywords | yara | deterministic |

### High (7)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| PI007 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The AI action `openai/openai-github-action@v1` receives the GitHub event expression `${{ github.event.pull_request.body }}` via the `prompt` input parameter. While this expression may not be the highest-risk source, it could still carry attacker-controlled content (e.g. via forked repositories, external contributors, or social engineering). | prompt-injection-scanner | deterministic |
| PI007 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The AI action `anthropic-ai/claude-action@v2` receives the GitHub event expression `${{ github.event.pull_request.title }}` via the `prompt` input parameter. While this expression may not be the highest-risk source, it could still carry attacker-controlled content (e.g. via forked repositories, external contributors, or social engineering). | prompt-injection-scanner | deterministic |
| PI007 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The AI action `langchain-ai/langchain-github-action@v1` receives the GitHub event expression `${{ github.event.issue.body }}` via the `issue-body` input parameter. While this expression may not be the highest-risk source, it could still carry attacker-controlled content (e.g. via forked repositories, external contributors, or social engineering). | prompt-injection-scanner | deterministic |
| PI007 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The AI action `langchain-ai/langchain-github-action@v1` receives the GitHub event expression `${{ github.event.comment.body }}` via the `comment-body` input parameter. While this expression may not be the highest-risk source, it could still carry attacker-controlled content (e.g. via forked repositories, external contributors, or social engineering). | prompt-injection-scanner | deterministic |
| PI010 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The ``run`` step embeds `${{ github.event.issue.title }}`, which originates from user-submitted issue or comment content. Any GitHub user can craft issue bodies or comments containing prompt injection payloads to manipulate AI agents triggered by this workflow. | prompt-injection-scanner | deterministic |
| obfuscated_javascript | tests/test_cli.py | 185 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | tests/test_reporter.py | 376 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |

## Errors

No errors recorded.

## Affected files

- `/tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml`
- `README.md`
- `prompt_injection_scanner/rules.py`
- `tests/test_cli.py`
- `tests/test_reporter.py`

