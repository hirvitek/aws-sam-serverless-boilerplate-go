package awsUtils

import (
	"boilerplate/sharedKernel/appError"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"strings"
)

func UnmarshalDynamoDBStreamImage(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {
	dbAttrMap := make(map[string]*dynamodb.AttributeValue)
	
	for k, v := range attribute {
		var dbAttr dynamodb.AttributeValue
		
		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return marshalErr
		}
		
		unmarshalErr := json.Unmarshal(bytes, &dbAttr)
		if unmarshalErr != nil {
			return unmarshalErr
		}
		dbAttrMap[k] = &dbAttr
	}
	
	return dynamodbattribute.UnmarshalMap(dbAttrMap, out)
}

type cognitoClaims struct {
	Sub                 string `json:"sub"`
	Profile             string `json:"profile"`
	Picture             string `json:"picture"`
	Email               string `json:"email"`
	HasUserRefundPolicy string `json:"custom:hasUsedRefundPolicy"`
}

func UnmarshalCognitoClaims(claims interface{}) (*cognitoClaims, error) {
	output := &cognitoClaims{}
	err := mapstructure.Decode(claims, &output)
	
	if err != nil {
		return &cognitoClaims{}, err
	}
	
	return output, nil
}

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APISuccessResponse(message string, data interface{}) (events.APIGatewayProxyResponse, error) {
	r := response{
		Message: message,
		Data:    data,
	}
	marshal, err := json.Marshal(r)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(marshal),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Methods":     "*",
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}

func APIFailureResponse(err appError.AppError) (events.APIGatewayProxyResponse, error) {
	r := response{
		Message: err.Error(),
	}
	marshal, err1 := json.Marshal(r)
	if err1 != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	
	return events.APIGatewayProxyResponse{
		StatusCode: err.GetStatusCode(),
		Body:       string(marshal),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Methods":     "*",
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}

func GetIdFromArn(arn string, service string) string {
	return strings.Split(arn, service+"/")[1]
}

func TrimRevisionFromId(id string) string {
	return strings.Split(id, ":")[0]
}

func GetRevisionVersionFromArn(arn string) string {
	res := strings.Split(arn, ":")
	return res[len(res)-1]
}

func ParseCloudwatchLogFromKinesis(data []byte) (events.CloudwatchLogsData, error) {
	var logData events.CloudwatchLogsData
	
	rawData := bytes.NewReader(data)
	r, err := gzip.NewReader(rawData)
	if err != nil {
		return events.CloudwatchLogsData{}, err
	}
	
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return events.CloudwatchLogsData{}, err
	}
	
	if err := json.Unmarshal(s, &logData); err != nil {
		return events.CloudwatchLogsData{}, err
	}
	
	return logData, nil
}
