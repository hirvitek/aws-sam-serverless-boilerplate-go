package main

import (
	"boilerplate/sharedKernel/awsUtils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type exampleResponse struct {
	Message string `json:"message"`
}

func function(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := exampleResponse{Message: "hello"}
	return awsUtils.APISuccessResponse("response", r)
}

func main() {
	lambda.Start(function)
}

