package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// MeFollowingInterestsController is the controller that is responsible for all
// /v1/me/following/interests/ endpoints in the Pinterest API.
type MeFollowingInterestsController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeFollowingInterestsController instantiates a new MeFollowingInterestsController
func newMeFollowingInterestsController(wc *wrecker.Wrecker) *MeFollowingInterestsController {
	return &MeFollowingInterestsController{
		wreckerClient: wc,
	}
}

// MeFollowingInterestsFetchOptionals is a struct that represents the optional
// parameters for the Fetch method
type MeFollowingInterestsFetchOptionals struct {
	Cursor string
}

// Fetch loads the authorized users interests
func (mfic *MeFollowingInterestsController) Fetch(optionals *MeFollowingInterestsFetchOptionals) (*[]models.Interest, *models.Page, error) {
	// Build request
	response := new(models.Response)
	response.Data = &[]models.Interest{}
	request := mfic.wreckerClient.Get("/me/following/interests/").
		URLParam("fields", "id,name").
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
	return response.Data.(*[]models.Interest), &response.Page, nil
}
