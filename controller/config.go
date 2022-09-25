package controller

import (
	"fmt"

	"github.com/dannywolfmx/twitch-chat-voice/view/screens"
	"github.com/gempir/go-twitch-irc/v3"
)

type configController struct {
	screen          *screens.Config
	Client          *twitch.Client
	routeHomeScreen func()
}

func NewConfigController(routeHomeScreen func(), c *twitch.Client) *configController {

	controller := &configController{
		routeHomeScreen: routeHomeScreen,
		Client:          c,
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
			c.Client.Join(name)
			fmt.Println(name)
		},
		OnBackButton: c.routeHomeScreen,
	}
}
