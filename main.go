package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/MatusOllah/slogcolor"
)

func main() {
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	slog.Info("Initializing")
	beforeInit := time.Now()

	a := app.New()

	a.Lifecycle().SetOnStarted(func() {
		slog.Info("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		slog.Info("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		slog.Info("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		slog.Info("Lifecycle: Exited Foreground")
	})

	slog.Info(fmt.Sprintf("gecfg-editor version %s", a.Metadata().Version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))

	reloadListItems()

	w := a.NewWindow(openFileName + " - " + a.Metadata().Name)
	w.SetMaster()
	w.Resize(fyne.NewSize(1280, 720))
	w.SetMainMenu(makeMainMenu(a, w))

	w.SetContent(makeUI(a, w))

	slog.Info(fmt.Sprintf("Initialization took %s", time.Since(beforeInit)))

	w.ShowAndRun()

	slog.Info("exiting")
	runtime.GC()
	os.Exit(0)
}
