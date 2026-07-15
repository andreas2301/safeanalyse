# safeanalyze Report: /tmp/safeanalyze-prompt-injection-scanner

- **Started:** 2026-07-15 02:03:58 UTC
- **Duration:** 2383 ms
- **Total findings:** 142

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 18 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 142 |
| Errors | 1 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 5 |
| High | 9 |
| Medium | 128 |

### Findings by source

| Source | Count |
| --- | --- |
| prompt-injection-scanner | 7 |
| yara | 135 |

## Findings

### Critical (5)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| PI010 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The ``run`` step embeds `${{ github.event.comment.body }}`, which originates from user-submitted issue or comment content. Any GitHub user can craft issue bodies or comments containing prompt injection payloads to manipulate AI agents triggered by this workflow. | prompt-injection-scanner | deterministic |
| PI010 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The ``run`` step embeds `${{ github.event.issue.body }}`, which originates from user-submitted issue or comment content. Any GitHub user can craft issue bodies or comments containing prompt injection payloads to manipulate AI agents triggered by this workflow. | prompt-injection-scanner | deterministic |
| prompt_injection_comment | README.md | 404 | 19 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 405 | 19 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | prompt_injection_scanner/rules.py | 997 | 55 | Comment containing prompt injection keywords | yara | deterministic |

### High (9)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| PI007 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The AI action `openai/openai-github-action@v1` receives the GitHub event expression `${{ github.event.pull_request.body }}` via the `prompt` input parameter. While this expression may not be the highest-risk source, it could still carry attacker-controlled content (e.g. via forked repositories, external contributors, or social engineering). | prompt-injection-scanner | deterministic |
| PI007 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The AI action `anthropic-ai/claude-action@v2` receives the GitHub event expression `${{ github.event.pull_request.title }}` via the `prompt` input parameter. While this expression may not be the highest-risk source, it could still carry attacker-controlled content (e.g. via forked repositories, external contributors, or social engineering). | prompt-injection-scanner | deterministic |
| PI007 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The AI action `langchain-ai/langchain-github-action@v1` receives the GitHub event expression `${{ github.event.issue.body }}` via the `issue-body` input parameter. While this expression may not be the highest-risk source, it could still carry attacker-controlled content (e.g. via forked repositories, external contributors, or social engineering). | prompt-injection-scanner | deterministic |
| PI007 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The AI action `langchain-ai/langchain-github-action@v1` receives the GitHub event expression `${{ github.event.comment.body }}` via the `comment-body` input parameter. While this expression may not be the highest-risk source, it could still carry attacker-controlled content (e.g. via forked repositories, external contributors, or social engineering). | prompt-injection-scanner | deterministic |
| PI010 | /tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml | 0 | 0 | The ``run`` step embeds `${{ github.event.issue.title }}`, which originates from user-submitted issue or comment content. Any GitHub user can craft issue bodies or comments containing prompt injection payloads to manipulate AI agents triggered by this workflow. | prompt-injection-scanner | deterministic |
| indirect_prompt_injection | prompt_injection_scanner/rules.py | 359 | 26 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| indirect_prompt_injection | prompt_injection_scanner/rules.py | 614 | 26 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| obfuscated_javascript | tests/test_cli.py | 185 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | tests/test_reporter.py | 376 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |

### Medium (128)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| template_injection | README.md | 152 | 51 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | README.md | 179 | 14 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | README.md | 214 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | README.md | 270 | 64 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | README.md | 371 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | README.md | 378 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | README.md | 392 | 14 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | README.md | 400 | 15 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | README.md | 411 | 14 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | prompt_injection_scanner/rules.py | 64 | 36 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | prompt_injection_scanner/rules.py | 224 | 59 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | prompt_injection_scanner/rules.py | 321 | 37 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | prompt_injection_scanner/rules.py | 359 | 39 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | prompt_injection_scanner/rules.py | 614 | 39 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | prompt_injection_scanner/rules.py | 703 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | prompt_injection_scanner/rules.py | 903 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | prompt_injection_scanner/rules.py | 1015 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/safe_workflow.yml | 86 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 32 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 33 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 50 | 48 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 56 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 70 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 71 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 79 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 80 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 86 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 87 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 95 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 96 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 110 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 111 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 112 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 119 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/fixtures/vulnerable_workflow.yml | 119 | 85 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_cli.py | 70 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_cli.py | 71 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_cli.py | 267 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_models.py | 38 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_models.py | 124 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_models.py | 128 | 38 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_models.py | 137 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_models.py | 139 | 42 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_reporter.py | 41 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_reporter.py | 327 | 57 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_reporter.py | 434 | 57 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 99 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 105 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 111 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 117 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 118 | 16 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 124 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 129 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 134 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 142 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 147 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 229 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 237 | 52 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 252 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 253 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 267 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 273 | 45 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 283 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 284 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 292 | 37 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 298 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 315 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 326 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 336 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 347 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 358 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 359 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 377 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 395 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 406 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 415 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 424 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 444 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 456 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 465 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 475 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 476 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 494 | 51 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 504 | 58 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 511 | 55 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 521 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 541 | 37 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 552 | 37 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 575 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 576 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 595 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 608 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 617 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 635 | 45 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 648 | 45 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 660 | 45 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 689 | 51 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 699 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 708 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 714 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 735 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 743 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 751 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 757 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 779 | 44 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 791 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 799 | 38 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 823 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 836 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 845 | 36 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 855 | 38 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 875 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 886 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 894 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 903 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 963 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 972 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 982 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_rules.py | 1013 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_scanner.py | 70 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_scanner.py | 86 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_scanner.py | 87 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_scanner.py | 102 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_scanner.py | 109 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_scanner.py | 317 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_scanner.py | 321 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_scanner.py | 400 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | tests/test_scanner.py | 404 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |

## Errors

- **[semgrep]** skipped: target contains 18 files, below 50-file threshold

## Affected files

- `/tmp/safeanalyze-prompt-injection-scanner/tests/fixtures/vulnerable_workflow.yml`
- `README.md`
- `prompt_injection_scanner/rules.py`
- `tests/fixtures/safe_workflow.yml`
- `tests/fixtures/vulnerable_workflow.yml`
- `tests/test_cli.py`
- `tests/test_models.py`
- `tests/test_reporter.py`
- `tests/test_rules.py`
- `tests/test_scanner.py`

