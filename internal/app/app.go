package app

import (
	"context"
	"fmt"

	"github.com/danielmesquitta/incognito-coder/internal/config/env"
	"github.com/danielmesquitta/incognito-coder/internal/domain/entity"
	"github.com/danielmesquitta/incognito-coder/internal/domain/usecase"
	"github.com/danielmesquitta/incognito-coder/internal/pkg/validator"
	"github.com/sashabaranov/go-openai"
)

type App struct {
	ctx             context.Context
	screenshots     []string
	currentLanguage string

	o  *openai.Client
	cs *usecase.CaptureScreenshot
	gs *usecase.GenerateSolution
	rs *usecase.RegisterShortcuts
	r  *usecase.Reset
}

func newApp(e *env.Env) *App {
	openAIClient := openai.NewClient(e.OpenAIAPIKey)

	return &App{
		ctx:             nil,
		screenshots:     []string{},
		currentLanguage: "golang",
		o:               openAIClient,
		cs:              usecase.NewCaptureScreenshot(e),
		gs:              usecase.NewGenerateSolution(openAIClient),
		rs:              usecase.NewRegisterShortcuts(),
		r:               usecase.NewReset(),
	}
}

func (a *App) Run(ctx context.Context) {
	a.ctx = ctx
	go a.rs.Execute(a.ctx)
}

func New() *App {
	val := validator.New()
	env := env.NewEnv(val)
	return newApp(env)
}

func (a *App) SetLanguage(lang string) {
	a.currentLanguage = lang
}

func (a *App) CaptureScreenshot() error {
	screenshots, err := a.cs.Execute(a.screenshots)
	if err != nil {
		return fmt.Errorf("failed to capture screenshot: %w", err)
	}
	a.screenshots = screenshots

	return nil
}

func (a *App) GenerateSolution() (*entity.Solution, error) {
	solution, err := a.gs.Execute(a.ctx, a.screenshots, a.currentLanguage)
	if err != nil {
		return nil, fmt.Errorf("failed to generate solution: %w", err)
	}

	return solution, nil
}

func (a *App) RegisterShortcuts() {
	a.rs.Execute(a.ctx)
}

func (a *App) Reset() error {
	err := a.r.Execute(a.ctx, a.screenshots)
	if err != nil {
		return fmt.Errorf("failed to reset: %w", err)
	}
	a.screenshots = []string{}

	return nil
}
