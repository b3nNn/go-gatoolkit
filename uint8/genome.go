package uint8

import (
	"errors"
	gat "github.com/b3nNn/go-gatoolkit"
	"gonum.org/v1/gonum/stat/distuv"
)

var dist distuv.Uniform

func init() {
	dist = distuv.Uniform{
		Min: 0,
		Max: 255,
	}
}

type Genome struct {
	Crosser IUInt8Crosser
}

func (Genome) CreateIndividual() gat.Individual {
	return &Individual{
		Gene:    NewGene(uint8(dist.Rand())),
		fitness: 0,
	}
}

func (u Genome) Crossover(i1, i2 gat.Individual) (gat.Individual, gat.Individual) {
	ind1, ok := i1.(*Individual)
	if !ok {
		panic(errors.New("invalid individual"))
	}

	ind2, ok := i2.(*Individual)
	if !ok {
		panic(errors.New("invalid individual"))
	}

	g1Cross, g2Corss := u.Crosser.Cross(ind1.Gene, ind2.Gene)

	return &Individual{
			Gene: g1Cross,
		}, &Individual{
			Gene: g2Corss,
		}
}
