package main

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func makeMainMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	// File
	file_new := fyne.NewMenuItem("New", func() {
		slog.Info("selected menu item File>New")
		//TODO: new
	})
	file_new.Icon = theme.FileIcon()

	file_open := fyne.NewMenuItem("Open", func() {
		slog.Info("selected menu item File>Open")
		//TODO: open
	})
	file_open.Icon = theme.FolderOpenIcon()

	file_save := fyne.NewMenuItem("Save", func() {
		slog.Info("selected menu item File>Save")
		//TODO: save
	})
	file_save.Icon = theme.DocumentSaveIcon()

	file_saveAs := fyne.NewMenuItem("Save As", func() {
		slog.Info("selected menu item File>Save As")
		//TODO: save as
	})
	file_saveAs.Icon = theme.DocumentSaveIcon()

	file_import := fyne.NewMenuItem("Import", func() {})
	file_import.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("JSON", func() {
			slog.Info("selected menu item File>Import>JSON")
			//TODO: import JSON
		}),
	)

	file_export := fyne.NewMenuItem("Export", func() {})
	file_export.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("JSON", func() {
			slog.Info("selected menu item File>Export>JSON")
			exportJSON(w)
		}),
		fyne.NewMenuItem("Go Source File", func() {
			slog.Info("selected menu item File>Export>Go Source File")
			//TODO: export go
		}),
	)

	file := fyne.NewMenu("File",
		file_new,
		fyne.NewMenuItemSeparator(),
		file_open,
		fyne.NewMenuItemSeparator(),
		file_save,
		file_saveAs,
		fyne.NewMenuItemSeparator(),
		file_import,
		fyne.NewMenuItemSeparator(),
		file_export,
	)

	// Edit
	edit_newValue := fyne.NewMenuItem("New Value", func() {
		slog.Info("selected menu item Edit>New Value")
		newValue(a, w)
	})
	edit_newValue.Icon = theme.ContentAddIcon()

	edit_delete := fyne.NewMenuItem("Delete", func() {
		slog.Info("selected menu item Edit>Delete")
		deleteSelected()
	})
	edit_delete.Icon = theme.ContentRemoveIcon()

	edit_edit := fyne.NewMenuItem("Edit", func() {
		slog.Info("selected menu item Edit>Edit")
		//TODO: edit
	})
	edit_edit.Icon = theme.DocumentCreateIcon()

	edit := fyne.NewMenu("Edit",
		edit_newValue,
		edit_delete,
		edit_edit,
	)

	// Help
	help_about := fyne.NewMenuItem("About", func() {
		slog.Info("selected menu item Help>About")
		showAboutDialog(a, w)
	})
	help_about.Icon = theme.InfoIcon()

	help := fyne.NewMenu("Help",
		help_about,
	)

	return fyne.NewMainMenu(file, edit, help)
}
