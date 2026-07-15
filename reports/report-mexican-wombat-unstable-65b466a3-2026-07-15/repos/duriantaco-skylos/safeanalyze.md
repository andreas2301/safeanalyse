# safeanalyze Report: /tmp/safeanalyze-skylos

- **Started:** 2026-07-15 01:40:56 UTC
- **Duration:** 24179 ms
- **Total findings:** 1371

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 1201 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 1371 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 57 |
| High | 258 |
| Medium | 727 |
| Low | 329 |

### Findings by source

| Source | Count |
| --- | --- |
| entropy | 707 |
| semgrep | 44 |
| trufflehog | 4 |
| yara | 616 |

## Findings

### Critical (57)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| backdoor_indicator | CHANGELOG.md | 237 | 25 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/audit/candidates.py | 388 | 65 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/debt/baseline.py | 259 | 26 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/debt/baseline.py | 280 | 41 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/debt/baseline.py | 449 | 41 | Common backdoor or persistence patterns | yara | deterministic |
| prompt_injection_comment | skylos/defend/owasp.py | 111 | 22 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | skylos/llm/prompts.py | 28 | 40 | Comment containing prompt injection keywords | yara | deterministic |
| backdoor_indicator | skylos/rules/catalog.py | 183 | 30 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/security/command_guard_exfil.py | 102 | 1 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/security/command_guard_exfil.py | 122 | 51 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/security/command_guard_policy.py | 83 | 20 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | skylos/security/command_guard_policy.py | 138 | 17 | Common backdoor or persistence patterns | yara | deterministic |
| backdoor_indicator | test/test_analyzer.py | 3011 | 16 | Common backdoor or persistence patterns | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 3926 | 27 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 3928 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 3951 | 39 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 3978 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_analyzer.py | 4027 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| backdoor_indicator | test/test_feedback.py | 39 | 17 | Common backdoor or persistence patterns | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 73 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 91 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 98 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 145 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 157 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 180 | 21 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 191 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 223 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 233 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 257 | 14 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 268 | 26 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 280 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 317 | 32 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 338 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 352 | 38 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 366 | 30 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 402 | 29 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 428 | 20 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 432 | 27 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 452 | 48 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 464 | 31 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 476 | 47 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 487 | 55 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_injection_scanner.py | 497 | 18 | Comment containing prompt injection keywords | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 25 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 25 | 39 | System prompt boundary marker or system instruction override | yara | deterministic |
| prompt_injection_comment | test/test_mcp_rules.py | 39 | 8 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | test/test_mcp_rules.py | 65 | 40 | Comment containing prompt injection keywords | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 89 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 89 | 28 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 101 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 101 | 52 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 114 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 114 | 20 | System prompt boundary marker or system instruction override | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 355 | 8 | System prompt boundary marker or system instruction override | yara | deterministic |
| prompt_injection_comment | test/test_mcp_rules.py | 355 | 16 | Comment containing prompt injection keywords | yara | deterministic |
| system_boundary | test/test_mcp_rules.py | 355 | 31 | System prompt boundary marker or system instruction override | yara | deterministic |
| backdoor_indicator | test/test_triage_learner.py | 497 | 11 | Common backdoor or persistence patterns | yara | deterministic |

