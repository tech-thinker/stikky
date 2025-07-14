package ui

import (
	"github.com/getlantern/systray"
	"github.com/tech-thinker/stikky/config"
	"github.com/tech-thinker/stikky/menu"
)

func RunSystray(cfg config.AppConfig) {
	st := menu.NewSystemTry(cfg)
	systray.Run(st.OnReady, st.OnExit)
}
