package gat

type IMuter interface {
	Mutate(individual Individual) Individual
}
