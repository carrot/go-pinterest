package main

import (
	"encoding/json"
	"fmt"
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/pinterest-go-client/controllers"
	"net/http"
	"time"
)

type PinterestClient struct {
	Users         *controllers.UsersController
	wreckerClient *wrecker.Wrecker
}

type PinterestErrorDetails struct {
	HttpStatusCode int
}

func NewClient() *PinterestClient {
	pinterestClient := &PinterestClient{
		wreckerClient: &wrecker.Wrecker{
			BaseURL: "https://api.pinterest.com/v1",
			HttpClient: &http.Client{
				Timeout: 10 * time.Second,
			},
			DefaultContentType: "application/json",
			RequestInterceptor: nil,
		},
	}
	pinterestClient.Users = controllers.NewUsersController(pinterestClient.wreckerClient)
	return pinterestClient
}

func (pc *PinterestClient) RegisterAccessToken(accessToken string) *PinterestClient {
	pc.wreckerClient.RequestInterceptor = func(req *wrecker.Request) error {
		req.URLParam("access_token", accessToken)
		return nil
	}
	return pc
}

func main() {
	user, err := NewClient().
		RegisterAccessToken("").
		Users.Fetch("BrandonRRomano")

	if err != nil {
		fmt.Println(err)
	} else {
		out, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}
}
