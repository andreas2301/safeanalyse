#!/usr/bin/env bash
# Download pinned prompt-injection evaluation datasets and convert them to
# testdata/eval/<set>.jsonl (dev) and testdata/eval/<set>-holdout.jsonl
# (holdout) for `safeanalyze eval`, with <set> in deepset, llmail,
# browsesafe, toolresp, nemotron, multilingual and spanish, plus the
# report-only stress set testdata/eval/promptprotection-stress.jsonl, and
# sealed sets testdata/eval/<set>-sealed.jsonl for llmail, browsesafe,
# toolresp and nemotron (disjoint from every dev and holdout set except the
# reused llmail benign emails; evaluated once per round end only).
# Holdout sets share no text with any dev set, except that llmail-holdout
# reuses the llmail benign emails (no unused ones exist). New sets share no
# text with the deepset/llmail/browsesafe dev and holdout sets.
#
# Every download is pinned to a Hugging Face (or GitHub) commit and verified
# by sha256.
# Sampling uses a fixed seed, so repeated runs produce byte-identical JSONL.
# Requires: bash, curl, sha256sum, python3 (stdlib only). No HF token needed.
# See testdata/eval/SOURCES.md for licenses and sampling rules.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OUT="$ROOT/testdata/eval"
CACHE="${SAFEANALYZE_EVAL_CACHE:-$OUT/.cache}"
SEED=1337
mkdir -p "$OUT" "$CACHE"

DEEPSET_REV=4f61ecb038e9c3fb77e21034b22511b523772cdd
LLMAIL_REV=1063bdf01ec8762b812d5e06ee768a06faa5a6f7
BROWSESAFE_REV=b506fb5bc7fd4472c8738055a67a0ef6406afdc9
TOOLRESP_REV=39050de797f2694f0d72e1086003154de1919379
NEMOTRON_REV=d738d4f361cc38bb4d7a42b9066776dade5332f5
PROMPTPROTECTION_REV=3c808d860711864f5dae99fb05f75a8139b8107d
PURPLELLAMA_REV=172c1074069eb88ec834124272c1b1c4f8893445
YANISMIRAOUI_REV=bd55359f2f332afc35f277ac3dd08f7111b024c9
EDITHNGALAME_REV=3e7180fbe4207547e3d0cebc44cedf6bc943a117

# fetch <dest-name> <url> <sha256>
fetch() {
  local dest="$CACHE/$1" url="$2" sum="$3"
  if [[ -f "$dest" ]] && echo "$sum  $dest" | sha256sum -c --status; then
    echo "cached   $1"
    return
  fi
  echo "fetching $1"
  curl -fsSL --retry 3 -o "$dest.part" "$url"
  if ! echo "$sum  $dest.part" | sha256sum -c --status; then
    echo "sha256 mismatch for $1 (expected $sum)" >&2
    rm -f "$dest.part"
    exit 1
  fi
  mv "$dest.part" "$dest"
}

HF=https://huggingface.co/datasets
fetch deepset-test.parquet \
  "$HF/deepset/prompt-injections/resolve/$DEEPSET_REV/data/test-00000-of-00001-701d16158af87368.parquet" \
  39ac797cabc157eeed58435a08593b2952bb6cb16fc394a2d383f447cc7b246e
fetch deepset-train.parquet \
  "$HF/deepset/prompt-injections/resolve/$DEEPSET_REV/data/train-00000-of-00001-9564e8b05b4757ab.parquet" \
  2e10bc7ab30f542c97e4e83e2a5683000b5057d25ec10908784c631d44124c04
fetch llmail-labelled_unique_submissions_phase2.json \
  "$HF/microsoft/llmail-inject-challenge/resolve/$LLMAIL_REV/data/labelled_unique_submissions_phase2.json" \
  f89af984e345430c3b357903890e30867bf4676f4ef10c138cc7bad218e890b8
fetch llmail-emails_for_fp_tests.json \
  "$HF/microsoft/llmail-inject-challenge/resolve/$LLMAIL_REV/data/emails_for_fp_tests.json" \
  4ddd950b5dbaa8548f5597c886d8e09a051ba07f80a9291fdcca9c2397d22abe
fetch browsesafe-test.parquet \
  "$HF/perplexity-ai/browsesafe-bench/resolve/$BROWSESAFE_REV/test.parquet" \
  00cbad96b60fee46e016d79af6981fb221384c61f12cf28b4f04b5a6420573d0
fetch toolresp-train.csv \
  "$HF/rgeada/tool-response-injections/resolve/$TOOLRESP_REV/tool_response_injections_train.csv" \
  4d7965bae0d0b1db3c50377596893923521f47a0119b0a21b83437617f387e6a
fetch toolresp-test.csv \
  "$HF/rgeada/tool-response-injections/resolve/$TOOLRESP_REV/tool_response_injections_test.csv" \
  ab5d1f4435b9652f74e9c09a5774d306b09bdce4bb96fe2f348c66eafaffbfc9
fetch nemotron-train.jsonl \
  "$HF/nvidia/Nemotron-RL-Agentic-Indirect-Prompt-Injection-v1/resolve/$NEMOTRON_REV/train.jsonl" \
  3329da17564a7eb287e2730fc7d6956e1f4fe51e8950ac4f110b3c37e78cf3b9
fetch promptprotection-agent-flows.jsonl \
  "$HF/promptprotection/agent-security-datasets/resolve/$PROMPTPROTECTION_REV/agent-flows.jsonl" \
  e2d46e679e89f0814f9f78a65226fccd5ba22897f164342a2f5f3e293dbc79fb
