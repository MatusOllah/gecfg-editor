package assets

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed icons/number.svg
var numberSVG []byte
var NumberSVG *fyne.StaticResource = &fyne.StaticResource{
	StaticName:    "number.svg",
	StaticContent: numberSVG,
}

//go:embed icons/string.svg
var stringSVG []byte
var StringSVG *fyne.StaticResource = &fyne.StaticResource{
	StaticName:    "string.svg",
	StaticContent: stringSVG,
}

//go:embed icons/binary.svg
var binarySVG []byte
var BinarySVG *fyne.StaticResource = &fyne.StaticResource{
	StaticName:    "binary.svg",
	StaticContent: binarySVG,
}
