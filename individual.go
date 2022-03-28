package gat

type Individual interface {
	GetFitness() float64
	SetFitness(fitness float64)
}
