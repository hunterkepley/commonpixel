// Collisions is a collection of basic collision functions to be used
// for a Pixel game

package commonPixel

import (
	"math"

	"github.com/faiface/pixel"
)

// AABBCollisionCheck returns a bool after checking if two pixel.Rect's for overlap
func AABBCollisionCheck(x pixel.Rect, y pixel.Rect) bool {
	if x.Min.X < y.Max.X &&
		x.Max.X > y.Min.X &&
		x.Min.Y < y.Max.Y &&
		x.Max.Y > y.Min.Y {
		return true
	}
	return false
}

// CalculateDistance returns the distance between two pixel.Vec's
func CalculateDistance(a pixel.Vec, b pixel.Vec) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}

// CircularCollisionCheck returns a bool after checking if two circles overlap
// it checks based off of two radii and the distance between them (can be gathered)
// using the CalculateDistance function or by other means
func CirclularCollisionCheck(r1 float64, r2 float64, d float64) bool {
	if d < r1+r2 {
		return true
	}
	return false
}
