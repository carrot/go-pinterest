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
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.Board{}
	request := mfbc.wreckerClient.Get("/me/following/boards/").
		URLParam("fields", models.BOARD_FIELDS).
		Into(resp)
	if optionals.Cursor != "" {
		request.URLParam("cursor", optionals.Cursor)
	}
	httpResp, err := request.Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, nil, err
	}

	// OK
	return resp.Data.(*[]models.Board), &resp.Page, nil
}

// Create follows a board for the authorized user
// Endpoint: [POST] /v1/me/following/boards/
func (mfbc *MeFollowingBoardsController) Create(boardSpec string) error {
	// Build + execute request
	resp := new(models.Response)
	httpResp, err := mfbc.wreckerClient.Post("/me/following/boards/").
		FormParam("board", boardSpec).
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return err
	}

	// OK
	return nil
}

// Delete unfollows a board for the authorized user
// Endpoint: [DELETE] /v1/me/following/boards/
func (mfbc *MeFollowingBoardsController) Delete(boardSpec string) error {
	// Build + execute request
	resp := new(models.Response)
	httpResp, err := mfbc.wreckerClient.Delete("/me/following/boards/" + boardSpec + "/").
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return err
	}

	// OK
	return nil
}
