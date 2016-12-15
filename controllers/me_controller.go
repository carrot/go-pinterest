package controllers

import (
	"github.com/BrandonRomano/wrecker"
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

func (*MeController) Fetch() {
	// TODO
}
