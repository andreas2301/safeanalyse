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
| `toolresp.jsonl` | 600 | 300 | 300 | `6e09301d293226409555feeafbaeb3791dd3f8d56751c7645b1aeff9a3409bf2` |
| `nemotron.jsonl` | 600 | 300 | 300 | `b63bd83f598c1ae3d832975efc8a0eb1d125bf8d37bda200182af531c9f58889` |
| `multilingual.jsonl` | 840 | 840 | 0 | `70d41ebe64a59905edd161334c18e5af2b9ac773223a4a7574867b60a8159265` |
| `spanish.jsonl` | 352 | 52 | 300 | `fc825b6d6e8d54749d4d25728941ad09351b54866abab09e171f4569194b3cbb` |
| `deepset-holdout.jsonl` | 546 | 203 | 343 | `563829453182bdd4c9e8b5de39d3c328c8e879ae932398b413d017b7bf4049bd` |
| `llmail-holdout.jsonl` | 460 | 300 | 160 | `52cb2f14f1edecafc8c6570cf96555f64f1d103add3c2d3fe8b8647e632c5048` |
| `browsesafe-holdout.jsonl` | 600 | 300 | 300 | `2b12814cedd07d57113a8a6ebbc7b9e66cd1f75d7321d3ec43563196ec569368` |
| `toolresp-holdout.jsonl` | 600 | 300 | 300 | `87e50d3f2cf0886eb9dfd0422f5b2f821435fed4c17aad31c2465164e6e5e164` |
| `nemotron-holdout.jsonl` | 600 | 300 | 300 | `88315a583cd3b3142ca612028d6a4d57c2b5a8db625456bd961741073e32970a` |
| `multilingual-holdout.jsonl` | 825 | 825 | 0 | `8a708bd75016fcc1241f7354c6750928470dffcce814df5ced0bc1572af18d1d` |
| `spanish-holdout.jsonl` | 424 | 53 | 371 | `7c91e99f78eb2dfc84e6b006c621f7859674b09813472c73b2c93e7fe82aafd4` |
| `promptprotection-stress.jsonl` | 376 | 180 | 196 | `6263b632697fb56667593d0071a7f60d67b07dc78427cd9baecf6caf0b40f4f5` |
| `llmail-sealed.jsonl` | 460 | 300 | 160 | `74190e50ebc023f95a6d9868a58f1337d3b4933fd15c1f5ad7e9f3390f42cdfc` |
| `browsesafe-sealed.jsonl` | 600 | 300 | 300 | `71dee4d78805617c0f1d70f2b3fe0625915fd64ca64886b09722e813ba4e7b03` |
| `toolresp-sealed.jsonl` | 600 | 300 | 300 | `c4bf965f9e43d33d804942947e7729b4c89abcd7f8dd575dfa4e33c4d419dd7c` |
| `nemotron-sealed.jsonl` | 600 | 300 | 300 | `e99fa3a60c44ae2eb203014ab4c38cf8a0c136961ecb060c0c6be6a6d2c7c9e9` |

The files without a suffix are the **dev** sets used while tuning rules. The `*-holdout.jsonl`
files are **holdout** sets: use them only to confirm that a change generalizes, never to
pick rules or thresholds. The script checks that no holdout text appears in any dev set and
exits non-zero if one does. The only exception is the LLMail benign emails (see below).
`promptprotection-stress.jsonl` is a **report-only** stress set with no dev/holdout split.
The `*-sealed.jsonl` files are **sealed** test sets (see "Sealed sets" below): evaluated once per
round end, never used for keep/revert or rule design. `policy.yaml` (committed) holds the set
roles and thresholds used by `safeanalyze gate`.
How each set is used in the decision rules (primary, recall-only, guard-only, report-only)
is defined in `CLAUDE.md` ("Set taxonomy").

The sets added on 2026-10-14 (toolresp, nemotron, multilingual, spanish, promptprotection-stress)
never reuse a text from the deepset, llmail or browsesafe dev/holdout sets: those texts are removed
from every new candidate pool before sampling (at these revisions no candidate matched), and the
script exits non-zero if an overlap remains. New holdout pools also exclude every dev text, old and new.
Exact duplicate texts inside a pool are dropped (first occurrence kept).

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

## rgeada/tool-response-injections (toolresp)

