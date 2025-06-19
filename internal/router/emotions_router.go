package router

import (
	"emoLog/internal/handler"
	"emoLog/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterEmoRoutes 注册所有路由
func RegisterEmoRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("emotions")
	err := container.Invoke(func(emoHandler *handler.EmoHandler) {
		// 创建
		router.POST("/", middleware.Jwt(true), emoHandler.Create)
		// 更新
		router.PUT("/:id", middleware.Jwt(true), emoHandler.Update)
		// 删除
		router.DELETE("/", middleware.Jwt(true), emoHandler.Delete)
		// emotions 数据
		router.GET("/", middleware.Jwt(true), emoHandler.List)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
