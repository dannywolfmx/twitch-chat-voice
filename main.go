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
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/ui"
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/joho/godotenv"
)

var screenText = make(chan string)
var userAvatar = make(chan image.Image)
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

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	bearer := os.Getenv(BEARER)
	client_id := os.Getenv(CLIENT_ID)

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
			img, err := getTwitchUserInfo(bearer, client_id, userName[0])
			if err == nil {
				userAvatar <- img
			}

		}
		screenText <- message
	})

	go func() {
		w := app.NewWindow(
			app.Title("Twitch text to voice"),
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

	theme := ui.NewTheme(gofont.Collection())
	theme.Bg = ui.NewColor(0x191E38FF)
	theme.Fg = ui.NewColor(0x2F365FFF)
	theme.ContrastBg = ui.NewColor(0x5661B3FF)
	theme.TextColor = ui.NewColor(0xE6E8FFFF)

	var ops op.Ops

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {

			case system.DestroyEvent:
				quit <- os.Interrupt
				return e.Err
			case system.FrameEvent:
				Layout(theme, &ops, e)

			}
		case m := <-screenText:
			texto = m

			w.Invalidate()
		case m := <-userAvatar:
			img = m
			w.Invalidate()
		}
	}
}

func Layout(theme *ui.Theme, ops *op.Ops, e system.FrameEvent) {
	gtx := layout.NewContext(ops, e)

	main := ui.Main{
		Theme:         theme,
		Texto:         texto,
		TwitchChannel: make(chan string),
		Next:          make(chan bool),
		Img:           img,
	}

	go func() {
		for t := range main.TwitchChannel {
			client.Join(t)
		}
	}()

	go func() {
		for {
			select {
			case <-main.Next:
				player.Next()
			}
		}
	}()

	main.Layout(gtx)

	e.Frame(gtx.Ops)
}
