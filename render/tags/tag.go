package tags

import "image/color"

type Tag struct {
	Height      *float32
	IsLine      bool
	Margin      float32
	Padding     float32
	TextContent string
	Name        string
	Color       *color.NRGBA
}