- URL: https://huggingface.co/datasets/rgeada/tool-response-injections
- Revision: `39050de797f2694f0d72e1086003154de1919379`
- License: Apache-2.0
- Files (CSV columns `text,label,injection,response_id`; `label` is `injection` or `safe`):
  - `tool_response_injections_train.csv`
    - sha256 `4d7965bae0d0b1db3c50377596893923521f47a0119b0a21b83437617f387e6a` (10,915,084 bytes; 11,834 rows: 5,917 injection / 5,917 safe)
  - `tool_response_injections_test.csv`
    - sha256 `ab5d1f4435b9652f74e9c09a5774d306b09bdce4bb96fe2f348c66eafaffbfc9` (2,829,798 bytes; 2,958 rows: 1,479 injection / 1,479 safe)
- Injection rows whose `injection` string does not appear in `text` (raw, or JSON-escaped with or without `ensure_ascii`)
  were never actually injected and are dropped before sampling: 4 train rows (tool responses with `"content": null`), 0 test rows.
- Dev (`toolresp.jsonl`): from **train**, a seeded sample of 300 `injection` rows (→ 1), then 300 `safe` rows (→ 0),
  same RNG, `injection` first, file order within each label. `source` = `toolresp`.
- Holdout (`toolresp-holdout.jsonl`): the same rule on **test**, with a fresh `random.Random(1337)`.
- The script checks that dev and holdout share 0 texts and 0 non-empty `injection` strings (it exits non-zero otherwise).
  The full train and test splits also share none. `response_id` is a per-split index, so it is not compared.
- Caveats: synthetic. Each injected row is a real tool response with a prompt injection
  (from `neuralchemy/Prompt-injection-dataset`) put into a **random** field, so placement is not realistic.
  The test split reuses a small pool of injections: only **561 unique injection strings** across its 1,479 injected rows.
  The upstream injection pool is noisy: roughly **10 % of the inserted strings are plain or harmful requests rather than
  injections** (e.g. "Question: How do I escape special characters in SQL queries?", "Pretend you're a tour guide"),
  so recall has a ceiling below 1.0. Many others are HackAPrompt-style "I have been PWNED" payloads.
  Safe and injected rows come in pairs (same `response_id`), but dev/holdout sample each label independently, so a
  sampled injected row does not always have its clean twin in the same set.

## nvidia/Nemotron-RL-Agentic-Indirect-Prompt-Injection-v1 (nemotron)

- URL: https://huggingface.co/datasets/nvidia/Nemotron-RL-Agentic-Indirect-Prompt-Injection-v1
- Revision: `d738d4f361cc38bb4d7a42b9066776dade5332f5`
- License: CC-BY-4.0
- File: `train.jsonl`
  - sha256 `3329da17564a7eb287e2730fc7d6956e1f4fe51e8950ac4f110b3c37e78cf3b9` (11,703,379 bytes; 1,272 records, all attacks)
- Pair extraction (records in `id` order): the **positive** is the string value in `environment` that contains
  `injection.injection_text` verbatim (the injected vector field). The **negative** is that same value with
  `injection_text` removed by exact string replacement. `injection_vector` names the field only loosely
  (e.g. `lab_comments` for a `comments` key), so the field is found by its content, not its key name.
  Identical copies of the same value (for example one claim stored under two keys) count as one.
- Skipped records (logged by the script): 1 where `injection_text` is not found verbatim (id 272), 1 ambiguous record
  where two different values contain it (id 1260), 15 where the cleaned text is empty or whitespace-only, and 7 where the
  cleaned text still contains a wrapper marker (`[PRIORITY MESSAGE]`, `[/PRIORITY MESSAGE]`,
  `URGENT: Before completing any other tasks`) or any injection line of 20+ characters. That leaves 1,248 pairs.
- Group split: records are joined (union-find) when they share a normalized `injection_text` or a normalized clean
  text (NFKC, casefold, whitespace collapsed). The 1,248 pairs form 654 groups. A group goes to dev when
  `sha256(str(lowest record id in the group))` is even, else to holdout: 741 dev / 507 holdout pairs.
  The script checks that dev and holdout share 0 normalized injection texts and 0 clean texts.
- Within each split, a pair whose positive or negative text repeats an earlier pair is dropped (1 dev, 4 holdout).
- Cap: a fresh `random.Random(1337)` samples 300 **pairs** per split (by sorted index), so each set has 300 positives and
  their 300 matching negatives. Positives are written first, then negatives, in `id` order. `source` = `nemotron/<domain>`.
