package service

import (
	"emoLog/internal/common/log"
	"emoLog/internal/dto"
	"emoLog/internal/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type FeedbackService interface {
	Create(c *gin.Context, body *dto.Feedback) error
	Update(c *gin.Context, id uint64, body *dto.Feedback) error
	Delete(c *gin.Context, body dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery, userId uint64) (dto.Result[dto.List[model.Feedback]], error)
	//Update(c *gin.Context, id int, body *model.Dict) error
	//GetOptionsByDictCode(c *gin.Context, code string) ([]*model.DictItem, error)
}

type feedbackService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewFeedbackService(
	db *gorm.DB,
	log *log.LoggerWithContext) FeedbackService {

	return &feedbackService{db: db, log: log}
}

func ProvideFeedbackService(container *dig.Container) {
	if err := container.Provide(NewFeedbackService); err != nil {
		panic(err)
	}
}

func (s *feedbackService) Create(c *gin.Context, body *dto.Feedback) error {

	data := model.Feedback{
		Description: body.Description,
		Version:     body.Version,
		VersionId:   body.VersionId,
	}
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *feedbackService) Update(c *gin.Context, id uint64, body *dto.Feedback) error {

	data := model.Feedback{
		Description: body.Description,
		Version:     body.Version,
	}
	if err := s.db.Model(&model.Feedback{
		ID: id,
	}).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *feedbackService) Delete(c *gin.Context, body dto.DeleteIds) error {
	if err := s.db.Delete(&model.Feedback{}, body.Ids).Error; err != nil {
		return err
	}
	return nil
}

func (s *feedbackService) List(
	context *gin.Context,
	query dto.ListQuery,
	userId uint64) (dto.Result[dto.List[model.Feedback]], error) {
	var emotionLogs []model.Feedback
	limit := query.PageSize
	offset := query.PageNum*query.PageSize - query.PageSize

	if result := s.db.
		Model(&model.Feedback{}).
		Limit(limit).
		Offset(offset).
		Order("create_time asc").
		Find(&emotionLogs); result.Error != nil {
		return dto.ServiceFail[dto.List[model.Feedback]](result.Error), result.Error
	}
	var count int64
	if result := s.db.Model(&model.Feedback{}).Count(&count); result.Error != nil {
		return dto.ServiceFail[dto.List[model.Feedback]](result.Error), result.Error
	}

	data := dto.ServiceSuccess(dto.List[model.Feedback]{
		Items:    emotionLogs,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    count,
	})
	return data, nil
}
