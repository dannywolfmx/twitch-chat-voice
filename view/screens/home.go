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

	sender  *canvas.Text
	message *canvas.Text
)

type Home struct {
	OnConfigTap, OnNextTap, OnStopTap func()

	GetMessage func() (string, string)
}

func (h *Home) Content() fyne.CanvasObject {
	sender = canvas.NewText("", senderColor)
	message = canvas.NewText("", textColor)

	return container.NewVBox(
		components.ToolbarLayout(nil, h.OnConfigTap),
		chatPart(sender, message),
		layout.NewSpacer(),
		playerButtonsLayout(h.OnStopTap, h.OnNextTap),
	)
}

func (h *Home) SetChatMessage(s, m string) {
	sender.Text = s
	message.Text = m

	sender.Refresh()
	message.Refresh()
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

func chatPart(sender, message *canvas.Text) *fyne.Container {
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
