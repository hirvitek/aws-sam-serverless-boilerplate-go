package main

import (
	"boilerplate/sharedKernel/awsUtils"
	"boilerplate/sharedKernel/logger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type exampleResponse struct {
	Message string `json:"message"`
}

var log = logger.New()

func function(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.SetMetadata(map[string]string{
		"requestId": request.RequestContext.RequestID,
		"requestTime": time.Now().String(),
	})
	log.Info("event", request.Headers)
	r := exampleResponse{Message: "hello"}
	return awsUtils.APISuccessResponse("response", r)
}

func main() {
	lambda.Start(function)
}

