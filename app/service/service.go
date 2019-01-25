package service

import "github.com/aws/aws-lambda-go/events"

type Service interface {
	Name() string
	Path() string
	Process(methodName string, args string) (events.APIGatewayProxyResponse, error)
}