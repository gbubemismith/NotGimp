package ui

import (
	"fyne.io/fyne/v2"
	"github.com/gbubemismith/NotGimp/options"
	"github.com/gbubemismith/NotGimp/swatch"
)

type UiOptions struct {
	Window   fyne.Window
	State    *options.State
	Swatches []*swatch.Swatch
}
