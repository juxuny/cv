package cv

import (
	"image"
)

func RGB2Gray(img image.Image) (dst image.Image) {
	bounds := img.Bounds()
	tmp := image.NewGray(img.Bounds())
	for x := 0; x <= bounds.Max.X; x++ {
		for y := 0; y <= bounds.Max.Y; y++ {
			c := tmp.ColorModel().Convert(img.At(x, y))
			tmp.Set(x, y, c)
		}
	}
	dst = tmp.SubImage(bounds)
	return
}
