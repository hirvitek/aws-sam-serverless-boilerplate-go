package notifications

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	functionName = os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	env          = os.Getenv("GO_ENV")
	region       = os.Getenv("AWS_REGION")
)

type ErrorFields struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type SlackRequestBody struct {
	Pretext string        `json:"pretext"`
	Color   string        `json:"color"`
	Fields  []ErrorFields `json:"fields"`
}

type slackAdapter struct {
	webhookUrl string
}

func NewSlack() *slackAdapter {
	return &slackAdapter{
		webhookUrl: os.Getenv("SLACK_WEBHOOK_URL"),
	}
}

func (s *slackAdapter) formatSlackErrorMessage(err error) string {
	body := SlackRequestBody{
		Pretext: fmt.Sprintf("Error from function %v, env: %v, region: %v", functionName, env, region),
		Color:   "#D00000",
		Fields: []ErrorFields{
			{
				Title: "Error message",
				Value: err.Error(),
				Short: false,
			},
		},
	}
	
	marshal, _ := json.Marshal(body)
	return string(marshal)
}

func (s *slackAdapter) sendSlackErrorNotification(msg string) error {
	req, err := http.NewRequest(http.MethodPost, s.webhookUrl, bytes.NewBuffer([]byte(msg)))
	if err != nil {
		return err
	}
	
	req.Header.Add("Content-Type", "application/json")
	
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("non-ok response returned from Slack")
	}
	
	return nil
}

func (s *slackAdapter) SendError(err error) {
	message := s.formatSlackErrorMessage(err)
	// TODO: maybe a go routine would be more effective to avoid failures due to Slack non-availability
	// This is not tested
	go s.sendSlackErrorNotification(message)
}
