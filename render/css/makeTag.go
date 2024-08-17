package css

import (
	"simpleWebBrowser/render"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

func (c *CSS) NumberDefault(value string, defaultValue float32) float32 {
	value = strings.ReplaceAll(value, "px", "")
	number, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultValue
	}
	return float32(number)
}

func (c *CSS) makeTag(element *parser.Element) *render.Tag {
	properties := element.Properties
	tag := render.TAGS[element.Type_]

	var textDimention fyne.Size

	tag.Width = c.NumberDefault(properties["width"], 0)
	tag.Height = c.NumberDefault(properties["height"], 0)
	tag.PaddingLeft = c.NumberDefault(properties["padding-left"], 0)
	tag.PaddingTop = c.NumberDefault(properties["padding-top"], 0)
	tag.PaddingBottom = c.NumberDefault(properties["padding-bottom"], 0)
	tag.PaddingBottom = c.NumberDefault(properties["padding-right"], 0)
	if properties["display"] != "" {
		tag.Display = properties["display"]
	}

	if properties["background"] != "" {
		color := c.Color(properties["background"])
		tag.Background = &color
	}

	if element.TextContent != "" {
		textDimention = fyne.MeasureText(tag.TextContent, tag.FontSize, fyne.TextStyle{})
	}

	if properties["width"] == "" && textDimention.Width > 0 {
		tag.Width = textDimention.Width
	}

	if properties["height"] == "" && textDimention.Height > 0 {
		tag.Height = textDimention.Height
	}

	if tag.Name == "button" {
		labelPosition := c.getLabelCenter(&tag)
		tag.TextX = labelPosition.X
		tag.TextY = labelPosition.Y
		tag.Width = labelPosition.W + tag.PaddingLeft + tag.PaddingRight
		tag.Height = labelPosition.H + tag.PaddingTop + tag.PaddingBottom
	}

	tag.TextContent = element.TextContent
	tag.Name = element.Type_
	tag.X = c.x 
	tag.Y = c.y

	return &tag
}
