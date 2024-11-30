package ui

import (
	"fmt"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/tech-thinker/stikky/ui/components"
)

var (
	theme *material.Theme
	quit  = &widget.Clickable{}
)

func RunWindow() {
	go func() {
		dashboard()
	}()
	app.Main()
}

func dashboard() {
	w := new(app.Window)
	w.Option(app.Title("Stikky"))

	// Theme
	theme = material.NewTheme()

	// Create a text editor widget
	var editor widget.Editor
	editor.SingleLine = true // Make it a single-line editor

	// Button widget
	var btn1, btn2 widget.Clickable

	// Main event loop
	ops := new(op.Ops)
	for {
		e := w.Event()
		switch e := e.(type) {
		case app.FrameEvent:

			gtx := app.NewContext(ops, e)

			// Layout and draw UI
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Stack layout to show the text box and button
				return layout.Flex{
					Axis:    layout.Vertical, // Arrange items vertically
					Spacing: layout.SpaceEnd,
				}.Layout(gtx,
					// Text Box
					layout.Rigid(components.EditorComponent(theme, &editor, "Type here...")),

					// Spacer
					layout.Rigid(components.VerticalSpacer(theme)),

					// Button
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal, // Arrange items horizontally
							Spacing: layout.SpaceEnd,
						}.Layout(gtx,
							layout.Rigid(components.ButtonComponent(theme, &btn1, "Click me 1", func() {
								data := editor.Text()
								fmt.Println("Button clicked! Text 1:", data)
							})),
							layout.Rigid(components.HorizontalSpacer(theme)),
							layout.Rigid(components.ButtonComponent(theme, &btn2, "Click me 2", func() {
								data := editor.Text()
								fmt.Println("Button clicked! Text 2:", data)
							})),
						)
					}),
				)
			})

			e.Frame(gtx.Ops)

		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}
