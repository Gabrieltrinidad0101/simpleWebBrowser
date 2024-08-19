package css

import (
	"fmt"
	"image/color"
	"math"
	"simpleWebBrowser/render"

	"fyne.io/fyne/v2"
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
	"black":  {R: 0, G: 0, B: 0, A: 255},
}

type BasicPosition struct {
	X float32
	Y float32
	W float32
	H float32
}

type CSS struct {
}

func (r *CSS) getLabelCenter(tag *render.Tag) *BasicPosition {
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

func (c *CSS) print(tag *render.Tag) {
	fmt.Println("name: ", tag.Name, " width: ", tag.Width, " height: ", tag.Height, "display: ", tag.Display, "x: ", tag.X, "y: ", tag.Y)
}

func (c *CSS) Color(colorStr string) color.NRGBA {
	colorRRBA, ok := DEFAULT_COLOR[colorStr]
	if !ok {
		return color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	}
	return colorRRBA
}

func (c *CSS) Run(root *parser.Element) *render.Tag {
	tag := c.run(root, &render.Root)
	c.resetPosition(tag, &render.Root)
	return tag
}

func (c *CSS) resetPosition(tag *render.Tag, parent *render.Tag) {
	c.print(tag)
	tag.ChildX = tag.X + tag.BorderWidth
	tag.ChildY = tag.Y + tag.BorderWidth

	gap := math.Max(float64((tag.Width-tag.BorderWidth*2-tag.ChildrenWidth)/float32(len(tag.Children)-1))+float64(tag.Gap), 0)

	for _, child := range tag.Children {
		child.X = tag.ChildX + child.MarginLeft
		child.Y = tag.ChildY + child.MarginTop

		if child.Display == "inline" || child.Display == "inline-block" {
			tag.ChildX += child.X - tag.ChildX + child.Width + child.MarginRight
		} else if tag.Display == "flex" {
			child.X = tag.ChildX
			child.Y = tag.ChildY
			tag.ChildX += child.Width + float32(gap)
		} else {
			tag.ChildY += child.Y + child.Height + child.MarginBottom
		}
	}

	for _, child := range tag.Children {
		c.resetPosition(child, tag)
	}

}

func (c *CSS) run(dom *parser.Element, parent *render.Tag) *render.Tag {
	tag := c.makeTag(dom, parent)
	biggerWidth := float32(0.0)
	biggerHeight := float32(0.0)
	totalChildrenWidth := float32(0.0)
	totalChildrenHeight := float32(0.0)

	for _, element := range dom.Children {
		childTag := c.run(element, tag)
		tag.Children = append(tag.Children, childTag)
		biggerWidth = float32(math.Max(float64(biggerWidth), float64(childTag.Width)))
		biggerHeight = float32(math.Max(float64(biggerHeight), float64(childTag.Height)))
		totalChildrenWidth += childTag.Width
		totalChildrenHeight += childTag.Height
	}

	tag.ChildrenWidth = totalChildrenWidth

	if tag.Display == "inline" && len(dom.Children) > 0 {
		tag.Width = biggerWidth
		tag.Height = biggerHeight
	}

	if tag.Width <= 0 && tag.Display == "inline-block" {
		tag.Width = totalChildrenWidth
	}

	if tag.Height <= 0 && tag.Display == "block" {
		tag.Height = totalChildrenHeight
	}

	tag.Height += tag.BorderWidth * 2
	tag.Width += tag.BorderWidth * 2

	return tag
}

func New() *CSS {
	return &CSS{}
}
