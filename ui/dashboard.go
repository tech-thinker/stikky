package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tech-thinker/stikky/ui/panes"
)

type Dashboard interface {
	GetWindow() fyne.Window
}

type dashboard struct {
	a        fyne.App
	menuPane *fyne.Container
	bodyPane *fyne.Container
}

func NewDashboard(a fyne.App) Dashboard {
	return &dashboard{
		a:        a,
		menuPane: container.NewVBox(),
		bodyPane: container.NewVBox(),
	}
}

func (screen *dashboard) SetMenuLayout(objs ...fyne.CanvasObject) {
	screen.menuPane.Objects = objs
}

func (screen *dashboard) SetBodyLayout(objs ...fyne.CanvasObject) {
	screen.bodyPane.Objects = objs
}

func (screen *dashboard) GetWindow() fyne.Window {
	w := screen.a.NewWindow("Demo")
	w.Resize(fyne.NewSize(800, 600))

	menuHome := widget.NewButton("Home", func() {
		home := panes.NewHomePane()
		screen.SetBodyLayout(home.Draw())
	})

	menuPublicKeys := widget.NewButton("Public Keys", func() {
		publicKey := panes.NewPublicKeyPane()
		screen.SetBodyLayout(publicKey)

	})

	menuSettings := widget.NewButton("Settings", func() {
		setting := panes.NewSettingPane()
		screen.SetBodyLayout(setting.Draw())
	})
	// leftSideDrawer := container.NewVBox()
	screen.SetMenuLayout(menuHome, menuPublicKeys, menuSettings)
	// SetBodyLayout()

	vs := container.NewVScroll(screen.bodyPane)
	content := container.NewHSplit(screen.menuPane, vs)
	content.SetOffset(0.25)
	w.SetContent(content)

	return w
}
