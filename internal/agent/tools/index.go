package tools

import (
	"encoding/json"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"go.uber.org/zap"
)

var AvailableTools = []llms.Tool{
	{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "GetCityPinyin",
			Description: "获取城市拼音",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name": map[string]any{
						"type":        "string",
						"description": "城市名称或者区县等名称，如广州，佛山，南海，佛山市南海区则取南海，匹配相关的区等",
					},
				},
				"required": []string{"name"},
			},
		},
	},
	{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "GetCityIDs",
			Description: "获取城市locationId",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"pinyin": map[string]any{
						"type":        "string",
						"description": "获取城市的locationId，根据城市拼音获取城市ID",
					},
				},
				"required": []string{"pinyin"},
			},
		},
	},
	{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "GetCurrentWeather",
			Description: "获取当前天气",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"location": map[string]any{
						"type":        "string",
						"description": "传递城市ID",
					},
				},
				"required": []string{"location"},
			},
		},
	},
}

func ExecuteToolCalls(
	messageHistory []llms.MessageContent,
	resp *llms.ContentResponse,
	logger *zap.Logger) []llms.MessageContent {

	// 维护一个累积器
	//var currentToolName string
	//var currentToolID string
	//var argsBuilder strings.Builder

	for _, choice := range resp.Choices {
		for _, toolCall := range choice.ToolCalls {
			// 记录 tool_call
			assistantResponse := llms.MessageContent{
				Role: llms.ChatMessageTypeAI,
				Parts: []llms.ContentPart{
					llms.ToolCall{
						ID:   toolCall.ID,
						Type: toolCall.Type,
						FunctionCall: &llms.FunctionCall{
							Name:      toolCall.FunctionCall.Name,
							Arguments: toolCall.FunctionCall.Arguments,
						},
					},
				},
			}
			messageHistory = append(messageHistory, assistantResponse)

			//// 如果 Function 名字为空，表示是上一个函数的 continuation
			//if toolCall.FunctionCall.Name != "" {
			//	currentToolName = toolCall.FunctionCall.Name
			//	currentToolID = toolCall.ID
			//	argsBuilder.Reset()
			//}
			//argsBuilder.WriteString(toolCall.FunctionCall.Arguments)

			// 当我们检测到 JSON 可能结束，才尝试 unmarshal
			if strings.HasSuffix(toolCall.FunctionCall.Arguments, "}") {
				//argsJSON := argsBuilder.String()
				switch toolCall.FunctionCall.Name {
				case "GetCurrentWeather":
					var args struct {
						Location string `json:"location"`
					}
					if err := json.Unmarshal([]byte(toolCall.FunctionCall.Arguments), &args); err != nil {
						logger.Error("Unmarshal GetCurrentWeather failed", zap.Error(err))
						continue
					}
					response, err := GetCurrentWeather(args.Location)
					if err != nil {
						logger.Error(err.Error())
					}
					weatherCallResponse := llms.MessageContent{
						Role: llms.ChatMessageTypeTool,
						Parts: []llms.ContentPart{
							llms.ToolCallResponse{
								ToolCallID: toolCall.ID,
								Name:       toolCall.FunctionCall.Name,
								Content:    response,
							},
						},
					}
					messageHistory = append(messageHistory, weatherCallResponse)

				case "GetCityPinyin":
					var args struct {
						Name string `json:"name"`
					}
					if err := json.Unmarshal([]byte(toolCall.FunctionCall.Arguments), &args); err != nil {
						logger.Error("Unmarshal GetCityPinyin failed", zap.Error(err))
						continue
					}
					response := GetCityPinyin(args.Name)
					weatherCallResponse := llms.MessageContent{
						Role: llms.ChatMessageTypeTool,
						Parts: []llms.ContentPart{
							llms.ToolCallResponse{
								ToolCallID: toolCall.ID,
								Name:       toolCall.FunctionCall.Name,
								Content:    response,
							},
						},
					}
					messageHistory = append(messageHistory, weatherCallResponse)

				case "GetCityIDs":
					var args struct {
						Pinyin string `json:"pinyin"`
					}
					if err := json.Unmarshal([]byte(toolCall.FunctionCall.Arguments), &args); err != nil {
						logger.Error("Unmarshal GetCityIDs failed", zap.Error(err))
						continue
					}
					response, err := GetCityIDs(args.Pinyin)
					if err != nil {
						logger.Error(err.Error())
					}
					responseJSON, _ := json.Marshal(response)
					weatherCallResponse := llms.MessageContent{
						Role: llms.ChatMessageTypeTool,
						Parts: []llms.ContentPart{
							llms.ToolCallResponse{
								ToolCallID: toolCall.ID,
								Name:       toolCall.FunctionCall.Name,
								Content:    string(responseJSON),
							},
						},
					}
					messageHistory = append(messageHistory, weatherCallResponse)

				default:
					logger.Warn("Unsupported tool", zap.String("tool", toolCall.FunctionCall.Name))
				}
			}
		}
	}
	return messageHistory
}
