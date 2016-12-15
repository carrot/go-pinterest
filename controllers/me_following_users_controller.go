package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// MeFollowingUsersController is the controller that is responsible for all
// /v1/me/following/users/ endpoints in the Pinterest API.
type MeFollowingUsersController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeFollowingUsersController instantiates a new MeFollowingUsersController
func newMeFollowingUsersController(wc *wrecker.Wrecker) *MeFollowingUsersController {
	return &MeFollowingUsersController{
		wreckerClient: wc,
	}
}

func (*MeFollowingUsersController) Fetch() {
	// TODO
}

func (*MeFollowingUsersController) Create() {
	// TODO
}

func (*MeFollowingUsersController) Delete() {
	// TODO
}
