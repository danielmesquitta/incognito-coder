package app

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

// GenerateSolution generates a solution based on the captured screenshots
func (a *App) GenerateSolution() (string, error) {
	if len(a.screenshots) == 0 {
		return "", fmt.Errorf("no screenshots available")
	}

	if a.openaiClient == nil {
		return "", fmt.Errorf("OpenAI client not initialized")
	}

	// Create a message for the system
	systemMessage := fmt.Sprintf(
		"You are a LeetCode expert. Analyze the provided screenshots and generate a solution in %s programming language.",
		a.currentLanguage,
	)

	// Create the chat completion request
	req := openai.ChatCompletionRequest{
		Model: openai.O4Mini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemMessage,
			},
		},
		MaxTokens: 1000,
	}

	// Add each screenshot as an image message
	for _, screenshot := range a.screenshots {
		imageData, err := os.ReadFile(screenshot)
		if err != nil {
			return "", fmt.Errorf("failed to read screenshot: %v", err)
		}

		imageMessage := openai.ChatCompletionMessage{
			Role: openai.ChatMessageRoleUser,
			MultiContent: []openai.ChatMessagePart{
				{
					Type: openai.ChatMessagePartTypeImageURL,
					ImageURL: &openai.ChatMessageImageURL{
						URL: fmt.Sprintf(
							"data:image/png;base64,%s",
							base64.StdEncoding.EncodeToString(imageData),
						),
					},
				},
			},
		}
		req.Messages = append(req.Messages, imageMessage)
	}

	// Add a final message requesting the solution
	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "Please analyze these screenshots and provide a solution to the problem shown.",
	})

	// Make the API call
	resp, err := a.openaiClient.CreateChatCompletion(a.ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to generate solution: %v", err)
	}

	// Return the generated solution
	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no solution generated")
}
