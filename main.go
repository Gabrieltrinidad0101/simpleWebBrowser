package main

import (
	htmlparser "simpleWebBrowser/htmlParser"
	"simpleWebBrowser/render/position"
	"simpleWebBrowser/render/tags"
	"simpleWebBrowser/render/tags/h1"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Max Layout")
	myWindow.Resize(fyne.NewSize(1000, 1000))

	elements := htmlparser.Init()

	pos := position.NewPosition(0, 0)

	ui := []fyne.CanvasObject{}

	for _, element := range elements {
		h1_ := h1.NewH1(tags.Tag{
			TextContent: element.TextContent,
		}).Render(pos)
		ui = append(ui, h1_)
	}

	myWindow.SetContent(container.NewWithoutLayout(ui...))
	myWindow.ShowAndRun()
}
