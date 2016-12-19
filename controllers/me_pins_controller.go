package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// MePinsController is the controller that is responsible for all
// /v1/me/pins/ endpoints in the Pinterest API.
type MePinsController struct {
	wreckerClient *wrecker.Wrecker
}

// newMePinsController instantiates a new MePinsController
func newMePinsController(wc *wrecker.Wrecker) *MePinsController {
	return &MePinsController{
		wreckerClient: wc,
	}
}

// MePinsFetchOptionals is a struct that represents the optional parameters
// that can be passed to the Fetch endpoint
type MePinsFetchOptionals struct {
	Cursor string
}

// Fetch loads all of the logged in user's Pins
// Endpoint: [GET] /v1/me/pins/
func (mpc *MePinsController) Fetch(optionals *MePinsFetchOptionals) (*[]models.Pin, *models.Page, error) {
	// Build request
	response := new(models.Response)
	response.Data = &[]models.Pin{}
	request := mpc.wreckerClient.Get("/me/pins/").
		URLParam("fields", "id,link,note,url,attribution,board,color,counts,created_at,creator,image,media,metadata,original_link").
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
	return response.Data.(*[]models.Pin), &response.Page, nil
}
