package javascript

import (
	"simpleWebBrowser/render"

	"github.com/Gabrieltrinidad0101/Make-Language/src/features/class"
	"github.com/Gabrieltrinidad0101/Make-Language/src/languageContext"
	"github.com/Gabrieltrinidad0101/Make-Language/src/parser/parserStructs"
)

type Element struct {
	tag *render.Tag
}

type Style struct {
	element *render.Tag
}

func NewElement(tag *render.Tag) class.Class {
	buildClass := class.NewBuildClass(languageContext.NewContext(nil))
	style := class.NewBuildClass(languageContext.NewContext(nil))
	style.AddProperty("background", parserStructs.StringNode{
		Value: "red",
	})
	buildClass.AddProperty("style", class.Class{
		Context: style.Context,
		Name:    "style",
	})

	return class.Class{
		Context: buildClass.Context,
		Name:    "element",
	}
}
