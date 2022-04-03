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

type FitnessStopper struct {
	fitness float64
}

func NewFitnessStopper(fitness float64) IStopper {
	return &FitnessStopper{
		fitness: fitness,
	}
}

func (i FitnessStopper) IsTerminated(_ int, population []Individual) bool {
	for _, ind := range population {
		if ind.GetFitness() >= i.fitness {
			return true
		}
	}

	return false
}
