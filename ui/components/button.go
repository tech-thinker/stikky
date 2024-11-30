package components

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ButtonComponent renders the button and handles click events.
func ButtonComponent(th *material.Theme, btn *widget.Clickable, text string, event func()) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		if btn.Clicked(gtx) {
			event()
		}
		return material.Button(th, btn, text).Layout(gtx)
	}
}
