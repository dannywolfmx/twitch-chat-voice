package main

import (
	"fmt"
	"image/color"
	"os"
	"os/signal"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/ui"
	"github.com/gempir/go-twitch-irc/v3"
)

var screenText = make(chan string)
var done = make(chan bool)

var texto string

var client *twitch.Client
var player *tts.TTS

func main() {
	client = twitch.NewAnonymousClient()
	player = tts.NewTTS("es")
	//client := twitch.NewClient("yourtwitchusername", "oauth:123123123")

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		m := fmt.Sprintf("%s: %s \n", message.User.Name, message.Message)
		player.Play(m)
	})

	go func() {
		for playing := range player.Playing() {
			select {
			case <-player.Done:
				return
			default:
				screenText <- playing.GetText()
				playing.Play()
			}
		}
	}()

	go func() {
		w := app.NewWindow(
			app.Title("Twitch text to voice"),
		)
		run(w, player, client)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		for {
			select {
			case <-quit:
				client.Disconnect()
				player.Stop()
			}
		}
	}()

	fmt.Println(client.Connect())
}

func run(w *app.Window, player *tts.TTS, client *twitch.Client) error {

	theme := material.NewTheme(gofont.Collection())
	theme.Bg = color.NRGBA{R: 54, G: 69, B: 79, A: 255}
	theme.Fg = color.NRGBA{
		R: 195,
		G: 206,
		B: 214,
		A: 255,
	}
	var ops op.Ops

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {

			case system.DestroyEvent:
				client.Disconnect()
				player.Stop()
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

func Layout(theme *material.Theme, ops *op.Ops, e system.FrameEvent) {
	gtx := layout.NewContext(ops, e)

	main := ui.Main{
		Theme:         theme,
		Texto:         texto,
		TwitchChannel: make(chan string),
		Skip:          make(chan bool),
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
			case <-main.Skip:
				player.Skip()
			}
		}
	}()

	main.Layout(gtx)

	e.Frame(gtx.Ops)
}
