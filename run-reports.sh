#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")" && pwd)"
OUT="$ROOT/reports/report-tactical-beaver-spork-cb4b88a9-2026-07-15"
SAFEANALYZE="$ROOT/safeanalyze"
CONFIG="/tmp/safeanalyze-test-config-v028.yaml"

mkdir -p "$OUT"

run_scan() {
    local name="$1"
    local target="$2"
    local dir="$OUT/$name"
    mkdir -p "$dir"
    echo "--- Scanning $name ($target) ---"
    local start end
    start=$(date +%s%N)
    if "$SAFEANALYZE" scan "$target" --mode thorough --config "$CONFIG" --output-dir "$dir" --verbose; then
        end=$(date +%s%N)
        echo "$(( (end-start)/1000000 ))" > "$dir/duration.txt"
    else
        end=$(date +%s%N)
        echo "$(( (end-start)/1000000 ))" > "$dir/duration.txt"
        echo "ERROR scanning $name" > "$dir/ERROR.txt"
    fi
}

run_scan repos/microsoft-bipia /tmp/safeanalyze-BIPIA
run_scan repos/uiuc-injecagent /tmp/safeanalyze-InjecAgent
run_scan repos/lakera-pint-benchmark /tmp/safeanalyze-pint-benchmark
run_scan repos/alexh-prompt-injection-scanner /tmp/safeanalyze-prompt-injection-scanner
run_scan repos/duriantaco-skylos /tmp/safeanalyze-skylos

# Docs are single HTML files; scan them as files.
run_scan docs/promptfoo-scenarios /tmp/safeanalyze-doc-scenarios.html
run_scan docs/promptfoo-webagents /tmp/safeanalyze-doc-webagents.html

echo "All scans complete. Outputs in $OUT"
