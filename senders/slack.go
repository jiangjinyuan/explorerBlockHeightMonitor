package senders

import (
	"bytes"
	"net/http"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	log "github.com/sirupsen/logrus"
)

type SlackSender struct{}

func NewSlackSender() *SlackSender {
	return &SlackSender{}
}

func (s *SlackSender) Send(text string) {
	data := "{\"text\":\"" + text + "\"}"
	_, err := http.Post(configs.Config.Slack.WebHookURL, "application/json;charset=utf-8",
		bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Error(err)
		return
	}
}
