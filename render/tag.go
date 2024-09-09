package render

import (
	"image/color"

	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

type Tag struct {
	Parent         *Tag
	BorderColor    color.Color
	BorderWidth    float32
	Height         *float32
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
	Color          *color.NRGBA
	Children       []*Tag
	X              float32
	Y              float32
	FontSize       *float32
	ChildX         float32
	ChildY         float32
	Id             string
	UUID           string
	ClassString    string
}

func (e *Tag) QuerySelector(textQuery string) *Tag {
	query := parser.NewQuery("")
	queries := query.Analyze(textQuery)
	tags := make([]*Tag, 0, len(queries))
	tags1 := e.querySelector(e, queries, 0, false, &tags)
	return (*tags1)[0]
}

func forEach(element *Tag, cb func(*Tag) bool) *Tag {
	for _, child := range element.Children {
		stop := cb(child)
		if stop {
			return child
		}
		elemt := forEach(child, cb)
		if elemt != nil {
			return elemt
		}
	}
	return nil
}

func (e *Tag) GetElementById(id string) *Tag {
	value := e.Id

	if value == id {
		return e
	}

	return forEach(e, func(e *Tag) bool {
		value := e.Id
		return value == id
	})
}

func (e *Tag) QuerySelectorAll(textQuery string) *[]*Tag {
	query := parser.NewQuery("")
	queries := query.Analyze(textQuery)
	elements := make([]*Tag, 0, len(queries))
	return e.querySelector(e, queries, 0, true, &elements)
}

func (e *Tag) setQueryFalses(queries []*parser.QueryData) {
	for _, query := range queries {
		query.IsFound = false
	}
}

func (e *Tag) querySelector(element *Tag, queries []*parser.QueryData, index int, getAll bool, elements *[]*Tag) *[]*Tag {
	for _, child := range element.Children {
		query := (queries)[index]

		if query.TypeSearch == "id" {
			query.IsFound = child.Id == query.Search
		}

		if query.TypeSearch == "class" {
			query.IsFound = child.ClassString == query.Search
		}

		if query.TypeSearch == "element" {
			query.IsFound = child.Name == query.Search
		}

		if query.IsFound {
			if index < len(queries)-1 {
				index++
			}
		}

		if (queries)[len(queries)-1].IsFound {
			*elements = append(*elements, child)
			e.setQueryFalses(queries)
			if !getAll {
				return elements
			}
		}

		elemt := e.querySelector(child, queries, index, getAll, elements)
		query.IsFound = elemt != nil

		if !getAll && len(*elemt) == 1 {
			return elemt
		}
	}
	return elements
}
