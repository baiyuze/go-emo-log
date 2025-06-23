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

	router := r.Group("version")
	err := container.Invoke(func(versionHandler *handler.VersionHandler) {
		// 数据同步
		router.GET("/dataSync", middleware.Jwt(true), versionHandler.List)
		// 检查更新
		router.POST("/checkUpdate", middleware.Jwt(true), versionHandler.Create)
		// 用户反馈
		router.DELETE("/userFeedback", middleware.Jwt(true), versionHandler.Delete)
		// 修改
		router.PUT("/:id", middleware.Jwt(true), versionHandler.Update)
		// 根据code获取options
		router.GET("/:code", middleware.Jwt(true), versionHandler.GetOptions)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
