package service

import (
	"context"
	"emoLog/internal/agent/tools"
	"emoLog/internal/common/log"
	"emoLog/internal/dto"
	"emoLog/internal/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type AiService interface {
	Create(c *gin.Context, body *dto.Feedback) error
	Update(c *gin.Context, id uint64, body *dto.Feedback) error
	Delete(c *gin.Context, body dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery, userId uint64) (dto.Result[dto.List[model.Feedback]], error)
	TestChat(c *gin.Context, msg string) []*llms.ContentChoice
	//Update(c *gin.Context, id int, body *model.Dict) error
	//GetOptionsByDictCode(c *gin.Context, code string) ([]*model.DictItem, error)
}

type aiService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
	llm *openai.LLM
}

func NewAiService(
	llm *openai.LLM,

	db *gorm.DB,
	log *log.LoggerWithContext) AiService {

	return &aiService{db: db, log: log, llm: llm}
}

func ProvideAiService(container *dig.Container) {
	if err := container.Provide(NewAiService); err != nil {
		panic(err)
	}
}

// TestChat 测试Ai
func (s *aiService) TestChat(c *gin.Context, msg string) []*llms.ContentChoice {
	ctx := context.Background()
	messageHistory := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "你是一个用来获取天气的智能助手，请用中文回答，多用表情符号，回答结尾需要，说下今天适合做什么"),
		llms.TextParts(llms.ChatMessageTypeHuman, msg),
	}
	// resp, err := s.llm.GenerateContent(ctx, messageHistory,
	// 	llms.WithTools(tools.AvailableTools),
	// 	llms.WithStreamingFunc(
	// 		func(ctx context.Context, chunk []byte) error {
	// 			fmt.Println(string(chunk))
	// 			return nil
	// 		}))
	resp, err := s.llm.GenerateContent(ctx, messageHistory, llms.WithTools(tools.AvailableTools))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Choices[0].Content, "----1")
	messageHistory = tools.ExecuteToolCalls(ctx, s.llm, messageHistory, resp)

	resp, err = s.llm.GenerateContent(ctx, messageHistory, llms.WithTools(tools.AvailableTools))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Choices[0].Content, "----2")
	// if err != nil {
	// 	// return err
	// 	fmt.Println(err)
	// }

	// fmt.Printf("---%+v---resp\n", resp.Choices[0].Content)

	return resp.Choices
}

func (s *aiService) Create(c *gin.Context, body *dto.Feedback) error {

	data := model.Feedback{
		Description: body.Description,
		Version:     body.Version,
	}
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *aiService) Update(c *gin.Context, id uint64, body *dto.Feedback) error {

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

func (s *aiService) Delete(c *gin.Context, body dto.DeleteIds) error {
	if err := s.db.Delete(&model.Feedback{}, body.Ids).Error; err != nil {
		return err
	}
	return nil
}

func (s *aiService) List(
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
