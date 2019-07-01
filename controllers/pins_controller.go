package controllers

import (
	"bufio"
	"encoding/base64"
	"github.com/BrandonRomano/wrecker"
	"github.com/jschwehn/go-pinterest/models"
	"os"
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
	// Build + execute request
	resp := new(models.Response)
	resp.Data = new(models.Pin)
	httpResp, err := pc.wreckerClient.Get("/pins/"+pinId+"/").
		URLParam("fields", models.PIN_FIELDS).
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*models.Pin), nil
}

// PinCreateOptionals is a struct that represents the optional parameters
// that can be passed to the Create endpoint
type PinCreateOptionals struct {
	Link     string
	ImageUrl string
	Image    *os.File
}

// Create creates a new pin
// Endpoint: [POST] /v1/pins/
func (pc *PinsController) Create(boardSpec string, note string, optionals *PinCreateOptionals) (*models.Pin, error) {
	// Build + execute request
	resp := new(models.Response)
	resp.Data = new(models.Pin)
	request := pc.wreckerClient.Post("/pins/").
		URLParam("fields", models.PIN_FIELDS).
		FormParam("board", boardSpec).
		FormParam("note", note).
		Into(resp)
	if optionals.Link != "" {
		request.FormParam("link", optionals.Link)
	}
	// Handle Image
	if optionals.ImageUrl != "" {
		request.FormParam("image_url", optionals.ImageUrl)
	} else if fileInfo, err := optionals.Image.Stat(); err == nil {
		// Create a new buffer based on file size
		var size = fileInfo.Size()
		buf := make([]byte, size)

		// Read file content into buffer
		fReader := bufio.NewReader(optionals.Image)
		fReader.Read(buf)

		// Convert the buffer bytes to base64 string
		imageBase64Str := base64.StdEncoding.EncodeToString(buf)
		request.FormParam("image_base64", imageBase64Str)
	}
	httpResp, err := request.Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*models.Pin), nil
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
	// Build + execute request
	resp := new(models.Response)
	resp.Data = new(models.Pin)
	request := pc.wreckerClient.Patch("/pins/"+pinId+"/").
		URLParam("fields", models.PIN_FIELDS).
		Into(resp)
	if optionals.Board != "" {
		request.FormParam("board", optionals.Board)
	}
	if optionals.Note != "" {
		request.FormParam("note", optionals.Note)
	}
	if optionals.Link != "" {
		request.FormParam("link", optionals.Link)
	}
	httpResp, err := request.Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return nil, err
	}

	// OK
	return resp.Data.(*models.Pin), nil
}

// Delete deletes an existing pin
// Endpoint: [DELETE] /v1/pins/<pin>/
func (pc *PinsController) Delete(pinId string) error {
	// Execute Request
	resp := new(models.Response)
	httpResp, err := pc.wreckerClient.Delete("/pins/" + pinId + "/").Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return err
	}

	// OK
	return nil
}
