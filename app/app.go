package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io"
	"net/http"

	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/controller"
	"github.com/dannywolfmx/twitch-chat-voice/oauth"
	"github.com/dannywolfmx/twitch-chat-voice/route"
	"github.com/dannywolfmx/twitch-chat-voice/view"
	"github.com/gempir/go-twitch-irc/v3"
)

var (
	config = make(chan struct{})
)

type MainApp struct {
	Auth       oauth.Oauth
	BearerToke string
	view       view.View
	Player     *tts.TTS
	Client     *twitch.Client
}

func (a *MainApp) events() {
}

const (
	HOME_SCREEN   = "home"
	CONFIG_SCREEN = "config"
)

func (a *MainApp) Run() error {
	a.events()
	config := view.ConfigView{
		WindowSize: view.Size{
			Width:  400,
			Height: 736,
		},
		Title: "Twitch App",
	}

	a.view = view.NewView(config)

	go func() {
		//Connect to the IRC twitch chat
		// warning the connect function is thread blocking
		// don't put function after this block
		if err := a.Client.Connect(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	route := route.NewRoute(a.view, true)
	{
		//Set home screen
		home := controller.NewHomeController(func() error { return route.Go(CONFIG_SCREEN) },
			a.Player,
			a.Client,
		)
		route.Set(HOME_SCREEN, home)

		config := controller.NewConfigController(func() { route.Go(HOME_SCREEN) },
			a.Client,
		)

		route.Set(CONFIG_SCREEN, config)

	}

	route.Go(HOME_SCREEN)
	a.view.ShowAndRun()

	return nil
}

func (a *MainApp) Quit() {
	//Close the window at the end to keep the running function until it stop
	a.view.Quit()
}

func (a *MainApp) Stop() {
	a.Client.Disconnect()
	a.Player.Stop()
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
	defer res.Body.Close()

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

	img, _, err := image.Decode(res.Body)
	defer res.Body.Close()

	return img, err

}

type TwitchUserInfo struct {
	Data []SingleData `json:"data"`
}

type SingleData struct {
	Image string `json:"profile_image_url"`
}
