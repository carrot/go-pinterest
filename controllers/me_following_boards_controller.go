package controllers

import "github.com/BrandonRomano/wrecker"

// MeFollowingBoardsController is the controller that is responsible for all
// /v1/me/following/boards/ endpoints in the Pinterest API.
type MeFollowingBoardsController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeFollowingBoardsController instantiates a new MeFollowingBoardsController
func newMeFollowingBoardsController(wc *wrecker.Wrecker) *MeFollowingBoardsController {
	return &MeFollowingBoardsController{
		wreckerClient: wc,
	}
}

func (*MeFollowingBoardsController) Fetch() {
	// TODO
}

func (*MeFollowingBoardsController) Create() {
	// TODO
}

func (*MeFollowingBoardsController) Delete() {
	// TODO
}
