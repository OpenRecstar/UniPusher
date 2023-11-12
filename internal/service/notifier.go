package service

// Notifier 通知服务接口
type Notifier interface {
	// SendNotification 发送通知方法
	// 参数:
	//   - message: 通知消息
	//   - target: 目标
	SendNotification(message, target string) error
}
