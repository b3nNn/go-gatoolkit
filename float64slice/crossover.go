package float64slice

import (
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

type IFloat64SliceCrosser interface {
	Cross(g1, g2 Gene) (Gene, Gene)
}

type BlendCrossover struct {
	alpha float64
}

func NewBlendCrossover(alpha float64) BlendCrossover {
	return BlendCrossover{
		alpha: alpha,
	}
}

func (c BlendCrossover) Cross(g1, g2 Gene) (Gene, Gene) {
	if len(g1.Val()) != len(g2.Val()) {
		panic("invalid gene")
	}

	dist := distuv.Uniform{
		Min: 0,
		Max: 1,
	}
	g1Cross := make([]float64, len(g1.Val()))
	copy(g1Cross, g1.Val())
	g2Cross := make([]float64, len(g2.Val()))
	copy(g2Cross, g2.Val())
	r := []float64{dist.Rand(), dist.Rand()}
	for i, _ := range g1.Val() {
		var low = math.Min(g1Cross[i], g2Cross[i]) - c.alpha*math.Abs(g2Cross[i]-g1Cross[i])
		var high = math.Max(g1Cross[i], g2Cross[i]) + c.alpha*math.Abs(g2Cross[i]-g1Cross[i])
		cross1 := low + r[0]*(high-low)
		cross2 := low + r[1]*(high-low)
		g1Cross[i] = cross1
		g2Cross[i] = cross2
	}

	var g1Offspring = NewGene(g1Cross)
	var g2Offspring = NewGene(g2Cross)

	return g1Offspring, g2Offspring
}
