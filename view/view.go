package view

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
	return container.NewVBox(
		toolbarLayout(),
		layout.NewSpacer(),
		container.NewMax(
			canvas.NewText("Prueba", theme.TextColor()),
		),
		layout.NewSpacer(),
		widget.NewIcon(theme.MediaPlayIcon()),
	)
}

func toolbarLayout() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MenuIcon(), func() {
			log.Println("Clicked")
		}),
	)
}
