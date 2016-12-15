package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// MeBoardsController is the controller that is responsible for all
// /v1/me/boards/ endpoints in the Pinterest API.
type MeBoardsController struct {
	wreckerClient *wrecker.Wrecker
	Suggested     *MeBoardsSuggestedController
}

// newMeBoardsController instantiates a new MeBoardsController
func newMeBoardsController(wc *wrecker.Wrecker) *MeBoardsController {
	return &MeBoardsController{
		wreckerClient: wc,
		Suggested:     newMeBoardsSuggestedController(wc),
	}
}

func (*MeBoardsController) Fetch() {
	// TODO
}
