package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(widget.NewLabel("Hello, World!"))
	w.Resize(fyne.NewSize(400, 600))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
