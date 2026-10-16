package icap

import (
	"bufio"
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime"
	"net"
	"net/url"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/user/safeanalyze/pkg/report"
)

// ScanFunc runs the detection suite on one payload and returns its findings.
type ScanFunc func(ctx context.Context, payload []byte) ([]report.Finding, error)

// Config configures the ICAP server.
type Config struct {
	Service       string        // service path, e.g. "/safeanalyze"
	MaxBodyBytes  int           // scan at most this many (decoded) body bytes
	Timeout       time.Duration // scan time budget per request
	IdleTimeout   time.Duration // max inactivity while reading from the client
	FailOpen      bool          // allow the response on errors/timeouts
	BlockSeverity string        // minimum finding severity that blocks
	ISTag         string        // service tag, without quotes
	PreviewBytes  int           // Preview size advertised in OPTIONS
	// MaxConcurrentScans bounds scans in flight, including scans abandoned
	// after a timeout that are still running. Default runtime.NumCPU()*2.
	MaxConcurrentScans int
	Scan               ScanFunc
	Logger             *slog.Logger
}

// Server is an ICAP RESPMOD server.
type Server struct {
	cfg       Config
	blockRank int
	sem       chan struct{} // scan slots

	mu     sync.Mutex
	ln     net.Listener
	conns  map[net.Conn]struct{}
	closed bool
	wg     sync.WaitGroup
}

