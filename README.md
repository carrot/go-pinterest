# Go Pinterest [![Build Status](https://travis-ci.org/carrot/go-pinterest.svg?branch=master)](https://travis-ci.org/carrot/go-pinterest) [![Coverage Status](https://coveralls.io/repos/github/carrot/go-pinterest/badge.svg?branch=br.coveralls)](https://coveralls.io/github/carrot/go-pinterest?branch=br.coveralls)

This is a Go client for the Pinterest API.  For more information about the Pinterest API, check out their [getting started](https://developers.pinterest.com/docs/api/overview/) guide.

## Creating a Client

To create a simple Pinterest client:

```go
import(
    "github.com/carrot/go-pinterest"
    // Below imports aren't needed immediately, but you'll want these soon after
	"github.com/carrot/go-pinterest/controllers"
	"github.com/carrot/go-pinterest/models"
)

func main() {
    client := pinterest.NewClient()
}
```

To create an authenticated Pinterest client you can just chain the `RegisterAccessToken` method:

```go
import(
    "github.com/carrot/go-pinterest"
    // Below imports aren't needed immediately, but you'll want these soon after
	"github.com/carrot/go-pinterest/controllers"
	"github.com/carrot/go-pinterest/models"
)

func main() {
    client := pinterest.NewClient().
        RegisterAccessToken("USERS_ACCESS_TOKEN")
}
```

## Calling API Methods

Calling specific API methods matches 1:1 with the Pinterest API URL schema.

The best way to illustrate this is with a few examples:

```go
// Fetch Board info:
// [GET] /v1/boards/<board_spec:board>/
board, err := client.Boards.Fetch("BrandonRRomano/go-pinterest")

// Unfollow a User:
// [DELETE] /v1/me/following/users/<user>/
err := client.Me.Following.Users.Delete("BrandonRRomano")

// Fetch the Pins on a Board:
// [GET] /v1/boards/<board_spec:board>/pins/
pins, err := client.Boards.Pins("BrandonRRomano/go-pinterest")
```

As you can see, you simply chain through the controllers following the URL in the Pinterest API. If there is a URL with a segment that is a parameter (see Fetch the Pins on a Board in above example), you simply skip that segment in the controller chaining; the parameter will be passed along in the parameters of the method.

Once you're at the final segment, you can call `Create`, `Fetch`, `Update`, or `Delete`, which will call the API's `POST`, `GET`, `PATCH`, or `DELETE` methods respectively.

All required parameters to the Pinterest API's methods will be parameters in the method.  All optional parameters will be stuffed in an `Optionals` object as the last parameter.

## Handling Errors

For all requests made via this library, there is the possibility of the Pinterest API throwing an error.

When the API does throw an error, this library makes it very easy to handle it.

If the error came from Pinterest (and not the [http.Client](https://golang.org/pkg/net/http/#Client)) this library wraps it in a custom error, PinterestError:

```go
type PinterestError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
```

Here's some example usage of this:

```go
// Fetch a board that doesn't exist
_, err := client.Boards.Fetch("BrandonRRomano/E20450921CE")

// Check error type
if pinterestError, ok := err.(*models.PinterestError); ok {
    if pinterestError.StatusCode == 404 {
        // Do something to handle it!
    } else {
        // something else!
    }
} else {
    // Was an error thrown by *http.Client!
    // Something is probably wrong with your network
}
```

## Boards Endpoints

### Create a Board

`[POST] /v1/boards/`

```go
board, err := client.Boards.Create(
    "My Test Board",
    &controllers.BoardCreateOptionals{
        Description: "This is a test!",
    },
)
```

### Delete a Board

`[DELETE] /v1/boards/<board_spec:board>/`

```go
err := client.Boards.Delete("BrandonRRomano/go-pinterest-test")
```

### Edit a Board

`[PATCH] /v1/boards/<board_spec:board>/`

```go
board, err := client.Boards.Update(
    "BrandonRRomano/go-pinterest-test",
    &controllers.BoardUpdateOptionals{
        Name:        "Some new name",
        Description: "Some new description",
    },
)
```

### Retrieve information about a Board

`[GET] /v1/boards/<board_spec:board>/`

```go
board, err := client.Boards.Fetch("BrandonRRomano/go-pinterest")
```

### Retrieve the Pins on a Board

`[GET] /v1/boards/<board_spec:board>/pins/`

```go
pins, err := client.Boards.Pins.Fetch(
    "BrandonRRomano/go-pinterest",
    &controllers.BoardPinsFetchOptionals{
        Cursor: "some-cursor-from-pinterest",
    },
)
```

## Me Endpoints

### Return the logged in user's information

`[GET] /v1/me/`

```go
user, err := client.Me.Fetch()
```

### Return the logged in user's Boards

`[GET] /v1/me/boards/`

```go
boards, err := client.Me.Boards.Fetch()
```

### Return Board suggestions for the logged in user

`[GET] /v1/me/boards/suggested/`

```go
boards, err := client.Me.Boards.Suggested.Fetch(
    &controllers.MeBoardsSuggestedFetchOptionals{
        Count: 10,
        Pin: "some-pin-id",
    },
)
```

### Return the users that follow the logged in user

`[GET] /v1/me/followers/`

```go
users, page, err := client.Me.Followers.Fetch(
    &controllers.MeFollowersFetchOptionals{
        Cursor: "some-cursor",
    },
)
```

### Get the Boards that the logged in user follows

`[GET] /v1/me/following/boards/`

```go
boards, page, err := client.Me.Following.Boards.Fetch(
    &controllers.MeFollowingBoardsFetchOptionals{
        Cursor: "some-cursor",
    },
)
```

### Follow a Board

`[POST] /v1/me/following/boards/`

```go
err := client.Me.Following.Boards.Create("pinterest/pinterest-100-for-2017")
```

### Unfollow a Board

`[DELETE] /v1/me/following/boards/<board_spec:board>/`

```go
err := client.Me.Following.Boards.Delete("pinterest/pinterest-100-for-2017")
```

### Return the Interests the logged in user follows

`[GET] /v1/me/following/interests/`

```go
interests, page, err := client.Me.Following.Interests.Fetch(
    &controllers.MeFollowingInterestsFetchOptionals{
        Cursor: "some-cursor",
    },
)
```

### Return the users that the logged in user follows

`[GET] /v1/me/following/users/`

```go
users, page, err := client.Me.Following.Users.Fetch(
    &controllers.FollowingUsersControllerFetchOptionals{
        Cursor: "some-cursor",
    },
)
```

### Follow a user

`[POST] /v1/me/following/users/`

```go
err := client.Me.Following.Users.Create("hhsnopek")
```

### Unfollow a user

`[DELETE] /v1/me/following/users/<user>/`

```go
err := client.Me.Following.Users.Delete("hhsnopek")
```

### Return Pins that the logged in user has liked

`[GET] /v1/me/likes/`

```go
pins, page, err := client.Me.Likes.Fetch(
    &controllers.MeLikesFetchOptionals{
        Cursor: "some-cursor",
    },
)
```

### Return the logged in user's Pins

`[GET] /v1/me/pins/`

```go
pins, page, err := client.Me.Pins.Fetch(
    &controllers.MePinsFetchOptionals{
        Cursor: "some-cursor",
    },
)
```

### Search the logged in user's Boards

`[GET] /v1/me/search/boards/`

```go
// Load first page
boards, page, err := client.Me.Search.Boards.Fetch(
    "Go Pinterest",
    &controllers.MeSearchBoardsFetchOptionals{
        Cursor: "some-cursor",
        Limit: 1,
    },
)
```

### Search the logged in user's Pins

`[GET] /v1/me/search/pins/`

```go
pins, page, err := client.Me.Search.Pins.Fetch(
    "Go Gopher",
    &controllers.MeSearchPinsFetchOptionals{
        Cursor: "some-cursor",
        Limit: 1,
    },
)
```

## Pins Endpoints

### Create a Pin

`[POST] /v1/pins/`

```go
pin, err := client.Pins.Create(
    "BrandonRRomano/go-pinterest-2",
    "This is a cat",
    &controllers.PinCreateOptionals{
        Link:     "http://www.google.com/",
        ImageUrl: "http://i.imgur.com/1olmVpO.jpg",
    },
)
```

### Delete a Pin

`[DELETE] /v1/pins/<pin>/`

```go
err := client.Pins.Delete("some-pin-id")
```

### Edit a Pin's information

`[PATCH] /v1/pins/<pin>/`

```go
pin, err := client.Pins.Update(
    "some-pin-id",
    &controllers.PinUpdateOptionals{
        Board: "BrandonRRomano/go-pinterest",
        Note:  "This is a new cat",
        Link:  "http://www.facebook.com/",
    },
)
```

### Return information about a Pin

`[GET] /v1/pins/<pin>/`

```go
pin, err := client.Pins.Fetch("some-pin-id")
```

## Users Endpoints

### Return a user's information

`[GET] /v1/users/<user>/`

```go
user, err := client.Users.Fetch("BrandonRRomano")
```

## License

[MIT](LICENSE.md) © Carrot Creative
