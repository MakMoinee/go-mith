package mithclientrest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Config struct {
	BaseURL string

	Headers map[string]string

	HTTPClient *http.Client

	Timeout time.Duration

	OperationResolver OperationResolver
}

type Client struct {
	config *Config
}

func New(c *Config) *Client {
	if c.HTTPClient == nil {
		timeout := c.Timeout

		if timeout == 0 {
			timeout = 30 * time.Second
		}

		c.HTTPClient = &http.Client{
			Timeout: timeout,
		}
	}

	return &Client{
		config: c,
	}
}

// Do invokes a REST operation and returns the raw response
// as a byte slice.
//
// The operation to perform is determined by the type of
// the provided payload.
func (c *Client) Do(
	ctx context.Context,
	payload any,
) ([]byte, error) {
	operation, err := c.config.OperationResolver.OperationFor(payload)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}

	url := c.config.BaseURL + operation.Path

	req, err := http.NewRequestWithContext(
		ctx,
		operation.Method,
		url,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range c.config.Headers {
		req.Header.Set(key, value)
	}

	for key, value := range operation.Headers {
		req.Header.Set(key, value)
	}

	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < http.StatusOK ||
		resp.StatusCode >= http.StatusMultipleChoices {
		return nil, parseFault(resp.StatusCode, responseBody)
	}

	return responseBody, nil
}

// Exec invokes a REST operation and discards the response.
func (c *Client) Exec(
	ctx context.Context,
	payload any,
) error {
	_, err := c.Do(ctx, payload)

	return err
}
