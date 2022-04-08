package main

import (
	"fmt"
	gat "github.com/b3nNn/go-gatoolkit"
	gatuint8 "github.com/b3nNn/go-gatoolkit/uint8"
	"math/rand"
	"time"
)

var (
	target uint8
)

func init() {
	rand.Seed(time.Now().UnixNano())
	target = uint8(rand.Intn(256))
}

func Eval(individual gat.Individual) float64 {
	ind, ok := individual.(*gatuint8.Individual)
	if !ok {
		panic("invalid individual")
	}

	if ind.Gene.Val() <= target {
		return 255. - float64(target-ind.Gene.Val())
	}

	return 255. - float64(ind.Gene.Val()-target)
}

func main() {
	ga := gat.NewGeneticAlgorithm()
	ga.Configure(10, 0.2, 0.8, 4)
	ga.WithGenome(gatuint8.Genome{
		Crosser: gatuint8.NewBlendCrossover(0.3),
	})
	ga.WithFitness(Eval)
	ga.WithMuter(gatuint8.NewRandomDeviationMuter(0, 1))
	ga.WithSelection(gat.NewRankSelection(2))
	ga.WithStopper(gat.NewFitnessStopper(255))

	result := ga.Simulate()

	best, ok := result.Best.(*gatuint8.Individual)
	if !ok {
		panic("invalid individual type")
	}

	fmt.Printf("Found %d after %d generation(s).\n", best.Gene.Val(), result.Iterations)
}
