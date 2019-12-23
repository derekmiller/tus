package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

const twitterURL = "https://mobile.twitter.com/"

type DesiredUsernames struct {
	Usernames []string `json:"usernames"`
}

type AvailableUsernames struct {
	Usernames []string `json:"usernames"`
}

type Result struct {
	username     string
	httpResponse *http.Response
	err          error
}

func handler(desiredUsernames DesiredUsernames) (AvailableUsernames, error) {
	ch := make(chan *Result)
	for _, username := range desiredUsernames.Usernames {
		go func(username string) {
			resp, err := http.Get(twitterURL + username)
			ch <- &Result{username: username, httpResponse: resp, err: err}
		}(username)
	}

	var results []*Result
	for result := range ch {
		results = append(results, result)
		if len(results) == len(desiredUsernames.Usernames) {
			break
		}
	}

	var availableUsernames []string
	for _, result := range results {
		log.Printf("username: %v, status code: %v", result.username, result.httpResponse.StatusCode)
		if result.httpResponse.StatusCode == http.StatusNotFound {
			availableUsernames = append(availableUsernames, result.username)
		}
	}

	if len(availableUsernames) == 0 {
		return AvailableUsernames{}, errors.New("no available usernames found")
	}

	return AvailableUsernames{Usernames: availableUsernames}, nil
}

func main() {
	lambda.Start(handler)
}