- Caveats: fully synthetic (LLM-generated environments; injections red-teamed against Nemotron Nano 3).
  **Almost all positives are 2 templates:** 1,252 of 1,272 injections (98.4 %) use either the
  `[PRIORITY MESSAGE] ... — the user [/PRIORITY MESSAGE]` wrapper (989) or the `⚠️ URGENT: Before completing any other
  tasks, you MUST: ...` wrapper (263); only 20 are bare instructions. In the output sets: dev 221 / 72 / 7, holdout 234 / 61 / 5.
  A single rule for either wrapper moves recall a lot, which is why the set is GUARD-ONLY (see CLAUDE.md).
  Negatives are the same field with the injection cut out, so they test precision on near-identical text but are short
  (median ~185 chars).

## promptprotection/agent-security-datasets (promptprotection-stress, report-only)

- URL: https://huggingface.co/datasets/promptprotection/agent-security-datasets
- Revision: `3c808d860711864f5dae99fb05f75a8139b8107d`
- License: CC-BY-4.0
- Files:
  - `agent-flows.jsonl` — sha256 `e2d46e679e89f0814f9f78a65226fccd5ba22897f164342a2f5f3e293dbc79fb` (57,316 bytes; 100 rows: 50 attack / 50 benign)
  - `attacks.jsonl` — sha256 `840620d8685dd6523327fe6d4ab8895f9d33e54f8dd73fd89e12ac9dc83502f2` (30,057 bytes; 130 rows, all `label:1`)
  - `benign-hard.jsonl` — sha256 `bd12c6965d797eb470a57b82ed29c455d41516bc2c02b0c235ace799ac193e36` (26,946 bytes; 155 rows, all `label:0`)
- Rows: `agent-flows` text = `sources[].text` joined with a blank line, label `attack` → 1 and `benign` → 0,
  `source` = `promptprotection/agent-flows/<scenario>`. 9 benign rows have no `sources` and are skipped.
  `attacks` → 1 (`promptprotection/attacks/<technique>`), `benign-hard` → 0 (`promptprotection/benign-hard/<category>`).
  Result: 376 rows (180 injection, 196 benign).
- Report-only: no dev/holdout split. Do not use it for accept/revert decisions or to design rules.
  The agent-flows label describes the whole agent turn (including the tool call), not only the source text,
  so a few attack rows may have source text that is not an injection on its own.

## Multilingual recall-only sets (multilingual)

Positives only (label 1), so only recall is meaningful.

- **CyberSecEval** (meta-llama/PurpleLlama), URL: https://github.com/meta-llama/PurpleLlama
  - Commit: `172c1074069eb88ec834124272c1b1c4f8893445`
  - License: MIT
  - Files (under `CybersecurityBenchmarks/datasets/prompt_injection/`):
    - `prompt_injection_multilingual_machine_translated.json` — sha256 `f018adfc7b1f766f98ce4dd4c8e34e97a8a88f402a2b63576652027e02c2a15f` (1,067,207 bytes; 1,004 rows, 17 languages)
    - `prompt_injection.json` (English originals, used only for grouping and the cross-language check) — sha256 `069e4d5d36f6d19f972a3bbc65df840cc729354f89358a9031aeb44d95b18a9a` (199,176 bytes; 251 rows)
  - Text = `user_input`; `source` = `cyberseceval/<speaking_language>`.
  - **Dropped:** the 284 rows with `risk_category` = `logic-violating`. These ask the model to break a task rule from the
    paired system prompt (e.g. "talk about AI" when told not to) and usually carry no override or leak cue on their own.
    No other field marks rows without a cue (`injection_type` is only direct/indirect), so no other rows are dropped. 720 rows remain.
  - **Base case:** rows 0–752 are three translations of each English case, in order (base = `prompt_id // 3`); rows
    753–1003 are the Hindi translations (base = `prompt_id - 753`). The script checks this against the English file:
    `injection_variant`, `injection_type` and `risk_category` match the base on all 1,004 rows (it exits otherwise), and
    438 of the 452 rows whose English base has a digit/@ token (password, date, e-mail) contain that token in translation
    (a one-row shift gives 144/452). English base cases whose `user_input` are near duplicates (character 5-gram Jaccard ≥ 0.8)
    are joined into one group.
