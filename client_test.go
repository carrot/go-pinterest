package main_test

import (
	pinterest "github.com/carrot/go-pinterest"
	"github.com/carrot/go-pinterest/controllers"
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
	board, err := suite.client.Boards.Fetch("BrandonRRomano/go-pinterest")

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
	_, err := suite.client.Boards.Fetch("BrandonRRomano/E20450921CE")

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
	_, err := suite.timeoutClient.Boards.Fetch("BrandonRRomano/go-pinterest")
	assert.NotEqual(suite.T(), nil, err)
}

// TestUnauthorizedBoardFetch tests that an error is appropriately thrown
// when the user makes an unauthorized request
func (suite *ClientTestSuite) TestUnauthorizedBoardFetch() {
	_, err := suite.unauthorizedClient.Boards.Fetch("BrandonRRomano/go-pinterest")
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
// ========== Boards.Create =========
// ==================================

// TestBadRequestBoardCreate tests that an error is appropriately thrown
// when a board is attempted to be created without a name
func (suite *ClientTestSuite) TestBadRequestBoardCreate() {
	_, err := suite.client.Boards.Create("", &controllers.BoardCreateOptionals{})
	assert.NotEqual(suite.T(), nil, err)

	// Check error type
	if pinterestError, ok := err.(*models.PinterestError); ok {
		// Should be a 400
		assert.Equal(suite.T(), http.StatusBadRequest, pinterestError.StatusCode)
	} else {
		// Make this error out, should always be a PinterestError
		assert.Equal(suite.T(), true, false)
	}
}

// TestTimeoutBoardCreate tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutBoardCreate() {
	_, err := suite.timeoutClient.Boards.Create("Some Board", &controllers.BoardCreateOptionals{})
	assert.NotEqual(suite.T(), nil, err)
}

// TestUnauthorizedBoardCreate tests that an error is appropriately thrown
// when the user makes an unauthorized request
func (suite *ClientTestSuite) TestUnauthorizedBoardCreate() {
	_, err := suite.unauthorizedClient.Boards.Create("Some Board", &controllers.BoardCreateOptionals{})
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
// ========== Boards.Update =========
// ==================================

// TestTimeoutBoardUpdate tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutBoardUpdate() {
	_, err := suite.timeoutClient.Boards.Update("brandonrromano/go-pinterest-test", &controllers.BoardUpdateOptionals{})
	assert.NotEqual(suite.T(), nil, err)
}

// TestUnauthorizedBoardUpdate tests that an error is appropriately thrown
// when the user makes an unauthorized request
func (suite *ClientTestSuite) TestUnauthorizedBoardUpdate() {
	_, err := suite.unauthorizedClient.Boards.Update("brandonrromano/go-pinterest-test", &controllers.BoardUpdateOptionals{})
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

// TestForbiddenBoardUpdate tests that you may not update boards
// that do not belong to you.
func (suite *ClientTestSuite) TestForbiddenBoardUpdate() {
	// Try to update Pinterests board!
	_, err := suite.client.Boards.Update("pinterest/pinterest-100-for-2017",
		&controllers.BoardUpdateOptionals{
			Name: "Hello World!",
		},
	)
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

// TestNotFoundBoardUpdate tests that you may not update boards
// that do not exist.
func (suite *ClientTestSuite) TestNotFoundBoardUpdate() {
	// Try to update Pinterests board!
	_, err := suite.client.Boards.Update("BrandonRRomano/E20450921CE",
		&controllers.BoardUpdateOptionals{
			Name: "Hello World!",
		},
	)
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

// ==================================
// ========== Boards.Delete =========
// ==================================

// TestUnauthorizedBoardDelete tests that an error is appropriately thrown
// when the user makes an unauthorized request
func (suite *ClientTestSuite) TestUnauthorizedBoardDelete() {
	err := suite.unauthorizedClient.Boards.Delete("brandonrromano/go-pinterest")
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

// TestTimeoutBoardDelete tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutBoardDelete() {
	err := suite.timeoutClient.Boards.Delete("brandonrromano/go-pinterest")
	assert.NotEqual(suite.T(), nil, err)
}

// TestNotFoundBoardDelete tests that an error is appropriately thrown
// when trying to delete a board that does not exist
func (suite *ClientTestSuite) TestNotFoundBoardDelete() {
	// Try to update Pinterests board!
	err := suite.client.Boards.Delete("BrandonRRomano/E20450921CE")
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

// TestNotFoundBoardDelete tests that an error is appropriately thrown
// when trying to delete a board that does not belong to the user
func (suite *ClientTestSuite) TestForbiddenBoardDelete() {
	// Try to update Pinterests board!
	err := suite.client.Boards.Delete("pinterest/pinterest-100-for-2017")
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

// =====================================================
// ========== Boards.Create / Update / Delete ==========
// =====================================================

// TestSuccessfulBoardCUD tests the successful flow of creating a board,
// updating a board, and then deleting that board.  These are all done together
// as there is no guarantee of order
func (suite *ClientTestSuite) TestSuccessfulBoardCUD() {
	// Creating the Board
	board, err := suite.client.Boards.Create("Go Pinterest Test",
		&controllers.BoardCreateOptionals{
			Description: "Go Pinterest Test!",
		},
	)

	// Assume there is no error / test result
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), board.Name, "Go Pinterest Test")
	assert.Equal(suite.T(), board.Description, "Go Pinterest Test!")
	assert.Equal(suite.T(), board.Creator.FirstName, "Brandon")
	assert.Equal(suite.T(), board.Creator.LastName, "Romano")

	// Updating the Board
	board, err = suite.client.Boards.Update("brandonrromano/go-pinterest-test",
		&controllers.BoardUpdateOptionals{
			Name:        "Go Pinterest Test2",
			Description: "Go Pinterest Test2!",
		},
	)

	// Assume there is no error / test result
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), board.Name, "Go Pinterest Test2")
	assert.Equal(suite.T(), board.Description, "Go Pinterest Test2!")

	// Deleting the board
	err = suite.client.Boards.Delete("brandonrromano/go-pinterest-test2")
	assert.Equal(suite.T(), nil, err)
}

// =====================================
// ========== BoardPins.Fetch ==========
// =====================================

// TestSuccessfulBoardPinsFetch tests that a boards pins can be
// fetched when everything was set up properly.
func (suite *ClientTestSuite) TestSuccessfulBoardPinsFetch() {
	pins, err := suite.client.Boards.Pins.Fetch("BrandonRRomano/go-pinterest", &controllers.BoardsPinsFetchOptionals{})

	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), len(*pins), 3)

	firstPin := (*pins)[0]
	assert.Equal(suite.T(), firstPin.Note, "Go Gopher Toy by Sean Tasdemir — Kickstarter")
	assert.Equal(suite.T(), firstPin.Board.Name, "Go Pinterest!")
	assert.Equal(suite.T(), firstPin.Creator.FirstName, "Brandon")
	assert.Equal(suite.T(), firstPin.Creator.FirstName, "Brandon")
}

