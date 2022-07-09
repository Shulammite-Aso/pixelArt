package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/Shulammite-Aso/pixelArt/swatch"
)

// The container in the fyne.container type are used to contain layouts, so we're goin gto be placing
//swatches into a layout and returning it as a conatiner
func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		s := swatch.NewSwatch(app.State, initialColor, i, func(s *swatch.Swatch) {
			// Unselect all the other swatches
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			app.State.SwatchSelected = s.SwatchIndex
			app.State.BrushColor = s.Color
		})

		// Whenever user opens the program, the 0 index swatch will be the one that is selected
		if i == 0 {
			s.Selected = true
			app.State.SwatchSelected = 0
			s.Refresh()
		}
		app.Swatches = append(app.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}

	// The reason we created the canvasSwatches variable containing canvas objects is to utilise it in
	//the NewGridWrap layout. Layouts such as GridWrap only operate on canvas objects
	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
