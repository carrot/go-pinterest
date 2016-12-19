package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// MeSearchController is the controller that is responsible for all
// /v1/me/search/ endpoints in the Pinterest API.
type MeSearchController struct {
	wreckerClient *wrecker.Wrecker
	Boards        *MeSearchBoardsController
	Pins          *MeSearchPinsController
}

// newMeSearchController instantiates a new MeSearchController
func newMeSearchController(wc *wrecker.Wrecker) *MeSearchController {
	return &MeSearchController{
		wreckerClient: wc,
		Boards:        newMeSearchBoardsController(wc),
		Pins:          newMeSearchPinsController(wc),
	}
}
