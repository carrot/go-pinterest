package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
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
	response := new(models.Response)
	response.Data = &[]models.Board{}
	resp, err := mbc.wreckerClient.Get("/me/boards/").
		URLParam("fields", models.BOARD_FIELDS).
		Into(response).
		Execute()

	// Error from Wrecker
	if err != nil {
		return nil, err
	}

	// Status code
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, &models.PinterestError{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
		}
	}

	// OK
	return response.Data.(*[]models.Board), nil
}
