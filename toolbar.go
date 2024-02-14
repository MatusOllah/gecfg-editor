package main

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeToolbar(a fyne.App, w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(widget.NewToolbar(
		widget.NewToolbarAction(theme.FileIcon(), func() {
			slog.Info("selected toolbar item New")
			//TODO: new
		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			slog.Info("selected toolbar item Open")
			//TODO: open file
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			slog.Info("selected toolbar item Save")
			//TODO: save file
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			slog.Info("selected toolbar item Add")
			newValue(a, w)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			slog.Info("selected toolbar item Delete")
			deleteSelected()
		}),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			slog.Info("selected toolbar item Edit")
			//TODO: edit
		}),
	), widget.NewSeparator())
}
