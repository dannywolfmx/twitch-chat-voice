package main

import (
	"embed"
	_ "embed"
	"fmt"
	"image"
	"os"
	"os/signal"

	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/app"
	"github.com/dannywolfmx/twitch-chat-voice/oauth"
	"github.com/dannywolfmx/twitch-chat-voice/repo"
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
)

var UpdateUI = make(chan bool)
var done = make(chan bool)

var texto string

var client *twitch.Client
var player *tts.TTS

var img image.Image

var quit = make(chan os.Signal, 1)

const (
	BEARER    = "BEARER"
	CLIENT_ID = "CLIENT_ID"
	TEST      = "TEST"
)

type MyPlayer struct {
	*tts.TTS
}

var bearerToken string

var auth *oauth.Twitch

//go:embed  .env
var envFile string

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	envVars, err := godotenv.Unmarshal(envFile)
	if err != nil {
		panic(err)
	}

	client_id, ok := envVars[CLIENT_ID]
	if !ok {
		panic(err)
	}

	player := tts.NewTTS("es")
	player.Play()

	repoConfig := repo.NewRepoConfigFile("config.json")

	a := &app.MainApp{
		Auth:       oauth.NewTwitchOAuth(client_id),
		Client:     twitch.NewAnonymousClient(),
		Player:     player,
		RepoConfig: repoConfig,
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
