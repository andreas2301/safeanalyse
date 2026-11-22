package icap

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Metrics tracks runtime counters and gauges for an ICAP Server in a
// thread-safe, zero-dependency manner and exposes them in Prometheus text
// exposition format alongside /healthz and /readyz endpoints.
type Metrics struct {
	reqAllow       atomic.Uint64
	reqBlock       atomic.Uint64
	reqShadowBlock atomic.Uint64
	reqSkip        atomic.Uint64
	reqError       atomic.Uint64
	reqOptions     atomic.Uint64

	activeConns   atomic.Int64
	rejectedConns atomic.Int64
	inflightScans atomic.Int64

	scanDurMicros atomic.Int64
	scanCount     atomic.Uint64

	mu       sync.Mutex
	ruleHits map[string]uint64
}

// NewMetrics creates a new Metrics collector.
func NewMetrics() *Metrics {
	return &Metrics{
		ruleHits: make(map[string]uint64),
	}
}

// RecordDecision increments the request counter for the given decision label.
func (m *Metrics) RecordDecision(decision string) {
	if m == nil {
		return
	}
	switch decision {
	case "allow":
		m.reqAllow.Add(1)
	case "block":
		m.reqBlock.Add(1)
	case "shadow_block":
		m.reqShadowBlock.Add(1)
	case "skip":
		m.reqSkip.Add(1)
	case "error":
		m.reqError.Add(1)
	case "options":
		m.reqOptions.Add(1)
	}
}

// DecisionCount returns the total number of requests recorded for decision.
func (m *Metrics) DecisionCount(decision string) uint64 {
	if m == nil {
		return 0
	}
	switch decision {
	case "allow":
		return m.reqAllow.Load()
	case "block":
		return m.reqBlock.Load()
	case "shadow_block":
		return m.reqShadowBlock.Load()
	case "skip":
		return m.reqSkip.Load()
	case "error":
		return m.reqError.Load()
	case "options":
		return m.reqOptions.Load()
	default:
		return 0
	}
}

// RecordScanDuration adds one completed scan's duration to the scan sum and count.
func (m *Metrics) RecordScanDuration(d time.Duration) {
	if m == nil {
		return
	}
	us := d.Microseconds()
	if us < 0 {
		us = 0
	}
	m.scanDurMicros.Add(us)
	m.scanCount.Add(1)
}

// RecordRuleHits increments the per-rule hit counter for each rule ID in rules.
func (m *Metrics) RecordRuleHits(rules []string) {
	if m == nil || len(rules) == 0 {
		return
	}
	m.mu.Lock()
	if m.ruleHits == nil {
		m.ruleHits = make(map[string]uint64, len(rules))
	}
	for _, r := range rules {
		if r != "" {
			m.ruleHits[r]++
		}
	}
	m.mu.Unlock()
}

// ActiveConns returns the current number of active TCP connections.
func (m *Metrics) ActiveConns() int64 {
	if m == nil {
		return 0
	}
	return m.activeConns.Load()
}

// RejectedConns returns the total number of rejected TCP connections.
func (m *Metrics) RejectedConns() int64 {
	if m == nil {
		return 0
	}
	return m.rejectedConns.Load()
}

// InflightScans returns the current number of scans in flight.
func (m *Metrics) InflightScans() int64 {
	if m == nil {
		return 0
	}
	return m.inflightScans.Load()
}

func (m *Metrics) incActiveConns() {
	if m != nil {
		m.activeConns.Add(1)
	}
}

func (m *Metrics) decActiveConns() {
	if m != nil {
		m.activeConns.Add(-1)
	}
}

func (m *Metrics) incRejectedConns() {
	if m != nil {
		m.rejectedConns.Add(1)
	}
}

func (m *Metrics) incInflightScans() {
	if m != nil {
		m.inflightScans.Add(1)
	}
}

func (m *Metrics) decInflightScans() {
	if m != nil {
		m.inflightScans.Add(-1)
	}
}

// Handler returns an http.Handler serving /healthz, /readyz, and /metrics.
func (m *Metrics) Handler(ready func() bool) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if ready == nil || ready() {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ready\n"))
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("not ready\n"))
	})
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(m.PrometheusText()))
	})
	return mux
}

