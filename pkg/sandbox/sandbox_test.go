package sandbox

import (
	"slices"
	"strings"
	"testing"
)

func TestDockerCommandHardeningFlags(t *testing.T) {
	r := NewRunner(Config{
		Mode:          "docker",
		DockerImage:   "alpine:3.20",
		ReadOnlyPaths: []string{"/tmp/ingest"},
	})
	cmd := r.Command("sh", "-c", "ls")
	if cmd == nil {
		t.Fatal("expected non-nil command")
	}
	args := cmd.Args
	for _, want := range []string{"--network", "none", "--read-only", "--cap-drop", "ALL", "no-new-privileges:true", "alpine:3.20"} {
		if !slices.Contains(args, want) {
			t.Errorf("docker args missing %q: %v", want, args)
		}
	}
	joined := strings.Join(args, " ")
	if !strings.Contains(joined, "/tmp/ingest:/tmp/ingest:ro") {
		t.Errorf("expected read-only mount, got: %s", joined)
	}
}

func TestFirejailCommandHardeningFlags(t *testing.T) {
	r := NewRunner(Config{
		Mode:            "firejail",
		FirejailProfile: "custom.profile",
		ReadOnlyPaths:   []string{"/tmp/ingest"},
	})
	cmd := r.Command("echo", "hi")
	if cmd == nil {
		t.Fatal("expected non-nil command")
	}
	for _, want := range []string{"--net=none", "--private", "--shell=none", "--caps.drop=all", "--profile=custom.profile"} {
		if !slices.Contains(cmd.Args, want) {
			t.Errorf("firejail args missing %q: %v", want, cmd.Args)
		}
	}
}

func TestNoneModeAvailable(t *testing.T) {
	r := NewRunner(Config{Mode: "none"})
	if !r.Available() {
		t.Error("expected mode=none to always be available")
	}
}
