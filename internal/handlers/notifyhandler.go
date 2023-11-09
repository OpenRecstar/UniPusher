package handlers

import (
	"github.com/OpenRecstar/UniPusher/internal/factory"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NotifyRequest 定义了推送通知请求的结构
type NotifyRequest struct {
	Message string `json:"message"`
	Target  string `json:"target"`
}

func NotifyHandler(c *gin.Context) {
	var req NotifyRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notifier := factory.GetNotifier("webhook")
	if notifier == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notifier type is not supported"})
		return
	}

	err := notifier.SendNotification(req.Message, req.Target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully"})
}
