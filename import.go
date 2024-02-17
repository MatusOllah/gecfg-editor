package main

import (
	"encoding/json"
	"log/slog"
	"maps"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

func importMap(m map[string]interface{}) {
	slog.Info("importing", "m", m)
	maps.Copy(theMap, m)
	reloadListItems()
}

func importJSON(w fyne.Window) {
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

		slog.Info("importing JSON", "path", path)

		content, err := os.ReadFile(path)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		var m map[string]interface{}
		err = json.Unmarshal(content, &m)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		importMap(m)

	}, w)
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
	dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
	dlg.Show()
}

func importYAML(w fyne.Window) {
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

		slog.Info("importing YAML", "path", path)

		content, err := os.ReadFile(path)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		var m map[string]interface{}
		err = yaml.Unmarshal(content, &m)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		importMap(m)

	}, w)
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".yaml", ".yml"}))
	dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
	dlg.Show()
}

func importTOML(w fyne.Window) {
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

		slog.Info("importing TOML", "path", path)

		content, err := os.ReadFile(path)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		var m map[string]interface{}
		err = toml.Unmarshal(content, &m)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		importMap(m)

	}, w)
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".toml"}))
	dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
	dlg.Show()
}
