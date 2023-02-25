package repo

import "github.com/dannywolfmx/twitch-chat-voice/model"

type RepoConfig interface {
	AddChat(chat *model.Chat) error
	AddMuttedUser(user model.User) ([]model.User, error)
	GetAnonymousUsername() string
	GetClientID() (string, error)
	GetConfig() *model.Config
	GetChats() []model.Chat
	GetLang() string
	GetMuttedUsers() []model.User
	GetTwitchToken() string
	GetTwitchUserInfo() model.TwitchUser
	RemoveChat(nameChannel string) error
	RemoveMuttedUser(user model.User) ([]model.User, error)
	SaveAnonymousUsername(username string) error
	SaveLang(lang string) error
	SaveTwitchInfo(twitchInfo model.TwitchInfo) error
}
