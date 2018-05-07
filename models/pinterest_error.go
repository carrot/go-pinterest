package models

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonRomano/wrecker"
)

// PinterestError is a custom error that is passed for all
// non 200 responses from the API.
type PinterestError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Limit      TypeRatelimit
}

func (e *PinterestError) Error() string {
	out, _ := json.Marshal(e)
	return "PinterestError: " + string(out)
}

// WrapPinterestError takes a *http.Response and a Response and returns a
// PinterestError if one should be returned.
func WrapPinterestError(httpResponse *http.Response, bodyResponse *Response, err error) error {
	if err != nil {
		if _, ok := err.(wrecker.ResponseError); !ok {
			return err
		}
	}

	if !(httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300) {
		return &PinterestError{
			StatusCode: httpResponse.StatusCode,
			Message:    bodyResponse.Message,
			Limit:      GetRatelimit(httpResponse),
		}
	}

	return nil
}
