package cv

import (
	"fmt"
)

type IOLayer struct {
	W, H  int
	Units []Matrix
}

func NewIOLayer(deep, w, h int) (l IOLayer) {
	l.W = w
	l.H = h
	l.Units = make([]Matrix, deep)

	for i := 0; i < deep; i++ {
		l.Units[i] = NewMatrix(w, h)
	}
	return
}

func (this IOLayer) String() (s string) {
	s = "\n"
	for d := 0; d < this.Deep(); d++ {
		s += fmt.Sprintf("deep: %v\n", d)
		s += this.Units[d].String()
	}
	return
}

func (this IOLayer) Walk(w, h, stepWidth, padding int, f func(deep, inLeft, inTop int, crop Matrix)) {
	if stepWidth < 1 {
		panic("step width can't smaller than one")
	}
	this.EachUnit(func(deep int, unit Matrix) {
		p := unit.PaddingWith(padding, 0)
		for j := 0; j+h-1 < p.H; j += stepWidth {
			for i := 0; i+w-1 < p.W; i += stepWidth {
				crop := p.Crop(i, j, i+w-1, j+h-1)
				f(deep, i, j, crop)
			}
		}
	})
}

func (this IOLayer) Deep() int {
	return len(this.Units)
}

func (this IOLayer) Width() int {
	return this.W
}

func (this IOLayer) Height() int {
	return this.H
}

func (this IOLayer) Clone() (l IOLayer) {
	l.W = this.W
	l.H = this.H
	l.Units = make([]Matrix, len(this.Units))
	for i := 0; i < len(this.Units); i++ {
		l.Units[i] = this.Units[i].Clone()
	}
	return
}

func (this IOLayer) Set(deep int, x, y int, v DataType) {
	this.Units[deep].Set(x, y, v)
}

func (this IOLayer) Add(deep, x, y int, v DataType) {
	this.Units[deep].Add(x, y, v)
}

func (this IOLayer) EachUnit(handler func(deep int, unit Matrix)) {
	for i := 0; i < len(this.Units); i++ {
		handler(i, this.Units[i])
	}
}

type ReLU struct{}

func (this ReLU) ComputeMatrix(in Matrix) (out Matrix) {
	out = NewMatrix(in.Width(), in.Height())
	in.Each(func(x, y int, v DataType) {
		if v > 0 {
			out.Set(x, y, v)
		} else {
			out.Set(x, y, 0)
		}
	})
	return
}

func (this ReLU) Compute(in IOLayer) (out IOLayer) {
	out = in.Clone()
	in.EachUnit(func(deep int, unit Matrix) {
		unit.Each(func(x, y int, v DataType) {
			if v > 0 {
				out.Set(deep, x, y, v)
			} else {
				out.Set(deep, x, y, 0)
			}
		})
	})
	return
}

type MaxPool struct {
	//Kernel width
	W int
}

func NewDefaultMaxPool() (p MaxPool) {
	return MaxPool{W: 2}
}

func (this MaxPool) Compute(in IOLayer) (out IOLayer) {
	out = NewIOLayer(in.Deep(), in.Width()>>1, in.Height()>>1)
	in.Walk(this.W, this.W, this.W, 0, func(deep, inLeft, inTop int, crop Matrix) {
		out.Set(deep, inLeft>>1, inTop>>1, crop.Max())
	})
	return
}

type AveragePool struct {
	W int
}

func NewDefaultAveragePool() (p AveragePool) {
	return AveragePool{W: 2}
}

func (this AveragePool) Compute(in IOLayer) (out IOLayer) {
	out = NewIOLayer(in.Deep(), in.Width()>>1, in.Height()>>1)
	in.Walk(this.W, this.W, this.W, 0, func(deep, inLeft, inTop int, crop Matrix) {
		out.Set(deep, inLeft>>1, inTop>>1, crop.Max())
	})
	return
}

type MinPool struct {
	W int
}

func (this MinPool) Compute(in IOLayer) (out IOLayer) {
	out = NewIOLayer(in.Deep(), in.Width()>>1, in.Height()>>1)
	in.Walk(this.W, this.W, this.W, 0, func(deep, inLeft, inTop int, crop Matrix) {
		out.Set(deep, inLeft>>1, inTop>>1, crop.Max())
	})
	return
}

type TransformLayer struct {
	fun func(DataType) DataType
}

func (this TransformLayer) ComputeMatrix(in Matrix) (out Matrix) {
	out = NewMatrix(in.Width(), in.Height())
	in.Each(func(x, y int, v DataType) {
		out.Set(x, y, this.fun(v))
	})
	return
}

func (this TransformLayer) Compute(in IOLayer) (out IOLayer) {
	out = NewIOLayer(in.Deep(), in.Width(), in.Height())
	in.EachUnit(func(deep int, unit Matrix) {
		out.Units[deep] = this.ComputeMatrix(unit)
	})
	return
}

type ConvLayer struct {
	filters []Filter
}

// num: the number of filter
// deep: deep for each filter
func NewDefaultConvLayer(num, deep, w, h int) (c ConvLayer) {
	c.filters = make([]Filter, num)
	for i := 0; i < num; i++ {
		c.filters[i] = NewFilter(deep, w, h)
		for j := 0; j < deep; j++ {
			c.filters[i].Units[j] = CreateGaussianMatrix(w, h)
			c.filters[i].Bias = RandValue()
		}
	}
	return
}

func (this ConvLayer) Deep() int {
	return len(this.filters)
}

func (this ConvLayer) Compute(in IOLayer, padding, stepWidth int) (out IOLayer) {
	ms := []Matrix{}
	w := 0
	h := 0
	for d := 0; d < len(this.filters); d++ {
		m := this.filters[d].Compute(in, padding, stepWidth)
		ms = append(ms, m)
		w = m.Width()
		h = m.Height()
	}
	out = IOLayer{W: w, H: h, Units: ms}
	return
}
