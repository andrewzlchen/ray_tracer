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
