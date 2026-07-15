# safeanalyze Report: /tmp/safeanalyze-BIPIA

- **Started:** 2026-07-15 02:03:50 UTC
- **Duration:** 3235 ms
- **Total findings:** 331

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 99 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 331 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 2 |
| High | 193 |
| Medium | 80 |
| Low | 56 |

### Findings by source

| Source | Count |
| --- | --- |
| entropy | 57 |
| hiddenchars | 152 |
| semgrep | 1 |
| yara | 121 |

## Findings

### Critical (2)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| prompt_injection_comment | defense/README.md | 77 | 137 | Comment containing prompt injection keywords | yara | deterministic |
| backdoor_indicator | defense/white_box/utils.py | 33 | 10 | Common backdoor or persistence patterns | yara | deterministic |

### High (193)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| data_exfiltration_email | SECURITY.md | 15 | 45 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| data_exfiltration_email | SECURITY.md | 15 | 50 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| obfuscated_javascript | benchmark/code/test.jsonl | 2 | 866 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmark/code/test.jsonl | 2 | 881 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmark/code/test.jsonl | 2 | 971 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmark/code/test.jsonl | 2 | 1034 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmark/code/test.jsonl | 2 | 1133 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | benchmark/code/test.jsonl | 2 | 1178 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 3 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 5 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 18 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 24 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 25 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 34 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 38 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 42 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 46 | 28 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 55 | 26 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 56 | 39 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 61 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| suspicious_shell | benchmark/code_attack_test.json | 61 | 116 | Suspicious shell command patterns | yara | deterministic |
| suspicious_shell | benchmark/code_attack_test.json | 62 | 130 | Suspicious shell command patterns | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 68 | 21 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_test.json | 69 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 3 | 28 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 6 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 7 | 36 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 20 | 26 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 24 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 25 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 35 | 11 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 42 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| data_exfiltration_email | benchmark/code_attack_train.json | 42 | 241 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 47 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 48 | 13 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 52 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 67 | 10 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| instruction_to_include_code | benchmark/code_attack_train.json | 70 | 11 | Natural-language instruction telling the model to include, merge, or execute a code snippet | yara | deterministic |
| data_exfiltration_email | benchmark/email/test.jsonl | 37 | 729 | Natural-language request to retrieve data and send it to an email address | yara | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 254 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 260 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 390 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 396 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 529 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 535 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 662 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 668 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 798 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 804 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 934 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 940 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1070 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1076 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1206 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1212 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1342 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1348 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1478 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1484 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1614 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1620 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1750 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1756 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1886 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 43 | 1892 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 304 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 310 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 467 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 473 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 630 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 636 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 793 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 799 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 956 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 962 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1119 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1125 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1282 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1288 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1445 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1451 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1608 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1614 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1771 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1777 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1932 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 125 | 1938 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 508 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 514 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 881 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 887 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 1250 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 1256 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 1631 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 1637 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 2004 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 2010 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 2376 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 163 | 2382 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 263 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 269 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 400 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 406 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 537 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 543 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 664 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 670 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 807 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 813 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 944 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 950 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 1081 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 1087 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 1218 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 458 | 1224 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 310 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 316 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 473 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 479 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 636 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 642 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 799 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 805 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 962 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 968 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1125 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1131 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1288 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1294 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1451 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1457 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1614 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1620 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1777 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1783 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1938 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 632 | 1944 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 258 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 264 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 395 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 401 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 532 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 538 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 659 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 665 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 802 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 808 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 939 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 945 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 1076 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 1082 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 1213 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 652 | 1219 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 508 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 514 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 881 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 887 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 1250 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 1256 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 1631 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 1637 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 2004 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 2010 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 2376 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 848 | 2382 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 263 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 269 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 399 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 405 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 538 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 544 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 671 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 677 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 807 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 813 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 943 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 949 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1079 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1085 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1215 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1221 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1351 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1357 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1487 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1493 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1623 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1629 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1759 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1765 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1895 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| hidden_char_zero_width | benchmark/table/train.jsonl | 853 | 1901 | BYTE ORDER MARK (zero_width) in benchmark/table/train.jsonl | hiddenchars | deterministic |
| suspicious_shell | bipia/metrics/regist.py | 267 | 38 | Suspicious shell command patterns | yara | deterministic |
| suspicious_shell | bipia/metrics/regist.py | 268 | 52 | Suspicious shell command patterns | yara | deterministic |

