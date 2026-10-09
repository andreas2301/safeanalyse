# safeanalyze Report: /tmp/safeanalyze-skylos

- **Started:** 2026-10-09 21:34:30 UTC
- **Duration:** 16771 ms
- **Total findings:** 2264

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 1951 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 2264 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 94 |
| High | 656 |
| Medium | 1128 |
| Low | 386 |

### Findings by source

| Source | Count |
| --- | --- |
| entropy | 780 |
| semgrep | 55 |
| trufflehog | 53 |
| yara | 1376 |

## Findings

### Critical (94)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| backdoor_indicator | CHANGELOG.md | 593 | 25 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | docs/agent-behavior-testing.md | 168 | 58 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | docs/agent-behavior-testing.md | 176 | 1 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/audit/candidates.py | 529 | 65 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/debt/baseline.py | 259 | 26 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/debt/baseline.py | 280 | 41 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/debt/baseline.py | 449 | 41 | Common backdoor or persistence patterns | yara | deterministic |
| prompt_injection_comment | skylos/defend/owasp.py | 111 | 22 | Comment containing prompt injection keywords | yara | deterministic |
| backdoor_indicator | skylos/llm/investigator/reviewer_packs.py | 204 | 95 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/llm/investigator/reviewer_packs.py | 221 | 30 | Common backdoor or persistence patterns | yara | deterministic |
| prompt_injection_comment | skylos/llm/prompts.py | 30 | 40 | Comment containing prompt injection keywords | yara | deterministic |
| backdoor_indicator | skylos/rules/ai_defect/module_facts_index.py | 123 | 10 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/rules/catalog.py | 280 | 30 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/security/command_guard_exfil.py | 102 | 1 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/security/command_guard_exfil.py | 125 | 51 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/security/command_guard_policy.py | 83 | 20 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/security/command_guard_policy.py | 138 | 17 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | test/test_analyzer.py | 4251 | 16 | Common backdoor or persistence patterns | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6324 | 27 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6324 | 27 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6326 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6326 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6349 | 39 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6349 | 39 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6376 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6376 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6425 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 6425 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| backdoor_indicator | test/test_feedback.py | 39 | 17 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | test/test_grep_cache.py | 253 | 15 | Common backdoor or persistence patterns | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 73 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 73 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 91 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 91 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 98 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 98 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 145 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 145 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 157 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 180 | 21 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 180 | 21 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 191 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 191 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 223 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 223 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 233 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 233 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 257 | 14 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 257 | 14 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 268 | 26 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 280 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 280 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 317 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 338 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 352 | 38 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 352 | 38 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 366 | 30 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 366 | 30 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 402 | 29 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 402 | 29 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 428 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 428 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 432 | 27 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 432 | 27 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 452 | 48 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 452 | 48 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 464 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 464 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 476 | 47 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 476 | 47 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 487 | 55 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 487 | 55 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 497 | 18 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 497 | 18 | Comment containing prompt injection keywords | yara | deterministic |
| system_boundary | test/test_mcp_reliability.py | 41 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| prompt_injection_comment | test/test_mcp_reliability.py | 41 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| system_boundary | test/test_mcp_reliability.py | 41 | 35 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 25 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| prompt_injection_comment | test/test_mcp_rules.py | 25 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 25 | 39 | System prompt boundary marker or system instruction override | yara | deterministic |
| prompt_injection_comment | test/test_mcp_rules.py | 39 | 8 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_mcp_rules.py | 39 | 8 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_mcp_rules.py | 65 | 40 | Comment containing prompt injection keywords | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 89 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| prompt_injection_comment | test/test_mcp_rules.py | 89 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 89 | 28 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 101 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 101 | 52 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 114 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 114 | 20 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 402 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| prompt_injection_comment | test/test_mcp_rules.py | 402 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 402 | 31 | System prompt boundary marker or system instruction override | yara | deterministic |
| backdoor_indicator | test/test_triage_learner.py | 497 | 11 | Common backdoor or persistence patterns | yara | deterministic |

