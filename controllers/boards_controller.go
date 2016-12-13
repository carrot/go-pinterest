package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// BoardsController is the controller that is responsible for all
// board specific endpoints in the Pinterest API.
// https://developers.pinterest.com/docs/api/boards/
type BoardsController struct {
	wreckerClient *wrecker.Wrecker
}

// NewBoardsController instantiates a new BoardsController.
func NewBoardsController(wc *wrecker.Wrecker) *BoardsController {
	return &BoardsController{
		wreckerClient: wc,
	}
}

// Fetch loads a board from the board_spec (username/board-slug)
func (bc *BoardsController) Fetch(boardSpec string) (*models.Board, error) {
	// Make request
	response := new(models.Response)
	response.Data = new(models.Board)
	resp, err := bc.wreckerClient.Get("/boards/"+boardSpec).
		URLParam("fields", "id,url,reason,counts,created_at,creator,description,image,privacy,name").
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
	return response.Data.(*models.Board), nil
}

// BoardCreateOptionals is a struct that represents the optional parameters
// that can be passed to the Create endpoint
type BoardCreateOptionals struct {
	Description string
}

// Create makes a new board
func (bc *BoardsController) Create(boardName string, optionals *BoardCreateOptionals) (*models.Board, error){
	// Make request
	response := new(models.Response)
	response.Data = new(models.Board)
	resp, err := bc.wreckerClient.Post("/boards/").
		URLParam("fields", "id,url,reason,counts,created_at,creator,description,image,privacy,name").
		FormParam("name", boardName).
		FormParam("description", optionals.Description).
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
	return response.Data.(*models.Board), nil
}

// BoardUpdateOptionals is a struct that represents the optional parameters
// that can be passed to the Update endpoint
type BoardUpdateOptionals struct {
	Name        string
	Description string
}

// Update updates an existing board
func (bc *BoardsController) Update(boardSpec string, optionals *BoardUpdateOptionals) (*models.Board, error) {
	// Make request
	response := new(models.Response)
	response.Data = new(models.Board)
	resp, err := bc.wreckerClient.Patch("/boards/" + boardSpec + "/").
		URLParam("fields", "id,url,reason,counts,created_at,creator,description,image,privacy,name").
		FormParam("name", optionals.Name).
		FormParam("description", optionals.Description).
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
	return response.Data.(*models.Board), nil
}

// Delete deletes an existing board
func (bc *BoardsController) Delete(boardSpec string) error {
	// Make request
	response := new(models.Response)
	response.Data = ""
	resp, err := bc.wreckerClient.Delete("/boards/" + boardSpec + "/").
		Into(response).
		Execute()

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
