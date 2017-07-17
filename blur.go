package cv

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

func GaussianFunction(mu, sigma float64, x float64) (p float64) {
	p = 1 / math.Sqrt(2*math.Pi*sigma*sigma) * math.Exp(-((x-mu)*(x-mu))/(2*sigma*sigma))
	return
}

func Distance(x1, y1, x2, y2 int) (d float64) {
	d = math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
	return
}

func abs(x int) (r int) {
	if x < 0 {
		r = -x
	} else {
		r = x
	}
	return
}

func convert16To8(x uint32) uint8 {
	return uint8(float64(x) / float64(1<<16-1) * float64(1<<8-1))
}

//func GaussianBlur(src image.Image, r int) (dst image.Image) {
//	bounds := src.Bounds()
//	w := bounds.Max.X
//	h := bounds.Max.Y
//	weight := make([][]float64, r*2+1)
//	for i := 0; i < len(weight); i++ {
//		weight[i] = make([]float64, r*2+1)
//	}
//	var sum float64 = 0
//	for x := 0; x < len(weight); x++ {
//		for y := 0; y < len(weight[x]); y++ {
//			weight[x][y] = GaussianFunction(0, float64(r)/3, Distance(x, y, r, r))
//			//			weight[x][y] = 1 / math.Pow(2*float64(r)+1, 2)
//			sum += weight[x][y]
//		}
//	}
//	for x := 0; x < len(weight); x++ {
//		for y := 0; y < len(weight[x]); y++ {
//			weight[x][y] /= sum
//		}
//	}
//	fmt.Println("weight matrix: ", weight)
//	tmp := image.NewNRGBA(bounds)
//	for x := 0; x <= w; x += 1 {
//		for y := 0; y <= h; y += 1 {
//			left := x - r
//			top := y - r
//			right := x + r
//			bottom := y + r
//			var R, G, B, A float64
//			for i := left; i <= right; i++ {
//				for j := top; j <= bottom; j++ {
//					tx := i
//					ty := j
//					if tx < 0 {
//						tx = x + x - tx
//					}
//					if ty < 0 {
//						ty = y + y - ty
//					}
//					if tx > w {
//						tx = w - (tx - w)
//					}
//					if ty > h {
//						ty = h - (ty - h)
//					}
//					c := src.At(tx, ty)
//					tR, tG, tB, tA := c.RGBA()

//					R += float64(tR) * weight[abs(tx-x)+r][abs(ty-y)+r]
//					G += float64(tG) * weight[abs(tx-x)+r][abs(ty-y)+r]
//					B += float64(tB) * weight[abs(tx-x)+r][abs(ty-y)+r]
//					A += float64(tA) * weight[abs(tx-x)+r][abs(ty-y)+r]
//				}
//			}
//			tmp.Set(x, y, color.NRGBA{R: convert16To8(uint32(R)), G: convert16To8(uint32(G)), B: convert16To8(uint32(B)), A: convert16To8(uint32(A))})
//		}
//		fmt.Println("process: ", x, "/", w)
//	}
//	dst = tmp.SubImage(src.Bounds())
//	return
//}

//func saveJpg(img image.Image, fileName string) (e error) {
//	f, e := os.Create(fileName)
//	if e != nil {
//		return
//	}
//	defer f.Close()
//	e = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
//	return
//}

//func BoxBlur(src image.Image, r int) (dst image.Image) {
//	bounds := src.Bounds()
//	w := bounds.Max.X
//	h := bounds.Max.Y
//	tmp := image.NewNRGBA(bounds)
//	for x := 0; x <= w; x += 1 {
//		for y := 0; y <= h; y += 1 {
//			left := x - r
//			top := y - r
//			right := x + r
//			bottom := y + r
//			var R, G, B, A float64
//			for i := left; i <= right; i++ {
//				for j := top; j <= bottom; j++ {
//					tx := i
//					ty := j
//					if tx < 0 {
//						tx = x + x - tx
//					}
//					if ty < 0 {
//						ty = y + y - ty
//					}
//					if tx > w {
//						tx = w - (tx - w)
//					}
//					if ty > h {
//						ty = h - (ty - h)
//					}
//					c := src.At(tx, ty)
//					tR, tG, tB, tA := c.RGBA()

//					R += float64(tR)
//					G += float64(tG)
//					B += float64(tB)
//					A += float64(tA)
//				}
//			}
//			R /= math.Pow(2*float64(r)+1, 2)
//			G /= math.Pow(2*float64(r)+1, 2)
//			B /= math.Pow(2*float64(r)+1, 2)
//			A /= math.Pow(2*float64(r)+1, 2)
//			tmp.Set(x, y, color.NRGBA{R: convert16To8(uint32(R)), G: convert16To8(uint32(G)), B: convert16To8(uint32(B)), A: convert16To8(uint32(A))})
//		}
//		fmt.Println("process: ", x, "/", w)
//	}
//	dst = tmp.SubImage(src.Bounds())
//	return
//}

