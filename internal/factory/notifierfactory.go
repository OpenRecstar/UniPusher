package factory

import (
	"github.com/OpenRecstar/UniPusher/internal/service"
	"github.com/OpenRecstar/UniPusher/internal/utils/log"
)

// GetNotifier 根据通知类型获取对应的通知发送器
func GetNotifier(notifierType string) service.Notifier {
	switch notifierType {
	case "webhook":
		log.Log.Info("Using webhook notifier")
		return service.NewWebhookNotifier()
	// 可以添加其他通知类型
	default:
		log.Log.Warn("Notifier type not found, using default notifier")
		return nil // 或者返回一个默认的通知发送器
	}
}
