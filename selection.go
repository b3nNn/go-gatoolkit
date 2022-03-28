package gat

type ISelection interface {
	Select(population []Individual) []Individual
}
