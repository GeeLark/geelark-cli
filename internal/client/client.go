package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/geelark-tech/geelark-cli/internal/build"
	"github.com/geelark-tech/geelark-cli/internal/config"
	"github.com/geelark-tech/geelark-cli/internal/output"
	"github.com/google/uuid"
)

// Client is the GeeLark API client.
type Client struct {
	cfg        *config.Config
	httpClient *http.Client
}

// APIResponse is the standard GeeLark API response envelope.
type APIResponse struct {
	TraceID string          `json:"traceId"`
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Data    json.RawMessage `json:"data,omitempty"`
}

// New creates a new GeeLark API client.
func New(cfg *config.Config) *Client {
	return &Client{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 30 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
			},
		},
	}
}

// Post sends a POST request to the given API path with the given body.
// Uses Bearer token authentication.
func (c *Client) Post(path string, body interface{}) (*APIResponse, error) {
	url := c.cfg.BaseURL + path

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	} else {
		bodyReader = bytes.NewReader([]byte("{}"))
	}

	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Bearer token auth
	traceID := strings.ToUpper(strings.ReplaceAll(uuid.New().String(), "-", ""))
	req.Header.Set("traceId", traceID)
	req.Header.Set("Authorization", "Bearer "+c.cfg.Token)
	req.Header.Set("geelark-cli", build.Version)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	var apiResp APIResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &apiResp, nil
}

// PostBrowser sends a POST request to the Browser API (local) endpoint.
func (c *Client) PostBrowser(path string, body interface{}) (*APIResponse, error) {
	url := c.cfg.BrowserBaseURL + path

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	} else {
		bodyReader = bytes.NewReader([]byte("{}"))
	}

	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Bearer token auth (shared with Cloud Phone API)
	traceID := strings.ToUpper(strings.ReplaceAll(uuid.New().String(), "-", ""))
	req.Header.Set("traceId", traceID)
	req.Header.Set("Authorization", "Bearer "+c.cfg.Token)
	req.Header.Set("geelark-cli", build.Version)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	var apiResp APIResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &apiResp, nil
}

// PostBrowserAndPrint sends a POST request to Browser API and returns formatted output.
func (c *Client) PostBrowserAndPrint(path string, body interface{}) (string, error) {
	resp, err := c.PostBrowser(path, body)
	if err != nil {
		return "", err
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format response: %w", err)
	}

	return output.FormatResponse(string(data)), nil
}

func (c *Client) PutFile(uploadURL string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}

	req, err := http.NewRequest("PUT", uploadURL, file)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %w", err)
	}
	req.ContentLength = stat.Size()

	putClient := &http.Client{
		Timeout: 10 * time.Minute,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   30 * time.Second,
			ResponseHeaderTimeout: 60 * time.Second,
			MaxIdleConns:          10,
			IdleConnTimeout:       90 * time.Second,
		},
	}

	resp, err := putClient.Do(req)
	if err != nil {
		return fmt.Errorf("upload request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// PostAndPrint sends a POST request and returns output formatted according to --format flag.
func (c *Client) PostAndPrint(path string, body interface{}) (string, error) {
	resp, err := c.Post(path, body)
	if err != nil {
		return "", err
	}

	// Always produce JSON first
	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format response: %w", err)
	}

	jsonStr := string(data)

	// Apply format transformation
	return output.FormatResponse(jsonStr), nil
}
