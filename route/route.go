package route

import (
	"errors"
	"fmt"
	"log"

	"github.com/dannywolfmx/twitch-chat-voice/controller"
	"github.com/dannywolfmx/twitch-chat-voice/view"
	"github.com/dannywolfmx/twitch-chat-voice/view/screens"
)

var (
	ERROR_ROUTE_NO_SCREEN = errors.New("no screen in route")
)

type Route struct {
	currentScreen string
	screens       map[string]controller.Controller
	window        view.View
	activeLog     bool
}

func NewRoute(window view.View, activeLog bool) *Route {
	return &Route{
		screens:   make(map[string]controller.Controller, 0),
		window:    window,
		activeLog: activeLog,
	}
}

func (r *Route) Set(name string, screen controller.Controller) {
	r.screens[name] = screen
}

func (r *Route) Go(name string) error {
	s, ok := r.screens[name]
	//No window
	if !ok {
		err := fmt.Errorf("%s: %s", ERROR_ROUTE_NO_SCREEN, name)
		if r.activeLog {
			log.Println(err)
		}

		return err
	}
	r.changeScreen(s.Screen())

	r.currentScreen = name
	return nil
}

func (r *Route) changeScreen(screen screens.Screen) {
	r.window.SetContent(screen.Content())
}
