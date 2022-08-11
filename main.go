package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/dannywolfmx/go-tts/tts"
	"github.com/gempir/go-twitch-irc/v3"
)

func main() {
	client := twitch.NewAnonymousClient()
	player := tts.NewTTS()
	go func() {
		player.Run()
	}()
	//client := twitch.NewClient("yourtwitchusername", "oauth:123123123")

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
		player.Play("es", message.Message)
	})

	client.Join("dannywolfmx2")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		for {
			select {
			case <-quit:
				player.Stop()
				client.Disconnect()
			}
		}
	}()

	fmt.Println(client.Connect())
}
