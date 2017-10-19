package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// OAuthTokenController is the controller that is responsible
// for all /v1/oauth/token endpoints in the Pinterest API.
type OAuthTokenController struct {
	wreckerClient *wrecker.Wrecker
}

// newOAuthTokenController instantiates a new OAuthTokenController
func newOAuthTokenController(wc *wrecker.Wrecker) *OAuthTokenController {
	return &OAuthTokenController{
		wreckerClient: wc,
	}
}

// Create generates an access token
// Endpoint: [POST] /v1/oauth/token
func (otc *OAuthTokenController) Create(clientId, clientSecret, accessCode string) (*models.AccessToken, error) {
	// Build + execute request
	accessToken := new(models.AccessToken)
	httpResp, err := otc.wreckerClient.Post("/oauth/token").
		URLParam("grant_type", "authorization_code").
		URLParam("client_id", clientId).
		URLParam("client_secret", clientSecret).
		URLParam("code", accessCode).
		Into(accessToken).
		Execute()

	if err != nil {
		if _, ok := err.(wrecker.ResponseError); !ok {
			return nil, err
		}
	}

	if !(httpResp.StatusCode >= 200 && httpResp.StatusCode < 300) {
		return nil, &models.PinterestError{
			StatusCode: httpResp.StatusCode,
			Message:    accessToken.Error,
		}
	}

	// OK
	return accessToken, nil
}
