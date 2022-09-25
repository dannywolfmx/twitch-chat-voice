package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomButton struct {
	widget.BaseWidget
	Resource fyne.Resource
	OnTapped func()
	tapAnim  *fyne.Animation
}

func NewCustomButton(resourceIcon fyne.Resource, tapped func()) *CustomButton {
	button := &CustomButton{
		OnTapped: tapped,
	}

	button.ExtendBaseWidget(button)
	button.SetResource(resourceIcon)
	return button
}

func (c *CustomButton) SetResource(res fyne.Resource) {
	c.Resource = res
	c.Refresh()
}

func (c *CustomButton) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	icon := canvas.NewImageFromResource(c.Resource)
	background := canvas.NewRectangle(color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x80})
	icon.FillMode = canvas.ImageFillContain
	r := &customButtonRender{
		button:     c,
		icon:       icon,
		background: background,
	}
	tapBG := canvas.NewRectangle(color.Transparent)
	c.tapAnim = newButtonTapAnimation(tapBG, c)
	c.tapAnim.Curve = fyne.AnimationEaseOut
	r.objects = []fyne.CanvasObject{icon, tapBG}

	return r
}

func (c *CustomButton) Tapped(p *fyne.PointEvent) {
	c.tapAnimation()
	if c.OnTapped != nil {
		c.OnTapped()
	}
}

func (c *CustomButton) MinSize() fyne.Size {
	c.ExtendBaseWidget(c)
	return c.BaseWidget.MinSize()
}

func (b *CustomButton) tapAnimation() {
	if b.tapAnim == nil {
		return
	}
	b.tapAnim.Stop()
	b.tapAnim.Start()
}

type customButtonRender struct {
	icon       *canvas.Image
	button     *CustomButton
	background *canvas.Rectangle
	objects    []fyne.CanvasObject
	size       fyne.Size
}

func (c *customButtonRender) Layout(size fyne.Size) {
	c.background.Move(c.icon.Position())
	c.background.Resize(size)
	c.icon.Resize(size)
	c.size = size
	//c.button.Resize(size)
}

func (c *customButtonRender) MinSize() fyne.Size {
	height := c.button.Size().Height
	size := theme.IconInlineSize()

	if c.size.Width > 128 {
		height = 128
	} else {
		height = c.size.Width
	}

	return fyne.NewSize(size, height)
}

func (c *customButtonRender) Refresh() {
	c.icon.Refresh()
	c.background.Refresh()
	c.Layout(c.button.Size())
}

func (c *customButtonRender) Objects() []fyne.CanvasObject {
	return c.objects
}
func (c *customButtonRender) Destroy() {
}

func newButtonTapAnimation(bg *canvas.Rectangle, w fyne.Widget) *fyne.Animation {
	return fyne.NewAnimation(canvas.DurationStandard, func(done float32) {
		mid := (w.Size().Width - theme.Padding()) / 2
		size := mid * done
		bg.Resize(fyne.NewSize(size*2, w.Size().Height-theme.Padding()))
		bg.Move(fyne.NewPos(mid-size, theme.Padding()/2))

		r, g, bb, a := theme.PressedColor().RGBA()
		aa := uint8(a)
		fade := aa - uint8(float32(aa)*done)
		bg.FillColor = &color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(bb), A: fade}
		canvas.Refresh(bg)
	})
}