### Medium (80)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| python.lang.security.audit.eval-detected.eval-detected | /tmp/safeanalyze-BIPIA/benchmark/qa/process.py | 21 | 32 | Detected the use of eval(). eval() can be dangerous if used to evaluate dynamic content. If this content can be input from outside the program, this may be a code injection vulnerability. Ensure evaluated content is not definable by external sources. | semgrep | deterministic |
| suspicious_imports | benchmark/code/train.jsonl | 9 | 2092 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code/train.jsonl | 9 | 2410 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code/train.jsonl | 9 | 2683 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 4 | 87 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 10 | 88 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 11 | 89 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 17 | 88 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 21 | 87 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 24 | 72 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 25 | 77 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 26 | 75 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 27 | 84 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 28 | 96 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 32 | 93 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 34 | 79 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 35 | 105 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 38 | 90 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 39 | 87 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 46 | 93 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 60 | 77 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 60 | 77 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 61 | 91 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 62 | 91 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 66 | 111 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 68 | 90 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_test.json | 69 | 81 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 10 | 104 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 12 | 98 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 19 | 100 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 20 | 98 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 32 | 75 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 45 | 74 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 46 | 78 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 48 | 80 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 49 | 90 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 55 | 87 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 61 | 73 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 63 | 77 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 66 | 72 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 67 | 77 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 68 | 79 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 69 | 82 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | benchmark/code_attack_train.json | 70 | 85 | Imports of known suspicious packages | yara | deterministic |
| account_access_request | benchmark/email/test.jsonl | 37 | 230 | Natural-language request to access a user account or service | yara | deterministic |
| account_access_request | benchmark/email/train.jsonl | 44 | 484 | Natural-language request to access a user account or service | yara | deterministic |
| entropy_base64_blob | bipia/metrics/eval/utils.py | 30 | 19 | entropy=5.86 len=58 type=base64_blob | entropy | heuristic |
| suspicious_imports | bipia/metrics/regist.py | 210 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 216 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 217 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 223 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 227 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 230 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 231 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 232 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 233 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 234 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 238 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 240 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 241 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 244 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 245 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 252 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 266 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 266 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 267 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 268 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 272 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 274 | 14 | Imports of known suspicious packages | yara | deterministic |
| suspicious_imports | bipia/metrics/regist.py | 275 | 14 | Imports of known suspicious packages | yara | deterministic |
| credential_hardcode | bipia/model/llm_worker.py | 217 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | bipia/model/llm_worker.py | 274 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | bipia/model/vllm_worker.py | 113 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | bipia/model/vllm_worker.py | 130 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | bipia/model/vllm_worker.py | 147 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | bipia/model/vllm_worker.py | 163 | 32 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | config/gpt35.yaml | 1 | 1 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | config/gpt35_wosys.yaml | 1 | 1 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | config/gpt4.yaml | 1 | 1 | Potential hardcoded credentials | yara | deterministic |
| credential_hardcode | config/gpt4_wosys.yaml | 1 | 1 | Potential hardcoded credentials | yara | deterministic |

