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

## License

[MIT](LICENSE.md) © Carrot Creative
