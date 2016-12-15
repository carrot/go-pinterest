package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// MeSearchBoardsController is the controller that is responsible for all
// /v1/me/search/boards/ endpoints in the Pinterest API.
type MeSearchBoardsController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeSearchBoardsController instantiates a new MeSearchBoardsController
func newMeSearchBoardsController(wc *wrecker.Wrecker) *MeSearchBoardsController {
	return &MeSearchBoardsController{
		wreckerClient: wc,
	}
}

func (MeSearchBoardsController) Fetch() {
	// TODO
}
