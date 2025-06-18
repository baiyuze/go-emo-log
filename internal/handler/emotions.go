package handler

import (
	"emoLog/internal/common/log"
	"emoLog/internal/grpc/container"
	"emoLog/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type EmoHandler struct {
	service service.EmoService
	log     *log.LoggerWithContext
	clients *container.Clients
}

func NewEmoHandler(
	s service.EmoService,
	l *log.LoggerWithContext,
	clients *container.Clients,
) *EmoHandler {
	return &EmoHandler{
		service: s,
		log:     l,
		clients: clients,
	}
}

func ProviderEmoHandler(container *dig.Container) {
	if err := container.Provide(NewEmoHandler); err != nil {
		return
	}
}

// Create emo 创建数据
// @Summary emo 创建数据
// @Tags emo数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/emotions [post]
func (h *EmoHandler) Create(c *gin.Context) {

}

// List emo 获取emo数据
// @Summary emo 更新emo数据
// @Tags emo数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/emotions [get]
func (h *EmoHandler) List(c *gin.Context) {

}

// Update emo 更新数据
// @Summary emo 更新数据
// @Tags emo数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/emotions [put]
func (h *EmoHandler) Update(c *gin.Context) {

}

// Delete emo
// @Summary emo 删除数据
// @Tags emo数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/emotions [delete]
func (h *EmoHandler) Delete(c *gin.Context) {

}
