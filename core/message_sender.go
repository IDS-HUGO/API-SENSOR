package core

type MessageSender interface {
	SendMessage(data *Data) error
}
