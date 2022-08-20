package oauth

import "github.com/toqueteos/webbrowser"

func OpenBrowserURL(url string) error {
	return webbrowser.Open(url)
}
