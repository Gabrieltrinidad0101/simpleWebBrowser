package render

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
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
}

func (r render) container(tag *Tag) {
	container := canvas.NewRectangle(tag.Color)
	container.Move(fyne.NewPos(tag.X, tag.Y))
	container.FillColor = tag.Background
	container.StrokeColor = tag.BorderColor
	container.StrokeWidth = tag.BorderWidth
	container.Resize(fyne.NewSize(tag.Width, tag.Height))
	*r.Uis = append(*r.Uis, container)
}

func (r render) entry(tag *Tag) {
	input := widget.NewEntry()

	e := &widget.Entry{
		Wrapping: fyne.TextWrap(fyne.TextTruncateClip),
		
	}
	e.ExtendBaseWidget(e)

	input.SetPlaceHolder("Enter text...")
	input.Resize(fyne.NewSize(200, 30))
	input.CreateRenderer()

	*r.Uis = append(*r.Uis, input)
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
