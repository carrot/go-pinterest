package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// MeSearchPinsController is the controller that is responsible for all
// /v1/me/search/pins/ endpoints in the Pinterest API.
type MeSearchPinsController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeSearchPinsController instantiates a new MeSearchPinsController
func newMeSearchPinsController(wc *wrecker.Wrecker) *MeSearchPinsController {
	return &MeSearchPinsController{
		wreckerClient: wc,
	}
}

func (*MeSearchPinsController) Fetch() {
	// TODO
}
