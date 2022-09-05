package components

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ToolbarLayout(onBackTap func(), onConfigTap func()) *widget.Toolbar {
	elements := []widget.ToolbarItem{}

	if onBackTap != nil {
		action := widget.NewToolbarAction(theme.NavigateBackIcon(), onBackTap)
		elements = append(elements, action)
	}

	elements = append(elements, widget.NewToolbarSpacer())

	if onConfigTap != nil {
		action := widget.NewToolbarAction(theme.MenuIcon(), onConfigTap)
		elements = append(elements, action)
	}
	return widget.NewToolbar(elements...)
}
