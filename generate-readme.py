#!/usr/bin/env python3
"""Generate the report branch README from scan outputs."""
import json
import os
import pathlib
import sys

ROOT = pathlib.Path(__file__).parent
REPORT_DIR = ROOT / "reports" / "report-mutant-sarcasm-taco-7c259ba7-2026-07-15"

TARGETS = [
    ("repos/microsoft-bipia", "https://github.com/microsoft/BIPIA", "repo"),
    ("repos/uiuc-injecagent", "https://github.com/uiuc-kang-lab/InjecAgent", "repo"),
    ("repos/lakera-pint-benchmark", "https://github.com/lakeraai/pint-benchmark", "repo"),
    ("repos/alexh-prompt-injection-scanner", "https://github.com/alexh-scrt/prompt-injection-scanner", "repo"),
    ("repos/duriantaco-skylos", "https://github.com/duriantaco/skylos", "repo"),
    ("docs/promptfoo-scenarios", "https://www.promptfoo.dev/docs/configuration/scenarios/", "doc"),
    ("docs/promptfoo-webagents", "https://www.promptfoo.dev/blog/indirect-prompt-injection-web-agents/", "doc"),
]

def load_summary(dir_path):
    json_file = dir_path / "safeanalyze.json"
    dur_file = dir_path / "duration.txt"
    error_file = dir_path / "ERROR.txt"
    if error_file.exists():
        return None, "error", None
    if not json_file.exists():
        return None, "missing", None
    with open(json_file) as f:
        data = json.load(f)
    duration_ms = data.get("duration_ms", 0)
    duration_s = duration_ms / 1000.0
    findings = len(data.get("findings", []))
    version = data.get("metadata", {}).get("safeanalyze_version", "unknown")
    counts = data.get("summary", {}).get("counts_by_severity", {})
    return {
        "findings": findings,
        "duration_s": duration_s,
        "version": version,
        "counts": counts,
    }, "ok", None

def main():
    rows = []
    version = "unknown"
    for rel, source, kind in TARGETS:
        summary, status, _ = load_summary(REPORT_DIR / rel)
        if summary:
            version = summary["version"]
        rows.append((rel, source, kind, summary, status))

    lines = [
        "# safeanalyze Test Reports",
        "",
        f"**Branch:** `report-mutant-sarcasm-taco-7c259ba7-2026-07-15`",
        f"**Device:** mutant-sarcasm-taco-7c259ba7",
        f"**Date:** 2026-07-15",
        f"**Tool version:** v{version}",
        "",
        "This directory contains scan reports produced by `safeanalyze` against a collection of public prompt-injection test repositories, datasets, and documentation pages.",
        "",
        "## Scanned targets",
        "",
        "### Repositories",
        "",
        "| Target | Source | Duration | Findings | Status |",
        "|--------|--------|----------|----------|--------|",
    ]

    repo_rows = [r for r in rows if r[2] == "repo"]
    doc_rows = [r for r in rows if r[2] == "doc"]

    for rel, source, kind, summary, status in repo_rows:
        name = rel.split("/")[-1]
        if summary:
            dur = f"{summary['duration_s']:.2f} s"
            findings = str(summary["findings"])
            status = "scanned"
        else:
            dur = "-"
            findings = "-"
        lines.append(f"| {name} | {source} | {dur} | {findings} | {status} |")

    lines.extend([
        "",
        "### Datasets",
        "",
        "| Target | Source | Status |",
        "|--------|--------|--------|",
        "| MAlmasabi/Indirect-Prompt-Injection-BIPIA-GPT | https://huggingface.co/datasets/MAlmasabi/Indirect-Prompt-Injection-BIPIA-GPT | gated / inaccessible |",
        "| deepkeep/Prompt_Injection_v2 | https://huggingface.co/datasets/deepkeep/Prompt_Injection_v2 | gated / inaccessible |",
        "",
        "### Documentation",
        "",
        "| Target | Source | Duration | Findings | Status |",
        "|--------|--------|----------|----------|--------|",
    ])

    for rel, source, kind, summary, status in doc_rows:
        name = rel.split("/")[-1]
        if summary:
            dur = f"{summary['duration_s']:.2f} s"
            findings = str(summary["findings"])
            status = "scanned"
        else:
            dur = "-"
            findings = "-"
        lines.append(f"| {name} | {source} | {dur} | {findings} | {status} |")

    lines.extend([
        "",
        "## Output formats",
        "",
        "Each successful scan produces:",
        "",
        "- `safeanalyze.json` — full machine-readable report",
        "- `safeanalyze.md` — human-readable Markdown summary",
        "- `safeanalyze.html` — self-contained dashboard",
        "- `safeanalyze.sarif` — SARIF v2.1.0 for IDE/security-tool consumption",
        "- `duration.txt` — wall-clock scan duration in milliseconds",
        "",
        "## Configuration used",
        "",
        "The scans were run with `--mode thorough` and a test configuration that enables deterministic checks (YARA, entropy, hidden chars), disables the stochastic ONNX classifier, and enables the `alexh-scrt/prompt-injection-scanner`, Semgrep (`p/security-audit`), and TruffleHog external bridges.",
        "",
        "See `/tmp/safeanalyze-test-config-v031.yaml` on the scanning host for the exact config.",
        "",
        "## Notes",
        "",
        "- Two Hugging Face datasets could not be accessed because they are gated/private.",
        "- Report durations are recorded as `duration_ms` in `safeanalyze.json`.",
        "- Binary files (images, PDFs, archives, fonts, compiled artifacts) are skipped by the text-based checks as of v0.2.6.",
        "- Detected credential-like values are redacted in report output as of v0.2.7.",
        "- Indirect prompt-injection, tool-injection, delimiter-breakout, encoded-injection, markdown-injection, and instruction-to-include-code rules were added in v0.2.8/v0.2.9/v0.3.0.",
        "- Semgrep and TruffleHog were enabled by default in thorough mode in v0.3.1.",
    ])

    README = REPORT_DIR / "README.md"
    README.parent.mkdir(parents=True, exist_ok=True)
    README.write_text("\n".join(lines) + "\n")
    print(f"Wrote {README}")

if __name__ == "__main__":
    main()
