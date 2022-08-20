package oauth

import (
	"net/url"
	"strings"

	"github.com/dannywolfmx/twitch-chat-voice/oauth/server"
)

const BASE_TWITCH_URL = "https://id.twitch.tv/oauth2/authorize"
const DEFAULT_REDIRECT_URI = "http://localhost:8080/twitch/oauth"

var ()

type Twitch struct {
	ClientID     string
	RedirectURI  string
	ResponseType string
	Scope        []string
	State        string
}

func NewTwitch(clientID, redirectURI, responseType, state string, scope ...string) *Twitch {
	return &Twitch{
		ClientID:     clientID,
		RedirectURI:  redirectURI,
		ResponseType: responseType,
		Scope:        scope,
		State:        state,
	}
}

func (t *Twitch) Connect() (string, error) {
	url, err := t.GetURL()
	if err != nil {
		return "", err
	}

	if err = OpenBrowserURL(url); err != nil {
		return "", err
	}

	return server.NewServer(":8080").Run("/twitch/oauth")
}

func (t *Twitch) GetURL() (string, error) {
	u, err := url.Parse(BASE_TWITCH_URL)
	if err != nil {
		return "", err
	}

	q := u.Query()

	q.Set("response_type", t.ResponseType)
	q.Set("client_id", t.ClientID)
	q.Set("redirect_uri", t.RedirectURI)
	q.Set("scope", strings.Join(t.Scope, " "))
	q.Set("state", t.State)

	u.RawQuery = q.Encode()

	return u.String(), nil
}
