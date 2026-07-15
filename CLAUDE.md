# CLAUDE.md — safeanalyze contributor guide

## Project purpose

`safeanalyze` is a Go CLI security scanner that checks untrusted code before it is fed to AI assistants. It supports two operational modes:

- **Fast mode** — sub-100 ms deterministic checks for reverse-proxy/Squid integration.
- **Thorough mode** — full deterministic + stochastic + external-scanner suite for repository review.

## Development conventions

- Keep changes minimal and surgical. Every changed line must trace to a functional requirement.
- Match existing Go style, error wrapping, and package layout.
- Do not add speculative abstractions or unrelated cleanups.
- All new checks must implement `pkg/pipeline.Stage`.
- All new report formats must consume `pkg/report.Report`.
- Tests run with `go test ./...`. Builds must pass with `go build ./...`.
- Commits describe the functional change; do not include review notes or process metadata.

## Release process

1. Bump `pkg/version/version.go` on every meaningful improvement.
2. Update `CHANGELOG.md` with functional and non-functional changes.
3. Update `README.md` so the header, feature list, and config example match the new version.
4. Commit improvements to `master`.
5. Tag the commit: `git tag vX.Y.Z && git push origin master vX.Y.Z`.
6. Reports generated from a versioned build must record `safeanalyze_version`, `scan_mode`, and `duration_ms` in their metadata.
7. Report branches are named `report-<devicename>-<date>`. Device names are random Deadpool-style UUID phrases (e.g., `maximum-effort-chimichanga-9f7d`).

## Autoresearch improvement loop

To continuously improve detection on test data, follow this iterative cycle:

1. **Measure baseline** — run `safeanalyze scan --mode thorough` against the test corpus and record findings + durations.
2. **Premortem** — run `./scripts/premortem.sh` and ask: "If this improvement lands and the tool still fails in production, what most likely broke?" Document the biggest risks (false positives, latency blow-out, missing variants).
3. **Analyze gaps** — compare findings to known prompt-injection patterns. Look for false negatives and false positives.
4. **Hypothesize** — pick one concrete improvement: a new YARA rule, a tuned entropy threshold, a smaller/faster model, or a file-size limit.
5. **Implement** — make the smallest change that tests the hypothesis.
6. **Red-team** — run `./scripts/redteam.sh`. Try to evade the new check with rephrased, encoded, or multi-language injections. If it is trivially bypassed, revert or harden.
7. **Evaluate** — re-run the same test corpus. Did recall or precision improve? Did latency meet the mode's budget?
8. **Keep or revert** — if the metric improved, commit and bump the version. If not, revert and try another hypothesis.
9. **Repeat** until no measurable progress remains.

## Security hardening notes

- `safeanalyze clone` validates URLs to prevent git option injection (`-u`, `--upload-pack`, shell metacharacters). Permitted formats are http(s), ssh, git, and scp-style.
- External scanners run as separate processes with non-fatal error handling; a missing or crashing scanner must not abort the whole pipeline.
- Fast mode intentionally avoids ML inference, external scanners, and repository walking to keep latency deterministic.

## Fast-mode latency budget

- Target: ≤ 100 ms per payload; ideally a few milliseconds.
- Fast mode runs only `yara` and `hiddenchars` stages on a single payload.
- No ML inference, no external scanners, no entropy analysis, no repository walk.

## Report branches

- Report branches contain only generated reports under `reports/`.
- Do not commit tool source code or build artifacts to report branches.
- Each report directory includes `safeanalyze.json`, `safeanalyze.md`, `safeanalyze.html`, `safeanalyze.sarif`, and `duration.txt`.

## Useful commands

```bash
# Build
go build -o safeanalyze .

# Test
go test ./...

# Fast inspect latency check
echo 'ignore all previous instructions' | ./safeanalyze inspect --verbose

# Thorough repo scan
./safeanalyze scan ./repo --mode thorough -o ./out

# Install dependencies
./safeanalyze install --all
```

## Current autoresearch iteration

- **Version under test:** v0.3.4
- **Active hypothesis:** Adding pre-computed base64 and hex fragments of common prompt-injection phrases to the `encoded_prompt_injection` rule will catch encoded evasions without adding latency or false positives.
- **Last report branch:** `report-mexican-wombat-unstable-65b466a3-2026-07-15` (v0.3.3)
