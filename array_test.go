package cv

import (
	"testing"
)

func TestConvolve(t *testing.T) {
	a := NewArray(3, 3, 4)
	b := NewArray(3, 2, 1)
	t.Log(Convolve(a, b))
}
