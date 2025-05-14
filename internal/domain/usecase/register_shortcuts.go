package usecase

import (
	"context"

	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type RegisterShortcuts struct{}

func NewRegisterShortcuts() *RegisterShortcuts {
	return &RegisterShortcuts{}
}

func (r *RegisterShortcuts) Execute(ctx context.Context) {
	hook.Register(
		hook.KeyDown,
		[]string{"p", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "screenshot")
		},
	)

	hook.Register(
		hook.KeyDown,
		[]string{"enter", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "generate")
		},
	)

	hook.Register(
		hook.KeyDown,
		[]string{"r", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "reset")
		},
	)

	s := hook.Start()
	<-hook.Process(s)
}
