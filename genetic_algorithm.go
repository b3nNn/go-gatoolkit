package gat

import "fmt"

type GeneticAlgorithm struct {
	Genome         IGenome
	Fitness        IFitness
	Selection      ISelection
	Population     []Individual
	populationSize int
}

func NewGeneticAlgorithm() *GeneticAlgorithm {
	return &GeneticAlgorithm{
		Genome:         nil,
		Fitness:        nil,
		Selection:      nil,
		Population:     []Individual{},
		populationSize: 10,
	}
}

func (g *GeneticAlgorithm) Configure(populationSize int) {
	g.populationSize = populationSize
}

func (g *GeneticAlgorithm) Simulate() {
	var population = g.Population[:]

	for len(population) < g.populationSize {
		population = append(population, g.Genome.CreateIndividual())
	}

	for i, ind := range population {
		fit := g.Fitness.Eval(ind)
		ind.SetFitness(fit)
		fmt.Printf("Population #%d\t%.2f\n", i, ind.GetFitness())
	}

	selection := g.Selection.Select(population)

	for i, ind := range selection {
		fmt.Printf("Selection #%d\t%.2f\n", i, ind.GetFitness())
	}
}
