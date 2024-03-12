package tuner

import (
	"math"
	"math/rand"
)

func InitUniform(rnd *rand.Rand, data []float64, variance float64) {
	var uniformVariance = 1.0 / 12
	var scale = math.Sqrt(variance / uniformVariance)
	for i := range data {
		data[i] = (rnd.Float64() - 0.5) * scale
	}
}

func ReverseSigmoid(x float64) float64 {
	return -math.Log(1/x - 1)
}