// NewServer validates cfg and returns a server.
func NewServer(cfg Config) (*Server, error) {
	if cfg.Scan == nil {
		return nil, errors.New("icap: Scan function is required")
	}
	if cfg.Service == "" {
		cfg.Service = "/safeanalyze"
	}
	if !strings.HasPrefix(cfg.Service, "/") {
		cfg.Service = "/" + cfg.Service
	}
	if cfg.MaxBodyBytes <= 0 {
		cfg.MaxBodyBytes = 2 << 20
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 2 * time.Second
	}
	if cfg.IdleTimeout <= 0 {
		cfg.IdleTimeout = 60 * time.Second
	}
	if cfg.BlockSeverity == "" {
		cfg.BlockSeverity = report.SeverityHigh
	}
	rank := report.SeverityRank(cfg.BlockSeverity)
	if rank == 0 {
		return nil, fmt.Errorf("icap: invalid block severity %q", cfg.BlockSeverity)
	}
	if cfg.ISTag == "" {
		cfg.ISTag = "safeanalyze"
	}
	if cfg.PreviewBytes <= 0 {
		cfg.PreviewBytes = 4096
	}
	if cfg.MaxConcurrentScans <= 0 {
		cfg.MaxConcurrentScans = runtime.NumCPU() * 2
	}
	if cfg.Logger == nil {
		cfg.Logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	return &Server{cfg: cfg, blockRank: rank, sem: make(chan struct{}, cfg.MaxConcurrentScans), conns: map[net.Conn]struct{}{}}, nil
}

// Serve accepts connections on ln until Close is called.
func (s *Server) Serve(ln net.Listener) error {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return net.ErrClosed
	}
	s.ln = ln
	s.mu.Unlock()
	for {
		conn, err := ln.Accept()
		if err != nil {
			s.mu.Lock()
			closed := s.closed
			s.mu.Unlock()
			if closed {
				return nil
			}
			var ne net.Error
			if errors.As(err, &ne) && ne.Timeout() {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			return err
		}
		s.mu.Lock()
		if s.closed {
			s.mu.Unlock()
			conn.Close()
			return nil
		}
		s.conns[conn] = struct{}{}
		s.wg.Add(1)
		s.mu.Unlock()
		go func() {
			defer s.wg.Done()
			s.serveConn(conn)
			s.mu.Lock()
			delete(s.conns, conn)
			s.mu.Unlock()
		}()
	}
}

// Close stops accepting, closes open connections and waits for handlers.
func (s *Server) Close() error {
	s.mu.Lock()
	s.closed = true
	var err error
	if s.ln != nil {
		err = s.ln.Close()
	}
	for c := range s.conns {
		c.Close()
	}
	s.mu.Unlock()
	s.wg.Wait()
	return err
}

// idleConnReader refreshes the read deadline before every read so a slow but
// active client (Squid streaming a large origin response) is not cut off,
// while a stalled one is.
type idleConnReader struct {
	conn net.Conn
	idle time.Duration
}

func (r idleConnReader) Read(p []byte) (int, error) {
	_ = r.conn.SetReadDeadline(time.Now().Add(r.idle))
	return r.conn.Read(p)
}

func (s *Server) serveConn(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReaderSize(idleConnReader{conn: conn, idle: s.cfg.IdleTimeout}, 64<<10)
	bw := bufio.NewWriterSize(conn, 32<<10)
	for {
		req, err := ReadRequest(br)
		if err != nil {
			if errors.Is(err, ErrMalformed) {
				s.cfg.Logger.Warn("icap request", "decision", "bad_request", "error", err.Error(), "remote", conn.RemoteAddr().String())
				s.writeStatus(conn, bw, 400, "Bad Request", true)
			}
			return
		}
		keep := s.handle(conn, br, bw, req)
		if !keep || strings.EqualFold(req.Header.Get("Connection"), "close") {
			return
		}
	}
}

// result describes the outcome of one request for logging.
type result struct {
	decision  string // allow, block, skip, error
	reason    string
	rules     []string
	bytes     int64
	scanned   int
	truncated bool
	status    int
}

func (s *Server) handle(conn net.Conn, br *bufio.Reader, bw *bufio.Writer, req *Request) (keepAlive bool) {
	start := time.Now()
	res := &result{}
	defer func() {
		attrs := []any{
			"method", req.Method,
			"status", res.status,
			"decision", res.decision,
			"rules", res.rules,
			"bytes", res.bytes,
			"scanned_bytes", res.scanned,
			"truncated", res.truncated,
			"ms", float64(time.Since(start).Microseconds()) / 1000,
		}
		if res.reason != "" {
			attrs = append(attrs, "reason", res.reason)
		}
		if u := requestURL(req); u != "" {
			attrs = append(attrs, "url", u)
		}
		s.cfg.Logger.Info("icap request", attrs...)
	}()

	u, err := url.Parse(req.URI)
	if err != nil || strings.TrimRight(u.Path, "/") != strings.TrimRight(s.cfg.Service, "/") {
		res.decision, res.reason, res.status = "error", "unknown_service", 404
		// The body (if any) was not read, so the connection cannot be reused.
		s.writeStatus(conn, bw, 404, "ICAP Service Not Found", true)
		return false
	}

	switch req.Method {
	case "OPTIONS":
		res.decision, res.status = "options", 200
		s.writeOptions(conn, bw)
		return true
	case "RESPMOD":
		return s.respmod(conn, br, bw, req, res)
	case "REQMOD":
		res.decision, res.reason, res.status = "error", "method_not_allowed", 405
		s.writeStatus(conn, bw, 405, "Method Not Allowed", true)
		return false
	default:
		res.decision, res.reason, res.status = "error", "not_implemented", 501
		s.writeStatus(conn, bw, 501, "Method Not Implemented", true)
		return false
	}
}

func (s *Server) respmod(conn net.Conn, br *bufio.Reader, bw *bufio.Writer, req *Request, res *result) bool {
	allow204 := req.Allows204()
	resHdr, err := ParseHTTPHeader(req.ResHdr)
	if err != nil {
		res.decision, res.reason, res.status = "error", "bad_http_header", 400
		s.writeStatus(conn, bw, 400, "Bad Request", true)
		return false
	}

	if !req.HasBody() {
		res.decision, res.reason = "allow", "no_body"
		res.status = s.writeAllow(conn, bw, req, allow204, true, nil)
		return true
	}

	scannable := isTextContentType(resHdr.Header.Get("Content-Type"))
	encoding := resHdr.Header.Get("Content-Encoding")
	inPreview := req.Preview >= 0
	// Echoing the response back (200 instead of 204) needs the body, which is
	// buffered up to MaxBodyBytes; see echoOverflow for larger bodies.
	sink := &chunkSink{keep: s.cfg.MaxBodyBytes, keepAll: !allow204}
	echo := &echoState{}
	armEcho := func() {
		if !req.Allows204() {
			sink.onFull = func(buffered []byte) { s.echoOverflow(conn, bw, req, res, sink, echo, buffered, scannable, encoding) }
		}
	}
	if !inPreview {
		armEcho()
	}

	readErr := func(err error) bool {
		if echo.streaming {
			// A partial 200 was already sent; the stream cannot be recovered.
			res.bytes, res.decision, res.reason = sink.total, "error", "read: "+err.Error()
			return false
		}
		return s.bodyError(conn, bw, req, res, sink, err)
	}

	ieof, err := readChunked(br, sink)
	if err != nil {
		return readErr(err)
	}
	if inPreview && !ieof {
		if !scannable {
			// 204 is always permitted in response to a preview.
			res.decision, res.reason, res.bytes = "skip", "content_type", sink.total
			res.status = s.writeAllow(conn, bw, req, true, true, nil)
			return true
		}
		if _, err := bw.WriteString("ICAP/1.0 100 Continue\r\n\r\n"); err == nil {
			_ = conn.SetWriteDeadline(time.Now().Add(s.cfg.IdleTimeout))
			err = bw.Flush()
		}
		if err != nil {
			res.decision, res.reason, res.status = "error", "write_failed", 0
			return false
		}
		armEcho()
		if _, err := readChunked(br, sink); err != nil {
			return readErr(err)
		}
	} else if inPreview {
		// The whole body fit in the preview; 204 stays permitted.
		allow204 = true
	}
	res.bytes = sink.total

	if sink.fired {
		return s.finishEchoOverflow(conn, bw, res, echo)
	}
	raw := sink.buf.Bytes()

	if !scannable {
		res.decision, res.reason = "skip", "content_type"
		res.status = s.writeAllow(conn, bw, req, allow204, true, raw)
		return true
	}

	res.truncated = sink.total > int64(s.cfg.MaxBodyBytes)
	blocking, err := s.scanBody(raw, encoding, res)
	if err != nil {
		return s.failDecision(conn, bw, req, res, allow204, raw)
	}
	if len(blocking) > 0 {
		res.decision, res.rules = "block", blocking
		res.status = s.writeBlock(conn, bw, blocking, "prompt-injection content detected")
		return true
	}
	res.decision = "allow"
	res.status = s.writeAllow(conn, bw, req, allow204, true, raw)
	return true
}

// scanBody decodes and scans at most MaxBodyBytes of raw and returns the rule
// IDs at or above the block severity. A non-nil error means the body could not
// be scanned; res.reason says why and the caller applies the fail policy.
func (s *Server) scanBody(raw []byte, encoding string, res *result) ([]string, error) {
	if len(raw) > s.cfg.MaxBodyBytes {
		raw = raw[:s.cfg.MaxBodyBytes]
	}
	decoded, decTrunc, err := decodeBody(raw, encoding, s.cfg.MaxBodyBytes)
	if err != nil {
		res.reason = "decode: " + err.Error()
		return nil, err
	}
	res.truncated = res.truncated || decTrunc
	res.scanned = len(decoded)

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.Timeout)
	defer cancel()
	findings, err := s.scan(ctx, decoded)
	if err != nil {
		switch {
		case errors.Is(err, errScanBusy):
			res.reason = "scan_queue_timeout"
		case errors.Is(err, context.DeadlineExceeded):
			res.reason = "scan_timeout"
		default:
			res.reason = "scan: " + err.Error()
		}
		return nil, err
	}
	res.rules = ruleIDs(findings, 0)
	return ruleIDs(findings, s.blockRank), nil
}

