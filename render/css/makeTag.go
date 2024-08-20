package css

import (
	"image/color"
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

func (c *CSS) border(border string) (float32, color.Color) {
	parts := strings.Split(border, " ")
	width := c.NumberDefault(parts[0], 0)
	color := c.Color(parts[2])
	return width, color
}

func (c *CSS) makeTag(element *parser.Element, parent *render.Tag) *render.Tag {
	properties := element.Properties
	tag := render.TAGS[element.Type_]
	tag.TextContent = element.TextContent

	var textDimention fyne.Size

	tag.Width = c.NumberDefault(properties["width"], 0)
	tag.Height = c.NumberDefault(properties["height"], 0)
	tag.PaddingLeft = c.NumberDefault(properties["padding-left"], tag.PaddingLeft)
	tag.PaddingTop = c.NumberDefault(properties["padding-top"], tag.PaddingTop)
	tag.PaddingBottom = c.NumberDefault(properties["padding-bottom"], tag.PaddingBottom)
	tag.PaddingRight = c.NumberDefault(properties["padding-right"], tag.PaddingRight)

	tag.MarginLeft = c.NumberDefault(properties["margin-left"], tag.MarginLeft)
	tag.MarginTop = c.NumberDefault(properties["margin-top"], tag.MarginTop)
	tag.MarginBottom = c.NumberDefault(properties["margin-bottom"], tag.MarginBottom)
	tag.MarginRight = c.NumberDefault(properties["margin-right"], tag.MarginRight)
	tag.JustifyContent = properties["justify-content"]

	if properties["display"] != "" {
		tag.Display = properties["display"]
	}

	if properties["background"] != "" {
		color := c.Color(properties["background"])
		tag.Background = &color
	}

	if properties["color"] != "" {
		color := c.Color(properties["color"])
		tag.Color = color
	} else if parent != nil {
		tag.Color = parent.Color
	}

	if properties["border"] != "" {
		width, color := c.border(properties["border"])
		tag.BorderWidth = width
		tag.BorderColor = color
	}

	if element.TextContent != "" {
		textDimention = fyne.MeasureText(tag.TextContent, tag.FontSize, fyne.TextStyle{})
	}

	if tag.Display == "block" && properties["width"] == "" && parent != nil {
		tag.Width = parent.Width
	}

	if properties["width"] == "" && textDimention.Width > 0 {
		tag.Width = textDimention.Width
	}

	if properties["height"] == "" && textDimention.Height > 0 {
		tag.Height = textDimention.Height
	}

	tag.Name = element.Type_
	return &tag
}
