package gapi

import "fmt"

// A GrafanaMessage contains the json error message received when http request failed
type GrafanaErrorMessage struct {
	Message string `json:"message"`
}

type GrafanaError struct {
	StatusCode int
	Message    string
}

func (ge GrafanaError) Error() string {
	return fmt.Sprintf("Request to Grafana returned %d status code with the following message: %s", ge.StatusCode, ge.Message)
}
