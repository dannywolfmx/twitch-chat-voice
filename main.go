package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"os/signal"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/ui"
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/oliamb/cutter"
)

var screenText = make(chan string)
var done = make(chan bool)

var texto string

var client *twitch.Client
var player *tts.TTS

var img image.Image

var quit = make(chan os.Signal, 1)

func main() {
	client = twitch.NewAnonymousClient()
	player = tts.NewTTS("es")
	url := "https://go.dev/blog/gopher/header.jpg"
	res, err := http.Get(url)

	if err != nil {
		panic("Error al descargar imagen")
	}

	img, err = jpeg.Decode(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic("Error al leer buffer de la imagen")
	}

	img, err = cutter.Crop(img, cutter.Config{
		Width:  256,
		Height: 256,
		Mode:   cutter.Centered,
	})

	//client := twitch.NewClient("yourtwitchusername", "oauth:123123123")
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		m := fmt.Sprintf("%s: %s \n", message.User.Name, message.Message)
		player.Add(m)
	})

	player.OnPlayerStart(func(message string) {
		screenText <- message
	})

	go func() {
		w := app.NewWindow(
			app.Title("Twitch text to voice"),
		)
		run(w, player, client)
	}()

	signal.Notify(quit, os.Interrupt, os.Kill)
	done := make(chan struct{}, 1)
	go func() {
		<-quit
		stopProgram()
		os.Exit(0)
		done <- struct{}{}
	}()

	fmt.Println(client.Connect())
	<-done
}

func stopProgram() {
	client.Disconnect()
	player.Stop()
	player.CleanCache()
}

func run(w *app.Window, player *tts.TTS, client *twitch.Client) error {

	theme := ui.NewTheme(gofont.Collection())
	theme.Bg = ui.NewColor(0x191E38FF)
	theme.Fg = ui.NewColor(0x2F365FFF)
	theme.ContrastBg = ui.NewColor(0x5661B3FF)
	theme.TextColor = ui.NewColor(0xE6E8FFFF)

	var ops op.Ops

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {

			case system.DestroyEvent:
				quit <- os.Interrupt
				return e.Err
			case system.FrameEvent:
				Layout(theme, &ops, e)

			}
		case m := <-screenText:
			texto = m

			w.Invalidate()
		}
	}
}

func Layout(theme *ui.Theme, ops *op.Ops, e system.FrameEvent) {
	gtx := layout.NewContext(ops, e)

	main := ui.Main{
		Theme:         theme,
		Texto:         texto,
		TwitchChannel: make(chan string),
		Next:          make(chan bool),
		Img:           img,
	}

	go func() {
		for t := range main.TwitchChannel {
			fmt.Println(t)
			client.Join(t)
		}
	}()

	go func() {
		for {
			select {
			case <-main.Next:
				player.Next()
			}
		}
	}()

	main.Layout(gtx)

	e.Frame(gtx.Ops)
}
