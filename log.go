package gologsnag

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// options for publishing logsnag events
type PublishOptions struct {
	Channel     string `json:"channel"`
	Event       string `json:"event"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Tags        *Tags  `json:"tags,omitempty"`
	Notify      bool   `json:"notify,omitempty"`
}

// extend the PublishOptions struct with the Porject ID
type publishOptionsRequest struct {
	// inherit PublishOptions
	PublishOptions

	// extend with Project ID
	Project string `json:"project"`
}

// tag type
type Tags map[string]interface{}

// extend the tags with additional values
func (t *Tags) Add(key string, value interface{}) {
	if t == nil {
		t = &Tags{}
	}
	(*t)[key] = value
}

// Publish a new event to LogSnag
// options: LogSnag options
func (l *LogSnag) Publish(ctx context.Context, options *PublishOptions) error {
	// add the project id to the options
	optionsWithProject := &publishOptionsRequest{
		PublishOptions: *options,
		Project:        l.project,
	}

	// get request body from options
	reqBody, err := json.Marshal(optionsWithProject)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", LOGSNAG_ENDPOINT+"/log", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	// add request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+l.token)

	// send request
	resp, err := l.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("LogSnag error code: %s", resp.Status)
	}

	return nil
}
