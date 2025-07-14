package components

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// TextFieldComponent renders the text box.
func TextFieldComponent(th *material.Theme, editor *widget.Editor, text string, readonly bool) layout.Widget {
	editor.ReadOnly = readonly
	return func(gtx layout.Context) layout.Dimensions {
		return material.Editor(th, editor, text).Layout(gtx)
	}
}
