package ui

import (
	"image/color"
)

func NewColor(hexa uint32) color.NRGBA {
	c := color.NRGBA{
		R: uint8(hexa >> 24),
		G: uint8(hexa >> 16),
		B: uint8(hexa >> 8),
		A: uint8(hexa),
	}

	return c
}
