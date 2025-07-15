package router

import (
	"emoLog/internal/handler"
	"emoLog/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterAiRoutes 注册所有路由
func RegisterAiRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("ai")
	err := container.Invoke(func(aiHandler *handler.AiHandler) {
		// 创建
		router.POST("/chat", middleware.Jwt(true), aiHandler.Chat)
		// 更新
		router.PUT("/chat/:id", middleware.Jwt(true), aiHandler.Delete)
		// 删除
		router.DELETE("/chat", middleware.Jwt(true), aiHandler.Delete)
		// feedback 数据
		router.GET("/chat", middleware.Jwt(true), aiHandler.List)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
