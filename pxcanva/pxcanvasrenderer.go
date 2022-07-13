package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	canvasImage  *canvas.Image // The one we would display on screen and is based on px img
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject
}

func (renderer *PxCanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	renderer.canvasCursor = objects
}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingArea
}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	// Would loop through the first four objects which represent the borders
	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}
	objects = append(objects, renderer.canvasImage)
	objects = append(objects, renderer.canvasCursor...)
	return objects
}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) Destroy() {}

func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	// We split our layout function into three separate functions
	// One for laying out image and the other for laying out border

	renderer.LayoutCanvas(size)
	renderer.LayoutBorder(size)
}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) Refresh() {
	if renderer.pxCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.pxCanvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.pxCanvas.reloadImage = false
	}
	//.Layout(renderer.pxCanvas.Size())
	renderer.Layout(renderer.pxCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}

// For laying out image
func (renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := renderer.pxCanvas.PxCols
	imgPxHeight := renderer.pxCanvas.PxRows
	pxSize := renderer.pxCanvas.PxSize

	renderer.canvasImage.Move(fyne.NewPos(renderer.pxCanvas.CanvasOffset.X, renderer.pxCanvas.CanvasOffset.Y))
	// Change our pixel size so they're not too small
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}

// For Laying out border
func (renderer *PxCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.pxCanvas.CanvasOffset
	imgHeight := renderer.canvasImage.Size().Height
	imgWidth := renderer.canvasImage.Size().Width

	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)

	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)

	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)
}
