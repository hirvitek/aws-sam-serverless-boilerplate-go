package queue

import (
	"boilerplate/sharedKernel/awsUtils"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsAdapter struct {
	client   *sqs.SQS
	queueUrl string
}

type message struct {
	Body      string `json:"body"`
	MessageId string `json:"messageId"`
}

func New(queueUrl string) *sqsAdapter {
	return &sqsAdapter{
		queueUrl: queueUrl,
		client:   awsUtils.NewSQS(),
	}
}

func (q sqsAdapter) PollMessages() ([]message, error) {
	params := &sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(10),
		QueueUrl:            aws.String(q.queueUrl),
		VisibilityTimeout:   nil,
		WaitTimeSeconds:     nil,
	}
	
	messages := make([]message, 0)
	
	response, err := q.client.ReceiveMessage(params)
	if err != nil {
		return messages, err
	}
	
	for _, msg := range response.Messages {
		messages = append(messages, message{
			Body:      *msg.Body,
			MessageId: *msg.ReceiptHandle,
		})
	}
	
	return messages, nil
}

func (q sqsAdapter) PushMessage(message interface{}) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBytes)),
		QueueUrl:    aws.String(q.queueUrl),
	}
	
	if _, err := q.client.SendMessage(params); err != nil {
		return err
	}
	
	return nil
}

func (q sqsAdapter) DeleteMessage(messageId string) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(q.queueUrl),
		ReceiptHandle: aws.String(messageId),
	}
	
	if _, err := q.client.DeleteMessage(params); err != nil {
		return err
	}
	
	return nil
}

type Port interface {
	PollMessages() ([]message, error)
	PushMessage(message interface{}) error
	DeleteMessage(messageId string) error
}
