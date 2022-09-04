package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dannywolfmx/twitch-chat-voice/view/custom"
)

var (
	//7DD3FC
	usernameColor = color.NRGBA{R: 0x7D, G: 0xD3, B: 0xFC, A: 0xFF}
	textColor     = color.NRGBA{R: 0xEF, G: 0xF6, B: 0xFF, A: 0xFF}
)

type View interface {
	ShowAndRun()
	Quit()
}

type viewFyne struct {
	mainApp fyne.App
	fyne.Window
	config chan struct{}
}

func NewView(config chan struct{}) *viewFyne {
	gui := app.New()
	gui.Settings().SetTheme(&CustomTheme{})

	w := gui.NewWindow("Twitch app")

	w.SetContent(contentLayout(config))

	active := true
	go func() {
		for range config {
			if !active {
				w.SetContent(contentLayout(config))
			} else {
				w.SetContent(widget.NewButton("Hola", func() {
					config <- struct{}{}
				}))
			}
			active = !active
		}
	}()

	w.Resize(fyne.NewSize(400, 736))

	return &viewFyne{
		Window:  w,
		mainApp: gui,
		config:  config,
	}
}

func (v *viewFyne) Quit() {
	v.mainApp.Quit()
}

func contentLayout(config chan struct{}) *fyne.Container {

	a := container.NewVBox(
		toolbarLayout(config),
		chatPart(),
		layout.NewSpacer(),
		playerButtonsLayout(),
	)
	return a
}

func toolbarLayout(config chan struct{}) *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MenuIcon(), func() {
			config <- struct{}{}
		}),
	)
}

// ref: https://stackoverflow.com/questions/60560906/how-make-expanded-and-stretched-layout-box-with-fyne
func playerButtonsLayout() *fyne.Container {
	b := custom.NewCustomButton(theme.MediaStopIcon(), nil)

	b2 := custom.NewCustomButton(theme.MediaFastForwardIcon(), nil)

	return fyne.NewContainerWithLayout(
		layout.NewAdaptiveGridLayout(2),
		b,
		b2,
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
