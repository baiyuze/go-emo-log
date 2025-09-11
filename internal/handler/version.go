package handler

import (
	errs "emoLog/internal/common/error"
	"emoLog/internal/common/log"
	"emoLog/internal/dto"
	"emoLog/internal/model"
	"emoLog/internal/service"
	"emoLog/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type VersionHandler struct {
	service service.VersionService
	log     *log.LoggerWithContext
}

func NewVersionHandler(
	s service.VersionService,
	l *log.LoggerWithContext,
) *VersionHandler {
	return &VersionHandler{
		service: s,
		log:     l,
	}
}

func ProviderVersionHandler(container *dig.Container) {
	if err := container.Provide(NewVersionHandler); err != nil {
		return
	}
}

func (h *VersionHandler) Create(c *gin.Context) {
	var body model.Version
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

// CheckUpdate 校验版本
// @Summary 校验版本
// @Tags 版本管理
// @Accept  json
// @Param data body model.Version true "body"
// @Success 200  {object} dto.Response[bool]
// @Router /api/versions/{versionName} [get]
func (h *VersionHandler) CheckUpdate(c *gin.Context) {
	versionName := c.Param("versionName")
	if len(versionName) == 0 {
		errs.FailWithJSON(c, errors.New("版本号不能为空"))
		return
	}
	version, err := h.service.CheckUpdate(c, versionName)
	if err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if version.ID == 0 {
		c.JSON(http.StatusOK, dto.Ok[any](nil))
		return
	}
	c.JSON(http.StatusOK, dto.Ok(version))
}

// Update 更新
// @Summary 更新
// @Tags 版本管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Success 200  {object} dto.Response[any]
// @Router /api/versions/{id} [put]
func (h *VersionHandler) Update(c *gin.Context) {
	var body model.Version
	var versionId uint64
	id := c.Param("id")
	if len(id) != 0 {
		result, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			errs.FailWithJSON(c, err)
			return
		}
		versionId = result
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if err := h.service.Update(c, versionId, &body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}

// List 查询
// @Summary 版本列表
// @Tags 版本列表
// @Accept json
// @Produce json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param version query string false "版本名称"
// @Success 200 {object} dto.Response[dto.List[[]model.Dict]]
// @Router /api/versions [get]
func (h *VersionHandler) List(c *gin.Context) {
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")
	version := c.Query("version")

	result, err := h.service.List(c, utils.HandleQuery(pageNum, pageSize), version)
	if err != nil {
		errs.FailWithJSON(c, err)
	} else {
		c.JSON(http.StatusOK, dto.Ok(result.Data))
	}
}

// Delete 删除
// @Summary 删除
// @Tags 版本管理
// @Accept  json
// @Param data body dto.DeleteIds true "删除ids"
// @Success 200  {object} dto.Response[any]
// @Router /api/versions [delete]
func (h *VersionHandler) Delete(c *gin.Context) {
	var ids dto.DeleteIds
	if err := c.ShouldBindJSON(&ids); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if err := h.service.Delete(c, ids); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}
