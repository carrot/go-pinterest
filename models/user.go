package models

import (
	"github.com/BrandonRomano/iso8601"
)

const USER_FIELDS = "first_name,last_name,url,account_type,bio,counts,created_at,id,image,username"

// User is a struct that represents an individual user
// from the Pinterest API.
type User struct {
	Id          string       `json:"id"`
	Username    string       `json:"username"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Bio         string       `json:"bio"`
	AccountType string       `json:"account_type"`
	Url         string       `json:"account_type"`
	CreatedAt   iso8601.Time `json:"created_at"`
	Counts      UserCounts   `json:"counts"`
	Image       Images       `json:"image"`
}

type UserCounts struct {
	Pins      int32 `json:"pins"`
	Following int32 `json:"following"`
	Followers int32 `json:"followers"`
	Boards    int32 `json:"boards"`
	Likes     int32 `json:"likes"`
}
