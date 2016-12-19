package models

import (
	"github.com/BrandonRomano/iso8601"
)

// Board is a struct that represents an individual board
// from the Pinterest API.
type Board struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Url         string       `json:"url"`
	Description string       `json:"description"`
	Creator     Creator      `json:"creator"`
	CreatedAt   iso8601.Time `json:"created_at"`
	Counts      BoardCounts  `json:"counts"`
	Image       Images       `json:"image"`
	Privacy     string       `json:"privacy"`
}

type BoardCounts struct {
	Pins          int32 `json:"pins"`
	Collaborators int32 `json:"collaborators"`
	Followers     int32 `json:"followers"`
}
