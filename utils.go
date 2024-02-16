package main

import (
	"cmp"
	"errors"
	"log/slog"
	"reflect"
	"slices"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func reloadListItems() {
	slog.Info("reloading list items")
	items.Set(nil)
	var _items []ListItem
	for k, v := range theMap {
		_items = append(_items, ListItem{reflect.TypeOf(v), k})
	}
	slices.SortFunc(_items, func(a ListItem, b ListItem) int {
		return cmp.Compare(a.Key, b.Key)
	})

	for _, i := range _items {
		items.Append(i)
	}
}

func newValue(a fyne.App, w fyne.Window) {
	keyEntry := widget.NewEntry()
	typeEntry := widget.NewSelectEntry([]string{
		"string",
		"bool",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"float32",
		"float64",
	})
	valueEntry := widget.NewEntry()

	form := []*widget.FormItem{
		widget.NewFormItem("Key", keyEntry),
		widget.NewFormItem("Type", typeEntry),
		widget.NewFormItem("Value", valueEntry),
	}

	dialog.NewForm("New Value", "OK", "Cancel", form, func(b bool) {
		if !b {
			return
		}

		switch typeEntry.Text {
		case "string":
			theMap[keyEntry.Text] = valueEntry.Text
		case "bool":
			v, err := strconv.ParseBool(valueEntry.Text)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = v
		case "int":
			v, err := strconv.Atoi(valueEntry.Text)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = v
		case "int8":
			v, err := strconv.ParseInt(valueEntry.Text, 10, 8)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = int8(v)
		case "int16":
			v, err := strconv.ParseInt(valueEntry.Text, 10, 16)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = int16(v)
		case "int32":
			v, err := strconv.ParseInt(valueEntry.Text, 10, 32)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = int32(v)
		case "int64":
			v, err := strconv.ParseInt(valueEntry.Text, 10, 64)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = int64(v)
		case "uint":
			v, err := strconv.ParseUint(valueEntry.Text, 10, 0)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = uint(v)
		case "uint8":
			v, err := strconv.ParseUint(valueEntry.Text, 10, 8)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = uint8(v)
		case "uint16":
			v, err := strconv.ParseUint(valueEntry.Text, 10, 16)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = uint16(v)
		case "uint32":
			v, err := strconv.ParseUint(valueEntry.Text, 10, 32)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = uint32(v)
		case "uint64":
			v, err := strconv.ParseUint(valueEntry.Text, 10, 64)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = uint64(v)
		case "float32":
			v, err := strconv.ParseFloat(valueEntry.Text, 10)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = float32(v)
		case "float64":
			v, err := strconv.ParseFloat(valueEntry.Text, 10)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w)
			}

			theMap[keyEntry.Text] = float64(v)
		default:
			slog.Error("invalid type", "type", typeEntry.Text)
			dialog.NewError(errors.New("invalid type: "+typeEntry.Text), w).Show()
		}
		reloadListItems()
	}, w).Show()
}

func deleteSelected() {
	if v, _ := curItemBinding.Get(); v == nil {
		return
	}
	item := NewListItemFromDataItem(curItemBinding)
	delete(theMap, item.Key)
	l.UnselectAll()
	reloadListItems()
}

// windowSizeToDialog scales the window size to a suitable dialog size.
func windowSizeToDialog(s fyne.Size) fyne.Size {
	return fyne.NewSize(s.Width*0.8, s.Height*0.8)
}
