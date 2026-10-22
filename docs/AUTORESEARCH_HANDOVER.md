# Handover prompt: autoresearch in safeanalyze

Paste everything below the line into a new Claude Code session started in
`/root/sideProjects/safeanalyse`.

---

You are continuing the autoresearch improvement loop of **safeanalyze** (Go CLI
prompt-injection scanner, repo `/root/sideProjects/safeanalyse`, GitHub
`andreas2301/safeanalyse`, branch `master`). Read `CLAUDE.md` first — it is the
source of truth for the loop, the set taxonomy and the decision rules. Then read
`docs/ASSESSMENT.md` (honest capability assessment) and `testdata/eval/SOURCES.md`.

## Status at handover (2026-10-22)

- Version v0.3.49 (fast mode: YARA-style regex + hidden-char checks; `icap` for
  Squid body scanning; `gate` for keep/revert; sealed + hard-benign sets).
- **The rule-adding loop is PAUSED** pending a strategy decision by the owner.
  Evidence: regex has a hard ceiling (human-written evasion set recall 0.17,
  multilingual 0.02, hard-benign FP 14 %), and round-4 gains partly overfit.
  Proposed strategy (not yet confirmed): safeanalyze becomes orchestrator +
  evaluation harness; a small classifier (Llama Prompt Guard 2 22M, ONNX, thorough
  mode) as main detector, LLM judge for uncertain items (see
  `docs/LLM_SCANNER_SPEC.md`), regex only as cheap pre-check. **Ask the owner
  before starting any new round.**

## How autoresearch works here (karpathy/autoresearch-style)

One iteration = one hypothesis, measured, then kept or reverted by rules in code.

1. **Prepare data** (once per machine; data is gitignored, pinned and sha256-checked):
   `./scripts/fetch_eval.sh` → `testdata/eval/*.jsonl` (dev, `-holdout`, `-sealed`).
   Committed sets: `hardbenign.jsonl`, `policy.yaml`. Corpus repos are pinned in
   SOURCES.md ("Thorough-mode corpus manifest"); scan config
   `/tmp/safeanalyze-test-config-eval.yaml` (trufflehog `no_verification: true`).
2. **Baseline**: `export PATH=/usr/local/go/bin:$PATH; go build -o ./safeanalyze .`
   then for every set n (dev, holdout, stress, hardbenign — NOT sealed):
   `./safeanalyze eval testdata/eval/<n>.jsonl --json $O/eval-<n>.json --fp-out $O/fp-<n>.jsonl --fn-out $O/fn-<n>.jsonl --predictions $O/predictions-<n>.jsonl`
   with `O=/tmp/safeanalyze-iter/<tag>`. Also `./scripts/redteam.sh` and the
   thorough corpus scan (regression/error check only).
3. **Hypothesize** from **PRIMARY dev** FN/FP files only (llmail, browsesafe,
   toolresp) plus multilingual dev. Never read holdout, sealed, GUARD-ONLY
   (deepset, nemotron, spanish) or REPORT-ONLY (promptprotection-stress,
   hardbenign) misclassifications. One small, generalizable change; no dataset
   strings. Run `./scripts/premortem.sh`.
4. **Implement + red-team** on a clean tree; unit tests; ≥10 evasion probes and
   ≥10 benign probes via `./safeanalyze inspect --body`.
5. **Measure** the candidate exactly like step 2 into a new tag dir.
6. **Decide with code, never by hand or in a prompt**:
   `./safeanalyze gate --baseline /tmp/safeanalyze-iter/<base> --candidate /tmp/safeanalyze-iter/<cand> --json /tmp/gate.json`
   Exit 0 = ACCEPT, 3 = REVERT, 1 = error. It enforces per-set non-regression,
   the single-sample tolerance, multilingual recall, holdout gain ≥ 0.5× dev gain,
   paired-bootstrap significance, and the 100 ms latency budget. Corpus counts,
   determinism and security hardening are still checked by hand (CLAUDE.md).
7. **Accept**: bump `pkg/version/version.go`, CHANGELOG, README, CLAUDE.md
   "Current autoresearch iteration"; commit with a Conventional Commit ≤72 chars
   ending in the version, e.g. `feat(yara): ... (v0.3.50)`; tag; push master+tag.
   **Revert**: restore only the iteration's files; tree must be clean.
8. **Round end** (stop after 2 consecutive no-gain iterations or the agreed cap):
   evaluate the **sealed** sets once, create ONE report branch
   (`report-<deadpool-name>-<date>`, worktree from master, COMPARISON.md vs round
   start incl. sealed + hardbenign + stress), push it, then do the process review
   and record it in CLAUDE.md.

## Running it with agents

Use one subagent per step (hypothesis → implement → measure → gate → release or
revert), sequentially; only the gate's exit code decides. Before each implement
step assert `git rev-parse HEAD` is the expected sha and `git status --porcelain`
is empty; stop the loop (do not count "no gain") if not. Measure before bumping the
version string, so label results by detection logic, not by `--version`.

## Pitfalls already hit (do not repeat)

- Summed-holdout checks and undocumented thresholds in loop scripts → use `gate`.
- Reusing the holdout every iteration overfits it → sealed sets once per round.
- Small sets overfit (spanish 52 positives: dev +0.43 vs holdout +0.15 F1).
- Dataset artifacts (JSON `\u` escapes trip `obfuscated_javascript`).
- Every regex adds ~2.5 ms p95; keep prefilter literals.
- Cheap-model research had wrong authors and invented numbers → verify citations.
- Do not author new evasive attack rewordings (declined); use published sets.
- Dirty trees break automation: never leave untracked files in the main checkout
  while a loop runs; agent worktrees go under `.claude/worktrees/` (gitignored
  locally). osEngineer runtime files are gitignored since v0.5.0 fixes.
- osEngineer guard: force-push and other denylisted commands need a human-authored
  `.osengineer/current-plan.md` or the owner running the command (`! ...`).
- Some sandboxes refuse commands containing the word "eval"; copy files to /tmp.

## Tooling

- osEngineer is installed (`.claude/agents`, `.claude/commands/osEngineer/*`).
  `/osEngineer:recipe-eval-harness` and the `eval-harness-engineer` agent encode
  this process generically; `bin/ose-eval-gate` (in the osEngineer repo) is the
  language-neutral gate, `safeanalyze gate` the Go one used here.
- Next candidate once the strategy is confirmed: Prompt Guard 2 22M spike in
  thorough mode, measured on all sets with `gate` (needs the owner's Hugging Face
  token and Llama 4 Community License acceptance).
