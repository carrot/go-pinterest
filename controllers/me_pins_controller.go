package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// MePinsController is the controller that is responsible for all
// /v1/me/pins/ endpoints in the Pinterest API.
type MePinsController struct {
	wreckerClient *wrecker.Wrecker
}

// newMePinsController instantiates a new MePinsController
func newMePinsController(wc *wrecker.Wrecker) *MePinsController {
	return &MePinsController{
		wreckerClient: wc,
	}
}

func (MePinsController) Fetch() {
	// TODO
}
