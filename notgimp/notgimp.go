package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"github.com/gbubemismith/NotGimp/options"
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

	appInit := ui.UiOptions{
		Window:   gWindow,
		State:    &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.Window.ShowAndRun()
}
