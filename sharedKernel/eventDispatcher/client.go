package eventDispatcher

import (
	"boilerplate/sharedKernel/awsUtils"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

type snsAdapter struct {
	topicArn string
	sns      *sns.SNS
}

func New(topicArn string) *snsAdapter {
	newSNS := awsUtils.NewSNS()
	return &snsAdapter{
		topicArn,
		newSNS,
	}
}

func (s *snsAdapter) marshalMessage(message interface{}) (string, error) {
	payload, err := json.Marshal(message)
	if err != nil {
		return "", err
	}
	
	return string(payload), nil
}

func (s *snsAdapter) send(message string) error {
	options := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(s.topicArn),
	}
	_, err := s.sns.Publish(options)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *snsAdapter) Dispatch(message interface{}) error {
	marshalMessage, err := s.marshalMessage(message)
	if err != nil {
		return err
	}
	if err := s.send(marshalMessage); err != nil {
		return err
	}
	return nil
}

func (s *snsAdapter) DispatchWithAttributes(message interface{}, messageAttributes map[string]string) error {
	marshalMessage, err := s.marshalMessage(message)
	if err != nil {
		return err
	}
	
	messageAttributeValues := make(map[string]*sns.MessageAttributeValue)
	
	if len(messageAttributes) > 0 {
		for key, value := range messageAttributes {
			if value != "" {
				messageAttributeValues[key] = &sns.MessageAttributeValue{
					DataType:    aws.String("String"),
					StringValue: aws.String(value),
				}
			}
		}
	}
	
	options := &sns.PublishInput{
		Message:           aws.String(marshalMessage),
		MessageAttributes: messageAttributeValues,
		TopicArn:          aws.String(s.topicArn),
	}
	if _, err := s.sns.Publish(options); err != nil {
		return err
	}
	
	return nil
}

type Port interface {
	Dispatch(message interface{}) error
	DispatchWithAttributes(message interface{}, messageAttributes map[string]string) error
}
