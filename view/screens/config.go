package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	OnBackButton func()
}

func (c *Config) Content() fyne.CanvasObject {
	return widget.NewButton("Hola", c.OnBackButton)
}
