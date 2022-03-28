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

	fmt.Printf("target: %d\n", target)
}

type Fitness struct {
}

func (f Fitness) Eval(individual gat.Individual) float64 {
	ind, ok := individual.(*gatuint8.Individual)
	if !ok {
		panic("invalid individual")
	}

	return float64(target - ind.Gene.Val())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ga := gat.NewGeneticAlgorithm()
	ga.Configure(5)
	ga.Genome = gatuint8.Genome{
		Crosser: gatuint8.NewBlendCrossover(0.3),
	}
	ga.Fitness = Fitness{}
	ga.Selection = &gatuint8.RankSelection{}

	ga.Simulate()
}
