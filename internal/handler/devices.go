package handler

import (
	errs "emoLog/internal/common/error"
	"emoLog/internal/common/log"
	"emoLog/internal/dto"
	"emoLog/internal/grpc/container"
	"emoLog/internal/model"
	"emoLog/internal/service"
	"emoLog/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type DevicesHandler struct {
	service service.DevicesService
	log     *log.LoggerWithContext
	clients *container.Clients
}

func NewDevicesHandlerHandler(
	s service.DevicesService,
	l *log.LoggerWithContext,
	clients *container.Clients,
) *DevicesHandler {
	return &DevicesHandler{
		service: s,
		log:     l,
		clients: clients,
	}
}

func ProviderDevicesHandler(container *dig.Container) {
	if err := container.Provide(NewDevicesHandlerHandler); err != nil {
		return
	}
}

// Create Devices 创建数据
// @Summary Devices 创建数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/Devices [post]
func (h *DevicesHandler) Create(c *gin.Context) {

	var body model.Device
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

// List Device 获取emo数据
// @Summary Device 更新emo数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/Devices [get]
func (h *DevicesHandler) List(c *gin.Context) {
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")

	result, err := h.service.List(c, utils.HandleQuery(pageNum, pageSize))
	if err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok(result.Data))
}

// Update Device 更新数据
// @Summary Device 更新数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/Devices/{id} [put]
func (h *DevicesHandler) Update(c *gin.Context) {
	queryId := c.Param("id")
	var body model.Device

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
// @Summary Device 删除数据
// @Tags 反馈数据管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Router /api/Devices [delete]
func (h *DevicesHandler) Delete(c *gin.Context) {
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
