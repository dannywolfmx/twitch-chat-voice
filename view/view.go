package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
	w.SetContent(widget.NewLabel("Hello World!"))

	return &viewFyne{
		Window:  w,
		mainApp: gui,
	}
}

func (v *viewFyne) Quit() {
	v.mainApp.Quit()
}
