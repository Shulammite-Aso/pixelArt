package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"github.com/Shulammite-Aso/pixelArt/apptype"
	"github.com/Shulammite-Aso/pixelArt/swatch"
	"github.com/Shulammite-Aso/pixelArt/ui"
)

func main() {
	pixelArtApp := app.New()
	pixelArtWindow := pixelArtApp.NewWindow("pixelArt")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixelArtWindow: pixelArtWindow,
		State:          &state,
		Swatches:       make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	// ShowAndRun function is part of the fyne toolkit and will display our application and run it
	appInit.PixelArtWindow.ShowAndRun()
}
