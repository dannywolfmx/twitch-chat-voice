package main

import (
	"fmt"
	"image/color"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/dannywolfmx/go-tts/tts"
	"github.com/gempir/go-twitch-irc/v3"
)

var screenText = make(chan string)
var done = make(chan bool)

var texto string

func main() {
	twitchChannelName := "dannywolfmx2"
	client := twitch.NewAnonymousClient()
	player := tts.NewTTS("es")
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

	client.Join(twitchChannelName)

	client.Say(twitchChannelName, "/vips")

	go func() {
		w := app.NewWindow()
		run(w, player, client)
	}()

	fmt.Println(client.Connect())
}

func run(w *app.Window, player *tts.TTS, client *twitch.Client) error {

	theme := material.NewTheme(gofont.Collection())
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
	layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}.Layout(gtx,
		layout.Rigid(layout.Spacer{Height: unit.Dp(25)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			title := material.H1(theme, texto)
			maroon := color.NRGBA{R: 127, A: 255}

			title.Color = maroon
			title.Alignment = text.Middle
			return title.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(25)}.Layout),
	)

	e.Frame(gtx.Ops)
}
