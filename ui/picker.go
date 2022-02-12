package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lusingander/colorpicker"
)

func SetupColorPicker(options *UiOptions) *fyne.Container {
	picker := colorpicker.New(200, colorpicker.StyleHue)
	picker.SetOnChanged(func(c color.Color) {
		options.State.BrushColor = c
		options.Swatches[options.State.SwatchSelected].SetColor(c)
	})
	return container.NewVBox(picker)
}
