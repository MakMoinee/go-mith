package mithclientrest

import (
	"encoding/json"
	"fmt"
)

type Fault struct {
	StatusCode int
	Message    string
	Body       []byte
}

func (f *Fault) Error() string {
	if f.Message != "" {
		return fmt.Sprintf(
			"REST API returned HTTP %d: %s",
			f.StatusCode,
			f.Message,
		)
	}

	return fmt.Sprintf(
		"REST API returned HTTP %d",
		f.StatusCode,
	)
}

func parseFault(statusCode int, body []byte) error {
	fault := &Fault{
		StatusCode: statusCode,
		Body:       body,
	}

	// Try to parse common API error formats.
	var response struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	if err := json.Unmarshal(body, &response); err == nil {
		switch {
		case response.Message != "":
			fault.Message = response.Message
		case response.Error != "":
			fault.Message = response.Error
		}
	}

	return fault
}
