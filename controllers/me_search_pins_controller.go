package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
	"strconv"
)

// MeSearchPinsController is the controller that is responsible for all
// /v1/me/search/pins/ endpoints in the Pinterest API.
type MeSearchPinsController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeSearchPinsController instantiates a new MeSearchPinsController
func newMeSearchPinsController(wc *wrecker.Wrecker) *MeSearchPinsController {
	return &MeSearchPinsController{
		wreckerClient: wc,
	}
}

// MeSearchPinsFetchOptionals is a struct that represents the optional parameters
// that can be passed to the Fetch endpoint
type MeSearchPinsFetchOptionals struct {
	Cursor string
	Limit  int
}

// Fetch searches the logged in user's Pins
// Endpoint: [GET] /v1/me/search/pins/
func (mspc *MeSearchPinsController) Fetch(query string, optionals *MeSearchPinsFetchOptionals) (*[]models.Pin, *models.Page, error) {
	// Build request
	response := new(models.Response)
	response.Data = &[]models.Pin{}
	request := mspc.wreckerClient.Get("/me/search/pins/").
		URLParam("fields", models.PIN_FIELDS).
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
	return response.Data.(*[]models.Pin), &response.Page, nil
}
