# safeanalyze Report: /tmp/safeanalyze-InjecAgent

- **Started:** 2026-07-15 01:08:20 UTC
- **Duration:** 4187 ms
- **Total findings:** 2620

## Summary

| Metric | Value |
| --- | --- |
| Files scanned | 22 |
| Files sanitized | 0 |
| Bytes before | 0 |
| Bytes after | 0 |
| Total findings | 2620 |
| Errors | 0 |

### Findings by severity

| Severity | Count |
| --- | --- |
| Critical | 2129 |
| High | 180 |
| Medium | 3 |
| Low | 308 |

### Findings by source

| Source | Count |
| --- | --- |
| entropy | 311 |
| yara | 2309 |

## Findings

### Critical (2129)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| prompt_injection_comment | README.md | 114 | 14 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | README.md | 114 | 51 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 540 | 60 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 540 | 97 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 961 | 62 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 961 | 99 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 1411 | 73 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 1411 | 110 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 1584 | 65 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 1584 | 102 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 1767 | 135 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 1767 | 172 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 2079 | 65 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 2079 | 102 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 2196 | 67 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 2196 | 104 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 2226 | 79 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 2226 | 116 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 2300 | 52 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/attacker_simulated_responses.json | 2300 | 89 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 15 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 15 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 30 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 30 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 45 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 45 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 60 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 60 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 75 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 75 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 90 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 90 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 105 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 105 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 120 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 120 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 135 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 135 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 150 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 150 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 165 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 165 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 180 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 180 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 195 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 195 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 210 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 210 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 225 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 225 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 240 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 240 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 255 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 255 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 270 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 270 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 285 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 285 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 300 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 300 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 315 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 315 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 330 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 330 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 345 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 345 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 360 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 360 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 375 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 375 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 390 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 390 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 405 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 405 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 420 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 420 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 435 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 435 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 450 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 450 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 465 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 465 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 480 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 480 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 495 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 495 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 510 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 510 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 525 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 525 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 540 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 540 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 555 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 555 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 570 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 570 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 585 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 585 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 600 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 600 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 615 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 615 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 630 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 630 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 645 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 645 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 660 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 660 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 675 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 675 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 690 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 690 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 705 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 705 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 720 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 720 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 735 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 735 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 750 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 750 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 765 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 765 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 780 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 780 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 795 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 795 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 810 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 810 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 825 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 825 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 840 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 840 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 855 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 855 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 870 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 870 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 885 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 885 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 900 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 900 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 915 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 915 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 930 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 930 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 945 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 945 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 960 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 960 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 975 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 975 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 990 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 990 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1005 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1005 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1020 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1020 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1035 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1035 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1050 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1050 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1065 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1065 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1080 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1080 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1095 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1095 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1110 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1110 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1125 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1125 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1140 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1140 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1155 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1155 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1170 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1170 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1185 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1185 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1200 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1200 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1215 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1215 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1230 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1230 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1245 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1245 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1260 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1260 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1275 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1275 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1290 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1290 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1305 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1305 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1320 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1320 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1335 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1335 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1350 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1350 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1365 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1365 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1380 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1380 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1395 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1395 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1410 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1410 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1425 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1425 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1440 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1440 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1455 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1455 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1470 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1470 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1485 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1485 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1500 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1500 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1515 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1515 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1530 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1530 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1545 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1545 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1560 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1560 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1575 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1575 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1590 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1590 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1605 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1605 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1620 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1620 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1635 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1635 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1650 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1650 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1665 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1665 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1680 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1680 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1695 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1695 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1710 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1710 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1725 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1725 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1740 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1740 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1755 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1755 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1770 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1770 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1785 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1785 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1800 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1800 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1815 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1815 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1830 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1830 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1845 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1845 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1860 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1860 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1875 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1875 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1890 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1890 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1905 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1905 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1920 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1920 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1935 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1935 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1950 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1950 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1965 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1965 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1980 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1980 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1995 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 1995 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2010 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2010 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2025 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2025 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2040 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2040 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2055 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2055 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2070 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2070 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2085 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2085 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2100 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2100 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2115 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2115 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2130 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2130 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2145 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2145 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2160 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2160 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2175 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2175 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2190 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2190 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2205 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2205 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2220 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2220 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2235 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2235 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2250 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2250 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2265 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2265 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2280 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2280 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2295 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2295 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2310 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2310 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2325 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2325 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2340 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2340 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2355 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2355 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2370 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2370 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2385 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2385 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2400 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2400 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2415 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2415 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2430 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2430 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2445 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2445 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2460 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2460 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2475 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2475 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2490 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2490 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2505 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2505 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2520 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2520 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2535 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2535 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2550 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2550 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2565 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2565 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2580 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2580 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2595 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2595 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2610 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2610 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2625 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2625 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2640 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2640 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2655 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2655 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2670 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2670 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2685 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2685 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2700 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2700 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2715 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2715 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2730 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2730 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2745 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2745 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2760 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2760 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2775 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2775 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2790 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2790 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2805 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2805 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2820 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2820 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2835 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2835 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2850 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2850 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2865 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2865 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2880 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2880 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2895 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2895 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2910 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2910 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2925 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2925 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2940 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2940 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2955 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2955 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2970 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2970 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2985 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 2985 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3000 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3000 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3015 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3015 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3030 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3030 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3045 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3045 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3060 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3060 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3075 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3075 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3090 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3090 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3105 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3105 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3120 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3120 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3135 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3135 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3150 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3150 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3165 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3165 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3180 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3180 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3195 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3195 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3210 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3210 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3225 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3225 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3240 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3240 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3255 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3255 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3270 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3270 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3285 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3285 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3300 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3300 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3315 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3315 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3330 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3330 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3345 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3345 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3360 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3360 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3375 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3375 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3390 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3390 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3405 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3405 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3420 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3420 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3435 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3435 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3450 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3450 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3465 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3465 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3480 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3480 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3495 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3495 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3510 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3510 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3525 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3525 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3540 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3540 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3555 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3555 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3570 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3570 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3585 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3585 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3600 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3600 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3615 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3615 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3630 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3630 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3645 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3645 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3660 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3660 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3675 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3675 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3690 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3690 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3705 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3705 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3720 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3720 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3735 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3735 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3750 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3750 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3765 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3765 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3780 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3780 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3795 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3795 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3810 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3810 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3825 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3825 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3840 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3840 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3855 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3855 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3870 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3870 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3885 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3885 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3900 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3900 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3915 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3915 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3930 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3930 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3945 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3945 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3960 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3960 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3975 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3975 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3990 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 3990 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4005 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4005 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4020 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4020 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4035 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4035 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4050 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4050 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4065 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4065 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4080 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4080 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4095 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4095 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4110 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4110 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4125 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4125 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4140 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4140 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4155 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4155 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4170 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4170 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4185 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4185 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4200 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4200 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4215 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4215 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4230 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4230 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4245 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4245 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4260 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4260 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4275 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4275 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4290 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4290 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4305 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4305 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4320 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4320 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4335 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4335 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4350 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4350 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4365 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4365 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4380 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4380 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4395 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4395 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4410 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4410 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4425 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4425 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4440 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4440 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4455 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4455 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4470 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4470 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4485 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4485 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4500 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4500 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4515 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4515 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4530 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4530 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4545 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4545 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4560 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4560 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4575 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4575 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4590 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4590 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4605 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4605 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4620 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4620 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4635 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4635 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4650 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4650 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4665 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4665 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4680 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4680 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4695 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4695 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4710 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4710 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4725 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4725 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4740 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4740 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4755 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4755 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4770 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4770 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4785 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4785 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4800 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4800 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4815 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4815 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4830 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4830 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4845 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4845 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4860 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4860 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4875 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4875 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4890 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4890 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4905 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4905 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4920 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4920 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4935 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4935 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4950 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4950 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4965 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4965 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4980 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4980 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4995 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 4995 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5010 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5010 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5025 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5025 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5040 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5040 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5055 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5055 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5070 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5070 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5085 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5085 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5100 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5100 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5115 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5115 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5130 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5130 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5145 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5145 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5160 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5160 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5175 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5175 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5190 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5190 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5205 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5205 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5220 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5220 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5235 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5235 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5250 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5250 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5265 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5265 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5280 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5280 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5295 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5295 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5310 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5310 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5325 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5325 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5340 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5340 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5355 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5355 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5370 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5370 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5385 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5385 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5400 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5400 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5415 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5415 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5430 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5430 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5445 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5445 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5460 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5460 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5475 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5475 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5490 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5490 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5505 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5505 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5520 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5520 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5535 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5535 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5550 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5550 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5565 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5565 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5580 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5580 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5595 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5595 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5610 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5610 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5625 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5625 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5640 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5640 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5655 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5655 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5670 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5670 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5685 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5685 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5700 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5700 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5715 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5715 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5730 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5730 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5745 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5745 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5760 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5760 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5775 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5775 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5790 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5790 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5805 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5805 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5820 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5820 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5835 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5835 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5850 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5850 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5865 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5865 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5880 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5880 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5895 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5895 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5910 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5910 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5925 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5925 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5940 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5940 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5955 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5955 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5970 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5970 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5985 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 5985 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6000 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6000 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6015 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6015 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6030 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6030 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6045 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6045 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6060 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6060 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6075 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6075 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6090 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6090 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6105 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6105 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6120 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6120 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6135 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6135 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6150 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6150 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6165 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6165 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6180 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6180 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6195 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6195 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6210 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6210 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6225 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6225 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6240 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6240 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6255 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6255 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6270 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6270 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6285 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6285 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6300 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6300 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6315 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6315 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6330 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6330 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6345 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6345 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6360 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6360 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6375 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6375 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6390 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6390 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6405 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6405 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6420 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6420 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6435 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6435 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6450 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6450 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6465 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6465 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6480 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6480 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6495 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6495 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6510 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6510 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6525 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6525 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6540 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6540 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6555 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6555 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6570 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6570 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6585 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6585 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6600 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6600 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6615 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6615 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6630 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6630 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6645 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6645 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6660 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6660 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6675 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6675 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6690 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6690 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6705 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6705 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6720 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6720 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6735 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6735 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6750 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6750 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6765 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6765 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6780 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6780 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6795 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6795 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6810 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6810 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6825 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6825 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6840 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6840 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6855 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6855 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6870 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6870 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6885 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6885 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6900 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6900 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6915 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6915 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6930 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6930 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6945 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6945 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6960 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6960 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6975 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6975 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6990 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 6990 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7005 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7005 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7020 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7020 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7035 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7035 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7050 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7050 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7065 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7065 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7080 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7080 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7095 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7095 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7110 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7110 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7125 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7125 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7140 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7140 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7155 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7155 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7170 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7170 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7185 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7185 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7200 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7200 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7215 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7215 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7230 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7230 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7245 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7245 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7260 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7260 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7275 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7275 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7290 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7290 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7305 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7305 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7320 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7320 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7335 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7335 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7350 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7350 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7365 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7365 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7380 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7380 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7395 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7395 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7410 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7410 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7425 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7425 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7440 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7440 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7455 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7455 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7470 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7470 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7485 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7485 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7500 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7500 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7515 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7515 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7530 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7530 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7545 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7545 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7560 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7560 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7575 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7575 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7590 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7590 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7605 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7605 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7620 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7620 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7635 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7635 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7650 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_dh_enhanced.json | 7650 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 16 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 16 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 32 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 32 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 48 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 48 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 64 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 64 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 80 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 80 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 96 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 96 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 112 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 112 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 128 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 128 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 144 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 144 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 160 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 160 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 176 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 176 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 192 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 192 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 208 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 208 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 224 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 224 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 240 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 240 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 256 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 256 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 272 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 272 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 288 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 288 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 304 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 304 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 320 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 320 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 336 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 336 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 352 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 352 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 368 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 368 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 384 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 384 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 400 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 400 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 416 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 416 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 432 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 432 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 448 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 448 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 464 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 464 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 480 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 480 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 496 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 496 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 512 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 512 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 528 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 528 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 544 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 544 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 560 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 560 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 576 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 576 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 592 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 592 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 608 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 608 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 624 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 624 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 640 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 640 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 656 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 656 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 672 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 672 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 688 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 688 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 704 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 704 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 720 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 720 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 736 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 736 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 752 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 752 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 768 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 768 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 784 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 784 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 800 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 800 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 816 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 816 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 832 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 832 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 848 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 848 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 864 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 864 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 880 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 880 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 896 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 896 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 912 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 912 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 928 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 928 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 944 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 944 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 960 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 960 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 976 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 976 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 992 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 992 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1008 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1008 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1024 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1024 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1040 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1040 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1056 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1056 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1072 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1072 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1088 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1088 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1104 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1104 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1120 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1120 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1136 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1136 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1152 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1152 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1168 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1168 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1184 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1184 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1200 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1200 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1216 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1216 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1232 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1232 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1248 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1248 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1264 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1264 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1280 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1280 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1296 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1296 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1312 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1312 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1328 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1328 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1344 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1344 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1360 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1360 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1376 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1376 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1392 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1392 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1408 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1408 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1424 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1424 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1440 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1440 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1456 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1456 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1472 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1472 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1488 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1488 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1504 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1504 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1520 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1520 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1536 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1536 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1552 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1552 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1568 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1568 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1584 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1584 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1600 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1600 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1616 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1616 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1632 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1632 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1648 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1648 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1664 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1664 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1680 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1680 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1696 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1696 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1712 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1712 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1728 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1728 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1744 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1744 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1760 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1760 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1776 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1776 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1792 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1792 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1808 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1808 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1824 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1824 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1840 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1840 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1856 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1856 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1872 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1872 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1888 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1888 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1904 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1904 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1920 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1920 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1936 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1936 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1952 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1952 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1968 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1968 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1984 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 1984 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2000 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2000 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2016 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2016 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2032 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2032 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2048 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2048 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2064 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2064 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2080 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2080 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2096 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2096 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2112 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2112 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2128 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2128 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2144 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2144 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2160 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2160 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2176 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2176 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2192 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2192 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2208 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2208 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2224 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2224 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2240 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2240 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2256 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2256 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2272 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2272 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2288 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2288 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2304 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2304 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2320 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2320 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2336 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2336 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2352 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2352 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2368 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2368 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2384 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2384 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2400 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2400 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2416 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2416 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2432 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2432 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2448 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2448 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2464 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2464 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2480 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2480 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2496 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2496 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2512 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2512 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2528 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2528 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2544 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2544 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2560 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2560 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2576 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2576 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2592 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2592 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2608 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2608 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2624 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2624 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2640 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2640 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2656 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2656 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2672 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2672 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2688 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2688 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2704 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2704 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2720 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2720 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2736 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2736 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2752 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2752 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2768 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2768 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2784 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2784 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2800 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2800 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2816 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2816 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2832 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2832 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2848 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2848 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2864 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2864 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2880 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2880 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2896 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2896 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2912 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2912 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2928 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2928 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2944 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2944 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2960 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2960 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2976 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2976 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2992 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 2992 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3008 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3008 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3024 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3024 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3040 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3040 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3056 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3056 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3072 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3072 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3088 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3088 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3104 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3104 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3120 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3120 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3136 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3136 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3152 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3152 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3168 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3168 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3184 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3184 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3200 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3200 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3216 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3216 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3232 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3232 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3248 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3248 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3264 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3264 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3280 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3280 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3296 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3296 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3312 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3312 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3328 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3328 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3344 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3344 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3360 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3360 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3376 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3376 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3392 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3392 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3408 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3408 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3424 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3424 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3440 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3440 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3456 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3456 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3472 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3472 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3488 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3488 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3504 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3504 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3520 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3520 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3536 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3536 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3552 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3552 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3568 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3568 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3584 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3584 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3600 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3600 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3616 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3616 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3632 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3632 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3648 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3648 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3664 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3664 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3680 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3680 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3696 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3696 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3712 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3712 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3728 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3728 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3744 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3744 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3760 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3760 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3776 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3776 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3792 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3792 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3808 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3808 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3824 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3824 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3840 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3840 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3856 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3856 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3872 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3872 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3888 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3888 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3904 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3904 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3920 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3920 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3936 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3936 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3952 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3952 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3968 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3968 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3984 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 3984 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4000 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4000 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4016 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4016 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4032 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4032 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4048 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4048 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4064 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4064 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4080 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4080 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4096 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4096 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4112 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4112 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4128 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4128 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4144 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4144 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4160 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4160 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4176 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4176 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4192 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4192 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4208 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4208 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4224 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4224 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4240 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4240 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4256 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4256 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4272 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4272 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4288 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4288 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4304 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4304 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4320 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4320 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4336 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4336 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4352 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4352 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4368 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4368 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4384 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4384 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4400 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4400 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4416 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4416 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4432 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4432 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4448 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4448 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4464 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4464 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4480 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4480 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4496 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4496 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4512 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4512 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4528 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4528 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4544 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4544 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4560 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4560 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4576 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4576 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4592 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4592 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4608 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4608 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4624 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4624 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4640 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4640 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4656 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4656 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4672 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4672 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4688 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4688 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4704 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4704 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4720 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4720 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4736 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4736 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4752 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4752 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4768 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4768 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4784 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4784 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4800 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4800 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4816 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4816 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4832 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4832 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4848 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4848 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4864 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4864 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4880 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4880 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4896 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4896 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4912 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4912 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4928 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4928 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4944 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4944 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4960 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4960 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4976 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4976 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4992 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 4992 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5008 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5008 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5024 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5024 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5040 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5040 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5056 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5056 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5072 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5072 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5088 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5088 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5104 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5104 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5120 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5120 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5136 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5136 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5152 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5152 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5168 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5168 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5184 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5184 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5200 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5200 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5216 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5216 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5232 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5232 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5248 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5248 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5264 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5264 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5280 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5280 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5296 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5296 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5312 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5312 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5328 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5328 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5344 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5344 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5360 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5360 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5376 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5376 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5392 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5392 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5408 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5408 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5424 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5424 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5440 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5440 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5456 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5456 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5472 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5472 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5488 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5488 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5504 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5504 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5520 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5520 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5536 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5536 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5552 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5552 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5568 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5568 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5584 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5584 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5600 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5600 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5616 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5616 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5632 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5632 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5648 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5648 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5664 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5664 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5680 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5680 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5696 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5696 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5712 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5712 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5728 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5728 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5744 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5744 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5760 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5760 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5776 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5776 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5792 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5792 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5808 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5808 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5824 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5824 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5840 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5840 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5856 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5856 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5872 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5872 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5888 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5888 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5904 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5904 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5920 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5920 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5936 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5936 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5952 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5952 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5968 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5968 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5984 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 5984 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6000 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6000 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6016 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6016 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6032 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6032 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6048 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6048 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6064 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6064 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6080 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6080 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6096 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6096 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6112 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6112 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6128 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6128 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6144 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6144 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6160 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6160 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6176 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6176 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6192 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6192 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6208 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6208 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6224 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6224 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6240 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6240 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6256 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6256 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6272 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6272 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6288 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6288 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6304 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6304 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6320 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6320 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6336 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6336 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6352 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6352 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6368 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6368 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6384 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6384 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6400 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6400 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6416 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6416 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6432 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6432 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6448 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6448 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6464 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6464 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6480 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6480 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6496 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6496 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6512 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6512 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6528 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6528 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6544 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6544 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6560 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6560 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6576 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6576 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6592 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6592 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6608 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6608 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6624 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6624 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6640 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6640 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6656 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6656 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6672 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6672 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6688 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6688 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6704 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6704 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6720 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6720 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6736 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6736 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6752 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6752 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6768 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6768 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6784 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6784 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6800 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6800 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6816 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6816 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6832 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6832 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6848 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6848 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6864 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6864 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6880 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6880 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6896 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6896 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6912 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6912 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6928 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6928 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6944 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6944 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6960 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6960 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6976 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6976 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6992 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 6992 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7008 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7008 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7024 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7024 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7040 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7040 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7056 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7056 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7072 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7072 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7088 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7088 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7104 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7104 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7120 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7120 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7136 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7136 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7152 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7152 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7168 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7168 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7184 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7184 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7200 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7200 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7216 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7216 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7232 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7232 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7248 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7248 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7264 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7264 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7280 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7280 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7296 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7296 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7312 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7312 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7328 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7328 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7344 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7344 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7360 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7360 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7376 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7376 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7392 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7392 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7408 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7408 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7424 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7424 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7440 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7440 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7456 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7456 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7472 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7472 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7488 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7488 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7504 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7504 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7520 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7520 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7536 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7536 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7552 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7552 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7568 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7568 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7584 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7584 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7600 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7600 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7616 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7616 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7632 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7632 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7648 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7648 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7664 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7664 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7680 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7680 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7696 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7696 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7712 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7712 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7728 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7728 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7744 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7744 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7760 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7760 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7776 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7776 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7792 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7792 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7808 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7808 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7824 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7824 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7840 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7840 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7856 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7856 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7872 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7872 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7888 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7888 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7904 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7904 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7920 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7920 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7936 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7936 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7952 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7952 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7968 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7968 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7984 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 7984 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8000 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8000 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8016 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8016 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8032 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8032 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8048 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8048 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8064 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8064 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8080 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8080 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8096 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8096 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8112 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8112 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8128 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8128 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8144 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8144 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8160 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8160 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8176 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8176 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8192 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8192 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8208 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8208 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8224 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8224 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8240 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8240 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8256 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8256 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8272 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8272 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8288 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8288 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8304 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8304 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8320 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8320 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8336 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8336 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8352 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8352 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8368 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8368 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8384 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8384 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8400 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8400 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8416 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8416 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8432 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8432 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8448 | 267 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8448 | 304 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8464 | 114 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8464 | 151 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8480 | 108 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8480 | 145 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8496 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8496 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8512 | 136 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8512 | 173 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8528 | 179 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8528 | 216 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8544 | 217 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8544 | 254 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8560 | 70 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8560 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8576 | 143 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8576 | 180 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8592 | 85 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8592 | 122 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8608 | 91 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8608 | 128 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8624 | 115 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8624 | 152 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8640 | 126 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8640 | 163 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8656 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8656 | 181 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8672 | 138 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8672 | 175 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8688 | 107 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8688 | 144 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8704 | 49 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | data/test_cases_ds_enhanced.json | 8704 | 86 | Comment containing prompt injection keywords | yara | deterministic |
| prompt_injection_comment | src/prompts/prompt_template.py | 17 | 42 | Comment containing prompt injection keywords | yara | deterministic |

