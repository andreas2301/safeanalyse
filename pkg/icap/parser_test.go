package icap

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
	"testing"
)

func TestParseEncapsulated(t *testing.T) {
	got, err := ParseEncapsulated("req-hdr=0, res-hdr=137, res-body=296")
	if err != nil {
		t.Fatal(err)
	}
	want := []EncapEntry{{"req-hdr", 0}, {"res-hdr", 137}, {"res-body", 296}}
	if len(got) != len(want) {
		t.Fatalf("got %v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("entry %d: got %v want %v", i, got[i], want[i])
		}
	}
	if got, err := ParseEncapsulated("null-body=0"); err != nil || got[0].Name != "null-body" {
		t.Fatalf("null-body: %v %v", got, err)
	}

	for _, bad := range []string{
		"", "res-hdr=0", "res-hdr=10, res-body=20", "res-hdr=0, res-body=x",
		"res-hdr=0, res-body=-1", "req-hdr=0, res-hdr=50, res-body=40",
		"res-body=0, res-hdr=10", "foo=0, res-body=1", "res-hdr0, res-body=1",
	} {
		if _, err := ParseEncapsulated(bad); !errors.Is(err, ErrMalformed) {
			t.Errorf("%q: expected ErrMalformed, got %v", bad, err)
		}
	}
}

func TestReadRequestEncapsulatedOffsets(t *testing.T) {
	reqHdr := "GET /page HTTP/1.1\r\nHost: example.com\r\n\r\n"
	resHdr := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n"
	raw := "RESPMOD icap://127.0.0.1:1344/safeanalyze ICAP/1.0\r\n" +
		"Host: 127.0.0.1\r\nAllow: 204\r\nPreview: 10\r\n" +
		"Encapsulated: req-hdr=0, res-hdr=" + itoa(len(reqHdr)) + ", res-body=" + itoa(len(reqHdr)+len(resHdr)) + "\r\n\r\n" +
		reqHdr + resHdr + "5\r\nhello\r\n0; ieof\r\n\r\n"
	br := bufio.NewReader(strings.NewReader(raw))
	req, err := ReadRequest(br)
	if err != nil {
		t.Fatal(err)
	}
	if req.Method != "RESPMOD" || string(req.ReqHdr) != reqHdr || string(req.ResHdr) != resHdr {
		t.Fatalf("bad split: method=%q req=%q res=%q", req.Method, req.ReqHdr, req.ResHdr)
	}
	if !req.HasBody() || req.BodyType != "res-body" || req.Preview != 10 || !req.Allows204() {
		t.Fatalf("bad request fields: %+v", req)
	}
	h, err := ParseHTTPHeader(req.ResHdr)
	if err != nil || h.Header.Get("Content-Type") != "text/html" || h.StartLine != "HTTP/1.1 200 OK" {
		t.Fatalf("http header: %+v %v", h, err)
	}
	sink := &chunkSink{keep: 100}
	ieof, err := readChunked(br, sink)
	if err != nil || !ieof || sink.buf.String() != "hello" {
		t.Fatalf("preview body: ieof=%v body=%q err=%v", ieof, sink.buf.String(), err)
	}
}

func TestReadRequestNullBody(t *testing.T) {
	resHdr := "HTTP/1.1 304 Not Modified\r\n\r\n"
	raw := "RESPMOD icap://x/safeanalyze ICAP/1.0\r\nEncapsulated: res-hdr=0, null-body=" + itoa(len(resHdr)) + "\r\n\r\n" + resHdr
	req, err := ReadRequest(bufio.NewReader(strings.NewReader(raw)))
	if err != nil {
		t.Fatal(err)
	}
	if req.HasBody() || req.BodyType != "null-body" || string(req.ResHdr) != resHdr || req.Preview != -1 || req.Allows204() {
		t.Fatalf("bad null-body request: %+v", req)
	}
}

func TestReadRequestMalformed(t *testing.T) {
	for _, raw := range []string{
		"GARBAGE\r\n\r\n",
		"RESPMOD icap://x/s HTTP/1.1\r\n\r\n",
		"RESPMOD icap://x/s ICAP/1.0\r\nHost: x\r\n\r\n", // no Encapsulated
		"RESPMOD icap://x/s ICAP/1.0\r\nEncapsulated: res-hdr=0, res-body=abc\r\n\r\n",
		"RESPMOD icap://x/s ICAP/1.0\r\nPreview: -3\r\nEncapsulated: null-body=0\r\n\r\n",
	} {
		if _, err := ReadRequest(bufio.NewReader(strings.NewReader(raw))); !errors.Is(err, ErrMalformed) {
			t.Errorf("%q: expected ErrMalformed, got %v", raw, err)
		}
	}
}

func TestReadChunked(t *testing.T) {
	raw := "4\r\nWiki\r\n5;name=val\r\npedia\r\nE\r\n in\r\n\r\nchunks.\r\n0\r\nTrailer: x\r\n\r\nNEXT"
	br := bufio.NewReader(strings.NewReader(raw))
	sink := &chunkSink{keep: 1 << 20}
	ieof, err := readChunked(br, sink)
	if err != nil || ieof {
		t.Fatalf("ieof=%v err=%v", ieof, err)
	}
	if got := sink.buf.String(); got != "Wikipedia in\r\n\r\nchunks." {
		t.Fatalf("body %q", got)
	}
	rest, _ := br.ReadString(0)
	if rest != "NEXT" {
		t.Fatalf("reader not positioned after body: %q", rest)
	}

	// Data beyond the keep limit is counted but dropped.
	sink = &chunkSink{keep: 3}
	if _, err := readChunked(bufio.NewReader(strings.NewReader("6\r\nabcdef\r\n0\r\n\r\n")), sink); err != nil {
		t.Fatal(err)
	}
	if sink.buf.String() != "abc" || sink.total != 6 {
		t.Fatalf("cap: body=%q total=%d", sink.buf.String(), sink.total)
	}

	for _, bad := range []string{"zz\r\nab\r\n0\r\n\r\n", "2\r\nabX\r\n0\r\n\r\n", "5\r\nab"} {
		if _, err := readChunked(bufio.NewReader(strings.NewReader(bad)), &chunkSink{keep: 10}); err == nil {
			t.Errorf("%q: expected error", bad)
		}
	}
}

func itoa(n int) string { return strconv.Itoa(n) }
