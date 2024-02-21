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
	case reflect.String:
		entry := widget.NewEntry()
		entry.TextStyle = fyne.TextStyle{Monospace: true}
		entry.SetText(fmt.Sprint(theMap[item.Key]))

		dialog.NewForm("Edit string", "OK", "Cancel", []*widget.FormItem{widget.NewFormItem("Value", entry)}, func(b bool) {
			if !b {
				return
			}

			v := entry.Text

			theMap[item.Key] = v
			reloadListItems()
			updateDetails()
			slog.Info("edited value", "item", item, "key", item.Key, "value", v)
		}, w).Show()
	case reflect.Bool:
		var v bool
		check := widget.NewCheck("", func(b bool) {
			v = b
		})
		check.SetChecked(theMap[item.Key].(bool))

		dialog.NewForm("Edit bool", "OK", "Cancel", []*widget.FormItem{widget.NewFormItem("Value", check)}, func(b bool) {
			if !b {
				return
			}

			theMap[item.Key] = v
			reloadListItems()
			updateDetails()
			slog.Info("edited value", "item", item, "key", item.Key, "value", v)
		}, w).Show()
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
	case reflect.Int8:
		entry := widget.NewEntry()
		entry.TextStyle = fyne.TextStyle{Monospace: true}
		entry.SetText(fmt.Sprint(theMap[item.Key]))

		dialog.NewForm("Edit int8", "OK", "Cancel", []*widget.FormItem{widget.NewFormItem("Value", entry)}, func(b bool) {
			if !b {
				return
			}

			v, err := strconv.ParseInt(entry.Text, 0, 8)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}

			theMap[item.Key] = int8(v)
			reloadListItems()
			updateDetails()
			slog.Info("edited value", "item", item, "key", item.Key, "value", v)
		}, w).Show()
	case reflect.Int16:
		entry := widget.NewEntry()
		entry.TextStyle = fyne.TextStyle{Monospace: true}
		entry.SetText(fmt.Sprint(theMap[item.Key]))

		dialog.NewForm("Edit int16", "OK", "Cancel", []*widget.FormItem{widget.NewFormItem("Value", entry)}, func(b bool) {
			if !b {
				return
			}

			v, err := strconv.ParseInt(entry.Text, 0, 16)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}

			theMap[item.Key] = int16(v)
			reloadListItems()
			updateDetails()
			slog.Info("edited value", "item", item, "key", item.Key, "value", v)
		}, w).Show()
	case reflect.Int32:
		entry := widget.NewEntry()
		entry.TextStyle = fyne.TextStyle{Monospace: true}
		entry.SetText(fmt.Sprint(theMap[item.Key]))

		dialog.NewForm("Edit int32", "OK", "Cancel", []*widget.FormItem{widget.NewFormItem("Value", entry)}, func(b bool) {
			if !b {
				return
			}

			v, err := strconv.ParseInt(entry.Text, 0, 32)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}

			theMap[item.Key] = int32(v)
			reloadListItems()
			updateDetails()
			slog.Info("edited value", "item", item, "key", item.Key, "value", v)
		}, w).Show()
	case reflect.Int64:
		entry := widget.NewEntry()
		entry.TextStyle = fyne.TextStyle{Monospace: true}
		entry.SetText(fmt.Sprint(theMap[item.Key]))

		dialog.NewForm("Edit int64", "OK", "Cancel", []*widget.FormItem{widget.NewFormItem("Value", entry)}, func(b bool) {
			if !b {
				return
			}

			v, err := strconv.ParseInt(entry.Text, 0, 64)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}

			theMap[item.Key] = int64(v)
			reloadListItems()
			updateDetails()
			slog.Info("edited value", "item", item, "key", item.Key, "value", v)
		}, w).Show()
	}
}
