package swatch

import "fyne.io/fyne/v2/driver/desktop"

func (s *Swatch) MouseDown(env *desktop.MouseEvent) {
	s.clickHandler(s)
	s.Selected = true
	s.Refresh()
}

func (s *Swatch) MouseUp(env *desktop.MouseEvent) {}
