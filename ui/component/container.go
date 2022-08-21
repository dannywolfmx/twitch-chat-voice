package component

import "gioui.org/layout"

var Container = layout.Rigid

var Row = func(w layout.Widget) layout.FlexChild {
	return layout.Flexed(2, w)
}
