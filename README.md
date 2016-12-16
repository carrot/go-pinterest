# Go Pinterest [![Build Status](https://travis-ci.org/carrot/go-pinterest.svg?branch=master)](https://travis-ci.org/carrot/go-pinterest) [![Coverage Status](https://coveralls.io/repos/github/carrot/go-pinterest/badge.svg?branch=br.coveralls)](https://coveralls.io/github/carrot/go-pinterest?branch=br.coveralls)

> Note: This project is in active development, and will not provide any stability until releases are drafted.  Keep an eye on the [releases](https://github.com/carrot/pinterest-go-client/releases)!

This is a Go client for the Pinterest API.  For more information about the Pinterest API, check out their [getting started](https://developers.pinterest.com/docs/api/overview/) guide.

## Roadmap to v1.0.0

Check out the [v1.0.0 milestone](https://github.com/carrot/go-pinterest/milestone/1)!

Before we hit v1.0.0 you probably shouldn't be using this library. If you decide to use it be sure to lock into specific commit, as there are possibly breaking changes between commits without notice.

## Creating a Client

To create a simple Pinterest client:

```go
import pinterest "github.com/carrot/go-pinterest"

func main() {
    client := pinterest.NewClient()
}
```

To create an authenticated Pinterest client you can just chain the `RegisterAccessToken` method:

```go
import pinterest "github.com/carrot/go-pinterest"

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

For all requests made via this library, there is the possibility of the API throwing an error.

When the API does throw an error, this library makes it very easy to handle it.

If the error came from Pinterest (and not the *http.Client) this library wraps it in a custom error, PinterestError:

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

## Me Endpoints

## Pins Endpoints

## Users Endpoints

## License

[MIT](LICENSE.md) © Carrot Creative
