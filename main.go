package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
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

func main() {
	client = twitch.NewAnonymousClient()
	player = tts.NewTTS("es")
	url := "https://go.dev/blog/gopher/header.jpg"
	res, err := http.Get(url)

	if err != nil {
		panic("Error al descargar imagen")
	}

	if err != nil {
		panic("Error al leer buffer de la imagen")
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

func Layout(theme *ui.Theme, ops *op.Ops, e system.FrameEvent) {
	gtx := layout.NewContext(ops, e)

	main := ui.Main{
		Theme:         theme,
		Texto:         texto,
		TwitchChannel: make(chan string),
		Skip:          make(chan bool),
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
			case <-main.Skip:
				player.Skip()
			}
		}
	}()

	main.Layout(gtx)

	e.Frame(gtx.Ops)
}
func drawCircle(img draw.Image, x0, y0, r int, c color.Color) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}
