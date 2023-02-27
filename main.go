package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"os/signal"

	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/app"
	"github.com/dannywolfmx/twitch-chat-voice/app/usecase"
	"github.com/dannywolfmx/twitch-chat-voice/oauth"
	"github.com/dannywolfmx/twitch-chat-voice/repo"
	"github.com/gempir/go-twitch-irc/v3"
)

const CONFIG_FILE string = "config.json"

var quit = make(chan os.Signal, 1)

// Note: the dist content will be created by wails dev command, is normal if you see
//
//	a empty directory.
//
//go:embed all:frontend/dist
var assets embed.FS

type MyPlayer struct {
	*tts.TTS
}

func main() {
	repoConfig, err := repo.NewRepoConfigFile(CONFIG_FILE)

	if err != nil {
		panic(err)
	}
	client := twitch.NewAnonymousClient()
	config := usecase.NewConfig(repoConfig, client)

	clientID, err := repoConfig.GetClientID()

	if err != nil {
		fmt.Print("Note: you're running the program without a client id. Social integration is disabled.")
	}

	fmt.Println(config.GetSampleRateOfTTS())
	player := tts.NewTTS(repoConfig.GetLang(), config.GetSampleRateOfTTS())
	player.Play()

	a := &app.MainApp{
		Auth:   oauth.NewTwitchOAuth(clientID),
		Client: client,
		Player: player,
		Config: config,
	}

	signal.Notify(quit, os.Interrupt, os.Kill)
	go func() {
		<-quit
		//a.Quit()
	}()

	if err := a.Run(assets); err != nil {
		fmt.Println(err)
	}
	a.Stop()

	if err != nil {
		println("Error:", err.Error())
	}
}
