package render

import (
	"image/color"
)

type Tag struct {
	BorderColor    color.Color
	BorderWidth    float32
	Height         float32
	Width          float32
	PaddingLeft    float32
	PaddingTop     float32
	PaddingBottom  float32
	PaddingRight   float32
	MarginLeft     float32
	MarginTop      float32
	MarginBottom   float32
	MarginRight    float32
	Padding        float32
	Display        string
	Name           string
	Gap            float32
	JustifyContent string
	Background     color.NRGBA
	TextContent    string
	ChildrenWidth  float32
	Color          color.NRGBA
	Children       []*Tag
	X              float32
	Y              float32
	FontSize       float32
	ChildX         float32
	ChildY         float32
}

var Root = Tag{
	Display:  "block",
	Height:   1000,
	Width:    1000,
	X:        0,
	Y:        0,
	Name:     "root",
	FontSize: 25,
	Color:    color.NRGBA{R: 0, G: 0, B: 0, A: 255},
}

var H1 = Tag{
	Display:    "block",
	Name:       "h1",
	FontSize:   30,
	MarginLeft: 10,
}

var Span = Tag{
	Display:  "inline",
	Name:     "h1",
	FontSize: 20,
}

var P = Tag{
	Display:  "block",
	Name:     "p",
	FontSize: 25,
}

var Div = Tag{
	Display:  "block",
	Name:     "div",
	FontSize: 25,
}

var Text = Tag{
	Display: "inline",
	Name:    "text",
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
	BorderColor:   DEFAULT_COLOR["black"],
	BorderWidth:   1,
	FontSize:      10,
	PaddingLeft:   5,
	PaddingTop:    5,
	PaddingRight:  5,
	PaddingBottom: 5,
}

var TAGS = map[string]Tag{
	"h1":     H1,
	"span":   Span,
	"p":      P,
	"button": Button,
	"div":    Div,
	"text":   Text,
	"root":   Root,
}
