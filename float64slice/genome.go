package float64slice

import (
	"errors"
	gat "github.com/b3nNn/go-gatoolkit"
	"gonum.org/v1/gonum/stat/distuv"
)

const (
	MaxFloat64             = 9.99e+10 // 2**1023 * (2**53 - 1) / 2**52
	SmallestNonzeroFloat64 = 9.99e-10 // 1 / 2**(1023 - 1 + 52)
)

var dist distuv.Uniform

func init() {
	dist = distuv.Uniform{
		Min: SmallestNonzeroFloat64,
		Max: MaxFloat64,
	}
}

type Genome struct {
	geneLen int
	Crosser IFloat64SliceCrosser
}

func NewGenome(geneLen int, crosser IFloat64SliceCrosser) *Genome {
	return &Genome{
		geneLen: geneLen,
		Crosser: crosser,
	}
}

func (u Genome) CreateIndividual() gat.Individual {
	g := make([]float64, u.geneLen)

	for i := 0; i < u.geneLen; i++ {
		g[i] = dist.Rand()
	}

	return &Individual{
		Gene:    NewGene(g),
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
