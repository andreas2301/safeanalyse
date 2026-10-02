# Eval dataset sources

Labeled data for `safeanalyze eval` is **not committed**. Regenerate it with:

```bash
./scripts/fetch_eval.sh
```

The script downloads pinned revisions over plain HTTPS (no Hugging Face token),
verifies each file's sha256, and writes the files below. Label `1` = injection
and `0` = benign. Each JSONL line looks like `{"text": ..., "label": 0|1, "source": ...}`.
Sampling uses Python's `random.Random(1337)`, so output is byte-identical between runs.
Downloads are cached in `testdata/eval/.cache/` (override with `SAFEANALYZE_EVAL_CACHE`).

| Output | Samples | Injection | Benign | Output sha256 |
|---|---|---|---|---|
| `deepset.jsonl` | 116 | 60 | 56 | `a326affea6c7b0f93b7b2b9707d16e5dbb58c68d8f5672616242a31339b062a3` |
| `llmail.jsonl` | 460 | 300 | 160 | `b82bae7a90f83bf17bd47d9ec9e9745ede7a86bbcf849bbb3c4536683c1fab8c` |
| `browsesafe.jsonl` | 600 | 300 | 300 | `f7547ff2b2f59aa4724bd150388598282fb0e097d496a1a40be0d60699e8df86` |
| `deepset-holdout.jsonl` | 546 | 203 | 343 | `563829453182bdd4c9e8b5de39d3c328c8e879ae932398b413d017b7bf4049bd` |
| `llmail-holdout.jsonl` | 460 | 300 | 160 | `52cb2f14f1edecafc8c6570cf96555f64f1d103add3c2d3fe8b8647e632c5048` |
| `browsesafe-holdout.jsonl` | 600 | 300 | 300 | `2b12814cedd07d57113a8a6ebbc7b9e66cd1f75d7321d3ec43563196ec569368` |

The first three files are the **dev** sets used while tuning rules. The `*-holdout.jsonl`
files are **holdout** sets: use them only to confirm that a change generalizes, never to
pick rules or thresholds. The script checks that no holdout text appears in any dev set and
exits non-zero if one does. The only exception is the LLMail benign emails (see below).

## deepset/prompt-injections

- URL: https://huggingface.co/datasets/deepset/prompt-injections
- Revision: `4f61ecb038e9c3fb77e21034b22511b523772cdd`
- License: Apache-2.0
- File: `data/test-00000-of-00001-701d16158af87368.parquet`
  - sha256 `39ac797cabc157eeed58435a08593b2952bb6cb16fc394a2d383f447cc7b246e` (10,892 bytes)
- Sampling: the full test split (116 rows). The dataset's own `label` (1 = injection) is kept. `source` = `deepset`.
- Holdout file: `data/train-00000-of-00001-9564e8b05b4757ab.parquet`
  - sha256 `2e10bc7ab30f542c97e4e83e2a5683000b5057d25ec10908784c631d44124c04` (40,323 bytes)
  - Holdout sampling: the full train split (546 rows: 203 injection, 343 benign). Rows whose text also appears in a dev set
    would be dropped; at this revision none do.

## microsoft/llmail-inject-challenge (LLMail-Inject)

- URL: https://huggingface.co/datasets/microsoft/llmail-inject-challenge
- Revision: `1063bdf01ec8762b812d5e06ee768a06faa5a6f7`
- License: MIT
- Files:
  - `data/labelled_unique_submissions_phase2.json`
    - sha256 `f89af984e345430c3b357903890e30867bf4676f4ef10c138cc7bad218e890b8` (68,586,935 bytes)
  - `data/emails_for_fp_tests.json`
    - sha256 `4ddd950b5dbaa8548f5597c886d8e09a051ba07f80a9291fdcca9c2397d22abe` (55,486 bytes)
- Phase 2 is used instead of phase 1 (448 MB) or the raw submissions (1.6 GB and 263 MB) so the download stays small.
- Positives: submissions with `attack_attempt` == True (21,007 of 37,303). `Unclear` and `False` are left out.
  The script takes 300 of them, stratified by `judge_category`; API-triggered submissions have no category and use `reason` (`api_triggered`).
  Each stratum gets a share in proportion to its size, rounded with the largest-remainder method.
  Strata are visited in sorted order, and texts are sorted before `random.sample`.
  `source` = `llmail/<stratum>`.
- Negatives: the 160 distinct benign emails from `emails_for_fp_tests.json` (203 entries; exact duplicate texts are dropped, keeping first-occurrence order). `source` = `llmail/benign`.
- Holdout: 300 *different* attack submissions with the same per-stratum allocation as dev
  (188 direct instructions, 62 obfuscation, 45 api_triggered, 4 social engineering, 1 direct instructions and obfuscation).
  They are drawn with a fresh `random.Random(1337)` from each stratum's sorted texts after removing every dev text.
