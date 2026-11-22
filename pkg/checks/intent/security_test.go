package intent

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestValidatePublicHTTPSURL(t *testing.T) {
	blocked := []string{
		"http://registry.npmjs.org/pkg.tgz",
		"https://localhost/pkg.tgz",
		"https://127.0.0.1/pkg.tgz",
		"https://[::1]/pkg.tgz",
		"https://169.254.169.254/latest/meta-data/",
		"https://10.0.0.5/pkg.tgz",
		"https://192.168.1.10/pkg.tgz",
		"https://172.16.0.1/pkg.tgz",
		"https://100.64.0.1/pkg.tgz",
		"https://metadata.google.internal/pkg.tgz",
	}
	for _, u := range blocked {
		if err := validatePublicHTTPSURL(u); err == nil {
			t.Errorf("expected %q to be rejected as SSRF/non-HTTPS", u)
		}
	}

	allowed := []string{
		"https://registry.npmjs.org/lodash/-/lodash-4.17.21.tgz",
		"https://1.1.1.1/pkg.tgz",
	}
	for _, u := range allowed {
		if err := validatePublicHTTPSURL(u); err != nil {
			t.Errorf("expected %q to be allowed, got: %v", u, err)
		}
	}
}

func TestUnpackTarballCapsFileCount(t *testing.T) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	for i := 0; i < maxTarballFiles+5; i++ {
		body := []byte("x")
		hdr := &tar.Header{
			Name:     fmt.Sprintf("package/f%d.js", i),
			Mode:     0644,
			Size:     int64(len(body)),
			Typeflag: tar.TypeReg,
		}
		if err := tw.WriteHeader(hdr); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write(body); err != nil {
			t.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}

	src := &depSource{id: "bomb@1.0.0", prefix: "deps/npm/bomb@1.0.0/"}
	err := unpackTarball(src, buf.Bytes())
	if err == nil || !strings.Contains(err.Error(), "maximum entry count") {
		t.Fatalf("expected entry count cap error, got %v", err)
	}
}

func TestFileLinesConcurrentRaceFree(t *testing.T) {
	f := NewFile("main.go", "line1\nline2\nline3\n")
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lines := f.Lines()
			if len(lines) != 4 {
				t.Errorf("expected 4 lines, got %d", len(lines))
			}
		}()
	}
	wg.Wait()
}
