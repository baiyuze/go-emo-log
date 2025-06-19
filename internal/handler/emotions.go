package handler

import (
	errs "emoLog/internal/common/error"
	"emoLog/internal/common/log"
	"emoLog/internal/dto"
	"emoLog/internal/grpc/container"
	"emoLog/internal/service"
	"emoLog/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"strconv"
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

	var body dto.EmotionLog
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	if err := h.service.Create(c, &body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}

// List emo 获取emo数据
// @Summary emo 更新emo数据
// @Tags emo数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/emotions [get]
func (h *EmoHandler) List(c *gin.Context) {
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")
	userId := c.Query("userId")

	if len(userId) == 0 {
		errs.FailWithJSON(c, errors.New("userId不能为空"))
		return
	}

	id, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	result, err := h.service.List(c, utils.HandleQuery(pageNum, pageSize), id)
	if err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok(result.Data))
	return
}

// Update emo 更新数据
// @Summary emo 更新数据
// @Tags emo数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/emotions/{id} [put]
func (h *EmoHandler) Update(c *gin.Context) {
	queryId := c.Param("id")
	var body dto.EmotionLog

	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	id, err := strconv.ParseUint(queryId, 10, 64)
	if err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	if err := h.service.Update(c, id, &body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Ok[any](nil))

}

// Delete emo
// @Summary emo 删除数据
// @Tags emo数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/emotions [delete]
func (h *EmoHandler) Delete(c *gin.Context) {
	var body dto.DeleteIds
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	if err := h.service.Delete(c, body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}
