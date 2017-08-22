package cv

import (
	"fmt"
)

type DataType float64

type Matrix struct {
	W, H int
	Data []DataType
}

func NewMatrix(w, h int) (m Matrix) {
	m.W = w
	m.H = h
	m.Data = make([]DataType, w*h)
	return
}

func CreateMatrix(w, h int, x ...DataType) (m Matrix) {
	if len(x) < w*h {
		panic("not enough data")
	}
	m = ConvertArrayToMatrix(w, h, x)
	return
}

func CreateGaussianMatrix(w, h int) (m Matrix) {
	n := w * h
	a := RandGaussianDistributionArray(n)
	m = CreateMatrix(w, h, a...)
	return
}

func (this *Matrix) Width() (w int) {
	return this.W
}

func (this *Matrix) Height() (h int) {
	return this.H
}

func (this *Matrix) Get(c, r int) (f DataType) {
	return this.Data[r*this.W+c]
}

func (this *Matrix) Set(c, r int, v DataType) {
	this.Data[r*this.Width()+c] = v
}

func (this *Matrix) Add(c, r int, v DataType) {
	this.Data[r*this.Width()+c] += v
}

func (this *Matrix) AddAll(v DataType) {
	for i := 0; i < len(this.Data); i++ {
		this.Data[i] += v
	}
}

func (this *Matrix) DotProduct(f Matrix) (r DataType) {
	r = 0
	for i, v := range this.Data {
		r += v * f.Data[i]
	}
	return
}

func (this *Matrix) Crop(l, t, r, b int) (m Matrix) {
	m = NewMatrix(r-l+1, b-t+1)
	for x := l; x <= r; x++ {
		for y := t; y <= b; y++ {
			m.Set(x-l, y-t, this.Get(x, y))
		}
	}
	return
}

func (this Matrix) Clone() (m Matrix) {
	m = NewMatrix(this.W, this.H)
	m.Data = make([]DataType, len(this.Data))
	for i := 0; i < len(this.Data); i++ {
		m.Data[i] = this.Data[i]
	}
	return
}

func (this *Matrix) Each(f func(x, y int, v DataType)) {
	for i := 0; i < len(this.Data); i++ {
		f(i%this.W, i/this.W, this.Data[i])
	}
}

func (this Matrix) Max() DataType {
	return Max(this.Data...)
}

func (this Matrix) PaddingWith(width int, v DataType) (out Matrix) {
	if width <= 0 {
		return this
	}
	out = NewMatrix(this.W+width*2, this.H+width*2)
	out.Each(func(x, y int, v DataType) {
		if x >= width && x < out.W-width && y >= width && y < out.H-width {
			out.Set(x, y, this.Get(x-width, y-width))
		} else {
			out.Set(x, y, v)
		}
	})
	return
}

func (this Matrix) String() (s string) {
	s += fmt.Sprintf("%v x %v\n", this.Width(), this.Height())
	for j := 0; j < this.Height(); j++ {
		for i := 0; i < this.Width(); i++ {
			s += fmt.Sprintf("%v", this.Get(i, j)) + " "
		}
		s += "\n"
	}
	return
}

func (this *Matrix) Apply(f func(x, y int, v DataType) DataType) {
	this.Each(func(x, y int, v DataType) {
		this.Set(x, y, f(x, y, v))
	})
}
