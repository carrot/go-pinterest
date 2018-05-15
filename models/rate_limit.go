package models

import (
	"net/http"
	"strconv"
	"fmt"
)

type TypeRatelimit struct {
	Remaining int
	Limit     int
	Refresh   int
}

func GetLimit(httpResp *http.Response, key string) int {
	if val, ok := httpResp.Header[key]; ok {
		rateLimit, err := strconv.Atoi(val[0])
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return rateLimit
	}
	return 0
}

func GetRatelimit(httpResp *http.Response) TypeRatelimit {
	return TypeRatelimit {
		Remaining: GetLimit(httpResp, "X-Ratelimit-Remaining"),
		Limit: GetLimit(httpResp, "X-Ratelimit-Limit"),
		Refresh: GetLimit(httpResp, "X-Ratelimit-Refresh"),
	}
}

