package cv

import (
	"fmt"
	"math"
)

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

func LogisticFunc(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func ConvertArrayToMatrix(w, h int, a []DataType) (m Matrix) {
	m = Matrix{W: w, H: h, Data: a}
	return
}
