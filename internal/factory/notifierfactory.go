package factory

import "github.com/OpenRecstar/UniPusher/internal/service"

func GetNotifier(notifierType string) service.Notifier {
	switch notifierType {
	case "webhook":
		return service.NewWebhookNotifier()
	// 可以添加其他通知类型
	default:
		return nil // 或者返回一个默认的通知发送器
	}
}
