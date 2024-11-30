package components

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// EditorComponent renders the text box.
func EditorComponent(th *material.Theme, editor *widget.Editor, text string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return material.Editor(th, editor, text).Layout(gtx)
	}
}
