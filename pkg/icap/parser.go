// Package icap implements the subset of the ICAP protocol (RFC 3507) needed
// to scan HTTP response bodies from Squid: OPTIONS and RESPMOD with
// Encapsulated header parsing, chunked bodies and Preview.
package icap

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/textproto"
	"strconv"
	"strings"
)

const (
	maxHeaderBytes   = 64 << 10 // ICAP header block and encapsulated HTTP headers
	maxChunkLineSize = 4096
)

// ErrMalformed marks a request that violates the protocol; the server answers
// it with 400 Bad Request and closes the connection.
var ErrMalformed = errors.New("malformed ICAP request")

func malformed(format string, args ...any) error {
	return fmt.Errorf("%w: %s", ErrMalformed, fmt.Sprintf(format, args...))
}

// EncapEntry is one "name=offset" element of the Encapsulated header.
type EncapEntry struct {
	Name   string
	Offset int
}

// Request is a parsed ICAP request head plus its encapsulated HTTP headers.
// The body (if any) is still unread on the connection.
type Request struct {
	Method       string
	URI          string
	Proto        string
	Header       textproto.MIMEHeader
	Encapsulated []EncapEntry
	ReqHdr       []byte // encapsulated HTTP request header block, if present
	ResHdr       []byte // encapsulated HTTP response header block, if present
	BodyType     string // res-body, req-body, opt-body, null-body or "" (no Encapsulated)
	Preview      int    // Preview size in bytes, -1 when the request is not a preview
}

// HasBody reports whether a chunked body follows the encapsulated headers.
func (r *Request) HasBody() bool {
	return r.BodyType == "res-body" || r.BodyType == "req-body" || r.BodyType == "opt-body"
}

// Allows204 reports whether the client sent "Allow: 204".
func (r *Request) Allows204() bool {
	for _, v := range r.Header.Values("Allow") {
		for _, tok := range strings.Split(v, ",") {
			if strings.TrimSpace(tok) == "204" {
				return true
			}
		}
	}
	return false
}

var validEncapNames = map[string]bool{
	"req-hdr": true, "res-hdr": true,
	"req-body": true, "res-body": true, "opt-body": true, "null-body": true,
}

// ParseEncapsulated parses an Encapsulated header value such as
// "req-hdr=0, res-hdr=137, res-body=296". Offsets must be non-decreasing,
// names must be known, and exactly one body entry must come last.
func ParseEncapsulated(v string) ([]EncapEntry, error) {
	v = strings.TrimSpace(v)
	if v == "" {
		return nil, malformed("empty Encapsulated header")
	}
	var out []EncapEntry
	prev := 0
	parts := strings.Split(v, ",")
	for i, part := range parts {
		name, off, ok := strings.Cut(strings.TrimSpace(part), "=")
		name = strings.ToLower(strings.TrimSpace(name))
		if !ok || !validEncapNames[name] {
			return nil, malformed("bad Encapsulated entry %q", part)
		}
		n, err := strconv.Atoi(strings.TrimSpace(off))
		if err != nil || n < 0 || n < prev {
			return nil, malformed("bad Encapsulated offset in %q", part)
		}
		isBody := strings.HasSuffix(name, "-body")
		last := i == len(parts)-1
		if isBody != last {
			return nil, malformed("Encapsulated body entry must be last: %q", v)
		}
		if i == 0 && n != 0 {
			return nil, malformed("first Encapsulated offset must be 0: %q", v)
		}
		prev = n
		out = append(out, EncapEntry{Name: name, Offset: n})
	}
	return out, nil
}

// ReadRequest reads an ICAP request line, the ICAP headers and the
// encapsulated HTTP header blocks. It returns io.EOF when the connection is
// closed cleanly before a new request starts.
func ReadRequest(br *bufio.Reader) (*Request, error) {
	tp := textproto.NewReader(br)
	line, err := tp.ReadLine()
	if err != nil {
		if errors.Is(err, io.EOF) && line == "" {
			return nil, io.EOF
		}
		return nil, err
	}
	if len(line) > maxChunkLineSize {
		return nil, malformed("request line too long")
	}
	parts := strings.Fields(line)
	if len(parts) != 3 || !strings.HasPrefix(parts[2], "ICAP/1.") {
		return nil, malformed("bad request line %q", truncate(line, 80))
	}
	hdr, err := readMIMEHeader(tp)
	if err != nil {
		return nil, err
	}
	req := &Request{Method: parts[0], URI: parts[1], Proto: parts[2], Header: hdr, Preview: -1}

	if p := hdr.Get("Preview"); p != "" {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil || n < 0 {
			return nil, malformed("bad Preview header %q", p)
		}
		req.Preview = n
	}

	enc := hdr.Get("Encapsulated")
	if enc == "" {
		if req.Method == "RESPMOD" || req.Method == "REQMOD" {
			return nil, malformed("%s without Encapsulated header", req.Method)
		}
		return req, nil
	}
	entries, err := ParseEncapsulated(enc)
	if err != nil {
		return nil, err
	}
	req.Encapsulated = entries
	body := entries[len(entries)-1]
	req.BodyType = body.Name
	if body.Offset > maxHeaderBytes {
		return nil, malformed("encapsulated headers too large (%d bytes)", body.Offset)
	}
	block := make([]byte, body.Offset)
	if _, err := io.ReadFull(br, block); err != nil {
		return nil, fmt.Errorf("reading encapsulated headers: %w", err)
	}
	for i := 0; i < len(entries)-1; i++ {
		sec := block[entries[i].Offset:entries[i+1].Offset]
		switch entries[i].Name {
		case "req-hdr":
			req.ReqHdr = sec
		case "res-hdr":
			req.ResHdr = sec
		}
	}
	return req, nil
}