// errScanBusy reports that no scan slot became free within the scan timeout.
var errScanBusy = errors.New("all scan slots busy")

// scan runs the scanner with the context deadline. It first waits for one of
// MaxConcurrentScans slots (within the same deadline). A scanner that ignores
// the context keeps running in the background and holds its slot until it
// returns, so abandoned scans never exceed the bound; its result is discarded.
func (s *Server) scan(ctx context.Context, payload []byte) ([]report.Finding, error) {
	select {
	case s.sem <- struct{}{}:
	case <-ctx.Done():
		return nil, fmt.Errorf("%w: %w", errScanBusy, ctx.Err())
	}
	type out struct {
		f   []report.Finding
		err error
	}
	ch := make(chan out, 1)
	go func() {
		defer func() { <-s.sem }()
		f, err := s.cfg.Scan(ctx, payload)
		ch <- out{f, err}
	}()
	select {
	case o := <-ch:
		return o.f, o.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// echoState tracks a no-Allow-204 response whose body exceeded MaxBodyBytes.
type echoState struct {
	streaming bool     // a 200 echo is being streamed to the client
	writeErr  error    // write failure while streaming
	rules     []string // blocking rules found in the first MaxBodyBytes
}

// echoOverflow runs once when a body that must be echoed (client sent no
// "Allow: 204") grows past MaxBodyBytes. The first MaxBodyBytes are buffered.
//   - fail-closed, text: the rest is drained and the response replaced by a 403.
//   - fail-open, or non-text: the buffered prefix is scanned (text only, like
//     the 204 path); on a finding the rest is drained and a 403 sent, otherwise
//     a 200 echo starts streaming immediately and the remaining chunks are
//     forwarded unscanned as they arrive, without buffering.
func (s *Server) echoOverflow(conn net.Conn, bw *bufio.Writer, req *Request, res *result, sink *chunkSink, echo *echoState, buffered []byte, scannable bool, encoding string) {
	res.truncated = true
	if scannable {
		if !s.cfg.FailOpen {
			res.reason = "echo_body_over_cap"
			return
		}
		blocking, err := s.scanBody(buffered, encoding, res)
		switch {
		case err != nil:
			res.decision = "error" // reason set by scanBody; fail-open
		case len(blocking) > 0:
			echo.rules = blocking
			return
		default:
			res.decision, res.reason = "allow", "echo_body_over_cap"
		}
	} else {
		res.decision, res.reason = "skip", "content_type"
	}

	echo.streaming = true
	var b bytes.Buffer
	b.WriteString("ICAP/1.0 200 OK\r\n")
	s.writeCommon(&b, false)
	fmt.Fprintf(&b, "Encapsulated: res-hdr=0, res-body=%d\r\n\r\n", len(req.ResHdr))
	b.Write(req.ResHdr)
	writeChunk(&b, buffered)
	s.streamWrite(conn, bw, echo, b.Bytes())
	sink.passthrough = func(p []byte) {
		if len(p) == 0 {
			return
		}
		var c bytes.Buffer
		writeChunk(&c, p)
		s.streamWrite(conn, bw, echo, c.Bytes())
	}
}

func (s *Server) streamWrite(conn net.Conn, bw *bufio.Writer, echo *echoState, p []byte) {
	if echo.writeErr != nil {
		return
	}
	_ = conn.SetWriteDeadline(time.Now().Add(s.cfg.IdleTimeout))
	_, echo.writeErr = bw.Write(p)
}

// finishEchoOverflow completes a response after echoOverflow ran.
func (s *Server) finishEchoOverflow(conn net.Conn, bw *bufio.Writer, res *result, echo *echoState) bool {
	switch {
	case echo.streaming:
		res.status = 200
		if echo.writeErr == nil {
			_ = conn.SetWriteDeadline(time.Now().Add(s.cfg.IdleTimeout))
			if _, err := bw.WriteString("0\r\n\r\n"); err == nil {
				echo.writeErr = bw.Flush()
			} else {
				echo.writeErr = err
			}
		}
		return echo.writeErr == nil
	case len(echo.rules) > 0:
		res.decision, res.rules = "block", echo.rules
		res.status = s.writeBlock(conn, bw, echo.rules, "prompt-injection content detected")
	default: // fail-closed
		res.decision = "block"
		res.status = s.writeBlock(conn, bw, nil, "response too large to scan")
	}
	return true
}

// bodyError handles a failure while reading the chunked body. The stream is
// out of sync afterwards, so the connection is always closed.
func (s *Server) bodyError(conn net.Conn, bw *bufio.Writer, req *Request, res *result, sink *chunkSink, err error) bool {
	res.bytes = sink.total
	if errors.Is(err, ErrMalformed) {
		res.decision, res.reason, res.status = "error", err.Error(), 400
		s.writeStatus(conn, bw, 400, "Bad Request", true)
		return false
	}
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		res.reason = "read_timeout"
	} else {
		res.reason = "read: " + err.Error()
	}
	if s.cfg.FailOpen {
		res.decision = "error"
		if req.Allows204() || req.Preview >= 0 {
			res.status = s.writeAllow(conn, bw, req, true, false, nil)
		} else {
			// The body is incomplete and cannot be echoed back.
			res.status = 500
			s.writeStatus(conn, bw, 500, "Server Error", true)
		}
		return false
	}
	res.decision = "block"
	res.status = s.writeBlockClose(conn, bw, nil, "response could not be scanned", true)
	return false
}

// failDecision applies the fail-open / fail-closed policy after the body was
// read completely.
func (s *Server) failDecision(conn net.Conn, bw *bufio.Writer, req *Request, res *result, allow204 bool, raw []byte) bool {
	if s.cfg.FailOpen {
		res.decision = "error"
		res.status = s.writeAllow(conn, bw, req, allow204, true, raw)
		return true
	}
	res.decision = "block"
	res.status = s.writeBlock(conn, bw, nil, "response could not be scanned")
	return true
}

// writeAllow tells the client to use the original response: 204 when allowed,
// otherwise a 200 echoing the original headers and body.
func (s *Server) writeAllow(conn net.Conn, bw *bufio.Writer, req *Request, allow204, keep bool, raw []byte) int {
	if allow204 {
		s.writeStatus(conn, bw, 204, "No Content", !keep)
		return 204
	}
	var b bytes.Buffer
	b.WriteString("ICAP/1.0 200 OK\r\n")
	s.writeCommon(&b, !keep)
	switch {
	case len(req.ResHdr) == 0:
		b.WriteString("Encapsulated: null-body=0\r\n\r\n")
	case !req.HasBody():
		fmt.Fprintf(&b, "Encapsulated: res-hdr=0, null-body=%d\r\n\r\n", len(req.ResHdr))
		b.Write(req.ResHdr)
	default:
		fmt.Fprintf(&b, "Encapsulated: res-hdr=0, res-body=%d\r\n\r\n", len(req.ResHdr))
		b.Write(req.ResHdr)
		writeChunked(&b, raw)
	}
	s.flush(conn, bw, b.Bytes())
	return 200
}

func (s *Server) writeBlock(conn net.Conn, bw *bufio.Writer, rules []string, why string) int {
	return s.writeBlockClose(conn, bw, rules, why, false)
}

// writeBlockClose replaces the response with a small HTTP 403. The payload is
// never echoed; only rule IDs are named.
func (s *Server) writeBlockClose(conn net.Conn, bw *bufio.Writer, rules []string, why string, closeConn bool) int {
	body := "Blocked by safeanalyze: " + why + ".\n"
	if len(rules) > 0 {
		body += "Rules: " + strings.Join(rules, ", ") + "\n"
	}
	var h bytes.Buffer
	h.WriteString("HTTP/1.1 403 Forbidden\r\n")
	h.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	fmt.Fprintf(&h, "Content-Length: %d\r\n", len(body))
	h.WriteString("Cache-Control: no-store\r\n")
	if len(rules) > 0 {
		h.WriteString("X-Safeanalyze-Rules: " + strings.Join(rules, ",") + "\r\n")
	}
	h.WriteString("\r\n")

	var b bytes.Buffer
	b.WriteString("ICAP/1.0 200 OK\r\n")
	s.writeCommon(&b, closeConn)
	fmt.Fprintf(&b, "Encapsulated: res-hdr=0, res-body=%d\r\n\r\n", h.Len())
	b.Write(h.Bytes())
	writeChunked(&b, []byte(body))
	s.flush(conn, bw, b.Bytes())
	return 403
}

func (s *Server) writeOptions(conn net.Conn, bw *bufio.Writer) {
	var b bytes.Buffer
	b.WriteString("ICAP/1.0 200 OK\r\n")
	s.writeCommon(&b, false)
	b.WriteString("Methods: RESPMOD\r\n")
	b.WriteString("Service: safeanalyze prompt-injection scanner\r\n")
	b.WriteString("Allow: 204\r\n")
	fmt.Fprintf(&b, "Preview: %d\r\n", s.cfg.PreviewBytes)
	b.WriteString("Transfer-Preview: *\r\n")
	b.WriteString("Options-TTL: 3600\r\n")
	b.WriteString("Encapsulated: null-body=0\r\n\r\n")
	s.flush(conn, bw, b.Bytes())
}

func (s *Server) writeStatus(conn net.Conn, bw *bufio.Writer, code int, text string, closeConn bool) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "ICAP/1.0 %d %s\r\n", code, text)
	s.writeCommon(&b, closeConn)
	b.WriteString("Encapsulated: null-body=0\r\n\r\n")
	s.flush(conn, bw, b.Bytes())
}

