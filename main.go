package main

import (
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

	pos := position.NewPosition(0, 0)

	h1_ := h1.NewH1(tags.Tag{
		TextContent: "Hello",
		IsLine:      true,
	}).Render(pos)

	h12_ := h1.NewH1(tags.Tag{
		TextContent: "Hello",
	}).Render(pos)

	myWindow.SetContent(container.NewWithoutLayout(
		h1_,
		h12_,
	))
	myWindow.ShowAndRun()
}
