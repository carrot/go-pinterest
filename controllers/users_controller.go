package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/pinterest-go-client/models"
)

type UsersController struct {
	wreckerClient *wrecker.Wrecker
}

func NewUsersController(wc *wrecker.Wrecker) *UsersController {
	return &UsersController{
		wreckerClient: wc,
	}
}

// https://developers.pinterest.com/docs/api/users
func (uc *UsersController) Fetch(user string) (*models.User, error) {
	response := new(models.Response)
	response.Data = new(models.User)

	resp, err := uc.wreckerClient.Get("/users/"+user+"/").
		URLParam("fields", "first_name,last_name,url,account_type,bio,counts,created_at,image,username").
		Into(response).
		Execute()

	// Error from Wrecker
	if err != nil {
		return nil, err
	}

	// Status code
	if resp.StatusCode != 200 {
		return nil, &models.PinterestError{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
			Type:       response.Type,
		}
	}

	// OK
	return response.Data.(*models.User), nil
}
