package components

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// PasswordFieldComponent renders the text box.
func PasswordFieldComponent(th *material.Theme, editor *widget.Editor, text string, isMasked bool) layout.Widget {
	if isMasked {
		editor.Mask = rune('*')
	} else {
		editor.Mask = 0
	}
	return func(gtx layout.Context) layout.Dimensions {
		return material.Editor(th, editor, text).Layout(gtx)
	}
}
