package screens

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/dannywolfmx/twitch-chat-voice/ui/component"
	"golang.org/x/image/draw"
)

type Home struct {
	Texto         string
	Editor        *widget.Editor
	TwitchChannel chan string
	Next          chan struct{}
	ConnectTwitch chan struct{}
	Img           image.Image
	*component.ComponentTheme
}

var emptyImage = image.NewNRGBA(image.Rect(0, 0, 128, 128))

func NewHomeScreen(next, connect chan struct{}, twitchChannel chan string) *Home {
	return &Home{
		Texto:          "",
		TwitchChannel:  twitchChannel,
		Next:           next,
		Img:            emptyImage,
		ComponentTheme: component.DefaultTheme,
		ConnectTwitch:  connect,
	}

}

func (m *Home) Render(gtx Context) Dimensions {
	flex := Flex{
		Axis: layout.Vertical,
	}

	elements := []FlexChild{
		component.Container(m.TwitchIcon),
		component.SpacerVertical(36),
		component.Row(m.messageText),
		component.SpacerVertical(36),
		component.Container(m.textInput),
		component.SpacerVertical(16),
		component.Container(m.buttonSkip),
		component.SpacerVertical(16),
		component.Container(m.buttonConnectTwitch),
	}

	return flex.Layout(gtx, elements...)
}

func (m *Home) buttonConnectTwitch(gtx Context) Dimensions {
	for component.ButtonTwitch.Clicked() {
		m.ConnectTwitch <- struct{}{}
	}
	return material.Button(m.Theme, component.ButtonTwitch, "Connect to twitch").Layout(gtx)
}

func (m *Home) buttonSkip(gtx Context) Dimensions {
	for component.Button.Clicked() {
		m.ConnectTwitch <- struct{}{}
	}
	return material.Button(m.Theme, component.Button, "Skip").Layout(gtx)
}

func (m *Home) textInput(gtx Context) Dimensions {
	for _, e := range component.Editor.Events() {
		if e, ok := e.(widget.SubmitEvent); ok {
			m.TwitchChannel <- e.Text

			component.Editor.SetText("")
		}
	}

	e := material.Editor(m.Theme, component.Editor, "Twitch channel")
	e.Font.Style = text.Italic
	e.Color = m.TextColor
	e.HintColor = m.TextColor

	c := m.Theme.Fg

	sizeX := gtx.Constraints.Min.X

	border := widget.Border{Color: c, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
	return border.Layout(gtx, func(gtx Context) Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx Context) Dimensions {
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 8).Push(gtx.Ops).Pop()
				paint.Fill(gtx.Ops, c)
				return Dimensions{Size: gtx.Constraints.Min}
			}),
			layout.Stacked(func(gtx Context) Dimensions {
				gtx.Constraints.Min.X = sizeX
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
			}),
		)
	})
}

func (m *Home) TwitchIcon(gtx Context) Dimensions {

	if m.Img == nil {
		m.Img = emptyImage
	}
	rec := image.Rect(0, 0, 128, 128)

	dst := image.NewNRGBA(image.Rect(0, 0, 128, 128))
	draw.NearestNeighbor.Scale(dst, dst.Rect, m.Img, m.Img.Bounds(), draw.Over, nil)

	return layout.Stack{Alignment: layout.Center}.Layout(gtx,
		layout.Stacked(func(gtx Context) Dimensions {
			defer clip.Ellipse{Max: rec.Max, Min: rec.Min}.Push(gtx.Ops).Pop()
			imageOpt := paint.NewImageOp(dst)
			imageOpt.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)

			return Dimensions{Size: rec.Max}
		}),
		layout.Stacked(func(gtx Context) Dimensions {
			defer clip.Ellipse{Max: rec.Max, Min: rec.Min}.Push(gtx.Ops).Pop()

			paint.Fill(gtx.Ops, component.NewColor(0xFFFFFF09))
			return Dimensions{Size: rec.Max}
		}),
	)
}

func (m *Home) messageText(gtx Context) Dimensions {
	title := material.H6(m.Theme, m.Texto)

	title.Color = m.TextColor
	title.Alignment = text.Start
	sizeX := gtx.Constraints.Min.X
	border := widget.Border{Color: m.Fg, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
	return border.Layout(gtx, func(gtx Context) Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx Context) Dimensions {
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Max}, 8).Push(gtx.Ops).Pop()
				paint.Fill(gtx.Ops, m.Fg)
				return Dimensions{Size: gtx.Constraints.Min}
			}),
			layout.Stacked(func(gtx Context) Dimensions {
				gtx.Constraints.Min.X = sizeX
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, title.Layout)
			}),
		)
	})
}
