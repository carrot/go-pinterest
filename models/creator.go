package models

type Creator struct {
	Url       string `json:"url"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Id        string `json:"id"`
}