fetch promptprotection-attacks.jsonl \
  "$HF/promptprotection/agent-security-datasets/resolve/$PROMPTPROTECTION_REV/attacks.jsonl" \
  840620d8685dd6523327fe6d4ab8895f9d33e54f8dd73fd89e12ac9dc83502f2
fetch promptprotection-benign-hard.jsonl \
  "$HF/promptprotection/agent-security-datasets/resolve/$PROMPTPROTECTION_REV/benign-hard.jsonl" \
  bd12c6965d797eb470a57b82ed29c455d41516bc2c02b0c235ace799ac193e36
fetch cyberseceval-multilingual.json \
  "https://raw.githubusercontent.com/meta-llama/PurpleLlama/$PURPLELLAMA_REV/CybersecurityBenchmarks/datasets/prompt_injection/prompt_injection_multilingual_machine_translated.json" \
  f018adfc7b1f766f98ce4dd4c8e34e97a8a88f402a2b63576652027e02c2a15f
fetch cyberseceval-en.json \
  "https://raw.githubusercontent.com/meta-llama/PurpleLlama/$PURPLELLAMA_REV/CybersecurityBenchmarks/datasets/prompt_injection/prompt_injection.json" \
  069e4d5d36f6d19f972a3bbc65df840cc729354f89358a9031aeb44d95b18a9a
fetch yanismiraoui-prompt_injections.csv \
  "$HF/yanismiraoui/prompt_injections/resolve/$YANISMIRAOUI_REV/prompt_injections.csv" \
  f4843f1841fa19b980f804796a68fc72f06841775eaba2723c768c7d772aabad
fetch edithngalame-test.parquet \
  "$HF/edithngalame/prompt-injection-en-es/resolve/$EDITHNGALAME_REV/data/test-00000-of-00001.parquet" \
  919f9ad2ae1d9e3ffed8ab293cc53e050459c71b2a0b51fe621e27ca749112dc
fetch edithngalame-test_native_es.parquet \
  "$HF/edithngalame/prompt-injection-en-es/resolve/$EDITHNGALAME_REV/data/test_native_es-00000-of-00001.parquet" \
  a267ba8c4f2785ff3a5bc31f5a0a9ad1cf02f4ddf3617c376ceca4009ab1c85f

python3 - "$CACHE" "$OUT" "$SEED" <<'PY'
import csv, gzip, hashlib, json, random, re, struct, sys, unicodedata

cache, out, seed = sys.argv[1], sys.argv[2], int(sys.argv[3])

# --- Minimal stdlib-only Parquet reader (flat schemas, BYTE_ARRAY/INT64,
# --- PLAIN and dictionary encodings, data page v1/v2, snappy/gzip/zstd).

class Thrift:
    """Thrift compact protocol decoder returning {field_id: value} dicts."""
    def __init__(self, buf, pos=0):
        self.b, self.p = buf, pos
    def byte(self):
        v = self.b[self.p]; self.p += 1; return v
    def varint(self):
        r = s = 0
        while True:
            x = self.byte(); r |= (x & 0x7F) << s; s += 7
            if not x & 0x80:
                return r
    def zigzag(self):
        v = self.varint(); return (v >> 1) ^ -(v & 1)
    def value(self, t):
        if t in (1, 2): return t == 1
        if t == 3: return struct.unpack('b', bytes([self.byte()]))[0]
        if t in (4, 5, 6): return self.zigzag()
        if t == 7:
            v = struct.unpack('<d', self.b[self.p:self.p + 8])[0]; self.p += 8; return v
        if t == 8:
            n = self.varint(); v = self.b[self.p:self.p + n]; self.p += n; return v
        if t in (9, 10):
            h = self.byte(); n = h >> 4; et = h & 0x0F
            if n == 15: n = self.varint()
            return [self.value(et) if et not in (1, 2) else self.byte() == 1 for _ in range(n)]
        if t == 12: return self.struct()
        raise ValueError('unsupported thrift type %d' % t)
    def struct(self):
        res, last = {}, 0
        while True:
            h = self.byte()
            if h == 0:
                return res
            t = h & 0x0F; d = h >> 4
            fid = last + d if d else self.zigzag()
            last = fid
            res[fid] = self.value(t)

def snappy_decompress(src):
    pos = 0; n = 0; shift = 0
    while True:
        x = src[pos]; pos += 1; n |= (x & 0x7F) << shift; shift += 7
        if not x & 0x80: break
    dst = bytearray()
    while pos < len(src):
        tag = src[pos]; pos += 1; kind = tag & 3
        if kind == 0:
            ln = tag >> 2
            if ln >= 60:
                nb = ln - 59; ln = int.from_bytes(src[pos:pos + nb], 'little'); pos += nb
            ln += 1; dst += src[pos:pos + ln]; pos += ln; continue
        if kind == 1:
            ln = ((tag >> 2) & 7) + 4; off = ((tag >> 5) << 8) | src[pos]; pos += 1
        elif kind == 2:
            ln = (tag >> 2) + 1; off = int.from_bytes(src[pos:pos + 2], 'little'); pos += 2
        else:
            ln = (tag >> 2) + 1; off = int.from_bytes(src[pos:pos + 4], 'little'); pos += 4
        start = len(dst) - off
        if off >= ln:
            dst += dst[start:start + ln]
        else:
            for i in range(ln): dst.append(dst[start + i])
    if len(dst) != n: raise ValueError('snappy length mismatch')
    return bytes(dst)

