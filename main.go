package main

import (
	"os"

	"github.com/tech-thinker/stikky/config"
	"github.com/tech-thinker/stikky/ui"
)

func main() {
	cfg := config.NewAppConfig()

	isFirstTime := false
	args := os.Args
	uiMode := false

	if len(args) >= 2 {
		uiMode = args[1] == "--ui"
	}

	if isFirstTime {
		uiMode = true
	}

	if uiMode {
		ui.RunWindow(cfg)
	} else {
		ui.RunSystray(cfg)
	}
}
