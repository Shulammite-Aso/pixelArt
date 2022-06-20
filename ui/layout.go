package ui

// Will setup all of our different ui elements
func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)

	app.PixelArtWindow.SetContent(swatchesContainer)
}