### Low (56)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| entropy_high_entropy | benchmark/code/train.jsonl | 45 | 962 | entropy=5.18 len=141 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/code/train.jsonl | 45 | 1851 | entropy=5.18 len=141 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/code_attack_test.json | 60 | 8 | entropy=5.01 len=307 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 2 | 12 | entropy=5.05 len=694 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 4 | 12 | entropy=5.19 len=643 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 8 | 12 | entropy=5.27 len=164 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 9 | 12 | entropy=5.20 len=512 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 10 | 12 | entropy=5.05 len=725 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 11 | 12 | entropy=5.20 len=620 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 17 | 12 | entropy=5.13 len=615 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 18 | 12 | entropy=5.07 len=597 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 19 | 12 | entropy=5.19 len=643 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 20 | 12 | entropy=5.15 len=227 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 23 | 12 | entropy=5.08 len=565 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 24 | 12 | entropy=5.17 len=595 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 26 | 12 | entropy=5.11 len=600 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 28 | 12 | entropy=5.08 len=547 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 30 | 12 | entropy=5.22 len=437 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 34 | 12 | entropy=5.14 len=778 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 38 | 12 | entropy=5.13 len=312 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 39 | 12 | entropy=5.12 len=557 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 40 | 12 | entropy=5.04 len=393 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 41 | 12 | entropy=5.14 len=778 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 42 | 12 | entropy=5.13 len=515 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 43 | 12 | entropy=5.17 len=595 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 44 | 12 | entropy=5.13 len=554 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 45 | 12 | entropy=5.04 len=722 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/test.jsonl | 47 | 12 | entropy=5.13 len=507 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 1 | 12 | entropy=5.02 len=440 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 6 | 12 | entropy=5.05 len=572 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 7 | 12 | entropy=5.16 len=607 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 11 | 12 | entropy=5.13 len=546 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 14 | 12 | entropy=5.08 len=565 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 16 | 12 | entropy=5.02 len=440 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 21 | 12 | entropy=5.09 len=523 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 26 | 12 | entropy=5.16 len=651 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 27 | 12 | entropy=5.07 len=546 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 28 | 12 | entropy=5.01 len=593 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 29 | 12 | entropy=5.15 len=461 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 31 | 12 | entropy=5.14 len=278 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 32 | 12 | entropy=5.20 len=587 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 36 | 12 | entropy=5.06 len=324 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 38 | 12 | entropy=5.20 len=512 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 39 | 12 | entropy=5.16 len=607 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 41 | 12 | entropy=5.12 len=489 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 42 | 12 | entropy=5.13 len=312 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 46 | 12 | entropy=5.14 len=641 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 47 | 12 | entropy=5.08 len=573 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 48 | 12 | entropy=5.14 len=278 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 49 | 12 | entropy=5.07 len=573 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | benchmark/email/train.jsonl | 50 | 12 | entropy=5.10 len=286 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | bipia/metrics/eval/utils.py | 30 | 19 | entropy=5.86 len=58 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | demo.ipynb | 52 | 6 | entropy=5.00 len=172 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | demo.ipynb | 53 | 6 | entropy=5.02 len=172 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | demo.ipynb | 163 | 6 | entropy=5.12 len=332 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | demo.ipynb | 226 | 6 | entropy=5.28 len=166 type=high_entropy | entropy | heuristic |

## Errors

No errors recorded.

## Affected files

- `/tmp/safeanalyze-BIPIA/benchmark/qa/process.py`
- `SECURITY.md`
- `benchmark/code/test.jsonl`
- `benchmark/code/train.jsonl`
- `benchmark/code_attack_test.json`
- `benchmark/code_attack_train.json`
- `benchmark/email/test.jsonl`
- `benchmark/email/train.jsonl`
- `benchmark/table/train.jsonl`
- `bipia/metrics/eval/utils.py`
- `bipia/metrics/regist.py`
- `bipia/model/llm_worker.py`
- `bipia/model/vllm_worker.py`
- `config/gpt35.yaml`
- `config/gpt35_wosys.yaml`
- `config/gpt4.yaml`
- `config/gpt4_wosys.yaml`
- `defense/README.md`
- `defense/white_box/utils.py`
- `demo.ipynb`

