package main

import (
	"github.com/OpenRecstar/UniPusher/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	// 健康检查路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 推送通知路由
	router.POST("/notify", handlers.NotifyHandler)

	// 启动HTTP服务，默认在8080端口
	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
