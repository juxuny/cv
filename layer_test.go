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
