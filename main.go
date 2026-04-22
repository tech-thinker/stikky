package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/tech-thinker/stikky/res"
	"github.com/tech-thinker/stikky/ui"
)

func main() {
	// cfg := config.NewAppConfig()

	a := app.New()
	icon, err := res.GetIcon()
	if err != nil {
		panic(err)
	}
	a.SetIcon(fyne.NewStaticResource("icon", icon))
	screen := ui.NewDashboard(a)
	w := screen.GetWindow()
	w.SetIcon(fyne.NewStaticResource("icon", icon))
	w.Show()

	a.Run()
}
