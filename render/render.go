package render

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type render struct {
	Uis *[]fyne.CanvasObject
}

var uiElements = map[string]fyne.CanvasObject{}

func UiElements() *map[string]fyne.CanvasObject {
	return &uiElements
}

func New() *render {
	uis := []fyne.CanvasObject{}
	return &render{
		Uis: &uis,
	}
}

func (r render) Render(tags *Tag) {
	for _, tag := range tags.Children {
		if tag.Name == "style" {
			continue
		}
		r.render(tag)
		r.Render(tag)
	}
}

func (r render) label(tag *Tag) {
	ui := canvas.NewText(tag.TextContent, tag.Color)
	ui.TextSize = tag.FontSize
	ui.Resize(fyne.NewSize(tag.Width, tag.Height))
	ui.Move(fyne.NewPos(tag.X, tag.Y))
	*r.Uis = append(*r.Uis, ui)
	uiElements[tag.Id] = ui
}

func (r render) container(tag *Tag) {
	container := canvas.NewRectangle(tag.Color)
	container.Move(fyne.NewPos(tag.X, tag.Y))
	container.FillColor = tag.Background
	container.StrokeColor = tag.BorderColor
	container.StrokeWidth = tag.BorderWidth
	container.Resize(fyne.NewSize(tag.Width, tag.Height))
	*r.Uis = append(*r.Uis, container)
	uiElements[tag.Id] = container
}

func (r render) entry(tag *Tag) {
	r.container(tag)
	input := NewCustomEntry()
	x := tag.X + tag.PaddingLeft + tag.BorderWidth
	y := tag.Y + tag.PaddingTop + tag.BorderWidth
	w := tag.Width - tag.BorderWidth*2
	h := tag.Height - tag.BorderWidth*2
	input.Move(fyne.NewPos(x, y))
	input.Resize(fyne.NewSize(w, h))
	*r.Uis = append(*r.Uis, input)
	uiElements[tag.Id] = input
}

func (r render) render(tag *Tag) {
	if tag.Name == "text" {
		r.label(tag)
	} else if tag.Name == "input" {
		r.entry(tag)
	} else {
		r.container(tag)
	}
}
