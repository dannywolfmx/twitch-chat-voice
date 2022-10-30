package repo

type RepoConfig interface {
	GetAnonymousUsername() string
	SaveAnonymousUsername(username string) error
}
