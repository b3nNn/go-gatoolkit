package uint8

type Gene struct {
	val uint8
}

func NewGene(val uint8) Gene {
	return Gene{val: val}
}

func (g Gene) Val() uint8 {
	return g.val
}
