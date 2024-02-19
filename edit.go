package main

import (
	"fmt"
	"log/slog"
	"reflect"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func editSelected(w fyne.Window) {
	if v, _ := curItemBinding.Get(); v == nil {
		return
	}
	item := NewListItemFromDataItem(curItemBinding)

	slog.Info("editing value", "item", item)

	switch item.Type.Kind() {
	case reflect.Int:
		entry := widget.NewEntry()
		entry.TextStyle = fyne.TextStyle{Monospace: true}
		entry.SetText(fmt.Sprint(theMap[item.Key]))

		dialog.NewForm("Edit int", "OK", "Cancel", []*widget.FormItem{widget.NewFormItem("Value", entry)}, func(b bool) {
			if !b {
				return
			}

			v, err := strconv.Atoi(entry.Text)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}

			theMap[item.Key] = v
			reloadListItems()
			updateDetails()
			slog.Info("edited value", "item", item, "key", item.Key, "value", v)
		}, w).Show()
	}
}
