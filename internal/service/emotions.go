package service

import (
	"emoLog/internal/common/log"
	"emoLog/internal/dto"
	"emoLog/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"time"
)

type EmoService interface {
	Create(c *gin.Context, body *dto.EmotionLog) error
	Update(c *gin.Context, id uint64, body *dto.EmotionLog) error
	Delete(c *gin.Context, body dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery, userId uint64) (dto.Result[dto.List[model.EmotionLog]], error)
	//Update(c *gin.Context, id int, body *model.Dict) error
	//GetOptionsByDictCode(c *gin.Context, code string) ([]*model.DictItem, error)
}

type emoService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewEmoService(
	db *gorm.DB,
	log *log.LoggerWithContext) EmoService {

	return &emoService{db: db, log: log}
}

func ProvideEmoService(container *dig.Container) {
	if err := container.Provide(NewEmoService); err != nil {
		panic(err)
	}
}

func (s *emoService) Create(c *gin.Context, body *dto.EmotionLog) error {
	layout := "2006-01-02 15:04:05"
	date, err := time.Parse(layout, body.Date)
	if err != nil {
		return fmt.Errorf("日期格式错误: %w", err)
	}
	data := model.EmotionLog{
		Title:   body.Title,
		Content: body.Content,
		UserID:  &body.UserID,
		Date:    date,
		Emo:     body.Emo,
	}
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *emoService) Update(c *gin.Context, id uint64, body *dto.EmotionLog) error {
	layout := "2006-01-02 15:04:05"
	date, err := time.Parse(layout, body.Date)
	if err != nil {
		return fmt.Errorf("日期格式错误: %w", err)
	}
	data := model.EmotionLog{
		Title:   body.Title,
		Content: body.Content,
		Date:    date,
		Emo:     body.Emo,
	}
	if err := s.db.Model(&model.EmotionLog{
		ID: id,
	}).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *emoService) Delete(c *gin.Context, body dto.DeleteIds) error {
	if err := s.db.Delete(&model.EmotionLog{}, body.Ids).Error; err != nil {
		return err
	}
	return nil
}

func (s *emoService) List(
	context *gin.Context,
	query dto.ListQuery,
	userId uint64) (dto.Result[dto.List[model.EmotionLog]], error) {
	var emotionLogs []model.EmotionLog
	limit := query.PageSize
	offset := query.PageNum*query.PageSize - query.PageSize

	if result := s.db.
		Model(&model.EmotionLog{}).
		Where("user_id = ?", userId).
		Limit(limit).
		Offset(offset).
		Order("create_time asc").
		Find(&emotionLogs); result.Error != nil {
		return dto.ServiceFail[dto.List[model.EmotionLog]](result.Error), result.Error
	}
	var count int64
	if result := s.db.Model(&model.EmotionLog{}).Count(&count); result.Error != nil {
		return dto.ServiceFail[dto.List[model.EmotionLog]](result.Error), result.Error
	}

	data := dto.ServiceSuccess(dto.List[model.EmotionLog]{
		Items:    emotionLogs,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    count,
	})
	return data, nil
}
