package float64slice

import (
	"errors"
	gat "github.com/b3nNn/go-gatoolkit"
	"gonum.org/v1/gonum/stat/distuv"
)

var (
	gaussDist distuv.Normal
)

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

	dist := distuv.Normal{Mu: r.mu, Sigma: r.sigma}
	mut := make([]float64, len(ind.GetGene().Val()))

	copy(mut, ind.GetGene().Val())
	ra := dist.Rand()
	for i := 0; i < len(mut); i++ {
		mut[i] = mut[i] + ra
	}

	gene := NewGene(mut)

	return &Individual{
		Gene: gene,
	}
}
