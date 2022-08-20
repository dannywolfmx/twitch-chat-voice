package oauth

import (
	"testing"
	"time"
)

func TestOpenBrowser(t *testing.T) {
	if err := OpenBrowserURL("https://go.dev/"); err != nil {
		t.Fail()
	}

	//The SO can stop the browser opening if the test stops before
	// we add a wait sleep to prevent this
	time.Sleep(1 * time.Second)
}