// TestNotFoundBoardPinsFetch tests that a 404 is thrown
// when trying to access the pins of a board that does not exist
func (suite *ClientTestSuite) TestNotFoundBoardPinsFetch() {
	_, err := suite.client.Boards.Pins.Fetch(
		"BrandonRRomano/E20450921CE",
		&controllers.BoardsPinsFetchOptionals{},
	)
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

// TestTimeoutBoardPinsFetch tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutBoardPinsFetch() {
	_, err := suite.timeoutClient.Boards.Pins.Fetch(
		"BrandonRRomano/go-pinterest",
		&controllers.BoardsPinsFetchOptionals{},
	)
	assert.NotEqual(suite.T(), nil, err)
}

// TestUnauthorizedBoardPinsFetch tests that an error is appropriately thrown
// when the user makes an unauthorized request
func (suite *ClientTestSuite) TestUnauthorizedBoardPinsFetch() {
	_, err := suite.unauthorizedClient.Boards.Pins.Fetch(
		"BrandonRRomano/go-pinterest",
		&controllers.BoardsPinsFetchOptionals{},
	)
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

// ================================
// ========== Pins.Fetch ==========
// ================================

// TestSuccessfulPinsFetch tests that Pins can be fetched when
// everything is set up appropriately
func (suite *ClientTestSuite) TestSuccessfulPinsFetch() {
	pin, err := suite.client.Pins.Fetch("192880796521721688")

	// Assume no error
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), "Go Pinterest!", pin.Board.Name)
	assert.Equal(suite.T(), "The Go Gopher - The Go Blog", pin.Note)
	assert.Equal(suite.T(), "Brandon", pin.Creator.FirstName)
	assert.Equal(suite.T(), "Romano", pin.Creator.LastName)
}

// TestNotFoundPinsFetch tests that a 404 is thrown when we try
// to call Fetch on a pin that doesn't exist
func (suite *ClientTestSuite) TestNotFoundPinsFetch() {
	_, err := suite.client.Pins.Fetch("9999999991234")

	// Check that there's an error
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

// TestTimeoutPinsFetch tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutPinsFetch() {
	_, err := suite.timeoutClient.Pins.Fetch("192880796521721688")
	assert.NotEqual(suite.T(), nil, err)
}

