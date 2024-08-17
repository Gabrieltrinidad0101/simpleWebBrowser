package main

import (
	htmlparser "simpleWebBrowser/htmlParser"
	"simpleWebBrowser/render"
	"simpleWebBrowser/render/css"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Max Layout")
	myWindow.Resize(fyne.NewSize(1000, 1000))

	dom := htmlparser.Init()

	css_ := css.New()
	tags := css_.Run(dom.Children)
	uiRender := render.New()
	uiRender.Render(tags)

	myWindow.SetContent(container.NewWithoutLayout(*uiRender.Uis...))
	myWindow.ShowAndRun()
}
