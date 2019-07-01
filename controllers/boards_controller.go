package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/jschwehn/go-pinterest/models"
)

// BoardsController is the controller that is responsible for all
// /v1/boards/ endpoints in the Pinterest API.
type BoardsController struct {
	wreckerClient *wrecker.Wrecker
	Pins          *BoardsPinsController
}

// NewBoardsController instantiates a new BoardsController.
func NewBoardsController(wc *wrecker.Wrecker) *BoardsController {
	return &BoardsController{
		wreckerClient: wc,
		Pins:          newBoardsPinsController(wc),
	}
}

// Fetch loads a board from the board_spec (username/board-slug)
// Endpoint: [GET] /v1/boards/<board_spec:board>/
func (bc *BoardsController) Fetch(boardSpec string) (*models.Board, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = new(models.Board)
	httpResp, err := bc.wreckerClient.Get("/boards/"+boardSpec).
		URLParam("fields", models.BOARD_FIELDS).
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*models.Board), nil
}

// BoardCreateOptionals is a struct that represents the optional parameters
// that can be passed to the Create endpoint
type BoardCreateOptionals struct {
	Description string
}

// Create makes a new board
// Endpoint: [POST] /v1/boards/
func (bc *BoardsController) Create(boardName string, optionals *BoardCreateOptionals) (*models.Board, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = new(models.Board)
	httpResp, err := bc.wreckerClient.Post("/boards/").
		URLParam("fields", models.BOARD_FIELDS).
		FormParam("name", boardName).
		FormParam("description", optionals.Description).
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*models.Board), nil
}

// BoardUpdateOptionals is a struct that represents the optional parameters
// that can be passed to the Update endpoint
type BoardUpdateOptionals struct {
	Name        string
	Description string
}

// Update updates an existing board
// Endpoint: [PATCH] /v1/boards/<board_spec:board>/
func (bc *BoardsController) Update(boardSpec string, optionals *BoardUpdateOptionals) (*models.Board, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = new(models.Board)
	httpResp, err := bc.wreckerClient.Patch("/boards/"+boardSpec+"/").
		URLParam("fields", models.BOARD_FIELDS).
		FormParam("name", optionals.Name).
		FormParam("description", optionals.Description).
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*models.Board), nil
}

// Delete deletes an existing board
// Endpoint: [DELETE] /v1/boards/<board_spec:board>/
func (bc *BoardsController) Delete(boardSpec string) error {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = ""
	httpResp, err := bc.wreckerClient.Delete("/boards/" + boardSpec + "/").
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return err
	}

	// OK
	return nil
}
