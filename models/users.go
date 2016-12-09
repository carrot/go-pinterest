package models

import (
	"github.com/joeshaw/iso8601"
)

type User struct {
	Id        string       `json:"id"`
	Username  string       `json:"username"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Bio       string       `json:"bio"`
	CreatedAt iso8601.Time `json:"created_at"`
	Counts    UserCounts   `json:"counts"`
	Image     UserImages   `json:"image"`
}

type UserImages struct {
	Size_60x60 string `60x60`
}

type Image struct {
	Url    string `json:"url"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
}

type UserCounts struct {
	Pins      int32 `json:"pins"`
	Following int32 `json:"following"`
	Followers int32 `json:"followers"`
	Boards    int32 `json:"boards"`
	Likes     int32 `json:"likes"`
}


