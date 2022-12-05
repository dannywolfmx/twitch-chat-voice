package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/oauth"
	"github.com/dannywolfmx/twitch-chat-voice/repo"
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	config = make(chan struct{})
)

type MainApp struct {
	Auth       oauth.Oauth
	BearerToke string
	Player     *tts.TTS
	Client     *twitch.Client
	ctx        context.Context
	RepoConfig repo.RepoConfig
}

func (a *MainApp) Run(assets fs.FS) error {
	clientID, err := a.RepoConfig.GetClientID()

	if err != nil {
		return err
	}

	c := &ConnectWithTwitch{
		Auth:               a.Auth,
		SaveTwitchUserinfo: a.RepoConfig.SaveTwitchUser,
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
		Title:            "myproject",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.startup,
		OnDomReady:       a.domready,
		Bind: []any{
			a,
			a.Player,
			a.RepoConfig,
			c,
		},
	})
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
		//Don't repeat the las name
		user := message.User.Name
		m := ""
		if lastUser == user {
			m = message.Message
		} else {
			lastUser = user
			m = fmt.Sprintf("%s ha dicho %s", user, message.Message)
		}
		fmt.Println(m)
		go a.Player.Add(m)
		runtime.EventsEmit(ctx, "OnNewMessage", message)
	})

	runtime.EventsOn(ctx, "OnConnectAnonymous", func(data ...interface{}) {
		if len(data) > 0 {
			username := data[0].(string)
			a.RepoConfig.SaveAnonymousUsername(username)
			runtime.EventsEmit(ctx, "IsLoggedIn", username != "")
			a.Client.Join(username)
		}
	})

	runtime.EventsOn(ctx, "OnIsLoggedIn", func(data ...interface{}) {
		username := a.RepoConfig.GetAnonymousUsername()
		runtime.EventsEmit(ctx, "IsLoggedIn", username != "")
	})

}

func (a *MainApp) domready(ctx context.Context) {
	username := a.RepoConfig.GetAnonymousUsername()
	if username != "" {
		a.Client.Join(username)
	}
}

type ConnectWithTwitch struct {
	Auth               oauth.Oauth
	SaveTwitchUserinfo func(username, token string) error
	clientID           string
}

func (c *ConnectWithTwitch) ConnectWithTwitch() bool {
	token, err := c.Auth.Connect()
	if err != nil {
		log.Println(err)
		return false
	}

	username, err := GetTwitchUsername(token, c.clientID)

	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println(username)

	if err = c.SaveTwitchUserinfo(username, token); err != nil {
		log.Println(err)
		return false
	}

	return true
}

func GetTwitchUsername(token, clientID string) (string, error) {

	request, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users", nil)

	if err != nil {
		return "", nil
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	request.Header.Set("Client-Id", clientID)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return "", nil
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return "", nil
	}

	return string(body), nil
}
