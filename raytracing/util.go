package raytracing

import "math/rand"

func clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func random(min float64, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
