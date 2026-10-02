#!/usr/bin/env bash
# Reproduce the v0.3.12 (iteration 2) measurements in this report.
# Run from the repository root of a v0.3.12 checkout with eval data fetched
# (./scripts/fetch_eval.sh) and corpus targets present under /tmp/.
set -euo pipefail
export PATH=/usr/local/go/bin:$PATH
OUT="${OUT:-/tmp/safeanalyze-iter/iter2}"
CFG="${CFG:-/tmp/safeanalyze-test-config-eval.yaml}"
mkdir -p "$OUT/corpus"

go build -o safeanalyze .
go test ./...
./scripts/redteam.sh

# Labeled fast-mode eval (dev + holdout).
for ds in deepset llmail browsesafe deepset-holdout llmail-holdout browsesafe-holdout; do
    ./safeanalyze eval "testdata/eval/$ds.jsonl" \
        --json "$OUT/eval-$ds.json" \
        --fp-out "$OUT/fp-$ds.jsonl" --fn-out "$OUT/fn-$ds.jsonl"
done

# Thorough corpus scan.
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
    start=$(date +%s%N)
    ./safeanalyze scan "${TARGETS[$name]}" --mode thorough -c "$CFG" -o "$dir" > "$dir/scan.log" 2>&1 || true
    end=$(date +%s%N)
    echo $(( (end - start) / 1000000 )) > "$dir/duration.txt"
done
