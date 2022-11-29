package repo

type RepoConfig interface {
	GetAnonymousUsername() string
	GetClientID() (string, error)
	GetLang() string
	SaveAnonymousUsername(username string) error
	SaveLang(lang string) error
}
