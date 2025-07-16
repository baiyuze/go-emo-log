package agent

import (
	"log"
	"os"

	"github.com/tmc/langchaingo/llms/openai"
	"go.uber.org/dig"
)

func ProvideAi(container *dig.Container) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	llm, err := openai.New(
		openai.WithToken(apiKey),

		openai.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1"), // 注意带 /v1
		openai.WithModel("deepseek-r1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := container.Provide(func() *openai.LLM {
		return llm
	}); err != nil {
		log.Fatal(err)
	}

}
