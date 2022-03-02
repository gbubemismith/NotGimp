package ui

import (
	"fyne.io/fyne/v2"
	"github.com/gbubemismith/NotGimp/options"
	"github.com/gbubemismith/NotGimp/pxcanvas"
	"github.com/gbubemismith/NotGimp/swatch"
)

type UiOptions struct {
	Canvas   *pxcanvas.PxCanvas
	Window   fyne.Window
	State    *options.State
	Swatches []*swatch.Swatch
}
