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
	//7DD3FC
	usernameColor = color.NRGBA{R: 0x7D, G: 0xD3, B: 0xFC, A: 0xFF}
	textColor     = color.NRGBA{R: 0xEF, G: 0xF6, B: 0xFF, A: 0xFF}
)

type Home struct {
	OnConfigTap, OnNextTap, OnStopTap func()
}

func (h *Home) Content() fyne.CanvasObject {

	return container.NewVBox(
		components.ToolbarLayout(nil, h.OnConfigTap),
		chatPart(),
		layout.NewSpacer(),
		playerButtonsLayout(h.OnStopTap, h.OnNextTap),
	)
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

func chatPart() *fyne.Container {
	//EFF6FF

	username := canvas.NewText("dannywolfmx2", usernameColor)
	username.TextSize = 24
	username.TextStyle.Bold = true

	//Text from the user
	text := canvas.NewText("Prueba", textColor)
	text.TextSize = 20
	text.TextStyle.Bold = true

	spacer := canvas.NewRectangle(color.Transparent)

	left := canvas.NewText("    ", usernameColor)
	left.TextSize = 24
	return fyne.NewContainerWithLayout(
		layout.NewAdaptiveGridLayout(1),
		spacer,
		spacer,
		spacer,
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(spacer, nil, left, nil),
			left,
			username,
		),
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(spacer, nil, left, nil),
			left,
			text,
		),
	)
}
