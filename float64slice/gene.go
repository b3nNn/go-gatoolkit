package float64slice

type Gene struct {
	values []float64
}

func NewGene(values []float64) Gene {
	return Gene{values: values[:]}
}

func (g Gene) Val() []float64 {
	return g.values
}
