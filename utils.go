package main

import (
	"cmp"
	"errors"
	"log/slog"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"github.com/MatusOllah/gecfg-editor/internal/config"
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
	updateDetails()
}

// windowSizeToDialog scales the window size to a suitable dialog size.
func windowSizeToDialog(s fyne.Size) fyne.Size {
	return fyne.NewSize(s.Width*0.8, s.Height*0.8)
}

func updateWindowTitle(a fyne.App, w fyne.Window) {
	w.SetTitle(openFileName + " - " + a.Metadata().Name)
}

func newFile(a fyne.App, w fyne.Window) {
	theMap = map[string]interface{}{} // wipe map
	reloadListItems()
	openFileName = "Untitled"
	updateWindowTitle(a, w)
}

func openFile(a fyne.App, w fyne.Window) {
	dlg := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		if uc == nil {
			return
		}

		path := uc.URI().Path()

		slog.Info("opening file", "path", path)

		cfg, err := config.Open(path)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}
		defer cfg.Close()

		openFileName = filepath.Base(path)
		openFilePath = path
		theMap = cfg.Data()
		reloadListItems()
		updateWindowTitle(a, w)
	}, w)
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".gecfg"}))
	dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
	dlg.Show()
}

func saveFile(a fyne.App, w fyne.Window) {
	if openFilePath == "" {
		saveFileAs(a, w)
		return
	}

	path := openFilePath

	slog.Info("saving file", "path", path)

	cfg, err := config.New(path)
	if err != nil {
		slog.Error(err.Error())
		dialog.NewError(err, w).Show()
		return
	}

	cfg.SetData(theMap)
	openFileName = filepath.Base(path)
	openFilePath = path
	updateWindowTitle(a, w)

	if err := cfg.Flush(); err != nil {
		slog.Error(err.Error())
		dialog.NewError(err, w).Show()
		return
	}
	if err := cfg.Close(); err != nil {
		slog.Error(err.Error())
		dialog.NewError(err, w).Show()
		return
	}
}

func saveFileAs(a fyne.App, w fyne.Window) {
	dlg := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		if uc == nil {
			return
		}

		path := uc.URI().Path()

		slog.Info("saving file", "path", path)

		cfg, err := config.Create(path)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		cfg.SetData(theMap)
		openFileName = filepath.Base(path)
		openFilePath = path
		updateWindowTitle(a, w)

		if err := cfg.Flush(); err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}
		if err := cfg.Close(); err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}
	}, w)
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".gecfg"}))
	dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
	dlg.Show()
}
