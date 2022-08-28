package oauth

type Oauth interface {
	Connect() (string, error)
}
