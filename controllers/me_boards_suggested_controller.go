package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/jschwehn/go-pinterest/models"
	"strconv"
)

// MeBoardsSuggestedController is the controller that is responsible for all
// /v1/me/boards/suggested/ endpoints in the Pinterest API.
type MeBoardsSuggestedController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeBoardsSuggestedController instantiates a new MeBoardsSuggestedController
func newMeBoardsSuggestedController(wc *wrecker.Wrecker) *MeBoardsSuggestedController {
	return &MeBoardsSuggestedController{
		wreckerClient: wc,
	}
}

// MeBoardsSuggestedFetchOptionals is a struct that represents the optional
// parameters for the Fetch method
type MeBoardsSuggestedFetchOptionals struct {
	Count int32
	Pin   string
}

// Fetch loads board suggestions for the logged in user
// Endpoint: [GET] /v1/me/boards/suggested/
func (mbsc *MeBoardsSuggestedController) Fetch(optionals *MeBoardsSuggestedFetchOptionals) (*[]models.Board, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.Board{}
	request := mbsc.wreckerClient.Get("/me/boards/suggested/").
		URLParam("fields", models.BOARD_FIELDS).
		Into(resp)
	if optionals.Count != 0 {
		request.URLParam("count", strconv.Itoa(int(optionals.Count)))
	}
	if optionals.Pin != "" {
		request.URLParam("pin", optionals.Pin)
	}
	httpResp, err := request.Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*[]models.Board), nil
}
