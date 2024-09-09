package render

import (
	"image/color"
	"simpleWebBrowser/utils"
)

var Root = Tag{
	Display:  "block",
	Height:   utils.RefNumber(1000),
	Width:    1000,
	X:        0,
	Y:        0,
	Name:     "root",
	FontSize: utils.RefNumber(25),
	Color:    &color.NRGBA{R: 0, G: 0, B: 0, A: 255},
}

var H1 = Tag{
	Display:    "block",
	Name:       "h1",
	FontSize:   utils.RefNumber(30),
	MarginLeft: 10,
	Height:     utils.RefNumber(0),
}

var Span = Tag{
	Display:  "inline",
	Name:     "h1",
	FontSize: utils.RefNumber(20),
}

var P = Tag{
	Display:  "block",
	Name:     "p",
	FontSize: utils.RefNumber(25),
}

var Div = Tag{
	Display:  "block",
	Name:     "div",
	FontSize: utils.RefNumber(25),
}

var Text = Tag{
	Display: "inline",
	Name:    "text",
	Height:  utils.RefNumber(0),
}

var Button = Tag{
	Display: "inline",
	Name:    "button",
	Background: color.NRGBA{
		R: 233,
		G: 233,
		B: 237,
		A: 255,
	},
	BorderColor:   utils.DEFAULT_COLOR["black"],
	BorderWidth:   1,
	FontSize:      utils.RefNumber(10),
	PaddingLeft:   5,
	PaddingTop:    5,
	PaddingRight:  5,
	PaddingBottom: 5,
}

var Input = Tag{
	Display:  "inline",
	Name:     "input",
	FontSize: utils.RefNumber(15),
	Width:    75,
	Height:   utils.RefNumber(30),
}

var TAGS = map[string]Tag{
	"h1":     H1,
	"span":   Span,
	"p":      P,
	"button": Button,
	"div":    Div,
	"text":   Text,
	"input":  Input,
	"root":   Root,
}
