package ui

import (
	"github.com/getlantern/systray"
	"github.com/tech-thinker/stikky/menu"
)

func RunSystray() {
	systray.Run(menu.OnReady, menu.OnExit)
}
