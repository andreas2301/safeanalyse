#!/usr/bin/env bash
# Reproduce the v0.3.42 (iteration 38) measurements in this report.
# Run from the repository root of a v0.3.42 checkout with eval data fetched
# (./scripts/fetch_eval.sh) and corpus targets present under /tmp/.
set -euo pipefail
export PATH=/usr/local/go/bin:$PATH
OUT="${OUT:-/tmp/safeanalyze-iter/iter38}"
CFG="${CFG:-/tmp/safeanalyze-test-config-eval.yaml}"
mkdir -p "$OUT/corpus"

go build -o safeanalyze .
go test ./...
./scripts/redteam.sh

# Labeled fast-mode eval (dev + holdout + report-only stress).
for ds in llmail browsesafe toolresp spanish multilingual deepset nemotron; do
    for s in "$ds" "$ds-holdout"; do
        ./safeanalyze eval "testdata/eval/$s.jsonl" \
            --json "$OUT/eval-$s.json" \
            --fp-out "$OUT/fp-$s.jsonl" --fn-out "$OUT/fn-$s.jsonl"
    done
done
./safeanalyze eval testdata/eval/promptprotection-stress.jsonl \
    --json "$OUT/eval-promptprotection-stress.json" \
    --fp-out "$OUT/fp-promptprotection-stress.jsonl" --fn-out "$OUT/fn-promptprotection-stress.jsonl"

# Thorough corpus scan (scan writes duration.txt itself since v0.3.34).
declare -A TARGETS=(
    [microsoft-bipia]=/tmp/safeanalyze-BIPIA
    [uiuc-injecagent]=/tmp/safeanalyze-InjecAgent
    [lakera-pint-benchmark]=/tmp/safeanalyze-pint-benchmark
    [alexh-prompt-injection-scanner]=/tmp/safeanalyze-prompt-injection-scanner
    [duriantaco-skylos]=/tmp/safeanalyze-skylos
    [promptfoo-scenarios]=/tmp/safeanalyze-doc-scenarios.html
    [promptfoo-webagents]=/tmp/safeanalyze-doc-webagents.html
)
for name in "${!TARGETS[@]}"; do
    dir="$OUT/corpus/$name"
    mkdir -p "$dir"
    ./safeanalyze scan "${TARGETS[$name]}" --mode thorough -c "$CFG" -o "$dir" > "$OUT/corpus/$name.log" 2>&1 || true
done