// TestUnauthorizedPinsFetch tests that an error is appropriately thrown
// when the user makes an unauthorized request
func (suite *ClientTestSuite) TestUnauthorizedPinsFetch() {
	_, err := suite.unauthorizedClient.Pins.Fetch("192880796521721688")
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

// ===================================================
// ========== Pins.Create / Update / Delete ==========
// ===================================================

// TestSuccessfulPinCUD tests that a pin can be created, updated,
// and deleted when called appropriately.
func (suite *ClientTestSuite) TestSuccessfulPinCUD() {
	// Create a Pin
	pin, err := suite.client.Pins.Create(
		"brandonrromano/go-pinterest-2",
		"This is a cat",
		&controllers.PinCreateOptionals{
			Link:     "http://www.google.com/",
			ImageUrl: "http://i.imgur.com/1olmVpO.jpg",
		},
	)
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), "This is a cat", pin.Note)
	assert.Equal(suite.T(), "http://www.google.com/", pin.OriginalLink)
	assert.NotEqual(suite.T(), "", pin.Image.Original.Url)

	// Update the Pin
	pin, err = suite.client.Pins.Update(
		pin.Id,
		&controllers.PinUpdateOptionals{
			Board: "brandonrromano/go-pinterest",
			Note:  "This is a new cat",
			Link:  "http://www.facebook.com/",
		},
	)
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), "This is a new cat", pin.Note)
	assert.Equal(suite.T(), "http://www.facebook.com/", pin.OriginalLink)
	assert.Equal(suite.T(), "Go Pinterest!", pin.Board.Name)

	// Delete the Pin
	err = suite.client.Pins.Delete(pin.Id)
	assert.Equal(suite.T(), nil, err)
}

// =================================
// ========== Pins.Create ==========
// =================================

// TestUnauthorizedPinCreate tests that a 401 error is thrown when
// a user is unauthorized and tries to update a pin
func (suite *ClientTestSuite) TestUnauthorizedPinCreate() {
	_, err := suite.unauthorizedClient.Pins.Create(
		"brandonrromano/go-pinterest",
		"Some note, wow",
		&controllers.PinCreateOptionals{
			ImageUrl: "http://i.imgur.com/1olmVpO.jpg",
		},
	)
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

// TestForbiddenPinCreate tests that a 401 error is thrown
// when a user is trying to update a pin that doesn't belong to them
func (suite *ClientTestSuite) TestForbiddenPinCreate() {
	_, err := suite.client.Pins.Create(
		"pinterest/pinterest-100-for-2017",
		"Some note, wow",
		&controllers.PinCreateOptionals{
			ImageUrl: "http://i.imgur.com/1olmVpO.jpg",
		},
	)
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

// TestTimeoutPinsCreate tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutPinsCreate() {
	_, err := suite.timeoutClient.Pins.Create(
		"pinterest/pinterest-100-for-2017",
		"Some note, wow",
		&controllers.PinCreateOptionals{
			ImageUrl: "http://i.imgur.com/1olmVpO.jpg",
		},
	)
	assert.NotEqual(suite.T(), nil, err)
}

// =================================
// ========== Pins.Update ==========
// =================================

// TestUnauthorizedPinUpdate tests that a 401 is thrown when a
// user isn't authorized and tries to update a pin
func (suite *ClientTestSuite) TestUnauthorizedPinUpdate() {
	_, err := suite.unauthorizedClient.Pins.Update(
		"192880796521721688",
		&controllers.PinUpdateOptionals{
			Note: "Hello Update!",
		},
	)
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

// TestUnauthorizedPinUpdate tests that a 401 is thrown when a
// user tries to update a pin that doesn't belong to them
func (suite *ClientTestSuite) TestForbiddenPinUpdate() {
	_, err := suite.client.Pins.Update(
		"424605071105031783",
		&controllers.PinUpdateOptionals{
			Note: "Hello Update!",
		},
	)
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

// TestTimeoutPinUpdate tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutPinUpdate() {
	_, err := suite.timeoutClient.Pins.Update(
		"192880796521721688",
		&controllers.PinUpdateOptionals{
			Note: "Hello Update!",
		},
	)
	assert.NotEqual(suite.T(), nil, err)
}

// =================================
// ========== Pins.Delete ==========
// =================================

// TestUnauthorizedPinDelete tests that a 401 is thrown when a user tries to delete
// a pin when they aren't authorized
func (suite *ClientTestSuite) TestUnauthorizedPinDelete() {
	err := suite.unauthorizedClient.Pins.Delete("192880796521721688")
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

// TestForbiddenPinDelete tests that a 401 is thrown when a user tries to delete
// a pin that doesn't belong to them.
func (suite *ClientTestSuite) TestForbiddenPinDelete() {
	err := suite.client.Pins.Delete("424605071105031783")
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

// TestTimeoutPinDelete tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutPinDelete() {
	err := suite.timeoutClient.Pins.Delete("192880796521721688")
	assert.NotEqual(suite.T(), nil, err)
}

// ==============================
// ========== Me.Fetch ==========
// ==============================

// TestSuccessfulMeFetch tests that a the logged in user can be
// fetched when everything was set up properly.
func (suite *ClientTestSuite) TestSuccessfulMeFetch() {
	user, err := suite.client.Me.Fetch()

	// Assume there is no error
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), user.FirstName, "Brandon")
	assert.Equal(suite.T(), user.LastName, "Romano")
}

// TestUnauthorizedMeFetch tests that a 401 is thrown
// when an unauthorized user tries to call a /me endpoint
func (suite *ClientTestSuite) TestUnauthorizedMeFetch() {
	_, err := suite.unauthorizedClient.Me.Fetch()

	// Check error type
	if pinterestError, ok := err.(*models.PinterestError); ok {
		// Should be a 401
		assert.Equal(suite.T(), http.StatusUnauthorized, pinterestError.StatusCode)
	} else {
		// Make this error out, should always be a PinterestError
		assert.Equal(suite.T(), true, false)
	}
}

// TestTimeoutMeFetch tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutMeFetch() {
	_, err := suite.timeoutClient.Me.Fetch()
	assert.NotEqual(suite.T(), nil, err)
}