### High (180)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| obfuscated_javascript | data/attacker_cases_dh.jsonl | 3 | 289 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_cases_ds.jsonl | 14 | 269 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_cases_ds.jsonl | 29 | 398 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_cases_ds.jsonl | 29 | 450 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 9 | 235 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 169 | 222 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 558 | 234 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 865 | 249 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 871 | 304 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1046 | 256 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1124 | 281 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1496 | 1120 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1688 | 52 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1688 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1688 | 64 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1688 | 70 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1721 | 446 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1767 | 107 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1767 | 113 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1767 | 119 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1767 | 125 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1773 | 95 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1773 | 101 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1773 | 107 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1831 | 94 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1831 | 100 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 103 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 109 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 180 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 186 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 215 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 221 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 227 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 233 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 239 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 273 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1833 | 279 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1834 | 72 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1834 | 78 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1834 | 84 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1834 | 90 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1905 | 109 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1938 | 167 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/attacker_simulated_responses.json | 1938 | 173 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 518 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 533 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 548 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 563 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 578 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 593 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 608 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 623 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 638 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 653 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 668 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 683 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 698 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 713 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 728 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 743 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_base.json | 758 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 518 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 533 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 548 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 563 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 578 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 593 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 608 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 623 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 638 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 653 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 668 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 683 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 698 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 713 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 728 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 743 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_dh_enhanced.json | 758 | 57 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3545 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3561 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3577 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3593 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3609 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3625 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3641 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3657 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3673 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3689 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3705 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3721 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3737 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3753 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3769 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3785 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 3801 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7625 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7625 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7641 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7641 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7657 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7657 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7673 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7673 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7689 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7689 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7705 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7705 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7721 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7721 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7737 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7737 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7753 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7753 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7769 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7769 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7785 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7785 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7801 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7801 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7817 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7817 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7833 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7833 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7849 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7849 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7865 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7865 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7881 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_base.json | 7881 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3545 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3561 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3577 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3593 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3609 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3625 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3641 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3657 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3673 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3689 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3705 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3721 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3737 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3753 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3769 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3785 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 3801 | 58 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7625 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7625 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7641 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7641 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7657 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7657 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7673 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7673 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7689 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7689 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7705 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7705 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7721 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7721 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7737 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7737 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7753 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7753 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7769 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7769 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7785 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7785 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7801 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7801 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7817 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7817 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7833 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7833 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7849 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7849 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7865 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7865 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7881 | 82 | Obfuscated or packed JavaScript patterns | yara | deterministic |
| obfuscated_javascript | data/test_cases_ds_enhanced.json | 7881 | 134 | Obfuscated or packed JavaScript patterns | yara | deterministic |

