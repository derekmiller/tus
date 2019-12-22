package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

const twitterURL = "https://twitter.com/"

type DesiredUsernames struct {
	Usernames []string `json:"usernames"`
}

type AvailableUsernames struct {
	Usernames []string `json:"usernames"`
}

func handler(desiredUsernames DesiredUsernames) (AvailableUsernames, error) {
	var availableUsernames []string
	for _, username := range desiredUsernames.Usernames {
		resp, err := http.Get(twitterURL + username)
		if err != nil {
			return AvailableUsernames{Usernames: availableUsernames}, fmt.Errorf("could not get url: %v", err)
		}
		if resp.StatusCode == http.StatusNotFound {
			availableUsernames = append(availableUsernames, username)
		}
	}

	return AvailableUsernames{Usernames: availableUsernames}, nil
}

func main() {
	lambda.Start(handler)
}
