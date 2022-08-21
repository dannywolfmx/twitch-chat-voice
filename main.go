package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/oauth"
	"github.com/dannywolfmx/twitch-chat-voice/ui/screens"
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

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	//bearer = os.Getenv(BEARER)
	client_id := os.Getenv(CLIENT_ID)

	auth = oauth.NewTwitchDefault(client_id)

	client = twitch.NewAnonymousClient()
	player = tts.NewTTS("es")

	//client := twitch.NewClient("yourtwitchusername", "oauth:123123123")
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		m := fmt.Sprintf("%s: %s \n", message.User.Name, message.Message)
		player.Add(m)
	})

	player.OnPlayerStart(func(message string) {
		userName := strings.Split(message, ":")
		if len(userName) == 2 {
			avatar, err := getTwitchUserInfo(bearerToken, client_id, userName[0])
			if err == nil {
				img = avatar
			}
		}
		texto = message
		UpdateUI <- true
	})

	go func() {
		w := app.NewWindow(
			app.Title("Twitch text to voice"),
			app.Size(356, 800),
		)
		run(w, player, client)
	}()

	signal.Notify(quit, os.Interrupt, os.Kill)
	done := make(chan struct{}, 1)
	go func() {
		<-quit
		stopProgram()
		os.Exit(0)
		done <- struct{}{}
	}()

	fmt.Println(client.Connect())
	<-done
}

type TwitchUserInfo struct {
	Data []SingleData `json:"data"`
}

type SingleData struct {
	Image string `json:"profile_image_url"`
}

func getTwitchUserInfo(bearer, client_id, username string) (image.Image, error) {
	if bearer == "" {
		return nil, errors.New("Empty bearer token")
	}
	url := fmt.Sprintf("https://api.twitch.tv/helix/users?login=%s", username)

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", bearer)},
		"Client-Id":     {client_id},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	buff, err := io.ReadAll(res.Body)
	res.Body.Close()

	userInfo := &TwitchUserInfo{}
	err = json.Unmarshal(buff, userInfo)

	if err != nil {
		return nil, err
	}

	if len(userInfo.Data) != 1 {
		return nil, errors.New("No image")
	}

	res, err = http.Get(userInfo.Data[0].Image)
	if err != nil {
		return nil, err
	}

	img, _, err = image.Decode(res.Body)
	defer res.Body.Close()

	return img, err

}

func stopProgram() {
	client.Disconnect()
	player.Stop()
	player.CleanCache()
}

func run(w *app.Window, player *tts.TTS, client *twitch.Client) error {
	var ops op.Ops
	next := make(chan struct{})
	twitchChannel := make(chan string)
	connectTwitch := make(chan struct{})

	go func() {
		for range next {
			player.Next()
		}
	}()

	go func() {
		for t := range twitchChannel {
			fmt.Println(t)
			client.Join(t)
		}
	}()

	go func() {
		for range connectTwitch {
			fmt.Println("Connecting")
			var err error
			bearerToken, err = auth.Connect()
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	home := screens.NewHomeScreen(next, connectTwitch, twitchChannel)

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {

			case system.DestroyEvent:
				quit <- os.Interrupt
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				screens.Show(gtx, home)

				e.Frame(gtx.Ops)
			}

		case <-UpdateUI:
			home.Img = img
			home.Texto = texto
			w.Invalidate()
		}
	}
}
