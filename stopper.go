package gat

type IStopper interface {
	IsTerminated(iterations int, population []Individual) bool
}

type IterationStopper struct {
	iterations int
}

func NewIterationStopper(iterations int) IStopper {
	return &IterationStopper{
		iterations: iterations,
	}
}

func (i IterationStopper) IsTerminated(iterations int, _ []Individual) bool {
	return iterations >= i.iterations
}
