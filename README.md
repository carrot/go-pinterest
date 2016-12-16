# Go Pinterest [![Build Status](https://travis-ci.org/carrot/go-pinterest.svg?branch=master)](https://travis-ci.org/carrot/go-pinterest) [![Coverage Status](https://coveralls.io/repos/github/carrot/go-pinterest/badge.svg?branch=br.coveralls)](https://coveralls.io/github/carrot/go-pinterest?branch=br.coveralls)

> Note: This project is in active development, and will not provide any stability until releases are drafted.  Keep an eye on the [releases](https://github.com/carrot/pinterest-go-client/releases)!

This is a Go client for the Pinterest API.  For more information about the Pinterest API, check out their [getting started](https://developers.pinterest.com/docs/api/overview/) guide.

## Roadmap to v1.0.0

Check out the [v1.0.0 milestone](https://github.com/carrot/go-pinterest/milestone/1)!

Before we hit v1.0.0 you probably shouldn't be using this library. If you decide to use it be sure to lock into specific commit, as there are possibly breaking changes between commits without notice.

## Creating a Client

To create a simple Pinterest client:

```go
import(
    pinterest "github.com/carrot/go-pinterest"
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
    pinterest "github.com/carrot/go-pinterest"
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
```

### Return the logged in user's Boards

`[GET] /v1/me/boards/`

```go
```

### Return Board suggestions for the logged in user

`[GET] /v1/me/boards/suggested/`

```go
```

### Return the users that follow the logged in user

`[GET] /v1/me/followers/`

```go
```

### Get the Boards that the logged in user follows

`[GET] /v1/me/following/boards/`

```go
```

### Follow a Board

`[POST] /v1/me/following/boards/`

```go
```

### Unfollow a Board

`[DELETE] /v1/me/following/boards/<board_spec:board>/`

```go
```

### Return the Interests the logged in user follows

`[GET] /v1/me/following/interests/`

```go
```

### Return the users that the logged in user follows

`[GET] /v1/me/following/users/`

```go
```

### Follow a user

`[POST] /v1/me/following/users/`

```go
```

### Unfollow a user

`[DELETE] /v1/me/following/users/<user>/`

```go
```

### Return Pins that the logged in user has liked

`[GET] /v1/me/likes/`

```go
```

### Return the logged in user's Pins

`[GET] /v1/me/pins/`

```go
```

### Search the logged in user's Boards

`[GET] /v1/me/search/boards/`

```go
```

### Search the logged in user's Pins

`[GET] /v1/me/search/pins/`

```go
```

## Pins Endpoints

### Create a Pin

`[POST] /v1/pins/`

```go
```

### Delete a Pin

`[DELETE] /v1/pins/<pin>/`

```go
```

### Edit a Pin's information

`[PATCH] /v1/pins/<pin>/`

```go
```

### Return information about a Pin

`[GET] /v1/pins/<pin>/`

```go
```

## Users Endpoints

### Return a user's information

`[GET] /v1/users/<user>/`

```go
```

## License

[MIT](LICENSE.md) © Carrot Creative
