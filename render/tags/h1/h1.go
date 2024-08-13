package h1

import (
	"image/color"
	"simpleWebBrowser/render/tags"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewH1() *h1 {
	return &h1{}
}

type h1 struct {
}

func (h *h1) Render(tag *tags.Tag) *fyne.Container {
	ui := []fyne.CanvasObject{}
	h1 := canvas.NewText(tag.TextContent, color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	h1.TextSize = 24
	if tag.Background != nil {
		rect := canvas.NewRectangle(*tag.Background)
		rect.Resize(fyne.NewSize(tag.Width, tag.Height))
		rect.Move(fyne.NewPos(tag.X, tag.Y))
		ui = append(ui, rect)
	}
	ui = append(ui, h1)
	container := container.NewWithoutLayout(ui...)
	h1.Move(fyne.NewPos(tag.X, tag.Y))
	return container
}
