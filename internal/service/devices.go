package service

import (
	"emoLog/internal/common/log"
	"emoLog/internal/dto"
	"emoLog/internal/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type DevicesService interface {
	Create(c *gin.Context, body *model.Device) error
	Update(c *gin.Context, id uint64, body *model.Device) error
	Delete(c *gin.Context, body dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery, userId uint64) (dto.Result[dto.List[model.Device]], error)
}

type devicesService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewDevicesService(
	db *gorm.DB,
	log *log.LoggerWithContext) DevicesService {

	return &devicesService{db: db, log: log}
}

func ProvideDevicesService(container *dig.Container) {
	if err := container.Provide(NewDevicesService); err != nil {
		panic(err)
	}
}

func (s *devicesService) Create(c *gin.Context, body *model.Device) error {

	data := model.Device{
		//Description: body.Description,
		//Version:     body.Version,
	}
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *devicesService) Update(c *gin.Context, id uint64, body *model.Device) error {

	data := model.Device{
		//Description: body.Description,
		//Version:     body.Version,
	}
	if err := s.db.Model(&model.Device{
		ID: id,
	}).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *devicesService) Delete(c *gin.Context, body dto.DeleteIds) error {
	if err := s.db.Delete(&model.Device{}, body.Ids).Error; err != nil {
		return err
	}
	return nil
}

func (s *devicesService) List(
	context *gin.Context,
	query dto.ListQuery,
	userId uint64) (dto.Result[dto.List[model.Device]], error) {
	var emotionLogs []model.Device
	limit := query.PageSize
	offset := query.PageNum*query.PageSize - query.PageSize

	if result := s.db.
		Model(&model.Device{}).
		Where("user_id = ?", userId).
		Limit(limit).
		Offset(offset).
		Order("create_time asc").
		Find(&emotionLogs); result.Error != nil {
		return dto.ServiceFail[dto.List[model.Device]](result.Error), result.Error
	}
	var count int64
	if result := s.db.Model(&model.Device{}).Count(&count); result.Error != nil {
		return dto.ServiceFail[dto.List[model.Device]](result.Error), result.Error
	}

	data := dto.ServiceSuccess(dto.List[model.Device]{
		Items:    emotionLogs,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    count,
	})
	return data, nil
}
