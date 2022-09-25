package screens

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/dannywolfmx/twitch-chat-voice/view/custom"
	"github.com/dannywolfmx/twitch-chat-voice/view/screens/components"
)

var (
	senderColor = color.NRGBA{R: 0x7D, G: 0xD3, B: 0xFC, A: 0xFF}
	textColor   = color.NRGBA{R: 0xEF, G: 0xF6, B: 0xFF, A: 0xFF}
)

type Home struct {
	OnConfigTap, OnNextTap, OnStopTap func()

	GetMessage func() (string, string)

	sender, message *canvas.Text
}

// Content is the main output to the user
func (h *Home) Content() fyne.CanvasObject {
	sender, message := h.GetMessage()

	h.sender = canvas.NewText(sender, senderColor)
	h.message = canvas.NewText(message, textColor)

	return container.NewVBox(
		components.ToolbarLayout(nil, h.OnConfigTap),
		chat(h.sender, h.message),
		layout.NewSpacer(),
		playerButtonsLayout(h.OnStopTap, h.OnNextTap),
	)
}

// Update Generate a call to the screan to update the
// sender and message text fields
func (h *Home) Update() {
	h.sender.Text, h.message.Text = h.GetMessage()
	h.sender.Refresh()
	h.message.Refresh()
}

// ref: https://stackoverflow.com/questions/60560906/how-make-expanded-and-stretched-layout-box-with-fyne
func playerButtonsLayout(OnStopTap, OnNextTap func()) *fyne.Container {
	stopButton := custom.NewCustomButton(theme.MediaStopIcon(), nil)
	stopButton.OnTapped = OnStopTap

	nextButton := custom.NewCustomButton(theme.MediaFastForwardIcon(), nil)
	nextButton.OnTapped = OnNextTap

	return fyne.NewContainerWithLayout(
		layout.NewAdaptiveGridLayout(2),
		stopButton,
		nextButton,
	)
}

// chat will fill the sender and message to the widget
func chat(sender, message *canvas.Text) *fyne.Container {
	//EFF6FF
	sender.Color = senderColor
	sender.TextSize = 24
	sender.TextStyle.Bold = true

	//Text from the user
	message.Color = textColor
	message.TextSize = 20
	message.TextStyle.Bold = true

	spacer := canvas.NewRectangle(color.Transparent)

	left := canvas.NewText("    ", senderColor)
	left.TextSize = 24
	return fyne.NewContainerWithLayout(
		layout.NewAdaptiveGridLayout(1),
		spacer,
		spacer,
		spacer,
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(spacer, nil, left, nil),
			left,
			sender,
		),
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(spacer, nil, left, nil),
			left,
			message,
		),
	)
}
