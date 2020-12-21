package raytracer

import (
	"math"
	"math/rand"
)

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func randomFloat(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

// clamp returns either returns x if min < x < max or min or max in order to return a value from [min, max]
func clamp(x, min, max float64) float64 {
	if x < min {
		return min
	} else if x > max {
		return max
	}
	return x
}
