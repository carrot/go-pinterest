package controllers

import (
	"github.com/BrandonRomano/wrecker"
)

// BoardPinsController is the controller that is responsible for all
// board pin specific endpoints in the Pinterest API.
// https://developers.pinterest.com/docs/api/pins/
type BoardPinsController struct {
	wreckerClient *wrecker.Wrecker
}

// NewBoardPinsController instantiates a new BoardsController.
func NewBoardPinsController(wc *wrecker.Wrecker) *BoardPinsController {
	return &BoardPinsController{
		wreckerClient: wc,
	}
}
