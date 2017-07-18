package cv

import (
	"image"
	"image/color"
	"math"
)

func SobelOperator(src image.Image) (dst image.Image) {
	bounds := src.Bounds()
	Dy := [][]int64{
		[]int64{-1, 0, 1},
		[]int64{-2, 0, 2},
		[]int64{-1, 0, 1},
	}
	Dx := [][]int64{
		[]int64{-1, -2, -1},
		[]int64{0, 0, 0},
		[]int64{-1, -2, -1},
	}
	G := image.NewGray(bounds)
	w := bounds.Max.X
	h := bounds.Max.Y
	for x := 0; x <= w; x++ {
		for y := 0; y <= h; y++ {
			left := x - 1
			right := x + 1
			top := y - 1
			bottom := y + 1
			var Gx int64 = 0
			var Gy int64 = 0
			for i := left; i <= right; i++ {
				for j := top; j <= bottom; j++ {
					dx := i
					dy := j
					if dx < 0 {
						dx = right
					}
					if dx > w {
						dx = left
					}
					if dy < 0 {
						dy = bottom
					}
					if dy > h {
						dy = top
					}
					c := src.At(dx, dy)
					gray, _, _, _ := c.RGBA()
					//					fmt.Println(dx, dy, x, y, gray)
					Gx += int64(Dx[abs(dx-x)+1][abs(dy-y)+1]) * int64(gray)
					Gy += int64(Dy[abs(dx-x)+1][abs(dy-y)]+1) * int64(gray)
				}
			}
			g := math.Sqrt(float64(Gx*Gx + Gy*Gy))
			G.Set(x, y, color.Gray{Y: convert16To8(uint32(g))})
		}
	}
	dst = G.SubImage(bounds)
	return
}
