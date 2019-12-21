package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) error {
	usernames := []string{"krudler", "krudlers"}
	url := "https://twitter.com/"
	for _, username := range usernames {
		resp, err := http.Get(url + username)
		if err != nil {
			return fmt.Errorf("could not get url: %v", err)
		}
		fmt.Printf("%v", resp.StatusCode)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
