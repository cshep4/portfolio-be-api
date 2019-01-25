package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"portfolio-be-api/app/request"
	"portfolio-be-api/app/router"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing Lambda req %s\n", req.RequestContext.RequestID)

	fmt.Printf("Chris::Request: %s\n", req)

	if len(req.Body) < 1 {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: map[string]string{"Content-Type": "application/json","Access-Control-Allow-Origin": "*" }}, nil
	}

	var body request.RequestBody
	err := json.Unmarshal([]byte(req.Body), &body)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: map[string]string{"Content-Type": "application/json","Access-Control-Allow-Origin": "*" }}, nil
	}

	return router.Process(req.Path, body)
}

func main() {
	lambda.Start(Handler)
}