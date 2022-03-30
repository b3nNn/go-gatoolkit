package gat

import (
	"fmt"
	"math/rand"
)

type GeneticAlgorithm struct {
	Genome               IGenome
	Fitness              IFitness
	Selection            ISelection
	Stopper              IStopper
	Population           []Individual
	populationSize       int
	crossoverProbability float64
}

func NewGeneticAlgorithm() *GeneticAlgorithm {
	return &GeneticAlgorithm{
		Genome:               nil,
		Fitness:              nil,
		Selection:            nil,
		Stopper:              nil,
		Population:           []Individual{},
		populationSize:       10,
		crossoverProbability: 0.5,
	}
}

func (g *GeneticAlgorithm) Configure(populationSize int, crossoverProbability float64) {
	g.populationSize = populationSize
	g.crossoverProbability = crossoverProbability
}

func (g *GeneticAlgorithm) Simulate() []Individual {
	var mustStop bool
	var it int
	population := g.Population[:]

	for len(population) < g.populationSize {
		population = append(population, g.Genome.CreateIndividual())
	}

	for !mustStop {
		nextGeneration := make([]Individual, 0)

		if it == 0 {
			nextGeneration = append(nextGeneration, population...)
		} else {
			for i := 0; i < len(population); {
				if i+1 >= len(population) {
					nextGeneration = append(nextGeneration, population[i])
					break
				}

				cp := 1 - g.crossoverProbability
				r := rand.Float64()

				if r >= cp {
					o1, o2 := g.Genome.Crossover(population[i], population[i+1])
					nextGeneration = append(nextGeneration, o1, o2)
					i += 2
				} else {
					nextGeneration = append(nextGeneration, population[i])
					i++
				}
			}
		}

		for _, ind := range population {
			fit := g.Fitness.Eval(ind)
			ind.SetFitness(fit)
		}

		population = g.Selection.Select(nextGeneration)
		it++

		for i, ind := range population {
			fmt.Printf("Population #%d/%d\t%.2f\n", it, i, ind.GetFitness())
		}

		if g.Stopper != nil {
			mustStop = g.Stopper.IsTerminated(it, population)
		}
	}

	return population
}
