package ui

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	Context    = layout.Context
	Dimensions = layout.Dimensions

	FlexChild = layout.FlexChild
	Flex      = layout.Flex
	Inset     = layout.Inset
	Widget    = layout.Widget
)

var editor = &widget.Editor{
	SingleLine: true,
	Submit:     true,
}

var button = &widget.Clickable{}

type Main struct {
	Theme         *Theme
	Texto         string
	Editor        *widget.Editor
	TwitchChannel chan string
	Skip          chan bool
	Img           image.Image
}

func (m *Main) Layout(gtx Context) Dimensions {

	c := color.NRGBA{
		R: m.Theme.Bg.R,
		G: m.Theme.Bg.G,
		B: m.Theme.Bg.B,
		A: m.Theme.Bg.A,
	}

	paint.ColorOp{Color: c}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	margin := Inset{
		Top:    unit.Dp(25),
		Bottom: unit.Dp(25),
		Right:  unit.Dp(35),
		Left:   unit.Dp(35),
	}

	flex := Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}

	elements := []FlexChild{
		Container(m.TwitchIcon),
		SpacerVertical(50),
		Container(m.messageText),
		SpacerVertical(50),
		Container(m.textInput),
		SpacerVertical(50),
		Container(m.buttonSkip),
		SpacerVertical(50),
	}

	return margin.Layout(gtx, func(gtx Context) Dimensions {
		return flex.Layout(gtx, elements...)
	})
}

func (m *Main) buttonSkip(gtx Context) Dimensions {
	for button.Clicked() {
		m.Skip <- true
	}
	return material.Button(m.Theme.Theme, button, "Skip").Layout(gtx)
}

func (m *Main) textInput(gtx Context) Dimensions {
	for _, e := range editor.Events() {
		if e, ok := e.(widget.SubmitEvent); ok {
			m.TwitchChannel <- e.Text

			editor.SetText("")
		}
	}

	e := material.Editor(m.Theme.Theme, editor, "Twitch channel")
	e.Font.Style = text.Italic
	e.Color = m.Theme.TextColor
	e.HintColor = m.Theme.TextColor

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

func (m *Main) TwitchIcon(gtx Context) Dimensions {

	sizeX := gtx.Constraints.Min.X
	c := NewColor(0xFFFFFFFF)

	border := widget.Border{Color: c, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}
	return border.Layout(gtx, func(gtx Context) Dimensions {
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Stacked(func(gtx Context) Dimensions {
				imageOpt := paint.NewImageOp(m.Img)
				imageOpt.Add(gtx.Ops)

				op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(4, 4)))
				op.Affine(f32.Affine2D{}.Shear(f32.Pt(0, 0), 1, 20))
				paint.PaintOp{}.Add(gtx.Ops)

				return Dimensions{Size: image.Pt(20, 20)}
			}),
			layout.Stacked(func(gtx Context) Dimensions {
				defer clip.UniformRRect(image.Rectangle{Max: image.Pt(256, 256)}, 8).Push(gtx.Ops).Pop()
				paint.Fill(gtx.Ops, c)
				gtx.Constraints.Min.X = sizeX
				return Dimensions{Size: image.Pt(256, 256)}
			}),
		)
	})
}

func (m *Main) messageText(gtx Context) Dimensions {
	title := material.H6(m.Theme.Theme, m.Texto)

	title.Color = m.Theme.TextColor
	title.Alignment = text.Start
	return title.Layout(gtx)
}

func SpacerVertical(dp unit.Dp) FlexChild {
	spacer := layout.Spacer{
		Height: unit.Dp(dp),
	}

	return layout.Rigid(spacer.Layout)
}

func Container(content Widget) FlexChild {
	return layout.Rigid(content)
}
