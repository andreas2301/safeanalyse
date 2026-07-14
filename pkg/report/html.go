package report

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"

	"github.com/user/safeanalyze/pkg/config"
)

// WriteHTML writes a self-contained HTML dashboard report to path.
func WriteHTML(report *Report, cfg config.OutputConfig, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	data := htmlReportData{
		Report:         report,
		Config:         cfg,
		SeverityOrder:  severityOrder(),
		SeverityCounts: report.Summary.CountsBySeverity,
		SourceCounts:   sortMap(report.Summary.CountsBySource),
		FindingsBySev:  groupBySeverityForHTML(report.Findings),
		AffectedFiles:  affectedFiles(report.Findings),
	}

	tmpl, err := template.New("report").Funcs(template.FuncMap{
		"severityClass": severityCSSClass,
		"add":           func(a, b int) int { return a + b },
	}).Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("parsing HTML template: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("creating HTML file: %w", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("executing HTML template: %w", err)
	}
	return nil
}

type htmlReportData struct {
	Report         *Report
	Config         config.OutputConfig
	SeverityOrder  []string
	SeverityCounts map[string]int
	SourceCounts   []keyValue
	FindingsBySev  map[string][]Finding
	AffectedFiles  []string
}

type keyValue struct {
	Key   string
	Value int
}

func sortMap(m map[string]int) []keyValue {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]keyValue, 0, len(keys))
	for _, k := range keys {
		out = append(out, keyValue{Key: k, Value: m[k]})
	}
	return out
}

func groupBySeverityForHTML(findings []Finding) map[string][]Finding {
	m := make(map[string][]Finding)
	for _, f := range findings {
		m[f.Severity] = append(m[f.Severity], f)
	}
	return m
}

func affectedFiles(findings []Finding) []string {
	seen := make(map[string]struct{})
	for _, f := range findings {
		seen[f.File] = struct{}{}
	}
	files := make([]string, 0, len(seen))
	for f := range seen {
		files = append(files, f)
	}
	sort.Strings(files)
	return files
}

