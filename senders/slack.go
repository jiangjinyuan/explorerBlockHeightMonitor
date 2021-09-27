package senders

import (
	"bytes"
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/client"

	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	log "github.com/sirupsen/logrus"
)

type SlackSender struct{}

func NewSlackSender() *SlackSender {
	return &SlackSender{}
}

func (s *SlackSender) Send(text string) {
	data := "{\"text\":\"" + text + "\"}"
	_, err := client.NewHTTPClient().Post(configs.Config.Slack.WebHookURL,
		bytes.NewBuffer([]byte(data)), "application/json;charset=utf-8")
	if err != nil {
		log.Error(err)
		return
	}
}
