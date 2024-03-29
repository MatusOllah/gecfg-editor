package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

func exportJSON(w fyne.Window) {
	var indent bool

	dialog.NewForm("Export JSON", "OK", "Cancel", []*widget.FormItem{widget.NewFormItem("Indent?", widget.NewCheck("", func(b bool) {
		indent = b
	}))}, func(b bool) {
		if !b {
			return
		}

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

			slog.Info("exporting JSON", "path", path, "indent", indent)

			var b []byte
			if indent {
				b, err = json.MarshalIndent(theMap, "", "\t")
				if err != nil {
					slog.Error(err.Error())
					dialog.NewError(err, w).Show()
					return
				}
			} else {
				b, err = json.Marshal(theMap)
				if err != nil {
					slog.Error(err.Error())
					dialog.NewError(err, w).Show()
					return
				}
			}

			if err := os.WriteFile(path, b, 0666); err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}
		}, w)
		dlg.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
		dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
		dlg.Show()

	}, w).Show()
}

func exportYAML(w fyne.Window) {
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

		slog.Info("exporting YAML", "path", path)

		b, err := yaml.Marshal(&theMap)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		if err := os.WriteFile(path, b, 0666); err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}
	}, w)
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".yaml", ".yml"}))
	dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
	dlg.Show()
}

func exportTOML(w fyne.Window) {
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

		slog.Info("exporting TOML", "path", path)

		var bf bytes.Buffer
		err = toml.NewEncoder(&bf).Encode(&theMap)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		if err := os.WriteFile(path, bf.Bytes(), 0666); err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}
	}, w)
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".toml"}))
	dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
	dlg.Show()
}

func exportGo(w fyne.Window) {
	tg := widget.NewTextGrid()

	pkgNameEntry := widget.NewEntry()
	pkgNameEntry.SetText("main")

	varNameEntry := widget.NewEntry()
	varNameEntry.SetText("m")

	var convert bool

	var indent bool
	indentCheck := widget.NewCheck("", func(b bool) {
		indent = b
		tg.SetText(string(mapToGoString(theMap, varNameEntry.Text, pkgNameEntry.Text, indent, convert, true)))
	})
	indentCheck.SetChecked(true)

	convertCheck := widget.NewCheck("", func(b bool) {
		convert = b
		tg.SetText(string(mapToGoString(theMap, varNameEntry.Text, pkgNameEntry.Text, indent, convert, true)))
	})
	convertCheck.SetChecked(false)

	pkgNameEntry.OnChanged = func(_ string) {
		tg.SetText(string(mapToGoString(theMap, varNameEntry.Text, pkgNameEntry.Text, indent, convert, true)))
	}

	varNameEntry.OnChanged = func(_ string) {
		tg.SetText(string(mapToGoString(theMap, varNameEntry.Text, pkgNameEntry.Text, indent, convert, true)))
	}

	tg.SetText(string(mapToGoString(theMap, varNameEntry.Text, pkgNameEntry.Text, indent, convert, true)))

	content := container.NewVBox(container.NewVBox(widget.NewLabelWithStyle("Preview", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), tg), widget.NewSeparator(), widget.NewForm(
		widget.NewFormItem("Package Name", pkgNameEntry),
		widget.NewFormItem("Variable Name", varNameEntry),
		widget.NewFormItem("Indent?", indentCheck),
		widget.NewFormItem("Convert values?", convertCheck),
	))

	dialog.NewCustomConfirm("Export Go Source File", "OK", "Cancel", content, func(b bool) {
		if !b {
			return
		}

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

			slog.Info("exporting Go Source File", "path", path)

			b := mapToGoString(theMap, varNameEntry.Text, pkgNameEntry.Text, indent, convert, false)

			if err := os.WriteFile(path, b, 0666); err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}
		}, w)
		dlg.SetFilter(storage.NewExtensionFileFilter([]string{".go"}))
		dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
		dlg.Show()
	}, w).Show()
}

func mapToGoString(m map[string]interface{}, varName string, pkgName string, indent bool, convert bool, preview bool) []byte {
	var bf bytes.Buffer

	fmt.Fprintln(&bf, "// AUTO-GENERATED by gecfg-editor; DO NOT EDIT")
	fmt.Fprintln(&bf, "")
	fmt.Fprintln(&bf, "package "+pkgName)
	fmt.Fprintln(&bf, "")
	fmt.Fprintln(&bf, "var "+varName+" map[string]interface{} = map[string]interface{}{")

	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		if indent {
			fmt.Fprint(&bf, "\t")
		}
		fmt.Fprint(&bf, "\""+k+"\": ")
		if convert {
			fmt.Fprintf(&bf, "%T(%#v)", m[k], m[k])
		} else {
			fmt.Fprintf(&bf, "%#v", m[k])
		}
		fmt.Fprintln(&bf, ",")

		if preview {
			if indent {
				fmt.Fprint(&bf, "\t")
			}
			fmt.Fprintln(&bf, "...")
			break
		}
	}

	fmt.Fprintln(&bf, "}")

	return bf.Bytes()
}
