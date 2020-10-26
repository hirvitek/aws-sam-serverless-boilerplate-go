package notifications

import (
	"boilerplate/sharedKernel/awsUtils"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"os"
	"strings"
)

type sesAdapter struct {
	sesClient       *ses.SES
	mainAddress     string
	internalAddress string
}

func NewSES() *sesAdapter {
	return &sesAdapter{
		sesClient:       awsUtils.NewSES(),
		mainAddress:     os.Getenv("MAIN_EMAIL_ADDRESS"),
	}
}

func createStandardCommunicationData(data StandardCommunicationTags, email string) (string, error) {
	name := strings.Split(email, "@")[0]
	marshal, err := json.Marshal(struct {
		Subject string `json:"subject"`
		Message string `json:"message"`
		Name    string `json:"name"`
	}{
		Subject: data.Subject,
		Message: data.Message,
		Name:    name,
	})
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

func (h *sesAdapter) sendWithTemplate(toAddresses []*string, messageData, templateName string) error {
	params := &ses.SendTemplatedEmailInput{
		Destination: &ses.Destination{
			ToAddresses: toAddresses,
		},
		Source:       aws.String(""),
		Template:     aws.String(templateName),
		TemplateData: aws.String(messageData),
	}
	
	if _, err := h.sesClient.SendTemplatedEmail(params); err != nil {
		return err
	}
	
	return nil
	
}

func (h *sesAdapter) SendCommunication(to string, data StandardCommunicationTags) error {
	communicationData, err := createStandardCommunicationData(data, to)
	if err != nil {
		return err
	}
	if err := h.sendWithTemplate([]*string{&to}, communicationData, StandardCommunicationTemplateName); err != nil {
		return err
	}
	
	return nil
}

type SESPort interface {
	SendCommunication(to string, data StandardCommunicationTags) error
}
