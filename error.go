package gapi

import "fmt"

// A GrafanaMessage contains the json error message received when http request failed.
// On a 412 Error, an additional Status field may be present explainin
type GrafanaErrorMessage struct {
	Message string `json:"message"`
	Status  string `json:status, omitempty`
}

func (gem GrafanaErrorMessage) String() string {
	if gem.Status != "" {
		return fmt.Sprintf("%s\" status=\"%s", gem.Message, gem.Status)
	}
	return gem.Message
}

type GrafanaError struct {
	StatusCode int
	Message    string
}

func (ge GrafanaError) Error() string {
	return fmt.Sprintf("Request to Grafana returned status-code=\"%d\" message=\"%s\"", ge.StatusCode, ge.Message)
}
