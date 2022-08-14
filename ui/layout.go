package ui

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var editor = &widget.Editor{
	SingleLine: true,
	Submit:     true,
}

var button = &widget.Clickable{}

type Main struct {
	Theme         *material.Theme
	Texto         string
	Editor        *widget.Editor
	TwitchChannel chan string
	Skip          chan bool
}

func (m *Main) Layout(gtx layout.Context) layout.Dimensions {

	c := color.NRGBA{
		R: m.Theme.Bg.R,
		G: m.Theme.Bg.G,
		B: m.Theme.Bg.B,
		A: m.Theme.Bg.A,
	}

	paint.ColorOp{Color: c}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	margin := layout.Inset{
		Top:    unit.Dp(25),
		Bottom: unit.Dp(25),
		Right:  unit.Dp(35),
		Left:   unit.Dp(35),
	}

	flex := layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}

	elements := []layout.FlexChild{
		Container(m.textField),
		SpacerVertical(50),
		Container(m.messageText),
		SpacerVertical(50),
		Container(m.buttonSkip),
		SpacerVertical(50),
	}

	return margin.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return flex.Layout(gtx, elements...)
	})
}

func (m *Main) buttonSkip(gtx layout.Context) layout.Dimensions {
	for button.Clicked() {
		fmt.Println("Fui clickeado")
		m.Skip <- true
	}
	return material.Button(m.Theme, button, "Skip").Layout(gtx)
}

func (m *Main) textField(gtx layout.Context) layout.Dimensions {
	for _, e := range editor.Events() {
		if e, ok := e.(widget.SubmitEvent); ok {
			m.TwitchChannel <- e.Text

			editor.SetText("")
		}
	}

	e := material.Editor(m.Theme, editor, "Twitch channel")
	e.Font.Style = text.Italic
	border := widget.Border{Color: color.NRGBA{R: 113, G: 140, B: 158, A: 255}, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
	})

}

func (m *Main) messageText(gtx layout.Context) layout.Dimensions {
	title := material.H4(m.Theme, m.Texto)

	title.Color = m.Theme.Fg
	title.Alignment = text.Middle
	return title.Layout(gtx)
}

func SpacerVertical(dp unit.Dp) layout.FlexChild {
	spacer := layout.Spacer{
		Height: unit.Dp(dp),
	}

	return layout.Rigid(spacer.Layout)
}

func Container(content layout.Widget) layout.FlexChild {
	return layout.Rigid(content)
}
