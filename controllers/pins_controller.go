package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// TODO doc
type PinsController struct {
	wreckerClient *wrecker.Wrecker
}

// TODO doc
func NewPinsController(wc *wrecker.Wrecker) *PinsController {
	return &PinsController{
		wreckerClient: wc,
	}
}

// TODO doc
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

type CreatePinOptionals struct {
	Link        string
	Image       string // TODO: file
	ImageUrl    string
	ImageBase64 string // TODO: better structure
}

// TODO doc
func (pc *PinsController) Create(pinId string, note string, optionals *CreatePinOptionals) (*models.Pin, error) {
	return nil, nil
}

// TODO doc
func (pc *PinsController) Update(pinId string) (*models.Pin, error) {
	return nil, nil
}

// TODO doc
func (pc *PinsController) Delete(pinId string) error {
	// TODO impl
	return nil, nil
}