### High (258)

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
| javascript.lang.security.detect-child-process.detect-child-process | /tmp/safeanalyze-skylos/editors/vscode/src/verifyCore.ts | 65 | 24 | Detected calls to child_process from a function argument `request`. This could lead to a command injection if the input is user controllable. Try to avoid calls to child_process, and if it is needed ensure user input is correctly sanitized or sandboxed.  | semgrep | deterministic |
| python.lang.security.use-defused-xml.use-defused-xml | /tmp/safeanalyze-skylos/skylos/rules/ai_defect/manifest_dependency_hallucination.py | 10 | 1 | The Python documentation recommends using `defusedxml` instead of `xml` because the native Python `xml` library is vulnerable to XML External Entity (XXE) attacks. These attacks can leak confidential data and "XML bombs" can cause denial of service. | semgrep | deterministic |
| gitlab | /tmp/safeanalyze-skylos/skylos/rules/config/cicd/gitlab_ci.py | 754 | 0 | Potential Gitlab secret detected | trufflehog | deterministic |
| stripe | /tmp/safeanalyze-skylos/test/test_mcp_rules.py | 359 | 0 | Potential Stripe secret detected | trufflehog | deterministic |
| stripe | /tmp/safeanalyze-skylos/test/test_mcp_rules.py | 367 | 0 | Potential Stripe secret detected | trufflehog | deterministic |
| gitlab | /tmp/safeanalyze-skylos/test/test_secrets.py | 51 | 0 | Potential Gitlab secret detected | trufflehog | deterministic |
| data_exfiltration_email | SECURITY.md | 15 | 1 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | benchmarks/agent_review/fixtures/debt_hotspot_service/ledger.py | 24 | 20 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | benchmarks/agent_review/fixtures/debt_hotspot_service/ledger.py | 46 | 28 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | benchmarks/agent_review/fixtures/extreme_grounding_framework/api.py | 14 | 32 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | benchmarks/dead_code/fixtures/extreme_grounding_framework/api.py | 14 | 32 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | benchmarks/security/fixtures/typescript_server_action_prisma_unsafe/actions.ts | 6 | 33 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration | editors/vscode/out/ai.js | 348 | 28 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration_email | editors/vscode/out/ai.js | 348 | 28 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration | editors/vscode/out/ai.js | 418 | 28 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration_email | editors/vscode/out/ai.js | 418 | 28 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration | editors/vscode/out/ai.js | 667 | 24 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration_email | editors/vscode/out/ai.js | 667 | 24 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration | editors/vscode/out/chatview.js | 148 | 28 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration_email | editors/vscode/out/chatview.js | 148 | 28 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | editors/vscode/out/codelens.js | 100 | 29 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration | editors/vscode/src/ai.ts | 375 | 24 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration_email | editors/vscode/src/ai.ts | 375 | 24 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration | editors/vscode/src/ai.ts | 459 | 24 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration_email | editors/vscode/src/ai.ts | 459 | 24 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration | editors/vscode/src/ai.ts | 741 | 22 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration_email | editors/vscode/src/ai.ts | 741 | 22 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration | editors/vscode/src/chatview.ts | 141 | 24 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration_email | editors/vscode/src/chatview.ts | 141 | 24 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | editors/vscode/src/codelens.ts | 85 | 21 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/api/__init__.py | 400 | 33 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | skylos/api/__init__.py | 409 | 20 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 805 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 805 | 54 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 805 | 81 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 807 | 56 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 832 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/cicd/review.py | 832 | 63 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/cicd/risk_passport.py | 224 | 24 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | skylos/commands/cache_cmd.py | 84 | 47 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/commands/cache_cmd.py | 115 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/commands/credits_cmd.py | 15 | 39 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | skylos/commands/credits_cmd.py | 18 | 20 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | skylos/commands/doctor_cmd.py | 88 | 40 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/commands/rules_cmd.py | 410 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/commands/scan_cmd.py | 894 | 34 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/core/api_symbol_truth.py | 261 | 33 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_language_strategies.py | 140 | 40 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_language_strategies.py | 140 | 73 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_python_strategy.py | 178 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_python_strategy.py | 178 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_python_strategy.py | 191 | 42 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/grep_verify_python_strategy.py | 191 | 62 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/core/js_api_surface_utils.py | 155 | 33 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/core/reference_index_store.py | 289 | 21 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | skylos/core/result_cache.py | 534 | 25 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | skylos/core/suite.py | 145 | 24 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
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
| data_exfiltration_email | skylos/rules/config/cicd/github_actions.py | 1817 | 28 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 39 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 40 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 41 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 42 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 43 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 44 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 45 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 46 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger.py | 47 | 6 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 8 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 14 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 20 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 26 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 41 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 7 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 42 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 43 | 7 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 43 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 43 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 43 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 43 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 44 | 7 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 44 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 44 | 19 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/rules/danger/danger_mcp/mcp_flow.py | 44 | 25 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/scale/parallel_static.py | 240 | 21 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
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
| obfuscated_javascript | skylos/ui/terminal_report.py | 69 | 38 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/ui/terminal_report.py | 69 | 43 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/ui/terminal_report.py | 69 | 47 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | skylos/ui/terminal_report.py | 69 | 52 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | skylos/ui/terminal_report.py | 315 | 32 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | skylos/visitors/languages/typescript/workspace.py | 310 | 71 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_cicd_workflow.py | 219 | 62 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| indirect_prompt_injection | test/test_community_rules.py | 53 | 13 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| indirect_prompt_injection | test/test_community_rules.py | 114 | 13 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| indirect_prompt_injection | test/test_community_rules.py | 135 | 13 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| obfuscated_javascript | test/test_dangerous.py | 49 | 66 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dangerous.py | 49 | 71 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dangerous.py | 49 | 77 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_dangerous.py | 153 | 30 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 427 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 427 | 28 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 459 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 459 | 61 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 766 | 31 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_debt.py | 766 | 35 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| indirect_prompt_injection | test/test_defend.py | 1821 | 5 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
| obfuscated_javascript | test/test_deserialization.py | 21 | 53 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_deserialization.py | 33 | 46 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_deserialization.py | 43 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_discover.py | 407 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| indirect_prompt_injection | test/test_discover.py | 556 | 5 | Content that may carry an indirect prompt-injection payload (user comment, email, web content, tool input delimiter) | yara | deterministic |
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
| data_exfiltration_email | test/test_llm_finding_evidence.py | 29 | 32 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | test/test_mcp_rules.py | 53 | 27 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pipeline.py | 149 | 23 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_pipeline.py | 149 | 37 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_rules_cmd.py | 112 | 23 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_rules_cmd.py | 122 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_rules_cmd.py | 123 | 14 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_security_taskflow.py | 229 | 59 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_security_taskflow.py | 229 | 67 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| suspicious_shell | test/test_shell_security.py | 311 | 1 | Suspicious shell command patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 105 | 32 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 106 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 106 | 50 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 117 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 118 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 119 | 17 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 120 | 21 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 120 | 36 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 126 | 54 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 144 | 13 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | test/test_terminal_report.py | 145 | 24 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| data_exfiltration_email | test/test_ts_danger_nextjs.py | 309 | 35 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 135 | 14 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 167 | 23 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 177 | 23 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 198 | 23 | Potential data exfiltration patterns | yara | deterministic |
| data_exfiltration | test/test_typescript_expanded.py | 1753 | 39 | Potential data exfiltration patterns | yara | deterministic |

