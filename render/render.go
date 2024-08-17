package render

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type render struct {
	Uis *[]fyne.CanvasObject
}

func New() *render {
	uis := []fyne.CanvasObject{}
	return &render{
		Uis: &uis,
	}
}

func (r render) Render(tags []*Tag) {
	for _, tag := range tags {
		if tag.Name == "style" {
			continue
		}
		r.render(tag)
		r.Render(tag.Children)
	}
}

func (r render) setBackgroundColor(tag *Tag) {
	if tag.Background != nil {
		rect := canvas.NewRectangle(tag.Background)
		rect.Resize(fyne.NewSize(tag.Width, tag.Height))
		rect.Move(fyne.NewPos(tag.X, tag.Y))
		*r.Uis = append(*r.Uis, rect)
	}
}

func (r render) label(tag *Tag) {
	ui := canvas.NewText(tag.TextContent, tag.Color)
	ui.TextSize = tag.FontSize
	ui.Resize(fyne.NewSize(tag.Width, tag.Height))
	ui.Move(fyne.NewPos(tag.X, tag.Y))
	*r.Uis = append(*r.Uis, ui)
}

func (r render) container(tag *Tag) {
	container := canvas.NewRectangle(tag.Color)
	container.Move(fyne.NewPos(tag.X, tag.Y))
	container.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	container.StrokeColor = tag.BorderColor
	container.StrokeWidth = tag.BorderWidth
	container.Resize(fyne.NewSize(tag.Width, tag.Height))
	*r.Uis = append(*r.Uis, container)
}

func (r render) button(tag *Tag) {
	button := canvas.NewRectangle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	button.Move(fyne.NewPos(tag.X, tag.Y))
	button.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	button.StrokeColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	button.StrokeWidth = 2
	button.Resize(fyne.NewSize(tag.Width, tag.Height))
	*r.Uis = append(*r.Uis, button)
}

func (r render) render(tag *Tag) {
	if tag.Name == "text" {
		r.label(tag)
	} else {
		r.container(tag)
	}
}
