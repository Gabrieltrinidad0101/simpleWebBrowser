package main

import (
	htmlparser "simpleWebBrowser/htmlParser"
	"simpleWebBrowser/javascript"
	"simpleWebBrowser/render"
	"simpleWebBrowser/render/css"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow := myApp.NewWindow("Max Layout")
	myWindow.Resize(fyne.NewSize(1000, 1000))

	dom := htmlparser.Init()
	dom.Properties = map[string]string{
		"width":  "1000px",
		"height": "1000px",
	}

	css_ := css.New()
	tags := css_.Run(dom)
	uiRender := render.New()
	uiRender.Render(tags)

	javascript.New(tags)

	myWindow.SetContent(container.NewWithoutLayout(*uiRender.Uis...))
	myWindow.ShowAndRun()
}
