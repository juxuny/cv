package cv

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

func (this IOLayer) Walk(w, h, stepWidth, padding int, f func(deep, inLeft, inTop int, crop Matrix)) {
	this.EachUnit(func(deep int, unit Matrix) {
		p := unit.PaddingWith(padding, 0)
		for i := 0; i+w < p.W; i += stepWidth {
			for j := 0; j+h < p.H; j += stepWidth {
				crop := p.Crop(i, j, i+w, j+h)
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

func (this IOLayer) EachUnit(handler func(deep int, unit Matrix)) {
	for i := 0; i < len(this.Units); i++ {
		handler(i, this.Units[i])
	}
}

type WeightLayer struct {
	W, H  int
	Units []Matrix
	Bias  DataType
}

type ReLU struct {
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
