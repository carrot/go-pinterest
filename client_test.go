package main_test

import (
	pinterest "github.com/carrot/go-pinterest"
	"github.com/carrot/go-pinterest/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
	"time"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

// A test suite for all of our tests against our Client
type ClientTestSuite struct {
	suite.Suite
	client             *pinterest.Client
	unauthorizedClient *pinterest.Client
	timeoutClient      *pinterest.Client
}

// SetupTest sets up our test suite.  All this really does is build us
// a client that is fed our AccessToken.
func (suite *ClientTestSuite) SetupTest() {
	// Create Standard Client
	suite.client = pinterest.NewClient().
		RegisterAccessToken(os.Getenv("PINTEREST_ACCESS_TOKEN"))

	// Create client without any AccessToken
	suite.unauthorizedClient = pinterest.NewClient()

	// Create a timeout client that can never make a request
	// (simulates no network connection)
	suite.timeoutClient = pinterest.NewClient().SetHttpClient(&http.Client{
		Timeout: 1 * time.Nanosecond,
	})
}

// =================================
// ========== Users.Fetch ==========
// =================================

// TestSuccessfulUserFetch tests that a user can be fetched when
// everything was set up properly.
func (suite *ClientTestSuite) TestSuccessfulUserFetch() {
	user, err := suite.client.Users.Fetch("BrandonRRomano")

	// Assume there is no error
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), user.FirstName, "Brandon")
	assert.Equal(suite.T(), user.LastName, "Romano")
}

// TestNotFoundUserFetch tests that a 404 is appropriately thrown
// when fetching a user that does not exist.
func (suite *ClientTestSuite) TestNotFoundUserFetch() {
	// Hopefully nobody ever makes this user
	_, err := suite.client.Users.Fetch("E20450921CE")

	// Assume there is an error
	assert.NotEqual(suite.T(), nil, err)
	assert.Equal(suite.T(), err.Error(), "PinterestError: {\"status_code\":404,\"message\":\"User not found.\"}")

	// Check error type
	if pinterestError, ok := err.(*models.PinterestError); ok {
		// Should be a 404
		assert.Equal(suite.T(), http.StatusNotFound, pinterestError.StatusCode)
	} else {
		// Make this error out, should always be a PinterestError
		assert.Equal(suite.T(), true, false)
	}
}

// TestTimeoutUserFetch tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutUserFetch() {
	_, err := suite.timeoutClient.Users.Fetch("BrandonRRomano")
	assert.NotEqual(suite.T(), nil, err)
}

// TestUnauthorizedUserFetch tests that an error is appropriately thrown
// when the user makes an unauthorized request
func (suite *ClientTestSuite) TestUnauthorizedUserFetch() {
	_, err := suite.unauthorizedClient.Users.Fetch("BrandonRRomano")
	assert.NotEqual(suite.T(), nil, err)

	// Check error type
	if pinterestError, ok := err.(*models.PinterestError); ok {
		// Should be a 401
		assert.Equal(suite.T(), http.StatusUnauthorized, pinterestError.StatusCode)
	} else {
		// Make this error out, should always be a PinterestError
		assert.Equal(suite.T(), true, false)
	}
}

// ==================================
// ========== Boards.Fetch ==========
// ==================================

// TestSuccessfulBoardFetch tests that a board can be fetched when
// everything was set up properly.
func (suite *ClientTestSuite) TestSuccessfulBoardFetch() {
	board, err := suite.client.Boards.Fetch("BrandonRRomano", "go-pinterest")

	// Assume there is no error
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), board.Name, "Go Pinterest!")
	assert.Equal(suite.T(), board.Creator.FirstName, "Brandon")
	assert.Equal(suite.T(), board.Creator.LastName, "Romano")
}

// TestNotFoundBoardFetch tests that a 404 is appropriately thrown
// when fetching a board that does not exist.
func (suite *ClientTestSuite) TestNotFoundBoardFetch() {
	// Fetch board that doesn't exist
	_, err := suite.client.Boards.Fetch("BrandonRRomano", "E20450921CE")

	// Assume there is an error
	assert.NotEqual(suite.T(), nil, err)

	// Check error type
	if pinterestError, ok := err.(*models.PinterestError); ok {
		// Should be a 404
		assert.Equal(suite.T(), http.StatusNotFound, pinterestError.StatusCode)
	} else {
		// Make this error out, should always be a PinterestError
		assert.Equal(suite.T(), true, false)
	}
}

// TestTimeoutBoardFetch tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutBoardFetch() {
	_, err := suite.timeoutClient.Boards.Fetch("BrandonRRomano", "go-pinterest")
	assert.NotEqual(suite.T(), nil, err)
}

// TestUnauthorizedBoardFetch tests that an error is appropriately thrown
// when the user makes an unauthorized request
func (suite *ClientTestSuite) TestUnauthorizedBoardFetch() {
	_, err := suite.unauthorizedClient.Boards.Fetch("BrandonRRomano", "go-pinterest")
	assert.NotEqual(suite.T(), nil, err)

	// Check error type
	if pinterestError, ok := err.(*models.PinterestError); ok {
		// Should be a 401
		assert.Equal(suite.T(), http.StatusUnauthorized, pinterestError.StatusCode)
	} else {
		// Make this error out, should always be a PinterestError
		assert.Equal(suite.T(), true, false)
	}
}
