package javascript

import (
	"simpleWebBrowser/render"

	"github.com/Gabrieltrinidad0101/Make-Language/src"
	"github.com/Gabrieltrinidad0101/Make-Language/src/features/class"
)

type Element struct {
	tag *render.Tag
}

type Style struct {
	element *render.Tag
}

func NewElement(makeLanguage *src.MakeLanguage, tag *render.Tag) class.Class {
	buildClass := class.NewBuildClass(nil)
	style := class.NewBuildClass(nil)
	buildClass.AddProperty("style")
}
