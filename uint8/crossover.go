package uint8

import (
	"math"
	"math/rand"
)

type IUInt8Crosser interface {
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
	var g1f = float64(g1.Val())
	var g2f = float64(g2.Val())
	var low = math.Min(g1f, g2f) - c.alpha*math.Abs(g2f-g1f)
	var high = math.Max(g1f, g2f) + c.alpha*math.Abs(g2f-g1f)
	var g1Cross = low + rand.Float64()*(high-low)
	var g2Cross = low + rand.Float64()*(high-low)
	var g1Offspring = NewGene(uint8(math.Max(0, math.Min(255, g1Cross))))
	var g2Offspring = NewGene(uint8(math.Max(0, math.Min(255, g2Cross))))

	return g1Offspring, g2Offspring
}
