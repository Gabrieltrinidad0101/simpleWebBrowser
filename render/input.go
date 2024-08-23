package render

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// CustomEntry is a customized Entry widget with a custom border.
type CustomEntry struct {
	widget.Entry
}

// CreateRenderer returns a new renderer for the custom entry.
func (e *CustomEntry) CreateRenderer() fyne.WidgetRenderer {
	// Call the default Entry renderer
	renderer := e.Entry.CreateRenderer()

	// Override the background and border
	for _, obj := range renderer.Objects() {
		if rect, ok := obj.(*canvas.Rectangle); ok {
			// Customize the border (this example removes it)
			rect.StrokeColor = color.Transparent // Remove border
			rect.FillColor = color.Transparent   // Remove background
		}
	}

	e.Entry.OnCursorChanged = func() {
		for _, obj := range renderer.Objects() {
			if rect, ok := obj.(*canvas.Rectangle); ok {
				// Customize the border (this example removes it)
				rect.StrokeColor = color.Transparent // Remove border
				rect.FillColor = color.Transparent   // Remove background
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
