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
	fc     ConvLayer
	output IOLayer
	e      error
)

func Train(fileName ...string) {
	for _, f := range fileName {
		input, e = LoadImage(f)
		if e != nil {
			log(e)
			continue
		}

	}
}
