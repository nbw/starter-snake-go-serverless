package main

import (
	"encoding/json"
  "log"

  "github.com/battlesnakeio/starter-snake-go/api"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse
type Request  events.APIGatewayProxyRequest

func Handler(request Request) (Response, error) {
  decoded := api.SnakeRequest{}

  err := json.Unmarshal([]byte(request.Body), &decoded)

	if err != nil {
		log.Printf("Bad start request: %v", err)
	}

  response := api.StartResponse{
		Color: "#75CEDD",
	}

  out, err := json.Marshal(response)
  if err != nil {
    panic (err)
  }
  
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
    Body: string(out),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "start-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

