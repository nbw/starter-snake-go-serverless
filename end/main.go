package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse
type Request  events.APIGatewayProxyRequest

// Note: reponse of /end endpoint is ignored
func Handler(request Request) (Response, error) {
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "Game over.",
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "end-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
