package loki

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// LokiClient handles sending log messages to Loki.
type LokiClient struct {
	URL    string
	Client *http.Client
}

type LokiLogEntry struct {
	Ts   string `json:"ts"`
	Line string `json:"line"`
}

type LokiStream struct {
	Stream map[string]string `json:"stream"`
	Values [][]string        `json:"values"`
}

type LokiPayload struct {
	Streams []LokiStream `json:"streams"`
}

// NewLokiClient creates a new instance of LokiClient.
func NewLokiClient(url string) *LokiClient {
	return &LokiClient{
		URL:    url,
		Client: &http.Client{},
	}
}

func InitializeLokiClient(url string) *LokiClient {
	return NewLokiClient(url)
}

func (c *LokiClient) SendLog(level, message string) error {
	// Prepare the payload
	payload := LokiPayload{
		Streams: []LokiStream{
			{
				Stream: map[string]string{"level": level},
				Values: [][]string{{time.Now().Format(time.RFC3339Nano), message}},
			},
		},
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return err
	}

	// Create the request
	req, err := http.NewRequest("POST", c.URL, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode >= 400 {
		return fmt.Errorf("Loki request failed with status code: %d", resp.StatusCode)
	}

	// Reset the buffer for reuse
	buf.Reset()

	return nil
}
