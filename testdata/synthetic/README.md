# Committed Synthetic Evaluation Suite

This directory holds **self-contained, hand-authored synthetic benchmarks** committed directly to the repository so CI and offline environments can verify `safeanalyze eval` without fetching external datasets.

- External fetched datasets (`testdata/eval/*.jsonl`) are external dependencies and are **not** modified here.
- `prompt_injection_suite.jsonl` contains 40 labeled payloads (20 adversarial injections + 20 hard-benign engineering, documentation, and tool-response texts) spanning:
  - Direct & multilingual overrides (`override`, `multilingual`)
  - Chat-template / delimiter breakouts (`chat_template`)
  - Invisible Unicode / bidi / tag-character smuggling (`unicode_smuggling`)
  - Email & tool-call exfiltration (`email_exfil`, `shell_injection`)
  - Obfuscation (`leetspeak`)
  - Paraphrase & non-Latin ceiling probes (`paraphrase_probe` — intentionally included so reports remain honest about deterministic regex limits vs. local classifier / LLM-judge sidecars)
  - Hard-benign engineering prose, incident docs, JSON responses, and emoji ZWJ sequences (`benign_*`)

Run offline:

```bash
go build -o ./safeanalyze .
./safeanalyze eval testdata/synthetic/prompt_injection_suite.jsonl
```
