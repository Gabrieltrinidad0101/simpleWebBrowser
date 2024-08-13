package render

var H1 = Tag{
	Height:   24,
	Display:  "block",
	Name:     "h1",
	Margin:   2,
	FontSize: 30,
}

var TAGS = map[string]Tag{
	"h1": H1,
}
