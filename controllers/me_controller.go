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
	response := new(models.Response)
	response.Data = new(models.User)
	resp, err := mc.wreckerClient.Get("/me/").
		URLParam("fields", "first_name,last_name,url,account_type,bio,counts,created_at,id,image,username").
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