func (s *Server) writeCommon(b *bytes.Buffer, closeConn bool) {
	fmt.Fprintf(b, "ISTag: \"%s\"\r\n", s.cfg.ISTag)
	b.WriteString("Date: " + time.Now().UTC().Format(time.RFC1123) + "\r\n")
	b.WriteString("Server: safeanalyze\r\n")
	if closeConn {
		b.WriteString("Connection: close\r\n")
	}
}

func (s *Server) flush(conn net.Conn, bw *bufio.Writer, p []byte) {
	_ = conn.SetWriteDeadline(time.Now().Add(s.cfg.IdleTimeout))
	if _, err := bw.Write(p); err == nil {
		_ = bw.Flush()
	}
}

// writeChunked writes data as one chunk plus the terminating zero chunk.
func writeChunked(b *bytes.Buffer, data []byte) {
	writeChunk(b, data)
	b.WriteString("0\r\n\r\n")
}

// writeChunk writes data as one chunk; empty data writes nothing.
func writeChunk(b *bytes.Buffer, data []byte) {
	if len(data) > 0 {
		fmt.Fprintf(b, "%x\r\n", len(data))
		b.Write(data)
		b.WriteString("\r\n")
	}
}

// ruleIDs returns the sorted, de-duplicated rule IDs of findings at or above
// minRank (0 = all).
func ruleIDs(findings []report.Finding, minRank int) []string {
	seen := map[string]bool{}
	var out []string
	for _, f := range findings {
		if minRank > 0 && report.SeverityRank(f.Severity) < minRank {
			continue
		}
		id := sanitizeToken(f.RuleID)
		if id == "" || seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}

func sanitizeToken(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || strings.ContainsRune("_-./:", r) {
			return r
		}
		return -1
	}, s)
}

