package oauth

import (
	"testing"
)

const (
	CLIENT_ID     = "1de7o2592oe1x32y9j0mpeybjkjj50"
	REDIRECT      = "http://localhost:8080/twitch/oauth"
	RESPONSE_TYPE = "token"
	testURL       = "https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=0000000000000000&redirect_uri=http://localhost:3000&scope=channel%3Amanage%3Apolls+channel%3Aread%3Apolls+user%3Aread%3Aemail+chat%3Aread+chat%3Aedit+channel%3Amoderate+whispers%3Aedit&state=c3ab8aa609ea11e793ae92361f002671"
	State         = "c3ab8aa609ea11e793ae92361f002671"
)

var Scopes = []string{
	"channel:manage:polls",
	"channel:read:polls",
	"user:read:email",
	"chat:read",
	"chat:edit",
	"channel:moderate",
	"whispers:edit",
}

func TestGetURL(t *testing.T) {
	testTwitchStruct := NewTwitch(CLIENT_ID, REDIRECT, RESPONSE_TYPE, State, Scopes...)
	res, err := testTwitchStruct.GetURL()

	if err != nil {
		t.Fatal(err)
	}

	if res != testURL {
		t.Fatal(res)
	}
}

func TestConnect(t *testing.T) {
	testToken := "3wj5i3hb2j0efdoh37iti1xpcnpfur"
	testTwitchStruct := NewTwitch(CLIENT_ID, REDIRECT, RESPONSE_TYPE, State, Scopes...)

	token, err := testTwitchStruct.Connect()

	if err != nil {
		t.Fatal(err)
	}

	if token != testToken {
		t.Fatal(token)
	}
}
