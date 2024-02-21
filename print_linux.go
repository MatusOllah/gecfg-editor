package main

import (
	"fmt"
	"image/color"
	"log/slog"
	"os/exec"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func print(a fyne.App, w fyne.Window) {
	slog.Info("printing")

	dlg := dialog.NewCustomWithoutButtons("Printing...", container.NewStack(canvas.NewRectangle(color.Transparent), widget.NewProgressBarInfinite()), w)
	dlg.Show()

	var lines []string
	lines = append(lines, "AUTO-GENERATED by gecfg-editor version "+a.Metadata().Version)
	lines = append(lines, "<github.com/MatusOllah/gecfg-editor>")
	lines = append(lines, "")
	lines = append(lines, "File name: "+openFileName)
	lines = append(lines, "")
	var keys []string
	for k, _ := range theMap {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		lines = append(lines, fmt.Sprintf("%s (%T) = %v", k, theMap[k], theMap[k]))
	}

	if err := printOneDocument(lines); err != nil {
		slog.Error(err.Error())
		dlg.Hide()
		dialog.NewError(err, w).Show()
		return
	}

	dlg.Hide()
}

func printOneDocument(lines []string) error {
	cmd := exec.Command("lp")
	slog.Info("executing", "path", cmd.Path, "args", cmd.Args)
	if err := cmd.Start(); err != nil {
		return err
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	for _, line := range lines {
		fmt.Fprintf(stdin, line+"\r\n")
	}

	return stdin.Close()
}
