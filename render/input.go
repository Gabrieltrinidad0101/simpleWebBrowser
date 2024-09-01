package render

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type CustomEntry struct {
	widget.Entry
}

func (e *CustomEntry) CreateRenderer() fyne.WidgetRenderer {
	renderer := e.Entry.CreateRenderer()
	for _, obj := range renderer.Objects() {
		if rect, ok := obj.(*canvas.Rectangle); ok {
			rect.StrokeColor = color.Transparent
			rect.FillColor = color.Transparent
		}
	}

	e.Entry.OnCursorChanged = func() {
		for _, obj := range renderer.Objects() {
			if rect, ok := obj.(*canvas.Rectangle); ok {
				rect.StrokeColor = color.Transparent
				rect.FillColor = color.Transparent
			}
		}
	}

	return renderer
}

func NewCustomEntry() *CustomEntry {
	entry := &CustomEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}
