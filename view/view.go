package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dannywolfmx/twitch-chat-voice/view/screens"
)

var route = make(map[int]screens.Screen, 0)

const (
	NONE_SCREEN = iota
	HOME_SCREEN
	CONFIG_SCREEN
)

type View interface {
	ShowAndRun()
	ChangeScreen(screen int)
	Quit()
}

//Fyne.Window will give us the ShowAndRun method
type viewFyne struct {
	mainApp fyne.App
	fyne.Window
	changeScreen func(screen int)
}

func NewView(config ConfigView) *viewFyne {

	gui := app.New()
	gui.Settings().SetTheme(&CustomTheme{})

	w := gui.NewWindow("Twitch app")
	w.Resize(fyne.NewSize(400, 736))

	route[HOME_SCREEN] = &screens.Home{
		OnConfigTap: config.OnConfigTap,
		OnStopTap:   config.OnStopTap,
		OnNextTap:   config.OnNextTap,
	}

	route[CONFIG_SCREEN] = &screens.Config{
		OnBackButton: func() {
			w.SetContent(route[HOME_SCREEN].Content())
		},
	}

	if config.DefaultScreen == NONE_SCREEN {
		config.DefaultScreen = HOME_SCREEN
	}

	w.SetContent(route[config.DefaultScreen].Content())

	return &viewFyne{
		Window:       w,
		mainApp:      gui,
		changeScreen: SetScreens(w, route),
	}
}

func SetScreens(window fyne.Window, screens map[int]screens.Screen) func(screen int) {
	return func(screen int) {
		s, ok := screens[screen]
		//No window
		if !ok {
			return
		}
		window.SetContent(s.Content())
	}
}

func (v *viewFyne) ChangeScreen(screen int) {
	v.changeScreen(screen)
}

func (v *viewFyne) Quit() {
	v.mainApp.Quit()
}
