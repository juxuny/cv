package cv

type Array []DataType

func NewArray(d ...DataType) (ret Array) {
	for i := 0; i < len(d); i++ {
		ret = append(ret, d[i])
	}
	return
}

func (this Array) Revert() (ret Array) {
	n := len(this)
	for i := 0; i < len(this); i++ {
		ret = append(ret, this[n-i-1])
	}
	return
}

func MatrixToArray(m Matrix) (ret Array) {
	for j := 0; j < m.Height(); j++ {
		for i := 0; i < m.Width(); i++ {
			ret = append(ret, m.Get(i, j))
		}
	}
	return
}

func Convolve(a, b Array) (ret DataType) {
	if len(a) != len(b) {
		panic("Convolve function can't compute convoluation between two Array without the same length")
	}
	n := len(a)
	for i := 0; i < len(a); i++ {
		ret += a[i] * b[n-i-1]
	}
	return
}
