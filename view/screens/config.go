package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/dannywolfmx/twitch-chat-voice/view/screens/components"
)

type Config struct {
	OnBackButton func()
}

func (c *Config) Content() fyne.CanvasObject {
	return container.NewVBox(
		components.ToolbarLayout(c.OnBackButton, nil),
		layout.NewSpacer(),
	)
}
