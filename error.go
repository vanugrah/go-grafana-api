package gapi

type ErrNotFound struct {
	Msg string
}

func (e *ErrNotFound) Error() string {
	return e.Msg
}

func NewErrNotFound(msg string) error {
	return &ErrNotFound {
		Msg: msg,
	}
}

// A GrafanaMessage contains the json error message received when http request failed
type GrafanaErrorMessage struct {
	Message string `json:"message"`
}
