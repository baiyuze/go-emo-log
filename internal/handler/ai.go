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
	"github.com/tmc/langchaingo/llms/openai"
	"go.uber.org/dig"
)

type AiHandler struct {
	service service.AiService
	log     *log.LoggerWithContext
	clients *container.Clients
	llm     *openai.LLM
}

func NewAiHandlerHandler(
	s service.AiService,
	l *log.LoggerWithContext,
	clients *container.Clients,
	llm *openai.LLM,

) *AiHandler {
	return &AiHandler{
		service: s,
		log:     l,
		clients: clients,
		llm:     llm,
	}
}

func ProviderAiHandler(container *dig.Container) {
	if err := container.Provide(NewAiHandlerHandler); err != nil {
		return
	}
}

// Chat ai 更新数据
// @Summary ai 更新数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/ai/chat{id} [put]
func (h *AiHandler) Chat(c *gin.Context) {

	msg := c.Query("msg")
	if len(msg) == 0 {
		errs.FailWithJSON(c, errors.New("msg不能为空"))
	}
	resp := h.service.TestChat(c, msg)
	c.JSON(http.StatusOK, dto.Ok(resp))

}

// Create ai 创建数据
// @Summary ai 创建数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/feedbacks [post]
func (h *AiHandler) Create(c *gin.Context) {

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

// List ai 获取emo数据
// @Summary ai 更新emo数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/feedbacks [get]
func (h *AiHandler) List(c *gin.Context) {
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

// Delete emo
// @Summary ai 删除数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/feedbacks [delete]
func (h *AiHandler) Delete(c *gin.Context) {
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
