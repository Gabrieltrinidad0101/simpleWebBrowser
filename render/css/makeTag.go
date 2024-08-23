package css

import (
	"image/color"
	"simpleWebBrowser/render"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

func (c *CSS) percentage(value string, container float32) (float32, error) {
	value = strings.ReplaceAll(value, "%", "")
	number, err := strconv.ParseFloat(value, 32)
	return float32(number) * container / 100, err
}

func (c *CSS) NumberDefault(value string, container float32, defaultValue float32) float32 {
	if strings.Contains(value, "px") {
		value = strings.ReplaceAll(value, "px", "")
		number, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return defaultValue
		}
		return float32(number)
	}

	if strings.Contains(value, "%") {
		number, err := c.percentage(value, container)
		if err != nil {
			return defaultValue
		}
		return float32(number)
	}

	return defaultValue

}

func (c *CSS) border(border string) (float32, color.Color) {
	parts := strings.Split(border, " ")
	width := c.NumberDefault(parts[0], 0, 0)
	color := c.Color(parts[2])
	return width, color
}

func (c *CSS) makeTag(element *parser.Element, parent *render.Tag) *render.Tag {
	properties := element.Properties
	tag := render.TAGS[element.Type_]
	tag.TextContent = element.TextContent

	var textDimention fyne.Size

	tag.Width = c.NumberDefault(properties["width"], parent.Width, tag.Width)
	tag.Height = c.NumberDefault(properties["height"], parent.Height, tag.Height)
	tag.PaddingLeft = c.NumberDefault(properties["padding-left"], parent.Width, tag.PaddingLeft)
	tag.PaddingTop = c.NumberDefault(properties["padding-top"], parent.Height, tag.PaddingTop)
	tag.PaddingBottom = c.NumberDefault(properties["padding-bottom"], parent.Height, tag.PaddingBottom)
	tag.PaddingRight = c.NumberDefault(properties["padding-right"], parent.Width, tag.PaddingRight)

	tag.MarginLeft = c.NumberDefault(properties["margin-left"], parent.Width, tag.MarginLeft)
	tag.MarginTop = c.NumberDefault(properties["margin-top"], parent.Height, tag.MarginTop)
	tag.MarginBottom = c.NumberDefault(properties["margin-bottom"], parent.Height, tag.MarginBottom)
	tag.MarginRight = c.NumberDefault(properties["margin-right"], parent.Width, tag.MarginRight)
	tag.JustifyContent = properties["justify-content"]
	tag.Gap = c.NumberDefault(properties["gap"], parent.Width, 0)

	if properties["display"] != "" {
		tag.Display = properties["display"]
	}

	if properties["background"] != "" {
		color := c.Color(properties["background"])
		tag.Background = color
	}

	if properties["color"] != "" {
		color := c.Color(properties["color"])
		tag.Color = color
	} else if parent != nil {
		tag.Color = parent.Color
	}

	if properties["font-size"] != "" {
		tag.FontSize = c.NumberDefault(properties["font-size"], parent.Height, tag.FontSize)
	} else if parent != nil {
		tag.FontSize = parent.FontSize
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
	tag.Id = string(time.Now().UnixMilli())
	return &tag
}
