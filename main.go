package main

import (
	"context"
	"embed"
	_ "embed"
	"fmt"
	"image"
	"os"
	"time"

	"github.com/dannywolfmx/go-tts/tts"
	"github.com/dannywolfmx/twitch-chat-voice/oauth"
	"github.com/gempir/go-twitch-irc/v3"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

var UpdateUI = make(chan bool)
var done = make(chan bool)

var texto string

var client *twitch.Client
var player *tts.TTS

var img image.Image

var quit = make(chan os.Signal, 1)

const (
	BEARER    = "BEARER"
	CLIENT_ID = "CLIENT_ID"
	TEST      = "TEST"
)

type MyPlayer struct {
	*tts.TTS
}

var bearerToken string

var auth *oauth.Twitch

//go:embed  .env
var envFile string

func main() {
	//	envVars, err := godotenv.Unmarshal(envFile)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	client_id, ok := envVars[CLIENT_ID]
	//	if !ok {
	//		panic(err)
	//	}
	//
	//	a := &app.MainApp{
	//		Auth:   oauth.NewTwitchOAuth(client_id),
	//		Client: twitch.NewAnonymousClient(),
	//		Player: tts.NewTTS("es"),
	//	}
	//
	//	signal.Notify(quit, os.Interrupt, os.Kill)
	//	go func() {
	//		<-quit
	//		a.Quit()
	//	}()
	//
	//	if err := a.Run(); err != nil {
	//		fmt.Println(err)
	//	}
	//	a.Stop()
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "myproject",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			&Prueba{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Daniel %s, It's show time!", name)
}

type Prueba struct {
	ctx context.Context
}

func (a *Prueba) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *Prueba) Saludo(name string) string {
	return fmt.Sprintf("Daniel %s, It's show time!", name)
}

func (a *Prueba) Suma(num1, num2 int) int {
	return num1 + num2
}

func (a *Prueba) Tiempo() string {
	return time.Now().GoString()
}
