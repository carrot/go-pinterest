package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/jschwehn/go-pinterest/models"

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
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.Pin{}
	request := mspc.wreckerClient.Get("/me/search/pins/").
		URLParam("fields", models.PIN_FIELDS).
		URLParam("query", query).
		Into(resp)
	if optionals.Cursor != "" {
		request.URLParam("cursor", optionals.Cursor)
	}
	if optionals.Limit != 0 {
		request.URLParam("limit", strconv.Itoa(optionals.Limit))
	}
	httpResp, err := request.Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, nil, err
	}

	// OK
	return resp.Data.(*[]models.Pin), &resp.Page, nil
}
