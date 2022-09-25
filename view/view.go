package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type View interface {
	fyne.Window
	Quit()
}

//Fyne.Window will give us the ShowAndRun method
type viewFyne struct {
	fyne.Window

	mainApp fyne.App
}

func NewView(config ConfigView) *viewFyne {

	gui := app.New()
	gui.Settings().SetTheme(&CustomTheme{})

	w := gui.NewWindow(config.Title)
	w.Resize(config.WindowSize)

	return &viewFyne{
		Window:  w,
		mainApp: gui,
	}
}

func (v *viewFyne) Quit() {
	v.mainApp.Quit()
}
