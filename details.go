package main

import (
	"fmt"
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeDetails(a fyne.App, w fyne.Window) fyne.CanvasObject {
	typeLbl := widget.NewLabel("")
	keyLbl := widget.NewLabel("")
	valueLbl := widget.NewLabel("")

	curItemBinding.AddListener(binding.NewDataListener(func() {
		var item ListItem = ListItem{nil, ""}
		_item, _ := curItemBinding.Get()
		if _item != nil {
			item = _item.(ListItem)
		}

		typeLbl.SetText(fmt.Sprint(item.Type))
		keyLbl.SetText(item.Key)
		valueLbl.SetText(fmt.Sprint(theMap[item.Key]))
	}))

	deleteBtn := widget.NewButtonWithIcon("Delete", theme.ContentRemoveIcon(), func() {
		slog.Info("clicked Delete button")
		deleteSelected()
	})
	deleteBtn.Importance = widget.DangerImportance

	return container.NewStack(container.NewVBox(
		container.NewBorder(nil, nil, widget.NewLabel("Type: "), nil, typeLbl),
		container.NewBorder(nil, nil, widget.NewLabel("Key: "), nil, keyLbl),
		container.NewBorder(nil, nil, widget.NewLabel("Value: "), nil, valueLbl),
		widget.NewSeparator(),
		container.NewHBox(
			widget.NewButtonWithIcon("Edit", theme.DocumentCreateIcon(), func() {
				slog.Info("clicked Edit button")
			}),
			deleteBtn,
		),
	))
}
