package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/gbubemismith/NotGimp/swatch"
)

func BuildSwatches(options *UiOptions) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(options.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		s := swatch.NewSwatch(options.State, initialColor, i, func(s *swatch.Swatch) {
			for j := 0; j < len(options.Swatches); j++ {
				options.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			options.State.SwatchSelected = s.SwatchIndex
			options.State.BrushColor = s.Color
		})
		if i == 0 {
			s.Selected = true
			options.State.SwatchSelected = 0
			s.Refresh()
		}

		options.Swatches = append(options.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}

	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
