package view

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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

	a := container.NewVBox(
		layout.NewSpacer(),
		playerButtonsLayout(),
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

// ref: https://stackoverflow.com/questions/60560906/how-make-expanded-and-stretched-layout-box-with-fyne
func playerButtonsLayout() *fyne.Container {
	//b := custom.NewCustomButton(theme.MediaStopIcon(), nil)

	//b2 := custom.NewCustomButton(theme.MediaFastForwardIcon(), nil)
	//b2.Resize(fyne.NewSize(200, 200))

	b := widget.NewButtonWithIcon("", theme.MediaFastForwardIcon(), nil)
	b2 := widget.NewButtonWithIcon("", theme.MediaFastForwardIcon(), nil)

	container.NewGridWithColumns(cols int, objects ...fyne.CanvasObject)
	return fyne.NewContainerWithLayout(
		layout.NewAdaptiveGridLayout(2),
		b,
		b2,
	)
}
