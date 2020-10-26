package awsUtils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-xray-sdk-go/xray"
)

func getDefaultSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("AWS_REGION"),
		},
	}))
}

func NewECS() *ecs.ECS {
	newSession := getDefaultSession()
	ecsClient := ecs.New(newSession)
	xray.AWS(ecsClient.Client)
	return ecsClient
}

func NewLambda() *lambda.Lambda {
	newSession := getDefaultSession()
	lambdaClient := lambda.New(newSession)
	xray.AWS(lambdaClient.Client)
	return lambdaClient
}

func NewDynamoDB() *dynamodb.DynamoDB {
	newSession := getDefaultSession()
	dynamoDBClient := dynamodb.New(newSession)
	xray.AWS(dynamoDBClient.Client)
	return dynamoDBClient
}

func NewSNS() *sns.SNS {
	newSession := getDefaultSession()
	snsClient := sns.New(newSession)
	xray.AWS(snsClient.Client)
	return snsClient
}

func NewEC2() *ec2.EC2 {
	newSession := getDefaultSession()
	ec2Client := ec2.New(newSession)
	xray.AWS(ec2Client.Client)
	return ec2Client
}

func NewCloudWatchLogs(region string) *cloudwatchlogs.CloudWatchLogs {
	newSession := getDefaultSession()
	cloudwatchClient := cloudwatchlogs.New(newSession)
	xray.AWS(cloudwatchClient.Client)
	return cloudwatchClient
}

func NewSQS() *sqs.SQS {
	newSession := getDefaultSession()
	sqsClient := sqs.New(newSession)
	xray.AWS(sqsClient.Client)
	return sqsClient
}

func NewAutoScaling() *autoscaling.AutoScaling {
	newSession := getDefaultSession()
	autoScalingClient := autoscaling.New(newSession)
	xray.AWS(autoScalingClient.Client)
	return autoScalingClient
}

func NewKMS() *kms.KMS {
	newSession := getDefaultSession()
	kmsClient := kms.New(newSession)
	xray.AWS(kmsClient.Client)
	return kmsClient
}

func NewSecretManager() *secretsmanager.SecretsManager {
	newSession := getDefaultSession()
	smClient := secretsmanager.New(newSession)
	xray.AWS(smClient.Client)
	return smClient
}

func NewCognito() *cognitoidentityprovider.CognitoIdentityProvider {
	newSession := getDefaultSession()
	cognitoClient := cognitoidentityprovider.New(newSession)
	xray.AWS(cognitoClient.Client)
	return cognitoClient
}

func NewSES() *ses.SES {
	newSession := getDefaultSession()
	sesClient := ses.New(newSession)
	xray.AWS(sesClient.Client)
	return sesClient
}

func NewRoute53() *route53.Route53 {
	newSession := getDefaultSession()
	route53Client := route53.New(newSession)
	xray.AWS(route53Client.Client)
	return route53Client
}
