package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/oauth"
	"github.com/dannywolfmx/twitch-chat-voice/view"
	"github.com/gempir/go-twitch-irc/v3"
)

var (
	config = make(chan struct{})
)

type MainApp struct {
	Auth       oauth.Oauth
	Player     *tts.TTS
	Client     *twitch.Client
	BearerToke string
	View       view.View
}

func (a *MainApp) events() {
	a.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		m := fmt.Sprintf("%s: %s \n", message.User.Name, message.Message)
		a.Player.Add(m)
	})

	a.Player.OnPlayerStart(func(message string) {
		userName := strings.Split(message, ":")
		if len(userName) > 1 {
		}
	})
}

func (a *MainApp) Run() error {
	a.events()
	twitchChannel := make(chan string)
	//connectTwitch := make(chan struct{})

	go func() {
		for t := range twitchChannel {
			fmt.Println(t)
			a.Client.Join(t)
		}
	}()

	onConfigTap := func() {
		a.View.ChangeScreen(view.CONFIG_SCREEN)
	}

	onNextTap := func() {
		a.Player.Next()
	}

	onStopTap := func() {
		fmt.Println("Stop")
	}

	config := view.ConfigView{
		OnConfigTap:   onConfigTap,
		OnStopTap:     onStopTap,
		OnNextTap:     onNextTap,
		DefaultScreen: view.CONFIG_SCREEN,
	}

	a.View = view.NewView(config)

	a.View.ShowAndRun()

	return nil
}

func (a *MainApp) Quit() {
	//Close the window at the end to keep the running function until it stop
	a.View.Quit()
}

func (a *MainApp) Stop() {
	a.Client.Disconnect()
	a.Player.Stop()
	a.Player.CleanCache()
	log.Println("Closed")
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
