package controllers

import (
	"github.com/BrandonRomano/wrecker"
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

func (*MeFollowersController) Fetch() {
	// TODO
}
