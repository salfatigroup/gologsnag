package gologsnag

import (
	"net/http"
)

const (
    LOGSNAG_ENDPOINT = "https://api.logsnag.com/v1"
)

// LogSnag client
type LogSnag struct {
    token string
    project string
    client *http.Client
}

// Create a new LogSnag client
// token: LogSnag API token
// project: LogSnag project name
func NewLogSnag(token string, project string) *LogSnag {
    return &LogSnag{
        token: token,
        project: project,
        client: &http.Client{},
    }
}
