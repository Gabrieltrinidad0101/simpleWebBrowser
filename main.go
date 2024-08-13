package main

import (
	htmlparser "simpleWebBrowser/htmlParser"
	"simpleWebBrowser/render/css"
	"simpleWebBrowser/render/tags"
	"simpleWebBrowser/render/tags/h1"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func render(tag *tags.Tag, ui *[]fyne.CanvasObject) *[]fyne.CanvasObject {
	tagUI := h1.NewH1().Render(tag)
	*ui = append(*ui, tagUI)
	for _, child := range tag.Children {
		render(child, ui)
	}

	return ui
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Max Layout")
	myWindow.Resize(fyne.NewSize(1000, 1000))

	dom := htmlparser.Init()

	css_ := css.New()
	tags := css_.Run(dom)
	ui := []fyne.CanvasObject{}

	for _, child := range tags.Children {
		render(child, &ui)
	}

	myWindow.SetContent(container.NewWithoutLayout(ui...))
	myWindow.ShowAndRun()
}
