package usecase

import (
	"github.com/dannywolfmx/twitch-chat-voice/model"
	"github.com/dannywolfmx/twitch-chat-voice/repo"
)

type Config interface {
	AddChat(chat *model.Chat) error
	GetAnonymousUsername() string
	GetClientID() (string, error)
	GetConfig() *model.Config
	GetChats() []model.Chat
	GetLang() string
	GetTwitchToken() string
	GetTwitchUserInfo() model.TwitchUser
	RemoveChat(nameChannel string) error
	SaveAnonymousUsername(username string) error
	SaveLang(lang string) error
	SaveTwitchInfo(twitchInfo model.TwitchInfo) error
}

type config struct {
	repository repo.RepoConfig
}

func NewConfig(repository repo.RepoConfig) *config {
	return &config{
		repository: repository,
	}
}

func (c *config) AddChat(chat *model.Chat) error {
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

func (c *config) GetTwitchToken() string {
	return c.repository.GetTwitchToken()
}

func (c *config) GetTwitchUserInfo() model.TwitchUser {
	return c.repository.GetTwitchUserInfo()
}

func (c *config) RemoveChat(nameChannel string) error {
	return c.repository.RemoveChat(nameChannel)
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
