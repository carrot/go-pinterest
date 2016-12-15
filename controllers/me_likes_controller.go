package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// MeLikesController is the controller that is responsible for all
// /v1/me/likes/ endpoints in the Pinterest API.
type MeLikesController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeController instantiates a new MeLikesController
func newMeLikesController(wc *wrecker.Wrecker) *MeLikesController {
	return &MeLikesController{
		wreckerClient: wc,
	}
}

func (*MeLikesController) Fetch() {
	// TODO
}
