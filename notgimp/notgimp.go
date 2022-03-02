package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/gbubemismith/NotGimp/options"
	"github.com/gbubemismith/NotGimp/pxcanvas"
	"github.com/gbubemismith/NotGimp/swatch"
	"github.com/gbubemismith/NotGimp/ui"
)

func main() {
	gAapp := app.New()
	gWindow := gAapp.NewWindow("Not Gimp")

	state := options.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	canvasConfig := options.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(750, 750),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}

	canvas := pxcanvas.NewPxCanvas(&state, canvasConfig)

	appInit := ui.UiOptions{
		Canvas:   canvas,
		Window:   gWindow,
		State:    &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.Window.ShowAndRun()
}