### High (656)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/extreme_grounding_framework/services/hooks.py | 6 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/extreme_grounding_framework/services/hooks.py | 26 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/extreme_grounding_framework/tools/admin.py | 8 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/flask_getter_shell/app.py | 11 | 38 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/flask_handler_security/app.py | 19 | 47 | Found 'subprocess' function 'check_output' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/shell_hook_runner/hooks.py | 12 | 38 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/static_blind_plugin_dispatch/plugins/audit.py | 7 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/static_blind_plugin_dispatch/plugins/audit.py | 12 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.requests.security.disabled-cert-validation.disabled-cert-validation | /tmp/safeanalyze-skylos/benchmarks/ai_code_defects/fixtures/disabled_security_control/app.py | 5 | 12 | Certificate verification has been explicitly disabled. This permits insecure connections to insecure servers. Re-enable certification validation. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/dead_code/fixtures/extreme_grounding_framework/services/hooks.py | 6 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/dead_code/fixtures/extreme_grounding_framework/services/hooks.py | 26 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/dead_code/fixtures/extreme_grounding_framework/tools/admin.py | 8 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/dead_code/fixtures/static_blind_plugin_dispatch/plugins/audit.py | 7 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/dead_code/fixtures/static_blind_plugin_dispatch/plugins/audit.py | 12 | 42 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.django.security.injection.command.command-injection-os-system.command-injection-os-system | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 99 | 5 | Request data detected in os.system. This could be vulnerable to a command injection and should be avoided. If this must be done, use the 'subprocess' module instead and pass the arguments as a list. See https://owasp.org/www-community/attacks/Command_Injection for more information. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 101 | 43 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.django.security.injection.command.command-injection-os-system.command-injection-os-system | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 107 | 5 | Request data detected in os.system. This could be vulnerable to a command injection and should be avoided. If this must be done, use the 'subprocess' module instead and pass the arguments as a list. See https://owasp.org/www-community/attacks/Command_Injection for more information. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 109 | 43 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.django.security.injection.command.command-injection-os-system.command-injection-os-system | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 115 | 5 | Request data detected in os.system. This could be vulnerable to a command injection and should be avoided. If this must be done, use the 'subprocess' module instead and pass the arguments as a list. See https://owasp.org/www-community/attacks/Command_Injection for more information. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 117 | 43 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.django.security.injection.command.command-injection-os-system.command-injection-os-system | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 123 | 5 | Request data detected in os.system. This could be vulnerable to a command injection and should be avoided. If this must be done, use the 'subprocess' module instead and pass the arguments as a list. See https://owasp.org/www-community/attacks/Command_Injection for more information. | semgrep | deterministic |
| python.lang.security.audit.subprocess-shell-true.subprocess-shell-true | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 125 | 43 | Found 'subprocess' function 'run' with 'shell=True'. This is dangerous because this call will spawn the command using a shell process. Doing so propagates current shell settings and variables, which makes it much easier for a malicious actor to execute commands. Use 'shell=False' instead. | semgrep | deterministic |
| python.requests.security.disabled-cert-validation.disabled-cert-validation | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 132 | 9 | Certificate verification has been explicitly disabled. This permits insecure connections to insecure servers. Re-enable certification validation. | semgrep | deterministic |
| python.flask.security.injection.user-eval.eval-injection | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 146 | 5 | Detected user data flowing into eval. This is code injection and should be avoided. | semgrep | deterministic |
| gitlab | /tmp/safeanalyze-skylos/docs/examples/gitlab-code-quality.yml | 2 | 0 | Potential Gitlab secret detected | trufflehog | deterministic |
| gitlab | /tmp/safeanalyze-skylos/docs/examples/gitlab-managed-upload.yml | 2 | 0 | Potential Gitlab secret detected | trufflehog | deterministic |
| javascript.lang.security.detect-child-process.detect-child-process | /tmp/safeanalyze-skylos/editors/vscode/src/verifyCore.ts | 65 | 24 | Detected calls to child_process from a function argument `request`. This could lead to a command injection if the input is user controllable. Try to avoid calls to child_process, and if it is needed ensure user input is correctly sanitized or sandboxed.  | semgrep | deterministic |
| gitlab | /tmp/safeanalyze-skylos/skylos/api/__init__.py | 830 | 0 | Potential Gitlab secret detected | trufflehog | deterministic |
| python.lang.security.use-defused-xml.use-defused-xml | /tmp/safeanalyze-skylos/skylos/done/runner.py | 31 | 1 | The Python documentation recommends using `defusedxml` instead of `xml` because the native Python `xml` library is vulnerable to XML External Entity (XXE) attacks. These attacks can leak confidential data and "XML bombs" can cause denial of service. | semgrep | deterministic |
| python.lang.security.use-defused-xml.use-defused-xml | /tmp/safeanalyze-skylos/skylos/rules/ai_defect/dependency_hallucination.py | 14 | 1 | The Python documentation recommends using `defusedxml` instead of `xml` because the native Python `xml` library is vulnerable to XML External Entity (XXE) attacks. These attacks can leak confidential data and "XML bombs" can cause denial of service. | semgrep | deterministic |
| python.lang.security.use-defused-xml.use-defused-xml | /tmp/safeanalyze-skylos/skylos/rules/ai_defect/manifest_dependency_hallucination.py | 10 | 1 | The Python documentation recommends using `defusedxml` instead of `xml` because the native Python `xml` library is vulnerable to XML External Entity (XXE) attacks. These attacks can leak confidential data and "XML bombs" can cause denial of service. | semgrep | deterministic |
| gitlab | /tmp/safeanalyze-skylos/skylos/rules/config/cicd/gitlab_ci.py | 754 | 0 | Potential Gitlab secret detected | trufflehog | deterministic |
| python.lang.security.use-defused-xml.use-defused-xml | /tmp/safeanalyze-skylos/skylos/rules/sca/vulnerability_scanner.py | 8 | 1 | The Python documentation recommends using `defusedxml` instead of `xml` because the native Python `xml` library is vulnerable to XML External Entity (XXE) attacks. These attacks can leak confidential data and "XML bombs" can cause denial of service. | semgrep | deterministic |
| python.lang.security.use-defused-xml.use-defused-xml | /tmp/safeanalyze-skylos/skylos/rules/sca/vulnerability_scanner.py | 11 | 1 | The Python documentation recommends using `defusedxml` instead of `xml` because the native Python `xml` library is vulnerable to XML External Entity (XXE) attacks. These attacks can leak confidential data and "XML bombs" can cause denial of service. | semgrep | deterministic |
| python.lang.security.use-defused-xml.use-defused-xml | /tmp/safeanalyze-skylos/skylos/visitors/languages/csharp/reachability.py | 14 | 1 | The Python documentation recommends using `defusedxml` instead of `xml` because the native Python `xml` library is vulnerable to XML External Entity (XXE) attacks. These attacks can leak confidential data and "XML bombs" can cause denial of service. | semgrep | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_agent_behavior_openai.py | 273 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_gitlab_report.py | 360 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_gitlab_report.py | 377 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_gitlab_report.py | 383 | 0 | Potential URI secret detected | trufflehog | deterministic |
| gitlab | /tmp/safeanalyze-skylos/test/test_gitlab_scan_receipt.py | 319 | 0 | Potential Gitlab secret detected | trufflehog | deterministic |
| stripe | /tmp/safeanalyze-skylos/test/test_mcp_rules.py | 406 | 0 | Potential Stripe secret detected | trufflehog | deterministic |
| stripe | /tmp/safeanalyze-skylos/test/test_mcp_rules.py | 413 | 0 | Potential Stripe secret detected | trufflehog | deterministic |
| stripe | /tmp/safeanalyze-skylos/test/test_mcp_rules.py | 414 | 0 | Potential Stripe secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_pipfile_lockfile.py | 15 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_pnpm_regressions.py | 216 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_poetry_lockfile.py | 193 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_poetry_lockfile.py | 208 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_poetry_lockfile.py | 278 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_sca_pipfile.py | 22 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_sca_poetry_yarn.py | 33 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_sca_poetry_yarn.py | 51 | 0 | Potential URI secret detected | trufflehog | deterministic |
| gitlab | /tmp/safeanalyze-skylos/test/test_secrets.py | 104 | 0 | Potential Gitlab secret detected | trufflehog | deterministic |
| gitlab | /tmp/safeanalyze-skylos/test/test_secrets.py | 110 | 0 | Potential Gitlab secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 1 | 0 | Potential URI secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 16 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 19 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| rabbitmq | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 22 | 0 | Potential RabbitMQ secret detected | trufflehog | deterministic |
| mongodb | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 23 | 0 | Potential MongoDB secret detected | trufflehog | deterministic |
| rabbitmq | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 25 | 0 | Potential RabbitMQ secret detected | trufflehog | deterministic |
| mongodb | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 26 | 0 | Potential MongoDB secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 33 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 35 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 36 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 38 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 40 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 41 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 41 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 42 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 43 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 46 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 47 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 48 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 122 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 135 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 157 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| postgres | /tmp/safeanalyze-skylos/test/test_secrets_credentials.py | 174 | 0 | Potential Postgres secret detected | trufflehog | deterministic |
| github | /tmp/safeanalyze-skylos/test/test_secrets_nonpy.py | 138 | 0 | Potential Github secret detected | trufflehog | deterministic |
| github | /tmp/safeanalyze-skylos/test/test_secrets_nonpy.py | 145 | 0 | Potential Github secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_shadow_compare.py | 741 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_uv_lockfile.py | 152 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_yarn_lockfile.py | 179 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_yarn_lockfile.py | 181 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_yarn_lockfile.py | 183 | 0 | Potential URI secret detected | trufflehog | deterministic |
| uri | /tmp/safeanalyze-skylos/test/test_yarn_lockfile.py | 200 | 0 | Potential URI secret detected | trufflehog | deterministic |
| data_exfiltration_email | SECURITY.md | 15 | 1 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/corpus/real_commits.json | 1121 | 72 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 13 | 207 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 17 | 211 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 20 | 198 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 23 | 220 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 291 | 227 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 291 | 244 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 292 | 227 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 292 | 244 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 324 | 226 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 324 | 243 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 325 | 226 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 325 | 243 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 349 | 255 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 349 | 272 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 350 | 255 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 350 | 272 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 433 | 235 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 433 | 252 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 438 | 216 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 439 | 216 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 440 | 216 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/findings.jsonl | 441 | 200 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmarks/agent-pr-bench/results/workspaces.json | 10123 | 69 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration | editors/vscode/out/ai.js | 348 | 28 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | editors/vscode/out/ai.js | 418 | 28 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | editors/vscode/out/ai.js | 667 | 24 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | editors/vscode/out/chatview.js | 148 | 28 | Potential data exfiltration patterns | yara | deterministic |
| obfuscated_javascript | editors/vscode/out/codelens.js | 100 | 29 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration | editors/vscode/src/ai.ts | 375 | 24 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | editors/vscode/src/ai.ts | 459 | 24 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | editors/vscode/src/ai.ts | 741 | 22 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | editors/vscode/src/chatview.ts | 141 | 24 | Potential data exfiltration patterns | yara | deterministic |
| obfuscated_javascript | editors/vscode/src/codelens.ts | 85 | 21 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | scripts/deep_audit_logic_benchmark.py | 26 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1106 | 10 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1106 | 34 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1106 | 42 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1106 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1106 | 55 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1106 | 60 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1107 | 7 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1107 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1107 | 28 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1107 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/analyzer.py | 1107 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/api/__init__.py | 523 | 33 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_paths.py | 52 | 28 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_paths.py | 52 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_paths.py | 54 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_paths.py | 54 | 44 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_paths.py | 54 | 48 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_preflight.py | 28 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_preflight.py | 28 | 44 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_preflight.py | 28 | 48 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_preflight.py | 62 | 40 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_preflight.py | 62 | 44 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_preflight.py | 62 | 48 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_preflight.py | 62 | 52 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_preflight.py | 62 | 56 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_preflight.py | 64 | 8 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 29 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 34 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 47 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 54 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 60 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 67 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/api/_upload_transport.py | 55 | 80 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/audit/investigator_tools/operations.py | 198 | 76 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/audit/investigator_tools/operations.py | 265 | 69 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/audit/investigator_tools/validation.py | 16 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/benchmarks/_jev_dead_code_dataset.py | 302 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/benchmarks/deep_audit_logic.py | 264 | 60 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/benchmarks/deep_audit_logic.py | 369 | 21 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/evidence.py | 85 | 23 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/evidence.py | 85 | 42 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/evidence.py | 86 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/evidence.py | 86 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 874 | 60 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 1412 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 1412 | 54 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 1412 | 81 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 1414 | 56 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 1439 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 1439 | 63 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/cache_cmd.py | 115 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/commands/credits_cmd.py | 15 | 39 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/commands/hook_cmd.py | 945 | 10 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/hook_cmd.py | 1315 | 60 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/preflight_cmd.py | 415 | 40 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/preflight_cmd.py | 415 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/review_cmd.py | 247 | 59 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/rules_cmd.py | 410 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/verify_verdict_cmd.py | 54 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/verify_verdict_cmd.py | 54 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/verify_verdict_cmd.py | 54 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/commands/verify_verdict_cmd.py | 54 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/api_symbol_truth.py | 311 | 33 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_common.py | 150 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_common.py | 150 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_common.py | 150 | 51 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_common.py | 150 | 59 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_common.py | 1650 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_common.py | 1650 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_common.py | 2088 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_common.py | 2088 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_language_strategies.py | 140 | 40 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_language_strategies.py | 140 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_python_strategy.py | 288 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_python_strategy.py | 288 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_python_strategy.py | 304 | 42 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_python_strategy.py | 304 | 62 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/js_api_surface_utils.py | 188 | 33 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/review_context.py | 214 | 20 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/review_context.py | 231 | 20 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/review_context.py | 282 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/review_context.py | 321 | 9 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/review_decisions.py | 230 | 20 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/review_decisions.py | 1014 | 56 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/deadcode/browser_refs.py | 621 | 56 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/done/base.py | 175 | 59 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/done/receipt.py | 49 | 29 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/done/receipt.py | 49 | 34 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/done/receipt.py | 49 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/engines/go/internal/analyzer/analyzer.go | 632 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/engines/go/internal/analyzer/analyzer.go | 632 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/engines/go/internal/analyzer/analyzer.go | 632 | 69 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| indirect_prompt_injection | skylos/engines/go/internal/analyzer/analyzer_exec_command_test.go | 25 | 2 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| indirect_prompt_injection | skylos/engines/go/internal/analyzer/analyzer_exec_command_test.go | 42 | 2 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| indirect_prompt_injection | skylos/engines/go/internal/analyzer/analyzer_exec_command_test.go | 60 | 2 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| indirect_prompt_injection | skylos/engines/go/internal/analyzer/analyzer_exec_command_test.go | 77 | 2 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| indirect_prompt_injection | skylos/engines/go/internal/analyzer/analyzer_exec_command_test.go | 122 | 2 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| obfuscated_javascript | skylos/engines/go/internal/symbols/symbols.go | 512 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/engines/go/internal/symbols/symbols.go | 515 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/engines/go/internal/symbols/symbols.go | 525 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/engines/go/internal/symbols/symbols.go | 528 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/engines/go/internal/symbols/typed_selectors.go | 76 | 20 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 795 | 33 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 795 | 66 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 795 | 70 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 795 | 74 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 795 | 78 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 808 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 808 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 808 | 47 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 808 | 51 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 808 | 62 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 808 | 66 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 809 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 809 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 810 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 810 | 40 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 812 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 812 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 812 | 40 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 812 | 44 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 813 | 34 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 813 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 813 | 42 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 813 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 997 | 66 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 1176 | 71 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/preflight/inspector.py | 1401 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/reporting/provenance.py | 196 | 28 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/reporting/provenance.py | 196 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/reporting/sarif.py | 73 | 16 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/ai_defect/dependency_hallucination.py | 431 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/ai_defect/pypi_wheel_modules.py | 45 | 23 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/ai_defect/pypi_wheel_modules.py | 45 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/ai_defect/pypi_wheel_modules.py | 48 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/ai_defect/pypi_wheel_modules.py | 48 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/ai_defect/pypi_wheel_modules.py | 471 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 40 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 41 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 42 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 43 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 44 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 45 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 46 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 47 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 48 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_fs/pytest_paths.py | 171 | 55 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 40 | 8 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 40 | 14 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 40 | 20 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 40 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 40 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 7 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 7 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 43 | 7 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 43 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 43 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 43 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/sca/yarn_lockfile.py | 319 | 48 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/secrets.py | 2494 | 47 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 9 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 10 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 11 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 12 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 13 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 14 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 15 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 16 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 17 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 18 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 19 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 20 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 36 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 37 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 38 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 39 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 40 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 41 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 42 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 43 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 44 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 45 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 46 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 47 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 48 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 49 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 50 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 51 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 52 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 53 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 54 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 55 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 56 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 57 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 58 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 59 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 60 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 61 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 62 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 63 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 64 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 65 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 66 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 67 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 68 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/canonicalize.py | 69 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/security/injection_scanner.py | 118 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/ui/terminal_report.py | 84 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/ui/terminal_report.py | 84 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/ui/terminal_report.py | 84 | 47 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/ui/terminal_report.py | 84 | 52 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/verdict.py | 62 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/verdict.py | 62 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 18 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 55 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/__init__.py | 32 | 59 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/danger.py | 59 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/visitors/languages/typescript/danger.py | 7809 | 30 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/source_lines.py | 6 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/source_lines.py | 6 | 56 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/type_safety.py | 168 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/workspace.py | 310 | 71 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| llm_tool_injection | test/test_agent_behavior_cli.py | 63 | 25 | Suspicious LLM tool or function-call payload | yara | deterministic |
| obfuscated_javascript | test/test_agent_behavior_cli.py | 123 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| llm_tool_injection | test/test_agent_behavior_contract.py | 486 | 29 | Suspicious LLM tool or function-call payload | yara | deterministic |
| obfuscated_javascript | test/test_analyzer.py | 5625 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_analyzer.py | 5625 | 60 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_analyzer.py | 5626 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_analyzer.py | 5626 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_analyzer.py | 5699 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_analyzer.py | 5699 | 44 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_analyzer.py | 5699 | 59 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_analyzer.py | 5699 | 63 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_audit_investigator_tools.py | 35 | 16 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_baseline_source.py | 163 | 45 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_behavior_baselines.py | 473 | 70 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_behavior_baselines.py | 474 | 68 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_behavior_changes.py | 106 | 61 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_behavior_changes.py | 131 | 16 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cicd_workflow.py | 223 | 62 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_gitlab.py | 101 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 311 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 317 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 324 | 53 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 328 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 328 | 45 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 333 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 340 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 357 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 357 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 357 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 357 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 357 | 66 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 359 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 359 | 40 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 359 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 359 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cli_preflight.py | 359 | 67 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| indirect_prompt_injection | test/test_community_rules.py | 53 | 13 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| indirect_prompt_injection | test/test_community_rules.py | 114 | 13 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| indirect_prompt_injection | test/test_community_rules.py | 135 | 13 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| obfuscated_javascript | test/test_dangerous.py | 50 | 66 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dangerous.py | 50 | 71 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dangerous.py | 50 | 77 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dangerous.py | 244 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 585 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 585 | 28 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 617 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 617 | 61 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 924 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 924 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_deep_audit_logic_benchmark_output.py | 108 | 10 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_deep_audit_logic_benchmark_output.py | 108 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_deep_audit_logic_benchmark_output.py | 121 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| indirect_prompt_injection | test/test_defend.py | 1821 | 5 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| obfuscated_javascript | test/test_dependency_providers.py | 605 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dependency_providers.py | 605 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dependency_providers.py | 605 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dependency_providers.py | 614 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dependency_providers.py | 614 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dependency_version_bump.py | 716 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_deserialization.py | 21 | 53 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_deserialization.py | 33 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_deserialization.py | 43 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_diff_dependencies.py | 347 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_discover.py | 407 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| indirect_prompt_injection | test/test_discover.py | 556 | 5 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| obfuscated_javascript | test/test_done_receipt_binding.py | 147 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_gitlab_delivery_receipt.py | 97 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_gitlab_report.py | 296 | 14 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_gitlab_report.py | 297 | 14 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_gitlab_report.py | 383 | 83 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_gitlab_report.py | 392 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_gitlab_report.py | 558 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_gitlab_report.py | 559 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_gitlab_report.py | 634 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_gitlab_upload_auth.py | 303 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_python_backend.py | 63 | 12 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_python_backend.py | 63 | 20 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_python_backend.py | 98 | 55 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_python_backend.py | 98 | 59 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_python_backend.py | 409 | 56 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_python_backend.py | 432 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_python_backend.py | 433 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_python_backend.py | 473 | 77 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 61 | 42 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 884 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 886 | 15 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 886 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 953 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1057 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1112 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1114 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1198 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1236 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1296 | 21 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1332 | 54 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1680 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1848 | 44 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1870 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1881 | 62 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1882 | 42 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1894 | 69 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1914 | 63 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1921 | 69 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1948 | 44 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1977 | 48 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 1998 | 63 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 2008 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 2097 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_grep_verify.py | 2623 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_hook_cmd.py | 1331 | 45 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 25 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 33 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 45 | 18 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 45 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 45 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 51 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 58 | 29 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 63 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 66 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 109 | 23 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 115 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 115 | 23 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 115 | 29 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 115 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 115 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 115 | 47 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 124 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 201 | 45 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 245 | 44 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| indirect_prompt_injection | test/test_injection_scanner.py | 280 | 26 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 329 | 28 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 392 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 508 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 517 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 528 | 34 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 542 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 559 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 559 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 559 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 559 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 559 | 55 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_injection_scanner.py | 559 | 61 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_java_property_resources.py | 46 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_java_property_resources.py | 46 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_java_property_resources.py | 52 | 12 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_java_property_resources.py | 52 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_java_property_resources.py | 53 | 14 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_java_property_resources.py | 53 | 20 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_jev_dataset_safety.py | 82 | 71 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_jev_dataset_safety.py | 140 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_jev_dataset_safety.py | 140 | 45 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_jev_triage.py | 179 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_mcp_rules.py | 53 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_npm_lockfile.py | 247 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_npm_lockfile.py | 374 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_osv_client.py | 169 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_osv_client.py | 361 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pipeline.py | 152 | 23 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pipeline.py | 152 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pnpm_lockfile.py | 806 | 55 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_preflight_inspector.py | 65 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_preflight_inspector.py | 804 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_preflight_inspector.py | 804 | 34 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_provenance.py | 182 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_provenance_attribution.py | 104 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 419 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 419 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 444 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 444 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 454 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 454 | 28 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 475 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 475 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 502 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pypi_wheel_inventory.py | 502 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pytest_path_parameters.py | 234 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_refactor_verify.py | 252 | 70 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_review_cmd.py | 226 | 45 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_review_cmd.py | 226 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_review_cmd.py | 243 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_review_cmd.py | 244 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_review_decisions.py | 478 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_review_decisions.py | 479 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_rules_cmd.py | 148 | 23 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_rules_cmd.py | 158 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_rules_cmd.py | 159 | 14 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sarif_exporter.py | 307 | 47 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sarif_exporter.py | 324 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sarif_exporter.py | 374 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sarif_exporter.py | 393 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sarif_exporter.py | 407 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sarif_exporter.py | 428 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sca_cache.py | 85 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sca_lockfiles.py | 182 | 15 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sca_pnpm.py | 451 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_sca_shrinkwrap.py | 188 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_secrets.py | 1788 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_secrets.py | 1798 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_secrets.py | 1876 | 56 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_secrets.py | 1877 | 54 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_secrets.py | 1948 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_secrets.py | 1959 | 48 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_secrets.py | 2356 | 53 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_secrets.py | 2360 | 47 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_secrets.py | 2365 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_security_taskflow.py | 229 | 59 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_security_taskflow.py | 229 | 67 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| suspicious_shell | test/test_shell_security.py | 312 | 1 | Suspicious shell command patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 272 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 273 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 273 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 284 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 285 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 286 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 287 | 21 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 287 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 293 | 54 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 311 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 312 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_trivy_image.py | 391 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_trivy_image.py | 391 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_ts_e003_entrypoints.py | 945 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 400 | 14 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 432 | 23 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 442 | 23 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 463 | 23 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 2384 | 39 | Potential data exfiltration patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 333 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 333 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 333 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 333 | 65 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 333 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 333 | 81 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 345 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 345 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 345 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 345 | 65 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 345 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 345 | 81 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 361 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 361 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 361 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 361 | 65 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 361 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 361 | 81 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 385 | 63 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 385 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 400 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 400 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 400 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 400 | 65 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 400 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 400 | 81 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 475 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_inline_ignores.py | 475 | 59 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_typescript_quality_signals.py | 977 | 21 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | test/test_typescript_quality_signals.py | 1073 | 23 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | test/test_upload_contract_fixtures.py | 107 | 42 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_contract_fixtures.py | 107 | 49 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_contract_fixtures.py | 351 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_contract_fixtures.py | 351 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 204 | 16 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 231 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 246 | 18 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 246 | 22 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 249 | 11 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 249 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 250 | 11 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 250 | 29 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 251 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 251 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 261 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 262 | 48 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 276 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 276 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 277 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 277 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 278 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 291 | 11 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 484 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 1008 | 33 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_upload_reliability.py | 1010 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_uv_lockfile.py | 327 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_render.py | 157 | 39 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_render.py | 157 | 65 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_render.py | 159 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_render.py | 165 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_render.py | 165 | 41 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_render.py | 166 | 42 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_verdict.py | 253 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_verdict.py | 253 | 62 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_verdict.py | 253 | 69 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_verify_verdict.py | 253 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_yarn_lockfile.py | 403 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |

