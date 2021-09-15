package senders

type Senders interface {
	Send(text string)
	SendText(text map[string]string, height string)
}
