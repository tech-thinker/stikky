package panes

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type HomePane interface {
	Draw() *fyne.Container
}

type homePane struct {
	pane *fyne.Container
}

func (p *homePane) Draw() *fyne.Container {
	// vb := container.NewVBox(
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 1")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 2")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 3")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 4")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 5")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 6")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 7")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 8")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 9")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 10")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 11")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 12")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 13")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 13")),
	// 	widget.NewCard("title", "", widget.NewLabel("Home page 15")),
	// )
	vb1 := canvas.NewCircle(color.RGBA{255, 255, 255, 255})
	vb1.Resize(fyne.NewSize(100, 100))
	vb2 := canvas.NewCircle(color.RGBA{255, 255, 255, 255})
	vb2.Resize(fyne.NewSize(100, 100))
	vb2.Move(fyne.NewPos(0, 110))
	x := container.NewWithoutLayout(vb1, vb2)
	go func() {
		FPS := 60
		t := time.Second / time.Duration(FPS)
		ticker := time.NewTicker(t)

		for range ticker.C {
			x := vb2.Position1.X + 5

			vb2.Move(fyne.NewPos(x, vb2.Position1.Y))

			if x > 500 {
				vb2.Move(fyne.NewPos(0, 110))
			}

			// 👇 THIS is the key in Fyne v2
			canvas.Refresh(vb2)
		}

	}()

	p.pane.Add(x)
	return p.pane
}

func NewHomePane() HomePane {
	return &homePane{
		pane: container.NewStack(),
	}
}
