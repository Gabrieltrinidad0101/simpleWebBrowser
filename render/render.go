package render

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type render struct{}

func New() *render {
	return &render{}
}

func (r render) Render(tag *Tag, ui *[]fyne.CanvasObject) *[]fyne.CanvasObject {
	tagUI := r.render(tag)
	*ui = append(*ui, tagUI)
	for _, child := range tag.Children {
		r.Render(child, ui)
	}

	return ui
}

func (r render) label(tag *Tag, uis *[]fyne.CanvasObject) {
	ui := canvas.NewText(tag.TextContent, color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	ui.TextSize = tag.FontSize
	ui.Resize(fyne.NewSize(tag.Width, tag.Height))
	if tag.Background != nil {
		rect := canvas.NewRectangle(*tag.Background)
		rect.Resize(fyne.NewSize(tag.Width, tag.Height))
		rect.Move(fyne.NewPos(tag.X, tag.Y))
		*uis = append(*uis, rect)
	}
	ui.Move(fyne.NewPos(tag.X, tag.Y))
	*uis = append(*uis, ui)
}

func (r render) button(tag *Tag, uis *[]fyne.CanvasObject) {
	label := &Tag{
		X:           tag.TextX,
		Y:           tag.TextY,
		TextContent: tag.TextContent,
		FontSize:    tag.FontSize,
	}
	button := canvas.NewRectangle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	button.Move(fyne.NewPos(tag.X, tag.Y))
	button.FillColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	button.StrokeColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	button.StrokeWidth = 2
	button.Resize(fyne.NewSize(tag.Width, tag.Height))
	*uis = append(*uis, button)
	r.label(label, uis)
}

func (r render) render(tag *Tag) *fyne.Container {
	uis := []fyne.CanvasObject{}
	if tag.Name == "button" {
		r.button(tag, &uis)
	} else {
		r.label(tag, &uis)
	}
	container := container.NewWithoutLayout(uis...)
	return container
}
