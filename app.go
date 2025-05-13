package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
	"github.com/sashabaranov/go-openai"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const tmpDir = "tmp"

// App struct
type App struct {
	ctx             context.Context
	openaiClient    *openai.Client
	screenshots     []string
	currentLanguage string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		screenshots:     make([]string, 0),
		currentLanguage: "golang",
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic(".env file not found")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY not set")
	}

	go a.registerKeyShortcuts()

	a.openaiClient = openai.NewClient(apiKey)
}

func (a *App) registerKeyShortcuts() {
	hook.Register(
		hook.KeyDown,
		[]string{"p", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(a.ctx, "global-shortcut", "screenshot")
		},
	)

	hook.Register(
		hook.KeyDown,
		[]string{"enter", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(a.ctx, "global-shortcut", "generate")
		},
	)

	hook.Register(
		hook.KeyDown,
		[]string{"r", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(a.ctx, "global-shortcut", "reset")
		},
	)

	s := hook.Start()
	<-hook.Process(s)
}

// CaptureScreenshot captures the screen and saves it
func (a *App) CaptureScreenshot() error {
	if len(a.screenshots) >= 5 {
		return fmt.Errorf("maximum number of screenshots (5) reached")
	}

	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return err
	}

	// Create screenshots directory if it doesn't exist
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return err
	}

	// Generate unique filename
	filename := filepath.Join(
		tmpDir,
		fmt.Sprintf("screenshot_%d.png", time.Now().Unix()),
	)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return err
	}

	a.screenshots = append(a.screenshots, filename)
	return nil
}

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
		"You are a helpful programming assistant. Analyze the provided screenshots and generate a solution in %s programming language.",
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

// SetLanguage sets the current programming language
func (a *App) SetLanguage(lang string) {
	a.currentLanguage = lang
}

// Reset clears all screenshots and resets the state
func (a *App) Reset() error {
	for _, screenshot := range a.screenshots {
		if err := os.Remove(screenshot); err != nil {
			return err
		}
	}
	a.screenshots = make([]string, 0)
	return nil
}
