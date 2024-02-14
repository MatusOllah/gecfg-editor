package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeStatusBar(a fyne.App, w fyne.Window) fyne.CanvasObject {
	return container.NewBorder(widget.NewSeparator(), nil, nil, nil, container.NewVBox(
		widget.NewLabel("v"+a.Metadata().Version),
	))
}
