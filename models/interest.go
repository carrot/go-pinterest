package models

const INTEREST_FIELDS = "id,name"

// Interest is a struct that represents an individual interest
// from the Pinterest API.
type Interest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
