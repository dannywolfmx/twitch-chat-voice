package screens

import "fyne.io/fyne/v2"

type Screen interface {
	Content() fyne.CanvasObject
	Update()
}
