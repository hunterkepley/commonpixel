// Letterbox makes a Pixel window scale properly using
// the letterbox approach (black bars) so you don't have to worry about
// resizing images or making images for different aspect ratios

package commonPixel

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// LetterBox is a function that adds black bars to a pixelgl Window
func LetterBox(win *pixelgl.Window, winWidth float32, winHeight float32) {
	sizeX := 1.
	sizeY := 1.

	if win.Bounds().H()-winHeight > win.Bounds().W()-winWidth {
		sizeX = win.Bounds().W() / winWidth
		sizeY = win.Bounds().W() / winWidth
	} else {
		sizeX = win.Bounds().H() / winHeight
		sizeY = win.Bounds().H() / winHeight
	}

	viewMatrix = pixel.IM.
		Moved(pixel.V(win.Bounds().Center().X, win.Bounds().Center().Y)).
		ScaledXY(pixel.V(win.Bounds().Center().X, win.Bounds().Center().Y), pixel.V(sizeX, sizeY))
}
