package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

type PinsController struct {
	wreckerClient *wrecker.Wrecker
}

func NewPinsController(wc *wrecker.Wrecker) *PinsController {
	return &PinsController{
		wreckerClient: wc,
	}
}

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
