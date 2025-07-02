package router

import (
	"emoLog/internal/handler"
	"emoLog/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterFeedbackRoutes 注册所有路由
func RegisterFeedbackRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("feedback")
	err := container.Invoke(func(feedbackHandler *handler.FeedbackHandler) {
		// 创建
		router.POST("/", middleware.Jwt(true), feedbackHandler.Create)
		// 更新
		router.PUT("/:id", middleware.Jwt(true), feedbackHandler.Update)
		// 删除
		router.DELETE("/", middleware.Jwt(true), feedbackHandler.Delete)
		// feedback 数据
		router.GET("/", middleware.Jwt(true), feedbackHandler.List)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
