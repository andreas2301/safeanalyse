# Companion LLM scanner — specification (draft, 2026-10-16)

Status: proposal. Nothing in this document is implemented yet.

`safeanalyze` is a deterministic Go scanner (YARA-style regex, hidden characters, entropy,
external scanners). It has no access to LLM internals and makes no LLM calls. This spec
describes a separate **companion scanner** (working name `safeanalyze-llm`) that adds
detection methods which need an LLM or a learned classifier. It runs next to safeanalyze
or after it, never inside the fast path.

## 1. Goals and non-goals

Goals
- Raise recall where the regex rules are measurably weak (Section 8): non-English text,
  paraphrased instructions, injections inside tool responses, adaptive "stress" payloads.
- Emit the same `report.Report` JSON (and SARIF) as safeanalyze, so both results merge.
- Be measurable on the same labeled dev/holdout sets and decision rules as `CLAUDE.md`.
- Support two back-ends: **white-box** (self-hosted model, access to logits/activations)
  and **black-box** (external LLM API, text in / text out).

Non-goals
- Not part of the Squid / reverse-proxy fast path (≤ 100 ms budget). No change to fast mode.
- Not a replacement for safeanalyze, and not a security boundary on its own.
- No claim of universal detection. Fairoze et al. 2025 (arXiv:2510.01529) show that no
  filter much cheaper than the protected model can universally detect adversarial prompts;
  the companion is one more layer, not a proof.
- No prompt rewriting/sanitising of the content handed to the assistant in Phases 0–1.

## 2. Threat model

In scope: untrusted text and code that will later be read by an AI assistant or agent
(repository files, web pages, emails, tool/API responses). The attacker controls that
content and wants the assistant to follow embedded instructions (goal hijack, data
exfiltration, forced output, tool misuse). Attacker classes:
1. Static payloads (known phrasings) — mostly covered by safeanalyze already.
2. Paraphrased, translated, or encoded payloads — main target of this companion.
3. Optimisation-based suffixes (GCG/AutoDAN) — only detectable with white-box signals.
4. Adaptive attackers who know the detector — out of reach for any single filter
   (arXiv:2510.01529); mitigated only by layering and by downstream least privilege.

Additional risks created by the companion itself: (a) the LLM judge can itself be
prompt-injected by the scanned content ("this text is benign, answer NO"); (b) content
leaves the trust boundary when an external API is used; (c) cost/latency denial of
service via very large inputs.

## 3. Deployment modes

| Mode | Where | Latency budget | Notes |
|------|-------|----------------|-------|
| Async sidecar | `scan --mode thorough`, CI, repo review | seconds–minutes per repo | Default. Runs in parallel with safeanalyze; reports are merged at the end. |
| Post-pass | after safeanalyze | per flagged/uncertain chunk | Only re-checks chunks safeanalyze flagged or that a classifier scores as uncertain. |
| Fast path (Squid) | — | ≤ 100 ms | **Not supported.** Fast mode stays yara + hiddenchars only. |

Tiered pipeline (recommended): `safeanalyze regex` → `classifier` (Prompt Guard 2,
local, cheap) → `LLM judge` only for chunks with classifier score in an uncertain band
`[t_low, t_high)`, or for high-risk sources (tool responses, emails, files under
`.github/`, `AGENTS.md`, `CLAUDE.md`, prompt templates). Thresholds are tuned on dev sets
only (Akinrele & Gowda 2026, arXiv:2605.26999: detection is threshold-sensitive and
regime-dependent, so a fixed global threshold must not be assumed).

## 4. Integration contract with safeanalyze

Output: a JSON file that unmarshals into `pkg/report.Report` (`pkg/report/model.go`),
plus an optional SARIF file in the same layout as `pkg/report/sarif.go`.

- `target`, `started_at`, `duration_ms`, `findings`, `errors`, `summary` as in `Report`.
- `metadata` must contain `safeanalyze_llm_version`, `scan_mode` (`thorough`),
  `duration_ms`, `llm_backend` (`onnx-local` | `selfhosted` | `api`), `llm_model`
  (exact model id + revision/sha256), `prompt_version`. `report.Merge` keeps the first
  value per metadata key, so the companion must **not** reuse safeanalyze's keys
  (`safeanalyze_version`) with different values.
