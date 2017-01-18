package models

// AccessToken is a struct that represents a Access Token
// response from the Pinterest API.
type AccessToken struct {
	AccessToken      string   `json:"access_token"`
	TokenType        string   `json:"token_type"`
	Scope            []string `json:"scope"`
	ErrorDescription string   `json:"error_description"`
	Error            string   `json:"error"`
}
