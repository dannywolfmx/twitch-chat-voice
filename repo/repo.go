package repo

type RepoConfig interface {
	GetAnonymousUsername() string
	GetClientID() (string, error)
	GetConfig() *Config
	GetLang() string
	GetTwitchToken() string
	SaveAnonymousUsername(username string) error
	SaveLang(lang string) error
	SaveTwitchUser(username, token string) error
}
