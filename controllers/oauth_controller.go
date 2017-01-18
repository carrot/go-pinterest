package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// OAuthController is the controller that is responsible
// for all /v1/oauth endpoints in the Pinterest API.
type OAuthController struct {
	wreckerClient *wrecker.Wrecker
	Token         *OAuthTokenController
}

// NewOAuthController instantiates a new OAuthController
func NewOAuthController(wc *wrecker.Wrecker) *OAuthController {
	return &OAuthController{
		wreckerClient: wc,
		Token:         newOAuthTokenController(wc),
	}
}