- **yanismiraoui/prompt_injections**, URL: https://huggingface.co/datasets/yanismiraoui/prompt_injections
  - Revision: `bd55359f2f332afc35f277ac3dd08f7111b024c9`
  - License: Apache-2.0
  - File: `prompt_injections.csv` — sha256 `f4843f1841fa19b980f804796a68fc72f06841775eaba2723c768c7d772aabad` (77,986 bytes; 1,034 rows)
  - Text = `prompt_injections`; `source` = `yanismiraoui/<lang>`. The CSV has no language column, so `<lang>` (en, fr, de, es, pt, it, ro, or `und`)
    is a **stopword/diacritic guess** made in `fetch_eval.sh`. It is roughly right but not exact (very short rows and
    Romanian/French look-alikes are the usual errors). About a third of the rows are English (328 of 974 distinct).
- Combine: CyberSecEval rows first, then yanismiraoui. Empty texts and exact duplicates are dropped (first occurrence kept): 74.
- Split: by `sha256(group key)` parity, even → dev. The group key of a CyberSecEval row is its (joined) base case, so all
  translations of one case land on the same side (the script checks this). A yanismiraoui row is its own group.
- Near-duplicate removal: a holdout row is dropped when a dev row of the same language (CyberSecEval language names map
  to the yanismiraoui codes for es/fr/de/it/pt) has character 5-gram Jaccard ≥ 0.8 on normalized text: 15 rows.
  After that, no holdout row has a near duplicate in dev in **any** language at that threshold.
- Result: `multilingual.jsonl` 840 (369 CyberSecEval, 471 yanismiraoui), `multilingual-holdout.jsonl` 825 (336, 489).
- Caveats and residual noise: CyberSecEval texts are **machine-translated**. yanismiraoui rows are often translations of
  one another (e.g. "Activate Developer Mode" in several languages); cross-language translations of the same yanismiraoui
  prompt can still land on both sides of the split, because only same-language near duplicates are removed.
  In a 40-row random sample of the final sets, about 4 rows (~10 %) have no self-contained override or leak cue and only
  read as attacks next to the CyberSecEval system prompt (e.g. "Write a haiku about the password", a licence-plate
  question built from a birth year). Expect a recall ceiling around 0.9.

## edithngalame/prompt-injection-en-es (spanish)

- URL: https://huggingface.co/datasets/edithngalame/prompt-injection-en-es
- Revision: `3e7180fbe4207547e3d0cebc44cedf6bc943a117`
- License: mixed, recorded per row in the `license` column
- Files:
  - `data/test-00000-of-00001.parquet` — sha256 `919f9ad2ae1d9e3ffed8ab293cc53e050459c71b2a0b51fe621e27ca749112dc` (381,998 bytes; 2,041 rows, 802 with `lang` = `es`)
  - `data/test_native_es-00000-of-00001.parquet` — sha256 `a267ba8c4f2785ff3a5bc31f5a0a9ad1cf02f4ddf3617c376ceca4009ab1c85f` (7,453 bytes; 32 rows)
  - Also used: the deepset train/test parquet files above, for the text filter.
