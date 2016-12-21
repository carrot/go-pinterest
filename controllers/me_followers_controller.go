package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// MeFollowersController is the controller that is responsible for all
// /v1/me/followers/ endpoints in the Pinterest API.
type MeFollowersController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeFollowersController instantiates a new MeFollowersController
func newMeFollowersController(wc *wrecker.Wrecker) *MeFollowersController {
	return &MeFollowersController{
		wreckerClient: wc,
	}
}

// MeFollowersFetchOptionals is a struct that represents the optional
// parameters for the Fetch method
type MeFollowersFetchOptionals struct {
	Cursor string
}

// Fetch loads the users that follow the logged in user
// Endpoint: [GET] /v1/me/boards/followers/
func (mfc *MeFollowersController) Fetch(optionals *MeFollowersFetchOptionals) (*[]models.User, *models.Page, error) {
	// Build request
	response := new(models.Response)
	response.Data = &[]models.User{}
	request := mfc.wreckerClient.Get("/me/followers/").
		URLParam("fields", models.USER_FIELDS).
		Into(response)
	if optionals.Cursor != "" {
		request.URLParam("cursor", optionals.Cursor)
	}

	// Execute request
	resp, err := request.Execute()

	// Error from Wrecker
	if err != nil {
		return nil, nil, err
	}

	// Status code
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, nil, &models.PinterestError{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
		}
	}

	// OK
	return response.Data.(*[]models.User), &response.Page, nil
}
