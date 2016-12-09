package models

import (
	"encoding/json"
)

type PinterestError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Type       string `json:"type"`
}

func (e *PinterestError) Error() string {
	out, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return "PinterestError: " + string(out)
}
