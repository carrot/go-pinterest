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

// Fetch loads a board from the board's username / boardName.
func (bc *BoardsController) Fetch(username string, boardName string) (*models.Board, error) {
	response := new(models.Response)
	response.Data = new(models.Board)

	resp, err := bc.wreckerClient.Get("/boards/"+username+"/"+boardName).
		URLParam("fields", "id,url,reason,counts,created_at,creator,description,image,privacy,name").
		Into(response).
		Execute()

	// Error from Wrecker
	if err != nil {
		return nil, err
	}

	// Status code
	if resp.StatusCode != 200 {
		return nil, &models.PinterestError{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
		}
	}

	// OK
	return response.Data.(*models.Board), nil
}
