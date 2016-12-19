package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
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
	// Build request
	response := new(models.Response)
	response.Data = &[]models.Board{}
	request := mbsc.wreckerClient.Get("/me/boards/suggested/").
		URLParam("fields", "id,name,url,counts,creator,description,created_at,image,privacy,reason").
		Into(response)
	if optionals.Count != 0 {
		request.URLParam("count", strconv.Itoa(int(optionals.Count)))
	}
	if optionals.Pin != "" {
		request.URLParam("pin", optionals.Pin)
	}

	// Execute request
	resp, err := request.Execute()

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
