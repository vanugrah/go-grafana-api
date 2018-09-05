package gapi

// A GrafanaMessage contains the json error message received when http request failed
type GrafanaErrorMessage struct {
	Message string `json:"message"`
}
