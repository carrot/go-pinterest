package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/jschwehn/go-pinterest/models"
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

// Fetch loads the authorized users boards
// Endpoint: [GET] /v1/me/boards/
func (mbc *MeBoardsController) Fetch() (*[]models.Board, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.Board{}
	httpResp, err := mbc.wreckerClient.Get("/me/boards/").
		URLParam("fields", models.BOARD_FIELDS).
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*[]models.Board), nil
}
