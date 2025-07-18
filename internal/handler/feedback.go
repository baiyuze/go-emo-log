package handler

import (
	errs "emoLog/internal/common/error"
	"emoLog/internal/common/log"
	"emoLog/internal/dto"
	"emoLog/internal/grpc/container"
	"emoLog/internal/service"
	"emoLog/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type FeedbackHandler struct {
	service service.FeedbackService
	log     *log.LoggerWithContext
	clients *container.Clients
}

func NewFeedbackHandlerHandler(
	s service.FeedbackService,
	l *log.LoggerWithContext,
	clients *container.Clients,
) *FeedbackHandler {
	return &FeedbackHandler{
		service: s,
		log:     l,
		clients: clients,
	}
}

func ProviderFeedbackHandler(container *dig.Container) {
	if err := container.Provide(NewFeedbackHandlerHandler); err != nil {
		return
	}
}

// Create feedback 创建数据
// @Summary feedback 创建数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/feedbacks [post]
func (h *FeedbackHandler) Create(c *gin.Context) {

	var body dto.Feedback
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

// List feedback 获取emo数据
// @Summary feedback 更新emo数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/feedbacks [get]
func (h *FeedbackHandler) List(c *gin.Context) {
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
}

// Update feedback 更新数据
// @Summary feedback 更新数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/feedbacks/{id} [put]
func (h *FeedbackHandler) Update(c *gin.Context) {
	queryId := c.Param("id")
	var body dto.Feedback

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
// @Summary feedback 删除数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/feedbacks [delete]
func (h *FeedbackHandler) Delete(c *gin.Context) {
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
