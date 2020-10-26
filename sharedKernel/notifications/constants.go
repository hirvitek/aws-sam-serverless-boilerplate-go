package notifications

// TODO create an interface for the notifications package
type Sender interface {
	Send(msg string) error
	SendEmail(to string, )
}

type StandardCommunicationTags struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
}

const StandardCommunicationTemplateName = "StandardCommunication"