### Medium (727)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| template_injection | .github/workflows/corpus.yml | 59 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/corpus.yml | 60 | 54 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-plus-claude-security.yml | 59 | 81 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-plus-claude-security.yml | 61 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-plus-claude-security.yml | 82 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-plus-claude-security.yml | 120 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/examples/skylos-tokenless-ci.yml | 44 | 90 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/parity.yml | 48 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/pr-title.yml | 17 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 24 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 35 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 42 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 67 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 74 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 75 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 121 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 244 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 257 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 262 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 298 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 299 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 304 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 305 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 306 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 307 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 327 | 73 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/publish.yml | 328 | 76 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/quality-benchmark.yml | 49 | 30 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/quality-benchmark.yml | 50 | 55 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 19 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 20 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 26 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 33 | 10 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 40 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/release-please.yml | 42 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/repo-map-pages.yml | 55 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 57 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 58 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 59 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 59 | 59 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 60 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 62 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 68 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 68 | 52 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 69 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 86 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 93 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 149 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 157 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 158 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/skylos.yaml | 163 | 33 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 14 | 25 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 43 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 53 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 56 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 86 | 10 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 91 | 18 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 91 | 72 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 92 | 40 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | .github/workflows/tests.yaml | 93 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
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
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/skylos/commands/rules_cmd.py | 196 | 14 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/skylos/rules/ai_defect/dependency_hallucination.py | 701 | 18 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected | /tmp/safeanalyze-skylos/skylos/rules/ai_defect/manifest_dependency_hallucination.py | 2131 | 10 | Detected a dynamic value being used with urllib. urllib supports 'file://' schemes, so a dynamic value controlled by a malicious actor may allow them to read arbitrary files. Audit uses of urllib calls to ensure user data cannot control the URLs, or consider using the 'requests' library instead. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/rules/quality/clones.py | 327 | 9 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/rules/quality/clones.py | 332 | 9 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/rules/quality/clones.py | 337 | 9 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| python.lang.security.insecure-hash-algorithms.insecure-hash-algorithm-sha1 | /tmp/safeanalyze-skylos/skylos/rules/quality/clones.py | 356 | 13 | Detected SHA1 hash algorithm which is considered insecure. SHA1 is not collision resistant and is therefore not suitable as a cryptographic signature. Use SHA256 or SHA3 instead. | semgrep | deterministic |
| template_injection | action.yml | 41 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 44 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 47 | 13 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 55 | 26 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 59 | 36 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 65 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 66 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 67 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 68 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 108 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 109 | 23 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 110 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 111 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 145 | 20 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | action.yml | 146 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
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
| entropy_base64_blob | editors/vscode/package-lock.json | 234 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 244 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 258 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 268 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 282 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 295 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 313 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 329 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 352 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 365 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 385 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 392 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 399 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 409 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 422 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 432 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 446 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 456 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 469 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 476 | 19 | entropy=5.53 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 499 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 509 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 516 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 529 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 536 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 543 | 19 | entropy=5.34 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 553 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 562 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 569 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 576 | 19 | entropy=5.23 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 582 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 597 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 644 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 663 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 677 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 691 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 705 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 719 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 733 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 747 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 761 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 775 | 19 | entropy=5.34 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 789 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 799 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 816 | 19 | entropy=5.25 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 832 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 845 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 861 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 868 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 878 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 885 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 896 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 903 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 925 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 941 | 19 | entropy=5.57 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 954 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 961 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 968 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 979 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 992 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1018 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1028 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1035 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1051 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1065 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1082 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1099 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1125 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1143 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1151 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1161 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1174 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1181 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1194 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1204 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1211 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1226 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1243 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1256 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1274 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1291 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1302 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1319 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1332 | 19 | entropy=5.59 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1345 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1355 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1366 | 19 | entropy=5.60 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1381 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1394 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1410 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1425 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1440 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1450 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1467 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1474 | 19 | entropy=5.29 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1488 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1499 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1512 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1525 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1535 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1545 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1558 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1574 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1585 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1592 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1609 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1626 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1636 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1649 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1666 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1683 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1691 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1706 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1716 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1741 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1755 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1763 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1788 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1801 | 19 | entropy=5.58 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1811 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1824 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1840 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1861 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1874 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1881 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1891 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1904 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1920 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1933 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1946 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1966 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1979 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 1993 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2007 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2020 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2042 | 19 | entropy=5.34 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2052 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2065 | 19 | entropy=5.26 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2073 | 19 | entropy=5.52 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2081 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2097 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2107 | 19 | entropy=5.52 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2117 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2130 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2149 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2159 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2175 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2182 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2200 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2216 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2223 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2246 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2253 | 19 | entropy=5.26 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2266 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2273 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2286 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2309 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2321 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2332 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2345 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2355 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2375 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2382 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2389 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2396 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2403 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2410 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2417 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2424 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2431 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2438 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2451 | 19 | entropy=5.23 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2479 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2489 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2496 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2506 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2520 | 19 | entropy=5.30 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2533 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2543 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2556 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2570 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2583 | 19 | entropy=5.23 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2594 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2604 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2612 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2619 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2626 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2634 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2648 | 19 | entropy=5.15 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2656 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2670 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2685 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2698 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2705 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2718 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2731 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2742 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2761 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2774 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2781 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2799 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2809 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2819 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2832 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2846 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2859 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2872 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2882 | 19 | entropy=5.52 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2899 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2909 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2922 | 19 | entropy=5.55 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2929 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2936 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2949 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2959 | 19 | entropy=5.67 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 2988 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3000 | 19 | entropy=5.62 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3010 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3026 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3047 | 19 | entropy=5.31 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3064 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3077 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3090 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3110 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3123 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3139 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3149 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3160 | 19 | entropy=5.39 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3173 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3197 | 19 | entropy=5.28 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3218 | 19 | entropy=5.33 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3225 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3235 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3257 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3270 | 19 | entropy=5.26 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3283 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3293 | 19 | entropy=5.35 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3313 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3330 | 19 | entropy=5.53 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3349 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3369 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3382 | 19 | entropy=5.62 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3404 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3431 | 19 | entropy=5.36 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3444 | 19 | entropy=5.45 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3462 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3473 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3480 | 19 | entropy=5.47 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3491 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3498 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3509 | 19 | entropy=5.52 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3524 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3534 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3547 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3563 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3574 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3584 | 19 | entropy=5.27 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3597 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3614 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3631 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3641 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3654 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3668 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3686 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3703 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3710 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3726 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3736 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3749 | 19 | entropy=5.42 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3756 | 19 | entropy=5.56 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3766 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3780 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3793 | 19 | entropy=5.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3805 | 19 | entropy=5.53 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3818 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3825 | 19 | entropy=5.49 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3832 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3842 | 19 | entropy=5.41 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3848 | 19 | entropy=5.43 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3861 | 19 | entropy=5.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3871 | 19 | entropy=5.38 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3878 | 19 | entropy=5.48 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3886 | 19 | entropy=5.40 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3897 | 19 | entropy=5.57 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3910 | 19 | entropy=5.57 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3924 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3934 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3950 | 19 | entropy=5.44 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3958 | 19 | entropy=5.37 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3974 | 19 | entropy=5.32 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3988 | 19 | entropy=5.65 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 3998 | 19 | entropy=5.54 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 4005 | 19 | entropy=5.64 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | editors/vscode/package-lock.json | 4019 | 19 | entropy=5.46 len=88 type=base64_blob | entropy | heuristic |
| output_constraint | editors/vscode/src/ai.ts | 603 | 116 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | editors/vscode/src/autoremediate.ts | 159 | 153 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| suspicious_imports | scripts/compare_codex_skylos_agent_review.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | scripts/compare_codex_skylos_demo_deadcode.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | scripts/compare_codex_skylos_quality.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/analyzer.py | 194 | 5 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/analyzer.py | 2516 | 17 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/analyzer.py | 2552 | 17 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/analyzer.py | 3031 | 21 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/api/__init__.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/api/_ai_detection.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/benchmarks/ai_code_defects.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/benchmarks/dead_code.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/benchmarks/security.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/benchmarks/verify_benchmark_runner.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/cicd/review.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/cicd/review.py | 439 | 97 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 60 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 61 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 115 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 135 | 50 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 148 | 45 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 236 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cicd/workflow.py | 272 | 31 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/cli.py | 47 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/cli.py | 3186 | 17 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/cloud/login.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/cloud/login.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/cloud/sync.py | 8 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/cloud/sync_setup.py | 96 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/cloud/sync_setup.py | 97 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/commands/agent_verify_cmd.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/commands/cicd_cmd.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| account_access_request | skylos/commands/scan_cmd.py | 900 | 61 | Natural-language request to access a user account or service | yara | deterministic |
| suspicious_imports | skylos/core/cli_shared.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/core/file_discovery.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/core/gatekeeper.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/core/grep_verify_common.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/core/grep_verify_language_strategies.py | 168 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/core/result_cache.py | 9 | 1 | Imports of known suspicious packages | yara | deterministic |
| output_constraint | skylos/debt/advisor.py | 60 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| credential_hardcode | skylos/engines/go/internal/analyzer/analyzer_symlink_test.go | 16 | 7 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | skylos/engines/go/internal/analyzer/analyzer_symlink_test.go | 26 | 7 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | skylos/engines/go/internal/symbols/symbols_symlink_test.go | 28 | 14 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | skylos/engines/go_runner.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| output_constraint | skylos/llm/agents.py | 460 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/agents.py | 477 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/cleanup_orchestrator.py | 193 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/cleanup_orchestrator.py | 233 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| credential_hardcode | skylos/llm/context.py | 664 | 8 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | skylos/llm/executor.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| output_constraint | skylos/llm/harness/ai_defect_challenge.py | 308 | 62 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 10 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 142 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 147 | 32 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 152 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 182 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 190 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 233 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 240 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 262 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 265 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 311 | 4 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 311 | 35 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/llm/prompts.py | 318 | 24 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 359 | 26 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 373 | 3 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/prompts.py | 377 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/security_verifier.py | 600 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/security_verifier.py | 613 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/verify_llm.py | 52 | 39 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | skylos/llm/verify_llm.py | 53 | 29 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 80 | 12 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 80 | 53 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 113 | 12 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 113 | 53 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 171 | 21 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 171 | 70 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 185 | 21 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 185 | 70 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 186 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| template_injection | skylos/llm/verify_llm.py | 378 | 1 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 407 | 12 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| output_constraint | skylos/llm/verify_llm.py | 407 | 61 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| suspicious_imports | skylos/llm/verify_orchestrator.py | 846 | 5 | Imports of known suspicious packages | yara | deterministic |
| template_injection | skylos/llm/verify_orchestrator.py | 1073 | 108 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | skylos/misc/update_mapping.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| output_constraint | skylos/pipeline.py | 74 | 1 | Instruction that constrains or suppresses the model's normal output | yara | deterministic |
| suspicious_imports | skylos/remediation/fixgen.py | 7 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/reporting/provenance.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/security/contracts.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | skylos/ui/tui.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/conftest.py | 218 | 13 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/fixtures/app.go | 13 | 7 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_agent_center.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_agent_integration.py | 1109 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_agent_integration.py | 1192 | 10 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_agent_integration.py | 1668 | 50 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_ai_pr_diff_rules.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_analyzer.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_api.py | 5 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_api.py | 485 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_api_signature_hallucination.py | 287 | 34 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_async_blocking.py | 61 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_audit_candidates.py | 28 | 34 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_candidates.py | 70 | 34 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_candidates.py | 90 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_export.py | 103 | 9 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_export.py | 104 | 25 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_processor.py | 294 | 22 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_processor.py | 346 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_audit_revalidator.py | 197 | 22 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_cicd_workflow.py | 188 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_cicd_workflow.py | 189 | 28 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_cleanup_orchestrator.py | 167 | 35 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cleanup_orchestrator.py | 252 | 35 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cleanup_orchestrator.py | 483 | 28 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cli.py | 703 | 22 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cli_llm_provider.py | 205 | 51 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cli_precommit.py | 536 | 10 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cli_precommit.py | 556 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_cli_precommit.py | 848 | 35 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_cli_terminal_probe.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cmd_injection.py | 21 | 13 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cmd_injection.py | 28 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_cmd_injection.py | 37 | 13 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 43 | 43 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 74 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 84 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 427 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 461 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dangerous.py | 504 | 13 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_dead_code_benchmark.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_debt.py | 802 | 13 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_defend.py | 890 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_defend.py | 2048 | 5 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_discover.py | 60 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_discover.py | 99 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_discover.py | 750 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_fast_parity.py | 6 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_file_discovery.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_file_discovery.py | 32 | 24 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_file_discovery.py | 54 | 24 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_fixgen.py | 3 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_fixgen_multilang.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_gatekeeper.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 49 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 181 | 22 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 276 | 17 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 297 | 12 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 299 | 21 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 387 | 15 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 413 | 14 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 422 | 19 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 426 | 27 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 461 | 54 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 509 | 47 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 564 | 59 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_github_actions_security.py | 565 | 14 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_go_security.py | 74 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_analyzer.py | 342 | 14 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_llm_analyzer.py | 826 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_analyzer.py | 879 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_analyzer.py | 1139 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_finding_evidence.py | 68 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_finding_evidence.py | 147 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_finding_evidence.py | 287 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_graph.py | 7 | 10 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_llm_graph.py | 26 | 10 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 496 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 632 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 666 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 733 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 784 | 59 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_llm_harness.py | 812 | 46 | Potential hardcoded credentials | yara | deterministic |
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
| suspicious_imports | test/test_provenance.py | 1 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_quality_rules.py | 375 | 1 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_quality_rules.py | 400 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_quality_rules.py | 415 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_quality_rules.py | 555 | 9 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_repo_map.py | 103 | 72 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_s102.py | 66 | 11 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_s102.py | 76 | 11 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_sanitizers.py | 58 | 13 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 24 | 19 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 58 | 11 | Potential hardcoded credentials | yara | deterministic |
| entropy_base64_blob | test/test_secrets.py | 67 | 10 | entropy=4.66 len=40 type=base64_blob | entropy | heuristic |
| credential_hardcode | test/test_secrets.py | 81 | 19 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 141 | 13 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 170 | 26 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 201 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 242 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets.py | 253 | 5 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_nonpy.py | 49 | 23 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_secrets_nonpy.py | 65 | 16 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_security_benchmark.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_security_contracts.py | 191 | 10 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_sync.py | 4 | 1 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | test/test_test_impact_ai_defect.py | 2 | 1 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_ts_exports.py | 904 | 8 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| suspicious_imports | test/test_typescript_expanded.py | 156 | 25 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_typescript_expanded.py | 1636 | 23 | Potential hardcoded credentials | yara | deterministic |
| template_injection | test/test_typescript_expanded.py | 1690 | 36 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| template_injection | test/test_typescript_expanded.py | 1691 | 41 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| credential_hardcode | test/test_typescript_expanded.py | 1865 | 23 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_verify_orchestrator.py | 423 | 10 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_verify_orchestrator.py | 1140 | 9 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_vibe_rules.py | 463 | 17 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 800 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1405 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1414 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1423 | 20 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1441 | 17 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1469 | 29 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1488 | 24 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | test/test_vibe_rules.py | 1497 | 17 | Potential hardcoded credentials | yara | deterministic |
| suspicious_imports | test/test_vibe_rules.py | 1710 | 9 | Imports of known suspicious packages | yara | deterministic |
| template_injection | test/test_xss_flow.py | 87 | 12 | Handlebars/Mustache/Jinja or GitHub Actions template interpolation that may inject instructions into a prompt | yara | deterministic |
| entropy_base64_blob | uv.lock | 17 | 12 | entropy=4.51 len=96 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 484 | 12 | entropy=4.53 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 490 | 12 | entropy=4.50 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 492 | 12 | entropy=4.52 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 507 | 12 | entropy=4.50 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 677 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 686 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 687 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 688 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 690 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 691 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 692 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 694 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 696 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 698 | 12 | entropy=4.55 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 701 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 702 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 703 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 706 | 12 | entropy=4.57 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 712 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 713 | 12 | entropy=4.57 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 714 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 717 | 12 | entropy=4.58 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 719 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 721 | 12 | entropy=4.56 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 725 | 12 | entropy=4.55 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 726 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 727 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 730 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 731 | 12 | entropy=4.55 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 734 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 735 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 738 | 12 | entropy=4.55 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 739 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 740 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 743 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 744 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 745 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 746 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 748 | 12 | entropy=4.54 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 749 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 750 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 751 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 753 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 755 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 757 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 760 | 12 | entropy=4.56 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 767 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 768 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 769 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 772 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 774 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 776 | 12 | entropy=4.52 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 779 | 12 | entropy=4.53 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 780 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 783 | 12 | entropy=4.56 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 785 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 934 | 16 | entropy=4.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 943 | 16 | entropy=4.51 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 945 | 12 | entropy=4.51 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 1617 | 12 | entropy=4.55 len=92 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 1764 | 12 | entropy=4.51 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 1793 | 12 | entropy=4.50 len=88 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 1983 | 12 | entropy=4.51 len=87 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 2412 | 16 | entropy=4.56 len=91 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 2414 | 12 | entropy=4.53 len=91 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 2633 | 12 | entropy=4.50 len=90 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | uv.lock | 2635 | 12 | entropy=4.51 len=90 type=base64_blob | entropy | heuristic |

