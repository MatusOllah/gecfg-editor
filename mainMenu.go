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
		newFile(a, w)
	})
	file_new.Icon = theme.FileIcon()

	file_open := fyne.NewMenuItem("Open", func() {
		slog.Info("selected menu item File>Open")
		openFile(a, w)
	})
	file_open.Icon = theme.FolderOpenIcon()

	file_save := fyne.NewMenuItem("Save", func() {
		slog.Info("selected menu item File>Save")
		saveFile(a, w)
	})
	file_save.Icon = theme.DocumentSaveIcon()

	file_saveAs := fyne.NewMenuItem("Save As", func() {
		slog.Info("selected menu item File>Save As")
		saveFileAs(a, w)
	})
	file_saveAs.Icon = theme.DocumentSaveIcon()

	file_import := fyne.NewMenuItem("Import", func() {})
	file_import.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("JSON", func() {
			slog.Info("selected menu item File>Import>JSON")
			importJSON(w)
		}),
		fyne.NewMenuItem("YAML", func() {
			slog.Info("selected menu item File>Import>YAML")
			importYAML(w)
		}),
		fyne.NewMenuItem("TOML", func() {
			slog.Info("selected menu item File>Import>YAML")
			importTOML(w)
		}),
	)

	file_export := fyne.NewMenuItem("Export", func() {})
	file_export.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("JSON", func() {
			slog.Info("selected menu item File>Export>JSON")
			exportJSON(w)
		}),
		fyne.NewMenuItem("YAML", func() {
			slog.Info("selected menu item File>Export>YAML")
			exportYAML(w)
		}),
		fyne.NewMenuItem("TOML", func() {
			slog.Info("selected menu item File>Export>TOML")
			exportTOML(w)
		}),
		fyne.NewMenuItem("Go Source File", func() {
			slog.Info("selected menu item File>Export>Go Source File")
			exportGo(w)
		}),
	)

	file_print := fyne.NewMenuItem("Print", func() {
		slog.Info("selected menu item File>Print")
		print(a, w)
	})
	file_print.Icon = theme.DocumentPrintIcon()

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
		fyne.NewMenuItemSeparator(),
		file_print,
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
		editSelected(w)
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
