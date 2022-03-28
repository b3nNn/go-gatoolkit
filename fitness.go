package gat

type IFitness interface {
	Eval(individual Individual) float64
}