func BoxBlur(src image.Image, r int) (dst image.Image) {
	bounds := src.Bounds()
	w := bounds.Max.X
	h := bounds.Max.Y
	tmp := image.NewRGBA(bounds)
	//	count := 2*r + 1
	pR := make([][]uint64, w+1)
	pG := make([][]uint64, w+1)
	pB := make([][]uint64, w+1)
	pA := make([][]uint64, w+1)
	for x := 0; x <= w; x++ {
		pR[x] = make([]uint64, h+1)
		pG[x] = make([]uint64, h+1)
		pB[x] = make([]uint64, h+1)
		pA[x] = make([]uint64, h+1)
	}
	for x := 0; x <= w; x++ {
		for y := 0; y <= h; y++ {
			tR, tG, tB, tA := src.At(x, y).RGBA()
			var (
				leftR, topR, leftTopR uint64
				leftG, topG, leftTopG uint64
				leftB, topB, leftTopB uint64
				leftA, topA, leftTopA uint64
			)
			if x-1 < 0 {
				leftR = 0
				leftG = 0
				leftB = 0
				leftA = 0
			} else {
				leftR = pR[x-1][y]
				leftG = pG[x-1][y]
				leftB = pB[x-1][y]
				leftA = pA[x-1][y]
			}
			if y-1 < 0 {
				topR = 0
				topG = 0
				topB = 0
				topA = 0
			} else {
				topR = pR[x][y-1]
				topG = pG[x][y-1]
				topB = pB[x][y-1]
				topA = pA[x][y-1]
			}
			if x-1 < 0 || y-1 < 0 {
				leftTopR = 0
				leftTopG = 0
				leftTopB = 0
				leftTopA = 0
			} else {
				leftTopR = pR[x-1][y-1]
				leftTopG = pG[x-1][y-1]
				leftTopB = pB[x-1][y-1]
				leftTopA = pA[x-1][y-1]
			}
			pR[x][y] = leftR + topR - leftTopR + uint64(convert16To8(uint32(tR)))
			pG[x][y] = leftG + topG - leftTopG + uint64(convert16To8(uint32(tG)))
			pB[x][y] = leftB + topB - leftTopB + uint64(convert16To8(uint32(tB)))
			pA[x][y] = leftA + topA - leftTopA + uint64(convert16To8(uint32(tA)))
		}
	}
	for x := 0; x <= w; x++ {
		for y := 0; y <= h; y++ {
			left := x - r
			top := y - r
			right := x + r
			bottom := y + r
			if left < 0 {
				left = 0
			}
			if right > w {
				right = w
			}
			if top < 0 {
				top = 0
			}
			if bottom > h {
				bottom = h
			}
			//			fmt.Println(x, y, left, top, right, bottom)
			var marginTopR, marginLeftR, marginCommonR uint64
			var marginTopG, marginLeftG, marginCommonG uint64
			var marginTopB, marginLeftB, marginCommonB uint64
			var marginTopA, marginLeftA, marginCommonA uint64
			if left-1 < 0 {
				marginLeftR = 0
				marginLeftG = 0
				marginLeftB = 0
				marginLeftA = 0
			} else {
				marginLeftR = pR[left-1][bottom]
				marginLeftG = pG[left-1][bottom]
				marginLeftB = pB[left-1][bottom]
				marginLeftA = pA[left-1][bottom]
			}
			if top-1 < 0 {
				marginTopR = 0
				marginTopG = 0
				marginTopB = 0
				marginTopA = 0
			} else {
				marginTopR = pR[right][top-1]
				marginTopG = pG[right][top-1]
				marginTopB = pB[right][top-1]
				marginTopA = pA[right][top-1]
			}
			if left-1 < 0 || top-1 < 0 {
				marginCommonR = 0
				marginCommonG = 0
				marginCommonB = 0
				marginCommonA = 0
			} else {
				marginCommonR = pR[left-1][top-1]
				marginCommonG = pG[left-1][top-1]
				marginCommonB = pB[left-1][top-1]
				marginCommonA = pA[left-1][top-1]
			}

			sR := pR[right][bottom] - marginLeftR - marginTopR + marginCommonR
			sG := pG[right][bottom] - marginLeftG - marginTopG + marginCommonG
			sB := pB[right][bottom] - marginLeftB - marginTopB + marginCommonB
			sA := pA[right][bottom] - marginLeftA - marginTopA + marginCommonA
			count := (right - left + 1) * (bottom - top + 1)
			tmp.Set(x, y, color.RGBA{
				R: uint8(float64(sR) / float64(count)),
				G: uint8(float64(sG) / float64(count)),
				B: uint8(float64(sB) / float64(count)),
				A: uint8(float64(sA) / float64(count)),
			})
		}
	}
	dst = tmp.SubImage(bounds)
	return
}

func GaussianBlur(src image.Image, r int) (dst image.Image) {
	dst = src
	for i := 0; i < 3; i++ {
		dst = BoxBlur(dst, r)
	}
	return
}

//func init() {
//	flag.StringVar(&inputFile, "i", "1.jpg", "input file")
//	flag.StringVar(&outputFile, "o", "out.jpg", "output file")
//	flag.IntVar(&r, "r", 5, "radius")
//}

//func main() {
//	flag.Parse()
//	f, e := os.Open(inputFile)
//	if e != nil {
//		panic(e)
//	}
//	defer f.Close()
//	img, e := jpeg.Decode(f)
//	if e != nil {
//		panic(e)
//	}
//	var dst image.Image = img
//	for i := 0; i < 1; i++ {
//		dst = GaussianBlur2(dst, r)
//	}
//	saveJpg(dst, outputFile)
//}
