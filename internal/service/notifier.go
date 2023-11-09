package service

type Notifier interface {
	SendNotification(message, target string) error
}
