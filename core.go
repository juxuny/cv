package cv

var (
	input  IOLayer
	c1     ConvLayer
	tran   TransformLayer = TransformLayer{fun: LogisticFunc}
	o1     IOLayer
	p1     AveragePool
	c2     ConvLayer
	o2     IOLayer
	p2     AveragePool
	fc1    ConvLayer
	fco1   IOLayer
	fc2    ConvLayer
	fco2   IOLayer
	fc3    ConvLayer
	fco3   IOLayer
	output IOLayer
	e      error
)

func init() {
	c1 = NewDefaultConvLayer(6, 3, 5, 5)
	p1 = NewDefaultAveragePool()
	c2 = NewDefaultConvLayer(16, 6, 5, 5)
	p2 = NewDefaultAveragePool()

	fc1 = NewDefaultConvLayer(160, 16, 5, 5)
	fc2 = NewDefaultConvLayer(84, 160, 1, 1)
	fc3 = NewDefaultConvLayer(10, 84, 1, 1)
}

func Train(fileName string, realResult Array) {
	input, e = LoadImage(fileName)
	if e != nil {
		log(e)
		return
	}
	input = Standardization(input)
	Test()
}

func Test() {
	o1 = c1.Compute(input, 0, 1)
	o1 = tran.Compute(o1)
	o1 = p1.Compute(o1)

	o2 = c2.Compute(o1, 0, 1)
	o2 = tran.Compute(o2)
	o2 = p2.Compute(o2)

	fco1 = fc1.Compute(o2, 0, 1)
	fco1 = tran.Compute(fco1)

	fco2 = fc2.Compute(fco1, 0, 1)
	fco2 = tran.Compute(fco2)

	fco3 = fc3.Compute(fco2, 0, 1)
	fco3 = tran.Compute(fco3)

	log(fco3)
}