### Low (329)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| entropy_high_entropy | .github/workflows/quality-benchmark.yml | 55 | 106 | entropy=5.16 len=45 type=high_entropy | entropy | heuristic |
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
| entropy_high_entropy | editors/vscode/package-lock.json | 234 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 244 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 258 | 19 | entropy=5.32 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 268 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 282 | 19 | entropy=5.31 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 295 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 313 | 19 | entropy=5.35 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 329 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 352 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 365 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 385 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 392 | 19 | entropy=5.34 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 399 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 409 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 422 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 432 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 446 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 456 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 469 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 476 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 499 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 509 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 516 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 529 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 536 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 543 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 553 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 562 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 569 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 576 | 19 | entropy=5.30 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 582 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 597 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 644 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 663 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 677 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 691 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 705 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 719 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 733 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 747 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 761 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 775 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 789 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 799 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 816 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 832 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 845 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 861 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 868 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 878 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 885 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 896 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 903 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 925 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 941 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 954 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 961 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 968 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 979 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 992 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1018 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1028 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1035 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1051 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1065 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1082 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1099 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1125 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1143 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1151 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1161 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1174 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1181 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1194 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1204 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1211 | 19 | entropy=5.34 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1226 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1243 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1274 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1291 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1302 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1319 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1332 | 19 | entropy=5.62 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1345 | 19 | entropy=5.38 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1355 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1366 | 19 | entropy=5.68 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1381 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1394 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1410 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1425 | 19 | entropy=5.66 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1440 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1450 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1467 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1474 | 19 | entropy=5.35 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1488 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1499 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1512 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1525 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1535 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1545 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1558 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1574 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1585 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1592 | 19 | entropy=5.62 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1609 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1626 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1636 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1649 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1666 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1683 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1691 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1706 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1716 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1741 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1755 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1763 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1788 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1801 | 19 | entropy=5.66 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1811 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1824 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1840 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1861 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1874 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1881 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1891 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1920 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1933 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1946 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1966 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1979 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 1993 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2007 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2020 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2042 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2052 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2065 | 19 | entropy=5.36 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2073 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2081 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2097 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2107 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2117 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2130 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2149 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2159 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2175 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2182 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2200 | 19 | entropy=5.63 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2216 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2223 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2246 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2253 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2266 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2273 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2286 | 19 | entropy=5.63 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2309 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2321 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2332 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2345 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2355 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2375 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2382 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2389 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2396 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2403 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2410 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2417 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2424 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2431 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2438 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2451 | 19 | entropy=5.28 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2479 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2489 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2496 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2506 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2520 | 19 | entropy=5.37 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2533 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2543 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2556 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2570 | 19 | entropy=5.36 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2583 | 19 | entropy=5.28 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2594 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2604 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2612 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2619 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2626 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2634 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2648 | 19 | entropy=5.22 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2656 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2670 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2685 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2698 | 19 | entropy=5.38 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2705 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2718 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2731 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2742 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2761 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2774 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2781 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2799 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2809 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2819 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2832 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2846 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2859 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2872 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2882 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2899 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2909 | 19 | entropy=5.42 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2922 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2929 | 19 | entropy=5.43 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2936 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2949 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2959 | 19 | entropy=5.69 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 2988 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3000 | 19 | entropy=5.65 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3010 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3026 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3047 | 19 | entropy=5.36 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3064 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3077 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3090 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3110 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3123 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3139 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3149 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3160 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3173 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3197 | 19 | entropy=5.38 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3218 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3225 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3235 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3257 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3270 | 19 | entropy=5.33 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3283 | 19 | entropy=5.58 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3293 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3313 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3330 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3349 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3369 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3382 | 19 | entropy=5.64 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3404 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3431 | 19 | entropy=5.45 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3444 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3462 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3473 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3480 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3491 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3509 | 19 | entropy=5.57 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3524 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3534 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3547 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3563 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3574 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3584 | 19 | entropy=5.34 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3597 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3614 | 19 | entropy=5.47 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3631 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3641 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3654 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3668 | 19 | entropy=5.67 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3686 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3703 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3710 | 19 | entropy=5.49 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3726 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3736 | 19 | entropy=5.54 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3749 | 19 | entropy=5.52 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3756 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3766 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3780 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3793 | 19 | entropy=5.56 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3805 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3818 | 19 | entropy=5.39 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3825 | 19 | entropy=5.51 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3832 | 19 | entropy=5.55 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3842 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3848 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3861 | 19 | entropy=5.59 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3871 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3878 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3886 | 19 | entropy=5.48 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3897 | 19 | entropy=5.64 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3910 | 19 | entropy=5.61 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3924 | 19 | entropy=5.46 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3934 | 19 | entropy=5.44 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3950 | 19 | entropy=5.53 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3958 | 19 | entropy=5.41 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3974 | 19 | entropy=5.40 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3988 | 19 | entropy=5.66 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 3998 | 19 | entropy=5.60 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 4005 | 19 | entropy=5.70 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/package-lock.json | 4019 | 19 | entropy=5.50 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | editors/vscode/src/dashboard.ts | 205 | 35 | entropy=5.01 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/analysis/file_processing.py | 265 | 13 | entropy=5.05 len=68 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/analyzer.py | 2225 | 21 | entropy=5.01 len=67 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/api/__init__.py | 2066 | 9 | entropy=5.13 len=116 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/visitors/languages/csharp/danger.py | 52 | 5 | entropy=5.00 len=55 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | skylos/visitors/languages/typescript/danger.py | 142 | 4 | entropy=6.07 len=67 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_evidence.py | 10 | 29 | entropy=5.17 len=36 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_evidence.py | 14 | 25 | entropy=5.00 len=32 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_cicd_review.py | 38 | 29 | entropy=5.17 len=36 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_java_security.py | 90 | 6 | entropy=6.02 len=65 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_mcp_rules.py | 287 | 30 | entropy=5.22 len=40 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_s102.py | 32 | 9 | entropy=5.44 len=55 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_s102.py | 76 | 9 | entropy=5.43 len=52 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 24 | 10 | entropy=5.02 len=59 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 51 | 9 | entropy=5.08 len=43 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 176 | 10 | entropy=5.02 len=40 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 201 | 12 | entropy=5.18 len=39 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 242 | 12 | entropy=5.18 len=39 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_secrets.py | 253 | 12 | entropy=5.18 len=39 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | test/test_typescript_expanded.py | 1636 | 15 | entropy=5.21 len=51 type=high_entropy | entropy | heuristic |

