package models

// Response is the base struct for all responses
// that come back from the Pinterest API.
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Page    Page        `json:"page"`
}

type Page struct {
	Cursor string `json:"cursor"`
	Next   string `json:"next"`
}