// isTextContentType reports whether a response with this Content-Type should
// be scanned. A missing type is scanned.
func isTextContentType(ct string) bool {
	if strings.TrimSpace(ct) == "" {
		return true
	}
	mt, _, err := mime.ParseMediaType(ct)
	if err != nil {
		mt = strings.ToLower(strings.TrimSpace(strings.SplitN(ct, ";", 2)[0]))
	}
	if strings.HasPrefix(mt, "text/") || strings.HasSuffix(mt, "+json") || strings.HasSuffix(mt, "+xml") {
		return true
	}
	switch mt {
	case "application/json", "application/xml", "application/javascript",
		"application/ecmascript", "application/x-javascript", "application/x-ndjson",
		"application/x-www-form-urlencoded", "application/yaml", "application/x-yaml",
		"application/graphql", "application/rss+xml", "application/atom+xml":
		return true
	}
	return false
}

// decodeBody undoes Content-Encoding for scanning. Decoded output is capped at
// limit bytes; a truncated (capped) compressed stream yields its decodable
// prefix.
func decodeBody(raw []byte, encoding string, limit int) (out []byte, truncated bool, err error) {
	enc := strings.ToLower(strings.TrimSpace(encoding))
	if enc == "" || enc == "identity" {
		return raw, false, nil
	}
	if strings.Contains(enc, ",") {
		return nil, false, fmt.Errorf("unsupported stacked encoding %q", encoding)
	}
	var r io.Reader
	switch enc {
	case "gzip", "x-gzip":
		zr, err := gzip.NewReader(bytes.NewReader(raw))
		if err != nil {
			return nil, false, err
		}
		r = zr
	case "deflate":
		// HTTP "deflate" is zlib-wrapped; some servers send raw deflate.
		if zr, err := zlib.NewReader(bytes.NewReader(raw)); err == nil {
			r = zr
		} else {
			r = flate.NewReader(bytes.NewReader(raw))
		}
	default:
		return nil, false, fmt.Errorf("unsupported content encoding %q", encoding)
	}
	out, err = io.ReadAll(io.LimitReader(r, int64(limit)+1))
	if len(out) > limit {
		return out[:limit], true, nil
	}
	if err != nil {
		if (errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF)) && len(out) > 0 {
			return out, true, nil
		}
		return nil, false, err
	}
	return out, false, nil
}

// requestURL extracts the request target from the encapsulated HTTP request
// header, for logging.
func requestURL(req *Request) string {
	if len(req.ReqHdr) == 0 {
		return ""
	}
	line, _, _ := strings.Cut(string(req.ReqHdr), "\n")
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return ""
	}
	return truncate(parts[1], 256)
}