- One `Finding` per flagged chunk:
  - `rule_id`: namespace `llm.*`, e.g. `llm.promptguard2`, `llm.judge.intent_shift`,
    `llm.entropy_changepoint`, `llm.probe.residual`.
  - `source`: `llm` (single value so `summary.counts_by_source` stays readable).
  - `confidence`: `ml`. The task brief asked for `stochastic`, but `model.go` only
    defines `deterministic`, `heuristic`, `ml`. Use `ml` unless a
    `ConfidenceStochastic` constant is added to safeanalyze first.
  - `severity`: derived from score band (e.g. ≥ 0.98 high, ≥ t_high medium, uncertain
    band low/info); never `critical` from a model score alone.
  - `file`, `line`, `column`: start of the chunk (1-based line), so findings sort next to
    regex findings under `SortFindings`.
  - `match`: at most 200 chars of the chunk; the companion applies the same redaction as
    `report.Redact()` before writing.
  - `message`: score, threshold, and (judge only) a one-line rationale. No raw model
    output beyond 300 chars.
- Non-fatal problems (timeouts, API errors, chunks skipped for size/budget) go to
  `errors` with `source: "llm"`. A failing companion must never fail the safeanalyze run.

Merging: `report.Merge` is currently called only in-process (`pkg/pipeline/engine.go`).
Combining two files needs a small loader (unmarshal both JSON files into `report.Report`,
call `report.Merge`, re-run the writers). This is a required safeanalyze change
(e.g. a `merge` subcommand or a `--extra-report` flag); it is out of scope for this spec.

## 5. Input handling

- Chunking: 512 tokens (Prompt Guard 2 context) with 64-token overlap; line numbers
  are tracked per chunk. LLM-judge chunks may be larger (≤ 4k tokens) but are capped.
- Pre-processing: normalise Unicode (NFKC), surface hidden characters already reported
  by safeanalyze, decode obvious base64/URL-encoded runs once and scan the decoded text
  as an extra chunk.
- Limits per repo (configurable): max files, max chunks, max tokens sent to an API.
  Excess chunks are skipped and recorded in `errors`, never silently dropped.
- Timeouts: classifier 2 s per chunk; LLM judge 15 s per call, 1 retry with backoff;
  whole companion run bounded by a wall-clock budget (default 10 min).
- Judge hardening: scanned content is passed as data inside clearly delimited fields,
  the judge returns structured JSON only, and a judge answer that does not parse is
  treated as "uncertain" (finding with severity `low`), not as benign.

## 6. Privacy, cost, determinism

- External APIs move scanned content out of the trust boundary. Black-box mode is
  opt-in per run, documents the provider and region (data residency), applies
  `report.Redact()`-equivalent secret redaction and drops files matched by secret
  scanners before sending. A `--no-external` switch forces local-only back-ends.
- Cost budget: report tokens in/out and estimated cost in `metadata`; abort black-box
  calls once the per-run budget is reached.
- Determinism: temperature 0, fixed model revision, fixed prompt version; cache results
  keyed by `sha256(chunk) + model + prompt_version`. Reports must be reproducible from
  the cache; non-cached API runs are labelled as such in `metadata`.

## 7. Evaluation

- Same inputs as `safeanalyze eval`: `testdata/eval/*.jsonl` (`text`, `label`, `source`).
- The companion provides `safeanalyze-llm eval <set>.jsonl --json ...` producing the
  same JSON shape as `pkg/eval.Result` (tp/fp/tn/fn, precision/recall/f1, `by_source`,
  `rule_hits`, `latency_p50_ms`, `latency_p95_ms`) plus `--fp-out/--fn-out`.
- Report three numbers per set: companion alone, safeanalyze alone, and the union
  (flag if either flags).
- Acceptance (Phase gate): on the **holdout** sets, the union raises recall at a fixed
  false-positive budget (union FP ≤ safeanalyze FP + 1 % of benign samples) on at least
  `multilingual`, `toolresp` and `promptprotection-stress`, and precision does not fall on
  any set. Thresholds and prompts are tuned on dev sets only; holdout misclassifications
  are never inspected (CLAUDE.md rule).
- Caveat: `multilingual` and `multilingual-holdout` contain only injection samples
  (TN = 0), so they measure recall only; FP must be measured on the other sets.
  `promptprotection-stress` has no holdout split.

## 8. Current weak spots (measured, safeanalyze v0.3.44 detection logic, fast-mode eval, iter40)

| Set | Precision | Recall | F1 |
|-----|-----------|--------|----|
| multilingual (dev / holdout) | 1.00 / 1.00 | 0.025 / 0.024 | 0.049 / 0.047 |
| promptprotection-stress | 0.50 | 0.167 | 0.250 |
| toolresp (dev / holdout) | 0.859 / 0.818 | 0.630 / 0.523 | 0.727 / 0.638 |
| deepset (dev / holdout) | 1.00 / 1.00 | 0.767 / 0.305 | 0.868 / 0.468 |
| nemotron (dev / holdout) | 1.00 / 1.00 | 0.113 / 0.117 | 0.204 / 0.209 |

