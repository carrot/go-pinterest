package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// MeFollowingController is the controller that is responsible for all
// /v1/me/following/ endpoints in the Pinterest API.
type MeFollowingController struct {
	wreckerClient *wrecker.Wrecker
	Boards        *MeFollowingBoardsController
	Interests     *MeFollowingInterestsController
	Users         *MeFollowingUsersController
}

// newMeFollowingController instantiates a new MeFollowingController
func newMeFollowingController(wc *wrecker.Wrecker) *MeFollowingController {
	return &MeFollowingController{
		wreckerClient: wc,
		Boards:        newMeFollowingBoardsController(wc),
		Interests:     newMeFollowingInterestsController(wc),
		Users:         newMeFollowingUsersController(wc),
	}
}