- **Holdout negatives reuse the dev benign emails.** The dataset has no other benign emails
  (`emails_for_fp_tests.json` is its only benign file, and all 160 distinct emails are already in dev).
  So `llmail-holdout.jsonl` repeats the same 160 benign rows. Its precision and FP counts are **not**
  independent of dev; only its recall is a true holdout measurement.

## perplexity-ai/browsesafe-bench (BrowseSafe-Bench)

- URL: https://huggingface.co/datasets/perplexity-ai/browsesafe-bench
- Revision: `b506fb5bc7fd4472c8738055a67a0ef6406afdc9`
- License: MIT
- File: `test.parquet`
  - sha256 `00cbad96b60fee46e016d79af6981fb221384c61f12cf28b4f04b5a6420573d0` (46,938,712 bytes; 3,680 rows: 1,824 yes / 1,856 no)
- Sampling: a seeded sample of 300 rows with `label=yes` (→ 1), then 300 rows with `label=no` (→ 0).
  Both draws use the same RNG, `yes` first, and rows are kept in file order within each label. `source` = `browsesafe`.
  Samples are full HTML pages (about 55 KB on average).
- Holdout: a fresh `random.Random(1337)` draws 300 `yes` and 300 `no` rows from the test rows that dev did not pick
  (and whose content does not match any dev text): 1,524 `yes` and 1,556 `no` candidates. The ordering rules are the same as dev.

Parquet files are decoded by a small stdlib-only reader inside `fetch_eval.sh`, so no pyarrow is needed.
It was cross-checked against the Hugging Face datasets-server `/rows` API.

## Thorough-mode corpus manifest

The corpus used for `safeanalyze scan --mode thorough` (see CLAUDE.md, autoresearch loop) was
**re-baselined on 2026-10-02** at the upstream HEAD commits below. To reproduce it, check out each repo at its SHA,
fetch each document and confirm its sha256. Corpus baseline results come from the v0.3.9 detection logic with
trufflehog `no_verification: true`, ML off, and yara, hidden_chars, entropy, prompt-injection-scanner and semgrep on.

```bash
git clone https://github.com/microsoft/BIPIA /tmp/safeanalyze-BIPIA && git -C /tmp/safeanalyze-BIPIA checkout <sha>
# ...same for each repo below...
curl -sSfL <url> -o /tmp/safeanalyze-doc-scenarios.html && sha256sum /tmp/safeanalyze-doc-scenarios.html
```

| Target | Local path | Upstream | Pin | Files |
|---|---|---|---|---|
| `repos/microsoft-bipia` | `/tmp/safeanalyze-BIPIA` | https://github.com/microsoft/BIPIA | `a004b69ec0dd446e0afd461d98cb5e96e120a5d0` (2024-04-15T10:08:17+08:00) | 99 |
| `repos/uiuc-injecagent` | `/tmp/safeanalyze-InjecAgent` | https://github.com/uiuc-kang-lab/InjecAgent | `f19c9f2c79a41046eb13c03c51a24c567a8ffa07` (2024-07-02T00:38:20-05:00) | 24 |
| `repos/lakera-pint-benchmark` | `/tmp/safeanalyze-pint-benchmark` | https://github.com/lakeraai/pint-benchmark | `0efab3f463eae9c823130d8faffb71b2e7c06e63` (2026-04-02T11:08:48+02:00) | 33 |
| `repos/alexh-prompt-injection-scanner` | `/tmp/safeanalyze-prompt-injection-scanner` | https://github.com/alexh-scrt/prompt-injection-scanner | `33dd171bf0096e9782dfe971ea21c4795c8eb9a6` (2026-02-26T09:49:17-05:00) | 18 |
| `repos/duriantaco-skylos` | `/tmp/safeanalyze-skylos` | https://github.com/duriantaco/skylos | `fc5fffd2bad9ebeb2c215b140868c9b84a4751be` (2026-10-02T14:16:56+08:00) | 1966 |
| `docs/promptfoo-scenarios` | `/tmp/safeanalyze-doc-scenarios.html` | https://www.promptfoo.dev/docs/configuration/scenarios/ | sha256 `a2d97060d5ec5cbc561b1d6090f43f13cfe7bd8fad13de25358db3cebc31ef48` (58870 bytes) | 1 |
| `docs/promptfoo-webagents` | `/tmp/safeanalyze-doc-webagents.html` | https://www.promptfoo.dev/blog/indirect-prompt-injection-web-agents/ | sha256 `60b8c971411752857b647c76404c67a30dc7ce587c138cc707e68483f2039ce2` (56541 bytes) | 1 |