func readMIMEHeader(tp *textproto.Reader) (textproto.MIMEHeader, error) {
	hdr, err := tp.ReadMIMEHeader()
	if err != nil {
		var pe textproto.ProtocolError
		if errors.As(err, &pe) {
			return nil, malformed("bad header: %v", err)
		}
		return nil, err
	}
	return hdr, nil
}

// HTTPHeader is a parsed encapsulated HTTP header block.
type HTTPHeader struct {
	StartLine string
	Header    textproto.MIMEHeader
}

// ParseHTTPHeader parses an encapsulated HTTP request or response header block.
func ParseHTTPHeader(b []byte) (*HTTPHeader, error) {
	if len(b) == 0 {
		return &HTTPHeader{Header: textproto.MIMEHeader{}}, nil
	}
	tp := textproto.NewReader(bufio.NewReader(bytes.NewReader(b)))
	line, err := tp.ReadLine()
	if err != nil {
		return nil, malformed("bad encapsulated HTTP start line: %v", err)
	}
	hdr, err := tp.ReadMIMEHeader()
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, malformed("bad encapsulated HTTP header: %v", err)
	}
	if hdr == nil {
		hdr = textproto.MIMEHeader{}
	}
	return &HTTPHeader{StartLine: line, Header: hdr}, nil
}

// chunkSink collects decoded chunk data. Bytes beyond keep are counted but
// dropped unless keepAll is set. With keepAll and onFull set, onFull is called
// once (with the first keep bytes) when the data outgrows keep; afterwards the
// data goes to passthrough if onFull set it, and is dropped otherwise.
type chunkSink struct {
	buf         bytes.Buffer
	keep        int
	keepAll     bool
	total       int64
	onFull      func(buffered []byte)
	passthrough func(p []byte)
	fired       bool
}

func (s *chunkSink) write(p []byte) {
	s.total += int64(len(p))
	if s.passthrough != nil {
		s.passthrough(p)
		return
	}
	if s.keepAll && s.onFull != nil && !s.fired && s.buf.Len()+len(p) > s.keep {
		s.fired = true
		if room := s.keep - s.buf.Len(); room > 0 {
			s.buf.Write(p[:room])
			p = p[room:]
		}
		s.keepAll = false
		s.onFull(s.buf.Bytes())
		if s.passthrough != nil {
			s.passthrough(p)
		}
		return
	}
	if s.keepAll {
		s.buf.Write(p)
		return
	}
	if room := s.keep - s.buf.Len(); room > 0 {
		if len(p) > room {
			p = p[:room]
		}
		s.buf.Write(p)
	}
}

// readChunked reads an ICAP chunked body up to and including the terminating
// zero-length chunk. It returns ieof=true when the last chunk carried the
// "ieof" extension (end of a preview that already holds the whole body).
func readChunked(br *bufio.Reader, sink *chunkSink) (ieof bool, err error) {
	for {
		line, err := readLine(br)
		if err != nil {
			return false, err
		}
		sizeStr, ext, _ := strings.Cut(line, ";")
		sizeStr = strings.TrimSpace(sizeStr)
		size, perr := strconv.ParseInt(sizeStr, 16, 64)
		if perr != nil || size < 0 || sizeStr == "" {
			return false, malformed("bad chunk size %q", truncate(line, 40))
		}
		if size == 0 {
			ieof = strings.TrimSpace(strings.ToLower(ext)) == "ieof"
			// Trailer section: lines until an empty one.
			for {
				t, err := readLine(br)
				if err != nil {
					return false, err
				}
				if t == "" {
					return ieof, nil
				}
			}
		}
		if err := copyChunk(br, size, sink); err != nil {
			return false, err
		}
		crlf, err := readLine(br)
		if err != nil {
			return false, err
		}
		if crlf != "" {
			return false, malformed("missing CRLF after chunk data")
		}
	}
}

func copyChunk(br *bufio.Reader, size int64, sink *chunkSink) error {
	var tmp [32 << 10]byte
	for size > 0 {
		n := int64(len(tmp))
		if size < n {
			n = size
		}
		if _, err := io.ReadFull(br, tmp[:n]); err != nil {
			return fmt.Errorf("reading chunk data: %w", err)
		}
		sink.write(tmp[:n])
		size -= n
	}
	return nil
}

// readLine reads one CRLF- (or LF-) terminated line with a length limit.
func readLine(br *bufio.Reader) (string, error) {
	var b []byte
	for {
		frag, isPrefix, err := br.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return "", io.ErrUnexpectedEOF
			}
			return "", err
		}
		b = append(b, frag...)
		if len(b) > maxChunkLineSize {
			return "", malformed("line too long")
		}
		if !isPrefix {
			return string(b), nil
		}
	}
}

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n] + "..."
	}
	return s
}
