package css

import (
	"fmt"
	"image/color"
	"simpleWebBrowser/render"
	"simpleWebBrowser/utils"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

func (c *CSS) border(border string) (float32, color.Color) {
	parts := strings.Split(border, " ")
	width := utils.NumberDefault(parts[0], 0, 0)
	color := utils.Color(parts[2])
	return width, color
}

func (c *CSS) makeTag(element *parser.Element, parent *render.Tag) *render.Tag {
	properties := element.Properties
	tag := render.TAGS[element.Type_]
	*tag.TextContent = element.TextContent
	tag.Parent = parent
	var textDimention fyne.Size

	tag.Width = utils.NumberDefault(properties["width"], parent.Width, tag.Width)
	*tag.Height = utils.NumberDefault(properties["height"], *parent.Height, *tag.Height)
	tag.PaddingLeft = utils.NumberDefault(properties["padding-left"], parent.Width, tag.PaddingLeft)
	tag.PaddingTop = utils.NumberDefault(properties["padding-top"], *parent.Height, tag.PaddingTop)
	tag.PaddingBottom = utils.NumberDefault(properties["padding-bottom"], *parent.Height, tag.PaddingBottom)
	tag.PaddingRight = utils.NumberDefault(properties["padding-right"], parent.Width, tag.PaddingRight)

	tag.MarginLeft = utils.NumberDefault(properties["margin-left"], parent.Width, tag.MarginLeft)
	tag.MarginTop = utils.NumberDefault(properties["margin-top"], *parent.Height, tag.MarginTop)
	tag.MarginBottom = utils.NumberDefault(properties["margin-bottom"], *parent.Height, tag.MarginBottom)
	tag.MarginRight = utils.NumberDefault(properties["margin-right"], parent.Width, tag.MarginRight)
	tag.JustifyContent = properties["justify-content"]
	tag.Gap = utils.NumberDefault(properties["gap"], parent.Width, 0)

	if properties["display"] != "" {
		tag.Display = properties["display"]
	}

	if properties["background"] != "" {
		color := utils.Color(properties["background"])
		tag.Background = color
	}

	if properties["color"] != "" {
		color := utils.Color(properties["color"])
		tag.Color = &color
	} else if parent != nil {
		tag.Color = parent.Color
	}

	if properties["font-size"] != "" {
		*tag.FontSize = utils.NumberDefault(properties["font-size"], *parent.Height, *tag.FontSize)
	} else if parent != nil {
		tag.FontSize = parent.FontSize
	}

	if properties["border"] != "" {
		width, color := c.border(properties["border"])
		tag.BorderWidth = width
		tag.BorderColor = color
	}

	if element.TextContent != "" {
		textDimention = fyne.MeasureText(*tag.TextContent, *tag.FontSize, fyne.TextStyle{})
	}

	if tag.Display == "block" && properties["width"] == "" && parent != nil {
		tag.Width = parent.Width
	}

	if properties["width"] == "" && textDimention.Width > 0 {
		tag.Width = textDimention.Width
	}

	if properties["height"] == "" && textDimention.Height > 0 {
		*tag.Height = textDimention.Height
	}
	tag.Id = properties["id"]
	tag.Name = element.Type_
	tag.UUID = fmt.Sprint(time.Now().UnixMilli())
	tag.Children = new([]*render.Tag)
	return &tag
}