### Medium (1128)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| template_injection | .github/workflows/corpus.yml | 59 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/corpus.yml | 60 | 54 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-plus-claude-security.yml | 59 | 81 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-plus-claude-security.yml | 61 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-plus-claude-security.yml | 82 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-plus-claude-security.yml | 120 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-tokenless-ci.yml | 44 | 90 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/liveness-primer.yml | 14 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/liveness-primer.yml | 27 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/liveness-primer.yml | 35 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/liveness-primer.yml | 41 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/liveness-primer.yml | 68 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/liveness-primer.yml | 100 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/liveness-primer.yml | 100 | 56 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/liveness-primer.yml | 102 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/liveness-primer.yml | 103 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/parity.yml | 48 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/pr-title.yml | 17 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 24 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 35 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 42 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 67 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 74 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 75 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 121 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 246 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 251 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 264 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 265 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 271 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 318 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 319 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 324 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 325 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 326 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 327 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 359 | 73 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 360 | 76 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/quality-benchmark.yml | 49 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/quality-benchmark.yml | 50 | 55 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 24 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 25 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 30 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 31 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 44 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 51 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 52 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 98 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 106 | 10 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 113 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 115 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/repo-map-pages.yml | 55 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 44 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 79 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 80 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 81 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 82 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 103 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 103 | 52 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 104 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 134 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 141 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 197 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 205 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 206 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 212 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 224 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 39 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 74 | 57 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 84 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 88 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 91 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 104 | 56 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 118 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 179 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 180 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 206 | 10 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 211 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 212 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 213 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 214 | 40 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 215 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 216 | 43 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/benchmarks/agent-pr-bench/harness/tools.py | 198 | 22 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/benchmarks/agent-pr-bench/harness/tools.py | 372 | 18 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/benchmarks/agent-pr-bench/harness/tools.py | 383 | 14 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| python.lang.security.deserialization.pickle.avoid-pickle | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/flask_pickle_deserialization/app.py | 15 | 15 | Avoid using `pickle`, which is known to lead to code execution vulnerabilities. When unpickling, the serialized data could be manipulated to run arbitrary code. Instead, consider serializing the relevant data as JSON or a similar text-based serialization format. | semgrep | deterministic |
| python.flask.security.audit.render-template-string.render-template-string | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/flask_reflected_xss/app.py | 12 | 12 | Found a template created with string formatting. This is susceptible to server-side template injection and cross-site scripting attacks. | semgrep | deterministic |
| python.flask.security.audit.render-template-string.render-template-string | /tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/flask_reflected_xss/app.py | 19 | 12 | Found a template created with string formatting. This is susceptible to server-side template injection and cross-site scripting attacks. | semgrep | deterministic |
| python.lang.security.deserialization.pickle.avoid-pickle | /tmp/safeanalyze-skylos/benchmarks/ai_code_defects/fixtures/repo_level_security_regression/app.py | 12 | 12 | Avoid using `pickle`, which is known to lead to code execution vulnerabilities. When unpickling, the serialized data could be manipulated to run arbitrary code. Instead, consider serializing the relevant data as JSON or a similar text-based serialization format. | semgrep | deterministic |
| python.django.security.audit.raw-query.avoid-raw-sql | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 93 | 5 | Detected the use of 'RawSQL' or 'raw' indicating the execution of a non-parameterized SQL query. This could lead to a SQL injection and therefore protected information could be leaked. Instead, use Django ORM and parameterized queries before raw SQL. An example of using the Django ORM is: `People.objects.get(name='Bob')` | semgrep | deterministic |
| python.django.security.injection.code.user-eval.user-eval | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 145 | 5 | Found user data in a call to 'eval'. This is extremely dangerous because it can enable an attacker to execute arbitrary remote code on the system. Instead, refactor your code to not use 'eval' and instead use a safe library for the specific functionality you need. | semgrep | deterministic |
| python.lang.security.audit.eval-detected.eval-detected | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 146 | 5 | Detected the use of eval(). eval() can be dangerous if used to evaluate dynamic content. If this content can be input from outside the program, this may be a code injection vulnerability. Ensure evaluated content is not definable by external sources. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 162 | 5 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| java.lang.security.audit.tainted-ldapi-from-http-request.tainted-ldapi-from-http-request | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/java_ldap_xpath_injection/App.java | 9 | 5 | Detected input from a HTTPServletRequest going into an LDAP query. This could lead to LDAP injection if the input is not properly sanitized, which could result in attackers modifying objects in the LDAP tree structure. Ensure data passed to an LDAP query is not controllable or properly sanitize the data. | semgrep | deterministic |
| java.lang.security.audit.tainted-xpath-from-http-request.tainted-xpath-from-http-request | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/java_ldap_xpath_injection/App.java | 13 | 12 | Detected input from a HTTPServletRequest going into a XPath evaluate or compile command. This could lead to xpath injection if variables passed into the evaluate or compile commands are not properly sanitized. Xpath injection could lead to unauthorized access to sensitive information in XML documents. Instead, thoroughly sanitize user input or use parameterized xpath queries if you can. | semgrep | deterministic |
| java.lang.security.audit.unvalidated-redirect.unvalidated-redirect | /tmp/safeanalyze-skylos/benchmarks/security/fixtures/java_open_redirect_protocol_relative/App.java | 5 | 3 | Application redirects to a destination URL specified by a user-supplied parameter that is not validated. This could direct users to malicious locations. Consider using an allowlist to validate URLs. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/api/_source_revision.py | 83 | 23 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/api/_source_revision.py | 105 | 23 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/skylos/commands/rules_cmd.py | 196 | 14 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/skylos/rules/ai_defect/dependency_hallucination.py | 2089 | 18 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/skylos/rules/ai_defect/manifest_dependency_hallucination.py | 2209 | 14 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/rules/quality/clones.py | 328 | 9 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/rules/quality/clones.py | 333 | 9 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/rules/quality/clones.py | 338 | 9 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/rules/quality/clones.py | 357 | 13 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/skylos/rules/sca/licenses.py | 604 | 10 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| template_injection | action.yml | 69 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 72 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 75 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 83 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 90 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 109 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 126 | 36 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 133 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 134 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 135 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 136 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 218 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 219 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 220 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 221 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 222 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 287 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 293 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 294 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 295 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 296 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 297 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 298 | 37 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 348 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 349 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 392 | 16 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | benchmarks/agent-pr-bench/bench.py | 21 | 1 | Imports of known suspicious packages | yara | deterministic |
| entropy_base64_blob | benchmarks/agent-pr-bench/corpus/real_commits.json | 84 | 13 | entropy=4.61 len=66 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | benchmarks/agent-pr-bench/corpus/real_commits.json | 576 | 13 | entropy=4.60 len=72 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | benchmarks/agent-pr-bench/corpus/real_commits.json | 630 | 13 | entropy=4.70 len=71 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | benchmarks/agent-pr-bench/corpus/real_commits.json | 930 | 13 | entropy=4.74 len=76 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | benchmarks/agent-pr-bench/corpus/real_commits.json | 1206 | 13 | entropy=4.55 len=66 type=base64_blob | entropy | heuristic |
| credential_hardcode | benchmarks/agent-pr-bench/corpus/seeded/cases/ex-09-secret-jwt.toml | 10 | 18 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | benchmarks/agent-pr-bench/corpus/seeded/cases/ex-09-secret-jwt.toml | 10 | 26 | entropy=5.17 len=40 type=base64_blob | entropy | heuristic |
| credential_hardcode | benchmarks/agent-pr-bench/corpus/seeded/cases/ex-09-secret-jwt.toml | 41 | 28 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | benchmarks/agent-pr-bench/corpus/seeded/cases/fa-11-cmdi.toml | 12 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | benchmarks/agent-pr-bench/corpus/seeded/cases/mb-14-secret-token.toml | 15 | 8 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | benchmarks/agent-pr-bench/corpus/seeded/cases/mb-c3-clean-subprocess.toml | 16 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/agent-pr-bench/harness/corpus.py | 19 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | benchmarks/agent-pr-bench/harness/corpus.py | 146 | 57 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | benchmarks/agent-pr-bench/harness/tools.py | 20 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | benchmarks/agent-pr-bench/results/findings.jsonl | 97 | 193 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | benchmarks/agent-pr-bench/results/workspaces.json | 2607 | 10 | entropy=4.61 len=66 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | benchmarks/agent-pr-bench/results/workspaces.json | 7271 | 10 | entropy=4.60 len=72 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | benchmarks/agent-pr-bench/results/workspaces.json | 7500 | 10 | entropy=4.70 len=71 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | benchmarks/agent-pr-bench/results/workspaces.json | 9398 | 10 | entropy=4.74 len=76 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | benchmarks/agent-pr-bench/results/workspaces.json | 10370 | 10 | entropy=4.55 len=66 type=base64_blob | entropy | heuristic |
| suspicious_imports | benchmarks/agent-pr-bench/scripts/select_real_commits.py | 32 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/agent_review/fixtures/extreme_grounding_framework/services/hooks.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/agent_review/fixtures/extreme_grounding_framework/tools/admin.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/agent_review/fixtures/flask_getter_shell/app.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/agent_review/fixtures/flask_handler_security/app.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | benchmarks/agent_review/fixtures/flask_reflected_xss/app.py | 19 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | benchmarks/agent_review/fixtures/shell_hook_runner/hooks.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/agent_review/fixtures/static_blind_plugin_dispatch/plugins/audit.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/agent_review/fixtures/static_blind_plugin_dispatch/plugins/tools.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/dead_code/fixtures/extreme_grounding_framework/services/hooks.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/dead_code/fixtures/extreme_grounding_framework/tools/admin.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/dead_code/fixtures/static_blind_plugin_dispatch/plugins/audit.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/dead_code/fixtures/static_blind_plugin_dispatch/plugins/tools.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmarks/security/fixtures/subprocess_alias_shell/app.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | benchmarks/verify_benchmark/fixtures/openai_sdk_member_drift/summarizer.py | 4 | 17 | Potential hardcoded credentials | yara | deterministic |
| template_injection | dictionary.md | 490 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | docs/container-image-reports.md | 41 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | docs/done-gate.md | 120 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | docs/done-gate.md | 125 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | editors/vscode/out/ai.js | 550 | 118 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | editors/vscode/out/autoremediate.js | 151 | 157 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| suspicious_imports | editors/vscode/out/commandcenter.js | 39 | 25 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | editors/vscode/out/scanner.js | 43 | 25 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | editors/vscode/out/verifyCore.js | 8 | 25 | Imports of known suspicious packages | yara | deterministic |
| entropy_base64_blob | editors/vscode/package-lock.json | 24 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 31 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 41 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 54 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 69 | 19 | entropy=5.34 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 88 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 107 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 120 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 135 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 158 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 172 | 19 | entropy=5.52 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 185 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 195 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 209 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 224 | 19 | entropy=5.25 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 234 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 248 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 258 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 272 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 285 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 303 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 319 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 342 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 355 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 375 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 382 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 389 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 399 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 412 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 422 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 436 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 446 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 459 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 466 | 19 | entropy=5.53 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 489 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 499 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 506 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 519 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 526 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 533 | 19 | entropy=5.34 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 543 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 552 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 559 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 566 | 19 | entropy=5.23 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 572 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 587 | 19 | entropy=5.58 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 634 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 653 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 667 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 681 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 695 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 709 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 723 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 737 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 751 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 765 | 19 | entropy=5.34 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 779 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 789 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 806 | 19 | entropy=5.25 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 822 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 835 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 851 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 858 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 868 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 875 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 886 | 19 | entropy=5.58 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 896 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 918 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 934 | 19 | entropy=5.57 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 947 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 954 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 961 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 974 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 987 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1013 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1023 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1030 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1046 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1060 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1077 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1094 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1120 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1138 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1146 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1156 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1169 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1176 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1189 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1199 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1216 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1229 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1247 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1264 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1275 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1292 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1305 | 19 | entropy=5.59 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1318 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1328 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1339 | 19 | entropy=5.60 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1354 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1367 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1383 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1398 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1413 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1423 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1440 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1447 | 19 | entropy=5.29 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1461 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1472 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1485 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1498 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1508 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1518 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1531 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1547 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1558 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1565 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1582 | 19 | entropy=5.29 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1599 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1609 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1622 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1639 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1647 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1662 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1672 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1697 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1711 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1719 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1737 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1750 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1771 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1784 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1791 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1801 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1814 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1830 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1843 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1856 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1876 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1889 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1903 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1917 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1930 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1952 | 19 | entropy=5.34 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1962 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1975 | 19 | entropy=5.26 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1983 | 19 | entropy=5.52 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1991 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2007 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2017 | 19 | entropy=5.52 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2027 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2040 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2059 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2069 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2085 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2103 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2110 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2133 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2140 | 19 | entropy=5.26 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2153 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2160 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2173 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2196 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2208 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2219 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2232 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2242 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2262 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2269 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2276 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2283 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2290 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2297 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2304 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2311 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2318 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2325 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2338 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2366 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2376 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2383 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2393 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2407 | 19 | entropy=5.30 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2420 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2430 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2443 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2457 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2473 | 19 | entropy=5.23 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2484 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2494 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2502 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2509 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2516 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2524 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2538 | 19 | entropy=5.15 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2546 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2560 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2575 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2588 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2595 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2608 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2621 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2632 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2651 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2664 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2682 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2692 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2702 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2715 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2729 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2742 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2755 | 19 | entropy=5.52 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2772 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2782 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2795 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2802 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2809 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2822 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2832 | 19 | entropy=5.67 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2861 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2873 | 19 | entropy=5.62 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2883 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2899 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2920 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2937 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2950 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2963 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2983 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2996 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3012 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3022 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3033 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3046 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3070 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3091 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3098 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3108 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3130 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3143 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3163 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3180 | 19 | entropy=5.53 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3199 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3219 | 19 | entropy=5.62 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3241 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3268 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3281 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3299 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3310 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3317 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3328 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3335 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3346 | 19 | entropy=5.52 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3361 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3371 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3384 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3400 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3411 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3421 | 19 | entropy=5.27 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3434 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3451 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3468 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3478 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3491 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3505 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3523 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3540 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3547 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3563 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3573 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3586 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3593 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3603 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3617 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3630 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3642 | 19 | entropy=5.53 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3655 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3662 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3669 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3679 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3685 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3698 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3708 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3715 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3723 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3734 | 19 | entropy=5.57 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3747 | 19 | entropy=5.57 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3761 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3771 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3779 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3795 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3809 | 19 | entropy=5.65 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3819 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3826 | 19 | entropy=5.64 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3840 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| output_constraint | editors/vscode/src/ai.ts | 603 | 116 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | editors/vscode/src/autoremediate.ts | 159 | 153 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| suspicious_imports | scripts/compare_codex_skylos_agent_review.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | scripts/compare_codex_skylos_demo_deadcode.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | scripts/compare_codex_skylos_quality.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/agents/evaluation/_openai_transport.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/analyzer.py | 10 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/api/__init__.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/api/_ai_detection.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/api/_source_revision.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/benchmarks/ai_code_defects.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/benchmarks/dead_code.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/benchmarks/security.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/benchmarks/verify_benchmark_runner.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/cicd/review.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/cicd/review.py | 916 | 97 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 68 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 69 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 123 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 143 | 50 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 156 | 45 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 244 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 280 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/cli.py | 49 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/cli.py | 4116 | 17 | Imports of known suspicious packages | yara | deterministic |
| output_constraint | skylos/cli_core/main_parser.py | 464 | 39 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| suspicious_imports | skylos/cloud/login.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/cloud/login.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/cloud/sync.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/cloud/sync_setup.py | 96 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cloud/sync_setup.py | 97 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/commands/agent_verify_cmd.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/commands/cicd_cmd.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/commands/compare_cmd.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/commands/hook_cmd.py | 719 | 5 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/commands/image_cmd.py | 9 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/commands/install_hooks_cmd.py | 15 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/commands/lint_cmd.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| account_access_request | skylos/commands/scan_cmd.py | 1493 | 61 | Natural-language request to access a user account or service | yara | deterministic |
| suspicious_imports | skylos/core/baseline_source.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/core/baseline_source.py | 74 | 62 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/core/cli_shared.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/core/file_discovery.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/core/gatekeeper.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/core/git_context.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/core/grep_verify_common.py | 12 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/core/grep_verify_language_strategies.py | 168 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/core/review_decisions.py | 9 | 1 | Imports of known suspicious packages | yara | deterministic |
| output_constraint | skylos/debt/advisor.py | 60 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| suspicious_imports | skylos/done/base.py | 14 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/done/base.py | 229 | 74 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/done/runner.py | 27 | 1 | Imports of known suspicious packages | yara | deterministic |
| output_constraint | skylos/done/runner.py | 547 | 35 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| credential_hardcode | skylos/engines/go/internal/analyzer/analyzer_symlink_test.go | 16 | 7 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | skylos/engines/go/internal/analyzer/analyzer_symlink_test.go | 26 | 7 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | skylos/engines/go/internal/symbols/symbols_symlink_test.go | 28 | 14 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | skylos/engines/go_runner.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/integrations/trivy_image.py | 30 | 76 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/agents.py | 498 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/agents.py | 515 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/cleanup_orchestrator.py | 193 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/cleanup_orchestrator.py | 233 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| credential_hardcode | skylos/llm/context.py | 664 | 8 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | skylos/llm/executor.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| output_constraint | skylos/llm/harness/ai_defect_challenge.py | 308 | 62 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 12 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 144 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 149 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 154 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 184 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 192 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 235 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 242 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 264 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 267 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 313 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 313 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 320 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 363 | 26 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 377 | 3 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 381 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/security_verifier.py | 600 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/security_verifier.py | 613 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/verification/llm.py | 54 | 39 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/llm/verification/llm.py | 55 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 82 | 12 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 82 | 53 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 115 | 12 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 115 | 53 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 173 | 21 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 173 | 70 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 187 | 21 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 187 | 70 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 188 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/verification/llm.py | 384 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 413 | 12 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verification/llm.py | 413 | 61 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| suspicious_imports | skylos/llm/verify_orchestrator.py | 847 | 5 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/llm/verify_orchestrator.py | 1074 | 108 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/misc/update_mapping.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| output_constraint | skylos/pipeline.py | 74 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| suspicious_imports | skylos/preflight/inspector.py | 11 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/remediation/fixgen.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/reporting/provenance.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/rules/ai_defect/dependency_bump_scan.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/rules/ai_defect/dependency_bump_scan.py | 136 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/rules/config/deployment/exposure.py | 561 | 14 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/rules/quality/code_health.py | 32 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/rules/secrets.py | 111 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/rules/secrets.py | 114 | 48 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/rules/secrets.py | 231 | 39 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/rules/secrets.py | 269 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/rules/secrets.py | 272 | 45 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| entropy_base64_blob | skylos/rules/secrets.py | 305 | 8 | entropy=4.66 len=40 type=base64_blob | entropy | heuristic |
| suspicious_imports | skylos/ui/tui.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/verification/context.py | 170 | 73 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/verification/refactor.py | 11 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/verification/refactor.py | 352 | 74 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/verify_change.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/verify_change.py | 210 | 63 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/visitors/languages/csharp/danger.py | 484 | 46 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/conftest.py | 260 | 13 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/fixtures/app.go | 13 | 7 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/fixtures/upload_contract/long-message.payload.json | 34 | 30 | Imports of known suspicious packages | yara | deterministic |
| entropy_base64_blob | test/fixtures/verdicts/verdict_failed.json | 9 | 15 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | test/fixtures/verdicts/verdict_passed.json | 9 | 15 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| suspicious_imports | test/test_action_image_scan.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_action_image_scan.py | 135 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_action_image_scan.py | 136 | 36 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_action_image_scan.py | 137 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_action_image_scan.py | 138 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_action_image_scan.py | 169 | 48 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_action_sarif_upload.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_action_sarif_upload.py | 107 | 52 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_action_sarif_upload.py | 118 | 44 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_action_sca.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_agent_behavior_cli.py | 932 | 5 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_agent_behavior_contract.py | 243 | 8 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_agent_behavior_contract.py | 284 | 8 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_agent_behavior_openai.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_agent_behavior_openai.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_agent_behavior_openai.py | 629 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_agent_behavior_openai.py | 672 | 5 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_agent_bench_python_detections.py | 107 | 48 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_agent_center.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_agent_center_review_memory.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_agent_integration.py | 1403 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_agent_integration.py | 1486 | 10 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_agent_integration.py | 1962 | 50 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_agent_llm_review_context.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_agent_llm_review_context.py | 89 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_agent_llm_review_context.py | 183 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_agent_llm_review_context.py | 464 | 13 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_agent_review_projection.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_ai_pr_diff_rules.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_analysis_review_context.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_analyzer.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| entropy_base64_blob | test/test_analyzer.py | 5257 | 4 | entropy=5.28 len=66 type=base64_blob | entropy | heuristic |
| suspicious_imports | test/test_api.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_api.py | 737 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_api_signature_hallucination.py | 296 | 34 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_api_symbol_truth.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_async_blocking.py | 61 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_audit_candidates.py | 35 | 34 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_candidates.py | 83 | 34 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_candidates.py | 103 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_export.py | 163 | 9 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_export.py | 164 | 25 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_export.py | 275 | 9 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_export.py | 468 | 9 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_audit_investigator_reviewer_packs.py | 180 | 77 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_audit_investigator_tools.py | 126 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_processor.py | 614 | 22 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_processor.py | 666 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_revalidator.py | 394 | 22 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_baseline_source.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_behavior_baselines.py | 9 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_behavior_changes.py | 9 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_behavior_comparison.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_cicd_evidence.py | 188 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cicd_evidence.py | 190 | 46 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cicd_evidence.py | 205 | 14 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cicd_review.py | 1016 | 5 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | test/test_cicd_review.py | 1017 | 15 | entropy=4.55 len=52 type=base64_blob | entropy | heuristic |
| template_injection | test/test_cicd_workflow.py | 192 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_cicd_workflow.py | 193 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_circular_deps.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_circular_deps.py | 997 | 43 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_cleanup_orchestrator.py | 167 | 35 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cleanup_orchestrator.py | 252 | 35 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cleanup_orchestrator.py | 483 | 28 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cli.py | 1354 | 22 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_cli_image_scan.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cli_image_scan.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_cli_llm_provider.py | 205 | 51 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_cli_precommit.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_cli_precommit.py | 932 | 10 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cli_precommit.py | 952 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cli_precommit.py | 1251 | 35 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_cli_sca_baseline.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cli_terminal_probe.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cli_trivy_image.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cli_trivy_image.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cmd_injection.py | 21 | 13 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cmd_injection.py | 28 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cmd_injection.py | 37 | 13 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_code_health_diff.py | 11 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_container_python_runtime.py | 43 | 42 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_cpp_scanner.py | 197 | 5 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_csharp_danger_inline.py | 129 | 62 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_csharp_raw_lex.py | 18 | 44 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_csharp_raw_lex.py | 55 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_csharp_raw_lex.py | 60 | 71 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_csharp_raw_lex.py | 70 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 44 | 43 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 165 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 175 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 647 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 681 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 724 | 13 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dead_code_benchmark.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_debt.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_debt.py | 960 | 13 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_defend.py | 890 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_defend.py | 2048 | 5 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dependency_version_bump_cli.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dependency_version_bump_cli.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dependency_version_bump_integration.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_deployment_exposure.py | 849 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_deployment_exposure.py | 1068 | 46 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_deployment_exposure.py | 1853 | 40 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_deployment_exposure.py | 1881 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_deployment_exposure.py | 1898 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_deployment_exposure.py | 1913 | 11 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_diff_working_tree.py | 10 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_diff_working_tree.py | 40 | 41 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_discover.py | 60 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_discover.py | 99 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_discover.py | 750 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_django_static_entrypoints.py | 62 | 65 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_django_static_entrypoints.py | 100 | 64 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_django_static_entrypoints.py | 286 | 65 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_done_gate.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_done_gate.py | 651 | 39 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_done_gate.py | 673 | 39 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_done_gate.py | 800 | 39 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_done_receipt_binding.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_fast_parity.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_file_discovery.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_file_discovery.py | 32 | 24 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_file_discovery.py | 54 | 24 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_finding_verification.py | 332 | 57 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_fixgen.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_fixgen_multilang.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_gatekeeper.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_git_context.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_git_hook_regressions.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_github_actions_finding_locations.py | 140 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 40 | 37 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 47 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 48 | 12 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 49 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 50 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 63 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 90 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 97 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 98 | 46 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 135 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 136 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 137 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 138 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 139 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 140 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_runner_labels.py | 141 | 48 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 55 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 187 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 282 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 303 | 12 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 305 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 393 | 15 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 419 | 14 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 428 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 432 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 469 | 54 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 528 | 48 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 567 | 47 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 649 | 11 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 687 | 53 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 715 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 716 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 717 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 718 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 735 | 59 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 736 | 14 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 763 | 51 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 777 | 57 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_gitlab_delivery_receipt.py | 97 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_gitlab_report.py | 349 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_gitlab_upload_auth.py | 342 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_go_security.py | 74 | 32 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_grep_python_backend.py | 13 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_grep_verify.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_hallucination_dependency.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_hook_cmd.py | 23 | 5 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_hook_cmd.py | 104 | 6 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_hook_cmd.py | 279 | 48 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_hook_cmd.py | 462 | 7 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_hook_cmd.py | 952 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_hook_cmd.py | 960 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_hook_cmd.py | 1838 | 41 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_java_source_helper_limits.py | 52 | 52 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_java_source_helpers.py | 34 | 46 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | test/test_jev_cascade_benchmark.py | 585 | 5 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| credential_hardcode | test/test_jev_dataset_safety.py | 113 | 36 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_jev_dead_code_benchmark.py | 278 | 9 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_jev_run_safety.py | 53 | 9 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_jev_run_safety.py | 106 | 9 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_js_barrel_exports.py | 246 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_js_barrel_exports.py | 246 | 69 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_js_barrel_exports.py | 380 | 44 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_js_barrel_exports.py | 459 | 73 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_js_barrel_exports.py | 497 | 73 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_js_barrel_exports.py | 517 | 40 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_licenses.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_lint_cmd.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_liveness_primer_workflow.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_liveness_primer_workflow.py | 44 | 36 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_liveness_primer_workflow.py | 85 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_liveness_primer_workflow.py | 95 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_liveness_primer_workflow.py | 119 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_liveness_primer_workflow.py | 147 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_liveness_primer_workflow.py | 162 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_liveness_primer_workflow.py | 162 | 57 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_liveness_primer_workflow.py | 166 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_liveness_primer_workflow.py | 167 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_llm_analyzer.py | 342 | 14 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_llm_analyzer.py | 826 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_analyzer.py | 879 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_analyzer.py | 1139 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_finding_evidence.py | 68 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_finding_evidence.py | 147 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_finding_evidence.py | 287 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_graph.py | 7 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_graph.py | 26 | 10 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 497 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 635 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 669 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 736 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 787 | 59 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 815 | 46 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_logic.py | 386 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_logic.py | 485 | 45 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_login.py | 143 | 9 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_login.py | 161 | 9 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 174 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 203 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 233 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 244 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 254 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 279 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 306 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 324 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 352 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 556 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 578 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 601 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 624 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_auth.py | 660 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_security.py | 132 | 39 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_security.py | 144 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_mcp_security.py | 246 | 47 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_mcp_summary.py | 586 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_method_reachability.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_module_facts_index.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_output_validation_flow.py | 235 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_output_validation_flow.py | 268 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_package_data.py | 11 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_pnpm_lockfile.py | 644 | 39 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_pnpm_regressions.py | 181 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_pnpm_regressions.py | 185 | 56 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_pnpm_regressions.py | 190 | 39 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_pnpm_regressions.py | 219 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_pnpm_regressions.py | 223 | 70 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_preflight.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_preflight.py | 9 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_preflight_inspector.py | 9 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_provenance.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_provenance_attribution.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_python_reachability.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_quality_rules.py | 441 | 1 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_quality_rules.py | 475 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_quality_rules.py | 490 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_quality_rules.py | 684 | 9 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_reachability_integration_qa.py | 10 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_refactor_verify.py | 10 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_repo_map.py | 103 | 72 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_review_cmd.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_review_decisions.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_review_memory_adversarial.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_review_mutation_guards.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_s102.py | 73 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_s102.py | 83 | 11 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_sanitizers.py | 58 | 13 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_sarif_exporter.py | 433 | 5 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | test/test_sarif_exporter.py | 435 | 15 | entropy=4.55 len=52 type=base64_blob | entropy | heuristic |
| suspicious_imports | test/test_sarif_exporter.py | 696 | 21 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_sarif_exporter.py | 706 | 21 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_sbom_shrinkwrap.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_sbom_shrinkwrap.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| entropy_base64_blob | test/test_secrets.py | 15 | 4 | entropy=5.28 len=66 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | test/test_secrets.py | 19 | 13 | entropy=4.70 len=44 type=base64_blob | entropy | heuristic |
| credential_hardcode | test/test_secrets.py | 29 | 9 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 30 | 12 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 83 | 19 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 117 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 141 | 19 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 201 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 230 | 26 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 261 | 5 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | test/test_secrets.py | 274 | 8 | entropy=5.95 len=62 type=base64_blob | entropy | heuristic |
| credential_hardcode | test/test_secrets.py | 288 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 312 | 5 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | test/test_secrets.py | 312 | 12 | entropy=5.95 len=62 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | test/test_secrets.py | 328 | 9 | entropy=5.70 len=52 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | test/test_secrets.py | 332 | 12 | entropy=5.95 len=62 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | test/test_secrets.py | 337 | 12 | entropy=5.95 len=62 type=base64_blob | entropy | heuristic |
| credential_hardcode | test/test_secrets.py | 426 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 516 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 739 | 5 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_secrets.py | 740 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 834 | 5 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_secrets.py | 835 | 36 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_secrets.py | 860 | 40 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_secrets.py | 870 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_secrets.py | 891 | 42 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| entropy_base64_blob | test/test_secrets.py | 1232 | 12 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | test/test_secrets.py | 1237 | 12 | entropy=5.18 len=60 type=base64_blob | entropy | heuristic |
| template_injection | test/test_secrets.py | 1500 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_secrets.py | 1517 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_secrets.py | 1552 | 74 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_secrets.py | 1554 | 75 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_secrets.py | 1573 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_secrets.py | 1594 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_secrets.py | 1821 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 2380 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 2391 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 2438 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 2452 | 5 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | test/test_secrets.py | 2467 | 8 | entropy=5.29 len=63 type=base64_blob | entropy | heuristic |
| credential_hardcode | test/test_secrets.py | 2519 | 19 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 2539 | 10 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | test/test_secrets.py | 2540 | 8 | entropy=4.66 len=40 type=base64_blob | entropy | heuristic |
| template_injection | test/test_secrets_credentials.py | 44 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_secrets_credentials.py | 72 | 29 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_credentials.py | 74 | 41 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_credentials.py | 78 | 35 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_credentials.py | 81 | 27 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_secrets_credentials.py | 95 | 34 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_secrets_credentials.py | 130 | 35 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_credentials.py | 143 | 26 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_nonpy.py | 55 | 40 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_nonpy.py | 56 | 34 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_nonpy.py | 86 | 40 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_nonpy.py | 105 | 46 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_nonpy.py | 145 | 40 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | test/test_secrets_nonpy.py | 164 | 8 | entropy=5.19 len=64 type=base64_blob | entropy | heuristic |
| credential_hardcode | test/test_secrets_nonpy.py | 184 | 23 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_nonpy.py | 200 | 16 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_security_benchmark.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_security_benchmark.py | 93 | 10 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_security_contracts.py | 190 | 10 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_ssrf.py | 198 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_ssrf.py | 405 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_ssrf.py | 418 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_sync.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_test_impact_ai_defect.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_trivy_image.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_trivy_image.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_ts_danger_d281_object_semantics_gaps.py | 240 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_object_semantics_gaps.py | 319 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_object_semantics_gaps.py | 388 | 14 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_object_semantics_gaps.py | 389 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_object_semantics_gaps.py | 390 | 15 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_object_semantics_gaps.py | 416 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_object_semantics_gaps.py | 417 | 15 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_object_semantics_gaps.py | 420 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_object_semantics_gaps.py | 444 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_promise_temporal.py | 268 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_redteam_regressions.py | 293 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_redteam_regressions.py | 336 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_redteam_regressions.py | 390 | 67 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_redteam_regressions.py | 414 | 67 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_redteam_regressions.py | 439 | 67 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_redteam_regressions.py | 461 | 67 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_redteam_regressions.py | 481 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_redteam_regressions.py | 560 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_d281_redteam_regressions.py | 561 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_release_blockers.py | 1294 | 40 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_release_blockers.py | 1318 | 66 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_danger_release_blockers.py | 1382 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_e003_entrypoints.py | 323 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_e003_entrypoints.py | 327 | 37 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_e003_entrypoints.py | 624 | 61 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_e003_entrypoints.py | 727 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_e003_entrypoints.py | 727 | 70 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_e003_entrypoints.py | 747 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_e003_entrypoints.py | 851 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_e003_entrypoints.py | 949 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_e003_entrypoints.py | 988 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_esbuild_provenance.py | 110 | 68 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_ts_exports.py | 1177 | 8 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_typescript_expanded.py | 421 | 25 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_typescript_expanded.py | 1320 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_typescript_expanded.py | 2225 | 23 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_typescript_expanded.py | 2279 | 36 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_expanded.py | 2280 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_typescript_expanded.py | 2496 | 23 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_typescript_expanded.py | 2716 | 66 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_expanded.py | 2735 | 66 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_framework.py | 395 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_framework.py | 396 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_framework.py | 397 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_framework.py | 436 | 43 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_framework.py | 443 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_typescript_inline_ignores.py | 280 | 16 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_typescript_inline_ignores.py | 302 | 16 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_typescript_inline_ignores.py | 313 | 16 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_typescript_inline_ignores.py | 350 | 16 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_typescript_inline_ignores.py | 391 | 18 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_typescript_inline_ignores.py | 408 | 18 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_typescript_quality_signals.py | 1053 | 44 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_quality_signals.py | 1117 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 520 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 1143 | 8 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 1156 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 1600 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 1912 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 1913 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 1959 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 2845 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_typescript_security_proof.py | 2874 | 28 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 2965 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 2970 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 3695 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_security_proof.py | 3702 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_test_callback_length.py | 40 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_test_callback_length.py | 48 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_test_callback_length.py | 57 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_test_callback_length.py | 57 | 52 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_test_callback_length.py | 65 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_test_callback_length.py | 112 | 46 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_test_callback_length.py | 121 | 43 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_test_callback_length.py | 146 | 39 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_unsafe_export_regressions.py | 397 | 43 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_unsafe_export_regressions.py | 561 | 62 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_unsafe_export_regressions.py | 586 | 62 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_unsafe_export_regressions.py | 758 | 43 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_upload_contract_fixtures.py | 206 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_upload_reliability.py | 15 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_upload_reliability.py | 59 | 1 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_upload_revision_attribution.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_uv_lockfile.py | 59 | 50 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_uv_lockfile.py | 389 | 48 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_uv_lockfile.py | 389 | 73 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_verify_change.py | 1363 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_verify_change.py | 1394 | 11 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_verify_change_dependency_bump.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_verify_change_dependency_bump.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_verify_cmd.py | 492 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_verify_cmd.py | 520 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_verify_orchestrator.py | 494 | 10 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_verify_orchestrator.py | 1211 | 9 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_verify_verdict.py | 199 | 5 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_vibe_rules.py | 473 | 17 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 896 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1537 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1546 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1555 | 20 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1573 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1601 | 29 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1620 | 24 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1629 | 17 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_vibe_rules.py | 1842 | 9 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_xss_flow.py | 87 | 12 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| entropy_base64_blob | uv.lock | 17 | 12 | entropy=4.51 len=96 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 62 | 12 | entropy=4.52 len=87 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 78 | 12 | entropy=4.50 len=87 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 500 | 12 | entropy=4.52 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 502 | 12 | entropy=4.51 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 507 | 12 | entropy=4.51 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 509 | 12 | entropy=4.54 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 510 | 12 | entropy=4.53 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 516 | 12 | entropy=4.51 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 518 | 12 | entropy=4.50 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 636 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 645 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 646 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 647 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 649 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 650 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 651 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 653 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 655 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 657 | 12 | entropy=4.55 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 660 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 661 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 662 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 665 | 12 | entropy=4.57 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 671 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 672 | 12 | entropy=4.57 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 673 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 676 | 12 | entropy=4.58 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 678 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 680 | 12 | entropy=4.56 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 684 | 12 | entropy=4.55 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 685 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 686 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 689 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 690 | 12 | entropy=4.55 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 693 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 694 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 697 | 12 | entropy=4.55 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 698 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 699 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 702 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 703 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 704 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 705 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 707 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 708 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 709 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 710 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 712 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 714 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 716 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 719 | 12 | entropy=4.56 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 726 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 727 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 728 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 731 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 733 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 735 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 738 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 739 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 742 | 12 | entropy=4.56 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 744 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 893 | 16 | entropy=4.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 1567 | 12 | entropy=4.55 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 1714 | 12 | entropy=4.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 1743 | 12 | entropy=4.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 1933 | 12 | entropy=4.51 len=87 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 2387 | 16 | entropy=4.56 len=91 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 2389 | 12 | entropy=4.53 len=91 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 2618 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 2620 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |

