package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// UsersController is the controller that is responsible for all
// /v1/users/ endpoints in the Pinterest API.
type UsersController struct {
	wreckerClient *wrecker.Wrecker
}

// NewUsersController instantiates a new UsersController.
func NewUsersController(wc *wrecker.Wrecker) *UsersController {
	return &UsersController{
		wreckerClient: wc,
	}
}

// Fetch loads a user from their username.
// Endpoint: [GET] /v1/users/<user>/
func (uc *UsersController) Fetch(username string) (*models.User, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = new(models.User)
	httpResp, err := uc.wreckerClient.Get("/users/"+username+"/").
		URLParam("fields", models.USER_FIELDS).
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*models.User), nil
}
