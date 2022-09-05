package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dannywolfmx/twitch-chat-voice/view/screens/components"
)

type Config struct {
	OnUserNameChange func(string)
	OnBackButton     func()
}

func (c *Config) Content() fyne.CanvasObject {
	return container.NewVBox(
		components.ToolbarLayout(c.OnBackButton, nil),
		c.inputUserName(),
		layout.NewSpacer(),
	)
}

func (c *Config) inputUserName() fyne.CanvasObject {
	input := widget.NewEntry()
	input.SetPlaceHolder("Twitch channel")
	input.Resize(fyne.NewSize(200, 200))

	return container.NewVBox(input, widget.NewButton("Conectar", func() {
		if input.Text != "" {
			c.OnUserNameChange(input.Text)
			input.SetText("")
		}
	}))
}
