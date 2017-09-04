package cv

import (
	"math"
)

var (
	input IOLayer
	//6@5x5x3
	c1   ConvLayer
	tran TransformLayer = TransformLayer{fun: LogisticFunc}
	o1   IOLayer
	p1   AveragePool
	//16@5x5x6
	c2 ConvLayer
	o2 IOLayer
	p2 AveragePool
	//10@5x5x16
	fc2  ConvLayer
	fco2 IOLayer
	e    error
)

func init() {
	c1 = NewDefaultConvLayer(6, 3, 5, 5)
	p1 = NewDefaultAveragePool()
	c2 = NewDefaultConvLayer(16, 6, 5, 5)
	p2 = NewDefaultAveragePool()
	fc2 = NewDefaultConvLayer(10, 16, 5, 5)
}

func Loss(out, realResult Array) (loss DataType) {
	if len(out) != len(realResult) {
		panic("can't handle two array without same length")
	}
	for i := 0; i < len(out); i++ {
		loss += DataType(math.Pow(float64(out[i]-realResult[i]), 2))
	}
	loss /= 2
	return
}

func Train(fileName string, realResult Array) {
	input, e = LoadImage(fileName)
	if e != nil {
		log(e)
		return
	}
	input = Standardization(input)
	Test()
	log("Loss: ", Loss(IOLayerToArray(fco2), realResult))
	UpdateWeight(realResult)
	Test()
	log("Loss: ", Loss(IOLayerToArray(fco2), realResult))
}

func Test() {
	o1 = c1.Compute(input, 0, 1)
	o1 = tran.Compute(o1)
	o1 = p1.Compute(o1)

	o2 = c2.Compute(o1, 0, 1)
	o2 = tran.Compute(o2)
	o2 = p2.Compute(o2)

	//	fco1 = fc1.Compute(o2, 0, 1)
	//	fco1 = tran.Compute(fco1)

	fco2 = fc2.Compute(o2, 0, 1)
	fco2 = tran.Compute(fco2)

	//	log(fco3)
}

func UpdateWeight(realResult Array) {
	//TODO
}
