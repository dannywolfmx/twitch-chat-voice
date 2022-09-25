package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/view/screens"
	"github.com/gempir/go-twitch-irc/v3"
)

type homeController struct {
	screen          *screens.Home
	Player          *tts.TTS
	Client          *twitch.Client
	Sender, Message string
	goConfigScreen  func() error
}

func NewHomeController(goConfigScreen func() error, p *tts.TTS, c *twitch.Client) *homeController {

	controller := &homeController{
		goConfigScreen: goConfigScreen,
		Player:         p,
		Client:         c,
	}

	controller.initScreen()

	controller.Connect()

	return controller
}

// Connect to the player and the twitch
func (c *homeController) Connect() {
	go func() {
		if err := c.Client.Connect(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	c.Client.Join("dannywolfmx2")
	c.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println("Prueba")
		m := fmt.Sprintf("%s: %s", message.User.Name, message.Message)
		c.Player.Add(m)
	})

	c.Player.OnPlayerStart(func(message string) {
		text := strings.Split(message, ":")
		if len(text) > 1 {
			//TODO add method to send data to the view
			//c.View.SetChatMessage(text[0], text[1])
		}
	})
}

// View events
//
// Event that send data from the view to the controller
func (c *homeController) EventConfigTap() {
	if err := c.goConfigScreen(); err != nil {
		//TODO log or report the error
	}
}

func (c *homeController) EventNext() {
	c.Player.Next()
}

func (c *homeController) EventStop() {
	c.Player.Stop()
}

func (c *homeController) GetMessage() (string, string) {
	return c.Sender, c.Message
}

// Main is the main screen of the controller
func (c *homeController) Screen() screens.Screen {
	return c.screen
}

// initScreen will generate the screen of the view
func (c *homeController) initScreen() {
	c.screen = &screens.Home{
		OnStopTap:   c.EventStop,
		OnNextTap:   c.EventNext,
		GetMessage:  c.GetMessage,
		OnConfigTap: c.EventConfigTap,
	}

	go func() {
		time.Sleep(time.Second * 5)
		c.Sender = "Sender"
		c.Message = "Message"
		c.screen.Update()
	}()
}
