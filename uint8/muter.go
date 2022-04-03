package uint8

import (
	"errors"
	gat "github.com/b3nNn/go-gatoolkit"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

var (
	gaussDist distuv.Normal
)

func init() {
	gaussDist = distuv.Normal{Mu: 0, Sigma: 1}
}

type RandomDeviationMuter struct {
	mu    float64
	sigma float64
}

func NewRandomDeviationMuter(mu float64, sigma float64) *RandomDeviationMuter {
	return &RandomDeviationMuter{
		mu:    mu,
		sigma: sigma,
	}
}

func (r *RandomDeviationMuter) Mutate(individual gat.Individual) gat.Individual {
	ind, ok := individual.(*Individual)
	if !ok {
		panic(errors.New("invalid individual"))
	}

	mut := math.Max(0, math.Min(255, float64(ind.GetGene().Val())+gaussDist.Rand()))

	return &Individual{
		Gene: NewGene(uint8(mut)),
	}
}
