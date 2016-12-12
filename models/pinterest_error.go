package models

import (
	"encoding/json"
)

// PinterestError is a custom error that is passed for all
// non 200 responses from the API.
type PinterestError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *PinterestError) Error() string {
	out, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return "PinterestError: " + string(out)
}
