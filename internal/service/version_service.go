package service

import (
	"emoLog/internal/common/log"
	"emoLog/internal/dto"
	"emoLog/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type VersionService interface {
	CheckUpdate(c *gin.Context, versionName string) (model.Version, error)
	Create(c *gin.Context, body *model.Version) error
	Delete(c *gin.Context, body dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery, versionName string) (dto.Result[dto.List[model.Version]], error)
	Update(c *gin.Context, id uint64, body *model.Version) error
}

type versionService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewVersionService(
	db *gorm.DB,
	log *log.LoggerWithContext) VersionService {

	return &versionService{db: db, log: log}
}

func ProvideVersionService(container *dig.Container) {
	if err := container.Provide(NewVersionService); err != nil {
		panic(err)
	}
}

// CheckUpdate 检查更新
func (s *versionService) CheckUpdate(c *gin.Context, versionName string) (model.Version, error) {
	var versionItem model.Version
	if err := s.db.Order("create_time desc").First(&versionItem).Error; err != nil {
		return model.Version{}, err
	}

	if versionName != versionItem.Version {
		return versionItem, nil
	} else {
		return model.Version{}, nil
	}
}

// Create 创建版本
func (s *versionService) Create(c *gin.Context, body *model.Version) error {
	if err := s.db.Create(&model.Version{
		Version:     body.Version,
		Description: body.Description,
	}).Error; err != nil {
		return err
	}

	return nil
}

// List 列表
func (s *versionService) List(c *gin.Context, query dto.ListQuery, versionName string) (dto.Result[dto.List[model.Version]], error) {
	logger := s.log.WithContext(c)
	var versions []model.Version
	db := s.db

	limit := query.PageSize
	offset := query.PageNum*query.PageSize - query.PageSize

	if len(versionName) != 0 {
		db = s.db.Where("version LIKE ?", "%"+versionName+"%")
	}
	if result := db.
		Limit(limit).
		Offset(offset).
		Order("create_time desc").
		Find(&versions); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[model.Version]](result.Error), result.Error
	}
	var count int64
	if result := s.db.Model(&model.Version{}).Count(&count); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[model.Version]](result.Error), result.Error
	}
	data := dto.ServiceSuccess(dto.List[model.Version]{
		Items:    versions,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    count,
	})
	return data, nil
}

// Update 更新字典数据,需要使用事务
func (s *versionService) Update(c *gin.Context, id uint64, body *model.Version) error {
	if err :=
		s.db.Model(&model.Version{
			ID: id,
		}).Updates(model.Version{
			Version:     body.Version,
			Description: body.Description,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (s *versionService) Delete(c *gin.Context, body dto.DeleteIds) error {
	if err := s.db.Delete(&model.Version{}, body.Ids).Error; err != nil {
		return err
	}
	return nil
}
