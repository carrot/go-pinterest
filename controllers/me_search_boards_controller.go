package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
	"strconv"
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

// MeSearchBoardsFetchOptionals is a struct that represents the optional parameters
// that can be passed to the Fetch endpoint
type MeSearchBoardsFetchOptionals struct {
	Cursor string
	Limit  int
}

// Fetch searches the logged in user's Boards
// Endpoint: [GET] /v1/me/search/boards/
func (msbc *MeSearchBoardsController) Fetch(query string, optionals *MeSearchBoardsFetchOptionals) (*[]models.Board, *models.Page, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.Board{}
	request := msbc.wreckerClient.Get("/me/search/boards/").
		URLParam("fields", models.BOARD_FIELDS).
		URLParam("query", query).
		Into(resp)
	if optionals.Cursor != "" {
		request.URLParam("cursor", optionals.Cursor)
	}
	if optionals.Limit != 0 {
		request.URLParam("limit", strconv.Itoa(optionals.Limit))
	}
	httpResp, err := request.Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, nil, err
	}

	// OK
	return resp.Data.(*[]models.Board), &resp.Page, nil
}