### Low (386)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| entropy_high_entropy | .github/workflows/publish.yml | 132 | 14 | entropy=5.03 len=85 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | .github/workflows/quality-benchmark.yml | 55 | 106 | entropy=5.16 len=45 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | .github/workflows/release-please.yml | 66 | 14 | entropy=5.20 len=84 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmarks/agent-pr-bench/corpus/seeded/cases/ex-09-secret-jwt.toml | 10 | 26 | entropy=5.17 len=40 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmarks/agent-pr-bench/labels/findings.json | 4187 | 13 | entropy=5.15 len=158 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/out/dashboard.js | 207 | 35 | entropy=5.01 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/out/dashboard.js | 343 | 35 | entropy=5.03 len=169 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 24 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 31 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 41 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 54 | 19 | entropy=5.62 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 69 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 88 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 107 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 120 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 135 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 158 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 172 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 185 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 195 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 209 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 224 | 19 | entropy=5.32 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 234 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 248 | 19 | entropy=5.32 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 258 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 272 | 19 | entropy=5.31 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 285 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 303 | 19 | entropy=5.35 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 319 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 342 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 355 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 375 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 382 | 19 | entropy=5.34 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 389 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 399 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 412 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 422 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 436 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 446 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 459 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 466 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 489 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 499 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 506 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 519 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 526 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 533 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 543 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 552 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 559 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 566 | 19 | entropy=5.30 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 572 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 587 | 19 | entropy=5.66 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 634 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 653 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 667 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 681 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 695 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 709 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 723 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 737 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 751 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 765 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 779 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 789 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 806 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 822 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 835 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 851 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 858 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 868 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 875 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 886 | 19 | entropy=5.66 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 896 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 918 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 934 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 947 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 954 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 961 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 974 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 987 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1013 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1023 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1030 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1046 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1060 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1077 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1094 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1120 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1138 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1146 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1156 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1169 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1176 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1189 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1199 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1216 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1247 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1264 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1275 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1292 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1305 | 19 | entropy=5.62 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1318 | 19 | entropy=5.38 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1328 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1339 | 19 | entropy=5.68 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1354 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1367 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1383 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1398 | 19 | entropy=5.66 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1413 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1423 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1440 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1447 | 19 | entropy=5.35 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1461 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1472 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1485 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1498 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1508 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1518 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1531 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1547 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1558 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1565 | 19 | entropy=5.62 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1582 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1599 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1609 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1622 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1639 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1647 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1662 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1672 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1697 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1711 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1719 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1737 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1750 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1771 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1784 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1791 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1801 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1830 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1843 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1856 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1876 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1889 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1903 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1917 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1930 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1952 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1962 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1975 | 19 | entropy=5.36 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1983 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1991 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2007 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2017 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2027 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2040 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2059 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2069 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2085 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2103 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2110 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2133 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2140 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2153 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2160 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2173 | 19 | entropy=5.63 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2196 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2208 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2219 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2232 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2242 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2262 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2269 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2276 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2283 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2290 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2297 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2304 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2311 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2318 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2325 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2338 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2366 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2376 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2383 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2393 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2407 | 19 | entropy=5.37 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2420 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2430 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2443 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2457 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2473 | 19 | entropy=5.28 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2484 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2494 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2502 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2509 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2516 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2524 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2538 | 19 | entropy=5.22 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2546 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2560 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2575 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2588 | 19 | entropy=5.38 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2595 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2608 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2621 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2632 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2651 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2664 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2682 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2692 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2702 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2715 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2729 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2742 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2755 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2772 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2782 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2795 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2802 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2809 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2822 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2832 | 19 | entropy=5.69 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2861 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2873 | 19 | entropy=5.65 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2883 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2899 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2920 | 19 | entropy=5.36 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2937 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2950 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2963 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2983 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2996 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3012 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3022 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3033 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3046 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3070 | 19 | entropy=5.38 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3091 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3098 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3108 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3130 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3143 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3163 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3180 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3199 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3219 | 19 | entropy=5.64 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3241 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3268 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3281 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3299 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3310 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3317 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3328 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3346 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3361 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3371 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3384 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3400 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3411 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3421 | 19 | entropy=5.34 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3434 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3451 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3468 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3478 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3491 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3505 | 19 | entropy=5.67 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3523 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3540 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3547 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3563 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3573 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3586 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3593 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3603 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3617 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3630 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3642 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3655 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3662 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3669 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3679 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3685 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3698 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3708 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3715 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3723 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3734 | 19 | entropy=5.64 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3747 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3761 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3771 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3779 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3795 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3809 | 19 | entropy=5.66 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3819 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3826 | 19 | entropy=5.70 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3840 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/src/dashboard.ts | 205 | 35 | entropy=5.01 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/analysis/file_processing.py | 280 | 13 | entropy=5.05 len=68 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/analyzer.py | 3782 | 21 | entropy=5.01 len=67 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/api/__init__.py | 3269 | 9 | entropy=5.13 len=116 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/commands/hook_cmd.py | 1261 | 9 | entropy=5.09 len=75 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/integrations/shadow.py | 143 | 13 | entropy=5.01 len=67 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/preflight/evaluator.py | 16 | 23 | entropy=5.03 len=54 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/preflight/inspector.py | 61 | 5 | entropy=5.13 len=68 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/preflight/inspector.py | 66 | 5 | entropy=5.29 len=80 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/rules/config/gpu/compatibility.py | 1535 | 9 | entropy=5.08 len=70 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/rules/secrets.py | 371 | 5 | entropy=5.21 len=80 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/rules/secrets.py | 375 | 5 | entropy=5.01 len=58 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/rules/secrets.py | 379 | 5 | entropy=5.27 len=68 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/rules/secrets.py | 382 | 5 | entropy=5.26 len=70 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/visitors/languages/csharp/danger.py | 51 | 5 | entropy=5.00 len=55 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/visitors/languages/typescript/danger.py | 191 | 4 | entropy=6.07 len=67 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos_mcp/server.py | 947 | 5 | entropy=5.03 len=63 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/fixtures/verdicts/verdict_failed.json | 9 | 15 | entropy=5.50 len=88 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/fixtures/verdicts/verdict_keys.json | 6 | 24 | entropy=5.03 len=116 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/fixtures/verdicts/verdict_passed.json | 9 | 15 | entropy=5.56 len=88 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_agent_bench_python_detections.py | 262 | 8 | entropy=5.05 len=80 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_agent_bench_python_detections.py | 277 | 8 | entropy=5.27 len=77 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_agent_bench_python_detections.py | 281 | 8 | entropy=5.34 len=87 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_agent_bench_python_detections.py | 301 | 27 | entropy=5.15 len=77 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_agent_bench_python_detections.py | 302 | 28 | entropy=5.09 len=58 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_analyzer.py | 5257 | 4 | entropy=5.39 len=73 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_audit_investigator_tools.py | 126 | 12 | entropy=5.04 len=33 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_evidence.py | 15 | 29 | entropy=5.17 len=36 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_evidence.py | 19 | 25 | entropy=5.00 len=32 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_evidence.py | 202 | 8 | entropy=5.19 len=41 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_evidence.py | 203 | 8 | entropy=5.23 len=42 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_evidence.py | 204 | 8 | entropy=5.20 len=50 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_evidence.py | 205 | 8 | entropy=5.06 len=39 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_evidence.py | 269 | 9 | entropy=5.39 len=61 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_review.py | 41 | 29 | entropy=5.17 len=36 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_container_python_runtime.py | 43 | 11 | entropy=5.13 len=57 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_deep_audit_logic_benchmark_output.py | 110 | 8 | entropy=5.12 len=56 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_dependency_version_bump.py | 155 | 8 | entropy=5.15 len=66 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_dependency_version_bump.py | 156 | 8 | entropy=5.02 len=70 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_gpu_compatibility.py | 251 | 8 | entropy=5.11 len=74 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_hallucination_dependency.py | 844 | 9 | entropy=5.10 len=57 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_java_security.py | 90 | 6 | entropy=6.02 len=65 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_mcp_rules.py | 287 | 30 | entropy=5.22 len=40 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_s102.py | 39 | 9 | entropy=5.44 len=55 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_s102.py | 83 | 9 | entropy=5.43 len=52 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_s102.py | 832 | 12 | entropy=5.45 len=77 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_s102.py | 836 | 12 | entropy=5.17 len=47 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_s102.py | 837 | 12 | entropy=5.29 len=48 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 15 | 4 | entropy=5.39 len=73 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 29 | 17 | entropy=5.14 len=38 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 83 | 10 | entropy=5.02 len=59 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 110 | 9 | entropy=5.08 len=43 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 236 | 10 | entropy=5.02 len=40 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 261 | 12 | entropy=5.18 len=39 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 274 | 8 | entropy=6.05 len=74 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 288 | 12 | entropy=6.00 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 312 | 12 | entropy=6.00 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 328 | 9 | entropy=5.73 len=53 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 332 | 12 | entropy=6.00 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 333 | 12 | entropy=5.14 len=38 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 336 | 12 | entropy=5.18 len=39 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 337 | 12 | entropy=6.00 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 443 | 19 | entropy=5.00 len=43 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 446 | 12 | entropy=5.11 len=55 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 450 | 12 | entropy=5.04 len=53 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 455 | 12 | entropy=5.13 len=46 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 459 | 12 | entropy=5.00 len=60 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 461 | 19 | entropy=5.07 len=49 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 462 | 19 | entropy=5.15 len=42 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 465 | 12 | entropy=5.07 len=66 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 560 | 17 | entropy=5.02 len=41 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 688 | 12 | entropy=5.07 len=56 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 783 | 10 | entropy=5.06 len=60 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 1232 | 12 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 1237 | 12 | entropy=5.42 len=93 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 2380 | 12 | entropy=5.18 len=39 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 2391 | 12 | entropy=5.18 len=39 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 2438 | 12 | entropy=5.18 len=39 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 2467 | 8 | entropy=5.35 len=70 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 2519 | 10 | entropy=5.47 len=58 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets_credentials.py | 174 | 11 | entropy=5.15 len=76 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets_nonpy.py | 56 | 8 | entropy=5.24 len=70 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets_nonpy.py | 86 | 8 | entropy=5.04 len=84 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets_nonpy.py | 119 | 13 | entropy=5.14 len=65 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets_nonpy.py | 132 | 13 | entropy=5.05 len=45 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets_nonpy.py | 145 | 8 | entropy=5.17 len=83 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets_nonpy.py | 164 | 8 | entropy=5.23 len=65 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_typescript_expanded.py | 2225 | 15 | entropy=5.21 len=51 type=high_entropy | entropy | heuristic |

