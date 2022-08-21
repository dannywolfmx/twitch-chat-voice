package component

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/widget/material"
)

type ComponentTheme struct {
	*material.Theme
	TextColor   color.NRGBA
	ButtonColor color.NRGBA
}

var DefaultTheme = NewTheme(gofont.Collection())

func NewTheme(font []text.FontFace) *ComponentTheme {
	material := material.NewTheme(font)
	material.Bg = NewColor(0x191E38FF)
	material.Fg = NewColor(0x2F365FFF)
	material.ContrastBg = NewColor(0x5661B3FF)

	return &ComponentTheme{
		Theme:     material,
		TextColor: NewColor(0xE6E8FFFF),
	}
}
