#!/usr/bin/env bash
# Download pinned prompt-injection evaluation datasets and convert them to
# testdata/eval/{deepset,llmail,browsesafe}.jsonl (dev) and
# testdata/eval/{deepset,llmail,browsesafe}-holdout.jsonl (holdout) for
# `safeanalyze eval`. Holdout sets share no text with any dev set, except
# that llmail-holdout reuses the llmail benign emails (no unused ones exist).
#
# Every download is pinned to a Hugging Face commit and verified by sha256.
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

python3 - "$CACHE" "$OUT" "$SEED" <<'PY'
import gzip, json, random, struct, sys

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
PY

(cd "$OUT" && sha256sum deepset.jsonl llmail.jsonl browsesafe.jsonl \
  deepset-holdout.jsonl llmail-holdout.jsonl browsesafe-holdout.jsonl)
