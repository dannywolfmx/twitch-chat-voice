package controller

import (
	"fmt"

	"github.com/dannywolfmx/twitch-chat-voice/view/screens"
)

type configController struct {
	screen *screens.Config
	//Client          *twitch.Client
	routeHomeScreen func()
}

func NewConfigController(routeHomeScreen func()) *configController {

	controller := &configController{
		routeHomeScreen: routeHomeScreen,
	}

	controller.initScreen()

	return controller
}

// Screen (view) of the actual controller
func (c *configController) Screen() screens.Screen {
	return c.screen
}

// initScreen will generate the screen of the view
func (c *configController) initScreen() {
	c.screen = &screens.Config{
		OnUserNameChange: func(name string) {
			fmt.Println(name)
		},
		OnBackButton: c.routeHomeScreen,
	}
}
