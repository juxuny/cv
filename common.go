package cv

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func log(x ...interface{}) {
	fmt.Println(x...)
}

type IntSet map[int]bool

func NewIntSet() (s IntSet) {
	s = make(map[int]bool)
	return
}

func (this IntSet) Each(f func(i int)) {
	for k, _ := range this {
		f(k)
	}
}

func (this IntSet) Add(i int) {
	this[i] = true
}

func (this IntSet) Remove(i int) {
	delete(this, i)
}

func (this IntSet) Has(i int) bool {
	_, b := this[i]
	return b
}

func Sum(x ...DataType) (v DataType) {
	v = 0
	for i := 0; i < len(x); i++ {
		v += x[i]
	}
	return
}

func Average(x ...DataType) (v DataType) {
	v = Sum(x...) / DataType(len(x))
	return
}

func Max(x ...DataType) (v DataType) {
	v = x[0]
	for i := 1; i < len(x); i++ {
		if x[i] > v {
			v = x[i]
		}
	}
	return
}

func Min(x ...DataType) (v DataType) {
	v = x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < v {
			v = x[i]
		}
	}
	return
}

func LogisticFunc(x DataType) DataType {
	return DataType(1 / (1 + math.Exp(float64(-x))))
}

func ConvertArrayToMatrix(w, h int, a []DataType) (m Matrix) {
	m = Matrix{W: w, H: h, Data: a}
	return
}

//the function approximates the CDF(Cumulative distribution function[edit])
//refer: https://en.wikipedia.org/wiki/Normal_distribution
func StandardNormalCDF(x float64) (r float64) {
	sum := x
	value := x
	for i := 1; i <= 1000; i++ {
		value = (value * x * x / (2*float64(i) + 1))
		sum += value
	}
	r = 0.5 + (sum/math.Sqrt(2*math.Pi))*math.Exp(-(x*x)/2)
	return
}

//refer: https://en.wikipedia.org/wiki/Gaussian_function
func StandardNormalPDF(mu, sigma, x float64) (r float64) {
	r = GaussianFunction(mu, sigma, x)
	return
}

func RandValue() (r DataType) {
	return DataType(GaussianFunction(0, 1, rand.Float64()))
}

func RandArray(n int) (r []DataType) {
	r = make([]DataType, n)
	for i := 0; i < n; i++ {
		r[i] = DataType(rand.Float64())
	}
	return
}

func RandGaussianDistributionArray(n int) (r []DataType) {
	r = make([]DataType, n)
	for i := 0; i < n; i++ {
		r[i] = DataType(GaussianFunction(0, 1, rand.Float64()))
	}
	return
}

func LoadImage(fileName string) (in IOLayer, e error) {
	f, e := os.Open(fileName)
	if e != nil {
		return
	}
	defer f.Close()
	img, _, e := image.Decode(f)
	if e != nil {
		return
	}
	b := img.Bounds()
	w := b.Max.X
	h := b.Max.Y
	in = NewIOLayer(3, w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c := img.At(i, j)
			r, g, b, _ := c.RGBA()
			in.Set(0, i, j, DataType(convert16To8(r)))
			in.Set(1, i, j, DataType(convert16To8(g)))
			in.Set(2, i, j, DataType(convert16To8(b)))
		}
	}
	return
}

func Standardization(in IOLayer) (ret IOLayer) {
	ret = in.Clone()
	in.EachUnit(func(deep int, unit Matrix) {
		unit.Each(func(x, y int, v DataType) {
			ret.Set(deep, x, y, v/255.0)
		})
	})
	return
}

func CreateResultArray(trueDigital int) (r Array) {
	d := make([]DataType, 10)
	d[trueDigital] = 1
	r = NewArray(d...)
	return
}
