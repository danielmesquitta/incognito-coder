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
	// Screenshot: Ctrl+Alt+P
	hook.Register(
		hook.KeyDown,
		[]string{"p", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "screenshot")
		},
	)

	// Generate solution: Ctrl+Alt+Enter
	hook.Register(
		hook.KeyDown,
		[]string{"enter", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "generate")
		},
	)

	// Reset: Ctrl+Alt+R
	hook.Register(
		hook.KeyDown,
		[]string{"r", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "reset")
		},
	)

	// Toggle visibility: Ctrl+Alt+V
	hook.Register(
		hook.KeyDown,
		[]string{"v", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "toggle-visibility")
		},
	)

	// Window movement: Ctrl+Alt+Arrow keys
	hook.Register(
		hook.KeyDown,
		[]string{"left", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "move-left")
		},
	)

	hook.Register(
		hook.KeyDown,
		[]string{"right", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "move-right")
		},
	)

	hook.Register(
		hook.KeyDown,
		[]string{"up", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "move-up")
		},
	)

	hook.Register(
		hook.KeyDown,
		[]string{"down", "ctrl", "alt"},
		func(e hook.Event) {
			runtime.EventsEmit(ctx, "global-shortcut", "move-down")
		},
	)

	s := hook.Start()
	<-hook.Process(s)
}
