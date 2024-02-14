GO = go
FYNE = $(GO) run fyne.io/fyne/v2/cmd/fyne@latest
WINRES = $(GO) run github.com/tc-hib/go-winres@latest

TARGET=windows
BINARY=./bin/$(TARGET)

all: build

build: clean
build:
	mkdir -p $(BINARY)
	$(GO) get
	$(WINRES) make
	$(FYNE) build --target $(TARGET) -o $(BINARY)

clean:
	rm -rf $(BINARY)
	rm -f rsrc_windows_amd64.syso
	rm -f rsrc_windows_386.syso
