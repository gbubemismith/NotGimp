package ui

import "fyne.io/fyne/v2/container"

func Setup(options *UiOptions) {
	swatchesContainer := BuildSwatches(options)
	colorPicker := SetupColorPicker(options)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker)

	options.Window.SetContent(appLayout)
}
