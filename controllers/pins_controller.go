package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// PinsController is the controller that is responsible for all
// /v1/pins/ endpoints in the Pinterest API.
type PinsController struct {
	wreckerClient *wrecker.Wrecker
}

// NewPinsController instantiates a new PinsController
func NewPinsController(wc *wrecker.Wrecker) *PinsController {
	return &PinsController{
		wreckerClient: wc,
	}
}

// Fetch loads a pin from the pin id
// Endpoint: [GET] /v1/pins/<pin>/
func (pc *PinsController) Fetch(pinId string) (*models.Pin, error) {
	// Make request
	response := new(models.Response)
	response.Data = new(models.Pin)

	resp, err := pc.wreckerClient.Get("/pins/"+pinId+"/").
		URLParam("fields", "id,link,note,url,attribution,board,color,counts,created_at,creator,image,media,metadata,original_link").
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
	return response.Data.(*models.Pin), nil
}

// PinCreateOptionals is a struct that represents the optional parameters
// that can be passed to the Create endpoint
type PinCreateOptionals struct {
	Link     string
	ImageUrl string
	// TODO
	// Image       string
	// ImageBase64 string
}

// Create creates a new pin
// Endpoint: [POST] /v1/pins/
func (pc *PinsController) Create(boardSpec string, note string, optionals *PinCreateOptionals) (*models.Pin, error) {
	// Build Request
	response := new(models.Response)
	response.Data = new(models.Pin)
	request := pc.wreckerClient.Post("/pins/").
		URLParam("fields", "id,link,note,url,attribution,board,color,counts,created_at,creator,image,media,metadata,original_link").
		FormParam("board", boardSpec).
		FormParam("note", note).
		Into(response)
	if optionals.Link != "" {
		request.FormParam("link", optionals.Link)
	}
	if optionals.ImageUrl != "" {
		request.FormParam("image_url", optionals.ImageUrl)
	}

	// Execute Request
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
	return response.Data.(*models.Pin), nil
}

// PinUpdateOptionals is a struct that represents the optional parameters
// that can be passed to the Update endpoint
type PinUpdateOptionals struct {
	Board string
	Note  string
	Link  string
}

// Update updates an existing pin
// Endpoint: [PATCH] /v1/pins/<pin>/
func (pc *PinsController) Update(pinId string, optionals *PinUpdateOptionals) (*models.Pin, error) {
	// Build request
	response := new(models.Response)
	response.Data = new(models.Pin)
	request := pc.wreckerClient.Patch("/pins/"+pinId+"/").
		URLParam("fields", "id,link,note,url,attribution,board,color,counts,created_at,creator,image,media,metadata,original_link").
		Into(response)
	if optionals.Board != "" {
		request.FormParam("board", optionals.Board)
	}
	if optionals.Note != "" {
		request.FormParam("note", optionals.Note)
	}
	if optionals.Link != "" {
		request.FormParam("link", optionals.Link)
	}

	// Execute Request
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
	return response.Data.(*models.Pin), nil
}

// Delete deletes an existing pin
// Endpoint: [DELETE] /v1/pins/<pin>/
func (pc *PinsController) Delete(pinId string) error {
	// Execute Request
	response := new(models.Response)
	resp, err := pc.wreckerClient.Delete("/pins/" + pinId + "/").Execute()

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
