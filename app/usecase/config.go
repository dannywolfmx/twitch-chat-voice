package usecase

import (
	"github.com/dannywolfmx/twitch-chat-voice/model"
	"github.com/dannywolfmx/twitch-chat-voice/repo"
	"github.com/gempir/go-twitch-irc/v3"
)

type Config interface {
	AddChat(chat *model.Chat) error
	GetAnonymousUsername() string
	GetClientID() (string, error)
	GetConfig() *model.Config
	GetChats() []model.Chat
	GetLang() string
	GetMuttedUsers() []model.User
	GetTwitchToken() string
	GetTwitchUserInfo() model.TwitchUser
	GetSampleRateOfTTS() int
	IsMutted(username model.User) bool
	RemoveChat(nameChannel string) error
	ToggleMuttedUser(username model.User) ([]model.User, error)
	SaveAnonymousUsername(username string) error
	SaveLang(lang string) error
	SaveTwitchInfo(twitchInfo model.TwitchInfo) error
}

type config struct {
	repository repo.RepoConfig
	client     *twitch.Client
}

func NewConfig(repository repo.RepoConfig, client *twitch.Client) *config {

	for _, chat := range repository.GetChats() {
		client.Join(chat.NameChannel)
	}

	return &config{
		repository: repository,
		client:     client,
	}
}

func (c *config) AddChat(chat *model.Chat) error {
	c.client.Join(chat.NameChannel)
	return c.repository.AddChat(chat)
}

func (c *config) GetAnonymousUsername() string {
	return c.repository.GetAnonymousUsername()
}

func (c *config) GetClientID() (string, error) {
	return c.repository.GetClientID()
}

func (c *config) GetConfig() *model.Config {
	return c.repository.GetConfig()
}

func (c *config) GetChats() []model.Chat {
	return c.repository.GetChats()
}

func (c *config) GetLang() string {
	return c.repository.GetLang()
}

func (c *config) GetMuttedUsers() []model.User {
	return c.repository.GetMuttedUsers()
}

func (c *config) GetTwitchToken() string {
	return c.repository.GetTwitchToken()
}

func (c *config) GetTwitchUserInfo() model.TwitchUser {
	return c.repository.GetTwitchUserInfo()
}
func (c *config) GetSampleRateOfTTS() int {
	sample := c.repository.GetSampleRateOfTTS()
	if sample == 0 {
		return 24000
	}
	return sample
}

func (c *config) IsMutted(username model.User) bool {
	var isMutted bool
	users := c.repository.GetMuttedUsers()

	for _, user := range users {
		if user == username {
			isMutted = true
			break
		}
	}

	return isMutted
}

func (c *config) RemoveChat(nameChannel string) error {
	c.client.Depart(nameChannel)
	return c.repository.RemoveChat(nameChannel)
}

func (c *config) ToggleMuttedUser(username model.User) ([]model.User, error) {
	if c.IsMutted(username) {
		return c.repository.RemoveMuttedUser(username)
	}

	return c.repository.AddMuttedUser(username)
}

func (c *config) SaveAnonymousUsername(username string) error {
	return c.repository.SaveAnonymousUsername(username)
}

func (c *config) SaveLang(lang string) error {
	return c.repository.SaveLang(lang)
}

func (c *config) SaveTwitchInfo(twitchInfo model.TwitchInfo) error {
	return c.repository.SaveTwitchInfo(twitchInfo)
}
