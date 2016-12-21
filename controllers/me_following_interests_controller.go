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
// Endpoint: [GET] /v1/me/following/interests/
func (mfic *MeFollowingInterestsController) Fetch(optionals *MeFollowingInterestsFetchOptionals) (*[]models.Interest, *models.Page, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.Interest{}
	request := mfic.wreckerClient.Get("/me/following/interests/").
		URLParam("fields", models.INTEREST_FIELDS).
		Into(resp)
	if optionals.Cursor != "" {
		request.URLParam("cursor", optionals.Cursor)
	}
	httpResp, err := request.Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, nil, err
	}

	// OK
	return resp.Data.(*[]models.Interest), &resp.Page, nil
}
