package render

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Tag struct {
	Height         float32 `json:"height"`
	Width          float32 `json:"width"`
	Margin         float32 `json:"margin"`
	Padding        float32 `json:"padding"`
	Display        string  `json:"display"`
	Name           string  `json:"name"`
	Gap            float64 `json:"gap"`
	JustifyContent string  `json:"justifyContent"`
	Background     *color.NRGBA
	TextContent    string
	ChildrenWidth  float32
	Color          color.NRGBA
	Children       []*Tag
	X              float32
	Y              float32
	FontSize       float32
}

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
	*uis = append(*uis, ui)
	ui.TextSize = 24
	if tag.Background != nil {
		rect := canvas.NewRectangle(*tag.Background)
		rect.Resize(fyne.NewSize(tag.Width, tag.Height))
		rect.Move(fyne.NewPos(tag.X, tag.Y))
		*uis = append(*uis, rect)
	}
	ui.Move(fyne.NewPos(tag.X, tag.Y))
}

func (r render) render(tag *Tag) *fyne.Container {
	uis := []fyne.CanvasObject{}
	r.label(tag, &uis)
	container := container.NewWithoutLayout(uis...)
	return container
}
