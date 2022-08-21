package screens

import (
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/dannywolfmx/twitch-chat-voice/ui/component"
)

var spaceV = unit.Dp(25)
var spaceH = unit.Dp(35)

var margin = Inset{
	Top:    spaceV,
	Bottom: spaceV,
	Right:  spaceH,
	Left:   spaceH,
}

func Show(gtx layout.Context, screen Screen) Dimensions {
	paint.ColorOp{Color: component.DefaultTheme.Bg}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return margin.Layout(gtx, screen.Render)
}