### Medium (3)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| entropy_base64_blob | data/attacker_simulated_responses.json | 893 | 90 | entropy=4.60 len=52 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | data/attacker_simulated_responses.json | 1139 | 68 | entropy=5.00 len=75 type=base64_blob | entropy | heuristic |
| entropy_base64_blob | data/attacker_simulated_responses.json | 2117 | 90 | entropy=4.68 len=60 type=base64_blob | entropy | heuristic |

### Low (308)

| Rule | File | Line | Column | Message | Source | Confidence |
| --- | --- | --- | --- | --- | --- | --- |
| entropy_high_entropy | data/attacker_simulated_responses.json | 14 | 1 | entropy=5.07 len=60 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 46 | 1 | entropy=5.06 len=62 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 62 | 1 | entropy=5.02 len=88 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 85 | 1 | entropy=5.19 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 145 | 96 | entropy=5.01 len=208 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 180 | 1 | entropy=5.01 len=84 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 182 | 1 | entropy=5.02 len=111 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 184 | 1 | entropy=5.03 len=80 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 185 | 1 | entropy=5.06 len=89 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 186 | 1 | entropy=5.14 len=62 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 187 | 1 | entropy=5.07 len=127 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 217 | 1 | entropy=5.07 len=82 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 218 | 1 | entropy=5.18 len=126 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 241 | 1 | entropy=5.04 len=58 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 279 | 1 | entropy=5.23 len=157 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 285 | 223 | entropy=5.00 len=778 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 308 | 1 | entropy=5.02 len=111 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 310 | 1 | entropy=5.08 len=121 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 385 | 91 | entropy=5.01 len=346 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 401 | 1 | entropy=5.02 len=61 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 408 | 1 | entropy=5.14 len=60 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 408 | 65 | entropy=5.05 len=114 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 411 | 1 | entropy=5.12 len=62 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 571 | 1 | entropy=5.02 len=92 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 576 | 1 | entropy=5.03 len=89 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 579 | 1 | entropy=5.09 len=91 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 580 | 1 | entropy=5.15 len=91 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 581 | 1 | entropy=5.07 len=62 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 582 | 1 | entropy=5.13 len=151 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 583 | 1 | entropy=5.18 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 584 | 1 | entropy=5.05 len=56 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 586 | 1 | entropy=5.03 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 587 | 1 | entropy=5.05 len=163 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 587 | 168 | entropy=5.05 len=93 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 588 | 1 | entropy=5.17 len=84 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 589 | 136 | entropy=5.03 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 661 | 1 | entropy=5.01 len=95 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 726 | 1 | entropy=5.06 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 784 | 1 | entropy=5.13 len=81 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 785 | 1 | entropy=5.11 len=86 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 787 | 1 | entropy=5.05 len=58 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 881 | 1 | entropy=5.02 len=63 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 882 | 1 | entropy=5.03 len=60 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 883 | 1 | entropy=5.19 len=66 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 893 | 1 | entropy=5.05 len=85 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 894 | 1 | entropy=5.04 len=145 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 895 | 1 | entropy=5.08 len=89 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 896 | 1 | entropy=5.07 len=78 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 896 | 83 | entropy=5.03 len=126 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 899 | 76 | entropy=5.08 len=407 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 914 | 1 | entropy=5.00 len=80 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 925 | 64 | entropy=5.09 len=324 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 980 | 1 | entropy=5.03 len=94 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 981 | 1 | entropy=5.02 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1093 | 1 | entropy=5.09 len=60 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1136 | 1 | entropy=5.03 len=62 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1163 | 88 | entropy=5.10 len=729 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1164 | 1 | entropy=5.02 len=56 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1189 | 1 | entropy=5.00 len=58 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1190 | 1 | entropy=5.03 len=56 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1232 | 1 | entropy=5.02 len=133 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1259 | 1 | entropy=5.07 len=60 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1260 | 1 | entropy=5.02 len=62 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1261 | 1 | entropy=5.02 len=66 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1274 | 69 | entropy=5.03 len=402 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1291 | 1 | entropy=5.08 len=60 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1334 | 1 | entropy=5.01 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1366 | 1 | entropy=5.11 len=114 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1367 | 1 | entropy=5.16 len=124 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1370 | 1 | entropy=5.06 len=132 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1447 | 1 | entropy=5.15 len=97 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1448 | 1 | entropy=5.13 len=80 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1478 | 1 | entropy=5.02 len=61 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1672 | 1 | entropy=5.16 len=58 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1791 | 1 | entropy=5.12 len=74 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1793 | 1 | entropy=5.03 len=80 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1856 | 1 | entropy=5.13 len=64 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1889 | 1 | entropy=5.12 len=78 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1951 | 1 | entropy=5.00 len=181 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1956 | 1 | entropy=5.01 len=72 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 1965 | 132 | entropy=5.04 len=279 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 2157 | 1 | entropy=5.10 len=92 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 2204 | 1 | entropy=5.15 len=60 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/attacker_simulated_responses.json | 2323 | 1 | entropy=5.03 len=167 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 105 | 19 | entropy=5.02 len=282 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 450 | 19 | entropy=5.01 len=162 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 525 | 19 | entropy=5.07 len=396 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 540 | 19 | entropy=5.08 len=242 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 570 | 19 | entropy=5.05 len=270 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 615 | 19 | entropy=5.17 len=345 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 630 | 19 | entropy=5.05 len=264 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 645 | 19 | entropy=5.01 len=475 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 660 | 19 | entropy=5.07 len=330 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 675 | 19 | entropy=5.03 len=219 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 690 | 19 | entropy=5.03 len=307 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 705 | 19 | entropy=5.13 len=290 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 720 | 19 | entropy=5.01 len=271 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 735 | 19 | entropy=5.10 len=298 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 870 | 19 | entropy=5.03 len=277 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 960 | 19 | entropy=5.02 len=222 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 1125 | 19 | entropy=5.04 len=265 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 1215 | 19 | entropy=5.06 len=210 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 1470 | 19 | entropy=5.04 len=186 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 2655 | 19 | entropy=5.02 len=326 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 2745 | 19 | entropy=5.03 len=271 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 2910 | 19 | entropy=5.01 len=285 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 2925 | 19 | entropy=5.00 len=204 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 3000 | 19 | entropy=5.06 len=230 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 3165 | 19 | entropy=5.01 len=360 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 3255 | 19 | entropy=5.03 len=305 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 3420 | 19 | entropy=5.04 len=319 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 3435 | 19 | entropy=5.05 len=238 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 3510 | 19 | entropy=5.09 len=264 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 4020 | 19 | entropy=5.01 len=245 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 4185 | 19 | entropy=5.02 len=281 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 4275 | 19 | entropy=5.07 len=226 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 4440 | 19 | entropy=5.01 len=271 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 4530 | 19 | entropy=5.06 len=216 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 5460 | 19 | entropy=5.01 len=262 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 5550 | 19 | entropy=5.05 len=207 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 6225 | 19 | entropy=5.01 len=259 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 6315 | 19 | entropy=5.05 len=204 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 6990 | 19 | entropy=5.02 len=306 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 7425 | 19 | entropy=5.03 len=163 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 7500 | 19 | entropy=5.09 len=266 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 7515 | 19 | entropy=5.02 len=185 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 7590 | 19 | entropy=5.12 len=211 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_base.json | 7620 | 19 | entropy=5.05 len=219 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 105 | 19 | entropy=5.08 len=378 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 120 | 19 | entropy=5.04 len=297 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 195 | 19 | entropy=5.07 len=323 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 225 | 19 | entropy=5.01 len=331 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 360 | 19 | entropy=5.06 len=313 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 375 | 19 | entropy=5.04 len=232 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 450 | 19 | entropy=5.11 len=258 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 525 | 19 | entropy=5.10 len=492 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 540 | 19 | entropy=5.12 len=338 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 570 | 19 | entropy=5.10 len=366 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 585 | 19 | entropy=5.02 len=438 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 600 | 19 | entropy=5.04 len=497 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 615 | 19 | entropy=5.21 len=441 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 630 | 19 | entropy=5.15 len=360 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 645 | 19 | entropy=5.08 len=571 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 660 | 19 | entropy=5.12 len=426 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 675 | 19 | entropy=5.10 len=315 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 690 | 19 | entropy=5.10 len=403 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 705 | 19 | entropy=5.20 len=386 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 720 | 19 | entropy=5.08 len=367 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 735 | 19 | entropy=5.15 len=394 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 750 | 19 | entropy=5.09 len=364 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 765 | 19 | entropy=5.05 len=271 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 870 | 19 | entropy=5.11 len=373 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 885 | 19 | entropy=5.06 len=292 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 915 | 19 | entropy=5.01 len=358 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 945 | 19 | entropy=5.00 len=335 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 960 | 19 | entropy=5.11 len=318 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 990 | 19 | entropy=5.05 len=326 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1050 | 19 | entropy=5.00 len=258 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1125 | 19 | entropy=5.13 len=361 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1140 | 19 | entropy=5.10 len=280 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1200 | 19 | entropy=5.00 len=323 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1215 | 19 | entropy=5.15 len=306 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1245 | 19 | entropy=5.07 len=314 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1260 | 19 | entropy=5.02 len=284 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1380 | 19 | entropy=5.07 len=337 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1395 | 19 | entropy=5.06 len=256 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1470 | 19 | entropy=5.11 len=282 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1500 | 19 | entropy=5.02 len=290 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1635 | 19 | entropy=5.10 len=337 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1650 | 19 | entropy=5.07 len=256 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1725 | 19 | entropy=5.13 len=282 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 1755 | 19 | entropy=5.05 len=290 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 2400 | 19 | entropy=5.01 len=374 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 2490 | 19 | entropy=5.05 len=319 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 2655 | 19 | entropy=5.07 len=422 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 2670 | 19 | entropy=5.04 len=341 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 2745 | 19 | entropy=5.08 len=367 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 2775 | 19 | entropy=5.02 len=375 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 2910 | 19 | entropy=5.06 len=381 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 2925 | 19 | entropy=5.07 len=300 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3000 | 19 | entropy=5.10 len=326 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3165 | 19 | entropy=5.06 len=456 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3180 | 19 | entropy=5.03 len=375 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3255 | 19 | entropy=5.07 len=401 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3285 | 19 | entropy=5.00 len=409 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3420 | 19 | entropy=5.08 len=415 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3435 | 19 | entropy=5.09 len=334 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3510 | 19 | entropy=5.12 len=360 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3540 | 19 | entropy=5.03 len=368 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3555 | 19 | entropy=5.01 len=338 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3765 | 19 | entropy=5.01 len=374 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3930 | 19 | entropy=5.03 len=396 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 3945 | 19 | entropy=5.03 len=315 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4020 | 19 | entropy=5.09 len=341 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4050 | 19 | entropy=5.02 len=349 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4185 | 19 | entropy=5.08 len=377 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4200 | 19 | entropy=5.06 len=296 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4275 | 19 | entropy=5.13 len=322 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4305 | 19 | entropy=5.02 len=330 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4440 | 19 | entropy=5.08 len=367 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4455 | 19 | entropy=5.07 len=286 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4530 | 19 | entropy=5.12 len=312 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4560 | 19 | entropy=5.01 len=320 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4695 | 19 | entropy=5.02 len=392 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4710 | 19 | entropy=5.00 len=311 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4785 | 19 | entropy=5.05 len=337 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4950 | 19 | entropy=5.05 len=354 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 4965 | 19 | entropy=5.01 len=273 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 5040 | 19 | entropy=5.06 len=299 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 5460 | 19 | entropy=5.06 len=358 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 5475 | 19 | entropy=5.02 len=277 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 5550 | 19 | entropy=5.08 len=303 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 5970 | 19 | entropy=5.01 len=369 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6060 | 19 | entropy=5.04 len=314 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6225 | 19 | entropy=5.08 len=355 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6240 | 19 | entropy=5.07 len=274 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6315 | 19 | entropy=5.12 len=300 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6345 | 19 | entropy=5.03 len=308 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6480 | 19 | entropy=5.04 len=391 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6495 | 19 | entropy=5.01 len=310 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6570 | 19 | entropy=5.05 len=336 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6825 | 19 | entropy=5.01 len=319 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 6990 | 19 | entropy=5.08 len=402 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7005 | 19 | entropy=5.05 len=321 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7080 | 19 | entropy=5.09 len=347 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7110 | 19 | entropy=5.02 len=355 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7410 | 19 | entropy=5.04 len=413 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7425 | 19 | entropy=5.05 len=259 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7455 | 19 | entropy=5.03 len=287 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7485 | 19 | entropy=5.00 len=418 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7500 | 19 | entropy=5.15 len=362 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7515 | 19 | entropy=5.14 len=281 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7530 | 19 | entropy=5.03 len=492 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7545 | 19 | entropy=5.04 len=347 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7560 | 19 | entropy=5.02 len=236 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7575 | 19 | entropy=5.05 len=324 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7590 | 19 | entropy=5.19 len=307 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7605 | 19 | entropy=5.01 len=288 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7620 | 19 | entropy=5.11 len=315 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7635 | 19 | entropy=5.07 len=285 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_dh_enhanced.json | 7650 | 19 | entropy=5.00 len=192 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 2016 | 19 | entropy=5.01 len=397 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 2112 | 19 | entropy=5.01 len=342 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 2288 | 19 | entropy=5.00 len=371 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 2384 | 19 | entropy=5.03 len=316 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 2928 | 19 | entropy=5.02 len=293 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 3472 | 19 | entropy=5.03 len=199 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 3744 | 19 | entropy=5.04 len=241 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 4016 | 19 | entropy=5.03 len=236 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 4288 | 19 | entropy=5.00 len=293 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 4464 | 19 | entropy=5.00 len=300 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 4480 | 19 | entropy=5.02 len=219 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 4560 | 19 | entropy=5.08 len=245 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 5104 | 19 | entropy=5.00 len=242 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 5920 | 19 | entropy=5.05 len=321 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_base.json | 8096 | 19 | entropy=5.01 len=357 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 1024 | 19 | entropy=5.01 len=329 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 1472 | 19 | entropy=5.00 len=415 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 1568 | 19 | entropy=5.05 len=360 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 1840 | 19 | entropy=5.02 len=350 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2016 | 19 | entropy=5.06 len=493 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2032 | 19 | entropy=5.04 len=412 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2112 | 19 | entropy=5.07 len=438 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2144 | 19 | entropy=5.02 len=446 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2288 | 19 | entropy=5.06 len=467 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2304 | 19 | entropy=5.04 len=386 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2384 | 19 | entropy=5.08 len=412 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2416 | 19 | entropy=5.01 len=420 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2576 | 19 | entropy=5.01 len=295 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2656 | 19 | entropy=5.05 len=321 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2832 | 19 | entropy=5.02 len=444 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2848 | 19 | entropy=5.04 len=363 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 2928 | 19 | entropy=5.07 len=389 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3200 | 19 | entropy=5.01 len=417 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3376 | 19 | entropy=5.04 len=350 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3392 | 19 | entropy=5.07 len=269 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3472 | 19 | entropy=5.11 len=295 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3504 | 19 | entropy=5.02 len=303 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3648 | 19 | entropy=5.04 len=392 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3664 | 19 | entropy=5.06 len=311 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3744 | 19 | entropy=5.09 len=337 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3776 | 19 | entropy=5.01 len=345 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3920 | 19 | entropy=5.03 len=387 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 3936 | 19 | entropy=5.03 len=306 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4016 | 19 | entropy=5.08 len=332 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4192 | 19 | entropy=5.01 len=444 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4208 | 19 | entropy=5.02 len=363 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4288 | 19 | entropy=5.05 len=389 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4464 | 19 | entropy=5.06 len=396 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4480 | 19 | entropy=5.09 len=315 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4560 | 19 | entropy=5.12 len=341 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4592 | 19 | entropy=5.04 len=349 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4608 | 19 | entropy=5.01 len=319 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 4832 | 19 | entropy=5.04 len=338 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 5008 | 19 | entropy=5.00 len=393 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 5024 | 19 | entropy=5.02 len=312 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 5104 | 19 | entropy=5.06 len=338 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 5552 | 19 | entropy=5.02 len=377 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 5568 | 19 | entropy=5.04 len=296 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 5648 | 19 | entropy=5.08 len=322 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 5824 | 19 | entropy=5.04 len=472 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 5840 | 19 | entropy=5.04 len=391 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 5920 | 19 | entropy=5.08 len=417 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 6192 | 19 | entropy=5.02 len=385 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 8000 | 19 | entropy=5.01 len=508 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 8016 | 19 | entropy=5.01 len=427 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 8096 | 19 | entropy=5.03 len=453 type=high_entropy | entropy | heuristic |
| entropy_high_entropy | data/test_cases_ds_enhanced.json | 8368 | 19 | entropy=5.01 len=394 type=high_entropy | entropy | heuristic |

## Errors

No errors recorded.

## Affected files

- `README.md`
- `data/attacker_cases_dh.jsonl`
- `data/attacker_cases_ds.jsonl`
- `data/attacker_simulated_responses.json`
- `data/test_cases_dh_base.json`
- `data/test_cases_dh_enhanced.json`
- `data/test_cases_ds_base.json`
- `data/test_cases_ds_enhanced.json`
- `src/prompts/prompt_template.py`

