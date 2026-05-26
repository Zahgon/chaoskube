package notifier

import (
	"net/http"
	"time"

	v1 "k8s.io/api/core/v1"
)

const NotifierSlack = "slack"

var NotificationColor = "#F35A00"
var DefaultTimeout = 10 * time.Second

type Slack struct {
	Webhook string
	Client  *http.Client
}

type slackMessage struct {
	Message     string       `json:"text"`
	Attachments []attachment `json:"attachments"`
}

type slackField struct {
	Title string `yaml:"title,omitempty" json:"title,omitempty"`
	Value string `yaml:"value,omitempty" json:"value,omitempty"`
	Short *bool  `yaml:"short,omitempty" json:"short,omitempty"`
}

type attachment struct {
	Title      string       `json:"title,omitempty"`
	TitleLink  string       `json:"title_link,omitempty"`
	Pretext    string       `json:"pretext,omitempty"`
	Text       string       `json:"text"`
	Fallback   string       `json:"fallback"`
	CallbackID string       `json:"callback_id"`
	Fields     []slackField `json:"fields,omitempty"`
	ImageURL   string       `json:"image_url,omitempty"`
	ThumbURL   string       `json:"thumb_url,omitempty"`
	Footer     string       `json:"footer"`
	Color      string       `json:"color,omitempty"`
	MrkdwnIn   []string     `json:"mrkdwn_in,omitempty"`
}

func NewSlackNotifier(webhook string) *Slack { _ = "STUB: not implemented"; return nil }

func (s Slack) NotifyPodTermination(pod v1.Pod) error { _ = "STUB: not implemented"; return nil }

func createSlackRequest(title string, text string, fields []slackField) slackMessage {
	_ = "STUB: not implemented"
	return *new(slackMessage)
}

func (s Slack) sendSlackMessage(message slackMessage) error { _ = "STUB: not implemented"; return nil }
