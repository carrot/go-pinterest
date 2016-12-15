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
	response := new(models.Response)
	response.Data = new(models.User)

	resp, err := uc.wreckerClient.Get("/users/"+username+"/").
		URLParam("fields", "first_name,last_name,url,account_type,bio,counts,created_at,image,username").
		Into(response).
		Execute()

	// Error from Wrecker
	if err != nil {
		return nil, err
	}

	// Status code
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, &models.PinterestError{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
		}
	}

	// OK
	return response.Data.(*models.User), nil
}
