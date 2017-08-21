package cv

type Connection []IntSet

func NewConnection(unitNum int) (c Connection) {
	c = make([]IntSet, unitNum)
	for i := 0; i < len(c); i++ {
		c[i] = NewIntSet()
	}
	return
}

func NewFullConnection(srcNum, dstNum int) (c Connection) {
	c = make([]IntSet, srcNum)
	for i := 0; i < len(c); i++ {
		c[i] = NewIntSet()
		for j := 0; j < dstNum; j++ {
			c[i].Add(j)
		}
	}
	return
}

func (this Connection) Add(from int, to ...int) {
	for i := 0; i < len(to); i++ {
		this[from].Add(to[i])
	}
}

func (this Connection) Each(f func(from int, to IntSet)) {
	for i := 0; i < len(this); i++ {
		f(i, this[i])
	}
}
