package ui

import (
	"fyne.io/fyne/v2"
	"github.com/Shulammite-Aso/pixelArt/apptype"
	"github.com/Shulammite-Aso/pixelArt/swatch"
)

type AppInit struct {
	PixelArtWindow fyne.Window
	State          *apptype.State
	Swatches       []*swatch.Swatch
}
