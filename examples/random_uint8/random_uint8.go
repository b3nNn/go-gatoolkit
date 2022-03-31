package main

import (
	"fmt"
	gat "github.com/b3nNn/go-gatoolkit"
	gatuint8 "github.com/b3nNn/go-gatoolkit/uint8"
	"math"
	"math/rand"
	"time"
)

var (
	target uint8
)

func init() {
	rand.Seed(time.Now().UnixNano())
	target = uint8(rand.Intn(256))

	fmt.Printf("target: %d\n", target)
}

type Fitness struct {
}

func (f Fitness) Eval(individual gat.Individual) float64 {
	ind, ok := individual.(*gatuint8.Individual)
	if !ok {
		panic("invalid individual")
	}

	return 255. - math.Abs(float64(ind.Gene.Val()-target))
}

func main() {
	ga := gat.NewGeneticAlgorithm()
	ga.Configure(20, 0.8, 0.3)
	ga.Genome = gatuint8.Genome{
		Crosser: gatuint8.NewBlendCrossover(0.3),
	}
	ga.Fitness = Fitness{}
	ga.Muter = gatuint8.NewRandomDeviationMuter(0, 1)
	ga.Selection = gat.NewRankSelection(2)
	ga.Stopper = gat.NewIterationStopper(5)

	population := ga.Simulate()

	for i, ind := range population {
		fmt.Printf("Selection #%d\t%.2f\n", i, ind.GetFitness())
	}
}
