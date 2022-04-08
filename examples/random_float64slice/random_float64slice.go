package main

import (
	"fmt"
	gat "github.com/b3nNn/go-gatoolkit"
	gatfloat64slice "github.com/b3nNn/go-gatoolkit/float64slice"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
	"math/rand"
	"time"
)

const (
	genLen int = 5
)

var (
	targets []float64
	dist    distuv.Uniform
)

func init() {
	dist = distuv.Uniform{
		Min: gatfloat64slice.SmallestNonzeroFloat64,
		Max: gatfloat64slice.MaxFloat64,
	}
	rand.Seed(time.Now().UnixNano())
	targets = make([]float64, genLen)
	for i := 0; i < genLen; i++ {
		targets[i] = dist.Rand()
	}
}

func Eval(individual gat.Individual) float64 {
	ind, ok := individual.(*gatfloat64slice.Individual)
	if !ok {
		panic("invalid individual")
	}

	var fitness float64

	for i, target := range targets {
		v := ind.GetGene().Val()[i]
		diff := math.Abs(v - target)
		fitness -= diff
	}
	return fitness
}

func main() {
	ga := gat.NewGeneticAlgorithm()
	ga.Configure(100, 0.6, 0.1, 4)
	ga.WithGenome(gatfloat64slice.NewGenome(genLen, gatfloat64slice.MaxFloat64, gatfloat64slice.SmallestNonzeroFloat64, gatfloat64slice.NewBlendCrossover(0.6)))
	ga.WithFitness(Eval)
	ga.WithMuter(gatfloat64slice.NewRandomDeviationMuter(genLen, 0, 1))
	ga.WithSelection(gat.NewRankSelection(1))
	ga.WithStopper(gat.NewFitnessStopper(-1))

	result := ga.Simulate()

	best, ok := result.Best.(*gatfloat64slice.Individual)
	if !ok {
		panic("invalid individual type")
	}

	fmt.Printf("Found %v after %d generation(s).\n", best.Gene.Val(), result.Iterations)

}
