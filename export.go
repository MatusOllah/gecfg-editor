package main

import (
	"encoding/json"
	"log/slog"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func exportJSON(w fyne.Window) {
	var indent bool

	dialog.NewForm("Export JSON", "OK", "Cancel", []*widget.FormItem{widget.NewFormItem("Indent?", widget.NewCheck("", func(b bool) {
		indent = b
	}))}, func(b bool) {
		if !b {
			return
		}

		dlg := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				slog.Info(err.Error())
				dialog.NewError(err, w).Show()
				return
			}

			path := uc.URI().Path()

			slog.Info("exporting JSON", "path", path, "indent", indent)

			var b []byte
			if indent {
				b, err = json.MarshalIndent(theMap, "", "\t")
				if err != nil {
					slog.Info(err.Error())
					dialog.NewError(err, w).Show()
					return
				}
			} else {
				b, err = json.Marshal(theMap)
				if err != nil {
					slog.Info(err.Error())
					dialog.NewError(err, w).Show()
					return
				}
			}

			if err := os.WriteFile(path, b, 0666); err != nil {
				slog.Info(err.Error())
				dialog.NewError(err, w).Show()
				return
			}
		}, w)
		dlg.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
		dlg.Show()

	}, w).Show()
}
