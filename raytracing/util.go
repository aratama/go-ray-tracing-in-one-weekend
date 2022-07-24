package raytracing

import (
	"math"
	"math/rand"
)

func clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func randomMinMax(min float64, max float64, random *rand.Rand) float64 {
	return min + (max-min)*random.Float64()
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}
