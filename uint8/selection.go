package uint8

import (
	gat "github.com/b3nNn/go-gatoolkit"
	"sort"
)

type RankSelection struct {
}

func (s *RankSelection) Select(population []gat.Individual) []gat.Individual {
	var pop = population[:]
	//var selection = make([]gat.Individual, 0)

	sort.SliceStable(pop, func(i, j int) bool {
		return pop[i].GetFitness() > pop[j].GetFitness()
	})

	return pop
}