// PrometheusText formats all metrics in Prometheus text exposition format.
func (m *Metrics) PrometheusText() string {
	var b strings.Builder

	b.WriteString("# HELP safeanalyze_icap_requests_total Total ICAP requests processed by decision.\n")
	b.WriteString("# TYPE safeanalyze_icap_requests_total counter\n")
	fmt.Fprintf(&b, "safeanalyze_icap_requests_total{decision=\"allow\"} %d\n", m.DecisionCount("allow"))
	fmt.Fprintf(&b, "safeanalyze_icap_requests_total{decision=\"block\"} %d\n", m.DecisionCount("block"))
	fmt.Fprintf(&b, "safeanalyze_icap_requests_total{decision=\"shadow_block\"} %d\n", m.DecisionCount("shadow_block"))
	fmt.Fprintf(&b, "safeanalyze_icap_requests_total{decision=\"skip\"} %d\n", m.DecisionCount("skip"))
	fmt.Fprintf(&b, "safeanalyze_icap_requests_total{decision=\"error\"} %d\n", m.DecisionCount("error"))
	fmt.Fprintf(&b, "safeanalyze_icap_requests_total{decision=\"options\"} %d\n", m.DecisionCount("options"))

	b.WriteString("# HELP safeanalyze_icap_active_connections Current number of active ICAP TCP connections.\n")
	b.WriteString("# TYPE safeanalyze_icap_active_connections gauge\n")
	fmt.Fprintf(&b, "safeanalyze_icap_active_connections %d\n", m.ActiveConns())

	b.WriteString("# HELP safeanalyze_icap_rejected_connections_total Total ICAP TCP connections rejected due to connection limit.\n")
	b.WriteString("# TYPE safeanalyze_icap_rejected_connections_total counter\n")
	fmt.Fprintf(&b, "safeanalyze_icap_rejected_connections_total %d\n", m.RejectedConns())

	b.WriteString("# HELP safeanalyze_icap_inflight_scans Current number of scans in flight.\n")
	b.WriteString("# TYPE safeanalyze_icap_inflight_scans gauge\n")
	fmt.Fprintf(&b, "safeanalyze_icap_inflight_scans %d\n", m.InflightScans())

	var durSumMs float64
	var durCount uint64
	if m != nil {
		durSumMs = float64(m.scanDurMicros.Load()) / 1000.0
		durCount = m.scanCount.Load()
	}
	b.WriteString("# HELP safeanalyze_icap_scan_duration_ms_sum Total scan duration in milliseconds.\n")
	b.WriteString("# TYPE safeanalyze_icap_scan_duration_ms_sum counter\n")
	fmt.Fprintf(&b, "safeanalyze_icap_scan_duration_ms_sum %.3f\n", durSumMs)

	b.WriteString("# HELP safeanalyze_icap_scan_duration_ms_count Total number of completed scans.\n")
	b.WriteString("# TYPE safeanalyze_icap_scan_duration_ms_count counter\n")
	fmt.Fprintf(&b, "safeanalyze_icap_scan_duration_ms_count %d\n", durCount)

	b.WriteString("# HELP safeanalyze_icap_rule_hits_total Total rule hits across scanned responses.\n")
	b.WriteString("# TYPE safeanalyze_icap_rule_hits_total counter\n")
	if m != nil {
		m.mu.Lock()
		rules := make([]string, 0, len(m.ruleHits))
		for k := range m.ruleHits {
			rules = append(rules, k)
		}
		sort.Strings(rules)
		counts := make([]uint64, len(rules))
		for i, k := range rules {
			counts[i] = m.ruleHits[k]
		}
		m.mu.Unlock()
		for i, k := range rules {
			fmt.Fprintf(&b, "safeanalyze_icap_rule_hits_total{rule=\"%s\"} %d\n", escapePromLabel(k), counts[i])
		}
	}

	return b.String()
}

var promLabelEscaper = strings.NewReplacer(`\`, `\\`, `"`, `\"`, "\n", `\n`)

func escapePromLabel(v string) string {
	return promLabelEscaper.Replace(v)
}
