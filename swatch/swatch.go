package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/Shulammite-Aso/pixelArt/apptype"
)

// Create widget structure
type Swatch struct {
	widget.BaseWidget // embed BaseWidget struct from fyne
	Selected          bool
	Color             color.Color
	SwatchIndex       int
	clickHandler      func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}

	// All this does is supply all the state infor to the fyne toolkit so we don't have to do it manually
	swatch.ExtendBaseWidget(swatch)
	return swatch
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}

	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}
