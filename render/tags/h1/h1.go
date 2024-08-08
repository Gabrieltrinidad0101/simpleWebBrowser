package h1

import (
	"image/color"
	"simpleWebBrowser/render/position"
	"simpleWebBrowser/render/tags"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewH1(tag tags.Tag) *h1 {

	if tag.Color == nil {
		tag.Color = &color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	}

	if tag.Height == nil {
		h := float32(24.0)
		tag.Height = &h
	}

	return &h1{
		Tag: tag,
	}
}

type h1 struct {
	tags.Tag
}

func (h *h1) Render(position *position.Position) *fyne.Container {
	h1 := canvas.NewText(h.TextContent, *h.Color)
	h1.TextSize = *h.Height
	container := container.NewWithoutLayout(h1)
	h1.Move(fyne.NewPos(position.X, position.Y))

	if h.IsLine {
		width := fyne.MeasureText(h1.Text, h1.TextSize, fyne.TextStyle{}).Width
		position.X += width
	} else {
		position.Y += *h.Height
	}

	return container
}