// =====================================
// ========== Me.Boards.Fetch ==========
// =====================================

// TestSuccessfulMeBoardsFetch tests that the logged in user
// can fetch their boards.
func (suite *ClientTestSuite) TestSuccessfulMeBoardsFetch() {
	boards, err := suite.client.Me.Boards.Fetch()

	// Assume there is no error
	assert.Equal(suite.T(), nil, err)
	assert.True(suite.T(), (len(*boards) > 0))
	assert.Equal(suite.T(), (*boards)[0].Creator.FirstName, "Brandon")
	assert.Equal(suite.T(), (*boards)[0].Creator.LastName, "Romano")
}

// TestUnauthorizedMeBoardsFetch tests that a 401 is thrown
// when an unauthorized user tries to call a /me endpoint
func (suite *ClientTestSuite) TestUnauthorizedMeBoardsFetch() {
	_, err := suite.unauthorizedClient.Me.Boards.Fetch()

	// Check error type
	if pinterestError, ok := err.(*models.PinterestError); ok {
		// Should be a 401
		assert.Equal(suite.T(), http.StatusUnauthorized, pinterestError.StatusCode)
	} else {
		// Make this error out, should always be a PinterestError
		assert.Equal(suite.T(), true, false)
	}
}

// TestTimeoutMeBoardsFetch tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutMeBoardsFetch() {
	_, err := suite.timeoutClient.Me.Boards.Fetch()
	assert.NotEqual(suite.T(), nil, err)
}

// ===============================================
// ========== Me.Boards.Suggested.Fetch ==========
// ===============================================

// TestSuccessfulMeBoardsSuggestedFetch tests that the logged in user
// can fetch suggested boards.
func (suite *ClientTestSuite) TestSuccessfulMeBoardsSuggestedFetch() {
	// Test simple Fetch
	boards, err := suite.client.Me.Boards.Suggested.Fetch(
		&controllers.MeBoardsSuggestedFetchOptionals{},
	)
	assert.Equal(suite.T(), nil, err)

	//  Test fetch w/ Count
	boards, err = suite.client.Me.Boards.Suggested.Fetch(
		&controllers.MeBoardsSuggestedFetchOptionals{
			Count: 1,
		},
	)
	assert.Equal(suite.T(), nil, err)
	assert.True(suite.T(), (len(*boards) == 1))
}

// TestUnauthorizedMeBoardsSuggestedFetch tests that a 401 is thrown
// when an unauthorized user tries to call a /me endpoint
func (suite *ClientTestSuite) TestUnauthorizedMeBoardsSuggestedFetch() {
	_, err := suite.unauthorizedClient.Me.Boards.Suggested.Fetch(
		&controllers.MeBoardsSuggestedFetchOptionals{},
	)

	// Check error type
	if pinterestError, ok := err.(*models.PinterestError); ok {
		// Should be a 401
		assert.Equal(suite.T(), http.StatusUnauthorized, pinterestError.StatusCode)
	} else {
		// Make this error out, should always be a PinterestError
		assert.Equal(suite.T(), true, false)
	}
}

// TestTimeoutMeBoardsSuggestedFetch tests that an error is appropriately thrown
// when a network timeout occurs
func (suite *ClientTestSuite) TestTimeoutMeBoardsSuggestedFetch() {
	_, err := suite.timeoutClient.Me.Boards.Suggested.Fetch(
		&controllers.MeBoardsSuggestedFetchOptionals{},
	)
	assert.NotEqual(suite.T(), nil, err)
}
