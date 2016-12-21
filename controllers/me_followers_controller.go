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
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.User{}
	request := mfc.wreckerClient.Get("/me/followers/").
		URLParam("fields", models.USER_FIELDS).
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
	return resp.Data.(*[]models.User), &resp.Page, nil
}
