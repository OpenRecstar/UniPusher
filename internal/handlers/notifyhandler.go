package handlers

import (
	"github.com/OpenRecstar/UniPusher/internal/factory"
	"github.com/OpenRecstar/UniPusher/internal/utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NotifyRequest 定义了推送通知请求的结构
type NotifyRequest struct {
	Message string `json:"message"`
	Target  string `json:"target"`
}

// NotifyHandler 处理推送通知的请求
func NotifyHandler(c *gin.Context) {
	var req NotifyRequest
	if err := c.BindJSON(&req); err != nil {
		log.Log.WithError(err).Error("Failed to bind JSON for notify request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notifier := factory.GetNotifier("webhook")
	if notifier == nil {
		log.Log.Error("Notifier type is not supported")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notifier type is not supported"})
		return
	}

	err := notifier.SendNotification(req.Message, req.Target)
	if err != nil {
		log.Log.WithError(err).Error("Failed to send notification")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
		return
	}

	log.Log.Info("Notification sent successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully"})
}
