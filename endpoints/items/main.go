package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type item struct {
	Title string `json:"title"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var responseBody string

	items := []item{
		{Title: "First Item"},
		{Title: "Second Item"},
		{Title: "Third Item"},
	}

	jsonBody, err := json.Marshal(items)
	if err != nil {
		responseBody = "failed to fetch strings"
	} else {
		responseBody = string(jsonBody)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       responseBody,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
