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
func (msbc *MeSearchBoardsController) Fetch(query string, optionals *MeSearchBoardsFetchOptionals) (*[]models.Board, *models.Page, error) {
	// Build request
	response := new(models.Response)
	response.Data = &[]models.Board{}
	request := msbc.wreckerClient.Get("/me/search/boards/").
		URLParam("fields", "id,name,url,counts,created_at,creator,description,privacy,image,reason").
		URLParam("query", query).
		Into(response)
	if optionals.Cursor != "" {
		request.URLParam("cursor", optionals.Cursor)
	}
	if optionals.Limit != 0 {
		request.URLParam("limit", strconv.Itoa(optionals.Limit))
	}

	// Execute request
	resp, err := request.Execute()

	// Error from Wrecker
	if err != nil {
		return nil, nil, err
	}

	// Status code
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, nil, &models.PinterestError{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
		}
	}

	// OK
	return response.Data.(*[]models.Board), &response.Page, nil
}