func severityCSSClass(sev string) string {
	switch sev {
	case SeverityCritical:
		return "severity-critical"
	case SeverityHigh:
		return "severity-high"
	case SeverityMedium:
		return "severity-medium"
	case SeverityLow:
		return "severity-low"
	case SeverityInfo:
		return "severity-info"
	default:
		return "severity-unknown"
	}
}

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>safeanalyze Dashboard: {{.Report.Target}}</title>
  <style>
    :root {
      --bg: #f6f8fa;
      --fg: #1f2328;
      --card: #ffffff;
      --border: #d0d7de;
      --critical: #da3633;
      --high: #cf222e;
      --medium: #9a6700;
      --low: #0969da;
      --info: #656d76;
      --unknown: #6e7781;
    }
    * { box-sizing: border-box; }
    body {
      font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
      background: var(--bg);
      color: var(--fg);
      margin: 0;
      padding: 2rem;
      line-height: 1.5;
    }
    .container { max-width: 1200px; margin: 0 auto; }
    header { margin-bottom: 1.5rem; }
    header h1 { margin: 0 0 0.25rem; }
    header p { margin: 0; color: var(--info); }
    .note {
      background: #ddf4ff;
      border: 1px solid #84d8ff;
      border-radius: 6px;
      padding: 0.75rem 1rem;
      margin-bottom: 1.5rem;
      font-size: 0.9rem;
    }
    .cards {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
      gap: 1rem;
      margin-bottom: 1.5rem;
    }
    .card {
      background: var(--card);
      border: 1px solid var(--border);
      border-radius: 8px;
      padding: 1rem;
      text-align: center;
    }
    .card .value {
      font-size: 2rem;
      font-weight: 700;
      display: block;
    }
    .card .label {
      color: var(--info);
      font-size: 0.85rem;
    }
    section {
      background: var(--card);
      border: 1px solid var(--border);
      border-radius: 8px;
      padding: 1.25rem;
      margin-bottom: 1.5rem;
    }
    section h2 {
      margin-top: 0;
      font-size: 1.25rem;
    }
    table {
      width: 100%;
      border-collapse: collapse;
      margin-top: 0.75rem;
    }
    th, td {
      text-align: left;
      padding: 0.5rem;
      border-bottom: 1px solid var(--border);
    }
    th { cursor: pointer; background: #f6f8fa; }
    th:hover { background: #eff2f5; }
    tr:hover { background: #f6f8fa; }
    .badge {
      display: inline-block;
      padding: 0.15rem 0.5rem;
      border-radius: 999px;
      font-size: 0.75rem;
      font-weight: 600;
      text-transform: uppercase;
      color: #fff;
    }
    .severity-critical { background: var(--critical); }
    .severity-high { background: var(--high); }
    .severity-medium { background: var(--medium); }
    .severity-low { background: var(--low); }
    .severity-info { background: var(--info); }
    .severity-unknown { background: var(--unknown); }
    .filter { margin-bottom: 0.75rem; }
    .filter input {
      padding: 0.4rem 0.6rem;
      border: 1px solid var(--border);
      border-radius: 4px;
      width: 100%;
      max-width: 320px;
    }
    .empty { color: var(--info); font-style: italic; }
    .errors li { margin-bottom: 0.5rem; }
    pre {
      background: #f6f8fa;
      border: 1px solid var(--border);
      border-radius: 4px;
      padding: 0.75rem;
      overflow-x: auto;
    }
    .file-list { columns: 2; }
    @media (max-width: 640px) {
      body { padding: 1rem; }
      .file-list { columns: 1; }
    }
  </style>
</head>
<body>
  <div class="container">
    <header>
      <h1>safeanalyze Dashboard</h1>
      <p>Target: <strong>{{.Report.Target}}</strong> &middot; {{.Report.CompletedAt.Format "2006-01-02 15:04:05 MST"}}</p>
    </header>

    <div class="note">
      This is the dashboard view of the SARIF and Markdown reports. For machine-readable output, consume the SARIF file; for review-friendly output, use the Markdown report.
    </div>

    <section>
      <h2>Executive summary</h2>
      <div class="cards">
        <div class="card"><span class="value">{{len .Report.Findings}}</span><span class="label">Total findings</span></div>
        <div class="card"><span class="value">{{.Report.Summary.FilesScanned}}</span><span class="label">Files scanned</span></div>
        <div class="card"><span class="value">{{.Report.Summary.FilesSanitized}}</span><span class="label">Files sanitized</span></div>
        <div class="card"><span class="value">{{len .Report.Errors}}</span><span class="label">Errors</span></div>
      </div>
      {{if .SeverityCounts}}
      <h3>Findings by severity</h3>
      <table>
        <thead><tr><th>Severity</th><th>Count</th></tr></thead>
        <tbody>
          {{range .SeverityOrder}}{{$count := index $.SeverityCounts .}}{{if $count}}<tr><td><span class="badge {{severityClass .}}">{{.}}</span></td><td>{{$count}}</td></tr>{{end}}{{end}}
        </tbody>
      </table>
      {{end}}
      {{if .SourceCounts}}
      <h3>Findings by source</h3>
      <table>
        <thead><tr><th>Source</th><th>Count</th></tr></thead>
        <tbody>
          {{range .SourceCounts}}<tr><td>{{.Key}}</td><td>{{.Value}}</td></tr>{{end}}
        </tbody>
      </table>
      {{end}}
    </section>

    <section>
      <h2>Findings</h2>
      <div class="filter">
        <input type="text" id="findingFilter" placeholder="Filter findings by rule, file, message, or source...">
      </div>
      {{if .Report.Findings}}
      <table id="findingsTable">
        <thead>
          <tr>
            <th onclick="sortTable(0)">Severity</th>
            <th onclick="sortTable(1)">Rule</th>
            <th onclick="sortTable(2)">File</th>
            <th onclick="sortTable(3)">Line</th>
            <th onclick="sortTable(4)">Column</th>
            <th onclick="sortTable(5)">Message</th>
            <th onclick="sortTable(6)">Source</th>
            <th onclick="sortTable(7)">Confidence</th>
          </tr>
        </thead>
        <tbody>
          {{range .SeverityOrder}}{{$group := index $.FindingsBySev .}}{{range $group}}
          <tr>
            <td><span class="badge {{severityClass .Severity}}">{{.Severity}}</span></td>
            <td>{{.RuleID}}</td>
            <td>{{.File}}</td>
            <td>{{.Line}}</td>
            <td>{{.Column}}</td>
            <td>{{.Message}}</td>
            <td>{{.Source}}</td>
            <td>{{.Confidence}}</td>
          </tr>
          {{end}}{{end}}
        </tbody>
      </table>
      {{else}}
      <p class="empty">No findings detected.</p>
      {{end}}
    </section>

    <section>
      <h2>Errors</h2>
      {{if .Report.Errors}}
      <ul class="errors">
        {{range .Report.Errors}}<li><strong>[{{.Source}}]</strong> {{.Message}}</li>{{end}}
      </ul>
      {{else}}
      <p class="empty">No errors recorded.</p>
      {{end}}
    </section>

    {{if .Config.IncludeFileTree}}
    <section>
      <h2>Affected files</h2>
      {{if .AffectedFiles}}
      <ul class="file-list">
        {{range .AffectedFiles}}<li><code>{{.}}</code></li>{{end}}
      </ul>
      {{else}}
      <p class="empty">No affected files.</p>
      {{end}}
    </section>
    {{end}}
  </div>

  <script>
    document.getElementById('findingFilter')?.addEventListener('input', function(e) {
      const term = e.target.value.toLowerCase();
      const rows = document.querySelectorAll('#findingsTable tbody tr');
      rows.forEach(function(row) {
        const text = row.textContent.toLowerCase();
        row.style.display = text.includes(term) ? '' : 'none';
      });
    });

    function sortTable(n) {
      const table = document.getElementById('findingsTable');
      const tbody = table.querySelector('tbody');
      const rows = Array.from(tbody.querySelectorAll('tr'));
      const asc = !table.dataset['sortAsc' + n];
      rows.sort(function(a, b) {
        let x = a.cells[n].textContent.trim();
        let y = b.cells[n].textContent.trim();
        const xi = parseInt(x, 10), yi = parseInt(y, 10);
        if (!isNaN(xi) && !isNaN(yi)) {
          return asc ? xi - yi : yi - xi;
        }
        return asc ? x.localeCompare(y) : y.localeCompare(x);
      });
      rows.forEach(function(r) { tbody.appendChild(r); });
      table.dataset['sortAsc' + n] = asc ? '1' : '';
    }
  </script>
</body>
</html>
`
