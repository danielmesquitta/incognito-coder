package app

import (
	"context"

	"github.com/danielmesquitta/incognito-coder/internal/config/env"
	"github.com/danielmesquitta/incognito-coder/internal/pkg/validator"
	"github.com/sashabaranov/go-openai"
)

type App struct {
	ctx             context.Context
	openaiClient    *openai.Client
	screenshots     []string
	currentLanguage string
}

func newApp(e *env.Env) *App {
	return &App{
		screenshots:     []string{},
		currentLanguage: "golang",
		openaiClient:    openai.NewClient(e.OpenAIAPIKey),
	}
}

func (a *App) Run(ctx context.Context) {
	a.ctx = ctx
	go a.registerKeyShortcuts()
}

func New() *App {
	val := validator.New()
	env := env.NewEnv(val)
	return newApp(env)
}
