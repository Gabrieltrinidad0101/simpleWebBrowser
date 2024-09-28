package javascript

import (
	"fmt"
	"simpleWebBrowser/render"
	"simpleWebBrowser/utils"

	"github.com/Gabrieltrinidad0101/Make-Language/src/features/class"
	"github.com/Gabrieltrinidad0101/Make-Language/src/features/str"
	"github.com/Gabrieltrinidad0101/Make-Language/src/interprete/interpreteStructs"
	"github.com/Gabrieltrinidad0101/Make-Language/src/languageContext"
)

type element struct {
	tag    *render.Tag
	render func()
}

func (e element) MakeStyle() class.Class {
	background := str.NewString("", nil)
	color := str.NewString("black", nil)
	fontSize := str.NewString(fmt.Sprint(e.tag.FontSize), nil)

	style := class.NewBuildClass(languageContext.NewContext(nil))

	style.AddProperty("background", &interpreteStructs.VarType{
		Value: background,
		OnUpdateVariable: func(value interface{}) {
			e.tag.Background = utils.Color(value.(string))
			e.render()
		},
	})

	style.AddProperty("color", &interpreteStructs.VarType{
		Value: color,
		OnUpdateVariable: func(value interface{}) {
			*e.tag.Color = utils.Color(value.(string))
			e.render()
		},
	})

	style.AddProperty("fontSize", &interpreteStructs.VarType{
		Value: fontSize,
		OnUpdateVariable: func(value interface{}) {
			*e.tag.FontSize = utils.NumberDefault(value.(string), *e.tag.Parent.Height, *e.tag.FontSize)
			e.render()
		},
	})

	styleClass := class.Class{
		Context: style.Context,
		Name:    "style",
	}

	return styleClass
}

func NewElement(tag *render.Tag, render func()) class.Class {
	elementClass := class.NewBuildClass(languageContext.NewContext(nil))

	ele := element{
		tag:    tag,
		render: render,
	}

	style := ele.MakeStyle()

	text := str.NewString(*tag.TextContent, nil)

	elementClass.AddProperty("style", &interpreteStructs.VarType{Value: style})
	elementClass.AddProperty("textContent", &interpreteStructs.VarType{
		Value: text,
		OnUpdateVariable: func(value interface{}) {
			textChild := (*tag.Children)[0]
			*textChild.TextContent = value.(string)
			render()
		},
	})
	return class.Class{
		Context: elementClass.Context,
		Name:    "element",
	}
}
