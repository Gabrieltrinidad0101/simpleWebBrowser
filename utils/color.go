package utils

import (
	"image/color"
	"simpleWebBrowser/render"
)

func Color(colorStr string) color.NRGBA {
	colorRRBA, ok := render.DEFAULT_COLOR[colorStr]
	if !ok {
		return color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	}
	return colorRRBA
}
