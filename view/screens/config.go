package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dannywolfmx/twitch-chat-voice/view/screens/components"
)

// Config is the config section where the user can set stuffs like:
//
// - User name of the twitch channel
type Config struct {
	//Event when the user name box is changed
	OnUserNameChange func(name string)

	//Event when the back button is pressed
	OnBackButton func()
}

func (c *Config) Content() fyne.CanvasObject {
	return container.NewVBox(
		components.ToolbarLayout(c.OnBackButton, nil),
		c.inputUserName(),
		layout.NewSpacer(),
	)
}

func (c *Config) Update() {

}

func (c *Config) inputUserName() fyne.CanvasObject {
	input := widget.NewEntry()
	input.SetPlaceHolder("Twitch channel")
	input.Resize(fyne.NewSize(200, 200))
	input.OnSubmitted = func(text string) {
		c.OnUserNameChange(text)
		input.SetText("")
	}

	return container.NewVBox(input, widget.NewButton("Conectar", func() {
		if input.Text != "" {
			c.OnUserNameChange(input.Text)
			input.SetText("")
		}
	}))
}
