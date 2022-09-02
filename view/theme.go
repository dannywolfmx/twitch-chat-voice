package view

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	darkModeColor = color.NRGBA{
		R: 0x4F,
		G: 0x46,
		B: 0xE5,
		A: 0xFF,
	}
)

type CustomTheme struct{}

func (c *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return darkModeColor
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (c *CustomTheme) Font(textStyle fyne.TextStyle) fyne.Resource {
	fmt.Println(textStyle)
	return theme.DefaultTheme().Font(textStyle)
}

func (c *CustomTheme) Icon(iconName fyne.ThemeIconName) fyne.Resource {

	return theme.DefaultTheme().Icon(iconName)
}

func (c *CustomTheme) Size(sizeName fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(sizeName)
}
