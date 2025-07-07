package router

import (
	"emoLog/internal/handler"
	"emoLog/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterDevicesRoutes 注册所有路由
func RegisterDevicesRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("devices")
	err := container.Invoke(func(devicesHandler *handler.DevicesHandler) {
		// 创建
		router.POST("/device", middleware.Jwt(true), devicesHandler.Create)
		// 更新
		router.PUT("/device/:id", middleware.Jwt(true), devicesHandler.Update)
		// 删除
		router.DELETE("/device", middleware.Jwt(true), devicesHandler.Delete)
		// feedback 数据
		router.GET("/device", middleware.Jwt(true), devicesHandler.List)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