def decompress(codec, data):
    if codec == 0: return data
    if codec == 1: return snappy_decompress(data)
    if codec == 2: return gzip.decompress(data)
    if codec == 6:
        from compression import zstd  # Python >= 3.14
        return zstd.decompress(data)
    raise ValueError('unsupported parquet codec %d' % codec)

def rle_hybrid(buf, bit_width, count):
    out, p, bw = [], 0, (bit_width + 7) // 8
    while len(out) < count and p < len(buf):
        h = 0; s = 0
        while True:
            x = buf[p]; p += 1; h |= (x & 0x7F) << s; s += 7
            if not x & 0x80: break
        if h & 1:
            groups = h >> 1; nbytes = groups * bit_width
            bits = int.from_bytes(buf[p:p + nbytes], 'little'); p += nbytes
            mask = (1 << bit_width) - 1
            for i in range(groups * 8):
                out.append((bits >> (i * bit_width)) & mask)
        else:
            run = h >> 1; v = int.from_bytes(buf[p:p + bw], 'little'); p += bw
            out.extend([v] * run)
    return out[:count]

def plain(ptype, buf, count):
    vals, p = [], 0
    for _ in range(count):
        if ptype == 6:  # BYTE_ARRAY
            n = struct.unpack_from('<i', buf, p)[0]; p += 4
            vals.append(buf[p:p + n].decode('utf-8')); p += n
        elif ptype == 2:  # INT64
            vals.append(struct.unpack_from('<q', buf, p)[0]); p += 8
        elif ptype == 1:  # INT32
            vals.append(struct.unpack_from('<i', buf, p)[0]); p += 4
        else:
            raise ValueError('unsupported physical type %d' % ptype)
    return vals

def read_parquet(path):
    data = open(path, 'rb').read()
    if data[:4] != b'PAR1' or data[-4:] != b'PAR1': raise ValueError('not parquet: ' + path)
    flen = struct.unpack('<i', data[-8:-4])[0]
    meta = Thrift(data, len(data) - 8 - flen).struct()
    schema = meta[2]
    optional = {el[4].decode(): el.get(3, 0) == 1 for el in schema[1:]}
    cols = {}
    for rg in meta[4]:
        for cc in rg[1]:
            md = cc[3]; ptype = md[1]; name = md[3][0].decode(); codec = md[4]
            remaining = md[5]
            pos = md.get(11) or md[9]
            if md.get(11) and md[9] < pos: pos = md[9]
            dictionary, vals = None, cols.setdefault(name, [])
            while remaining > 0:
                t = Thrift(data, pos); ph = t.struct(); body = data[t.p:t.p + ph[3]]; pos = t.p + ph[3]
                ptype_page = ph[1]
                if ptype_page == 2:  # DICTIONARY_PAGE
                    dictionary = plain(ptype, decompress(codec, body), ph[7][1]); continue
                if ptype_page == 0:  # DATA_PAGE v1
                    dh = ph[5]; n = dh[1]; enc = dh[2]; page = decompress(codec, body); p = 0
                    if optional[name]:
                        ln = struct.unpack_from('<i', page, 0)[0]
                        defs = rle_hybrid(page[4:4 + ln], 1, n); p = 4 + ln
                    else:
                        defs = [1] * n
                    payload = page[p:]
                elif ptype_page == 3:  # DATA_PAGE v2
                    dh = ph[8]; n = dh[1]; enc = dh[4]; dl = dh[5]; rl = dh[6]
                    levels = body[:rl + dl]; rest = body[rl + dl:]
                    if dh.get(7, True): rest = decompress(codec, rest)
                    defs = rle_hybrid(levels[rl:], 1, n) if optional[name] else [1] * n
                    payload = rest
                else:
                    continue
                present = sum(defs)
                if enc in (2, 8):
                    idx = rle_hybrid(payload[1:], payload[0], present) if present else []
                    got = [dictionary[i] for i in idx]
                elif enc == 0:
                    got = plain(ptype, payload, present)
                else:
                    raise ValueError('unsupported encoding %d' % enc)
                it = iter(got)
                vals.extend(next(it) if d else None for d in defs)
                remaining -= n
    if any(len(v) != meta[3] for v in cols.values()): raise ValueError('row count mismatch in ' + path)
    return cols

def write_jsonl(name, rows):
    path = '%s/%s.jsonl' % (out, name)
    with open(path, 'w', encoding='utf-8', newline='\n') as f:
        for text, label, source in rows:
            f.write(json.dumps({'text': text, 'label': label, 'source': source}, ensure_ascii=False) + '\n')
    pos = sum(1 for r in rows if r[1] == 1)
    print('wrote %s: %d samples (%d injection, %d benign)' % (path, len(rows), pos, len(rows) - pos))

# deepset/prompt-injections: full test split, label already 1=injection.
c = read_parquet(cache + '/deepset-test.parquet')
dev_deepset = [(t, int(l), 'deepset') for t, l in zip(c['text'], c['label'])]
write_jsonl('deepset', dev_deepset)

# LLMail-Inject phase 2: stratified (by judge_category, or api_triggered)
# sample of 300 attack submissions + all benign FP-test emails.
subs = json.load(open(cache + '/llmail-labelled_unique_submissions_phase2.json', encoding='utf-8'))
strata = {}
for text in sorted(subs):
    v = subs[text]
    if str(v.get('attack_attempt')) != 'True':
        continue
    key = v.get('judge_category') or v.get('reason') or 'unknown'
    strata.setdefault(key, []).append(text)