## Errors

No errors recorded.

## Affected files

- `.github/workflows/corpus.yml`
- `.github/workflows/examples/skylos-plus-claude-security.yml`
- `.github/workflows/examples/skylos-tokenless-ci.yml`
- `.github/workflows/parity.yml`
- `.github/workflows/pr-title.yml`
- `.github/workflows/publish.yml`
- `.github/workflows/quality-benchmark.yml`
- `.github/workflows/release-please.yml`
- `.github/workflows/repo-map-pages.yml`
- `.github/workflows/skylos.yaml`
- `.github/workflows/tests.yaml`
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
- `/tmp/safeanalyze-skylos/editors/vscode/src/verifyCore.ts`
- `/tmp/safeanalyze-skylos/skylos/commands/rules_cmd.py`
- `/tmp/safeanalyze-skylos/skylos/rules/ai_defect/dependency_hallucination.py`
- `/tmp/safeanalyze-skylos/skylos/rules/ai_defect/manifest_dependency_hallucination.py`
- `/tmp/safeanalyze-skylos/skylos/rules/config/cicd/gitlab_ci.py`
- `/tmp/safeanalyze-skylos/skylos/rules/quality/clones.py`
- `/tmp/safeanalyze-skylos/test/test_mcp_rules.py`
- `/tmp/safeanalyze-skylos/test/test_secrets.py`
- `CHANGELOG.md`
- `SECURITY.md`
- `action.yml`
- `benchmarks/agent_review/fixtures/debt_hotspot_service/ledger.py`
- `benchmarks/agent_review/fixtures/extreme_grounding_framework/api.py`
- `benchmarks/agent_review/fixtures/extreme_grounding_framework/services/hooks.py`
- `benchmarks/agent_review/fixtures/extreme_grounding_framework/tools/admin.py`
- `benchmarks/agent_review/fixtures/flask_getter_shell/app.py`
- `benchmarks/agent_review/fixtures/flask_handler_security/app.py`
- `benchmarks/agent_review/fixtures/flask_reflected_xss/app.py`
- `benchmarks/agent_review/fixtures/shell_hook_runner/hooks.py`
- `benchmarks/agent_review/fixtures/static_blind_plugin_dispatch/plugins/audit.py`
- `benchmarks/agent_review/fixtures/static_blind_plugin_dispatch/plugins/tools.py`
- `benchmarks/dead_code/fixtures/extreme_grounding_framework/api.py`
- `benchmarks/dead_code/fixtures/extreme_grounding_framework/services/hooks.py`
- `benchmarks/dead_code/fixtures/extreme_grounding_framework/tools/admin.py`
- `benchmarks/dead_code/fixtures/static_blind_plugin_dispatch/plugins/audit.py`
- `benchmarks/dead_code/fixtures/static_blind_plugin_dispatch/plugins/tools.py`
- `benchmarks/security/fixtures/intentional_vulnerable_flask_app/app.py`
- `benchmarks/security/fixtures/subprocess_alias_shell/app.py`
- `benchmarks/security/fixtures/typescript_server_action_prisma_unsafe/actions.ts`
- `benchmarks/verify_benchmark/fixtures/openai_sdk_member_drift/summarizer.py`
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
- `skylos/analysis/file_processing.py`
- `skylos/analyzer.py`
- `skylos/api/__init__.py`
- `skylos/api/_ai_detection.py`
- `skylos/audit/candidates.py`
- `skylos/benchmarks/ai_code_defects.py`
- `skylos/benchmarks/dead_code.py`
- `skylos/benchmarks/security.py`
- `skylos/benchmarks/verify_benchmark_runner.py`
- `skylos/cicd/review.py`
- `skylos/cicd/risk_passport.py`
- `skylos/cicd/workflow.py`
- `skylos/cli.py`
- `skylos/cloud/login.py`
- `skylos/cloud/sync.py`
- `skylos/cloud/sync_setup.py`
- `skylos/commands/agent_verify_cmd.py`
- `skylos/commands/cache_cmd.py`
- `skylos/commands/cicd_cmd.py`
- `skylos/commands/credits_cmd.py`
- `skylos/commands/doctor_cmd.py`
- `skylos/commands/rules_cmd.py`
- `skylos/commands/scan_cmd.py`
- `skylos/core/api_symbol_truth.py`
- `skylos/core/cli_shared.py`
- `skylos/core/file_discovery.py`
- `skylos/core/gatekeeper.py`
- `skylos/core/grep_verify_common.py`
- `skylos/core/grep_verify_language_strategies.py`
- `skylos/core/grep_verify_python_strategy.py`
- `skylos/core/js_api_surface_utils.py`
- `skylos/core/reference_index_store.py`
- `skylos/core/result_cache.py`
- `skylos/core/suite.py`
- `skylos/debt/advisor.py`
- `skylos/debt/baseline.py`
- `skylos/defend/owasp.py`
- `skylos/engines/go/internal/analyzer/analyzer.go`
- `skylos/engines/go/internal/analyzer/analyzer_exec_command_test.go`
- `skylos/engines/go/internal/analyzer/analyzer_symlink_test.go`
- `skylos/engines/go/internal/symbols/symbols.go`
- `skylos/engines/go/internal/symbols/symbols_symlink_test.go`
- `skylos/engines/go/internal/symbols/typed_selectors.go`
- `skylos/engines/go_runner.py`
- `skylos/llm/agents.py`
- `skylos/llm/cleanup_orchestrator.py`
- `skylos/llm/context.py`
- `skylos/llm/executor.py`
- `skylos/llm/harness/ai_defect_challenge.py`
- `skylos/llm/prompts.py`
- `skylos/llm/security_verifier.py`
- `skylos/llm/verify_llm.py`
- `skylos/llm/verify_orchestrator.py`
- `skylos/misc/update_mapping.py`
- `skylos/pipeline.py`
- `skylos/remediation/fixgen.py`
- `skylos/reporting/provenance.py`
- `skylos/rules/catalog.py`
- `skylos/rules/config/cicd/github_actions.py`
- `skylos/rules/danger/danger.py`
- `skylos/rules/danger/danger_mcp/mcp_flow.py`
- `skylos/scale/parallel_static.py`
- `skylos/security/canonicalize.py`
- `skylos/security/command_guard_exfil.py`
- `skylos/security/command_guard_policy.py`
- `skylos/security/contracts.py`
- `skylos/security/injection_scanner.py`
- `skylos/ui/terminal_report.py`
- `skylos/ui/tui.py`
- `skylos/visitors/languages/csharp/danger.py`
- `skylos/visitors/languages/typescript/danger.py`
- `skylos/visitors/languages/typescript/workspace.py`
- `test/conftest.py`
- `test/fixtures/app.go`
- `test/test_agent_center.py`
- `test/test_agent_integration.py`
- `test/test_ai_pr_diff_rules.py`
- `test/test_analyzer.py`
- `test/test_api.py`
- `test/test_api_signature_hallucination.py`
- `test/test_async_blocking.py`
- `test/test_audit_candidates.py`
- `test/test_audit_export.py`
- `test/test_audit_processor.py`
- `test/test_audit_revalidator.py`
- `test/test_cicd_evidence.py`
- `test/test_cicd_review.py`
- `test/test_cicd_workflow.py`
- `test/test_cleanup_orchestrator.py`
- `test/test_cli.py`
- `test/test_cli_llm_provider.py`
- `test/test_cli_precommit.py`
- `test/test_cli_terminal_probe.py`
- `test/test_cmd_injection.py`
- `test/test_community_rules.py`
- `test/test_dangerous.py`
- `test/test_dead_code_benchmark.py`
- `test/test_debt.py`
- `test/test_defend.py`
- `test/test_deserialization.py`
- `test/test_discover.py`
- `test/test_fast_parity.py`
- `test/test_feedback.py`
- `test/test_file_discovery.py`
- `test/test_fixgen.py`
- `test/test_fixgen_multilang.py`
- `test/test_gatekeeper.py`
- `test/test_github_actions_security.py`
- `test/test_go_security.py`
- `test/test_injection_scanner.py`
- `test/test_java_security.py`
- `test/test_llm_analyzer.py`
- `test/test_llm_finding_evidence.py`
- `test/test_llm_graph.py`
- `test/test_llm_harness.py`
- `test/test_login.py`
- `test/test_mcp_auth.py`
- `test/test_mcp_rules.py`
- `test/test_mcp_security.py`
- `test/test_pipeline.py`
- `test/test_provenance.py`
- `test/test_quality_rules.py`
- `test/test_repo_map.py`
- `test/test_rules_cmd.py`
- `test/test_s102.py`
- `test/test_sanitizers.py`
- `test/test_secrets.py`
- `test/test_secrets_nonpy.py`
- `test/test_security_benchmark.py`
- `test/test_security_contracts.py`
- `test/test_security_taskflow.py`
- `test/test_shell_security.py`
- `test/test_sync.py`
- `test/test_terminal_report.py`
- `test/test_test_impact_ai_defect.py`
- `test/test_triage_learner.py`
- `test/test_ts_danger_nextjs.py`
- `test/test_ts_exports.py`
- `test/test_typescript_expanded.py`
- `test/test_verify_orchestrator.py`
- `test/test_vibe_rules.py`
- `test/test_xss_flow.py`
- `uv.lock`

