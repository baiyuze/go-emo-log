package service

import (
	"emoLog/internal/common/log"
	"emoLog/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type EmoService interface {
	Create(c *gin.Context, body *model.Dict) error
	//Delete(c *gin.Context, body dto.DeleteIds) error
	//List(context *gin.Context, query dto.ListQuery, name string) (dto.Result[dto.List[model.Dict]], error)
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

func (s *emoService) Create(c *gin.Context, body *model.Dict) error {
	return nil
}
