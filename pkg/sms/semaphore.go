package sms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const semaphoreAPIURL = "https://api.semaphore.co/api/v4/messages"

// Client represents a Semaphore SMS API client.
type Client struct {
	APIKey     string
	SenderName string
	HTTPClient *http.Client
}

// Result represents the response returned by the Semaphore API.
type Result struct {
	StatusCode int             `json:"status_code"`
	Body       json.RawMessage `json:"body,omitempty"`
	Raw        string          `json:"raw,omitempty"`
}

// NewClient creates a new Semaphore SMS client.
func NewClient(apiKey, senderName string) *Client {
	return &Client{
		APIKey:     apiKey,
		SenderName: senderName,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) SendSMS(ctx context.Context, recipientNumber string, message string) (*Result, error) {
	if c == nil {
		return nil, fmt.Errorf("SMS client is nil")
	}

	if strings.TrimSpace(c.APIKey) == "" {
		return nil, fmt.Errorf("Semaphore API key is required")
	}

	if strings.TrimSpace(recipientNumber) == "" {
		return nil, fmt.Errorf("recipient number is required")
	}

	if strings.TrimSpace(message) == "" {
		return nil, fmt.Errorf("SMS message is required")
	}

	if c.HTTPClient == nil {
		return nil, fmt.Errorf("HTTP client is not configured")
	}

	// Prepare the form payload.
	payload := url.Values{}
	payload.Set("apikey", c.APIKey)
	payload.Set("number", recipientNumber)
	payload.Set("message", message)

	if c.SenderName != "" {
		payload.Set("sendername", c.SenderName)
	}

	// Create the HTTP POST request.
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		semaphoreAPIURL,
		bytes.NewBufferString(payload.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("create SMS request: %w", err)
	}

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	// Send the request.
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send SMS request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body.
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read SMS API response: %w", err)
	}

	result := &Result{
		StatusCode: resp.StatusCode,
		Raw:        string(responseBody),
	}

	// Check the HTTP status code.
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf(
			"SMS API returned HTTP %d: %s",
			resp.StatusCode,
			string(responseBody),
		)
	}

	// Parse the JSON response.
	if !json.Valid(responseBody) {
		// Preserve the raw response if it isn't valid JSON.
		return result, nil
	}

	result.Body = json.RawMessage(responseBody)

	return result, nil
}
