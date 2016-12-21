package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// MeController is the controller that is responsible for all
// /v1/me/ endpoints in the Pinterest API.
type MeController struct {
	wreckerClient *wrecker.Wrecker
	Boards        *MeBoardsController
	Followers     *MeFollowersController
	Following     *MeFollowingController
	Likes         *MeLikesController
	Pins          *MePinsController
	Search        *MeSearchController
}

// NewMeController instantiates a new MeController
func NewMeController(wc *wrecker.Wrecker) *MeController {
	return &MeController{
		wreckerClient: wc,
		Boards:        newMeBoardsController(wc),
		Followers:     newMeFollowersController(wc),
		Following:     newMeFollowingController(wc),
		Likes:         newMeLikesController(wc),
		Pins:          newMePinsController(wc),
		Search:        newMeSearchController(wc),
	}
}

// Fetch loads the authorized users info
// Endpoint: [GET] /v1/me/
func (mc *MeController) Fetch() (*models.User, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = new(models.User)
	httpResp, err := mc.wreckerClient.Get("/me/").
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
