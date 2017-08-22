package cv

type Filter struct {
	W, H  int
	Units []Matrix
	Bias  DataType
}

func NewFilter(deep, w, h int) (r Filter) {
	r.W = w
	r.H = h
	r.Units = make([]Matrix, deep)
	for i := 0; i < len(r.Units); i++ {
		r.Units[i] = NewMatrix(w, h)
	}
	return
}

func (this Filter) Deep() int {
	return len(this.Units)
}

func (this Filter) Width() int {
	return this.W
}

func (this Filter) Height() int {
	return this.H
}

func (this Filter) Compute(in IOLayer, padding, stepWidth int) (out Matrix) {
	out = NewMatrix((in.Width()-this.Width()+2*padding)/stepWidth+1, (in.Height()-this.Height()+2*padding)/stepWidth+1)
	in.Walk(this.Width(), this.Height(), stepWidth, padding, func(deep, inLeft, inTop int, unit Matrix) {
		out.Add(inLeft/stepWidth, inTop/stepWidth, unit.DotProduct(this.Units[deep]))
	})
	out.AddAll(this.Bias)
	return
}
