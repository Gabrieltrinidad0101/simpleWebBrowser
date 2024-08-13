package tags

import (
	"image/color"
	"strconv"
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
}

func Number(value string, defaultValue float32) float32 {
	number, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultValue
	}
	return float32(number)
}
