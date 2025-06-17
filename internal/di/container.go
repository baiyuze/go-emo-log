package di

import (
	"emoLog/config"
	"emoLog/internal/common/log"
	grpcContainer "emoLog/internal/grpc/container"
	"emoLog/internal/handler"
	"emoLog/internal/repo"
	"emoLog/internal/service"

	"go.uber.org/dig"
)

// NewContainer 创建并初始化 DI 容器
func NewContainer() *dig.Container {
	// 注册各模块的依赖
	container := dig.New()
	//// 注入logger
	//if err := container.Provide(func() *zap.Logger {
	//	return logger
	//}); err != nil {
	//	logger.Fatal("日志注入失败", zap.Error(err))
	//}
	// 公共日志管理器
	log.NewProvideLogger(container)
	// 获取客户端grpc
	grpcContainer.NewProvideClients(container)
	// 配置
	config.ProvideConfig(container)
	// 数据库
	repo.ProvideDB(container)
	// 服务
	service.Provide(container)
	// controller
	handler.Provide(container)

	return container
}
