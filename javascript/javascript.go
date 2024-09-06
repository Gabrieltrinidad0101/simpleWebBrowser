package javascript

import (
	"fmt"
	"simpleWebBrowser/render"

	"github.com/Gabrieltrinidad0101/Make-Language/src"
	"github.com/Gabrieltrinidad0101/Make-Language/src/api"
	"github.com/Gabrieltrinidad0101/Make-Language/src/interprete/interpreteStructs"
	"github.com/Gabrieltrinidad0101/Make-Language/src/parser/parserStructs"
)

type Javascript struct{}

type Dom struct {
	root   *render.Tag
	render func()
}

func (d *Dom) GetElementById(params *[]interpreteStructs.IBaseElement) interface{} {
	query := (*params)[0]
	queryStr := query.GetValue().(string)
	tag := d.root.GetElementById(queryStr)
	if tag == nil {
		return parserStructs.NullNode{}
	}
	return NewElement(tag, d.render)
}

func consoleLog(params *[]interpreteStructs.IBaseElement) interface{} {
	query := (*params)[0]
	fmt.Println(query)
	return parserStructs.NullNode{}
}

func New(root *render.Tag, render func()) {
	makeLanguage := src.NewMakeLanguage("/home/gabriel/Desktop/go/simpleWebBrowser/javascript/conf.json", "/home/gabriel/Desktop/go/simpleWebBrowser/javascript/index.js")

	dom := &Dom{root: root, render: render}

	methods := api.Methods{
		"getElementById": dom.GetElementById,
	}

	makeLanguage.AddClass("Document", api.CustomClassValues{
		Methods: methods,
	})

	makeLanguage.AddFunction("console", consoleLog)

	err := makeLanguage.Run()
	if err != nil {
		panic(err)
	}
}
