package oauth

import (
	"net/url"
	"strings"

	"github.com/dannywolfmx/twitch-chat-voice/oauth/server"
)

const BASE_TWITCH_URL = "https://id.twitch.tv/oauth2/authorize"
const DEFAULT_REDIRECT_URI = "http://localhost:8080/twitch/oauth"

type Twitch struct {
	ClientID     string
	RedirectURI  string
	ResponseType string
	Scopes       []string
	State        string
}

func NewTwitchOAuth(clientID string) *Twitch {

	scopes := []string{
		"channel:manage:polls",
		"channel:read:polls",
		"user:read:email",
		"chat:read",
		"chat:edit",
		"channel:moderate",
		"whispers:edit",
	}

	return &Twitch{
		Scopes:       scopes,
		ClientID:     clientID,
		RedirectURI:  DEFAULT_REDIRECT_URI,
		ResponseType: "token",
		State:        "c3ab8aa609ea11e793ae92361f002671",
	}
}

func (t *Twitch) Connect() (string, error) {
	url, err := t.getURL()
	if err != nil {
		return "", err
	}

	if err = OpenBrowserURL(url); err != nil {
		return "", err
	}

	s := server.NewServer(":8080")

	return s.Run("/twitch/oauth")
}

func (t *Twitch) getURL() (string, error) {
	u, err := url.Parse(BASE_TWITCH_URL)
	if err != nil {
		return "", err
	}

	q := u.Query()

	q.Set("response_type", t.ResponseType)
	q.Set("client_id", t.ClientID)
	q.Set("redirect_uri", t.RedirectURI)
	q.Set("scope", strings.Join(t.Scopes, " "))
	q.Set("state", t.State)

	u.RawQuery = q.Encode()

	return u.String(), nil
}
