package main

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	theMap         map[string]interface{} = map[string]interface{}{}
	items          binding.UntypedList    = binding.NewUntypedList()
	curItemBinding binding.Untyped        = binding.NewUntyped()
	l              *widget.List
	openFileName   string = "Untitled"
	openFilePath   string
)
