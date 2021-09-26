package senders

import (
	"github.com/jiangjinyuan/explorerBlockHeightMonitor/configs"
	"gopkg.in/gomail.v2"
)

type EmailSender struct{}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

func (e *EmailSender) Send(mess string) {
	m := gomail.NewMessage()
	//sender
	m.SetAddressHeader("From", configs.Config.Email.SenderName, "explorerBlockHeightMonitor")
	//receiver
	m.SetHeader("To", m.FormatAddress(configs.Config.Email.SenderName, "BTC.com"))
	m.SetHeader("Subject", "Attention from explorerBlockHeightMonitor!")
	m.SetBody("text/html", mess)

	d := gomail.NewDialer(configs.Config.Email.Host, configs.Config.Email.Port, configs.Config.Email.SenderName, configs.Config.Email.SenderPassword)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
