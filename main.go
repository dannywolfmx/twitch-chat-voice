package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	htgotts "github.com/hegedustibor/htgo-tts"
)

func main() {
	TwitchChat()
}

func TwitchChat() {
	client := twitch.NewAnonymousClient()
	//client := twitch.NewClient("yourtwitchusername", "oauth:123123123")

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
		Speak(message.Message)
	})

	client.Join("dannywolfmx2")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func Speak(texto string) {
	speech := htgotts.Speech{Folder: "audio", Language: "es"}
	speech.Speak(texto)
}
