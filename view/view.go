package view

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dannywolfmx/twitch-chat-voice/view/custom"
)

type View interface {
	ShowAndRun()
	Quit()
}

type viewFyne struct {
	mainApp fyne.App
	fyne.Window
}

func NewView() *viewFyne {
	gui := app.New()
	gui.Settings().SetTheme(&CustomTheme{})
	w := gui.NewWindow("Twitch app")

	w.SetContent(contentLayout())

	w.Resize(fyne.NewSize(400, 736))

	return &viewFyne{
		Window:  w,
		mainApp: gui,
	}
}

func (v *viewFyne) Quit() {
	v.mainApp.Quit()
}

func contentLayout() *fyne.Container {

	bo := custom.NewCustomButton(theme.HomeIcon(), func() {
		fmt.Println("Prueba")
	})
	bo.Resize(fyne.NewSize(200, 200))
	a := container.NewWithoutLayout(
		bo,
	)
	return a
}

func toolbarLayout() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MenuIcon(), func() {
			log.Println("Clicked s")
		}),
	)
}

func playerButtonsLayout() *fyne.Container {
	b := widget.NewButtonWithIcon("", theme.HomeIcon(), nil)
	b.Resize(fyne.NewSize(200, 200))

	b2 := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)
	n := container.New(
		layout.NewMaxLayout(),
		b,
	)

	n2 := container.New(
		layout.NewMaxLayout(),
		b2,
	)

	return container.New(
		layout.NewHBoxLayout(),
		n, n2,
	)
}
