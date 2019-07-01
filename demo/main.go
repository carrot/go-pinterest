package main

import (
	"fmt"
	pinterest "go-pinterest"
)

const Read_public = "read_public"
const Write_public = "write_public"
const Read_relationships = "read_relationships"
const Write_relationships = "write_relationships"

func main() {
	fmt.Println(getCredentials())
	scopes := []string{Read_public, Read_relationships, Write_relationships, Write_public}
	fmt.Print(getLoginUrl("http://", scopes))
	client := pinterest.NewClient()
	client.OAuth.Token.Re

}

func getLoginUrl(redirectUrl string, scops []string) string {

	return fmt.Sprintf("%v\n", scops)
}