## Errors

No errors recorded.

## Affected files

- `.github/workflows/corpus.yml`
- `.github/workflows/examples/skylos-plus-claude-security.yml`
- `.github/workflows/examples/skylos-tokenless-ci.yml`
- `.github/workflows/liveness-primer.yml`
- `.github/workflows/parity.yml`
- `.github/workflows/pr-title.yml`
- `.github/workflows/publish.yml`
- `.github/workflows/quality-benchmark.yml`
- `.github/workflows/release-please.yml`
- `.github/workflows/repo-map-pages.yml`
- `.github/workflows/skylos.yaml`
- `.github/workflows/tests.yaml`
- `/tmp/safeanalyze-skylos/benchmarks/agent-pr-bench/harness/tools.py`
- `/tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/extreme_grounding_framework/services/hooks.py`
- `/tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/extreme_grounding_framework/tools/admin.py`
- `/tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/flask_getter_shell/app.py`
- `/tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/flask_handler_security/app.py`
- `/tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/flask_pickle_deserialization/app.py`
- `/tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/flask_reflected_xss/app.py`
- `/tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/shell_hook_runner/hooks.py`
- `/tmp/safeanalyze-skylos/benchmarks/agent_review/fixtures/static_blind_plugin_dispatch/plugins/audit.py`
- `/tmp/safeanalyze-skylos/benchmarks/ai_code_defects/fixtures/disabled_security_control/app.py`
- `/tmp/safeanalyze-skylos/benchmarks/ai_code_defects/fixtures/repo_level_security_regression/app.py`
- `/tmp/safeanalyze-skylos/benchmarks/dead_code/fixtures/extreme_grounding_framework/services/hooks.py`
- `/tmp/safeanalyze-skylos/benchmarks/dead_code/fixtures/extreme_grounding_framework/tools/admin.py`
- `/tmp/safeanalyze-skylos/benchmarks/dead_code/fixtures/static_blind_plugin_dispatch/plugins/audit.py`
- `/tmp/safeanalyze-skylos/benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py`
- `/tmp/safeanalyze-skylos/benchmarks/security/fixtures/java_ldap_xpath_injection/App.java`
- `/tmp/safeanalyze-skylos/benchmarks/security/fixtures/java_open_redirect_protocol_relative/App.java`
- `/tmp/safeanalyze-skylos/docs/examples/gitlab-code-quality.yml`
- `/tmp/safeanalyze-skylos/docs/examples/gitlab-managed-upload.yml`
- `/tmp/safeanalyze-skylos/editors/vscode/src/verifyCore.ts`
- `/tmp/safeanalyze-skylos/skylos/api/__init__.py`
- `/tmp/safeanalyze-skylos/skylos/api/_source_revision.py`
- `/tmp/safeanalyze-skylos/skylos/commands/rules_cmd.py`
- `/tmp/safeanalyze-skylos/skylos/done/runner.py`
- `/tmp/safeanalyze-skylos/skylos/rules/ai_defect/dependency_hallucination.py`
- `/tmp/safeanalyze-skylos/skylos/rules/ai_defect/manifest_dependency_hallucination.py`
- `/tmp/safeanalyze-skylos/skylos/rules/config/cicd/gitlab_ci.py`
- `/tmp/safeanalyze-skylos/skylos/rules/quality/clones.py`
- `/tmp/safeanalyze-skylos/skylos/rules/sca/licenses.py`
- `/tmp/safeanalyze-skylos/skylos/rules/sca/vulnerability_scanner.py`
- `/tmp/safeanalyze-skylos/skylos/visitors/languages/csharp/reachability.py`
- `/tmp/safeanalyze-skylos/test/test_agent_behavior_openai.py`
- `/tmp/safeanalyze-skylos/test/test_gitlab_report.py`
- `/tmp/safeanalyze-skylos/test/test_gitlab_scan_receipt.py`
- `/tmp/safeanalyze-skylos/test/test_mcp_rules.py`
- `/tmp/safeanalyze-skylos/test/test_pipfile_lockfile.py`
- `/tmp/safeanalyze-skylos/test/test_pnpm_regressions.py`
- `/tmp/safeanalyze-skylos/test/test_poetry_lockfile.py`
- `/tmp/safeanalyze-skylos/test/test_sca_pipfile.py`
- `/tmp/safeanalyze-skylos/test/test_sca_poetry_yarn.py`
- `/tmp/safeanalyze-skylos/test/test_secrets.py`
- `/tmp/safeanalyze-skylos/test/test_secrets_credentials.py`
- `/tmp/safeanalyze-skylos/test/test_secrets_nonpy.py`
- `/tmp/safeanalyze-skylos/test/test_shadow_compare.py`
- `/tmp/safeanalyze-skylos/test/test_uv_lockfile.py`
- `/tmp/safeanalyze-skylos/test/test_yarn_lockfile.py`
- `CHANGELOG.md`
- `SECURITY.md`
- `action.yml`
- `benchmarks/agent-pr-bench/bench.py`
- `benchmarks/agent-pr-bench/corpus/real_commits.json`
- `benchmarks/agent-pr-bench/corpus/seeded/cases/ex-09-secret-jwt.toml`
- `benchmarks/agent-pr-bench/corpus/seeded/cases/fa-11-cmdi.toml`
- `benchmarks/agent-pr-bench/corpus/seeded/cases/mb-14-secret-token.toml`
- `benchmarks/agent-pr-bench/corpus/seeded/cases/mb-c3-clean-subprocess.toml`
- `benchmarks/agent-pr-bench/harness/corpus.py`
- `benchmarks/agent-pr-bench/harness/tools.py`
- `benchmarks/agent-pr-bench/labels/findings.json`
- `benchmarks/agent-pr-bench/results/findings.jsonl`
- `benchmarks/agent-pr-bench/results/workspaces.json`
- `benchmarks/agent-pr-bench/scripts/select_real_commits.py`
- `benchmarks/agent_review/fixtures/extreme_grounding_framework/services/hooks.py`
- `benchmarks/agent_review/fixtures/extreme_grounding_framework/tools/admin.py`
- `benchmarks/agent_review/fixtures/flask_getter_shell/app.py`
- `benchmarks/agent_review/fixtures/flask_handler_security/app.py`
- `benchmarks/agent_review/fixtures/flask_reflected_xss/app.py`
- `benchmarks/agent_review/fixtures/shell_hook_runner/hooks.py`
- `benchmarks/agent_review/fixtures/static_blind_plugin_dispatch/plugins/audit.py`
- `benchmarks/agent_review/fixtures/static_blind_plugin_dispatch/plugins/tools.py`
- `benchmarks/dead_code/fixtures/extreme_grounding_framework/services/hooks.py`
- `benchmarks/dead_code/fixtures/extreme_grounding_framework/tools/admin.py`
- `benchmarks/dead_code/fixtures/static_blind_plugin_dispatch/plugins/audit.py`
- `benchmarks/dead_code/fixtures/static_blind_plugin_dispatch/plugins/tools.py`
- `benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py`
- `benchmarks/security/fixtures/subprocess_alias_shell/app.py`
- `benchmarks/verify_benchmark/fixtures/openai_sdk_member_drift/summarizer.py`
- `dictionary.md`
- `docs/agent-behavior-testing.md`
- `docs/container-image-reports.md`
- `docs/done-gate.md`
- `editors/vscode/out/ai.js`
- `editors/vscode/out/autoremediate.js`
- `editors/vscode/out/chatview.js`
- `editors/vscode/out/codelens.js`
- `editors/vscode/out/commandcenter.js`
- `editors/vscode/out/dashboard.js`
- `editors/vscode/out/scanner.js`
- `editors/vscode/out/verifyCore.js`
- `editors/vscode/package-lock.json`
- `editors/vscode/src/ai.ts`
- `editors/vscode/src/autoremediate.ts`
- `editors/vscode/src/chatview.ts`
- `editors/vscode/src/codelens.ts`
- `editors/vscode/src/dashboard.ts`
- `scripts/compare_codex_skylos_agent_review.py`
- `scripts/compare_codex_skylos_demo_deadcode.py`
- `scripts/compare_codex_skylos_quality.py`
- `scripts/deep_audit_logic_benchmark.py`
- `skylos/agents/evaluation/_openai_transport.py`
- `skylos/analysis/file_processing.py`
- `skylos/analyzer.py`
- `skylos/api/__init__.py`
- `skylos/api/_ai_detection.py`
- `skylos/api/_source_revision.py`
- `skylos/api/_upload_paths.py`
- `skylos/api/_upload_preflight.py`
- `skylos/api/_upload_transport.py`
- `skylos/audit/candidates.py`
- `skylos/audit/investigator_tools/operations.py`
- `skylos/audit/investigator_tools/validation.py`
- `skylos/benchmarks/_jev_dead_code_dataset.py`
- `skylos/benchmarks/ai_code_defects.py`
- `skylos/benchmarks/dead_code.py`
- `skylos/benchmarks/deep_audit_logic.py`
- `skylos/benchmarks/security.py`
- `skylos/benchmarks/verify_benchmark_runner.py`
- `skylos/cicd/evidence.py`
- `skylos/cicd/review.py`
- `skylos/cicd/workflow.py`
- `skylos/cli.py`
- `skylos/cli_core/main_parser.py`
- `skylos/cloud/login.py`
- `skylos/cloud/sync.py`
- `skylos/cloud/sync_setup.py`
- `skylos/commands/agent_verify_cmd.py`
- `skylos/commands/cache_cmd.py`
- `skylos/commands/cicd_cmd.py`
- `skylos/commands/compare_cmd.py`
- `skylos/commands/credits_cmd.py`
- `skylos/commands/hook_cmd.py`
- `skylos/commands/image_cmd.py`
- `skylos/commands/install_hooks_cmd.py`
- `skylos/commands/lint_cmd.py`
- `skylos/commands/preflight_cmd.py`
- `skylos/commands/review_cmd.py`
- `skylos/commands/rules_cmd.py`
- `skylos/commands/scan_cmd.py`
- `skylos/commands/verify_verdict_cmd.py`
- `skylos/core/api_symbol_truth.py`
- `skylos/core/baseline_source.py`
- `skylos/core/cli_shared.py`
- `skylos/core/file_discovery.py`
- `skylos/core/gatekeeper.py`
- `skylos/core/git_context.py`
- `skylos/core/grep_verify_common.py`
- `skylos/core/grep_verify_language_strategies.py`
- `skylos/core/grep_verify_python_strategy.py`
- `skylos/core/js_api_surface_utils.py`
- `skylos/core/review_context.py`
- `skylos/core/review_decisions.py`
- `skylos/deadcode/browser_refs.py`
- `skylos/debt/advisor.py`
- `skylos/debt/baseline.py`
- `skylos/defend/owasp.py`
- `skylos/done/base.py`
- `skylos/done/receipt.py`
- `skylos/done/runner.py`
- `skylos/engines/go/internal/analyzer/analyzer.go`
- `skylos/engines/go/internal/analyzer/analyzer_exec_command_test.go`
- `skylos/engines/go/internal/analyzer/analyzer_symlink_test.go`
- `skylos/engines/go/internal/symbols/symbols.go`
- `skylos/engines/go/internal/symbols/symbols_symlink_test.go`
- `skylos/engines/go/internal/symbols/typed_selectors.go`
- `skylos/engines/go_runner.py`
- `skylos/integrations/shadow.py`
- `skylos/integrations/trivy_image.py`
- `skylos/llm/agents.py`
- `skylos/llm/cleanup_orchestrator.py`
- `skylos/llm/context.py`
- `skylos/llm/executor.py`
- `skylos/llm/harness/ai_defect_challenge.py`
- `skylos/llm/investigator/reviewer_packs.py`
- `skylos/llm/prompts.py`
- `skylos/llm/security_verifier.py`
- `skylos/llm/verification/llm.py`
- `skylos/llm/verify_orchestrator.py`
- `skylos/misc/update_mapping.py`
- `skylos/pipeline.py`
- `skylos/preflight/evaluator.py`
- `skylos/preflight/inspector.py`
- `skylos/remediation/fixgen.py`
- `skylos/reporting/provenance.py`
- `skylos/reporting/sarif.py`
- `skylos/rules/ai_defect/dependency_bump_scan.py`
- `skylos/rules/ai_defect/dependency_hallucination.py`
- `skylos/rules/ai_defect/module_facts_index.py`
- `skylos/rules/ai_defect/pypi_wheel_modules.py`
- `skylos/rules/catalog.py`
- `skylos/rules/config/deployment/exposure.py`
- `skylos/rules/config/gpu/compatibility.py`
- `skylos/rules/danger/danger.py`
- `skylos/rules/danger/danger_fs/pytest_paths.py`
- `skylos/rules/danger/danger_mcp/mcp_flow.py`
- `skylos/rules/quality/code_health.py`
- `skylos/rules/sca/yarn_lockfile.py`
- `skylos/rules/secrets.py`
- `skylos/security/canonicalize.py`
- `skylos/security/command_guard_exfil.py`
- `skylos/security/command_guard_policy.py`
- `skylos/security/injection_scanner.py`
- `skylos/ui/terminal_report.py`
- `skylos/ui/tui.py`
- `skylos/verdict.py`
- `skylos/verification/context.py`
- `skylos/verification/refactor.py`
- `skylos/verify_change.py`
- `skylos/visitors/languages/csharp/danger.py`
- `skylos/visitors/languages/typescript/__init__.py`
- `skylos/visitors/languages/typescript/danger.py`
- `skylos/visitors/languages/typescript/source_lines.py`
- `skylos/visitors/languages/typescript/type_safety.py`
- `skylos/visitors/languages/typescript/workspace.py`
- `skylos_mcp/server.py`
- `test/conftest.py`
- `test/fixtures/app.go`
- `test/fixtures/upload_contract/long-message.payload.json`
- `test/fixtures/verdicts/verdict_failed.json`
- `test/fixtures/verdicts/verdict_keys.json`
- `test/fixtures/verdicts/verdict_passed.json`
- `test/test_action_image_scan.py`
- `test/test_action_sarif_upload.py`
- `test/test_action_sca.py`
- `test/test_agent_behavior_cli.py`
- `test/test_agent_behavior_contract.py`
- `test/test_agent_behavior_openai.py`
- `test/test_agent_bench_python_detections.py`
- `test/test_agent_center.py`
- `test/test_agent_center_review_memory.py`
- `test/test_agent_integration.py`
- `test/test_agent_llm_review_context.py`
- `test/test_agent_review_projection.py`
- `test/test_ai_pr_diff_rules.py`
- `test/test_analysis_review_context.py`
- `test/test_analyzer.py`
- `test/test_api.py`
- `test/test_api_signature_hallucination.py`
- `test/test_api_symbol_truth.py`
- `test/test_async_blocking.py`
- `test/test_audit_candidates.py`
- `test/test_audit_export.py`
- `test/test_audit_investigator_reviewer_packs.py`
- `test/test_audit_investigator_tools.py`
- `test/test_audit_processor.py`
- `test/test_audit_revalidator.py`
- `test/test_baseline_source.py`
- `test/test_behavior_baselines.py`
- `test/test_behavior_changes.py`
- `test/test_behavior_comparison.py`
- `test/test_cicd_evidence.py`
- `test/test_cicd_review.py`
- `test/test_cicd_workflow.py`
- `test/test_circular_deps.py`
- `test/test_cleanup_orchestrator.py`
- `test/test_cli.py`
- `test/test_cli_gitlab.py`
- `test/test_cli_image_scan.py`
- `test/test_cli_llm_provider.py`
- `test/test_cli_precommit.py`
- `test/test_cli_preflight.py`
- `test/test_cli_sca_baseline.py`
- `test/test_cli_terminal_probe.py`
- `test/test_cli_trivy_image.py`
- `test/test_cmd_injection.py`
- `test/test_code_health_diff.py`
- `test/test_community_rules.py`
- `test/test_container_python_runtime.py`
- `test/test_cpp_scanner.py`
- `test/test_csharp_danger_inline.py`
- `test/test_csharp_raw_lex.py`
- `test/test_dangerous.py`
- `test/test_dead_code_benchmark.py`
- `test/test_debt.py`
- `test/test_deep_audit_logic_benchmark_output.py`
- `test/test_defend.py`
- `test/test_dependency_providers.py`
- `test/test_dependency_version_bump.py`
- `test/test_dependency_version_bump_cli.py`
- `test/test_dependency_version_bump_integration.py`
- `test/test_deployment_exposure.py`
- `test/test_deserialization.py`
- `test/test_diff_dependencies.py`
- `test/test_diff_working_tree.py`
- `test/test_discover.py`
- `test/test_django_static_entrypoints.py`
- `test/test_done_gate.py`
- `test/test_done_receipt_binding.py`
- `test/test_fast_parity.py`
- `test/test_feedback.py`
- `test/test_file_discovery.py`
- `test/test_finding_verification.py`
- `test/test_fixgen.py`
- `test/test_fixgen_multilang.py`
- `test/test_gatekeeper.py`
- `test/test_git_context.py`
- `test/test_git_hook_regressions.py`
- `test/test_github_actions_finding_locations.py`
- `test/test_github_actions_runner_labels.py`
- `test/test_github_actions_security.py`
- `test/test_gitlab_delivery_receipt.py`
- `test/test_gitlab_report.py`
- `test/test_gitlab_upload_auth.py`
- `test/test_go_security.py`
- `test/test_gpu_compatibility.py`
- `test/test_grep_cache.py`
- `test/test_grep_python_backend.py`
- `test/test_grep_verify.py`
- `test/test_hallucination_dependency.py`
- `test/test_hook_cmd.py`
- `test/test_injection_scanner.py`
- `test/test_java_property_resources.py`
- `test/test_java_security.py`
- `test/test_java_source_helper_limits.py`
- `test/test_java_source_helpers.py`
- `test/test_jev_cascade_benchmark.py`
- `test/test_jev_dataset_safety.py`
- `test/test_jev_dead_code_benchmark.py`
- `test/test_jev_run_safety.py`
- `test/test_jev_triage.py`
- `test/test_js_barrel_exports.py`
- `test/test_licenses.py`
- `test/test_lint_cmd.py`
- `test/test_liveness_primer_workflow.py`
- `test/test_llm_analyzer.py`
- `test/test_llm_finding_evidence.py`
- `test/test_llm_graph.py`
- `test/test_llm_harness.py`
- `test/test_logic.py`
- `test/test_login.py`
- `test/test_mcp_auth.py`
- `test/test_mcp_reliability.py`
- `test/test_mcp_rules.py`
- `test/test_mcp_security.py`
- `test/test_mcp_summary.py`
- `test/test_method_reachability.py`
- `test/test_module_facts_index.py`
- `test/test_npm_lockfile.py`
- `test/test_osv_client.py`
- `test/test_output_validation_flow.py`
- `test/test_package_data.py`
- `test/test_pipeline.py`
- `test/test_pnpm_lockfile.py`
- `test/test_pnpm_regressions.py`
- `test/test_preflight.py`
- `test/test_preflight_inspector.py`
- `test/test_provenance.py`
- `test/test_provenance_attribution.py`
- `test/test_pypi_wheel_inventory.py`
- `test/test_pytest_path_parameters.py`
- `test/test_python_reachability.py`
- `test/test_quality_rules.py`
- `test/test_reachability_integration_qa.py`
- `test/test_refactor_verify.py`
- `test/test_repo_map.py`
- `test/test_review_cmd.py`
- `test/test_review_decisions.py`
- `test/test_review_memory_adversarial.py`
- `test/test_review_mutation_guards.py`
- `test/test_rules_cmd.py`
- `test/test_s102.py`
- `test/test_sanitizers.py`
- `test/test_sarif_exporter.py`
- `test/test_sbom_shrinkwrap.py`
- `test/test_sca_cache.py`
- `test/test_sca_lockfiles.py`
- `test/test_sca_pnpm.py`
- `test/test_sca_shrinkwrap.py`
- `test/test_secrets.py`
- `test/test_secrets_credentials.py`
- `test/test_secrets_nonpy.py`
- `test/test_security_benchmark.py`
- `test/test_security_contracts.py`
- `test/test_security_taskflow.py`
- `test/test_shell_security.py`
- `test/test_ssrf.py`
- `test/test_sync.py`
- `test/test_terminal_report.py`
- `test/test_test_impact_ai_defect.py`
- `test/test_triage_learner.py`
- `test/test_trivy_image.py`
- `test/test_ts_danger_d281_object_semantics_gaps.py`
- `test/test_ts_danger_d281_promise_temporal.py`
- `test/test_ts_danger_d281_redteam_regressions.py`
- `test/test_ts_danger_release_blockers.py`
- `test/test_ts_e003_entrypoints.py`
- `test/test_ts_esbuild_provenance.py`
- `test/test_ts_exports.py`
- `test/test_typescript_expanded.py`
- `test/test_typescript_framework.py`
- `test/test_typescript_inline_ignores.py`
- `test/test_typescript_quality_signals.py`
- `test/test_typescript_security_proof.py`
- `test/test_typescript_test_callback_length.py`
- `test/test_typescript_unsafe_export_regressions.py`
- `test/test_upload_contract_fixtures.py`
- `test/test_upload_reliability.py`
- `test/test_upload_revision_attribution.py`
- `test/test_uv_lockfile.py`
- `test/test_verify_change.py`
- `test/test_verify_change_dependency_bump.py`
- `test/test_verify_cmd.py`
- `test/test_verify_orchestrator.py`
- `test/test_verify_render.py`
- `test/test_verify_verdict.py`
- `test/test_vibe_rules.py`
- `test/test_xss_flow.py`
- `test/test_yarn_lockfile.py`
- `uv.lock`

