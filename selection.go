package gat

import (
	"gonum.org/v1/gonum/stat/distuv"
	"sort"
)

type ISelection interface {
	Select(population []Individual) []Individual
}

type RankSelection struct {
	eliteSize int
}

func NewRankSelection(eliteSize int) *RankSelection {
	return &RankSelection{
		eliteSize: eliteSize,
	}
}

func (s *RankSelection) Select(population []Individual) []Individual {
	var ranksSum float64
	pop := make([]Individual, len(population))
	copy(pop, population)
	rankDistance := 1. / float64(len(pop))
	ranks := make([]float64, 0)

	for i := 0; i < len(pop); i++ {
		var rank = 1. - float64(i)*rankDistance
		ranks = append(ranks, rank)
		ranksSum += rank
	}

	sort.SliceStable(pop, func(i, j int) bool {
		return pop[i].GetFitness() > pop[j].GetFitness()
	})

	selection := make([]Individual, 0)
	selection = append(selection, pop[:s.eliteSize]...)
	limit := len(pop) - s.eliteSize
	dist := distuv.Uniform{
		Min: 0,
		Max: 1,
	}
	for i := 0; i < limit; i++ {
		var sum float64
		rest := dist.Rand() * ranksSum

		for ii := 0; ii < len(pop); ii++ {
			sum += ranks[ii]

			if sum > rest {
				selection = append(selection, pop[ii])
				break
			}
		}
	}

	return selection
}
