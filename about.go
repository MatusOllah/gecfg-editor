package main

import (
	"fmt"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func showAboutDialog(a fyne.App, w fyne.Window) {
	dialog.NewInformation("About gecfg-editor", fmt.Sprintf(
		"%s version %s\nGo version %s\n\nBuilt with ❤️ by Matúš Ollah <github.com/MatusOllah>",
		a.Metadata().Name,
		a.Metadata().Version,
		runtime.Version(),
	), w).Show()
}
