package gat

type IGenome interface {
	CreateIndividual() Individual
	Crossover(i1, i2 Individual) (Individual, Individual)
}
