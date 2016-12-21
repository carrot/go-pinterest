package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// MeFollowingUsersController is the controller that is responsible for all
// /v1/me/following/users/ endpoints in the Pinterest API.
type MeFollowingUsersController struct {
	wreckerClient *wrecker.Wrecker
}

// newMeFollowingUsersController instantiates a new MeFollowingUsersController
func newMeFollowingUsersController(wc *wrecker.Wrecker) *MeFollowingUsersController {
	return &MeFollowingUsersController{
		wreckerClient: wc,
	}
}

// FollowingUsersControllerFetchOptionals is a struct that represents the optional
// parameters for the Fetch method
type FollowingUsersControllerFetchOptionals struct {
	Cursor string
}

// Fetch loads the users that the authorized user follows
// Endpoint: [GET] /v1/me/following/users/
func (c *MeFollowingUsersController) Fetch(optionals *FollowingUsersControllerFetchOptionals) (*[]models.User, *models.Page, error) {
	// Build request
	response := new(models.Response)
	response.Data = &[]models.User{}
	request := c.wreckerClient.Get("/me/following/users/").
		URLParam("fields", models.USER_FIELDS).
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
	return response.Data.(*[]models.User), &response.Page, nil
}

// Create follows a user
// Endpoint: [POST] /v1/me/following/users/
func (c *MeFollowingUsersController) Create(user string) error {
	// Build request
	response := new(models.Response)
	request := c.wreckerClient.Post("/me/following/users/").
		FormParam("user", user).
		Into(response)

	// Execute request
	resp, err := request.Execute()

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

// Delete unfollows a user
// Endpoint: [DELETE] /v1/me/following/users/
func (c *MeFollowingUsersController) Delete(user string) error {
	// Build request
	response := new(models.Response)
	request := c.wreckerClient.Delete("/me/following/users/" + user + "/").
		Into(response)

	// Execute request
	resp, err := request.Execute()

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
