package main

import (
	"log/slog"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MatusOllah/gecfg-editor/assets"
)

func makeList(w fyne.Window) fyne.CanvasObject {
	l = widget.NewListWithData(
		items,
		func() fyne.CanvasObject {
			return NewIconLabel("", theme.ErrorIcon())
		},
		func(i binding.DataItem, co fyne.CanvasObject) {
			item := NewListItemFromDataItem(i)
			co.(*IconLabel).Label.SetText(item.Key)

			switch item.Type.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Uintptr:
				co.(*IconLabel).Icon.Resource = theme.NewThemedResource(assets.NumberSVG)
			case reflect.String:
				co.(*IconLabel).Icon.Resource = theme.NewThemedResource(assets.StringSVG)
			case reflect.Bool:
				co.(*IconLabel).Icon.Resource = theme.NewThemedResource(assets.BinarySVG)
			case reflect.Array, reflect.Slice:
				co.(*IconLabel).Icon.Resource = theme.ListIcon()
			case reflect.Invalid:
				co.(*IconLabel).Icon.Resource = theme.ErrorIcon()
			default:
				co.(*IconLabel).Icon.Resource = theme.QuestionIcon()
			}

			co.(*IconLabel).Icon.Refresh()
		},
	)

	l.OnSelected = func(id widget.ListItemID) {
		theItems, _ := items.Get()
		slog.Info("selected item", "id", id, "item", theItems[id])
		curItemBinding.Set(theItems[id])
	}

	return l
}
