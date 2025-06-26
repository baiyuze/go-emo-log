package router

import (
	"emoLog/internal/handler"
	"emoLog/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterVersionRoutes 注册所有路由
func RegisterVersionRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("versions")
	err := container.Invoke(func(versionHandler *handler.VersionHandler) {
		// 获取版本列表
		router.GET("/", middleware.Jwt(true), versionHandler.List)
		// 数据同步
		router.GET("/dataSync", middleware.Jwt(true), versionHandler.List)
		// 检查更新
		router.GET("/:versionName", middleware.Jwt(true), versionHandler.CheckUpdate)
		// 创建版本
		router.POST("/", middleware.Jwt(true), versionHandler.Create)
		// 用户反馈
		router.POST("/userFeedback", middleware.Jwt(true), versionHandler.Delete)

		router.DELETE("/", middleware.Jwt(true), versionHandler.Delete)
		// 修改
		router.PUT("/:id", middleware.Jwt(true), versionHandler.Update)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
