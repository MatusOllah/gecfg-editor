package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type IconLabel struct {
	widget.BaseWidget
	Icon  *canvas.Image
	Label *widget.Label
}

func NewIconLabel(text string, icon fyne.Resource) *IconLabel {
	img := canvas.NewImageFromResource(icon)
	img.FillMode = canvas.ImageFillOriginal

	il := &IconLabel{
		Icon:  img,
		Label: widget.NewLabel(text),
	}
	il.ExtendBaseWidget(il)

	return il
}

func (il *IconLabel) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, il.Icon, nil, il.Label)
	return widget.NewSimpleRenderer(c)
}
