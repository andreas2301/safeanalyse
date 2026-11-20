package sync

import (
	"bytes"
	"net/http"
)

func Upload(endpoint string, body []byte) error {
	_, err := http.Post(endpoint, "application/json", bytes.NewReader(body))
	return err
}
