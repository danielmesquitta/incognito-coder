package usecase

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/danielmesquitta/incognito-coder/internal/domain/entity"
	"github.com/sashabaranov/go-openai"
)

type GenerateSolution struct {
	o *openai.Client
}

func NewGenerateSolution(o *openai.Client) *GenerateSolution {
	return &GenerateSolution{
		o: o,
	}
}

// GenerateSolution generates a solution based on the captured screenshots
func (g *GenerateSolution) Execute(
	ctx context.Context,
	screenshots []string,
	currentLanguage string,
) (*entity.Solution, error) {
	if len(screenshots) == 0 {
		return nil, fmt.Errorf("no screenshots available")
	}

	if g.o == nil {
		return nil, fmt.Errorf("OpenAI client not initialized")
	}

	// Create a message for the system
	systemMessage := fmt.Sprintf(`
			You are a LeetCode expert. Analyze the provided screenshots and generate a solution in %s programming language, return your answer as a JSON object with exactly these fields:

			- **code**: A complete, working Go implementation.
			- **time_complexity**: A brief Big-O analysis of your solution’s time complexity.
			- **space_complexity**: A brief Big-O analysis of your solution’s auxiliary space usage.
			- **thoughts**: A step-by-step explanation of your approach and trade-offs considered.
			
			Remember to return the JSON object only, without any other text or comments.
		`,
		currentLanguage,
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
	}

	// Add each screenshot as an image message
	for _, screenshot := range screenshots {
		imageData, err := os.ReadFile(screenshot)
		if err != nil {
			return nil, fmt.Errorf("failed to read screenshot: %v", err)
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

	// Make the API call
	resp, err := g.o.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate solution: %v", err)
	}

	// Return the generated solution
	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no solution generated")
	}

	solution := &entity.Solution{}
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), solution)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal solution: %v", err)
	}

	if solution.Code == "" ||
		solution.TimeComplexity == "" ||
		solution.SpaceComplexity == "" ||
		solution.Thoughts == "" {
		return nil, fmt.Errorf("invalid solution generated")
	}

	return solution, nil
}
