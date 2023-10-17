package service

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type RecommendationInterface interface {
	GetRecommendation(request string) (string, error)
}

type RecommendationModel struct {
	OpenAPIKey string
}

func NewRecommendation(OpenAPIKey string) RecommendationInterface {
	return &RecommendationModel{OpenAPIKey: OpenAPIKey}
}

func (service *RecommendationModel) GetRecommendation(input string) (string, error) {
	context := context.Background()
	openAPIKey := service.OpenAPIKey
	client := openai.NewClient(openAPIKey)

	res, err := client.CreateChatCompletion(context, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: "you're here to help people choose their laptops from your recommendations."},
			{Role: openai.ChatMessageRoleUser, Content: input},
		},
	})

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v\n", err)
	}

	return res.Choices[0].Message.Content, nil
}
