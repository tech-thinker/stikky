package components

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ImageButtonComponent renders the button and handles click events.
func ImageButtonComponent(th *material.Theme, btn *widget.Clickable, icon *widget.Icon, text string, event func()) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		if btn.Clicked(gtx) {
			event()
		}
		ibs := material.IconButton(th, btn, icon, text)
		ibs.Size = 12
		return ibs.Layout(gtx)
	}
}
