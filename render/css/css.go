package css

import (
	"image/color"
	"math"
	"simpleWebBrowser/render/tags"
	"strconv"
	"strings"

	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

var DEFAULT_COLOR = map[string]color.NRGBA{
	"red":    {R: 255, G: 0, B: 0, A: 255},
	"blue":   {R: 0, G: 0, B: 255, A: 255},
	"green":  {R: 0, G: 255, B: 0, A: 255},
	"yellow": {R: 255, G: 255, B: 0, A: 255},
	"orange": {R: 255, G: 165, B: 0, A: 255},
	"purple": {R: 128, G: 0, B: 128, A: 255},
	"pink":   {R: 255, G: 192, B: 203, A: 255},
}

type CSS struct {
}

func (c *CSS) Number(value string, defaultValue float32) float32 {
	value = strings.ReplaceAll(value, "px", "")
	number, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultValue
	}
	return float32(number)
}

func (c *CSS) mapToStruct(properties map[string]string) *tags.Tag {
	tag := &tags.Tag{}
	if strings.Contains(properties["width"], "px") {
		tag.Width = c.Number(properties["width"], 0)
	}

	if strings.Contains(properties["height"], "px") {
		tag.Height = c.Number(properties["height"], 0)
	}
	tag.Display = properties["display"]
	if properties["background"] != "" {
		color := c.Color(properties["background"])
		tag.Background = &color
	}
	return tag
}

func (c *CSS) Color(colorStr string) color.NRGBA {
	colorRRBA, ok := DEFAULT_COLOR[colorStr]
	if !ok {
		return color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	}
	return colorRRBA
}

func (c *CSS) Run(dom *parser.Element) *tags.Tag {
	tag := c.mapToStruct(dom.Properties)
	tag.TextContent = dom.TextContent

	for _, element := range dom.Children {
		childTag := c.Run(element)
		tag.Children = append(tag.Children, childTag)
		tag.ChildrenWidth += childTag.Width
		if dom.Properties["width"] == "" {
			tag.Width += childTag.Width
		}

		if dom.Properties["height"] == "" {
			tag.Height += childTag.Height
		}
	}

	if tag.Display == "flex" {
		gap := math.Max(float64((tag.Width-tag.ChildrenWidth)/float32(len(tag.Children)-1)), 0+float64(tag.Gap))
		lastWidth := float32(0.0)
		for _, child := range tag.Children {
			child.X = lastWidth
			child.Y = tag.Y
			lastWidth += child.Width + float32(gap)
		}
	}

	return tag
}

func New() *CSS {
	return &CSS{}
}