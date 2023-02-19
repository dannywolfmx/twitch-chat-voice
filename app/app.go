package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/app/usecase"
	"github.com/dannywolfmx/twitch-chat-voice/model"
	"github.com/dannywolfmx/twitch-chat-voice/oauth"
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var config = make(chan struct{})

type MainApp struct {
	Auth       oauth.Oauth
	BearerToke string
	Player     *tts.TTS
	Client     *twitch.Client
	ctx        context.Context
	Config     usecase.Config
}

func (a *MainApp) Run(assets fs.FS) error {
	clientID, err := a.Config.GetClientID()

	if err != nil {
		return err
	}

	c := &ConnectWithTwitch{
		Auth:               a.Auth,
		SaveTwitchUserinfo: a.Config.SaveTwitchInfo,
		clientID:           clientID,
	}

	go func() {
		//Connect to the IRC twitch chat
		// warning the connect function is thread blocking
		// don't put function after this block
		if err := a.Client.Connect(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	// Create application with options
	return wails.Run(&options.App{
		Title:            "Chat to voice",
		Width:            400,
		Height:           500,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.startup,
		Bind: []any{
			a,
			a.Player,
			a.Config,
			c,
		},
	})
}

func (a *MainApp) Stop() {
	a.Client.Disconnect()
	a.Player.Stop()
}

type TwitchUserInfo struct {
	Data []SingleData `json:"data"`
}

type SingleData struct {
	Image string `json:"profile_image_url"`
}

type Message struct {
	Message string `json:"Message"`
	User    string `json:"user"`
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *MainApp) startup(ctx context.Context) {
	lastUser := ""
	a.ctx = ctx

	a.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		//Prevent the tts repeat the las name
		user := message.User.Name
		m := ""
		if lastUser == user {
			m = message.Message
		} else {
			lastUser = user
			m = fmt.Sprintf("%s ha dicho %s", user, message.Message)
		}
		go a.Player.Add(m)
		runtime.EventsEmit(ctx, "OnNewMessage", message)
	})

	runtime.EventsOn(ctx, "OnConnectAnonymous", func(data ...interface{}) {
		if len(data) > 0 {
			username := data[0].(string)
			a.Config.SaveAnonymousUsername(username)
			runtime.EventsEmit(ctx, "IsLoggedIn", username != "")
		}
	})

	runtime.EventsOn(ctx, "OnIsLoggedIn", func(data ...interface{}) {
		username := a.Config.GetAnonymousUsername()
		runtime.EventsEmit(ctx, "IsLoggedIn", username != "")
	})

}

type ConnectWithTwitch struct {
	Auth               oauth.Oauth
	SaveTwitchUserinfo func(info model.TwitchInfo) error
	clientID           string
}

type getDataTwitch struct {
	Data []model.TwitchUser `json:"data"`
}

func (c *ConnectWithTwitch) ConnectWithTwitch() bool {
	token, err := c.Auth.Connect()
	if err != nil {
		log.Println(err)
		return false
	}

	rawUserInfo, err := GetTwitchUserInfo(token, c.clientID)

	if err != nil {
		log.Println(err)
		return false
	}

	twitchData := getDataTwitch{}

	err = json.Unmarshal(rawUserInfo, &twitchData)

	if err != nil {
		log.Println(err)
		return false
	}

	if len(twitchData.Data) == 0 {
		return false
	}

	userData := twitchData.Data[0]

	userInfo := model.TwitchInfo{
		Token:      token,
		TwitchUser: userData,
	}

	if err = c.SaveTwitchUserinfo(userInfo); err != nil {
		log.Println(err)
		return false
	}

	return true
}

func GetTwitchUserInfo(token, clientID string) ([]byte, error) {

	request, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users", nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	request.Header.Set("Client-Id", clientID)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	return body, nil
}