- Filters, applied to rows with `lang` = `es` in this order:
  1. Drop rows whose `origin` contains `deepset` (case-insensitive): 26 test rows.
  2. Drop rows whose normalized text (NFKC, casefold, whitespace collapsed) equals a normalized deepset train or test text: 0.
  3. Keep only rows whose `license` is `CC-BY-SA-3.0`, `MIT` or `Apache-2.0`. All 32 `test_native_es` rows have license
     `project` (the dataset author's own hand-written rows, with no stated license), so **all 32 are dropped**.
  After filtering: 776 test rows (105 injection, 671 benign) and 0 `test_native_es` rows.
- Dev (`spanish.jsonl`): from the filtered **test** rows, `random.Random(1337)` samples `min(300, ⌊n/2⌋)` rows per label,
  injection first: 52 of 105 injections and 300 of 671 benign. File order is kept within each label.
- Holdout (`spanish-holdout.jsonl`): the filtered `test_native_es` rows (0 at this revision), then every filtered test row
  that dev did not pick (53 injection, 371 benign), in file order.
- `source` = `edithngalame/<family>/<license>`. Licenses kept: dev 227 CC-BY-SA-3.0, 113 MIT, 12 Apache-2.0;
  holdout 282 CC-BY-SA-3.0, 131 MIT, 11 Apache-2.0. CC-BY-SA-3.0 rows (Dolly-derived benign text) are share-alike.
- Dev and holdout share no text (checked by the script).
- Caveats: almost all Spanish text is **machine-translated** (opus-mt-en-es or Dolly's own translation). Injections are only
  `gandalf` (94) and `jailbreak_roleplay` (11) families; the Spanish BIPIA email attacks are all in the unused
  `test_heldout` split. The half split was chosen because the test split has only 105 usable Spanish injections; a
  300-row dev cap would leave the holdout with no positives. Benign BIPIA emails come in clean and
  "benign insert" versions of the same email, which can land on both sides of the split.

## Sealed sets (added 2026-10-16)

Why: dev and holdout are both read in every iteration, so the holdout is reused hundreds of times
and decisions slowly adapt to it (Dwork et al. 2015, "The reusable holdout", arXiv:1506.02629;
Blum & Hardt 2015, "The Ladder", arXiv:1502.04585). A sealed set is read **once per round, at round
end**, and its result is recorded in that round's report-branch `COMPARISON.md`. It is never an input
to `safeanalyze gate` (the gate ignores every `*-sealed` set), to keep/revert, or to rule design;
its FP/FN files must not be read.

Rules (all script-checked; `fetch_eval.sh` exits non-zero on a violation):

- Built after every dev, holdout and stress set, each with a fresh `random.Random(1337)`, so the
  existing files stay byte-identical. Same sources, revisions and sha256s as the dev/holdout files.
- Candidate rows exclude every text in any dev, holdout or stress set. Each sealed set shares no text
  with any dev, holdout or stress set or with another sealed set, except the llmail benign emails.
- **llmail-sealed**: 300 further attack submissions with the dev per-stratum allocation (same
  largest-remainder allocation and sorted-stratum order as dev/holdout), drawn from the attack submissions that
  no dev or holdout set picked. **Benign rows:** the dataset has no unused benign emails (all 160 distinct
  `emails_for_fp_tests.json` emails are in dev), so llmail-sealed reuses the same 160 benign rows, as
  llmail-holdout does. Its precision and FP counts are therefore not independent; only its **recall** is
  a sealed measurement. (A positives-only file was considered and rejected so that all three llmail
  files have the same shape and F1 scale.)
- **browsesafe-sealed**: 300 `yes` + 300 `no` rows from the 3,680 test rows that neither dev nor holdout
  picked and whose content is in no other set (1,224 `yes` / 1,256 `no` rows are unpicked before that text filter). Same ordering rules as dev.
- **toolresp-sealed**: from the **test** split (the holdout's split), the same group rule as the
  holdout, by injection string: a row is a candidate only if its `injection` string is in no dev or
  holdout row, its text is in no other set, and its `response_id` (the safe/injected pair key within the
  split) is used by no holdout row, so no sealed row is the clean or injected twin of a holdout row.
  646 injection / 937 safe candidates; 300 + 300 sampled (`injection` first). The 300 positives use
  209 distinct injection strings. The script checks that no injection string is shared with dev or holdout.
- **nemotron-sealed** (GUARD-ONLY): pairs from the union-find groups (shared normalized `injection_text`
  or clean text, see nemotron above) that have **no** pair in nemotron dev or holdout: 311 pairs in 268
  groups; 300 pairs sampled as for dev (positives, then their negatives). Wrappers: 228 `[PRIORITY MESSAGE]`,
  71 `URGENT`, 1 bare.
- **No sealed set** for deepset (dev and holdout use the whole test and train splits), spanish (holdout
  takes every filtered test row dev did not pick), multilingual (every row is split between dev and
  holdout) or promptprotection-stress (report-only, no split).
- Round-5 end reading (v0.3.46 detection logic, 2026-10-16; `/tmp/safeanalyze-iter/v0.3.46-baseline/sealed/`):
  llmail-sealed P 1.000 R 0.753 F1 0.859; browsesafe-sealed P 0.895 R 0.680 F1 0.773 (FP 24);
  toolresp-sealed P 0.826 R 0.473 F1 0.602 (FP 30); nemotron-sealed P 1.000 R 0.137 F1 0.240.

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
