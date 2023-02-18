package repo

type RepoConfig interface {
	AddChat(chat *Chat) error
	GetAnonymousUsername() string
	GetClientID() (string, error)
	GetConfig() *Config
	GetChats() []Chat
	GetLang() string
	GetTwitchToken() string
	GetTwitchUserInfo() TwitchUser
	RemoveChat(nameChannel string) error
	SaveAnonymousUsername(username string) error
	SaveLang(lang string) error
	SaveTwitchInfo(twitchInfo TwitchInfo) error
}
