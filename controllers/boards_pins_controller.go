package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/jschwehn/go-pinterest/models"
)

// BoardsPinsController is the controller that is responsible for all
// /v1/boards/<board_spec:board>/pins/ endpoints in the Pinterest API.
type BoardsPinsController struct {
	wreckerClient *wrecker.Wrecker
}

// newBoardsPinsController instantiates a new BoardsPinsController.
func newBoardsPinsController(wc *wrecker.Wrecker) *BoardsPinsController {
	return &BoardsPinsController{
		wreckerClient: wc,
	}
}

// BoardsPinsFetchOptionals is a struct that represents the optional parameters
// that can be passed to the Fetch endpoint
type BoardsPinsFetchOptionals struct {
	Cursor string
}

// Fetch loads a board from the board_spec (username/board-slug)
// Endpoint: [GET] /v1/boards/<board_spec:board>/pins/
func (bpc *BoardsPinsController) Fetch(boardSpec string, optionals *BoardsPinsFetchOptionals) (*[]models.Pin, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.Pin{}
	httpResp, err := bpc.wreckerClient.Get("/boards/"+boardSpec+"/pins/").
		URLParam("fields", models.PIN_FIELDS).
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*[]models.Pin), nil
}