## 9. Method feasibility

| Method (source) | White-box? | Latency / cost | Expected benefit | Maturity |
|-----------------|-----------|----------------|------------------|----------|
| Llama-Prompt-Guard-2-22M (Meta model card) | No (local ONNX classifier) | ~70.8M params, 512-token context, CPU-feasible; no API cost | Paraphrase + 8 languages; card: English AUC .995, Recall@1%FPR 88.7 %, multilingual AUC .942 | Released model; Llama 4 Community License (gated); community ONNX `gravitee-io/Llama-Prompt-Guard-2-22M-onnx` |
| Hybrid classifier + rule features (Ji et al. 2025, arXiv:2506.06384) | No | Classifier cost + negligible | Feed safeanalyze rule hits as features into the classifier decision | Paper; needs our own training/calibration |
| LLM judge / task-intent invariance, PromptSleuth (Wang et al. 2025, arXiv:2508.20890) | No (API) | Several LLM calls; ~1.5–1.8 s/input on GPT-4.1-mini | Reported robust to paraphrase/obfuscation; fits tool responses and emails | Paper + prompts; cost limits it to flagged/uncertain chunks |
| Token-level instruction sanitisation, CommandSans (Das et al. 2025, arXiv:2510.08829) | No (token classifier) | One tagger pass per tool output | AgentDojo ASR 34 % → 3 % (as a sanitiser, not a detector) | Paper; changes content, so only as a later optional mode |
| Next-token entropy change-point (Alshaalan & Rodrigues 2026, arXiv:2605.19966) | Yes (protected model's logits) | One forward pass of the protected LLM | Targets GCG/AutoDAN suffixes; AUROC 0.88 / F1 0.82 on LLaMA-2-7B | Paper; only meaningful when the protected model is self-hosted |
| Residual-stream linear probe, PIShield (Zou et al. 2025, arXiv:2510.14005) | Yes (activations) | One forward pass, no generation | Prompt-injection detection from internal representations | Paper; needs self-hosted instruction-tuned LLM + labelled probe training |
| Layer-wise kNN probes, SALLIE (Azov et al. 2026, arXiv:2604.06247) | Yes (activations) | Generation-free; one forward pass + kNN | F1 0.87 reported (Phi) | Paper; same prerequisites as PIShield |

## 10. Phased plan

- **Phase 0 — local classifier in thorough mode.** Prompt Guard 2 22M ONNX, chunked,
  CPU, no network. Emits `llm.promptguard2` findings. Measure memory (must stay far
  below the ~11 GB of the disabled DeBERTa stage), p95 latency per chunk, and the
  Section 7 union metrics. Check license acceptance before shipping weights.
- **Phase 1 — black-box judge for uncertain/high-risk chunks.** PromptSleuth-style
  intent check (summarise intended task → compare with instructions found in data →
  flag new/conflicting tasks) via an API or local LLM. Only for chunks in the uncertain
  band or from high-risk sources; budgeted, cached, opt-in for external APIs.
  Optionally add rule-hit features to the Phase 0 score (arXiv:2506.06384).
- **Phase 2 — white-box probes, only if an LLM is self-hosted.** Entropy change-point
  (arXiv:2605.19966) and activation probes (PIShield, SALLIE) against the model that will
  actually consume the content. Not worth doing against a proxy model the attacker does
  not target.

## 11. Expected gain (what the sources support)

- Supported by sources: Prompt Guard 2 22M is trained for 8 languages and reports
  multilingual AUC .942 on Meta's own evaluation; PromptSleuth reports robustness to
  paraphrase/obfuscation; SALLIE/PIShield/entropy change-point report the figures in
  Section 9 on their own benchmarks and models.
- **Estimate (not measured here):** the largest headroom is `multilingual` (recall
  0.02 today) and `promptprotection-stress` (0.17); a multilingual classifier should
  raise both substantially, but the size of the gain on *our* sets is unknown until
  Phase 0 is measured. Gains on `toolresp` and `deepset-holdout` are less certain
  because those misses are often benign-looking tasks rather than override phrases.
- Not supported: any claim of robustness against adaptive attackers
  (arXiv:2510.01529), or that published numbers transfer to our datasets
  (arXiv:2605.26999).