total, want = sum(len(v) for v in strata.values()), 300
alloc = {k: want * len(v) // total for k, v in strata.items()}
rema = sorted(strata, key=lambda k: (-(want * len(strata[k]) % total), k))
for k in rema[:want - sum(alloc.values())]:
    alloc[k] += 1
rng = random.Random(seed)
rows = []
for k in sorted(strata):
    for text in rng.sample(strata[k], alloc[k]):
        rows.append((text, 1, 'llmail/' + k))
# Drop exact duplicate benign emails (file order kept) so repeats don't
# get extra weight in false-positive counts.
for text in dict.fromkeys(json.load(open(cache + '/llmail-emails_for_fp_tests.json', encoding='utf-8'))):
    rows.append((text, 0, 'llmail/benign'))
write_jsonl('llmail', rows)
rows_llmail = rows

# BrowseSafe-Bench: test split, seeded sample of 300 label=yes + 300 label=no.
c = read_parquet(cache + '/browsesafe-test.parquet')
rng = random.Random(seed)
rows = []
for lab, num in (('yes', 1), ('no', 0)):
    idx = [i for i, l in enumerate(c['label']) if l == lab]
    for i in sorted(rng.sample(idx, 300)):
        rows.append((c['content'][i], num, 'browsesafe'))
write_jsonl('browsesafe', rows)
dev = {'deepset': dev_deepset, 'llmail': rows_llmail, 'browsesafe': rows}

# ---------------- Holdout sets (disjoint from every dev set) ----------------
dev_texts = {r[0] for rs in dev.values() for r in rs}

# deepset holdout: the full train split, minus rows whose text also appears
# in a dev set (the upstream train/test splits share a few texts).
c = read_parquet(cache + '/deepset-train.parquet')
hold = {}
hold['deepset'] = [(t, int(l), 'deepset') for t, l in zip(c['text'], c['label']) if t not in dev_texts]
print('deepset-holdout: dropped %d train rows that overlap dev' % (len(c['text']) - len(hold['deepset'])))

# LLMail holdout: 300 more attack submissions with the same per-stratum
# allocation, drawn from the submissions dev did not pick, + the benign emails.
rng = random.Random(seed)
rows_h = []
for k in sorted(strata):
    pool = [t for t in strata[k] if t not in dev_texts]
    if len(pool) < alloc[k]:
        raise SystemExit('llmail stratum %r too small for holdout' % k)
    for text in rng.sample(pool, alloc[k]):
        rows_h.append((text, 1, 'llmail/' + k))
rows_h.extend(r for r in rows_llmail if r[1] == 0)
hold['llmail'] = rows_h

# BrowseSafe holdout: seeded 300 yes + 300 no from test rows dev did not pick.
c = read_parquet(cache + '/browsesafe-test.parquet')
rng = random.Random(seed)
rows_h = []
for lab, num in (('yes', 1), ('no', 0)):
    idx = [i for i, l in enumerate(c['label']) if l == lab and c['content'][i] not in dev_texts]
    for i in sorted(rng.sample(idx, 300)):
        rows_h.append((c['content'][i], num, 'browsesafe'))
hold['browsesafe'] = rows_h

for name, rs in hold.items():
    write_jsonl(name + '-holdout', rs)

# Verify disjointness: no holdout text may appear in any dev set, except the
# llmail benign emails that llmail-holdout deliberately reuses.
reused = {r[0] for r in rows_llmail if r[1] == 0}
for name, rs in hold.items():
    shared = {r[0] for r in rs} & dev_texts
    if name == 'llmail':
        if any(r[0] in dev_texts and r[1] == 1 for r in rs):
            raise SystemExit('llmail-holdout attack text overlaps dev')
        shared -= reused
    if shared:
        raise SystemExit('%s-holdout shares %d texts with dev' % (name, len(shared)))
print('disjointness check passed (llmail-holdout reuses %d benign dev emails)' % len(reused))

# ---------------- Additional sets (see SOURCES.md) ----------------
# Texts already in a deepset/llmail/browsesafe dev or holdout set are removed
# from every new candidate pool before sampling, so new sets never repeat them.
old_texts = dev_texts | {r[0] for rs in hold.values() for r in rs}
new_dev, new_hold = {}, {}
csv.field_size_limit(1 << 30)

def fresh(rows, name, exclude):
    """Drop rows whose text is in `exclude` or repeats an earlier row."""
    seen, kept = set(), []
    for r in rows:
        if r[0] in exclude or r[0] in seen:
            continue
        seen.add(r[0]); kept.append(r)
    if len(kept) != len(rows):
        print('%s: dropped %d candidates (exact text in an earlier set or repeated)' % (name, len(rows) - len(kept)))
    return kept

def take(rows, n, rng):
    """Seeded sample of n rows (all if fewer), kept in input order."""
    if len(rows) <= n:
        return list(rows)
    return [rows[i] for i in sorted(rng.sample(range(len(rows)), n))]

def read_csv(path):
    with open(path, encoding='utf-8', newline='') as f:
        return list(csv.DictReader(f))

# rgeada/tool-response-injections: seeded 300 injection + 300 safe from train
# (dev) and from test (holdout). Injection rows whose `injection` string is not
# in `text` (raw or JSON-escaped) were never actually injected and are dropped.
def injected(r):
    inj = r['injection']
    return any(v in r['text'] for v in (inj, json.dumps(inj)[1:-1], json.dumps(inj, ensure_ascii=False)[1:-1]))
for name, split, excl in (('toolresp', 'train', old_texts), ('toolresp-holdout', 'test', None)):
    if excl is None:
        excl = old_texts | {r[0] for rs in new_dev.values() for r in rs}
    rows = read_csv(cache + '/toolresp-%s.csv' % split)
    bad = sum(1 for r in rows if r['label'] == 'injection' and not injected(r))
    print('%s: dropped %d injection rows whose injection string is not in the text' % (name, bad))
    rows = [r for r in rows if r['label'] != 'injection' or injected(r)]
    rng = random.Random(seed)
    out_rows = []
    for lab, num in (('injection', 1), ('safe', 0)):
        pool = fresh([(r['text'], num, 'toolresp', r['injection']) for r in rows if r['label'] == lab], name, excl)
        out_rows += take(pool, 300, rng)
    (new_dev if split == 'train' else new_hold)['toolresp'] = out_rows
d, h = new_dev['toolresp'], new_hold['toolresp']
shared_text = {r[0] for r in d} & {r[0] for r in h}
shared_inj = {r[3] for r in d if r[3]} & {r[3] for r in h if r[3]}
if shared_text or shared_inj:
    raise SystemExit('toolresp dev/holdout share %d texts, %d injection strings' % (len(shared_text), len(shared_inj)))
print('toolresp: dev/holdout share 0 texts and 0 injection strings')
new_dev['toolresp'] = [r[:3] for r in d]
new_hold['toolresp'] = [r[:3] for r in h]

# Nemotron-RL-Agentic-Indirect-Prompt-Injection-v1: positive = the environment
# string that contains injection_text; negative = that string with
# injection_text removed. Pairs whose clean text still holds a wrapper marker
# or an injection line are dropped. Records that share a normalized
# injection_text or clean text are grouped (union-find); each group goes to
# dev or holdout by sha256(str(lowest record id in the group)) parity.
def norm(t):
    return ' '.join(unicodedata.normalize('NFKC', t).casefold().split())
def leaves(o):
    if isinstance(o, dict):
        for v in o.values(): yield from leaves(v)
    elif isinstance(o, list):
        for v in o: yield from leaves(v)
    elif isinstance(o, str):
        yield o
NEM_MARKERS = ('[PRIORITY MESSAGE]', '[/PRIORITY MESSAGE]', 'URGENT: Before completing any other tasks')
skip = {'not_found': 0, 'ambiguous': 0, 'empty_clean': 0, 'leftover_marker_or_fragment': 0}
with open(cache + '/nemotron-train.jsonl', encoding='utf-8') as f:
    recs = sorted((json.loads(l) for l in f if l.strip()), key=lambda r: r['id'])
cands = []
for r in recs:
    inj = r['injection']['injection_text']
    hits = list(dict.fromkeys(v for v in leaves(r['environment']) if inj in v))
    if not hits:
        skip['not_found'] += 1; continue
    if len(hits) > 1:
        skip['ambiguous'] += 1; continue
    clean = hits[0].replace(inj, '')
    if not clean.strip():
        skip['empty_clean'] += 1; continue
    frags = [l.strip() for l in inj.split('\n') if len(l.strip()) >= 20]
    if any(m in clean for m in NEM_MARKERS) or any(x in clean for x in frags):
        skip['leftover_marker_or_fragment'] += 1; continue
    cands.append((r['id'], hits[0], clean, 'nemotron/' + r['domain'], norm(inj)))
parent, owner = {}, {}
def find(x):
    while parent[x] != x:
        parent[x] = parent[parent[x]]; x = parent[x]
    return x
for rid, pos, clean, src, ninj in cands:
    parent[rid] = rid
    for key in (('inj', ninj), ('clean', norm(clean))):
        if key in owner:
            a, b = find(rid), find(owner[key])
            parent[max(a, b)] = min(a, b)
        else:
            owner[key] = rid
pairs = {0: [], 1: []}
for rid, pos, clean, src, ninj in cands:
    part = int(hashlib.sha256(str(find(rid)).encode()).hexdigest(), 16) % 2
    pairs[part].append((pos, clean, src, ninj))
print('nemotron: %d records, %d groups, %d dev pairs, %d holdout pairs, skipped %s' % (
    len(recs), len({find(c[0]) for c in cands}), len(pairs[0]), len(pairs[1]), skip))
for part, name in ((0, 'nemotron'), (1, 'nemotron-holdout')):
    excl = old_texts | ({r[0] for rs in new_dev.values() for r in rs} if part else set())
    ok, seen = [], set(excl)
    for p in pairs[part]:
        if p[0] in seen or p[1] in seen:
            continue
        seen.update(p[:2]); ok.append(p)
    if len(ok) != len(pairs[part]):
        print('%s: dropped %d pairs (text in an earlier set or an earlier pair)' % (name, len(pairs[part]) - len(ok)))
    chosen = take(ok, 300, random.Random(seed))
    rows = fresh([(p[0], 1, p[2]) for p in chosen] + [(p[1], 0, p[2]) for p in chosen], name, set())
    (new_hold if part else new_dev)['nemotron'] = rows
shared = {p[3] for p in pairs[0]} & {p[3] for p in pairs[1]}
shared |= {norm(p[1]) for p in pairs[0]} & {norm(p[1]) for p in pairs[1]}
if shared:
    raise SystemExit('nemotron dev/holdout share %d injection or clean texts' % len(shared))
print('nemotron: dev/holdout share 0 normalized injection_text and 0 clean texts')
nem_cands = [(find(c[0]),) + c[1:] for c in cands]  # (group, positive, clean, source, ninj), for the sealed set

# Multilingual recall-only sets: CyberSecEval multilingual user_input +
# yanismiraoui/prompt_injections (language guessed from stopwords, since the
# CSV has no language column). CyberSecEval logic-violating rows are dropped.
# Exact duplicates dropped (CyberSecEval first). Split by sha256(group key)
# parity, even -> dev: all translations of one CyberSecEval base case share a
# key; a yanismiraoui row is its own group. Holdout rows that are near
# duplicates (char 5-gram Jaccard >= 0.8) of a same-language dev row are dropped.
YM_STOP = {
    'en': 'the you your and to of is are what for with this that me my all in on be it not do previous instructions ignore',
    'fr': 'le la les de des et est vous votre tu ton ta tes un une pour que qui dans sur pas ne mes instructions ignore quelle quel',
    'de': 'der die das und ist du dein deine sie ihr nicht ein eine zu mit auf für ich mir mich von alle anweisungen ignoriere was',
    'es': 'el la los las de y es tu tus usted que en un una para por con no del al mis instrucciones ignora qué cuál',
    'pt': 'o a os as de e é você seu sua tu teu tua que em um uma para por com não do da dos das ao instruções ignore qual',
    'it': 'il lo la i gli le di e è tu tuo tua che in un una per con non del della dei al istruzioni ignora qual quale',
    'ro': 'și este tu ta tău să în un o pentru cu nu de la care instrucțiunile ignoră ce ești din pe',
}
YM_STOP = {k: set(v.split()) for k, v in YM_STOP.items()}
YM_HINTS = (('[ăşșţț]', 'ro', 3), ('[ãõç]', 'pt', 2), ('[ñ¿¡]', 'es', 3), ('[äöüß]', 'de', 3),
            ('[èàù]', 'it', 1), ('[èàù]', 'fr', 1), ('[êëœç]', 'fr', 1))
def guess_lang(t):
    t = t.lower()
    words = re.findall(r'[^\W\d_]+', t)
    sc = {k: sum(w in v for w in words) for k, v in YM_STOP.items()}
    for pat, k, n in YM_HINTS:
        if re.search(pat, t): sc[k] += n
    best = max(sorted(sc), key=lambda k: sc[k])
    return best if sc[best] else 'und'
# CyberSecEval multilingual rows 0..752 are three translations of each English
# base case in order (base = prompt_id // 3); rows 753..1003 are the Hindi
# translations (base = prompt_id - 753). Checked against the English file.
cse = json.load(open(cache + '/cyberseceval-multilingual.json', encoding='utf-8'))
cse_en = json.load(open(cache + '/cyberseceval-en.json', encoding='utf-8'))
def cse_base(pid):
    return pid // 3 if pid < 3 * len(cse_en) else pid - 3 * len(cse_en)
if len(cse) != 4 * len(cse_en):
    raise SystemExit('cyberseceval: unexpected row count %d' % len(cse))
for r in cse:
    e = cse_en[cse_base(int(r['prompt_id']))]
    if any(r[k] != e[k] for k in ('injection_variant', 'injection_type', 'risk_category')):
        raise SystemExit('cyberseceval: row %s does not match its English base' % r['prompt_id'])
def tok_hit(r):
    e = cse_en[cse_base(int(r['prompt_id']))]
    toks = re.findall(r"[A-Za-z0-9@._-]*[0-9@][A-Za-z0-9@._-]*", e['judge_question'] + ' ' + e['test_case_prompt'])
    if not toks:
        return None
    hay = ' '.join(r[k] for k in ('judge_question', 'test_case_prompt', 'user_input'))
    return any(t.strip("'.") in hay for t in toks)
hits = [h for h in map(tok_hit, cse) if h is not None]
print('cyberseceval: cross-language check: %d/%d rows with a digit/@ token in the English base contain it' % (sum(hits), len(hits)))
def grams(t, n=5):
    t = norm(t)
    return {t[i:i + n] for i in range(max(1, len(t) - n + 1))}
def jacc(a, b):
    return len(a & b) / len(a | b) if a or b else 1.0
# Union English base cases whose user_input are near duplicates.
eg = [grams(e['user_input']) for e in cse_en]
bp = list(range(len(cse_en)))
def bfind(x):
    while bp[x] != x:
        bp[x] = bp[bp[x]]; x = bp[x]
    return x
for i in range(len(cse_en)):
    for j in range(i + 1, len(cse_en)):
        if jacc(eg[i], eg[j]) >= 0.8:
            a, b = bfind(i), bfind(j); bp[max(a, b)] = min(a, b)
CSE_LANG = {'Spanish': 'es', 'French': 'fr', 'German': 'de', 'Italian': 'it', 'Portuguese': 'pt'}
logic = sum(1 for r in cse if r['risk_category'] == 'logic-violating')
cands = [(r['user_input'], 1, 'cyberseceval/' + r['speaking_language'],
          'cse:%d' % bfind(cse_base(int(r['prompt_id']))), CSE_LANG.get(r['speaking_language'], r['speaking_language']))
         for r in cse if r['risk_category'] != 'logic-violating']
print('cyberseceval: dropped %d logic-violating rows' % logic)
for r in read_csv(cache + '/yanismiraoui-prompt_injections.csv'):
    t = r['prompt_injections']; lang = guess_lang(t)
    cands.append((t, 1, 'yanismiraoui/' + lang, 'ym:' + t, lang))
cands = fresh([r for r in cands if r[0].strip()], 'multilingual', old_texts | {r[0] for rs in new_dev.values() for r in rs})
even = lambda t: int(hashlib.sha256(t.encode()).hexdigest(), 16) % 2 == 0
ml_dev = [r for r in cands if even(r[3])]
ml_hold = [r for r in cands if not even(r[3])]
dev_grams = {}
for r in ml_dev:
    dev_grams.setdefault(r[4], []).append(grams(r[0]))
near = [r for r in ml_hold if any(jacc(grams(r[0]), g) >= 0.8 for g in dev_grams.get(r[4], ()))]
print('multilingual-holdout: dropped %d same-language near duplicates of dev rows' % len(near))
near = {r[0] for r in near}
if {r[3] for r in ml_dev if r[3].startswith('cse:')} & {r[3] for r in ml_hold if r[3].startswith('cse:')}:
    raise SystemExit('multilingual: a CyberSecEval base case is on both sides')
new_dev['multilingual'] = [r[:3] for r in ml_dev]
new_hold['multilingual'] = [r[:3] for r in ml_hold if r[0] not in near]

# edithngalame/prompt-injection-en-es, Spanish rows only. Drop deepset-derived
# rows (origin contains "deepset" or normalized text equal to any deepset
# train/test text) and rows whose license is not CC-BY-SA-3.0, MIT or Apache-2.0.
deepset_norm = {norm(t) for f in ('deepset-test', 'deepset-train') for t in read_parquet('%s/%s.parquet' % (cache, f))['text']}
PERMISSIVE = ('CC-BY-SA-3.0', 'MIT', 'Apache-2.0')
def spanish(split):
    c = read_parquet('%s/edithngalame-%s.parquet' % (cache, split))
    rows, drop = [], {'origin_deepset': 0, 'text_deepset': 0, 'license': 0}
    for i in range(len(c['text'])):
        if c['lang'][i] != 'es':
            continue
        if 'deepset' in c['origin'][i].lower():
            drop['origin_deepset'] += 1
        elif norm(c['text'][i]) in deepset_norm:
            drop['text_deepset'] += 1
        elif c['license'][i] not in PERMISSIVE:
            drop['license'] += 1
        else:
            rows.append((c['text'][i], int(c['label'][i]), 'edithngalame/%s/%s' % (c['family'][i], c['license'][i])))
    print('edithngalame %s: kept %d Spanish rows, dropped %s' % (split, len(rows), drop))
    return rows
test_es = fresh(spanish('test'), 'spanish', old_texts | {r[0] for rs in new_dev.values() for r in rs})
# Dev: seeded min(300, half) of each label from the test split, so that the
# holdout keeps the other half of the scarce Spanish test injections.
rng = random.Random(seed)
dev_es = []
for num in (1, 0):
    pool = [r for r in test_es if r[1] == num]
    dev_es += take(pool, min(300, len(pool) // 2), rng)
new_dev['spanish'] = dev_es
used = {r[0] for r in dev_es}
hold_es = spanish('test_native_es') + [r for r in test_es if r[0] not in used]
new_hold['spanish'] = fresh(hold_es, 'spanish-holdout', old_texts | {r[0] for rs in new_dev.values() for r in rs})

for name, rs in new_dev.items():
    write_jsonl(name, rs)
for name, rs in new_hold.items():
    write_jsonl(name + '-holdout', rs)

# promptprotection/agent-security-datasets: report-only stress set.
def jl(name):
    with open('%s/promptprotection-%s.jsonl' % (cache, name), encoding='utf-8') as f:
        return [json.loads(l) for l in f if l.strip()]
stress = [('\n\n'.join(s['text'] for s in r['sources']), 1 if r['label'] == 'attack' else 0,
           'promptprotection/agent-flows/' + r['scenario']) for r in jl('agent-flows')]
print('promptprotection agent-flows: %d rows without sources skipped' % sum(1 for r in stress if not r[0].strip()))
stress = [r for r in stress if r[0].strip()]
stress += [(r['text'], 1, 'promptprotection/attacks/' + r['technique']) for r in jl('attacks')]
stress += [(r['text'], 0, 'promptprotection/benign-hard/' + r['category']) for r in jl('benign-hard')]
write_jsonl('promptprotection-stress', fresh(stress, 'promptprotection-stress', old_texts))

# Final check: no new holdout text in any dev set, no new set text in any
# existing deepset/llmail/browsesafe set.
all_dev = dev_texts | {r[0] for rs in new_dev.values() for r in rs}
for name, rs in new_hold.items():
    if {r[0] for r in rs} & all_dev:
        raise SystemExit('%s-holdout shares text with a dev set' % name)
for name, rs in list(new_dev.items()) + [(n + '-holdout', r) for n, r in new_hold.items()]:
    if {r[0] for r in rs} & old_texts:
        raise SystemExit('%s shares text with an existing set' % name)
print('new-set disjointness check passed')

# ---------------- Sealed sets (disjoint from every dev and holdout set) ----------------
# Evaluated once per round end, never used for keep/revert or rule design.
# Each sealed set is drawn from rows that no dev, holdout or stress set uses,
# with the holdout's group-split rules, and a fresh random.Random(seed), so
# adding them leaves every dev/holdout file byte-identical.
used_texts = all_dev | old_texts | {r[0] for rs in new_hold.values() for r in rs} | {r[0] for r in stress}
sealed = {}

# LLMail: 300 more attack submissions with the dev per-stratum allocation, from
# submissions no dev or holdout set picked. No unused benign emails exist, so
# the 160 dev benign emails are reused (as in llmail-holdout): only recall is
# an independent sealed measurement.
rng = random.Random(seed)
rows_s = []
for k in sorted(strata):
    pool = [t for t in strata[k] if t not in used_texts]
    if len(pool) < alloc[k]:
        raise SystemExit('llmail stratum %r too small for sealed' % k)
    for text in rng.sample(pool, alloc[k]):
        rows_s.append((text, 1, 'llmail/' + k))
rows_s.extend(r for r in rows_llmail if r[1] == 0)
sealed['llmail'] = rows_s

# BrowseSafe: seeded 300 yes + 300 no from test rows no dev or holdout set picked.
c = read_parquet(cache + '/browsesafe-test.parquet')
rng = random.Random(seed)
rows_s = []
for lab, num in (('yes', 1), ('no', 0)):
    idx = [i for i, l in enumerate(c['label']) if l == lab and c['content'][i] not in used_texts]
    for i in sorted(rng.sample(idx, 300)):
        rows_s.append((c['content'][i], num, 'browsesafe'))
sealed['browsesafe'] = rows_s

# toolresp: from the test split (the holdout's split), rows whose injection
# string is in no dev or holdout row and whose response_id (the safe/injected
# pair key within the split) no holdout row uses; seeded 300 + 300.
rows = [r for r in read_csv(cache + '/toolresp-test.csv') if r['label'] != 'injection' or injected(r)]
hold_tr = {r[0] for r in new_hold['toolresp']}
hold_ids = {r['response_id'] for r in rows if r['text'] in hold_tr}
used_inj = {r[3] for r in d + h if r[3]}
rng = random.Random(seed)
rows_s = []
for lab, num in (('injection', 1), ('safe', 0)):
    pool = [(r['text'], num, 'toolresp', r['injection']) for r in rows if r['label'] == lab
            and r['response_id'] not in hold_ids and r['injection'] not in used_inj]
    pool = fresh(pool, 'toolresp-sealed', used_texts)
    print('toolresp-sealed: %d unused %s candidates' % (len(pool), lab))
    rows_s += take(pool, 300, rng)
if {r[3] for r in rows_s if r[3]} & used_inj:
    raise SystemExit('toolresp-sealed shares an injection string with dev/holdout')
sealed['toolresp'] = [r[:3] for r in rows_s]

# nemotron: pairs from union-find groups with no pair in nemotron dev or holdout.
nem_used = {r[0] for r in new_dev['nemotron'] + new_hold['nemotron']}
used_groups = {g for g, pos, clean, src, ninj in nem_cands if pos in nem_used or clean in nem_used}
ok, seen = [], set(used_texts)
for g, pos, clean, src, ninj in nem_cands:
    if g in used_groups or pos in seen or clean in seen:
        continue
    seen.update((pos, clean)); ok.append((pos, clean, src, ninj))
print('nemotron-sealed: %d unused pairs in %d unused groups' % (len(ok), len({g for g, *_ in nem_cands} - used_groups)))
if len(ok) >= 100:
    chosen = take(ok, 300, random.Random(seed))
    sealed['nemotron'] = [(p[0], 1, p[2]) for p in chosen] + [(p[1], 0, p[2]) for p in chosen]
    used_n = {norm(p[3]) for p in pairs[0] + pairs[1] if p[0] in nem_used} | {norm(p[1]) for p in pairs[0] + pairs[1] if p[0] in nem_used}
    if {p[3] for p in chosen} & used_n or {norm(p[1]) for p in chosen} & used_n:
        raise SystemExit('nemotron-sealed shares an injection or clean text with dev/holdout')
# deepset, multilingual and spanish use every usable row in dev or holdout, and
# promptprotection-stress has no split, so they have no sealed set.

for name, rs in sealed.items():
    write_jsonl(name + '-sealed', rs)
reused = {r[0] for r in rows_llmail if r[1] == 0}
for name, rs in sealed.items():
    shared = {r[0] for r in rs} & used_texts
    if name == 'llmail':
        shared -= reused
    if shared:
        raise SystemExit('%s-sealed shares %d texts with a dev, holdout or stress set' % (name, len(shared)))
names = sorted(sealed)
for i, a in enumerate(names):
    for b in names[i + 1:]:
        if {r[0] for r in sealed[a]} & {r[0] for r in sealed[b]}:
            raise SystemExit('%s-sealed and %s-sealed share text' % (a, b))
print('sealed disjointness check passed (llmail-sealed reuses %d benign dev emails)' % len(reused))
PY

(cd "$OUT" && sha256sum deepset.jsonl llmail.jsonl browsesafe.jsonl \
  toolresp.jsonl nemotron.jsonl multilingual.jsonl spanish.jsonl \
  deepset-holdout.jsonl llmail-holdout.jsonl browsesafe-holdout.jsonl \
  toolresp-holdout.jsonl nemotron-holdout.jsonl multilingual-holdout.jsonl \
  spanish-holdout.jsonl promptprotection-stress.jsonl \
  llmail-sealed.jsonl browsesafe-sealed.jsonl toolresp-sealed.jsonl \
  nemotron-sealed.jsonl)
