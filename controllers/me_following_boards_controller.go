package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// MeFollowingBoardsController is the controller that is responsible for all
// /v1/me/following/boards/ endpoints in the Pinterest API.
type MeFollowingBoardsController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeFollowingBoardsController instantiates a new MeFollowingBoardsController
func newMeFollowingBoardsController(wc *wrecker.Wrecker) *MeFollowingBoardsController {
	return &MeFollowingBoardsController{
		wreckerClient: wc,
	}
}

// MeFollowingBoardsFetchOptionals is a struct that represents the optional
// parameters for the Fetch method
type MeFollowingBoardsFetchOptionals struct {
	Cursor string
}

// Fetch loads the boards that the authorized user follows
// Endpoint: [GET] /v1/me/following/boards/
func (mfbc *MeFollowingBoardsController) Fetch(optionals *MeFollowingBoardsFetchOptionals) (*[]models.Board, *models.Page, error) {
	// Build request
	response := new(models.Response)
	response.Data = &[]models.Board{}
	request := mfbc.wreckerClient.Get("/me/following/boards/").
		URLParam("fields", "id,name,url,counts,created_at,creator,description,image,privacy,reason").
		Into(response)
	if optionals.Cursor != "" {
		request.URLParam("cursor", optionals.Cursor)
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

// Create follows a board for the authorized user
// Endpoint: [POST] /v1/me/following/boards/
func (mfbc *MeFollowingBoardsController) Create(boardSpec string) error {
	// Build request
	response := new(models.Response)
	request := mfbc.wreckerClient.Post("/me/following/boards/").
		FormParam("board", boardSpec).
		Into(response)

	// Execute request
	resp, err := request.Execute()

	// Error from Wrecker
	if err != nil {
		return err
	}

	// Status code
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return &models.PinterestError{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
		}
	}

	// OK
	return nil
}

// Delete unfollows a board for the authorized user
// Endpoint: [DELETE] /v1/me/following/boards/
func (mfbc *MeFollowingBoardsController) Delete(boardSpec string) error {
	// Build request
	response := new(models.Response)
	request := mfbc.wreckerClient.Delete("/me/following/boards/" + boardSpec + "/").
		Into(response)

	// Execute request
	resp, err := request.Execute()

	// Error from Wrecker
	if err != nil {
		return err
	}

	// Status code
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return &models.PinterestError{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
		}
	}

	// OK
	return nil
}
