package float64slice

import (
	"errors"
	gat "github.com/b3nNn/go-gatoolkit"
	"gonum.org/v1/gonum/stat/distuv"
)

var ()

type RandomDeviationMuter struct {
	deviations int
	mu         float64
	sigma      float64
}

func NewRandomDeviationMuter(deviations int, mu float64, sigma float64) *RandomDeviationMuter {
	return &RandomDeviationMuter{
		deviations: deviations,
		mu:         mu,
		sigma:      sigma,
	}
}

func (r *RandomDeviationMuter) Mutate(individual gat.Individual) gat.Individual {
	ind, ok := individual.(*Individual)
	if !ok {
		panic(errors.New("invalid individual"))
	}

	gaussDist := distuv.Normal{Mu: r.mu, Sigma: r.sigma}
	mut := make([]float64, len(ind.GetGene().Val()))

	copy(mut, ind.GetGene().Val())

	idx := make(map[int]bool, r.deviations)

	dist := distuv.Uniform{Min: 0, Max: float64(len(mut))}
	for len(idx) < r.deviations {
		i := int(dist.Rand())

		if _, ok := idx[i]; !ok {
			idx[i] = true
		}
	}

	for i, _ := range idx {
		ra := gaussDist.Rand()
		mut[i] = mut[i] + ra
	}

	gene := NewGene(mut)

	return &Individual{
		Gene: gene,
	}
}
