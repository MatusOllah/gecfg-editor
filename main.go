package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/MatusOllah/gecfg-editor/internal/config"
	"github.com/MatusOllah/slogcolor"
)

func getLogLevel(s string) slog.Leveler {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func main() {
	path := flag.String("open-file", "", "Open a file")
	logLevel := flag.String("log-level", "info", "Log level (\"debug\", \"info\", \"warn\", \"error\")")
	flag.Parse()

	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, &slogcolor.Options{
		Level:       getLogLevel(*logLevel),
		TimeFormat:  time.DateTime,
		SrcFileMode: slogcolor.ShortFile,
	})))
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

	// open file
	if *path != "" {
		slog.Info("opening file", "path", *path)

		cfg, err := config.Open(*path)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		defer cfg.Close()

		openFileName = filepath.Base(*path)
		openFilePath = *path
		theMap = cfg.Data()
		reloadListItems()
		updateWindowTitle(a, w)
	}

	w.ShowAndRun()

	slog.Info("exiting")
	runtime.GC()
	os.Exit(0)
}
