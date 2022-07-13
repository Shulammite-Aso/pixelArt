package ui

import (
	"fyne.io/fyne/v2"
	"github.com/Shulammite-Aso/pixelArt/apptype"
	pxcanvas "github.com/Shulammite-Aso/pixelArt/pxcanva"
	"github.com/Shulammite-Aso/pixelArt/swatch"
)

type AppInit struct {
	PixelArtCanvas *pxcanvas.PxCanvas
	PixelArtWindow fyne.Window
	State          *apptype.State
	Swatches       []*swatch.Swatch
}
