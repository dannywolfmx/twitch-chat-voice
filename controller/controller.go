package controller

import (
	"github.com/dannywolfmx/twitch-chat-voice/view/screens"
)

type Controller interface {
	Screen() screens.Screen
}
