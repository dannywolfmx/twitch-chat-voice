package component

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

func SpacerVertical(dp unit.Dp) layout.FlexChild {
	spacer := layout.Spacer{
		Height: unit.Dp(dp),
	}

	return layout.Rigid(spacer.Layout)
}
