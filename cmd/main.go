package main

import (
	"github.com/OpenRecstar/UniPusher/internal/handlers"
	"github.com/OpenRecstar/UniPusher/internal/utils/log"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志配置
	log.InitLogger()

	router := gin.Default()

	// 健康检查路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
		log.Log.Info("Health check route called")
	})

	// 推送通知路由
	router.POST("/notify", handlers.NotifyHandler)

	// 启动HTTP服务，默认在8080端口
	if err := router.Run(); err != nil {
		log.Log.WithError(err).Fatalf("Failed to start the server")
	}
}
