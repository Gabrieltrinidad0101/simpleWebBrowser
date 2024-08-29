package main

import (
	"simpleWebBrowser/render"

	"github.com/Gabrieltrinidad0101/Make-Language/src"
	"github.com/Gabrieltrinidad0101/Make-Language/src/interprete/interpreteStructs"
	"github.com/Gabrieltrinidad0101/Make-Language/src/parser/parserStructs"
)

type Javascript struct{}

type Dom struct {
	root *render.Tag
}

func (d *Dom) GetElementById(params *[]interpreteStructs.IBaseElement) interface{} {
	query := (*params)[0]
	queryStr := query.GetValue().(string)
	tag := d.root.GetElementById(queryStr)
	if tag == nil {
		return parserStructs.NullNode{}
	}
}

func New(root *render.Tag) {
	makeLanguage := src.NewMakeLanguage("./conf.json", "./index.js")
	makeLanguage.AddClass("Document")
}
