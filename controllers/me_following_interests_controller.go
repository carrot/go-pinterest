package controllers

import (
	"github.com/BrandonRomano/wrecker"
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

func (MeFollowingInterestsController) Fetch() {
	// TODO
}