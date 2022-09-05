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
	SetChatMessage(sender, message string)
	Quit()
}

//Fyne.Window will give us the ShowAndRun method
type viewFyne struct {
	mainApp fyne.App
	fyne.Window
	changeScreen  func(screen int)
	currentScreen int
}

func NewView(config ConfigView) *viewFyne {

	gui := app.New()
	gui.Settings().SetTheme(&CustomTheme{})

	w := gui.NewWindow("Twitch app")
	w.Resize(fyne.NewSize(400, 736))

	if config.DefaultScreen == NONE_SCREEN {
		config.DefaultScreen = HOME_SCREEN
	}

	v := &viewFyne{
		Window:        w,
		mainApp:       gui,
		changeScreen:  SetScreens(w, route),
		currentScreen: config.DefaultScreen,
	}

	route[HOME_SCREEN] = &screens.Home{
		OnConfigTap: config.OnConfigTap,
		OnStopTap:   config.OnStopTap,
		OnNextTap:   config.OnNextTap,
	}

	route[CONFIG_SCREEN] = &screens.Config{
		OnBackButton: func() {
			v.ChangeScreen(HOME_SCREEN)
		},
		OnUserNameChange: config.OnUserNameChange,
	}

	v.ChangeScreen(config.DefaultScreen)

	return v

}

func (v *viewFyne) SetChatMessage(sender, message string) {
	if v.currentScreen != HOME_SCREEN {
		return
	}
	s, ok := route[HOME_SCREEN].(*screens.Home)

	if !ok {
		//error on cast
		return
	}

	s.SetChatMessage(sender, message)
}

func (v *viewFyne) ChangeScreen(screen int) {
	v.currentScreen = screen
	v.changeScreen(screen)
}

func (v *viewFyne) Quit() {
	v.mainApp.Quit()
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
