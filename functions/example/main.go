package main

import (
	"boilerplate/sharedKernel/appError"
	"boilerplate/sharedKernel/awsUtils"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)



func function(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	apiErr := appError.NewApi(errors.New("something went wrong"), 500)
	return awsUtils.APIFailureResponse(apiErr)
}

func main() {
	lambda.Start(function)
}

