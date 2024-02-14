package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func makeUI(a fyne.App, w fyne.Window) fyne.CanvasObject {
	return container.NewBorder(
		makeToolbar(a, w),
		makeStatusBar(a, w),
		nil,
		nil,
		container.NewHSplit(makeList(w), makeDetails(a, w)),
	)
}
