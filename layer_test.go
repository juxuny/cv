package cv

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestMaxPool(t *testing.T) {
	in := NewIOLayer(1, 4, 4)
	t.Log("new IOLayer")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			in.Set(0, i, j, DataType(rand.Float64()))
		}
	}
	m := NewDefaultMaxPool()
	t.Log(in)
	out := m.Compute(in)
	t.Log(out)
}

//Refer: http://cs231n.github.io/convolutional-networks/?from=singlemessage&isappinstalled=0
func TestFilter(t *testing.T) {
	in := NewIOLayer(3, 5, 5)
	in.Units = []Matrix{
		CreateMatrix(5, 5,
			1, 1, 1, 0, 0,
			1, 0, 0, 2, 2,
			1, 0, 2, 1, 2,
			1, 0, 1, 1, 1,
			1, 1, 2, 2, 1,
		),
		CreateMatrix(5, 5,
			2, 0, 0, 1, 1,
			0, 0, 1, 2, 2,
			2, 0, 2, 2, 2,
			0, 0, 1, 2, 2,
			0, 1, 2, 2, 0,
		),
		CreateMatrix(5, 5,
			1, 1, 0, 0, 2,
			0, 2, 2, 1, 0,
			2, 1, 0, 1, 1,
			2, 1, 2, 0, 2,
			2, 1, 2, 2, 1,
		),
	}
	w := NewFilter(3, 3, 3)
	w.Units = []Matrix{
		CreateMatrix(
			3, 3,
			0, 1, 1,
			1, 0, 0,
			1, 0, 0,
		),
		CreateMatrix(
			3, 3,
			1, 1, 0,
			0, -1, 0,
			1, 1, -1,
		),
		CreateMatrix(
			3, 3,
			1, -1, -1,
			1, 0, -1,
			0, -1, -1,
		),
	}
	w.Bias = 1
	out := w.Compute(in, 1, 2)
	t.Log(out)

	w.Units = []Matrix{
		CreateMatrix(
			3, 3,
			0, 0, 0,
			1, 1, -1,
			0, -1, 0,
		),
		CreateMatrix(
			3, 3,
			0, 1, 1,
			0, 0, -1,
			0, -1, 0,
		),
		CreateMatrix(
			3, 3,
			-1, 1, 0,
			-1, -1, -1,
			-1, 1, -1,
		),
	}
	w.Bias = 0
	out = w.Compute(in, 1, 2)
	t.Log(out)
	actLayer := TransformLayer{fun: LogisticFunc}
	out = actLayer.ComputeMatrix(out)
	t.Log(out)
	out = ReLU{}.ComputeMatrix(out)
	t.Log(out)
}

func TestTrain(t *testing.T) {
	Train("res/2/1.png", CreateResultArray(2))
}