Notes from the manifest: the v0.3.7 corpus SHAs were never recorded, so drift since then is possible. skylos
is the clearest case: it now has 1,966 files, and v0.3.7 scanned 1,201. Compare corpus results only against runs on these pins.

Full manifest (`/tmp/safeanalyze-corpus-manifest.json`, copied verbatim):

```json
{
  "date": "2026-10-02",
  "source_evidence": "report README tables (report-*-2026-07-15) + original kimi-code session clone/curl commands",
  "config": {
    "path": "/tmp/safeanalyze-test-config-v037.yaml",
    "sha256": "cfc9294c94af454ed308b16275d2f81a8854ba1d3802f7079ca902297821d1aa",
    "bytes": 1154,
    "provenance": "exact: kimi session Write of v031 (1154 B), cp chain v031->v032->v033->v034->v035->v037; size matches original ls listing"
  },
  "note": "v0.3.7 baseline corpus SHAs were never recorded; upstream drift possible (skylos changed: 1966 files now vs 1201 scanned, skylos/llm/verify_llm.py removed).",
  "targets": [
    {
      "label": "repos/microsoft-bipia",
      "local_path": "/tmp/safeanalyze-BIPIA",
      "upstream_url": "https://github.com/microsoft/BIPIA",
      "clone": "git clone --depth 1",
      "commit_sha": "a004b69ec0dd446e0afd461d98cb5e96e120a5d0",
      "commit_date": "2024-04-15T10:08:17+08:00",
      "file_count": 99,
      "v037_files_scanned": 99
    },
    {
      "label": "repos/uiuc-injecagent",
      "local_path": "/tmp/safeanalyze-InjecAgent",
      "upstream_url": "https://github.com/uiuc-kang-lab/InjecAgent",
      "clone": "git clone --depth 1",
      "commit_sha": "f19c9f2c79a41046eb13c03c51a24c567a8ffa07",
      "commit_date": "2024-07-02T00:38:20-05:00",
      "file_count": 24,
      "v037_files_scanned": 22
    },
    {
      "label": "repos/lakera-pint-benchmark",
      "local_path": "/tmp/safeanalyze-pint-benchmark",
      "upstream_url": "https://github.com/lakeraai/pint-benchmark",
      "clone": "git clone --depth 1",
      "commit_sha": "0efab3f463eae9c823130d8faffb71b2e7c06e63",
      "commit_date": "2026-04-02T11:08:48+02:00",
      "file_count": 33,
      "v037_files_scanned": 32
    },
    {
      "label": "repos/alexh-prompt-injection-scanner",
      "local_path": "/tmp/safeanalyze-prompt-injection-scanner",
      "upstream_url": "https://github.com/alexh-scrt/prompt-injection-scanner",
      "clone": "git clone --depth 1",
      "commit_sha": "33dd171bf0096e9782dfe971ea21c4795c8eb9a6",
      "commit_date": "2026-02-26T09:49:17-05:00",
      "file_count": 18,
      "v037_files_scanned": 18
    },
    {
      "label": "repos/duriantaco-skylos",
      "local_path": "/tmp/safeanalyze-skylos",
      "upstream_url": "https://github.com/duriantaco/skylos",
      "clone": "git clone --depth 1",
      "commit_sha": "fc5fffd2bad9ebeb2c215b140868c9b84a4751be",
      "commit_date": "2026-10-02T14:16:56+08:00",
      "file_count": 1966,
      "v037_files_scanned": 1201
    },
    {
      "label": "docs/promptfoo-scenarios",
      "local_path": "/tmp/safeanalyze-doc-scenarios.html",
      "upstream_url": "https://www.promptfoo.dev/docs/configuration/scenarios/",
      "fetch": "curl -sSfL",
      "sha256": "a2d97060d5ec5cbc561b1d6090f43f13cfe7bd8fad13de25358db3cebc31ef48",
      "bytes": 58870,
      "file_count": 1,
      "v037_files_scanned": 1
    },
    {
      "label": "docs/promptfoo-webagents",
      "local_path": "/tmp/safeanalyze-doc-webagents.html",
      "upstream_url": "https://www.promptfoo.dev/blog/indirect-prompt-injection-web-agents/",
      "fetch": "curl -sSfL",
      "sha256": "60b8c971411752857b647c76404c67a30dc7ce587c138cc707e68483f2039ce2",
      "bytes": 56541,
      "file_count": 1,
      "v037_files_scanned": 1
    }
  ]
}
```
