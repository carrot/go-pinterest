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
	// Build + execute request
	resp := new(models.Response)
	resp.Data = &[]models.User{}
	request := c.wreckerClient.Get("/me/following/users/").
		URLParam("fields", models.USER_FIELDS).
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
	return resp.Data.(*[]models.User), &resp.Page, nil
}

// Create follows a user
// Endpoint: [POST] /v1/me/following/users/
func (c *MeFollowingUsersController) Create(user string) error {
	// Build + execute request
	resp := new(models.Response)
	httpResp, err := c.wreckerClient.Post("/me/following/users/").
		FormParam("user", user).
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return err
	}

	// OK
	return nil
}

// Delete unfollows a user
// Endpoint: [DELETE] /v1/me/following/users/
func (c *MeFollowingUsersController) Delete(user string) error {
	// Build + execute request
	resp := new(models.Response)
	httpResp, err := c.wreckerClient.Delete("/me/following/users/" + user + "/").
		Into(resp).
		Execute()

	// Check Error
	if err = models.WrapPinterestError(httpResp, resp, err); err != nil {
		return err
	}

	// OK
	return nil
}
