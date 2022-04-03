package gat

import (
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
)

type GeneticAlgorithm struct {
	Genome               IGenome
	Fitness              Fitness
	Muter                IMuter
	Selection            ISelection
	Stopper              IStopper
	Population           []Individual
	populationSize       int
	crossoverProbability float64
	mutationProbability  float64
}

func NewGeneticAlgorithm() *GeneticAlgorithm {
	return &GeneticAlgorithm{
		Genome:               nil,
		Fitness:              nil,
		Muter:                nil,
		Selection:            nil,
		Stopper:              nil,
		Population:           []Individual{},
		populationSize:       0,
		crossoverProbability: 0,
		mutationProbability:  0,
	}
}

func (g *GeneticAlgorithm) Configure(populationSize int, crossoverProbability float64, mutationProbability float64) {
	g.populationSize = populationSize
	g.crossoverProbability = crossoverProbability
	g.mutationProbability = mutationProbability
}

func (g *GeneticAlgorithm) Simulate() *SimulationResult {
	var mustStop bool
	var it int
	population := make([]Individual, len(g.Population))
	copy(population, g.Population)

	for len(population) < g.populationSize {
		population = append(population, g.Genome.CreateIndividual())
	}

	for !mustStop {
		offsprings := make([]Individual, 0)

		if it > 0 {
			for i := 0; i < len(population); {
				if i+1 >= len(population) {
					offsprings = append(offsprings, population[i])
					break
				}

				r := rand.Float64()
				if r < g.crossoverProbability {
					o1, o2 := g.Genome.Crossover(population[i], population[i+1])
					offsprings = append(offsprings, o1, o2)
					i += 2
				} else {
					offsprings = append(offsprings, population[i])
					i++
				}
			}

			population = make([]Individual, 0)
			dist := distuv.Uniform{
				Min: 0,
				Max: 1,
			}
			for i := 0; i < len(offsprings); i++ {
				if dist.Rand() < g.mutationProbability {
					mut := g.Muter.Mutate(offsprings[i])
					population = append(population, mut)
				} else {
					population = append(population, offsprings[i])
				}
			}
		}

		for _, ind := range population {
			fit := g.Fitness(ind)
			ind.SetFitness(fit)
		}

		population = g.Selection.Select(population)
		it++

		if g.Stopper != nil {
			mustStop = g.Stopper.IsTerminated(it, population)
		}
	}

	result := NewSimulationResult()
	result.Population = population[:]
	result.Iterations = it
	result.Best = result.Population[0]

	return result
}

func (g *GeneticAlgorithm) WithGenome(genome IGenome) *GeneticAlgorithm {
	g.Genome = genome

	return g
}

func (g *GeneticAlgorithm) WithFitness(fitness Fitness) *GeneticAlgorithm {
	g.Fitness = fitness

	return g
}

func (g *GeneticAlgorithm) WithMuter(muter IMuter) *GeneticAlgorithm {
	g.Muter = muter

	return g
}

func (g *GeneticAlgorithm) WithSelection(selection ISelection) *GeneticAlgorithm {
	g.Selection = selection

	return g
}

func (g *GeneticAlgorithm) WithStopper(stopper IStopper) *GeneticAlgorithm {
	g.Stopper = stopper

	return g
}
