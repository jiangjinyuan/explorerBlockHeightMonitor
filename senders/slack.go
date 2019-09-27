package senders

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"bytes"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type SlackSender struct {

}

var SlackPoster SlackSender

func (s SlackSender) SendText(text map[string]string) {
	if !configs.Config.Slack.IsEnable{
		return
	}
	textBody:=""
	temp:=""
	for key,result:=range text{
		if key =="0"{
			temp=result+"\n"
		}else{
			result=result+"\n"
			textBody+=result
		}
	}
	textBody=temp+textBody
	data:="{\"text\":\"" + textBody + "\"}"
	_, err := http.Post(configs.Config.Slack.WebHookURL, "application/json;charset=utf-8",
		bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Error(err)
		return
	}

	log.Info(textBody)
}

func (s SlackSender) Send(text string) {
	if !configs.Config.Slack.IsEnable{
		return
	}
	data:="{\"text\":\"" + text + "\"}"
	_, err := http.Post(configs.Config.Slack.WebHookURL, "application/json;charset=utf-8",
		bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Error(err)
		return
	}
}