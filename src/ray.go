package raytracer

import (
	"fmt"
)

// Ray is a struc that contains a origin and a direction and can be described the formula
// P(t) = A + tb where
// - P represents the position of a point in 3D space
// - t is a real number that scales the point along the direction vector
// - b is a vector that represents the direction that the ray is facing
type Ray struct {
	// Origin is the location of a point in 3 dimensional space, denoted by a Vec3
	origin *Vec3
	// Direction is a Vec3 that represents the direction of the ray
	direction *Vec3
}

// NewRay constructs a new ray from an origin and direction vector
func NewRay(origin, direction *Vec3) *Ray {
	return &Ray{
		origin:    origin,
		direction: direction,
	}
}

// Origin returns a vector representing the origin of the ray
func (r *Ray) Origin() *Vec3 {
	return r.origin
}

// Direction returns a vector representing the direction of the ray
func (r *Ray) Direction() *Vec3 {
	return r.direction
}

// At returns the point on the ray given the coefficient, t
func (r *Ray) At(t int64) *Vec3 {
	return r.Origin().
		AddVector(
			r.Direction().
				MultiplyFloat(float64(t)),
		)
}

// Color computes the color of the ray.
// This function blends color to be a linear gradient on the Y direction.
// We first get the unit vector, of which, the values will be -1 < x < 1.
// We then add 1 and divide by 2 to scale the values to be from 0 to 1.
// 0 = white
// 1 = blue
func (r *Ray) Color() (*Vec3, error) {
	unitDirection, err := r.Direction().Unit()
	if err != nil {
		return nil, fmt.Errorf("cannot get unit vector of vector: %s", err)
	}

	// scale unit vector values to be 0 < x < 1 to determine what color from white to blue to choose
	blueness := 0.5 * (unitDirection.Y + 1.0)

	blue := NewVec3(1.0, 1.0, 1.0)
	shadeOfBlue := blue.MultiplyFloat(1.0 - blueness)

	// This is similar to a seed in a RNG to determine the base level of blueness
	baseBlue := NewVec3(0.5, 0.7, 1.0).MultiplyFloat(blueness)

	return shadeOfBlue.AddVector(baseBlue), nil
}
