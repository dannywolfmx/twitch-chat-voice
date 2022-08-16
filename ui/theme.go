package ui

import (
	"image/color"

	"gioui.org/text"
	"gioui.org/widget/material"
)

type Theme struct {
	*material.Theme
	TextColor   color.NRGBA
	ButtonColor color.NRGBA
}

func NewTheme(font []text.FontFace) *Theme {
	materialTheme := material.NewTheme(font)

	return &Theme{
		Theme: materialTheme,
	}
}
