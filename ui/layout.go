package ui

import (
	"fyne.io/fyne/v2/container"
)

// Will setup all of our different ui elements
func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	colorpicker := SetupColorPicker(app)
	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorpicker, app.PixelArtCanvas)
	app.PixelArtWindow.SetContent(appLayout)
}
