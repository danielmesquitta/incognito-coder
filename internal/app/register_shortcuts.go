package app

import (
	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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
