#!/usr/bin/env bash
# scripts/premortem.sh
# Before shipping an improvement, list the most likely failure modes.
# This is the "pre-mortem" step of the autoresearch loop described in CLAUDE.md.
set -euo pipefail

cat <<'EOF'
# Premortem — safeanalyze improvement

Imagine the next improvement has shipped and, one month later, the tool failed in production.
Answer the following honestly. The risks identified here become regression tests or guardrails.

## 1. What is the most likely reason a real prompt-injection payload was missed?

Examples:
- New rephrasing/encoding not covered by YARA rules.
- Binary/non-UTF8 payload bypassed text-based checks.
- ML model was disabled, not installed, or exceeded its timeout.
- External scanner crashed or was not installed.
- Long payload exceeded chunk/file-size limits and was truncated.

## 2. What is the most likely source of false-positive noise?

Examples:
- Test fixtures / legitimate documentation contain trigger phrases.
- Hidden-char rule fires on localization/emoji data.
- Entropy rule flags minified JS or lockfiles.
- Dependency code (node_modules/vendor) is scanned in thorough mode.

## 3. What latency or resource failure will hit first at scale?

Examples:
- Fast mode exceeds the 100 ms budget on large payloads.
- ML model loading uses >2 GB RAM and OOMs the Squid helper.
- External scanners serialize on one CPU core.
- Large repos generate reports larger than GitHub's push limit.

## 4. What single guardrail would have prevented the failure?

Pick the cheapest guardrail that addresses #1–#3 and add it before the improvement is released.

## 5. Regression test

Write one payload or repo scenario that would have caught the failure, then run:

    ./scripts/redteam.sh
    ./safeanalyze scan <test-target> --mode thorough

EOF
