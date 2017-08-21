package cv

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
	this.Data[r*this.W+c] = v
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
	out = NewMatrix(this.W+width, this.H+width)
	out.Each(func(x, y int, v DataType) {
		if x >= width && x < out.W-width && y >= width && y < out.H-width {
			out.Set(x, y, this.Get(x-width, y-width))
		} else {
			out.Set(x, y, v)
		}
	})
	return
}
