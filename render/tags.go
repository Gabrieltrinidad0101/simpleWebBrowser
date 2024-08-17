package render

import "image/color"

type Tag struct {
	BorderColor    color.Color
	BorderWidth         float32
	Height         float32
	Width          float32
	PaddingLeft    float32
	PaddingTop     float32
	PaddingBottom  float32
	PaddingRight   float32
	Padding        float32
	Display        string
	Name           string
	Gap            float32
	JustifyContent string
	Background     *color.NRGBA
	TextContent    string
	ChildrenWidth  float32
	Color          color.NRGBA
	Children       []*Tag
	TextX          float32
	TextY          float32
	X              float32
	Y              float32
	FontSize       float32
}

var H1 = Tag{
	Display:  "block",
	Name:     "h1",
	FontSize: 30,
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
	Display:  "inline",
	Name:     "text",
	FontSize: 25,
}

var Button = Tag{
	Display:       "inline",
	Name:          "button",
	FontSize:      20,
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
}
