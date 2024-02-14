package main

import (
	"reflect"

	"fyne.io/fyne/v2/data/binding"
)

type ListItem struct {
	Type reflect.Type
	Key  string
}

func NewListItemFromDataItem(item binding.DataItem) ListItem {
	v, err := item.(binding.Untyped).Get()
	if err != nil {
		panic(err)
	}

	return v.(ListItem)
}
