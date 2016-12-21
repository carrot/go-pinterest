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
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.Pin{}
	request := mpc.wreckerClient.Get("/me/pins/").
		URLParam("fields", models.PIN_FIELDS).
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
	return resp.Data.(*[]models.Pin), &resp.Page, nil
}
