package senders

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"gopkg.in/gomail.v2"
)

type EmailSender struct {
}

var EmailPublisher EmailSender

func (e *EmailSender) SendText(text map[string]string) {
	if !configs.Config.Email.IsEnable {
		return
	}
	textBody := ""
	temp := ""
	for key, result := range text {
		if key == "0" {
			temp = "<p>" + result + "<p>"
		} else {
			result = "<p>" + result + "<p>"
			textBody += result
		}
	}
	textBody = temp + textBody
	m := gomail.NewMessage()
	//sender
	m.SetAddressHeader("From", configs.Config.Email.SenderName, "sender")
	//receiver
	m.SetHeader("To", m.FormatAddress(configs.Config.Email.SenderName, "receiver"))
	m.SetHeader("Subject", "Attention from blockHeightMonitor!")
	m.SetBody("text/html", textBody)

	d := gomail.NewDialer(configs.Config.Email.Host, configs.Config.Email.Port, configs.Config.Email.SenderName, configs.Config.Email.SenderPassword)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func (e *EmailSender) Send(mess string) {
	if !configs.Config.Email.IsEnable {
		return
	}
	m := gomail.NewMessage()
	//sender
	m.SetAddressHeader("From", configs.Config.Email.SenderName, "sender")
	//receiver
	m.SetHeader("To", m.FormatAddress(configs.Config.Email.SenderName, "receiver"))
	m.SetHeader("Subject", "Attention from blockHeightMonitor!")
	m.SetBody("text/html", mess)

	d := gomail.NewDialer(configs.Config.Email.Host, configs.Config.Email.Port, configs.Config.Email.SenderName, configs.Config.Email.SenderPassword)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
