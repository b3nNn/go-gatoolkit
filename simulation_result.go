package gat

type SimulationResult struct {
	Best       Individual
	Iterations int
	Population []Individual
}

func NewSimulationResult() *SimulationResult {
	return &SimulationResult{
		Best:       nil,
		Iterations: 0,
		Population: []Individual{},
	}
}
