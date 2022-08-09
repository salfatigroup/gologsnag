package gologsnag

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// options to publish a new insight to logsnag
type InsightOptions struct {
    Title string `json:"title"`
    Value interface{} `json:"value"`
    Icon string `json:"icon,omitempty"`
}

// extend the InsightOptions struct with the Porject ID
type insightOptionsRequest struct {
    // inherit InsightOptions
    InsightOptions

    // extend with Project ID
    Project string `json:"project"`
}

// publish a new insight to LogSnag
func (l *LogSnag) Insight(ctx context.Context, options *InsightOptions) error {
    // add the project id to the options
    optionsWithProject := &insightOptionsRequest{
        InsightOptions: *options,
        Project: l.project,
    }

    // get request body from options
    reqBody, err := json.Marshal(optionsWithProject)
    if err != nil {
        return err
    }

    req, err := http.NewRequestWithContext(ctx, "POST", LOGSNAG_ENDPOINT+"/insight", bytes.NewBuffer(reqBody))
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
        return fmt.Errorf("LogSnag error code: %d", resp.StatusCode)
    }

    return nil
}
