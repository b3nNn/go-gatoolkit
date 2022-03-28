package uint8

type Individual struct {
	Gene    Gene
	fitness float64
}

func (i *Individual) GetGene() Gene {
	return i.Gene
}

func (i *Individual) GetFitness() float64 {
	return i.fitness
}

func (i *Individual) SetFitness(fitness float64) {
	i.fitness = fitness
}
