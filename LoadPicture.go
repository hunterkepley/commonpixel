// LoadPicture loads a picture and prepares it for use in Pixel

package commonPixel

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

// LoadPicture loads an image file (.png for example) and returns a pixel.Picture
func LoadPicture(path string) pixel.Picture {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil
	}
	return pixel.PictureDataFromImage(img)
}
