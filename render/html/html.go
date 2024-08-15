package html

import (
	"math"
	"simpleWebBrowser/render"

	"fyne.io/fyne/v2"
	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

type HTML struct {
	x                  float32
	y                  float32
	lastIsInline       bool
	lastIsInlineHeight float32
}

type BasicPosition struct {
	X float32
	Y float32
	W float32
	H float32
}

func (r HTML) getLabelCenter(tag *render.Tag) *BasicPosition {
	textDimention := fyne.MeasureText(tag.TextContent, tag.FontSize, fyne.TextStyle{})

	x := tag.X + tag.PaddingLeft
	y := tag.Y + tag.PaddingTop

	return &BasicPosition{
		X: x,
		Y: y,
		W: textDimention.Width,
		H: textDimention.Height,
	}
}

func (c *HTML) makeTag(element *parser.Element) *render.Tag {
	tag := render.TAGS[element.Type_]
	tag.TextContent = element.TextContent
	tag.Name = element.Type_
	tag.X = c.x
	tag.Y = c.y

	if tag.Name == "button" {
		labelPosition := c.getLabelCenter(&tag)
		tag.TextX = labelPosition.X
		tag.TextY = labelPosition.Y
		tag.Width = labelPosition.W + tag.PaddingLeft + tag.PaddingRight
		tag.Height = labelPosition.H + tag.PaddingTop + tag.PaddingBottom
	}

	return &tag
}

func (c *HTML) Run(dom *parser.Element, parent *render.Tag) *render.Tag {
	tag := c.makeTag(dom)

	for _, element := range dom.Children {
		childTag := c.Run(element, tag)
		tag.Children = append(tag.Children, childTag)
		tag.ChildrenWidth += childTag.Width

		if dom.Properties["width"] == "" {
			tag.Width += childTag.Width
		}

		if dom.Properties["height"] == "" {
			tag.Height += childTag.Height
		}
	}

	if tag.Display == "inline" {
		c.x += tag.Width
		c.lastIsInlineHeight = float32(math.Max(float64(c.lastIsInlineHeight), float64(tag.Height)))
	} else if parent != nil {
		if c.lastIsInline {
			c.y += c.lastIsInlineHeight
			tag.Y = c.y
		}
		c.y += tag.Height
		tag.X = 0
		c.x = 0
	}
	c.lastIsInline = tag.Display == "inline"

	return tag
}

func New() *HTML {
	return &HTML{}
}
