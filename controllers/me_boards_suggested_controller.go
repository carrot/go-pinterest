package controllers

import "github.com/BrandonRomano/wrecker"

// MeBoardsSuggestedController is the controller that is responsible for all
// /v1/me/boards/suggested/ endpoints in the Pinterest API.
type MeBoardsSuggestedController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeBoardsController instantiates a new MeBoardsController
func newMeBoardsSuggestedController(wc *wrecker.Wrecker) *MeBoardsSuggestedController {
	return &MeBoardsSuggestedController{
		wreckerClient: wc,
	}
}

func (*MeBoardsSuggestedController) Fetch() {
	// TODO
}
