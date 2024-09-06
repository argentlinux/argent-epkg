package main

import (
	"epkg-go/pkg/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("EPKG GUI Manage")

	w.SetContent(ui.CreateAppLayout(w))

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
