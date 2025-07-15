package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
)

var AvailableTools = []llms.Tool{
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
						"description": "城市名称，如：北京",
					},
				},
				"required": []string{"location"},
			},
		},
	},
}

func ExecuteToolCalls(ctx context.Context, llm llms.Model, messageHistory []llms.MessageContent, resp *llms.ContentResponse) []llms.MessageContent {
	for _, choice := range resp.Choices {
		for _, toolCall := range choice.ToolCalls {

			// Append tool_use to messageHistory
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
			fmt.Println(toolCall.FunctionCall.Name, "--->")
			switch toolCall.FunctionCall.Name {
			case "GetCurrentWeather":
				var args struct {
					Location string `json:"location"`
				}
				if err := json.Unmarshal([]byte(toolCall.FunctionCall.Arguments), &args); err != nil {
					log.Fatal(err)
				}

				// Perform Function Calling
				response, err := GetCurrentWeather(args.Location)
				if err != nil {
					log.Fatal(err)
				}

				// Append tool_result to messageHistory
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
			default:
				log.Fatalf("Unsupported tool: %s", toolCall.FunctionCall.Name)
			}
		}
	}

	return messageHistory
}
